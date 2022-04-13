package utils

// Returns a slice with all the unique elements of the input slice.
func Unique[T comparable](slice []T) []T {
	m := make(map[T]bool)
	for _, v := range slice {
		m[v] = true
	}
	unique := make([]T, 0, len(m))
	for k := range m {
		unique = append(unique, k)
	}
	return unique
}
