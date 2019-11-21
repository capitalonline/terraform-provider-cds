package utils

func MergeSlice(s1 []*string, s2 []*string) []*string {
	slice := make([]*string, len(s1)+len(s2))
	copy(slice, s1)
	copy(slice[len(s1):], s2)
	return slice
}
