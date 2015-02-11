package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

var DefaultFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "data",
		Usage: "Body data",
	},
}

var Commands = []cli.Command{
	commandGet,
	commandPost,
	commandPut,
	commandDelete,
}

var commandGet = cli.Command{
	Name:      "get",
	ShortName: "g",
	Usage:     "Make GET request",
	Description: `
`,
	Flags:  DefaultFlags,
	Action: doGet,
}

var commandPost = cli.Command{
	Name:      "post",
	ShortName: "p",
	Usage:     "Make POST request",
	Description: `
`,
	Flags:  DefaultFlags,
	Action: doPost,
}

var commandPut = cli.Command{
	Name:  "put",
	Usage: "Make PUT request",
	Description: `
`,
	Flags:  DefaultFlags,
	Action: doPut,
}

var commandDelete = cli.Command{
	Name:      "delete",
	ShortName: "d",
	Usage:     "Make DELETE request",
	Description: `
`,
	Flags:  DefaultFlags,
	Action: doDelete,
}

func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func doGet(c *cli.Context) {
	doRequest(c, "GET")
}

func doPost(c *cli.Context) {
	doRequest(c, "POST")
}

func doPut(c *cli.Context) {
	doRequest(c, "PUT")
}

func doDelete(c *cli.Context) {
	doRequest(c, "DELETE")
}

func doRequest(ctx *cli.Context, method string) {
	url := getUrl(ctx)
	debug("url = %s", url)
	body := strings.NewReader(ctx.String("data"))
	req, _ := http.NewRequest(method, url, body)

	reqDump, _ := httputil.DumpRequestOut(req, true)
	debug("request = ", string(reqDump))

	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	resDump, _ := httputil.DumpResponse(res, true)
	debug("response = ", string(resDump))

	byteArray, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(byteArray))
}

func getUrl(ctx *cli.Context) string {
	if len(ctx.Args()) < 1 {
		log.Fatal("require URL" + string(len(ctx.Args())))
		os.Exit(1)
	}
	return ctx.Args()[0]
}
