package generator

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

// Generator 代码生成器
type Generator struct {
	pkgPath string
	exports map[string]*ast.FuncDecl
}

// NewGenerator 创建生成器
func NewGenerator(pkgPath string) *Generator {
	return &Generator{
		pkgPath: pkgPath,
		exports: make(map[string]*ast.FuncDecl),
	}
}

// Generate 生成头文件
func (g *Generator) Generate() (string, error) {
	// 扫描导出函数
	if err := g.scanExports(); err != nil {
		return "", err
	}

	// 生成头文件内容
	var b strings.Builder

	// 写入头部
	b.WriteString(`/* Generated by py_go_bridge */
#ifndef PY_GO_BRIDGE_H
#define PY_GO_BRIDGE_H

#include <stdlib.h>

#ifdef __cplusplus
extern "C" {
#endif

`)

	// 写入函数声明
	for name, fn := range g.exports {
		b.WriteString(g.generateFuncDecl(name, fn))
		b.WriteString("\n")
	}

	// 写入尾部
	b.WriteString(`
#ifdef __cplusplus
}
#endif

#endif /* PY_GO_BRIDGE_H */
`)

	return b.String(), nil
}

// scanExports 扫描导出函数
func (g *Generator) scanExports() error {
	fset := token.NewFileSet()

	// 解析包
	pkgs, err := parser.ParseDir(fset, g.pkgPath, nil, 0)
	if err != nil {
		return err
	}

	// 遍历所有文件
	for _, pkg := range pkgs {
		for _, file := range pkg.Files {
			// 遍历所有声明
			for _, decl := range file.Decls {
				if fn, ok := decl.(*ast.FuncDecl); ok {
					// 检查是否有export注释
					if g.hasExportComment(fn) {
						g.exports[fn.Name.Name] = fn
					}
				}
			}
		}
	}

	return nil
}

// hasExportComment 检查是否有export注释
func (g *Generator) hasExportComment(fn *ast.FuncDecl) bool {
	if fn.Doc == nil {
		return false
	}

	for _, comment := range fn.Doc.List {
		if strings.Contains(comment.Text, "//export") {
			return true
		}
	}

	return false
}

// generateFuncDecl 生成函数声明
func (g *Generator) generateFuncDecl(name string, fn *ast.FuncDecl) string {
	// TODO: 实现函数声明生成
	return fmt.Sprintf("extern void* %s();", name)
}
