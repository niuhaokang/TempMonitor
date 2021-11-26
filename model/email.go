package model

type Email struct {
	Subject string
	Body string
}

func (e *Email) Update(subject, body string) {
	e.Body = body
	e.Subject = subject
}