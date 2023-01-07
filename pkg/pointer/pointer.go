package pointer

func ToString(value string) *string {
	return &value
}

func ToInt(value int) *int {
	return &value
}
