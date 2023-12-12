package util

type Number interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

func Sum[N Number](values []N) N {
	var sum N
	for _, value := range values {
		sum += value
	}
	return sum
}

func SumOf[T any, N Number](list []T, fn func(T) N) N {
	var sum N
	for _, item := range list {
		sum += fn(item)
	}
	return sum
}
