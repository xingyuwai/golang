/*
 * @Author: Mr Bian
 * @Date: 2018-08-19 19:19:05
 * @LastEditors: Mr Bian
 * @LastEditTime: 2018-08-19 22:21:56
 * @Description:
 * @version:
 */

package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (err ErrNegativeSqrt) Error() string {
	/**
		* fmt.Sprint(e) will call e.Error() to convert the
		* value e to a string. If the Error() method calls fmt.Sprint(e),
		* then the program recurses until out of memory.
	    * You can break the recursion by converting the e
	    * to a value without a String or Error method
	*/
	return fmt.Sprintf("Cannot Sqrt negative number: %v\n", float64(err))
}

func Sqrt(x float64) (float64, error) {
	if x < 0.0 {
		return 0, ErrNegativeSqrt(x)
	}

	result := 0.0
	for temp := x; ; temp = result {
		result = temp - (temp*temp-x)/(2*temp)
		if temp-result < 0.000001 && temp-result > -0.000001 {
			break
		}

	}
	return result, nil
}
