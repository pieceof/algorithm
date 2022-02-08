package bkdrhash

import (
	"strconv"
)

type HasherBKDR struct {
	hasherNum uint64
}

func NewHasherBKDR() *HasherBKDR{
	return &HasherBKDR{
		hasherNum: 0,
	}
}

func (hasher *HasherBKDR) AddStr(str string) {
	val := uint64(hasher.hasherNum)
	for i := 0;i < len(str);i++{
		val = val * 131 + uint64(str[i])
	}
	hasher.hasherNum = val
	return
}

func (hasher *HasherBKDR) AddInt(num uint64){
	hasher.AddStr(strconv.FormatUint(num,10))
	return
}

func (hasher *HasherBKDR) Val() uint64{
	return hasher.hasherNum
}