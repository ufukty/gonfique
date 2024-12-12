package bijective

func Bijective26(i int) string {
	const q = 26
	i++
	s := ""
	for i > 0 {
		i--
		s = string(byte(int(byte('A'))+i%q)) + s
		i /= q
	}
	return s
}
