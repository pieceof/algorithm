package item

import (
	"fmt"
	"strconv"
)

type Uint64 uint64
func (x Uint64) Less(than Item) (bool, error) {
	value, ok := than.(Uint64)
	if !ok {
		return false, fmt.Errorf("error format, not Uint64.")
	}
	return x < value, nil
}
func (x Uint64) Equal(than Item) (bool, error) {
	value, ok := than.(Uint64)
	if !ok {
		return false, fmt.Errorf("error format, not Uint64.")
	}
	return x == value, nil
}
func (x Uint64) More(than Item) (bool, error) {
	value, ok := than.(Uint64)
	if !ok {
		return false, fmt.Errorf("error format, not Uint64.")
	}
	return x > value, nil
}
func (x Uint64) GetValue() string{
	return strconv.FormatUint(uint64(x),10)
}
