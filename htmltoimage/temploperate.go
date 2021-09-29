package htmltoimage

import (
	"bufio"
	"bytes"
	"errors"
	"html/template"
	"reflect"
)

type Node struct {
	Contact_id  int
	Employer_id int
	First_name  string
	Middle_name string
	Last_name   string
}

var templateFuncs = template.FuncMap{"rangeStruct": RangeStructer}

// In the template, we use rangeStruct to turn our struct values
// into a slice we can iterate over
var htmlTemplate = `{{range .}}<tr>
{{range rangeStruct .}} <td>{{.}}</td>
{{end}}</tr>
{{end}}`

func TestTemplateForTable() bytes.Buffer {
	container := []Node{
		{1, 12, "Accipiter", "ANisus", "Nisus"},
		{2, 42, "Hello", "my", "World"},
	}

	// We create the template and register out template function
	t := template.New("t").Funcs(templateFuncs)
	t, err := t.Parse(htmlTemplate)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	testbuf := bufio.NewWriter(&buf)
	err = t.Execute(testbuf, container)
	testbuf.Flush()

	if err != nil {
		panic(err)
	}
	return buf
}

func RangeStructer(args ...interface{}) ([]interface{}, error) {

	if len(args) == 0 {
		return []interface{}{""}, errors.New("no input data")
	}

	v := reflect.ValueOf(args[0])
	if v.Kind() != reflect.Struct {
		return []interface{}{""}, errors.New("input is not a struct")
	}

	out := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		out[i] = v.Field(i).Interface()
	}
	return out, nil
}

//note method comes from https://stackoverflow.com/questions/19991124/go-template-html-iteration-to-generate-table-from-struct
