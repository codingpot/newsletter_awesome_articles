package internal

import (
	"fmt"
	"strings"
	"time"

	"github.com/adrg/frontmatter"
)

func ParseDate(date string) (time.Time, error) {
	layout := "2006-01-02 15:04"
	parsedDate, err := time.Parse(layout, date)
	return parsedDate, err
}

func ParseArticle(article string) Article {
	newArticle := Article{}

	_, err := frontmatter.Parse(strings.NewReader(article), &newArticle)
	if err != nil {
		fmt.Println(err)
	}

	return newArticle
}
