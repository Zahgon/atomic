// Copyright (c) 2020-2022 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// gen-atomicwrapper generates wrapper types around other atomic types.
//
// It supports plugging in functions which convert the value inside the atomic
// type to the user-facing value. For example,
//
// Given, atomic.Value and the functions,
//
//	func packString(string) interface{}
//	func unpackString(interface{}) string
//
// We can run the following command:
//
//	gen-atomicwrapper -name String -wrapped Value \
//	  -type string -pack fromString -unpack tostring
//
// This will generate approximately,
//
//	type String struct{ v Value }
//
//	func (s *String) Load() string {
//	  return unpackString(v.Load())
//	}
//
//	func (s *String) Store(s string) {
//	  return s.v.Store(packString(s))
//	}
//
// The packing/unpacking logic allows the stored value to be different from
// the user-facing value.
package main

import (
	"embed"
	"log"
	"os"
	"text/template"
)

func main() {
	log.SetFlags(0)
	if err := run(os.Args[1:]); err != nil {
		log.Fatalf("%+v", err)
	}
}

type stringList []string

func (sl *stringList) String() string { _ = "STUB: not implemented"; return "" }

func (sl *stringList) Set(s string) error { _ = "STUB: not implemented"; return nil }

func run(args []string) error { _ = "STUB: not implemented"; return nil }

// Required flags

// Optional flags

// Switches for individual methods. Underlying atomics must support
// these.

// Import encoding/json if needed.

var (
	//go:embed *.tmpl
	_tmplFS embed.FS

	_tmpl = template.Must(template.New("atomicwrapper").ParseFS(_tmplFS, "*.tmpl"))
)
