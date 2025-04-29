package slice

func Convert[T, T2 any](inputSlice []T, caster func(T) T2) []T2 {
	s := make([]T2, len(inputSlice))
	for i, element := range inputSlice {
		s[i] = caster(element)
	}
	return s
}
