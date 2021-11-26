package model

type Machine struct {
	ServerName string
	SafeTemp int
	Label []string
}

func NewMachine() *Machine {
	return &Machine{}
}

func (m *Machine) SetServerName(servername string) {
	m.ServerName = servername
}

func (m *Machine) SetSafeTemp(temp int) {
	m.SafeTemp = temp
}