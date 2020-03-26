package gotest

import "testing"

func TestGetArea(t *testing.T) {
	area := GetArea(40, 50)
	if area != 200 {
		t.Error("测试失败")
	} else {
		t.Log("测试成功")
	}
}
