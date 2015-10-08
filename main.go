package main

import (
	"os"
	"fmt"
	"io/ioutil"

	"github.com/BurntSushi/toml"
	"github.com/codegangsta/cli"
	"github.com/mitchellh/go-homedir"
	"github.com/russross/blackfriday"
)

type Configuration struct {
	Username string
	Token string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var config Configuration

	app := cli.NewApp()
	app.Name = "Medium Cli"
	app.Version = "0.0.1"
	app.Usage = "Publish to Medium.com using the command line"

	app.Flags = []cli.Flag {
		cli.StringFlag {
			Name: "status, s",
			Value: "publish",
			Usage: "Status of the post",
		},
		cli.StringFlag {
			Name: "markdown, m",
			Usage: "Markdown file to post to Medium",
		},
	}

	app.Action = func(c *cli.Context) {
		status := c.String("status")
		markdownFile := c.String("markdown")

		// check post status
		if status != "publish" {
			println(status + " is not a supported status")
			return
		}

		// get post as markdown file
		data, err := ioutil.ReadFile(markdownFile);
		check(err)

		markdown := blackfriday.MarkdownCommon([]byte(data))

		fmt.Println("Markdown: " + string(markdown))

		dir, err := homedir.Dir()
		check(err)

		println("Get config file from " + dir)

		if _, err := toml.DecodeFile(dir + "/medium-cli-conf.toml", &config); err != nil {
			fmt.Println(err)
			return
		}

		println(status + " something")
		fmt.Println("Found token: " + config.Token)
	}

	app.Run(os.Args)
}
