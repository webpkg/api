// Copyright 2023 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// https://gostartkit.com
package helper

// Clean make path clean
// /user/role/ to user/role
// /user/role/123 to user/role
func Clean(path string) string {
	l := len(path)
	prev := 0
	str := make([]byte, 0, l)

	for pos := 0; pos < l; pos++ {
		r := path[pos]
		n := false

		switch r {
		case '/', ' ', '\t', '_':
			if pos > prev {
				key := path[prev:pos]
				if !isDigit(key) {
					str = append(str, key...)
					n = true
				}
			}

			prev = pos + 1

			if n && pos > 0 && prev < l {
				str = append(str, '/')
			}
		}
	}

	if prev < l {
		key := path[prev:]
		if !isDigit(key) {
			str = append(str, key...)
		}
	}

	return string(str)
}

// ParsePath parse path to clean name and view name
// /user/role/ to user/role
// /user/role/123 to user/role edit
// /user/role/0 to user/role new
func ParsePath(path string, v string) (string, string) {
	l := len(path)
	prev := 0
	str := make([]byte, 0, l)

	for pos := 0; pos < l; pos++ {
		r := path[pos]
		n := false

		switch r {
		case '/', ' ', '\t', '_':
			if pos > prev {
				key := path[prev:pos]
				if !isDigit(key) {
					str = append(str, key...)
					n = true
				}
			}

			prev = pos + 1

			if n && pos > 0 && prev < l {
				str = append(str, '/')
			}
		}
	}

	if prev < l {
		key := path[prev:]
		if isDigit(key) {
			if v == "" {
				if len(key) == 1 && key[0] == '0' {
					v = "new"
				} else {
					v = "edit"
				}
			}
			str = str[:len(str)-1]
		} else {
			str = append(str, key...)
		}
	}

	return string(str), v
}

func isDigit(val string) bool {

	for i := 0; i < len(val); i++ {

		if !('0' <= val[i] && val[i] <= '9') {
			return false
		}
	}

	return true
}
