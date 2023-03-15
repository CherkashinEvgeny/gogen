package gen

const (
	Break    = keywordsRenderer("break")
	Continue = keywordsRenderer("continue")
	Return   = keywordsRenderer("return")
)

var _ Code = keywordsRenderer("")

type keywordsRenderer string

func (r keywordsRenderer) GetContext() Code {
	return nil
}

func (r keywordsRenderer) SetContext(_ Code) {
}

func (r keywordsRenderer) Render(w Writer) {
	w.Write(string(r))
}
