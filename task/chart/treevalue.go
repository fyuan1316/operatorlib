package chart

//
//import "strings"
//
//func TreeValue(values map[string]interface{}) map[string]interface{} {
//	result := make(map[string]interface{})
//	for k, v := range values {
//		setTreeValue(result, k, v)
//	}
//	return result
//}
//
//func setTreeValue(tree map[string]interface{}, key string, value interface{}) {
//	tmp := tree
//	keys := strings.Split(key, ".")
//	count := len(keys)
//	for i, k := range keys {
//		if i < count-1 {
//			if _, ok := tmp[k]; !ok {
//				tmp[k] = map[string]interface{}{}
//			}
//			tmp = tmp[k].(map[string]interface{})
//		} else {
//			tmp[k] = value
//		}
//	}
//}
