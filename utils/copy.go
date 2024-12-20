package utils

// in 2d slice underlying container still share memory
// so we need to deep copy
func CopyGrid(org [][]string) [][]string {
	fscpy := make([][]string, len(org))
	for i := range org {
		fscpy[i] = make([]string, len(org[i]))
		copy(fscpy[i], org[i])
	}
	return fscpy
}
