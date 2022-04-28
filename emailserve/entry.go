package emailserve

import "gopkg.in/gomail.v2"

// Info email info
type Info struct {
	host     string
	port     int
	username string
	password string
	message  *gomail.Message
}

func (i *Info) Host() string {
	return i.host
}

func (i *Info) Port() int {
	return i.port
}

func (i *Info) Username() string {
	return i.username
}

func (i *Info) Password() string {
	return i.password
}

func (i *Info) Message() *gomail.Message {
	return i.message
}

// InfoBuilder email info builder
type InfoBuilder struct {
	info *Info
}

// NewInfoBuilder returns a new instance of type InfoBuilder with the empty info
func NewInfoBuilder() *InfoBuilder {
	return &InfoBuilder{
		info: &Info{},
	}
}

// Host set the host by info builder
func (b *InfoBuilder) Host(host string) *InfoBuilder {
	b.info.host = host
	return b
}

// Port set the port by info builder
func (b *InfoBuilder) Port(port int) *InfoBuilder {
	b.info.port = port
	return b
}

// Username set the username by info builder
func (b *InfoBuilder) Username(username string) *InfoBuilder {
	b.info.username = username
	return b
}

// Password set the password by info builder
func (b *InfoBuilder) Password(password string) *InfoBuilder {
	b.info.password = password
	return b
}

// Message set the message by info builder
func (b *InfoBuilder) Message(message *gomail.Message) *InfoBuilder {
	b.info.message = message
	return b
}

// Build returns a new instance of type Info
func (b *InfoBuilder) Build() *Info {
	// if b.info.Host == "" {
	// 	panic("host is required")
	// }
	// if b.info.Port == 0 {
	// 	panic("port is required")
	// }
	// if b.info.Username == "" {
	// 	panic("username is required")
	// }
	// if b.info.Password == "" {
	// 	panic("password is required")
	// }
	// if b.info.Message == nil {
	// 	b.info.Message = gomail.NewMessage()
	// }
	return b.info
}

// ToBuilder returns a new instance of type InfoBuilder with info
func (i *Info) ToBuilder() *InfoBuilder {
	return &InfoBuilder{
		info: i,
	}
}
