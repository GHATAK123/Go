package main

type author struct {
	name, branch      string
	particles, salary int
}

// Method with a receiver
// of author type
func (a author) show() (string, string, int, int) {
	return a.name,
		a.branch,
		a.particles,
		a.salary
}

func returnMany(a int) (int, int, int, int) {
	return a + 2, a - 2, a * 2, a / 2
}

func calculate(p, q int) (add, sub, mul, div int) { // bare return
	add = p + q
	sub = p - q
	mul = p * q
	div = p / q
	return
}
