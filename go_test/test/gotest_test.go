package test

import "testing"

func TestDivision1(t *testing.T){
	if i, e := Division(6, 2); i != 3 || e != nil{
		t.Error("测试没有通过")
	}else{
		t.Log("第一个测试通过了")
	}
}

func TestDivision2(t *testing.T){
	if _, e := Division(6, 0); e == nil{
		t.Error("Division did not work as expected")
	}else{
		t.Log("one test passed.", e)
	}
}
