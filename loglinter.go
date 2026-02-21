package loglinter

import (
	"fmt"
	"go/ast"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	_ "go.uber.org/zap"
)

func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "loglinter",
		Doc: "checks that all logs uses english letters, starts with lowercasw letter, doesnt have special symbols, or emoji, doesnt have critical info",
		Run: run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	
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
		if checkMethod(method) {
			checkLogMessage(pass, call)
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

func checkMethod(method string) bool {
	switch method {
		case "Info", "Warn", "Error", "Debug",
			"Print", "Println", "Printf", "Fatal", "Panic", "String":
			return true

		case "With", "WithContext", "WithGroup", "Sugar", "Default":
			return false

		default:
			return false
    }
}

func checkLogMessage(pass *analysis.Pass, call *ast.CallExpr) {
	problems := []string{}
	dangerous_words := []string{"password", "token", "key", "api_key", "jwt", "secret", "bearer", "private_key", "jwt_token", "auth_token", "bearer", "apiKey"}
	for _, arg := range call.Args {
		switch expr := arg.(type) {
			case *ast.BinaryExpr:
				concatWords := getConcats(expr)
				for _, word := range concatWords {
					lower := strings.ToLower(word.Name)
					for _, keyword := range dangerous_words {
						if strings.Contains(lower, keyword) {
							problems = append(problems, fmt.Sprintf("log message should not contain critical information like %s", lower))
							
						} 
				
					}
				}
			case *ast.BasicLit:
				message := []rune(strings.Trim(expr.Value, "\"`"))
				if unicode.IsUpper(message[0]) {
					problems = append(problems, fmt.Sprintf("log message '%s' should be named '%s'", string(message), strings.ToLower(string(message[0])) + string(message[1:])))
				}
				
				if res := useCyrillic(message); res {
					problems = append(problems, fmt.Sprintf("log message '%s' should not use cyrillic characters", string(message)))
				}
				if res := useSpecial(message); res {
					problems = append(problems, fmt.Sprintf("log message '%s' should not use special symbols", string(message)))
				}
			default:
				continue
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
}

func getConcats(binaryExpr ast.Expr) []*ast.Ident {
	concatElements := []*ast.Ident{}
	for binaryExpr != nil {
		switch elem := binaryExpr.(type) {
			case *ast.BinaryExpr:
				binaryExpr = elem.X
				concatElements = append(concatElements, elem.Y.(*ast.Ident))

			default:
				binaryExpr = nil
		}
	}
	return concatElements
}

func useCyrillic(message []rune) bool {
	for _, r := range message {
		if unicode.Is(unicode.Cyrillic, r) {
			return true
		}
	}
	return false
}


func useSpecial(message []rune) bool {
	for _, r := range message {
		if (r != ' ') && (!((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z'))) && !unicode.Is(unicode.Cyrillic, r){
				return true
		}
	}
	return false
}