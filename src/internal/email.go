package internal

type Head struct {
	Title                 string
	Description           string
	ImageURL              string
	CommunityLink         string
	CommunityLinkBtnTitle string
}

type Section struct {
	Title string
}

type ArticleTuple struct {
	Article1 Article
	Article2 Article
}

type Email struct {
	Title         string
	FooterTitle   string
	Header        Head
	FirstSection  Section
	ArticleTuples []ArticleTuple // has to be
}

// type Request struct {
// 	from    string
// 	to      []string
// 	subject string
// 	body    string
// }

// const (
// 	MIME = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
// )

// func NewRequest(to []string, subject string) *Request {
// 	return &Request{
// 		to:      to,
// 		subject: subject,
// 	}
// }

// func (r *Request) parseTemplate(fileName string, data interface{}) error {
// 	t, err := template.ParseFiles(fileName)
// 	if err != nil {
// 		return err
// 	}
// 	buffer := new(bytes.Buffer)
// 	if err = t.Execute(buffer, data); err != nil {
// 		return err
// 	}
// 	r.body = buffer.String()
// 	return nil
// }

// func (r *Request) sendMail() bool {
// 	body := "To: " + r.to[0] + "\r\nSubject: " + r.subject + "\r\n" + MIME + "\r\n" + r.body
// 	SMTP := fmt.Sprintf("%s:%d", config.Server, config.Port)
// 	if err := smtp.SendMail(SMTP, smtp.PlainAuth("", config.Email, config.Password, config.Server), config.Email, r.to, []byte(body)); err != nil {
// 		return false
// 	}
// 	return true
// }

// func (r *Request) Send(templateName string, items interface{}) {
// 	err := r.parseTemplate(templateName, items)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if ok := r.sendMail(); ok {
// 		log.Printf("Email has been sent to %s\n", r.to)
// 	} else {
// 		log.Printf("Failed to send the email to %s\n", r.to)
// 	}
// }
