package brainfuck

import "io"

type Machine struct {
	code   []*Instruction
	ip     int
	memory [30000]int
	dp     int
	input  io.Reader
	output io.Writer
	buf    []byte
}

func NewMachine(code []*Instruction, in io.Reader, out io.Writer) *Machine {
	return &Machine{
		code:   code,
		input:  in,
		output: out,
		buf:    make([]byte, 1),
	}
}

func (m *Machine) Execute() {
	for m.ip < len(m.code) {
		ins := m.code[m.ip]

		switch ins.Type {
		case Plus:
			m.memory[m.dp] += ins.Argument
		case Minus:
			m.memory[m.dp] -= ins.Argument
		case Right:
			m.dp += ins.Argument
		case Left:
			m.dp -= ins.Argument
		case PutChar:
			for i := 0; i < ins.Argument; i++ {
				m.putChar()
			}
		case ReadChar:
			for i := 0; i < ins.Argument; i++ {
				m.readChar()
			}
		case JumpIfZero:
			if m.memory[m.dp] == 0 {
				m.ip = ins.Argument
				continue
			}
		case ']':
			if m.memory[m.dp] != 0 {
				m.ip = ins.Argument
				continue
			}
		}

		m.ip++
	}
}

func (m *Machine) readChar() {
	n, err := m.input.Read(m.buf)
	if err != nil {
		panic(err)
	}

	if n != 1 {
		panic("wrong number of bytes read")
	}

	m.memory[m.dp] = int(m.buf[0])
}

func (m *Machine) putChar() {
	m.buf[0] = byte(m.memory[m.dp])

	n, err := m.output.Write(m.buf)
	if err != nil {
		panic(err)
	}

	if n != 1 {
		panic("wrong number of bytes written")
	}
}
