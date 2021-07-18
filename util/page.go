package util

// IndexToPage return page by index and size
func IndexToPage(index, size int) int {
	return index * size
}
