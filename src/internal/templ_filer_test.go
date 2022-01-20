package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeaderTemplate(t *testing.T) {
	// 	mockArticle := `
	// ---
	// date: 2022-01-17 15:34
	// author: 박찬성
	// title: 첫 아티클
	// thumbnail: https://github.com/codingpot/newsletter_awesome_articles/blob/main/assets/overview.png
	// link: https://github.com/codingpot/newsletter_awesome_articles
	// tags: ["greeting", "mock", "oh-my"]
	// ---
	// summary: Coding Pot Newsletter Platform
	// opinion: Looks amazing!
	// `

	// 	article := ParseArticle(mockArticle)
	// mockData := MockData{
	// 	HeaderImage: "hello world",
	// }

	// tmpl, _ := template.ParseFiles("../../templates/head_section.gohtml")
	// tmpl.Execute(os.Stdout, mockData)

	// // tmpl.Execute("header", mockData)
	TT()
}

func TestAllTemplates(t *testing.T) {
	templateFilenames := GetTemplatesInDir("../../templates")
	assert.Equal(t, 5, len(templateFilenames), "wrong length")

	// tmpl := template.Must(template.ParseFiles(templateFilenames...))
}
