package main

import (
    "go/parser"
    goast "go/ast"
    "strconv"
    "errors"
    "fmt"
    "os"
    "math"
    // "github.com/davecgh/go-spew"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Printf("ERRO: Adicione contas a serem feitas\n")
        os.Exit(1)
    }
    for i := 1; i < len(os.Args); i++ {
        ast, err := parser.ParseExpr(os.Args[i]);
        if err != nil {
            panic(err)
        }
        res, err := expandStatement(ast)
        if err != nil {
            panic(err)
        }
        fmt.Printf("%d > %s -> %f\n", i, os.Args[i], res)
    }
    // spew.Dump(ast)
}

func expandStatement(ast goast.Expr) (float64, error) {
    switch v := ast.(type) {
    case *goast.BasicLit:
        return strconv.ParseFloat(v.Value, 64)
    case *goast.BinaryExpr:
        return processSides(v)
    case *goast.ParenExpr:
        return expandStatement(v.X)
    default:
        return 0, errors.New(fmt.Sprintf("Undefined statement member: %T", v))
    }
}

func processSides(ast *goast.BinaryExpr) (float64, error) {
    va, err := expandStatement(ast.X)
    if err != nil {
        return 0, err
    }
    vb, err := expandStatement(ast.Y)
    if err != nil {
        return 0, err
    }
    switch ast.Op.String() {
        case "+": 
            return va + vb, nil
        case "-": 
            return va - vb, nil
        case "*": 
            return va * vb, nil
        case "/": 
            return va / vb, nil
        case "^":
            return math.Pow(va, vb), nil
        default: 
        return 0, errors.New(fmt.Sprintf("undefined operation: %s", ast.Op.String()))
    }
}
