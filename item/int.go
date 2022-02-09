package item

import (
	"fmt"
	"strconv"
)

type Int int
func (x Int) Less(than Item) (bool, error) {
	value, ok := than.(Int)
	if !ok {
		return false, fmt.Errorf("error format, not Int.")
	}
	return x < value, nil
}
func (x Int) Equal(than Item) (bool, error) {
	value, ok := than.(Int)
	if !ok {
		return false, fmt.Errorf("error format, not Int.")
	}
	return x == value, nil
}
func (x Int) More(than Item) (bool, error) {
	value, ok := than.(Int)
	if !ok {
		return false, fmt.Errorf("error format, not Int.")
	}
	return x > value, nil
}
func (x Int) GetValue() string {
	return strconv.FormatInt(int64(x),10)
}

