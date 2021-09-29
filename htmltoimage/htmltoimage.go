package htmltoimage

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	wk "github.com/shezadkhan137/go-wkhtmltoimage"
)

var templ *template.Template

func TestHtmlToImage(testString []byte, filename string) {
	templ, err := templ.ParseGlob("htmltoimage/templates/*.html")
	if err != nil {
		fmt.Println("failed to get template err:", err)
	}
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{
		Flags:     htmlFlags,
		Generator: `  <meta charset="utf-8"/`, // `  <meta name="GENERATOR" content="github.com/gomarkdown/markdown markdown processor for Go`
	}
	renderer := html.NewRenderer(opts)
	html := markdown.ToHTML(testString, nil, renderer)
	_ = html
	wk.Init()
	defer wk.Destroy()

	converter, err := wk.NewConverter(
		&wk.Config{
			Quality:          100,
			Fmt:              "png",
			EnableJavascript: false,
		})
	if err != nil {
		log.Fatal(err)
	}

	// testChString := "<html><meta charset=\"utf-8\"/><body><p>This is some html 中文</p></body></html>"

	var buf bytes.Buffer

	testout := bufio.NewWriter(&buf)
	templ.ExecuteTemplate(testout, "template.html", struct {
		IndicatorInfo []MinerIndicatorInfo
		CreateTime    string
	}{
		IndicatorInfo: []MinerIndicatorInfo{
			{
				Indicator: "指标1",
				F02301:    "0.004 FIL",
				F03223:    "0.004 FIL",
				F0143858:  "0.004 FIL",
				F0240185:  "0.004 FIL",
			},
			{
				Indicator: "指标1",
				F02301:    "0.004 FIL",
				F03223:    "0.004 FIL",
				F0143858:  "0.004 FIL",
				F0240185:  "0.004 FIL",
			},
			{
				Indicator: "指标1",
				F02301:    "0.004 FIL",
				F03223:    "0.004 FIL",
				F0143858:  "0.004 FIL",
				F0240185:  "0.004 FIL",
			},
			{
				Indicator: "指标1",
				F02301:    "0.004 FIL",
				F03223:    "0.004 FIL",
				F0143858:  "0.004 FIL",
				F0240185:  "0.004 FIL",
			},
			{
				Indicator: "指标1",
				F02301:    "0.004 FIL",
				F03223:    "0.004 FIL",
				F0143858:  "0.004 FIL",
				F0240185:  "0.004 FIL",
			},
			{
				Indicator: "指标1",
				F02301:    "0.004 FIL",
				F03223:    "0.004 FIL",
				F0143858:  "0.004 FIL",
				F0240185:  "0.004 FIL",
			},
			{
				Indicator: "指标1",
				F02301:    "0.004 FIL",
				F03223:    "0.004 FIL",
				F0143858:  "0.004 FIL",
				F0240185:  "0.004 FIL",
			},
			{
				Indicator: "指标1",
				F02301:    "0.004 FIL",
				F03223:    "0.004 FIL",
				F0143858:  "0.004 FIL",
				F0240185:  "0.004 FIL",
			},
		},
		CreateTime: "创建时间: 2021-09-14",
	},
	)
	testout.Flush()

	outFile, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	err = converter.Run(buf.String(), outFile)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("html.......output :%v\n", buf.String())
	//发送消息
	// dingtalk.SendToDingTalk("robot\n" + buf.String())

}

type PersonData struct {
	Id         int
	Username   string
	Password   string
	Created_at string
	Updated_at string
	AdminD     int
	Miner4     string
}

type TableHeader struct {
	TableItemDesc string
	Miner1        string
	Miner2        string
	Miner3        string
	Miner4        string
}

//miner info
type MinerIndicatorInfo struct {
	Indicator string
	F02301    string
	F03223    string
	F0143858  string
	F0240185  string
}
