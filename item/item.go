package item

type Item interface {
	Less(than Item) (bool, error)
	Equal(than Item) (bool, error)
	More(than Item) (bool, error)
	GetValue() (string)
}