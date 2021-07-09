package monsoon


/**
*	返回所有的map键值
*/
func MapGetKeys(m map[string]interface{}) []string {
	j := 0
	keys := make([]string, len(m))
	for k := range m {
		keys[j] = k
		j++
	}
	return keys
}

