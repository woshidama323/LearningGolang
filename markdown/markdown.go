package markdown

import (
	"io/ioutil"
	"log"

	"github.com/atsushinee/go-markdown-generator/doc"
)

type MarkDownTemplate struct {
	Doc       *doc.MarkDownDoc
	MinerList []string
	Indicator []string
}

func NewMarkDownTemplate() *MarkDownTemplate {
	book := doc.NewMarkDown()
	return &MarkDownTemplate{
		Doc: book,
	}
}

func (Mt *MarkDownTemplate) Exmaple() {
	code, _ := ioutil.ReadFile("main.go")
	book := doc.NewMarkDown()
	book.WriteTitle("Go-MarkDownDoc-Generator", doc.LevelTitle).
		WriteLines(2)

	book.WriteMultiCode(string(code), "go")

	book.WriteTitle("Author", doc.LevelNormal).
		WriteCodeLine("lichun")

	book.WriteTitle("Website", doc.LevelNormal)
	book.WriteLinkLine("lichunorz", "https://lichunorz.com")

	t := doc.NewTable(4, 4)
	t.SetTitle(0, "Version")
	t.SetTitle(1, "Date")
	t.SetTitle(2, "Creator")
	t.SetTitle(3, "Remarks")
	t.SetContent(0, 0, "v1")
	t.SetContent(0, 1, "2019-11-21")
	t.SetContent(0, 2, "Lee")
	t.SetContent(0, 3, "无")
	book.WriteTable(t)
	err := book.Export("harrytest.md")
	if err != nil {
		log.Fatal(err)
	}
}

func (Mt *MarkDownTemplate) MinerFeeInfo() {

	book := Mt.Doc
	//
	book.WriteTitle("Go-MarkDownDoc-Generator", doc.LevelTitle).
		WriteLines(2)

	book.WriteTitle("Author", doc.LevelNormal).
		WriteCodeLine("lichun")

	book.WriteTitle("Website", doc.LevelNormal)
	book.WriteLinkLine("lichunorz", "https://lichunorz.com")

	t := doc.NewTable(len(Mt.Indicator)+1, len(Mt.MinerList)+1)

	//todo 可以优化 自定义
	t.SetTitle(0, "费用指标/对象")

	for i := 1; i <= len(Mt.MinerList); i++ {
		t.SetTitle(i, Mt.MinerList[i-1])
	}

	for j := 0; j < len(Mt.Indicator); j++ {
		t.SetContent(j, 0, Mt.Indicator[j])
	}

	//用户
	for j := 0; j < len(Mt.Indicator); j++ {
		t.SetContent(j, 0, Mt.Indicator[j])
	}

	//
	//指标1 昨日出块奖励
	// t.SetContent(0, 0, "昨日出块奖励")

	// t.SetContent(0, 1, "2019-11-21")
	// t.SetContent(0, 2, "Lee")
	// t.SetContent(0, 3, "无")
	// t.SetContent(0, 4, "无")

	// //指标2 昨日出块数
	// t.SetContent(1, 0, "昨日出块数")

	// t.SetContent(1, 1, "2019-11-21")
	// t.SetContent(1, 2, "Lee")
	// t.SetContent(1, 3, "无")
	// t.SetContent(1, 4, "无")

	// //指标3 昨日25%释放
	// t.SetContent(2, 0, "昨日25%释放")

	// t.SetContent(2, 1, "2019-11-21")
	// t.SetContent(2, 2, "Lee")
	// t.SetContent(2, 3, "无")
	// t.SetContent(2, 4, "无")

	// //指标4 昨日75%锁仓
	// t.SetContent(3, 0, "v1")

	// t.SetContent(3, 1, "2019-11-21")
	// t.SetContent(3, 2, "Lee")
	// t.SetContent(3, 3, "无")
	// t.SetContent(3, 4, "无")

	// //指标5 昨日180天总线性释放
	// t.SetContent(4, 0, "昨日180天总线性释放")

	// t.SetContent(4, 1, "2019-11-21")
	// t.SetContent(4, 2, "Lee")
	// t.SetContent(4, 3, "无")
	// t.SetContent(4, 4, "无")

	// //指标6 昨日总释放
	// t.SetContent(5, 0, "昨日总释放")

	// t.SetContent(5, 1, "2019-11-21")
	// t.SetContent(5, 2, "Lee")
	// t.SetContent(5, 3, "无")
	// t.SetContent(5, 4, "无")

	// //指标7 昨日总增加算力
	// t.SetContent(6, 0, "昨日总增加算力")

	// t.SetContent(6, 1, "2019-11-21")
	// t.SetContent(6, 2, "Lee")
	// t.SetContent(6, 3, "无")
	// t.SetContent(6, 4, "无")

	// //指标8 昨日总质押费
	// t.SetContent(7, 0, "昨日总质押费")

	// t.SetContent(7, 1, "2019-11-21")
	// t.SetContent(7, 2, "Lee")
	// t.SetContent(7, 3, "无")
	// t.SetContent(7, 4, "无")

	// //指标9 SubmitWindowedPoSt 扇区
	// t.SetContent(8, 0, "SubmitWindowedPoSt 扇区")

	// t.SetContent(8, 1, "2019-11-21")
	// t.SetContent(8, 2, "Lee")
	// t.SetContent(8, 3, "无")
	// t.SetContent(8, 4, "无")

	// //指标10 矿工手续费
	// t.SetContent(9, 0, "矿工手续费")

	// t.SetContent(9, 1, "2019-11-21")
	// t.SetContent(9, 2, "Lee")
	// t.SetContent(9, 3, "无")
	// t.SetContent(9, 4, "无")

	// //指标11 GAS燃烧费
	// t.SetContent(10, 0, "GAS燃烧费")

	// t.SetContent(10, 1, "2019-11-21")
	// t.SetContent(10, 2, "Lee")
	// t.SetContent(10, 3, "无")
	// t.SetContent(10, 4, "无")

	book.WriteTable(t)
	err := book.Export("harrytestagain.md")
	if err != nil {
		log.Fatal(err)
	}
}
