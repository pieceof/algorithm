package bitmap

import (
	"algorithm/bkdrhash"
)

type Bitmap struct {
	words  []uint64
	length uint64
	count  uint64
}

func NewBitmap() *Bitmap {
	return &Bitmap{
		words:  nil,
		length: 0,
		count:  0,
	}
}
func hash(str string) uint64 {
	hasher := bkdrhash.NewHasherBKDR()
	hasher.AddStr(str)
	return hasher.Val()
}

func getNum(num uint64) uint64 {
	ans := uint64(0)
	for ; num != 0; num >>= 1 {
		ans += num & 1
	}
	return ans
}
func (bitmap *Bitmap) Count() uint64 {
	return bitmap.count
}
func(bitmap *Bitmap) Len() uint64{
	return bitmap.length
}
func (bitmap *Bitmap) Add(num uint64) {
	if bitmap.Check(num) {
		return
	}
	index, bit := num>>6, num&0x3f
	for bitmap.length <= index {
		bitmap.words = append(bitmap.words, 0)
		bitmap.length++
	}
	bitmap.words[index] |= 1 << bit
	bitmap.count++
}
func (bitmap *Bitmap) AddStr(str string) {
	bitmap.Add(hash(str))
}
func (bitmap *Bitmap) Min(num uint64) uint64 {
	index, bit := num>>6, num&0x3f
	if index >= bitmap.length {
		return bitmap.count
	}
	ans := uint64(0)
	ans += getNum(bitmap.words[index] & ((1 << bit) - 1))
	for i := uint64(0); i < index; i++ {
		ans += getNum(bitmap.words[i])
	}
	return ans
}

func (bitmap *Bitmap) Max(num uint64) uint64 {
	index, bit := num>>6, num&0x3f
	if index >= bitmap.length {
		return 0
	}
	ans := uint64(0)
	ans += getNum(bitmap.words[index] >> bit)
	for i := index + 1; i < bitmap.length; i++ {
		ans += getNum(bitmap.words[i])
	}
	return ans
}
func (bitmap *Bitmap) Check(num uint64) bool {
	index, bit := num>>6, num&0x3f
	return (bitmap.length > index) && (bitmap.words[index]&(1<<bit) != 0)
}
func (bitmap *Bitmap) CheckStr(str string) bool {
	return bitmap.Check(hash(str))
}

