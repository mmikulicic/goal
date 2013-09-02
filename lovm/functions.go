package lovm

import (
	"fmt"
	"io"
)

type Sequence int

type Function struct {
	*Module
	Indent string
	Tmps   Sequence
	Labels Sequence

	Blocks []*Block
	Values map[Value]bool
	Type   FuncType
	Name   string
}

func (s *Sequence) Next() Sequence {
	res := *s
	(*s)++
	return res
}

func (fun *Function) NewBlock() *Block {
	res := NewBlock(fun)
	fun.Blocks = append(fun.Blocks, res)
	return res
}

func (fun *Function) Emitf(format string, args ...interface{}) {
	io.WriteString(fun.Writer, fun.Indent)
	fmt.Fprintf(fun.Writer, format, args...)
	io.WriteString(fun.Writer, "\n")
}

func (fun *Function) Emit() {
	for _, b := range fun.Blocks {
		b.Prepare(fun)
	}

	fun.Type.EmitDef(fun.Writer, fun.Name, func() {
		for _, b := range fun.Blocks {
			b.Emit(fun)
		}
	})
}