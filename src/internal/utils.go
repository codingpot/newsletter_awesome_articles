package internal

import (
	"bufio"
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

	rest, err := frontmatter.Parse(strings.NewReader(article), &newArticle.Front)
	if err != nil {
		fmt.Println(err)
	}

	trimmed := strings.Trim(string(rest), "\n")
	trimmed = strings.Trim(trimmed, " ")

	scanner := bufio.NewScanner(strings.NewReader(trimmed))
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, ":")
		category := strings.Trim(tokens[0], " ")
		content := strings.Trim(strings.Join(tokens[1:], ":"), " ")

		switch category {
		case "summary":
			newArticle.Summary = content
		case "opinion":
			newArticle.Opinion = content
		}
	}

	return newArticle
}
