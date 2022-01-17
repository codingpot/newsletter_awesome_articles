package internal

import (
	"time"
)

func ParseDate(date string) (time.Time, error) {
	layout := "2006-01-02 15:04"
	parsedDate, err := time.Parse(layout, date)
	return parsedDate, err
}

func ParseHashtags(tags string) []string {
	return []string{}
}

func ParseArticle(article string) Article {
	return Article{}
}
