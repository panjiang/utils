package xreflect

import "reflect"

// GetFields reflects the fields of a struct object
func GetFields(obj interface{}) map[string]interface{} {
	objVal := reflect.ValueOf(obj).Elem()
	fieldsNum := objVal.NumField()

	fieldsMap := make(map[string]interface{}, fieldsNum)
	for i := 0; i < fieldsNum; i++ {
		fieldVal := objVal.Field(i)
		fieldType := objVal.Type().Field(i)
		if fieldType.Tag == "ignore" {
			continue
		}
		fieldsMap[fieldType.Name] = fieldVal.Interface()

	}
	return fieldsMap
}
