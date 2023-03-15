package gen

type Code interface {
	GetContext() Code
	SetContext(ctx Code)
	Render(w Writer)
}

type Sequence interface {
	Code
	Len() int
	At(i int) Code
	Add(items ...Code)
}
