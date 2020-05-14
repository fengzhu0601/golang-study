package data

import (
	"testing"
)

const name = "fengzhu"

func TestAdd(t *testing.T){
	s := Add(name)
	if s == ""{
		t.Errorf("Test.Add error! ")
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(name)
	}
}

