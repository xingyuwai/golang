/**
 * @Author: Mr Bian
 * @Date: 2018-08-14 16:51:15
 * @LastEditors: Mr Bian
 * @LastEditTime: 2018-08-14 16:57:05
 * @Description:
 * @version:
 */

// Package stringutil contains utility functions for working with strings.
package stringutil

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
