package internal

type Article struct {
	Date      string
	Title     string
	Author    string
	Thumbnail string
	Link      string
	Tags      []string
	Summary   string
	Opinion   string
}

type Newsletter struct {
	Articles []Article
}
