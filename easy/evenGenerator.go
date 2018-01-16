package main

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() uint {
		tmp := i
		i += 2
		return tmp
	}
}
