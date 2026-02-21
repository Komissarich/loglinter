package loglinter

import (
	"fmt"
	"go/ast"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
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

		// fmt.Println("args: ", call.Args[0].(*ast.BinaryExpr).X.(*ast.Ident).Name)
	
		// fmt.Println(base.Name, method, cleanLogMessage)

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
			"Print", "Println", "Printf", "Fatal", "Panic":
			return true

		case "With", "WithContext", "WithGroup", "Sugar", "Default":
			return false

		default:
			return false
    }
}

func checkLogMessage(pass *analysis.Pass, call *ast.CallExpr) {
	dangerous_words := []string{"password", "token", "key", "api_key", "jwt", "secret", "bearer", "private_key", "jwt_token", "auth_token", "bearer"}
	for _, arg := range call.Args {
		fmt.Println(arg)
		dangerous_concat, ok := arg.(*ast.BinaryExpr)
		if !ok {
			fmt.Println("helloas")
			logMessage, ok := arg.(*ast.BasicLit)
			if !ok {
				continue
			}
			message := []rune(strings.Trim(logMessage.Value, "\"`"))
			if unicode.IsUpper(message[0]) {
				pass.Reportf(call.Pos(), "log message '%s' should be named '%s'",
				string(message), strings.ToLower(string(message[0])) + string(message[1:]))
			}
			for _, r := range message {
				if !unicode.Is(unicode.Latin, r) {
					if unicode.Is(unicode.Cyrillic, r) {
						pass.Reportf(call.Pos(), "log message '%s' should not use cyrillic characters",
						string(message))
					}
					if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')) {
						pass.Reportf(call.Pos(), "log message '%s' should not use special symbols",
						string(message))
					}
				}
			}
			continue
		} 
		dangerous_var, ok := dangerous_concat.Y.(*ast.Ident)
		if !ok {
			dangerous_var, ok = dangerous_concat.X.(*ast.Ident)
			if !ok {
				continue
			}
		}
		lower := strings.ToLower(dangerous_var.Name)
		for _, keyword := range dangerous_words {
			if strings.Contains(lower, keyword) {
				pass.Reportf(call.Pos(), "log message should not use contain critical information like %s", lower)
			} else {
				continue
			}

		fmt.Println("again")
	}
	
	}
}