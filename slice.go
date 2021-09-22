package supportlib

func AppendStrUnique(slice []string, item string) (res []string) {
	for _, i := range slice {
		if i == item {
			return slice
		}
	}
	return append(slice, item)
}
