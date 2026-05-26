// Copyright (c) 2020-2023 Uber Technologies, Inc.
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

package atomic

//go:generate bin/gen-atomicwrapper -name=String -type=string -wrapped Value -pack packString -unpack unpackString -compareandswap -swap -file=string.go

func packString(s string) interface{} { _ = "STUB: not implemented"; return nil }

func unpackString(v interface{}) string { _ = "STUB: not implemented"; return "" }

// String returns the wrapped value.
func (s *String) String() string {
	_ = "STUB: not implemented"

	// MarshalText encodes the wrapped string into a textual form.
	//
	// This makes it encodable as JSON, YAML, XML, and more.
	return ""
}

func (s *String) MarshalText() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// UnmarshalText decodes text and replaces the wrapped string with it.
		//
		// This makes it decodable from JSON, YAML, XML, and more.
		nil
}

func (s *String) UnmarshalText(b []byte) error { _ = "STUB: not implemented"; return nil }
