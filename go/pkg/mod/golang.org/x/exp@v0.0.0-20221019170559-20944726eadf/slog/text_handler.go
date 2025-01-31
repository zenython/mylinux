// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slog

import (
	"encoding"
	"fmt"
	"io"
	"unicode"
	"unicode/utf8"
)

// TextHandler is a Handler that writes Records to an io.Writer as a
// sequence of key=value pairs separated by spaces and followed by a newline.
type TextHandler struct {
	*commonHandler
}

// NewTextHandler creates a TextHandler that writes to w,
// using the default options.
func NewTextHandler(w io.Writer) *TextHandler {
	return (HandlerOptions{}).NewTextHandler(w)
}

// NewTextHandler creates a TextHandler with the given options that writes to w.
func (opts HandlerOptions) NewTextHandler(w io.Writer) *TextHandler {
	return &TextHandler{
		&commonHandler{
			json: false,
			w:    w,
			opts: opts,
		},
	}
}

// Enabled reports whether the handler handles records at the given level.
// The handler ignores records whose level is lower.
func (h *TextHandler) Enabled(level Level) bool {
	return h.commonHandler.enabled(level)
}

// With returns a new TextHandler whose attributes consists
// of h's attributes followed by attrs.
func (h *TextHandler) WithAttrs(attrs []Attr) Handler {
	return &TextHandler{commonHandler: h.commonHandler.withAttrs(attrs)}
}

func (h *TextHandler) WithGroup(name string) Handler {
	return &TextHandler{commonHandler: h.commonHandler.withGroup(name)}
}

// Handle formats its argument Record as a single line of space-separated
// key=value items.
//
// If the Record's time is zero, the time is omitted.
// Otherwise, the key is "time"
// and the value is output in RFC3339 format with millisecond precision.
//
// If the Record's level is zero, the level is omitted.
// Otherwise, the key is "level"
// and the value of [Level.String] is output.
//
// If the AddSource option is set and source information is available,
// the key is "source" and the value is output as FILE:LINE.
//
// The message's key "msg".
//
// To modify these or other attributes, or remove them from the output, use
// [HandlerOptions.ReplaceAttr].
//
// If a value implements [encoding.TextMarshaler], the result of MarshalText is
// written. Otherwise, the result of fmt.Sprint is written.
//
// Keys and values are quoted with [strconv.Quote] if they contain Unicode space
// characters, non-printing characters, '"' or '='.
//
// Keys inside groups consist of components (keys or group names)
// separated by the Unicode middle dot character, '·'. No further escaping is
// performed. If it is necessary to reconstruct the group structure of a key
// even in the presence of middle dots inside components, use
// [HandlerOptions.ReplaceAttr] to escape the keys.
//
// Each call to Handle results in a single serialized call to
// io.Writer.Write.
func (h *TextHandler) Handle(r Record) error {
	return h.commonHandler.handle(r)
}

func appendTextValue(s *handleState, v Value) error {
	switch v.Kind() {
	case StringKind:
		s.appendString(v.str())
	case TimeKind:
		s.appendTime(v.time())
	case AnyKind:
		if tm, ok := v.any.(encoding.TextMarshaler); ok {
			data, err := tm.MarshalText()
			if err != nil {
				return err
			}
			// TODO: avoid the conversion to string.
			s.appendString(string(data))
			return nil
		}
		s.appendString(fmt.Sprint(v.Any()))
	default:
		*s.buf = v.append(*s.buf)
	}
	return nil
}

func needsQuoting(s string) bool {
	for i := 0; i < len(s); {
		b := s[i]
		if b < utf8.RuneSelf {
			if needsQuotingSet[b] {
				return true
			}
			i++
			continue
		}
		r, size := utf8.DecodeRuneInString(s[i:])
		if r == utf8.RuneError || unicode.IsSpace(r) || !unicode.IsPrint(r) {
			return true
		}
		i += size
	}
	return false
}

var needsQuotingSet = [utf8.RuneSelf]bool{
	'"': true,
	'=': true,
}

func init() {
	for i := 0; i < utf8.RuneSelf; i++ {
		r := rune(i)
		if unicode.IsSpace(r) || !unicode.IsPrint(r) {
			needsQuotingSet[i] = true
		}
	}
}
