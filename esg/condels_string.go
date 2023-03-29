// Code generated by "stringer -type=CondEls"; DO NOT EDIT.

package esg

import (
	"errors"
	"strconv"
)

var _ = errors.New("dummy error")

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CRule-0]
	_ = x[And-1]
	_ = x[Or-2]
	_ = x[Not-3]
	_ = x[SubCond-4]
	_ = x[CondElsN-5]
}

const _CondEls_name = "CRuleAndOrNotSubCondCondElsN"

var _CondEls_index = [...]uint8{0, 5, 8, 10, 13, 20, 28}

func (i CondEls) String() string {
	if i < 0 || i >= CondEls(len(_CondEls_index)-1) {
		return "CondEls(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _CondEls_name[_CondEls_index[i]:_CondEls_index[i+1]]
}

func (i *CondEls) FromString(s string) error {
	for j := 0; j < len(_CondEls_index)-1; j++ {
		if s == _CondEls_name[_CondEls_index[j]:_CondEls_index[j+1]] {
			*i = CondEls(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: CondEls")
}
