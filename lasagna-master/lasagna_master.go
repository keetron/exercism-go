package lasagna

func PreparationTime(layers []string, t int) int {
	if t == 0 {
		t = 2
	}
	return len(layers) * t
}

func Quantities(l []string) (n int, s float64) {
	for i := 0; i < len(l); i++ {
		switch l[i] {
		case "noodles":
			n = n + 50
		case "sauce":
			s = s + 0.2
		}
	}
	return n, s
}

func AddSecretIngredient(f, m []string) {
	m[len(m)-1] = f[len(f)-1]
}

func ScaleRecipe(q []float64, p int) []float64 {
	var a []float64
	for i := 0; i < len(q); i++ {
		x := q[i] * (float64(p) / 2)
		a = append(a, x)
	}
	return a
}
