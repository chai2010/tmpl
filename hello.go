// Copyright 2019 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/chai2010/template"
)

func main() {
	fmt.Println(
		template.MustRender(`Hello, {{.}}`, "Neo"),
	)
	fmt.Println(
		template.MustRender(`Hello, {{index . 0}}`, []string{"Go"}),
	)
	fmt.Println(
		template.MustRender(`Hello, {{index . "Name" "SubName"}}`,
			map[string]interface{}{
				"Name": map[string]string{
					"SubName": "凹(Wa)",
				},
			},
		),
	)

	fmt.Println(
		template.MustRender(`Hello, {{.Name}}`, map[string]string{
			"Name": "Lua",
		}),
	)
	fmt.Println(
		template.MustRender(`Hello, {{.Name}}`, struct{ Name string }{
			Name: "Ruby",
		}),
	)

	fmt.Println(
		template.MustRender(
			`Hello, {{upper .Name}}, Age: {{Age}}; Point: {{Point.X}}, {{Map.Name.SubName}}`,
			struct{ Name string }{
				Name: "chai2010",
			},
			template.FuncMap{
				"upper": strings.ToUpper,
				"Age":   func() int { return 20 },
				"Point": func() image.Point { return image.Pt(123, 456) },
				"Map": func() map[string]interface{} {
					return map[string]interface{}{
						"Name": struct{ SubName string }{
							SubName: "chai2010",
						},
					}
				},
			},
		),
	)

	fmt.Println(
		template.MustRender(
			`{{range $i, $v := .}}{{$v.Book}}{{end}}`,
			[]struct{ Name, Book string }{
				{Name: "chai2010", Book: "《Go语言高级编程》"},
				{Name: "chai2010 & ending", Book: "《WebAssembly标准入门》"},
				{Name: "ending & chai2010", Book: "《C/C++面向WebAssembly编程》"},
			},
		),
	)

	fmt.Println(
		template.MustRender(
			`{{range .}}<{{.}}>{{end}}`, []string{"hello", "world"},
		),
	)

	fmt.Println(
		template.MustRenderWithDelims(`Hello, {{<<.))}}`, `<<`, `))`, "Neo"),
	)

	fmt.Println(
		template.MustRender(
			`{{.}}: {{A}}-{{B}}-{{C}}, {{if A}}if A == true{{end}}`,
			"Self", template.FuncMap{
				"A": func() bool { return true },
				"B": func() int { return 9527 },
				"C": func() string { return "C-Value" },
			},
		),
	)

	fmt.Println(
		template.MustRender(
			`slice: {{range $i, $v := slice}}{{$i}}:{{$v}} {{end}}`, nil,
			template.FuncMap{
				"slice": func() []string {
					return []string{"A", "B", "C"}
				},
			},
		),
	)

	fmt.Println(
		template.MustRenderMap("{{.Say}}, {{Year}}", map[string]interface{}{
			"Say":  "Hello",
			"Year": func() interface{} { return 2019 },
		}),
	)

	fmt.Println(
		template.MustRenderFile("hello.tmpl", map[string]string{"Name": "World"}),
	)

	// Output:
	// Hello, Neo
	// Hello, Go
	// Hello, 凹(Wa)
	// Hello, Lua
	// Hello, Ruby
	// Hello, CHAI2010, Age: 20; Point: 123, chai2010
	// 《Go语言高级编程》《WebAssembly标准入门》《C/C++面向WebAssembly编程》
	// Hello, {{Neo}}
	// Self: true-9527-C-Value, if A == true
	// slice: 0:A 1:B 2:C
	// Hello, World
}
