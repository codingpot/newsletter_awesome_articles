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
	mockTags := "#greeting #mock #oh-my"

	tags := ParseHashtags(mockTags)
	assert.Equal(t, 3, len(tags), "wrong length")
	assert.Equal(t, "#greeting", tags[0], "first item is wrong")
	assert.Equal(t, "#mock", tags[1], "second item is wrong")
	assert.Equal(t, "#oh-my", tags[2], "third item is wrong")
}

func TestParsingMarkdown(t *testing.T) {
	mockArticle := `
---
date: 2022-01-17 15:34
author: 박찬성
title: 첫 아티클
thumbnail: https://github.com/codingpot/newsletter_awesome_articles/blob/main/assets/overview.png 
link: https://github.com/codingpot/newsletter_awesome_articles 
summary: Coding Pot Newsletter Platform
opinion: Looks amazing!
hashtags: #greeting #mock #oh-my
---
	`

	article := ParseArticle(mockArticle)
	assert.Equal(t, 2022, article.Date.Year(), "wrong date")
	assert.Equal(t, "박찬성", article.Author, "wrong author")
	assert.Equal(t, "첫 아티클", article.Title, "wrong title")
	assert.Equal(t, "https://github.com/codingpot/newsletter_awesome_articles/blob/main/assets/overview.png", article.Thumbnail, "wrong thumbnail")
	assert.Equal(t, "https://github.com/codingpot/newsletter_awesome_articles", article.Link, "wrong link")
	assert.Equal(t, "Coding Pot Newsletter Platform", article.Summary, "wrong summary")
	assert.Equal(t, "Looks amazing!", article.Opinion, "wrong opinion")
	assert.Equal(t, 3, len(article.Tags), "wrong length")
	assert.Equal(t, "#greeting", article.Tags[0], "first item is wrong")
	assert.Equal(t, "#mock", article.Tags[1], "second item is wrong")
	assert.Equal(t, "#oh-my", article.Tags[2], "third item is wrong")
}
