// Code generated by "stringer -type=state"; DO NOT EDIT.

package fee

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[UnderInspection-1]
	_ = x[Approved-2]
	_ = x[InAction-3]
	_ = x[Suspended-4]
	_ = x[UnderReturn-5]
	_ = x[End-6]
}

const _state_name = "UnderInspectionApprovedInActionSuspendedUnderReturnEnd"

var _state_index = [...]uint8{0, 15, 23, 31, 40, 51, 54}

func (i state) String() string {
	i -= 1
	if i < 0 || i >= state(len(_state_index)-1) {
		return "state(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _state_name[_state_index[i]:_state_index[i+1]]
}
