// Code generated by "stringer -type=ButtonSignals"; DO NOT EDIT.

package gi

import (
	"errors"
	"strconv"
)

var _ = errors.New("dummy error")

const _ButtonSignals_name = "ButtonClickedButtonPressedButtonReleasedButtonToggledButtonSignalsN"

var _ButtonSignals_index = [...]uint8{0, 13, 26, 40, 53, 67}

func (i ButtonSignals) String() string {
	if i < 0 || i >= ButtonSignals(len(_ButtonSignals_index)-1) {
		return "ButtonSignals(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ButtonSignals_name[_ButtonSignals_index[i]:_ButtonSignals_index[i+1]]
}

func (i *ButtonSignals) FromString(s string) error {
	for j := 0; j < len(_ButtonSignals_index)-1; j++ {
		if s == _ButtonSignals_name[_ButtonSignals_index[j]:_ButtonSignals_index[j+1]] {
			*i = ButtonSignals(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: ButtonSignals")
}
