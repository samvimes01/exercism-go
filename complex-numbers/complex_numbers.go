package complexnumbers

import "math"

// Define the Number type here.
type Number struct {
	real, imaginary float64
}

func (n Number) Real() float64 {
	return n.real
}

func (n Number) Imaginary() float64 {
	return n.imaginary
}

func (n1 Number) Add(n2 Number) Number {
	return Number{n1.real + n2.real, n1.imaginary + n2.imaginary}
}

func (n1 Number) Subtract(n2 Number) Number {
	return Number{n1.real - n2.real, n1.imaginary - n2.imaginary}
}

func (n1 Number) Multiply(n2 Number) Number {
	return Number{n1.real*n2.real - n1.imaginary*n2.imaginary, n1.imaginary*n2.real + n1.real*n2.imaginary}
}

func (n Number) Times(factor float64) Number {
	return Number{n.real * factor, n.imaginary * factor}
}

func (n1 Number) Divide(n2 Number) Number {

	den := n2.real*n2.real + n2.imaginary*n2.imaginary

	return Number{(n1.real*n2.real + n1.imaginary*n2.imaginary) / den, (n1.imaginary*n2.real - n1.real*n2.imaginary) / den}

}

func (n Number) Conjugate() Number {

	return Number{n.real, -n.imaginary}

}

func (n Number) Abs() float64 {

	return math.Sqrt(n.real*n.real + n.imaginary*n.imaginary)

}

func (n Number) Exp() Number {

	return Number{math.Exp(n.real) * math.Cos(n.imaginary), math.Exp(n.real) * math.Sin(n.imaginary)}

}
