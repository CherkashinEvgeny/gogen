# gogen

Code generation toolkit.

## Library api is unstable. If you need stable api, you should fork this project or find an [alternative one](#underwater-rocks).

## About The Project

Gogen - is golang code generation toolkit.

Toolkit consists of three parts:

1. [core](https://github.com/CherkashinEvgeny/gogen) provides core functions for code generation
2. [types](https://github.com/CherkashinEvgeny/gogen/types) integration with `go/types` package
3. [utils](https://github.com/CherkashinEvgeny/gogen/utils) contains functions, that do not lead directly to code
   generation, but can help keep code concise

Features:
- Reach code generation api
- Embedded formatting
- Autoimport

## Usage
Hello world:
```
import (
	"fmt"
	. "github.com/CherkashinEvgeny/gogen"
)

func main() {
	mainFunc := Func("main", Sign(In(), Out()), Lines(
		Call(
			SmartQual("fmt", "fmt", "Println"),
			Val("Hello world(:"),
		),
	))
	pkg := Pkg("", "test", Imports(), mainFunc)
	fmt.Println(Stringify(pkg))
}
```

## Underwater rocks

## Similar projects

- [dave/jennifer](https://github.com/dave/jennifer)

## License

Retry is licensed under the Apache License, Version 2.0. See [LICENSE](./LICENCE.md)
for the full license text.

## Contact

- Email: `cherkashin.evgeny.viktorovich@gmail.com`
- Telegram: `@evgeny_cherkashin`
