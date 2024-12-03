func addSpaces(s string, spaces []int) string {
	str := strings.Builder{}
	str.Grow(len(s) + len(spaces))

	cursor := 0
	for _, spacePosition := range spaces {
		str.WriteString(s[cursor:spacePosition])
		str.WriteString(" ")
		cursor = spacePosition
	}

	str.WriteString(s[cursor:])

	return str.String()
}