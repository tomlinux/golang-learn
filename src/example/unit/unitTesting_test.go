package unit

import "testing"

func TestAdd(t *testing.T) {
	ret := Add(3, 5)
	if ret != 8 {
		t.Errorf("期望值是%d，实际的得到是%d", 8, ret)
	}

}

func TestMul(t *testing.T) {
	ret := Mul(5,2)
	if ret != 10 {
		t.Errorf("期望值是%d，实际的得到是%d", 10, ret)
	}
}