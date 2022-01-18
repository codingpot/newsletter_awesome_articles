package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewArticle(t *testing.T) {
	article := Article{
		Title:     "My Article",
		Author:    "Chansung",
		Thumbnail: "https://github.com/codingpot/newsletter_awesome_articles/blob/main/assets/overview.png",
		Link:      "https://github.com/codingpot/newsletter_awesome_articles",
		Summary:   "Coding Pot Newsletter Platform",
		Opinion:   "Looks amazing!",
		Tags:      []string{"blog", "oss"},
	}

	assert.Equal(t, "My Article", article.Title, "Wrong Title")
}

func TestNewNewsletter(t *testing.T) {
	article1 := Article{
		Title: "My Article1",
	}

	article2 := Article{
		Title: "My Article2",
	}

	newsletter := Newsletter{
		Articles: []Article{article1, article2},
	}

	assert.Equal(t, "My Article1", newsletter.Articles[0].Title, "Wrong Number")
	assert.Equal(t, "My Article2", newsletter.Articles[1].Title, "Wrong Number")
}
