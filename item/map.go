package item

import "fmt"

type Map struct {
	key String
	val Item
}
func (x Map) Less(than Item) (bool, error){
	value, ok := than.(Map)
	if !ok {
		return false, fmt.Errorf("error format, not Map.")
	}
	return x.key.Less(value.key)
}
func (x Map) Equal(than Item) (bool, error) {
	value, ok := than.(Map)
	if !ok {
		return false, fmt.Errorf("error format, not Map.")
	}
	return x.key.Equal(value.key)
}
func (x Map) More(than Item) (bool, error) {
	value, ok := than.(Map)
	if !ok {
		return false, fmt.Errorf("error format, not Map.")
	}
	return x.key.More(value.key)
}
func (x Map) GetValue() string{
	return fmt.Sprintf("key: %v; value: %v", x.key.GetValue(), x.val.GetValue())
}

