// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pflag

import "strconv"

// -- int8 Value
type int8Value int8

func newInt8Value(val int8, p *int8) *int8Value {
	*p = val
	return (*int8Value)(p)
}

func (i *int8Value) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 8)
	*i = int8Value(v)
	return err
}

func (i *int8Value) Type() string {
	return "int8"
}

func (i *int8Value) String() string { return strconv.FormatInt(int64(*i), 10) }

func int8Conv(sval string) (interface{}, error) {
	v, err := strconv.ParseInt(sval, 0, 8)
	if err != nil {
		return 0, err
	}
	return int8(v), nil
}

// GetInt8 return the int8 value of a flag with the given name
func (f *FlagSet) GetInt8(name string) (int8, error) {
	val, err := f.getFlagType(name, "int8", int8Conv)
	if err != nil {
		return 0, err
	}
	return val.(int8), nil
}

// MustGetInt8 is like GetInt8, but panics on error.
func (f *FlagSet) MustGetInt8(name string) int8 {
	val, err := f.GetInt8(name)
	if err != nil {
		panic(err)
	}
	return val
}

// Int8Var defines an int8 flag with specified name, default value, and usage string.
// The argument p points to an int8 variable in which to store the value of the flag.
func (f *FlagSet) Int8Var(p *int8, name string, value int8, usage string) {
	f.Int8VarP(p, name, "", value, usage)
}

// Int8VarP is like Int8Var, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Int8VarP(p *int8, name, shorthand string, value int8, usage string) {
	f.VarP(newInt8Value(value, p), name, shorthand, usage)
}

// Int8VarS is like Int8Var, but accepts a shorthand letter that can be used after a single dash, alone.
func (f *FlagSet) Int8VarS(p *int8, name, shorthand string, value int8, usage string) {
	f.VarS(newInt8Value(value, p), name, shorthand, usage)
}

// Int8Var defines an int8 flag with specified name, default value, and usage string.
// The argument p points to an int8 variable in which to store the value of the flag.
func Int8Var(p *int8, name string, value int8, usage string) {
	CommandLine.Int8Var(p, name, value, usage)
}

// Int8VarP is like Int8Var, but accepts a shorthand letter that can be used after a single dash.
func Int8VarP(p *int8, name, shorthand string, value int8, usage string) {
	CommandLine.Int8VarP(p, name, shorthand, value, usage)
}

// Int8VarS is like Int8Var, but accepts a shorthand letter that can be used after a single dash, alone.
func Int8VarS(p *int8, name, shorthand string, value int8, usage string) {
	CommandLine.Int8VarS(p, name, shorthand, value, usage)
}

// Int8 defines an int8 flag with specified name, default value, and usage string.
// The return value is the address of an int8 variable that stores the value of the flag.
func (f *FlagSet) Int8(name string, value int8, usage string) *int8 {
	return f.Int8P(name, "", value, usage)
}

// Int8P is like Int8, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Int8P(name, shorthand string, value int8, usage string) *int8 {
	p := new(int8)
	f.Int8VarP(p, name, shorthand, value, usage)
	return p
}

// Int8S is like Int8, but accepts a shorthand letter that can be used after a single dash, alone.
func (f *FlagSet) Int8S(name, shorthand string, value int8, usage string) *int8 {
	p := new(int8)
	f.Int8VarS(p, name, shorthand, value, usage)
	return p
}

// Int8 defines an int8 flag with specified name, default value, and usage string.
// The return value is the address of an int8 variable that stores the value of the flag.
func Int8(name string, value int8, usage string) *int8 {
	return CommandLine.Int8(name, value, usage)
}

// Int8P is like Int8, but accepts a shorthand letter that can be used after a single dash.
func Int8P(name, shorthand string, value int8, usage string) *int8 {
	return CommandLine.Int8P(name, shorthand, value, usage)
}

// Int8S is like Int8, but accepts a shorthand letter that can be used after a single dash, alone.
func Int8S(name, shorthand string, value int8, usage string) *int8 {
	return CommandLine.Int8S(name, shorthand, value, usage)
}
