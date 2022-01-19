package internal

import "time"

type FrontPart struct {
	Date      string
	Title     string
	Author    string
	Thumbnail string
	Link      string
	Tags      []string
}

type Article struct {
	Front   FrontPart
	Date    time.Time
	Summary string
	Opinion string
}

type Newsletter struct {
	Articles []Article
}
