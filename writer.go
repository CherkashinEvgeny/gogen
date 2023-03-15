package gen

type Writer interface {
	Write(str string)
	Br()
	Indent(n int)
}
