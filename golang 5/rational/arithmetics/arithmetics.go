package arithmetics

type Fraction struct {
	A int
	B int
}
//import сокращение дроби
func abs(n int) int {
	if n < 0 {return -n}
	return n
}

func (n Fraction) lcm(m Fraction) int {
	return (abs(n.B) * abs(m.B)) / gcd(n.B, m.B)
}

func gcd(a, b int) int {
	for ; b != 0; {a, b = b, a % b}
	return a
}

func (x *Fraction) Inc(y Fraction) {
	c := x.lcm(y)
	xDiv, yDiv := c/(*x).B, c/y.B
	if (*x).B != y.B {
		(*x).A *= xDiv
		(*x).B *= xDiv
		y.A *= yDiv
		y.B *= yDiv
	}
	(*x).A += y.A
	c = gcd((*x).A, (*x).B)
}

func (x *Fraction) Dec(y Fraction) {
	c := x.lcm(y)
	xDiv, yDiv := c/(*x).B, c/y.B
	if (*x).B != y.B {
		(*x).A *= xDiv
		(*x).B *= xDiv
		y.A *= yDiv
		y.B *= yDiv
	}
	(*x).A -= y.A
}

func (x *Fraction) Mult(y Fraction) {
	(*x).A *= y.A
	(*x).B *= y.B
}

func (x *Fraction) Div(y Fraction) {
	y.A, y.B = y.B, y.A
	(*x).A *= y.A
	(*x).B *= y.B
}
