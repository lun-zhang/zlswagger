package main

import (
	"encoding/json"
	"flag"
	"github.com/lun-zhang/go-tool/api/com/parser"
	"io/ioutil"
	"os"
)

var (
	mainFileDir string
	title       string
	desc        string
	contactName string
)

func init() {
	flag.StringVar(&mainFileDir, "dir", "./", "main.go的目录地址，默认当前目录")
	flag.StringVar(&title, "title", "tmp", "项目标题，默认是tmp")
	flag.StringVar(&desc, "desc", "", "项目描述")
	flag.StringVar(&contactName, "cn", "your@email.com", "联系人")
	flag.Parse()
	if mainFileDir == "" || title == "" {
		panic("项目目录或标题必传")
	}
}

/*
./zlswagger \
-dir="改成你的main.go所在的目录" \
-title="示例标题" \
-desc="示例描述"
*/
func main() {
	_, a, err := parser.ParseApis(
		mainFileDir,
		true,
		true)
	if err != nil {
		panic(err)
	}

	swgSpc := parser.NewSwaggerSpec()
	swgSpc.Apis(a)
	swgSpc.Info(
		title,
		desc,
		"1",
		contactName,
	)
	swgSpc.Schemes([]string{"http"})
	swgSpc.Host("localhost")
	err = swgSpc.ParseApis()
	if nil != err {
		panic(err)
	}

	out, err := json.MarshalIndent(swgSpc.Swagger, "", "	")
	if nil != err {
		panic(err)
	}
	if err := ioutil.WriteFile("./swagger.json", out, os.ModePerm); err != nil {
		panic(err)
	}
}
