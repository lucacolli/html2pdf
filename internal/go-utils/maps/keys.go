package maps

func GetKeys(m map[string]interface{}) []string {
	var ks []string
	for k, _ := range m {
		ks = append(ks, k)
	}
	return ks
}
