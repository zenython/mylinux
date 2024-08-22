// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pflag

import "strconv"

// -- uint64 Value
type uint64Value uint64

func newUint64Value(val uint64, p *uint64) *uint64Value {
	*p = val
	return (*uint64Value)(p)
}

func (i *uint64Value) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, 64)
	*i = uint64Value(v)
	return err
}

func (i *uint64Value) Type() string {
	return "uint64"
}

func (i *uint64Value) String() string { return strconv.FormatUint(uint64(*i), 10) }

func uint64Conv(sval string) (interface{}, error) {
	v, err := strconv.ParseUint(sval, 0, 64)
	if err != nil {
		return 0, err
	}
	return uint64(v), nil
}

// GetUint64 return the uint64 value of a flag with the given name
func (f *FlagSet) GetUint64(name string) (uint64, error) {
	val, err := f.getFlagType(name, "uint64", uint64Conv)
	if err != nil {
		return 0, err
	}
	return val.(uint64), nil
}

// MustGetUint64 is like GetUint64, but panics on error.
func (f *FlagSet) MustGetUint64(name string) uint64 {
	val, err := f.GetUint64(name)
	if err != nil {
		panic(err)
	}
	return val
}

// Uint64Var defines a uint64 flag with specified name, default value, and usage string.
// The argument p points to a uint64 variable in which to store the value of the flag.
func (f *FlagSet) Uint64Var(p *uint64, name string, value uint64, usage string) {
	f.Uint64VarP(p, name, "", value, usage)
}

// Uint64VarP is like Uint64Var, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Uint64VarP(p *uint64, name, shorthand string, value uint64, usage string) {
	f.VarP(newUint64Value(value, p), name, shorthand, usage)
}

// Uint64VarS is like Uint64Var, but accepts a shorthand letter that can be used after a single dash, alone.
func (f *FlagSet) Uint64VarS(p *uint64, name, shorthand string, value uint64, usage string) {
	f.VarS(newUint64Value(value, p), name, shorthand, usage)
}

// Uint64Var defines a uint64 flag with specified name, default value, and usage string.
// The argument p points to a uint64 variable in which to store the value of the flag.
func Uint64Var(p *uint64, name string, value uint64, usage string) {
	CommandLine.Uint64Var(p, name, value, usage)
}

// Uint64VarP is like Uint64Var, but accepts a shorthand letter that can be used after a single dash.
func Uint64VarP(p *uint64, name, shorthand string, value uint64, usage string) {
	CommandLine.Uint64VarP(p, name, shorthand, value, usage)
}

// Uint64VarS is like Uint64Var, but accepts a shorthand letter that can be used after a single dash, alone.
func Uint64VarS(p *uint64, name, shorthand string, value uint64, usage string) {
	CommandLine.Uint64VarS(p, name, shorthand, value, usage)
}

// Uint64 defines a uint64 flag with specified name, default value, and usage string.
// The return value is the address of a uint64 variable that stores the value of the flag.
func (f *FlagSet) Uint64(name string, value uint64, usage string) *uint64 {
	return f.Uint64P(name, "", value, usage)
}

// Uint64P is like Uint64, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Uint64P(name, shorthand string, value uint64, usage string) *uint64 {
	p := new(uint64)
	f.Uint64VarP(p, name, shorthand, value, usage)
	return p
}

// Uint64S is like Uint64, but accepts a shorthand letter that can be used after a single dash, alone.
func (f *FlagSet) Uint64S(name, shorthand string, value uint64, usage string) *uint64 {
	p := new(uint64)
	f.Uint64VarS(p, name, shorthand, value, usage)
	return p
}

// Uint64 defines a uint64 flag with specified name, default value, and usage string.
// The return value is the address of a uint64 variable that stores the value of the flag.
func Uint64(name string, value uint64, usage string) *uint64 {
	return CommandLine.Uint64(name, value, usage)
}

// Uint64P is like Uint64, but accepts a shorthand letter that can be used after a single dash.
func Uint64P(name, shorthand string, value uint64, usage string) *uint64 {
	return CommandLine.Uint64P(name, shorthand, value, usage)
}

// Uint64S is like Uint64, but accepts a shorthand letter that can be used after a single dash, alone.
func Uint64S(name, shorthand string, value uint64, usage string) *uint64 {
	return CommandLine.Uint64S(name, shorthand, value, usage)
}
