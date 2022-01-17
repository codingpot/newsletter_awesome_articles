package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseDateFormat(t *testing.T) {
	mockDate := "2022-01-17 15:34"
	parsedDate, err := ParseDate(mockDate)

	assert.Equal(t, nil, err, err)
	assert.Equal(t, 2022, parsedDate.Year(), "Year doesn't match")
	assert.Equal(t, "January", parsedDate.Month().String(), "Month string doesn't match")
	assert.Equal(t, 1, int(parsedDate.Month()), "Month integer doesn't match")
	assert.Equal(t, 17, parsedDate.Day(), "Day doesn't match")
	assert.Equal(t, 15, parsedDate.Hour(), "Hour doesn't match")
	assert.Equal(t, 34, parsedDate.Minute(), "Hour doesn't match")

	mockDate = "2022/01/17"
	_, err = ParseDate(mockDate)
	assert.NotEqual(t, nil, err, err)

	mockDate = "2022-01-17 15:34:02"
	_, err = ParseDate(mockDate)
	assert.NotEqual(t, nil, err, err)
}

func TestParseHashtags(t *testing.T) {

}

func TestParsingMarkdown(t *testing.T) {
	mockArticle := `
	date: 2022-01-17 15:34
	author: 박찬성
	title: 첫 아티클
	thumbnail: https://github.com/codingpot/newsletter_awesome_articles/blob/main/assets/overview.png 
	link: https://github.com/codingpot/newsletter_awesome_articles 
	summary: Coding Pot Newsletter Platform
	opinion: Looks amazing!
	hashtags: #greeting #mock #oh-my
	`
}
