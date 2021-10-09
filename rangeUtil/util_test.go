package rangeUtil

import (
	"testing"
)

func TestRemoveBracket(t *testing.T) {
	input := ""
	expect := ""
	get := removeBracket(input)
	if get != expect {
		t.Errorf("Expect to get %s, but get %s instead.", expect, get)
		return
	}

	input = "[1,5]"
	expect = "1,5"
	get = removeBracket(input)
	if get != expect {
		t.Errorf("Expect to get %s, but get %s instead.", expect, get)
		return
	}

	input = "[1,5)"
	expect = "1,5"
	get = removeBracket(input)
	if get != expect {
		t.Errorf("Expect to get %s, but get %s instead.", expect, get)
		return
	}

	input = "(1,5]"
	expect = "1,5"
	get = removeBracket(input)
	if get != expect {
		t.Errorf("Expect to get %s, but get %s instead.", expect, get)
		return
	}

	input = "(1,5)"
	expect = "1,5"
	get = removeBracket(input)
	if get != expect {
		t.Errorf("Expect to get %s, but get %s instead.", expect, get)
		return
	}

	input = "{1,5}"
	expect = "{1,5}"
	get = removeBracket(input)
	if get != expect {
		t.Errorf("Expect to get %s, but get %s instead.", expect, get)
		return
	}

	input = "[1,5}"
	expect = "1,5}"
	get = removeBracket(input)
	if get != expect {
		t.Errorf("Expect to get %s, but get %s instead.", expect, get)
		return
	}

	input = "{1,5]"
	expect = "{1,5"
	get = removeBracket(input)
	if get != expect {
		t.Errorf("Expect to get %s, but get %s instead.", expect, get)
		return
	}

	input = "(1,5}"
	expect = "1,5}"
	get = removeBracket(input)
	if get != expect {
		t.Errorf("Expect to get %s, but get %s instead.", expect, get)
		return
	}

	input = "{1,5)"
	expect = "{1,5"
	get = removeBracket(input)
	if get != expect {
		t.Errorf("Expect to get %s, but get %s instead.", expect, get)
		return
	}
}

func TestIsEmptyRange(t *testing.T) {
	input := ""
	expect := true
	get := IsEmptyRange(input)
	if get != expect {
		t.Errorf("Expect to get %v, but get %v instead.", expect, get)
		return
	}

	input = "[]"
	expect = true
	get = IsEmptyRange(input)
	if get != expect {
		t.Errorf("Expect to get %v, but get %v instead.", expect, get)
		return
	}

	input = "[)"
	expect = true
	get = IsEmptyRange(input)
	if get != expect {
		t.Errorf("Expect to get %v, but get %v instead.", expect, get)
		return
	}

	input = "(]"
	expect = true
	get = IsEmptyRange(input)
	if get != expect {
		t.Errorf("Expect to get %v, but get %v instead.", expect, get)
		return
	}

	input = "()"
	expect = true
	get = IsEmptyRange(input)
	if get != expect {
		t.Errorf("Expect to get %v, but get %v instead.", expect, get)
		return
	}

	input = "[}"
	expect = false
	get = IsEmptyRange(input)
	if get != expect {
		t.Errorf("Expect to get %v, but get %v instead.", expect, get)
		return
	}

	input = "[1,2]"
	expect = false
	get = IsEmptyRange(input)
	if get != expect {
		t.Errorf("Expect to get %v, but get %v instead.", expect, get)
		return
	}

	input = "[a]"
	expect = false
	get = IsEmptyRange(input)
	if get != expect {
		t.Errorf("Expect to get %v, but get %v instead.", expect, get)
		return
	}
}

func TestIsValidFormat(t *testing.T) {
	input := ""
	expect := false
	get := IsValidFormat(input)
	if get != expect {
		t.Errorf("Expect to get %v, but get %v instead.", expect, get)
		return
	}

	input = "["
	expect = false
	get = IsValidFormat(input)
	if get != expect {
		t.Errorf("Expect to get %v, but get %v instead.", expect, get)
		return
	}

	input = "[]"
	expect = false
	get = IsValidFormat(input)
	if get != expect {
		t.Errorf("Expect to get %v, but get %v instead.", expect, get)
		return
	}

	input = "[1]"
	expect = false
	get = IsValidFormat(input)
	if get != expect {
		t.Errorf("Expect to get %v, but get %v instead.", expect, get)
		return
	}

	input = "[0,5]"
	expect = true
	get = IsValidFormat(input)
	if get != expect {
		t.Errorf("Expect to get %v, but get %v instead.", expect, get)
		return
	}

	input = "(1,5]"
	expect = true
	get = IsValidFormat(input)
	if get != expect {
		t.Errorf("Expect to get %v, but get %v instead.", expect, get)
		return
	}

	input = "[1,5)"
	expect = true
	get = IsValidFormat(input)
	if get != expect {
		t.Errorf("Expect to get %v, but get %v instead.", expect, get)
		return
	}

	input = "(1,5)"
	expect = true
	get = IsValidFormat(input)
	if get != expect {
		t.Errorf("Expect to get %v, but get %v instead.", expect, get)
		return
	}

	input = "(1,100)"
	expect = true
	get = IsValidFormat(input)
	if get != expect {
		t.Errorf("Expect to get %v, but get %v instead.", expect, get)
		return
	}

	input = "(100,10000)"
	expect = true
	get = IsValidFormat(input)
	if get != expect {
		t.Errorf("Expect to get %v, but get %v instead.", expect, get)
		return
	}
}
