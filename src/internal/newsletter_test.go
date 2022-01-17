package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewArticle(t *testing.T) {
	article := Article{
		Date:      "2022-01-17",
		Title:     "My Article",
		Author:    "Chansung",
		Thumbnail: "https://github.com/codingpot/newsletter_awesome_articles/blob/main/assets/overview.png",
		Link:      "https://github.com/codingpot/newsletter_awesome_articles",
		Summary:   "Coding Pot Newsletter Platform",
		Opinion:   "Looks amazing!",
		Tags:      []string{"blog", "oss"},
	}

	assert.Equal(t, "2022-01-17", article.Date, "기대값과 결과값이 다릅니다.")
}
