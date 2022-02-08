package bkdrhash

import (
	"reflect"
	"testing"
)

func TestHasherBKDR_AddInt(t *testing.T) {
	type fields struct {
		hasherNum uint64
	}
	type args struct {
		num uint64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hasher := &HasherBKDR{
				hasherNum: tt.fields.hasherNum,
			}
		})
	}
}

func TestHasherBKDR_AddStr(t *testing.T) {
	type fields struct {
		hasherNum uint64
	}
	type args struct {
		str string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hasher := &HasherBKDR{
				hasherNum: tt.fields.hasherNum,
			}
		})
	}
}

func TestHasherBKDR_Val(t *testing.T) {
	type fields struct {
		hasherNum uint64
	}
	tests := []struct {
		name   string
		fields fields
		want   uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hasher := &HasherBKDR{
				hasherNum: tt.fields.hasherNum,
			}
			if got := hasher.Val(); got != tt.want {
				t.Errorf("Val() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewHasherBKDR(t *testing.T) {
	tests := []struct {
		name string
		want *HasherBKDR
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHasherBKDR(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHasherBKDR() = %v, want %v", got, tt.want)
			}
		})
	}
}
