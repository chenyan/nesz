package core

const (
	FC uint8 = 1 << iota // carry
	FZ                   // zero
	FI                   // interrupt
	FD                   // decimal
	FB                   // break
	_                    //  _
	FV                   // overflow
	FN                   // negative
)

type CPU struct {
	Acc, X, Y, Sp, P uint8
	Pc               uint16

	M *Memory

	insts map[uint8]Instruction
}

func (c *CPU) Execute() {
}
