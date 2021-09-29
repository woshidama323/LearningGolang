package ascii

import (
	"fmt"
	"reflect"
)

func TestAscii() error {
	x := struct {
		Harry string
		Test  string
	}{
		Harry: "harry",
		Test:  "test",
	}
	typeNum := reflect.TypeOf(x).NumField()
	for i := int('A'); i < int('A')+typeNum; i++ {
		fmt.Printf("test the code %c\n", i)
	}

	return nil
}
