package calculator

func Add(a float64, b float64) float64 { return a + b }
func Sub(a float64, b float64) float64 { return a - b }
func Mul(a float64, b float64) float64 { return a * b }
func Div(a float64, b float64) float64 { return a / b }

type OpFuncType func(float64, float64) float64

var OpMap = map[string]OpFuncType{
	"+": Add,
	"-": Sub,
	"*": Mul,
	"/": Div,
}
