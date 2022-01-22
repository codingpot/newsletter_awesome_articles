/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/codingpot/newsletter_awesome_articles/internal"
	"github.com/spf13/cobra"
)

// currentCmd represents the current command
var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// - how to get argument
		// topk, _ := cmd.Flags().GetInt("topk")

		// article := internal.Article{}
		// 1. get YAML files in current directory
		filenames := internal.GetListOfFilesAt("../current", ".yaml")
		fmt.Println(filenames)
		// 2. open each files & parse them to Article
		articles := []internal.Article{}
		for _, filename := range filenames {
			buf, _ := ioutil.ReadFile(filename)
			articles = append(articles, internal.ParseArticle(string(buf)))
		}
		articleTuples := internal.ZipArticleTuples(articles)
		fmt.Println(articleTuples)
		// 3. fill out template
		email := internal.Email{
			Title:       "코딩맛집 뉴스레터",
			FooterTitle: "코딩맛집",
			Header: internal.Head{
				Title:                 "코딩맛집에서 발행하는 뉴스레터 입니다",
				Description:           "이것은 뭥미?",
				ImageURL:              "http://...",
				CommunityLink:         "https://discord.gg/HGPnfzDdkG",
				CommunityLinkBtnTitle: "커뮤니티 둘러보고 가입하러 가기",
			},
			FirstSection: internal.Section{
				Title: "이거슨.... 첫 번째 섹션 타이틀",
			},
			ArticleTuples: articleTuples,
		}
		fmt.Println(email)

		f, err := os.Create("output.html")
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()

		// 4. send email
		subject := "Get latest Tech News directly to your inbox"
		receiver := "deep.diver.csp@gmail.com"
		r := internal.NewRequest("codingpot.newsletter@gmail.com", "fkgickspzlxibxkm", []string{receiver}, subject)
		r.Send("../templates", email)

		// 5. archive
		archive_dest, _ := filepath.Abs("../archive")
		archive_destinations := internal.MoveFiles(filenames, archive_dest)

		tag_dest, _ := filepath.Abs("../tags")
		for i, archive_destination := range archive_destinations {
			internal.RecordArticleByTags(articles[i], tag_dest, archive_destination)
		}
	},
}

func init() {
	publishCmd.AddCommand(currentCmd)
}
