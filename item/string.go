package item

import (
	"algorithm/bkdrhash"
	"fmt"
)

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
func (x String) GetValue() string{
	return x.val
}
