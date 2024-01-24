package bloom_filter

type Contract interface {
	Add(value string)
	IsValue(value string) bool
	Hash1(value string) int
	Hash2(value string) int
	GetStatusIsValue() int
}
