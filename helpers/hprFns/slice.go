package hprFns

import "slices"

func AppendIfNotExist[T comparable](s []T, vs ...T) []T {
	for _, v := range vs {
		if !slices.Contains[[]T, T](s, v) {
			s = append(s, v)
		}
	}
	return s
}

func AppendIfNotNilAndNotExist[T comparable](s []T, vs ...*T) []T {
	for _, v := range vs {
		if v != nil && !slices.Contains[[]T, T](s, *v) {
			s = append(s, *v)
		}
	}
	return s
}
