package tests

import (
	"errors"
	"time"
)

// @memo
func Test(a string, b string) string {
	return "something"
}

// @memo
func TestWithArgs(a string, b string) {

}

// @memo
func TestNothing() {
	time.Sleep(time.Second)
}

// @memo
func TestWithOneReturn() string {
	return "a"
}

// @memo
func TestWithMultipleReturnValues() (string, int, error) {
	return "a", 2, errors.New("test")
}
