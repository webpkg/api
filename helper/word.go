// Copyright 2023 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// https://gostartkit.com
package helper

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

var (
	plural1 = regexp.MustCompile("(?P<keep>[^aeiou])y$")
	plural2 = regexp.MustCompile("(?P<keep>[aeiou]y)$")
	plural3 = regexp.MustCompile("(?P<keep>[sxzh])$")
	plural4 = regexp.MustCompile("(?P<keep>[^sxzhy])$")
	plural5 = regexp.MustCompile("(?P<keep>[^aeiou])ies$")
	plural6 = regexp.MustCompile("(?P<keep>[aeiou]y)s$")
	plural7 = regexp.MustCompile("(?P<keep>[sxzh])es$")
	plural8 = regexp.MustCompile("(?P<keep>[^sxzhyu])s$")

	_prefixKeys = []string{"Ip"}
	_suffixKeys = []string{"Id"}
)

// MakePlural convert word to plural
func MakePlural(name string) string {

	if plural1.MatchString(name) {
		return plural1.ReplaceAllString(name, "${keep}ies")
	}

	if plural2.MatchString(name) {
		return plural2.ReplaceAllString(name, "${keep}s")
	}

	if plural3.MatchString(name) {
		return plural3.ReplaceAllString(name, "${keep}es")
	}

	if plural4.MatchString(name) {
		return plural4.ReplaceAllString(name, "${keep}s")
	}

	return name
}

// MakeSingle convert word to single
func MakeSingle(name string) string {

	if plural5.MatchString(name) {
		return plural5.ReplaceAllString(name, "${keep}y")
	}

	if plural6.MatchString(name) {
		return plural6.ReplaceAllString(name, "${keep}")
	}

	if plural7.MatchString(name) {
		return plural7.ReplaceAllString(name, "${keep}")
	}

	if plural8.MatchString(name) {
		return plural8.ReplaceAllString(name, "${keep}")
	}

	return name
}

// MakeSnake converts the given val to snake case
// like: UserKey,userKey to user_key
func MakeSnake(val string) string {

	val = fixKey(val)

	return makeLowerCase(parseWord(&val), func(s *strings.Builder) {
		s.WriteByte('_')
	})
}

// MakeStudly converts the given val to studly case
// like: user_key, userKey to UserKey
func MakeStudly(val string) string {

	val = fixKey(val)

	return fixEnd(makePascalCase(parseWord(&val), nil))
}

// MakeCamel converts the given val to camel case
// like: user_key to userKey
func MakeCamel(val string) string {

	val = fixKey(val)

	return fixEnd(makeCamelCase(parseWord(&val), nil))
}

// MakeKebab converts the given val to kebab case
// like: UserKey to user-key
func MakeKebab(val string) string {

	val = fixKey(val)

	return makeLowerCase(parseWord(&val), func(s *strings.Builder) {
		s.WriteByte('-')
	})
}

// MakeUrl converts the given val to url
// like: UserKey to user/key
func MakeUrl(val string) string {

	val = fixKey(val)

	return makeLowerCase(parseWord(&val), func(s *strings.Builder) {
		s.WriteByte('/')
	})
}

// ModuleName get ModuleName from  "go.mod"
func ModuleName() string {

	filename := "go.mod"

	r, err := os.Open(filename)

	if err != nil {
		return ""
	}

	defer r.Close()

	scanner := bufio.NewScanner(r)

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {

		line := strings.TrimSpace(scanner.Text())

		if line != "module" {
			continue
		}

		if scanner.Scan() {
			return strings.TrimSpace(scanner.Text())
		}
	}

	return ""
}

// makeLowerCase convert val to lower case then join to string with delimiter
func makeLowerCase(val []string, delimiter func(*strings.Builder)) string {
	var str strings.Builder

	var m byte = 'a' - 'A'

	for i := 0; i < len(val); i++ {
		if i > 0 && delimiter != nil {
			delimiter(&str)
		}

		toLowerCase(&str, &val[i], m)
	}

	return str.String()
}

// makePascalCase convert val to pascal case then join to string with delimiter
func makePascalCase(val []string, delimiter func(*strings.Builder)) string {
	var str strings.Builder

	var m byte = 'a' - 'A'

	for i := 0; i < len(val); i++ {
		if i > 0 && delimiter != nil {
			delimiter(&str)
		}

		toPascalCase(&str, &val[i], m)
	}

	return str.String()
}

// makeCamelCase convert val to camel case then join to string with delimiter
func makeCamelCase(val []string, delimiter func(*strings.Builder)) string {
	var str strings.Builder

	var m byte = 'a' - 'A'

	for i := 0; i < len(val); i++ {
		if i > 0 && delimiter != nil {
			delimiter(&str)
		}

		if i == 0 {
			toLowerCase(&str, &val[i], m)
		} else {
			toPascalCase(&str, &val[i], m)
		}
	}

	return str.String()
}

// MakeSafeString make string safe for sql
func MakeSafeString(val string) string {
	var str strings.Builder

	for i := 0; i < len(val); i++ {

		c := val[i]

		switch c {
		case '\'':
			str.WriteByte(c)
			str.WriteByte(c)
		default:
			str.WriteByte(c)
		}
	}

	return str.String()
}

// toLowerCase convert val to lower case
func toLowerCase(str *strings.Builder, val *string, m byte) {

	for i := 0; i < len(*val); i++ {
		c := (*val)[i]

		if 'A' <= c && c <= 'Z' {
			c += m
		}

		str.WriteByte(c)
	}
}

// toUpperCase convert val to upper case
func toUpperCase(str *strings.Builder, val *string, m byte) {

	for i := 0; i < len(*val); i++ {
		c := (*val)[i]

		if 'a' <= c && c <= 'z' {
			c -= m
		}

		str.WriteByte(c)
	}
}

// toPascalCase convert val to pascal case
func toPascalCase(str *strings.Builder, val *string, m byte) {

	for i := 0; i < len(*val); i++ {
		c := (*val)[i]

		if i == 0 {
			if 'a' <= c && c <= 'z' {
				c -= m
			}
		} else {
			if 'A' <= c && c <= 'Z' {
				c += m
			}
		}

		str.WriteByte(c)
	}
}

// parseWord parse the given val to words
// like: user_key,user-key to [user, key]
// like: userKey to [user, Key]
// like: UserKey to [User, Key]
func parseWord(val *string) []string {
	l := len(*val)

	prev := 0

	words := make([]string, 0)

	for pos := 0; pos < l; pos++ {
		r := (*val)[pos]

		switch r {
		case '_', '-', ' ', '\t':

			if pos > prev {
				words = append(words, (*val)[prev:pos])
			}

			prev = pos + 1
		default:
			if r >= 'A' && r <= 'Z' {
				if pos > prev {
					words = append(words, (*val)[prev:pos])
				}

				prev = pos
			}
		}
	}

	if prev < l {
		words = append(words, (*val)[prev:])
	}

	return words
}

func fixKey(val string) string {

	for _, key := range _prefixKeys {
		if strings.HasPrefix(val, strings.ToUpper(key)) {
			val = key + val[len(key):]
			break
		}
	}

	for _, key := range _suffixKeys {
		if strings.HasSuffix(val, strings.ToUpper(key)) {
			val = val[:len(val)-len(key)] + key
			break
		}
	}

	return val
}

func fixEnd(val string) string {

	for _, key := range _suffixKeys {
		if strings.HasSuffix(val, key) {
			val = val[:len(val)-len(key)] + strings.ToUpper(key)
			break
		}
	}

	return val
}
