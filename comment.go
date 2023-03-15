package gen

import "strings"

var _ Code = CommentRenderer("")

type CommentRenderer string

func Comment(str string) CommentRenderer {
	return CommentRenderer(str)
}

func (r CommentRenderer) GetContext() Code {
	return nil
}

func (r CommentRenderer) SetContext(_ Code) {
}

func (r CommentRenderer) Render(w Writer) {
	lines := strings.Split(string(r), "\n")
	for i, line := range lines {
		if i != 0 {
			w.Br()
		}
		if !strings.HasPrefix(line, "//") {
			w.Write("// ")
		}
		w.Write(line)
	}
}
