package emailserve

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

// NewInfo Returns a new instance of type Info with the given parameters: host, port, user name, and password
func NewInfo(host string, port int, username, password string) *Info {
	return &Info{host: host, port: port, username: username, password: password}
}

// SendEmail send email with type Info
func SendEmail(info *Info) error {
	return DialAndSend(info.host, info.port, info.username, info.password, info.message)
}

// DialAndSend send email with params: host, port, user name, password and message
func DialAndSend(host string, port int, username string, password string, message *gomail.Message) error {
	d := gomail.NewDialer(host, port, username, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return d.DialAndSend(message)
}
