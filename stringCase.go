package main

import "unicode"

/*
Define String.prototype.toAlternatingCase (or a similar function/method such as to_alternating_case/toAlternatingCase/ToAlternatingCase
in your selected language; see the initial solution for details) such that each lowercase letter becomes uppercase and
each uppercase letter becomes lowercase
*/
func ToAlternatingCase(str string) string {
	result := []rune{}
	for _, ch := range str {
		if unicode.IsUpper(ch) {
			result = append(result, unicode.ToLower(ch))
		} else if unicode.IsLower(ch) {
			result = append(result, unicode.ToUpper(ch))
		} else {
			result = append(result, ch)
		}
	}

	return string(result)
}
