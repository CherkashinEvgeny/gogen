package gen

import "strings"

type Builder struct {
	indent   int
	indented bool
	sb       strings.Builder
}

func (b *Builder) Write(str string) {
	if strings.Count(str, "\n") == 0 {
		b.writeLine(str)
		return
	}
	lines := strings.Split(str, "\n")
	for i, line := range lines {
		if i != 0 {
			b.Br()
		}
		b.writeLine(line)
	}
	return
}

func (b *Builder) writeLine(line string) {
	if line == "" {
		return
	}
	if !b.indented {
		b.writeIndent()
	}
	b.sb.WriteString(line)
}

func (b *Builder) writeIndent() {
	for i := 0; i < b.indent; i++ {
		b.sb.WriteString("\t")
	}
	b.indented = true
}

func (b *Builder) Br() {
	b.sb.WriteString("\n")
	b.indented = false
}

func (b *Builder) Indent(n int) {
	b.indent += n
	if b.indent < 0 {
		b.indent = 0
	}
}

func (b *Builder) String() string {
	return b.sb.String()
}
