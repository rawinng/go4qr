package utilities

import (
	"testing"
)

func TestToJson_ShouldBeSuccess(t *testing.T) {
	testJson := TestJson{
		Name: "test",
		Age:  18,
	}
	actual := ToJson(testJson)
	expect := "{\"name\":\"test\",\"age\":18}"
	if actual != expect {
		t.Errorf("ToJson() failed, actual: %s, expect: %s", actual, expect)
	}
}

func TestFromJson_ShouldBeSuccess(t *testing.T) {
	jsonString := "{\"name\":\"test\",\"age\":18}"
	expect := TestJson{
		Name: "test",
		Age:  18,
	}
	var actual TestJson
	FromJson(jsonString, &actual)
	if actual != expect {
		t.Errorf("FromJson() failed, actual: %v, expect: %v", actual, expect)
	}
}

// structure for test
type TestJson struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
