package table

import (
	"bufio"
	"bytes"
	"log"

	"github.com/olekukonko/tablewriter"
)

func Tabletest() bytes.Buffer {

	data := [][]string{
		[]string{"1/1/2014", "中文支持", "f02233", "$10.98"},
		[]string{"1/1/2014", "中文支持", "2233", "$54.95"},
		[]string{"1/4/2014", "February Hosting", "2233", "$51.00"},
		[]string{"1/4/2014", "February Extra Bandwidth", "2233", "$30.00"},
	}

	var buf bytes.Buffer

	testout := bufio.NewWriter(&buf)
	table := tablewriter.NewWriter(testout)
	table.SetHeader([]string{"Date", "Description", "CV2", "header"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data) // Add Bulk Data
	table.Render()

	if err := testout.Flush(); err != nil {
		log.Printf("++++ %v", err)
	}
	log.Printf("++++ %v", buf.String())
	// dingtalk.SendToDingTalk("robot" + buf.String())
	return buf
}

func TableTestOtherFormat() {
	data := [][]string{
		[]string{"1/1/2014", "Domain name", "2233", "$10.98"},
		[]string{"1/1/2014", "January Hosting", "2233", "$54.95"},
		[]string{"1/4/2014", "因噎废食", "2233", "$51.00"},
		[]string{"1/4/2014", "February Extra Bandwidth hahahhahah ahah a ", "2233", "$30.00"},
	}

	var buf bytes.Buffer

	testout := bufio.NewWriter(&buf)
	table := tablewriter.NewWriter(testout)
	table.SetHeader([]string{"Date", "Description", "CV2", "Amount"})
	table.SetFooter([]string{"", "", "Total", "$146.93"}) // Add Footer
	table.SetBorder(false)                                // Set Border to false
	table.AppendBulk(data)                                // Add Bulk Data
	table.Render()
	if err := testout.Flush(); err != nil {
		log.Printf("++++ %v", err)
	}
	log.Printf("++++\n%v\n", buf.String())
	// dingtalk.SendToDingTalk("robot\n" + buf.String())

}
