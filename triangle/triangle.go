package triangle


// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
    // Pick values for the following identifiers used by the test program.
    NaT = iota// not a triangle
    Equ // equilateral
    Iso // isosceles
    Sca // scalene
)

// KindFromSides return a kind of a triangle.
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}
	if a + b <= c || a + c <= b || b + c <= a { 
		return NaT
	}
	if a == b && b == c {
		k = Equ
	} else if a == b || a == c || b == c {
		k = Iso
	} else if a != b && b != c && a != c {
		k = Sca
	}

	return k
}
