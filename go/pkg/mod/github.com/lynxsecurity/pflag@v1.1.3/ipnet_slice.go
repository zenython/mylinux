// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pflag

import (
	"fmt"
	"io"
	"net"
	"strings"
)

// -- ipNetSlice Value
type ipNetSliceValue struct {
	value   *[]net.IPNet
	changed bool
}

func newIPNetSliceValue(val []net.IPNet, p *[]net.IPNet) *ipNetSliceValue {
	ipnsv := new(ipNetSliceValue)
	ipnsv.value = p
	*ipnsv.value = val
	return ipnsv
}

// Set converts, and assigns, the comma-separated IPNet argument string representation as the []net.IPNet value of this flag.
// If Set is called on a flag that already has a []net.IPNet assigned, the newly converted values will be appended.
func (s *ipNetSliceValue) Set(val string) error {

	// remove all quote characters
	rmQuote := strings.NewReplacer(`"`, "", `'`, "", "`", "")

	// read flag arguments with CSV parser
	ipNetStrSlice, err := readAsCSV(rmQuote.Replace(val))
	if err != nil && err != io.EOF {
		return err
	}

	// parse ip values into slice
	out := make([]net.IPNet, 0, len(ipNetStrSlice))
	for _, ipNetStr := range ipNetStrSlice {
		_, n, err := net.ParseCIDR(strings.TrimSpace(ipNetStr))
		if err != nil {
			return fmt.Errorf("invalid string being converted to CIDR: %s", ipNetStr)
		}
		out = append(out, *n)
	}

	if !s.changed {
		*s.value = out
	} else {
		*s.value = append(*s.value, out...)
	}

	s.changed = true

	return nil
}

// Type returns a string that uniquely represents this flag's type.
func (s *ipNetSliceValue) Type() string {
	return "ipNetSlice"
}

// String defines a "native" format for this net.IPNet slice flag value.
func (s *ipNetSliceValue) String() string {

	ipNetStrSlice := make([]string, len(*s.value))
	for i, n := range *s.value {
		ipNetStrSlice[i] = n.String()
	}

	out, _ := writeAsCSV(ipNetStrSlice)
	return "[" + out + "]"
}

func ipNetSliceConv(val string) (interface{}, error) {
	val = strings.Trim(val, "[]")
	// Empty string would cause a slice with one (empty) entry
	if len(val) == 0 {
		return []net.IPNet{}, nil
	}
	ss := strings.Split(val, ",")
	out := make([]net.IPNet, len(ss))
	for i, sval := range ss {
		_, n, err := net.ParseCIDR(strings.TrimSpace(sval))
		if err != nil {
			return nil, fmt.Errorf("invalid string being converted to CIDR: %s", sval)
		}
		out[i] = *n
	}
	return out, nil
}

// GetIPNetSlice returns the []net.IPNet value of a flag with the given name
func (f *FlagSet) GetIPNetSlice(name string) ([]net.IPNet, error) {
	val, err := f.getFlagType(name, "ipNetSlice", ipNetSliceConv)
	if err != nil {
		return []net.IPNet{}, err
	}
	return val.([]net.IPNet), nil
}

// MustGetIPNetSlice is like GetIPNetSlice, but panics on error.
func (f *FlagSet) MustGetIPNetSlice(name string) []net.IPNet {
	val, err := f.GetIPNetSlice(name)
	if err != nil {
		panic(err)
	}
	return val
}

// IPNetSliceVar defines a ipNetSlice flag with specified name, default value, and usage string.
// The argument p points to a []net.IPNet variable in which to store the value of the flag.
func (f *FlagSet) IPNetSliceVar(p *[]net.IPNet, name string, value []net.IPNet, usage string) {
	f.IPNetSliceVarP(p, name, "", value, usage)
}

// IPNetSliceVarP is like IPNetSliceVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) IPNetSliceVarP(p *[]net.IPNet, name, shorthand string, value []net.IPNet, usage string) {
	f.VarP(newIPNetSliceValue(value, p), name, shorthand, usage)
}

// IPNetSliceVarS is like IPNetSliceVar, but accepts a shorthand letter that can be used after a single dash, alone.
func (f *FlagSet) IPNetSliceVarS(p *[]net.IPNet, name, shorthand string, value []net.IPNet, usage string) {
	f.VarS(newIPNetSliceValue(value, p), name, shorthand, usage)
}

// IPNetSliceVar defines a []net.IPNet flag with specified name, default value, and usage string.
// The argument p points to a []net.IPNet variable in which to store the value of the flag.
func IPNetSliceVar(p *[]net.IPNet, name string, value []net.IPNet, usage string) {
	CommandLine.IPNetSliceVar(p, name, value, usage)
}

// IPNetSliceVarP is like IPNetSliceVar, but accepts a shorthand letter that can be used after a single dash.
func IPNetSliceVarP(p *[]net.IPNet, name, shorthand string, value []net.IPNet, usage string) {
	CommandLine.IPNetSliceVarP(p, name, shorthand, value, usage)
}

// IPNetSliceVarS is like IPNetSliceVar, but accepts a shorthand letter that can be used after a single dash, alone.
func IPNetSliceVarS(p *[]net.IPNet, name, shorthand string, value []net.IPNet, usage string) {
	CommandLine.IPNetSliceVarS(p, name, shorthand, value, usage)
}

// IPNetSlice defines a []net.IPNet flag with specified name, default value, and usage string.
// The return value is the address of a []net.IPNet variable that stores the value of that flag.
func (f *FlagSet) IPNetSlice(name string, value []net.IPNet, usage string) *[]net.IPNet {
	return f.IPNetSliceP(name, "", value, usage)
}

// IPNetSliceP is like IPNetSlice, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) IPNetSliceP(name, shorthand string, value []net.IPNet, usage string) *[]net.IPNet {
	p := []net.IPNet{}
	f.IPNetSliceVarP(&p, name, shorthand, value, usage)
	return &p
}

// IPNetSliceS is like IPNetSlice, but accepts a shorthand letter that can be used after a single dash, alone.
func (f *FlagSet) IPNetSliceS(name, shorthand string, value []net.IPNet, usage string) *[]net.IPNet {
	p := []net.IPNet{}
	f.IPNetSliceVarS(&p, name, shorthand, value, usage)
	return &p
}

// IPNetSlice defines a []net.IPNet flag with specified name, default value, and usage string.
// The return value is the address of a []net.IP variable that stores the value of the flag.
func IPNetSlice(name string, value []net.IPNet, usage string) *[]net.IPNet {
	return CommandLine.IPNetSlice(name, value, usage)
}

// IPNetSliceP is like IPNetSlice, but accepts a shorthand letter that can be used after a single dash.
func IPNetSliceP(name, shorthand string, value []net.IPNet, usage string) *[]net.IPNet {
	return CommandLine.IPNetSliceP(name, shorthand, value, usage)
}

// IPNetSliceS is like IPNetSlice, but accepts a shorthand letter that can be used after a single dash, alone.
func IPNetSliceS(name, shorthand string, value []net.IPNet, usage string) *[]net.IPNet {
	return CommandLine.IPNetSliceS(name, shorthand, value, usage)
}
