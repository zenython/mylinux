// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pflag

// -- string Value
type stringValue string

func newStringValue(val string, p *string) *stringValue {
	*p = val
	return (*stringValue)(p)
}

func (s *stringValue) Set(val string) error {
	*s = stringValue(val)
	return nil
}
func (s *stringValue) Type() string {
	return "string"
}

func (s *stringValue) String() string { return string(*s) }

func stringConv(sval string) (interface{}, error) {
	return sval, nil
}

// GetString return the string value of a flag with the given name
func (f *FlagSet) GetString(name string) (string, error) {
	val, err := f.getFlagType(name, "string", stringConv)
	if err != nil {
		return "", err
	}
	return val.(string), nil
}

// MustGetString is like GetString, but panics on error.
func (f *FlagSet) MustGetString(name string) string {
	val, err := f.GetString(name)
	if err != nil {
		panic(err)
	}
	return val
}

// StringVar defines a string flag with specified name, default value, and usage string.
// The argument p points to a string variable in which to store the value of the flag.
func (f *FlagSet) StringVar(p *string, name string, value string, usage string) {
	f.StringVarP(p, name, "", value, usage)
}

// StringVarP is like StringVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) StringVarP(p *string, name, shorthand string, value string, usage string) {
	f.VarP(newStringValue(value, p), name, shorthand, usage)
}

// StringVarS is like StringVar, but accepts a shorthand letter that can be used after a single dash, alone.
func (f *FlagSet) StringVarS(p *string, name, shorthand string, value string, usage string) {
	f.VarS(newStringValue(value, p), name, shorthand, usage)
}

// StringVar defines a string flag with specified name, default value, and usage string.
// The argument p points to a string variable in which to store the value of the flag.
func StringVar(p *string, name string, value string, usage string) {
	CommandLine.StringVar(p, name, value, usage)
}

// StringVarP is like StringVar, but accepts a shorthand letter that can be used after a single dash.
func StringVarP(p *string, name, shorthand string, value string, usage string) {
	CommandLine.StringVarP(p, name, shorthand, value, usage)
}

// StringVarS is like StringVar, but accepts a shorthand letter that can be used after a single dash, alone.
func StringVarS(p *string, name, shorthand string, value string, usage string) {
	CommandLine.StringVarS(p, name, shorthand, value, usage)
}

// String defines a string flag with specified name, default value, and usage string.
// The return value is the address of a string variable that stores the value of the flag.
func (f *FlagSet) String(name string, value string, usage string) *string {
	return f.StringP(name, "", value, usage)
}

// StringP is like String, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) StringP(name, shorthand string, value string, usage string) *string {
	p := new(string)
	f.StringVarP(p, name, shorthand, value, usage)
	return p
}

// StringS is like String, but accepts a shorthand letter that can be used after a single dash, alone.
func (f *FlagSet) StringS(name, shorthand string, value string, usage string) *string {
	p := new(string)
	f.StringVarS(p, name, shorthand, value, usage)
	return p
}

// String defines a string flag with specified name, default value, and usage string.
// The return value is the address of a string variable that stores the value of the flag.
func String(name string, value string, usage string) *string {
	return CommandLine.String(name, value, usage)
}

// StringP is like String, but accepts a shorthand letter that can be used after a single dash.
func StringP(name, shorthand string, value string, usage string) *string {
	return CommandLine.StringP(name, shorthand, value, usage)
}

// StringS is like String, but accepts a shorthand letter that can be used after a single dash, alone.
func StringS(name, shorthand string, value string, usage string) *string {
	return CommandLine.StringS(name, shorthand, value, usage)
}
