package calculator

func Add(a float64, b float64) float64 { return a + b }
func Sub(a float64, b float64) float64 { return a - b }
func Mul(a float64, b float64) float64 { return a * b }
func Div(a float64, b float64) float64 { return a / b }

var OpMap = map[string]func(float64, float64) float64{
	"+": Add,
	"-": Sub,
	"*": Mul,
	"/": Div,
}
