// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pflag

import (
	"fmt"
	"strconv"
	"strings"
)

// -- intSlice Value
type intSliceValue struct {
	value   *[]int
	changed bool
}

func newIntSliceValue(val []int, p *[]int) *intSliceValue {
	isv := new(intSliceValue)
	isv.value = p
	*isv.value = val
	return isv
}

func (s *intSliceValue) Set(val string) error {
	ss := strings.Split(val, ",")
	out := make([]int, len(ss))
	for i, d := range ss {
		var err error
		out[i], err = strconv.Atoi(d)
		if err != nil {
			return err
		}

	}
	if !s.changed {
		*s.value = out
	} else {
		*s.value = append(*s.value, out...)
	}
	s.changed = true
	return nil
}

func (s *intSliceValue) Type() string {
	return "intSlice"
}

func (s *intSliceValue) String() string {
	out := make([]string, len(*s.value))
	for i, d := range *s.value {
		out[i] = fmt.Sprintf("%d", d)
	}
	return "[" + strings.Join(out, ",") + "]"
}

func (s *intSliceValue) Append(val string) error {
	i, err := strconv.Atoi(val)
	if err != nil {
		return err
	}
	*s.value = append(*s.value, i)
	return nil
}

func (s *intSliceValue) Replace(val []string) error {
	out := make([]int, len(val))
	for i, d := range val {
		var err error
		out[i], err = strconv.Atoi(d)
		if err != nil {
			return err
		}
	}
	*s.value = out
	return nil
}

func (s *intSliceValue) GetSlice() []string {
	out := make([]string, len(*s.value))
	for i, d := range *s.value {
		out[i] = strconv.Itoa(d)
	}
	return out
}

func intSliceConv(val string) (interface{}, error) {
	val = strings.Trim(val, "[]")
	// Empty string would cause a slice with one (empty) entry
	if len(val) == 0 {
		return []int{}, nil
	}
	ss := strings.Split(val, ",")
	out := make([]int, len(ss))
	for i, d := range ss {
		var err error
		out[i], err = strconv.Atoi(d)
		if err != nil {
			return nil, err
		}

	}
	return out, nil
}

// GetIntSlice return the []int value of a flag with the given name
func (f *FlagSet) GetIntSlice(name string) ([]int, error) {
	val, err := f.getFlagType(name, "intSlice", intSliceConv)
	if err != nil {
		return []int{}, err
	}
	return val.([]int), nil
}

// MustGetIntSlice is like GetIntSlice, but panics on error.
func (f *FlagSet) MustGetIntSlice(name string) []int {
	val, err := f.GetIntSlice(name)
	if err != nil {
		panic(err)
	}
	return val
}

// IntSliceVar defines a intSlice flag with specified name, default value, and usage string.
// The argument p points to a []int variable in which to store the value of the flag.
func (f *FlagSet) IntSliceVar(p *[]int, name string, value []int, usage string) {
	f.IntSliceVarP(p, name, "", value, usage)
}

// IntSliceVarP is like IntSliceVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) IntSliceVarP(p *[]int, name, shorthand string, value []int, usage string) {
	f.VarP(newIntSliceValue(value, p), name, shorthand, usage)
}

// IntSliceVarS is like IntSliceVar, but accepts a shorthand letter that can be used after a single dash, alone.
func (f *FlagSet) IntSliceVarS(p *[]int, name, shorthand string, value []int, usage string) {
	f.VarS(newIntSliceValue(value, p), name, shorthand, usage)
}

// IntSliceVar defines a int[] flag with specified name, default value, and usage string.
// The argument p points to a int[] variable in which to store the value of the flag.
func IntSliceVar(p *[]int, name string, value []int, usage string) {
	CommandLine.IntSliceVar(p, name, value, usage)
}

// IntSliceVarP is like IntSliceVar, but accepts a shorthand letter that can be used after a single dash.
func IntSliceVarP(p *[]int, name, shorthand string, value []int, usage string) {
	CommandLine.IntSliceVarP(p, name, shorthand, value, usage)
}

// IntSliceVarS is like IntSliceVar, but accepts a shorthand letter that can be used after a single dash, alone.
func IntSliceVarS(p *[]int, name, shorthand string, value []int, usage string) {
	CommandLine.IntSliceVarS(p, name, shorthand, value, usage)
}

// IntSlice defines a []int flag with specified name, default value, and usage string.
// The return value is the address of a []int variable that stores the value of the flag.
func (f *FlagSet) IntSlice(name string, value []int, usage string) *[]int {
	return f.IntSliceP(name, "", value, usage)
}

// IntSliceP is like IntSlice, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) IntSliceP(name, shorthand string, value []int, usage string) *[]int {
	p := []int{}
	f.IntSliceVarP(&p, name, shorthand, value, usage)
	return &p
}

// IntSliceS is like IntSlice, but accepts a shorthand letter that can be used after a single dash, alone.
func (f *FlagSet) IntSliceS(name, shorthand string, value []int, usage string) *[]int {
	p := []int{}
	f.IntSliceVarS(&p, name, shorthand, value, usage)
	return &p
}

// IntSlice defines a []int flag with specified name, default value, and usage string.
// The return value is the address of a []int variable that stores the value of the flag.
func IntSlice(name string, value []int, usage string) *[]int {
	return CommandLine.IntSlice(name, value, usage)
}

// IntSliceP is like IntSlice, but accepts a shorthand letter that can be used after a single dash.
func IntSliceP(name, shorthand string, value []int, usage string) *[]int {
	return CommandLine.IntSliceP(name, shorthand, value, usage)
}

// IntSliceS is like IntSlice, but accepts a shorthand letter that can be used after a single dash, alone.
func IntSliceS(name, shorthand string, value []int, usage string) *[]int {
	return CommandLine.IntSliceS(name, shorthand, value, usage)
}
