package markdown

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/atsushinee/go-markdown-generator/doc"
)

type MarkDownTemplate struct {
	Doc *doc.MarkDownDoc

	CostInfo  CostInfoDo
	BillInfo  BillInfoDo
	MinerList []string
	Indicator []string
	MinerMap  map[string]map[string]string //miner-> indicator -> 具体数值
}

type MinerIndicatorInfo struct {
}

func NewMarkDownTemplate() (*MarkDownTemplate, error) {
	book := doc.NewMarkDown()

	//从json文件中获取到指标信息
	mkd := &MarkDownTemplate{
		Doc: book,
	}

	mkd.MinerMap = make(map[string]map[string]string)
	// mkd.MinerMap["f02301"] = make(map[string]string)

	var result map[string]string
	err := json.Unmarshal([]byte(TemplateData), &result)
	if err != nil {
		return nil, err
	}

	//todo 增加config 读入方法

	// mkd.MinerMap[""] = result

	for i := range result {
		mkd.Indicator = append(mkd.Indicator, i)
	}

	mkd.MinerMap["f02301"] = result
	return mkd, nil
}

// func (Mt *MarkDownTemplate) Exmaple() {
// 	code, _ := ioutil.ReadFile("main.go")
// 	book := doc.NewMarkDown()
// 	book.WriteTitle("Go-MarkDownDoc-Generator", doc.LevelTitle).
// 		WriteLines(2)

// 	book.WriteMultiCode(string(code), "go")

// 	book.WriteTitle("Author", doc.LevelNormal).
// 		WriteCodeLine("lichun")

// 	book.WriteTitle("Website", doc.LevelNormal)
// 	book.WriteLinkLine("lichunorz", "https://lichunorz.com")

// 	t := doc.NewTable(4, 4)
// 	t.SetTitle(0, "Version")
// 	t.SetTitle(1, "Date")
// 	t.SetTitle(2, "Creator")
// 	t.SetTitle(3, "Remarks")
// 	t.SetContent(0, 0, "v1")
// 	t.SetContent(0, 1, "2019-11-21")
// 	t.SetContent(0, 2, "Lee")
// 	t.SetContent(0, 3, "无")
// 	book.WriteTable(t)
// 	err := book.Export("harrytest.md")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func (Mt *MarkDownTemplate) MinerFeeInfo() string {

	book := Mt.Doc
	book.WriteTitle("数据统计时间:", doc.LevelTitle).
		WriteLines(2)

	t := doc.NewTable(len(Mt.Indicator)+1, len(Mt.MinerList)+1)
	//todo 可以优化 自定义
	for i := 0; i <= len(Mt.MinerList); i++ {

		if i == 0 {
			t.SetTitle(i, "费用指标/对象")
		} else {
			t.SetTitle(i, Mt.MinerList[i-1])

		}
	}

	for index, value := range Mt.Indicator {
		t.SetContent(index, 0, value)
		for mindex, mvalue := range Mt.MinerList {
			t.SetContent(index, mindex+1, Mt.MinerMap[mvalue][value])
		}
	}
	book.WriteTable(t)

	fmt.Println(book.String())
	err := book.Export("harrytestagain.md")
	if err != nil {
		log.Fatal(err)
	}

	return book.String()
}

//
func (Mt *MarkDownTemplate) CalculateOrderInfo() {

	//数据库中按照时间去查
	// Mt.MinerMap["f02301"]["昨日出块奖励"] =

	// 昨日出块奖励
}
