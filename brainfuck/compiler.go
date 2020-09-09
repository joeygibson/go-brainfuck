package brainfuck

type Compiler struct {
	code         string
	codeLength   int
	position     int
	instructions []*Instruction
}

func NewCompiler(code string) *Compiler {
	return &Compiler{
		code:         code,
		codeLength:   len(code),
		instructions: []*Instruction{},
	}
}

func (c *Compiler) Compile() []*Instruction {
	loopStack := []int{}

	for c.position < c.codeLength {
		current := c.code[c.position]

		switch current {
		case '+':
			c.CompileFoldableInstruction(current, Plus)
		case '-':
			c.CompileFoldableInstruction(current, Minus)
		case '>':
			c.CompileFoldableInstruction(current, Right)
		case '<':
			c.CompileFoldableInstruction(current, Left)
		case '.':
			c.CompileFoldableInstruction(current, PutChar)
		case ',':
			c.CompileFoldableInstruction(current, ReadChar)
		case '[':
			insPos := c.EmitWithArg(JumpIfZero, 0)
			loopStack = append(loopStack, insPos)
		case ']':
			// pop position of last JumpIfZero ("[") off the stack
			openInstruction := loopStack[len(loopStack)-1]
			loopStack = loopStack[:len(loopStack)-1]

			// emit the new JumpIfNotZero ("]") instruction,
			// with correct position argument
			closeInstructionPos := c.EmitWithArg(JumpIfNotZero, openInstruction)

			// Patch the old JumpIfZero with the new position
			c.instructions[openInstruction].Argument = closeInstructionPos
		}

		c.position++
	}

	return c.instructions
}

func (c *Compiler) CompileFoldableInstruction(char byte, insType InsType) {
	count := 1

	for c.position < c.codeLength-1 && c.code[c.position+1] == char {
		count++
		c.position++
	}

	c.EmitWithArg(insType, count)
}

func (c *Compiler) EmitWithArg(insType InsType, arg int) int {
	ins := &Instruction{
		Type:     insType,
		Argument: arg,
	}

	c.instructions = append(c.instructions, ins)

	return len(c.instructions) - 1
}
