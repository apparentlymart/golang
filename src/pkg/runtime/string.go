// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import (
	"unsafe"
)

func concatstrings(a []string) string {
	idx := 0
	l := 0
	count := 0
	for i, x := range a {
		n := len(x)
		if n == 0 {
			continue
		}
		if l+n < l {
			panic("string concatenation too long")
		}
		l += n
		count++
		idx = i
	}
	if count == 0 {
		return ""
	}
	if count == 1 {
		return a[idx]
	}
	s, b := rawstring(l)
	l = 0
	for _, x := range a {
		copy(b[l:], x)
		l += len(x)
	}
	return s
}

//go:nosplit
func concatstring2(a [2]string) string {
	return concatstrings(a[:])
}

//go:nosplit
func concatstring3(a [3]string) string {
	return concatstrings(a[:])
}

//go:nosplit
func concatstring4(a [4]string) string {
	return concatstrings(a[:])
}

//go:nosplit
func concatstring5(a [5]string) string {
	return concatstrings(a[:])
}

func slicebytetostring(b []byte) string {
	if raceenabled && len(b) > 0 {
		fn := slicebytetostring
		racereadrangepc(unsafe.Pointer(&b[0]),
			len(b),
			gogetcallerpc(unsafe.Pointer(&b)),
			**(**uintptr)(unsafe.Pointer(&fn)))
	}
	s, c := rawstring(len(b))
	copy(c, b)
	return s
}

func slicebytetostringtmp(b []byte) string {
	// Return a "string" referring to the actual []byte bytes.
	// This is only for use by internal compiler optimizations
	// that know that the string form will be discarded before
	// the calling goroutine could possibly modify the original
	// slice or synchronize with another goroutine.
	// Today, the only such case is a m[string(k)] lookup where
	// m is a string-keyed map and k is a []byte.

	if raceenabled && len(b) > 0 {
		fn := slicebytetostringtmp
		racereadrangepc(unsafe.Pointer(&b[0]),
			len(b),
			gogetcallerpc(unsafe.Pointer(&b)),
			**(**uintptr)(unsafe.Pointer(&fn)))
	}
	return *(*string)(unsafe.Pointer(&b))
}

func stringtoslicebyte(s string) []byte {
	b := rawbyteslice(len(s))
	copy(b, s)
	return b
}

func stringtoslicerune(s string) []rune {
	// two passes.
	// unlike slicerunetostring, no race because strings are immutable.
	n := 0
	t := s
	for len(s) > 0 {
		_, k := charntorune(s)
		s = s[k:]
		n++
	}
	a := rawruneslice(n)
	n = 0
	for len(t) > 0 {
		r, k := charntorune(t)
		t = t[k:]
		a[n] = r
		n++
	}
	return a
}

func slicerunetostring(a []rune) string {
	if raceenabled && len(a) > 0 {
		fn := slicerunetostring
		racereadrangepc(unsafe.Pointer(&a[0]),
			len(a)*int(unsafe.Sizeof(a[0])),
			gogetcallerpc(unsafe.Pointer(&a)),
			**(**uintptr)(unsafe.Pointer(&fn)))
	}
	var dum [4]byte
	size1 := 0
	for _, r := range a {
		size1 += runetochar(dum[:], r)
	}
	s, b := rawstring(size1 + 3)
	size2 := 0
	for _, r := range a {
		// check for race
		if size2 >= size1 {
			break
		}
		size2 += runetochar(b[size2:], r)
	}
	return s[:size2]
}

type stringStruct struct {
	str *byte
	len int
}

func cstringToGo(str uintptr) (s string) {
	i := 0
	for ; ; i++ {
		if *(*byte)(unsafe.Pointer(str + uintptr(i))) == 0 {
			break
		}
	}
	t := (*stringStruct)(unsafe.Pointer(&s))
	t.str = (*byte)(unsafe.Pointer(str))
	t.len = i
	return
}

func intstring(v int64) string {
	s, b := rawstring(4)
	n := runetochar(b, rune(v))
	return s[:n]
}

// stringiter returns the index of the next
// rune after the rune that starts at s[k].
func stringiter(s string, k int) int {
	if k >= len(s) {
		// 0 is end of iteration
		return 0
	}

	c := s[k]
	if c < runeself {
		return k + 1
	}

	// multi-char rune
	_, n := charntorune(s[k:])
	return k + n
}

// stringiter2 returns the rune that starts at s[k]
// and the index where the next rune starts.
func stringiter2(s string, k int) (int, rune) {
	if k >= len(s) {
		// 0 is end of iteration
		return 0, 0
	}

	c := s[k]
	if c < runeself {
		return k + 1, rune(c)
	}

	// multi-char rune
	r, n := charntorune(s[k:])
	return k + n, r
}
