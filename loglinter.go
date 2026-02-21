package loglinter

import (
	"fmt"
	"go/ast"
	"slices"
	"strings"
	"unicode"

	"github.com/Komissarich/loglinter/config"
	_ "go.uber.org/zap"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)


func NewAnalyzer() *analysis.Analyzer {
	cfg, err := config.New()
	if err != nil {
		return nil
	}
	return &analysis.Analyzer{
			Name: "loglinter",
			Doc: "Checks that all logs uses english letters, starts with lowercase letter, doesnt have special symbols or emoji, doesnt have critical info",
			Run: func(p *analysis.Pass) (any, error) {
				return run(cfg, p)
			},
			Requires: []*analysis.Analyzer{inspect.Analyzer},
	}
}
func run(cfg *config.Config, pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}
	inspector.Preorder(nodeFilter, func(node ast.Node) {
		call := node.(*ast.CallExpr)
		selector, ok := call.Fun.(*ast.SelectorExpr)
		if !ok {
			return
		}
		base := getBase(selector.X)
		if base == nil {
			return
		}
		method := selector.Sel.Name
		if len(call.Args) == 0 {
        	return 
    	}
		if checkMethod(cfg, method) {
			checkLogMessage(cfg, pass, call)
		}

	})
	return nil, nil
}

func getBase(expr ast.Expr) *ast.Ident {
    for expr != nil {
        switch e := expr.(type) {
			case *ast.Ident:
				return e
			case *ast.SelectorExpr:
				expr = e.X
			case *ast.CallExpr:
				expr = e.Fun
			case *ast.ParenExpr:
				expr = e.X
			default:
				return nil
			}
    }
    return nil
}

func checkMethod(cfg *config.Config, method string) bool {
	if slices.Contains(cfg.AllowedMethods, method) {
		return true
	}
	return false
}

func checkLogMessage(cfg *config.Config, pass *analysis.Pass, call *ast.CallExpr) {
	problems := []string{}
	for _, arg := range call.Args {
		expr, ok := arg.(*ast.BinaryExpr)
		if ok {
			if cfg.Rules.CriticalInfoCheck {
				concatVars, concatStrings := getConcats(expr)
				fmt.Println("concatVars", concatVars, "concatStrings", concatStrings)
				for _, word := range concatVars {
					lower := strings.ToLower(word.Name)
					for _, keyword := range cfg.DangerousWords {
						if strings.Contains(lower, keyword) {
							problems = append(problems, fmt.Sprintf("log message should not contain critical information like %s", lower))
						} 
					}
				}
				for _, str := range concatStrings {
					if problem := performChecks(cfg, []rune(strings.Trim(str, "\"`"))); len(problem) != 0 {
						problems = append(problems, problem...)
					}
				}
			}
			break
		} else if expr, ok := arg.(*ast.BasicLit); ok {
			message := []rune(strings.Trim(expr.Value, "\"`"))
			if problem := performChecks(cfg, message); len(problem) != 0 {
				problems = append(problems, problem...)
			}
		}	
	
	}

		 fmt.Println(problems)		
	if len(problems) != 0 {
			pass.Report(
			analysis.Diagnostic{
				Pos:    call.Pos(),
				Message: strings.Join(problems, ";"),
		},)
	}
}

func getConcats(binaryExpr ast.Expr) ([]*ast.Ident, []string) {
	concatVariables := []*ast.Ident{}
	concatStrings := []string{}
	for binaryExpr != nil {
		switch elem := binaryExpr.(type) {
			case *ast.BinaryExpr:
				binaryExpr = elem.X
				ident, ok := elem.Y.(*ast.Ident)
				if ok {
					concatVariables = append(concatVariables, ident)
				} else {
					concatStrings = append(concatStrings, elem.Y.(*ast.BasicLit).Value)
				}
				
			case *ast.BasicLit:
				concatStrings = append(concatStrings, elem.Value)
				binaryExpr = nil
			default:
				fmt.Println("default value", binaryExpr)
				binaryExpr = nil
		}
	}
	return concatVariables, concatStrings
}

func performChecks(cfg *config.Config, message []rune) []string {
	fmt.Println("received string: ", string(message))
	result := []string{}
	if res := checkUpper(cfg, message); res != "" {
		result = append(result, fmt.Sprintf("log message '%s' should be named '%s'", string(message), strings.ToLower(string(message[0])) + string(message[1:])))
	}
	if res := checkCyrillic(cfg, message); res != "" {
		result = append(result,fmt.Sprintf("log message '%s' should not use cyrillic characters", string(message)))
	}
	if res := checkSpecial(cfg, message); res != "" {
		result = append(result,"log message should not use special symbols")
	}
	fmt.Println("found problems:", result)
	return result
}

func checkUpper(cfg *config.Config, message []rune) string {
	if cfg.Rules.UpperCaseCheck {
		if unicode.IsUpper(message[0]) {
			return fmt.Sprintf("log message '%s' should be named '%s'", string(message), strings.ToLower(string(message[0])) + string(message[1:]))
		}
	}
	return ""
}

func checkCyrillic(cfg *config.Config, message []rune) string {
	if cfg.Rules.UpperCaseCheck {
		for _, r := range message {
			if unicode.Is(unicode.Cyrillic, r) {
				return fmt.Sprintf("log message '%s' should not use cyrillic characters", string(message))
			}
		}
	}
	return ""
}

func checkSpecial(cfg *config.Config, message []rune) string {
	for _, r := range message {
		if (r != ' ') && (!((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r<= '9'))) && (!unicode.Is(unicode.Cyrillic, r)){
			return "log message should not use special symbols"
		}
	}
	return ""
}