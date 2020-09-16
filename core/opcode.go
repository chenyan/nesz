package core

type Instruction struct {
	Opcode uint8
	Nmae   string
	F      func() uint16
}

func (c *CPU) adc() uint16 {

}
