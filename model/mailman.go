package model

type MailMan struct {
	Sender string
	Receiver string
	Host string
	Port int
	Password string
	*Email
}

func NewMailMan() *MailMan {
	return &MailMan{}
}

func (m *MailMan) SetSender(sender string) {
	m.Sender = sender
}

func (m *MailMan) SetReceiver(receiver string) {
	m.Receiver = receiver
}

func (m *MailMan) SetEmail(email *Email) {
	m.Email = email
}

func (m *MailMan) SetHost(host string) {
	m.Host = host
}

func (m *MailMan) SetPort(port int) {
	m.Port = port
}

func (m *MailMan) SetPassword(password string) {
	m.Password = password
}

