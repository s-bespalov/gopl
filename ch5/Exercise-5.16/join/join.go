package join

// variadic version od strings.Join
func join(sep string, elements ...string) string {
	if len(elements) == 0 {
		return ""
	}
	r := elements[0]
	for i := 1; i < len(elements); i++ {
		r = r + sep + elements[i]
	}
	return r
}
