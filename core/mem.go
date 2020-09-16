package core

type Memory struct {
}

func (m *Memory) Read(addr uint16) uint8 {
	return 0
}

func (m *Memory) Write(addr uint16, v uint8) {}
