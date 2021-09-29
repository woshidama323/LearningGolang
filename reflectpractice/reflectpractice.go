package reflectpractice

import (
	"fmt"
	"reflect"
)

//已知一个struct 现在想通过string 动态找到对应的field tag 还有value

func TestGetField() {

	forref := &MinerIndicatorInfo{
		Indicator: "why",
	}

	v := reflect.ValueOf(forref).Elem().FieldByName("F02301")
	fmt.Printf(".....v:%v\n", forref)
	if v.IsValid() {
		v.SetString("harry")
	}
	fmt.Printf(".....forref:%v\n", forref)

}

type MinerIndicatorInfo struct {
	Indicator string
	F02301    string
	F03223    string
	F0143858  string
	F0240185  string
}

type TemplVariable struct {
	CreateTime    string
	IndicatorInfo []MinerIndicatorInfo
}
