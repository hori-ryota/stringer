// Generated by: setup
// TypeWriter: stringer
// Directive: +test on Num

package main

import (
	"fmt"
)

const _Num_name = "m_2m_1m0m1m2"

var _Num_index = [...]uint8{0, 3, 6, 8, 10, 12}

func (i Num) String() string {
	i -= -2
	if i < 0 || i+1 >= Num(len(_Num_index)) {
		return fmt.Sprintf("Num(%d)", i+-2)
	}
	return _Num_name[_Num_index[i]:_Num_index[i+1]]
}