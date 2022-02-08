package item

import (
	"algorithm/bkdrhash"
	"fmt"
)

type Item interface {
	Less(than Item) (bool, error)
	Equal(than Item) (bool, error)
	More(than Item) (bool, error)
}

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

type String struct {
	hashVal uint64
	val string
}
func NewString(str string) String{
	hasher := bkdrhash.NewHasherBKDR()
	hasher.AddStr(str)
	return String{
		hashVal: hasher.Val(),
		val : str,
	}
}
func (x String) Less(than Item) (bool, error) {
	value, ok := than.(String)
	if !ok {
		return false, fmt.Errorf("error format, not String.")
	}
	return x.hashVal < value.hashVal, nil
}
func (x String) Equal(than Item) (bool, error) {
	value, ok := than.(String)
	if !ok {
		return false, fmt.Errorf("error format, not String.")
	}
	return x.hashVal == value.hashVal, nil
}
func (x String) More(than Item) (bool, error) {
	value, ok := than.(String)
	if !ok {
		return false, fmt.Errorf("error format, not String.")
	}
	return x.hashVal > value.hashVal, nil
}
