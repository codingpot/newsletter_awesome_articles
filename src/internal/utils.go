package internal

import (
	"bufio"
	"regexp"
	"strings"
	"time"
)

func ParseDate(date string) (time.Time, error) {
	layout := "2006-01-02 15:04"
	parsedDate, err := time.Parse(layout, date)
	return parsedDate, err
}

func ParseHashtags(tags string) []string {
	tags_a := strings.Split(tags, " ")
	return tags_a
}

func ParseArticle(article string) Article {
	newArticle := Article{}

	r, _ := regexp.Compile(`---*([\S\s]+)*---`)
	match := r.FindString(article)
	trimmed := strings.ReplaceAll(match, "---", "")
	trimmed = strings.Trim(trimmed, "\n")
	trimmed = strings.Trim(trimmed, " ")

	scanner := bufio.NewScanner(strings.NewReader(trimmed))
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, ":")
		category := strings.Trim(tokens[0], " ")
		content := strings.Trim(strings.Join(tokens[1:], ":"), " ")

		switch category {
		case "title":
			newArticle.Title = content
		case "date":
			newArticle.Date, _ = ParseDate(content)
		case "author":
			newArticle.Author = content
		case "thumbnail":
			newArticle.Thumbnail = content
		case "link":
			newArticle.Link = content
		case "summary":
			newArticle.Summary = content
		case "opinion":
			newArticle.Opinion = content
		case "hashtags":
			newArticle.Tags = ParseHashtags(content)
		}
	}

	return newArticle
}
