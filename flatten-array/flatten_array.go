package flatten
import "reflect"


func Flatten(nested interface{}) []interface{} {
	var flattened = []interface{}{}
	return flat(nested, flattened)
}
func flat(el interface{}, list []interface{}) []interface{} {
	if isList(el) {
		for _, e := range el.([]interface{}) {
			list = flat(e, list)
		}
	} else if el != nil {
		list = append(list, el)
	}
	return list
}

func isList(el interface{}) bool {
	kind := reflect.ValueOf(el).Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		return true
	} else  {
		return false
	}
}