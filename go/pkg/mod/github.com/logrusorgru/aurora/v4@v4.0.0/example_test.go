//
// Copyright (c) 2016-2022 The Aurora Authors. All rights reserved.
// This program is free software. It comes without any warranty,
// to the extent permitted by applicable law. You can redistribute
// it and/or modify it under the terms of the Unlicense. See LICENSE
// file for more details or see below.
//

//
// This is free and unencumbered software released into the public domain.
//
// Anyone is free to copy, modify, publish, use, compile, sell, or
// distribute this software, either in source code form or as a compiled
// binary, for any purpose, commercial or non-commercial, and by any
// means.
//
// In jurisdictions that recognize copyright laws, the author or authors
// of this software dedicate any and all copyright interest in the
// software to the public domain. We make this dedication for the benefit
// of the public at large and to the detriment of our heirs and
// successors. We intend this dedication to be an overt act of
// relinquishment in perpetuity of all present and future rights to this
// software under copyright law.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
// OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
// ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.
//
// For more information, please refer to <http://unlicense.org/>
//

package aurora

import (
	"fmt"
)

func ExampleRed() {
	fmt.Println("value exceeds min-threshold:", Red(3.14))

	// Output: value exceeds min-threshold: [31m3.14[0m
}

func ExampleBold() {
	fmt.Println("value:", Bold(Green(99)))

	// Output: value: [1;32m99[0m
}

func ExampleNew_no_colors() {
	var a = New(WithColors(false), WithHyperlinks(false))
	fmt.Println(a.Red("Not red"))

	// Output: Not red
}

func ExampleNew_colors() {
	var a = New()
	fmt.Println(a.Red("Red"))

	// Output: [31mRed[0m
}

func Example_printf() {
	fmt.Printf("%d %s", Blue(100), BgBlue("cats"))

	// Output: [34m100[0m [44mcats[0m
}

func ExampleSprintf() {
	fmt.Print(
		Sprintf(
			Blue("we've got %d cats, but want %d"), // <- blue format
			Cyan(5),
			Bold(Magenta(25)),
		),
	)

	// Output: [34mwe've got [0;36m5[0;34m cats, but want [0;1;35m25[0;34m[0m
}

func ExampleHyperlink() {
	fmt.Println(Hyperlink(Red("Example"), "http://example.com/"))

	// Output: ]8;;http://example.com/\[31mExample[0m]8;;\
}
