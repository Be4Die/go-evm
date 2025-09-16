// vm/psw.go
package vm

type PSW struct {
	ip    uint16
	sp    uint8
	flags uint16
}

func NewPSW() *PSW {
	return &PSW{
		sp: 31, // Initialize SP to top of stack (31 words)
	}
}

func (p *PSW) GetIP() uint16 {
	return p.ip
}

func (p *PSW) SetIP(ip uint16) {
	p.ip = ip
}

func (p *PSW) GetSP() uint8 {
	return p.sp
}

func (p *PSW) SetSP(sp uint8) {
	p.sp = sp
}

func (p *PSW) GetFlags() uint16 {
	return p.flags
}

func (p *PSW) SetFlags(flags uint16) {
	p.flags = flags
}

func (p *PSW) SetFlag(bit uint8, value bool) {
	if value {
		p.flags |= 1 << bit
	} else {
		p.flags &^= 1 << bit
	}
}

func (p *PSW) GetFlag(bit uint8) bool {
	return (p.flags>>bit)&1 == 1
}