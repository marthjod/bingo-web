// Code generated by "stringer -type=Degree"; DO NOT EDIT.

package comparison

import "fmt"

const _Degree_name = "PositiveComparativeSuperlativeUnknown"

var _Degree_index = [...]uint8{0, 8, 19, 30, 37}

func (i Degree) String() string {
	if i < 0 || i >= Degree(len(_Degree_index)-1) {
		return fmt.Sprintf("Degree(%d)", i)
	}
	return _Degree_name[_Degree_index[i]:_Degree_index[i+1]]
}
