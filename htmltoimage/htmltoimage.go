package htmltoimage

import (
	"log"
	"os"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	wk "github.com/shezadkhan137/go-wkhtmltoimage"
	"github.com/woshidama323/LearningGolang/dingtalk"
)

func TestHtmlToImage(testString, filename string) {
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)
	html := markdown.ToHTML([]byte(testString), nil, renderer)

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

	// testString := "<html><body><p>This is some html</p></body></html>"

	outFile, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	err = converter.Run(string(html), outFile)
	if err != nil {
		log.Fatal(err)
	}

	//发送消息
	dingtalk.SendToDingTalk("robot\n" + buf.String())

}
