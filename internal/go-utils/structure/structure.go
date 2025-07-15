package structure

func PresentOrDefault(variable *string, def string) {
	if len(*variable) == 0 {
		*variable = def
	}
	return
}
