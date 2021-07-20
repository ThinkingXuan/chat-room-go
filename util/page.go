package util

// IndexToPage return page by index and size
func IndexToPage(index, size int) int {
	if index < 0 {
		index = -index - 1
	}
	return index * size
}
