package emailserve

import (
	"testing"

	"gopkg.in/gomail.v2"
)

func TestSendEmail(t *testing.T) {
	m := gomail.NewMessage()
	m.SetHeader("From", "xxx@gmail.com")
	m.SetHeader("To", "xxx@gmail.com")
	m.SetHeader("Subject", "test")
	m.SetBody("text/html", "<h1>test</h1><p:<b>test</b></p>")
	i := NewInfoBuilder().Host("smtp.gmail.com").Port(465).Username("xxx@gmail.com").Password("xxx").Message(m).Build()
	if err := SendEmail(i); err != nil {
		t.Error(err)
	}
}
