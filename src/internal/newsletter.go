package internal

import "time"

type Article struct {
	Date      time.Time
	Title     string
	Author    string
	Thumbnail string
	Link      string
	Summary   string
	Opinion   string
	Tags      []string
}

type Newsletter struct {
	Articles []Article
}
