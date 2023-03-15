package gen

import "strings"

func Stringify(code Code) string {
	b := &Builder{}
	code.Render(b)
	return b.String()
}

func RuneCount(code Code) int {
	c := &runeCounter{}
	code.Render(c)
	return c.Count()
}

type runeCounter struct {
	indent   int
	indented bool
	counter  int
}

func (c *runeCounter) Write(str string) {
	if strings.Count(str, "\n") == 0 {
		c.writeLine(str)
		return
	}
	lines := strings.Split(str, "\n")
	for i, line := range lines {
		if i != 0 {
			c.Br()
		}
		c.writeLine(line)
	}
	return
}

func (c *runeCounter) writeLine(line string) {
	if line == "" {
		return
	}
	if !c.indented {
		c.writeIndent()
	}
	c.counter += len([]rune(line))
}

func (c *runeCounter) writeIndent() {
	c.counter += c.indent
	c.indented = true
}

func (c *runeCounter) Br() {
	c.counter++
	c.indented = false
}

func (c *runeCounter) Indent(n int) {
	c.indent += n
	if c.indent < 0 {
		c.indent = 0
	}
}

func (c *runeCounter) Count() int {
	return c.counter
}

func IsMultiline(code Code) bool {
	return LinesCount(code) > 0
}

func LinesCount(code Code) int {
	c := &lineCounter{}
	code.Render(c)
	return c.Count()
}

type lineCounter struct {
	counter int
}

func (c *lineCounter) Write(str string) {
	c.counter += strings.Count(str, "\n")
	return
}

func (c *lineCounter) Br() {
	c.counter++
}

func (c *lineCounter) Indent(n int) {
}

func (c *lineCounter) Count() int {
	return c.counter
}
