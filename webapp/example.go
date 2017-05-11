// Copyright 2011-2016 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"html/template"
	"net/http"
)

func init() {
	http.HandleFunc("/example", exampleHandler)
	http.HandleFunc("/example2", exampleHandler)
}

var exampleTemplate = template.Must(template.ParseFiles("static/example.html"))
var exampleTemplate2 = template.Must(template.ParseFiles("_examples/index.html"))

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	exampleTemplate.Execute(w, nil)
}

func exampleHandler2(w http.ResponseWriter, r *http.Request) {
	exampleTemplate2.Execute(w, nil)
}
