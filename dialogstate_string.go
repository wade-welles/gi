// Code generated by "stringer -type=DialogState"; DO NOT EDIT.

package gi

import (
	"fmt"
	"strconv"
)

const _DialogState_name = "DialogExistsDialogOpenModalDialogOpenModelessDialogAcceptedDialogCanceledDialogStateN"

var _DialogState_index = [...]uint8{0, 12, 27, 45, 59, 73, 85}

func (i DialogState) String() string {
	if i < 0 || i >= DialogState(len(_DialogState_index)-1) {
		return "DialogState(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _DialogState_name[_DialogState_index[i]:_DialogState_index[i+1]]
}

func StringToDialogState(s string) (DialogState, error) {
	for i := 0; i < len(_DialogState_index)-1; i++ {
		if s == _DialogState_name[_DialogState_index[i]:_DialogState_index[i+1]] {
			return DialogState(i), nil
		}
	}
	return 0, fmt.Errorf("String %v is not a valid option for type DialogState", s)
}
