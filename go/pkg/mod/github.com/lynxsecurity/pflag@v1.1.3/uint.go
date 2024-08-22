// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pflag

import "strconv"

// -- uint Value
type uintValue uint

func newUintValue(val uint, p *uint) *uintValue {
	*p = val
	return (*uintValue)(p)
}

func (i *uintValue) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, 64)
	*i = uintValue(v)
	return err
}

func (i *uintValue) Type() string {
	return "uint"
}

func (i *uintValue) String() string { return strconv.FormatUint(uint64(*i), 10) }

func uintConv(sval string) (interface{}, error) {
	v, err := strconv.ParseUint(sval, 0, 0)
	if err != nil {
		return 0, err
	}
	return uint(v), nil
}

// GetUint return the uint value of a flag with the given name
func (f *FlagSet) GetUint(name string) (uint, error) {
	val, err := f.getFlagType(name, "uint", uintConv)
	if err != nil {
		return 0, err
	}
	return val.(uint), nil
}

// MustGetUint is like GetUint, but panics on error.
func (f *FlagSet) MustGetUint(name string) uint {
	val, err := f.GetUint(name)
	if err != nil {
		panic(err)
	}
	return val
}

// UintVar defines a uint flag with specified name, default value, and usage string.
// The argument p points to a uint variable in which to store the value of the flag.
func (f *FlagSet) UintVar(p *uint, name string, value uint, usage string) {
	f.UintVarP(p, name, "", value, usage)
}

// UintVarP is like UintVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) UintVarP(p *uint, name, shorthand string, value uint, usage string) {
	f.VarP(newUintValue(value, p), name, shorthand, usage)
}

// UintVarS is like UintVar, but accepts a shorthand letter that can be used after a single dash, alone.
func (f *FlagSet) UintVarS(p *uint, name, shorthand string, value uint, usage string) {
	f.VarS(newUintValue(value, p), name, shorthand, usage)
}

// UintVar defines a uint flag with specified name, default value, and usage string.
// The argument p points to a uint  variable in which to store the value of the flag.
func UintVar(p *uint, name string, value uint, usage string) {
	CommandLine.UintVar(p, name, value, usage)
}

// UintVarP is like UintVar, but accepts a shorthand letter that can be used after a single dash.
func UintVarP(p *uint, name, shorthand string, value uint, usage string) {
	CommandLine.UintVarP(p, name, shorthand, value, usage)
}

// UintVarS is like UintVar, but accepts a shorthand letter that can be used after a single dash, alone.
func UintVarS(p *uint, name, shorthand string, value uint, usage string) {
	CommandLine.UintVarS(p, name, shorthand, value, usage)
}

// Uint defines a uint flag with specified name, default value, and usage string.
// The return value is the address of a uint  variable that stores the value of the flag.
func (f *FlagSet) Uint(name string, value uint, usage string) *uint {
	return f.UintP(name, "", value, usage)
}

// UintP is like Uint, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) UintP(name, shorthand string, value uint, usage string) *uint {
	p := new(uint)
	f.UintVarP(p, name, shorthand, value, usage)
	return p
}

// UintS is like Uint, but accepts a shorthand letter that can be used after a single dash, alone.
func (f *FlagSet) UintS(name, shorthand string, value uint, usage string) *uint {
	p := new(uint)
	f.UintVarS(p, name, shorthand, value, usage)
	return p
}

// Uint defines a uint flag with specified name, default value, and usage string.
// The return value is the address of a uint  variable that stores the value of the flag.
func Uint(name string, value uint, usage string) *uint {
	return CommandLine.Uint(name, value, usage)
}

// UintP is like Uint, but accepts a shorthand letter that can be used after a single dash.
func UintP(name, shorthand string, value uint, usage string) *uint {
	return CommandLine.UintP(name, shorthand, value, usage)
}

// UintS is like Uint, but accepts a shorthand letter that can be used after a single dash, alone.
func UintS(name, shorthand string, value uint, usage string) *uint {
	return CommandLine.UintS(name, shorthand, value, usage)
}
