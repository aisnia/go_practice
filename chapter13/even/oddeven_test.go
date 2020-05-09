package even

import "testing"

func TestEven(t *testing.T) {
	//测试奇数，如歌不是则失败Fail
	if Even(10) {
		t.Log(" 10 must be even!")
		t.Fail()
	}
	if Even(7) {
		t.Log(" 7 is not even!")
		t.Fail()
	}

}

func TestOdd(t *testing.T) {
	//测试偶数 如果不是则 Fail失败，并且记录日志
	if !Odd(11) {
		t.Log(" 11 must be odd!")
		t.Fail()
	}
	if Odd(10) {
		t.Log(" 10 is not odd!")
		t.Fail()
	}
}
