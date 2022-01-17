package internal

type Article struct {
	Date      string
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
