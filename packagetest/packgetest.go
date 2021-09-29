package packagetest

import (
	"fmt"

	wk "github.com/woshidama323/go-wkhtmltoimage"
)

func TestPackagetest() {
	err := wk.Init()
	if err != nil {
		fmt.Println("....,", err)
		return
	}

	defer wk.Destroy()
}
