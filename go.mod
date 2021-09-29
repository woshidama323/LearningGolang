module github.com/woshidama323/LearningGolang

go 1.16

require (
	github.com/Techbinator/go-table-image v0.0.0-20190913132030-9b3a8fdf94e8
	github.com/atsushinee/go-markdown-generator v0.0.0-20191121114853-83f9e1f68504
	github.com/aws/aws-sdk-go v1.32.11
	github.com/filecoin-project/go-address v0.0.5
	github.com/filecoin-project/go-jsonrpc v0.1.5
	github.com/filecoin-project/lotus v1.11.1
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/gomarkdown/markdown v0.0.0-20210914093620-23ec7da3bafc
	github.com/gorilla/mux v1.7.4
	github.com/multiformats/go-multiaddr v0.3.3
	github.com/olekukonko/tablewriter v0.0.0-20170122224234-a0225b3f23b5
	github.com/shezadkhan137/go-wkhtmltoimage v0.0.0-20191029041329-db2a8e59cfaf
	github.com/urfave/cli/v2 v2.3.0
	github.com/woshidama323/go-wkhtmltoimage v0.0.0-20210922071421-0f6eda71acf4
	github.com/xuri/excelize/v2 v2.4.1
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
)

replace github.com/Techbinator/go-table-image => ./extern/go-table-image
