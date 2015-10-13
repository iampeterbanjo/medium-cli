package main

import (
//	"os"
	"fmt"
//	"io/ioutil"
//	"log"

	"github.com/docopt/docopt-go"

//	"github.com/BurntSushi/toml"
//	"github.com/codegangsta/cli"
//	"github.com/mitchellh/go-homedir"
//	"github.com/russross/blackfriday"
//
//	medium "github.com/Medium/medium-sdk-go"
)

type Configuration struct {
	Token string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}


func main() {
	usage := `Not a serious example.
Usage:
  calculator_example <value> ( ( + | - | * | / ) <value> )...
  calculator_example <function> <value> [( , <value> )]...
  calculator_example (-h | --help)
Examples:
  calculator_example 1 + 2 + 3 + 4 + 5
  calculator_example 1 + 2 '*' 3 / 4 - 5    # note quotes around '*'
  calculator_example sum 10 , 20 , 30 , 40
Options:
  -h, --help
`
	arguments, _ := docopt.Parse(usage, nil, true, "", false)
	fmt.Println(arguments)
	fmt.Println(arguments["value"])
	fmt.Println(arguments["function"])

//	var config Configuration
//
//	app := cli.NewApp()
//	app.Name = "Medium Cli"
//	app.Version = "0.0.1"
//	app.Usage = "Publish to Medium.com using the command line"
//
//	app.Flags = []cli.Flag {
//		cli.StringFlag {
//			Name: "status, s",
//			Value: "publish",
//			Usage: "Status of the post",
//		},
//		cli.StringFlag {
//			Name: "title, t",
//			Usage: "Post tile",
//		},
//		cli.StringFlag {
//			Name: "markdown, m",
//			Usage: "Markdown file to post to Medium",
//		},
//	}
//
//	app.Action = func(c *cli.Context) {
//		status := c.String("status")
//		title := c.String("title")
//		markdownFile := c.String("markdown")
//
//		// check post status
//		if status != "publish" {
//			println(status + " is not a supported status")
//			return
//		}
//
//		if markdownFile == "" {
//			println("Please specify a markdown file")
//			return
//		}
//
//		if title == "" {
//			println("Please specify a title")
//			return
//		}
//
//		// get post as markdown file
//		data, err := ioutil.ReadFile(markdownFile);
//		check(err)
//
//		markdown := string(blackfriday.MarkdownCommon([]byte(data)))
//
//		fmt.Println("Markdown: " + markdown)
//
//		dir, err := homedir.Dir()
//		check(err)
//
//		println("Get config file from " + dir)
//
//		if _, err := toml.DecodeFile(dir + "/medium-cli-conf.toml", &config); err != nil {
//			fmt.Println(err)
//			return
//		}
//
//		println(status + " something")
//		fmt.Println("Found token: " + config.Token)
//
//		mediumClient := medium.NewClientWithAccessToken(config.Token)
//
//		mediumUser, err := mediumClient.GetUser()
//		check(err)
//
//		post, err := mediumClient.CreatePost(medium.CreatePostOptions{
//			UserID: 				mediumUser.ID,
//			Title:					title,
//			Content:				markdown,
//			ContentFormat:	medium.ContentFormatHTML,
//			PublishStatus:	medium.PublishStatusDraft,
//		})
//
//		check(err)
//
//		log.Println(mediumUser, post)
//	}

//	app.Run(os.Args)
}
