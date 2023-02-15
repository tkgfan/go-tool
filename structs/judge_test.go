// author lby
// date 2023/2/15

package structs

import "testing"

func TestIsNil(t *testing.T) {
	// 1
	if IsNil(nil) != true {
		t.Error("IsNil has error")
	}

	// 2
	if IsNil("") == true {
		t.Error("IsNil has error")
	}

	// 3
	var a *int
	if IsNil(a) != true {
		t.Error("IsNil has error")
	}

	// 4
	a = nil
	var b interface{} = a
	if IsNil(b) != true {
		t.Error("IsNil has error")
	}
}
