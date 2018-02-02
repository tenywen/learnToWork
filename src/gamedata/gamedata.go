package gamedata

import (
	"fmt"
	"strconv"
)

import (
	. "logger"
)

func get(tableName string, index int32, fieldName string) string {
	if t, ok := tables[tableName]; ok {
		if _, ok := t.fields[index]; ok { // find index
			for _, v := range t.fields[index] {
				if v.name == fieldName {
					return v.value
				}
			}
			DEBUG("gamedata", tableName, "id ", index, "not exist filedName:", fieldName)
			return ""
		} else {
			DEBUG("gamedata", tableName, " not exist id:", index)
			return ""
		}
	}
	DEBUG("gamedata tableName not exist", tableName)
	return ""
}

func GetKeys(tableName string) (result []int32) {
	if t, ok := tables[tableName]; ok {
		result = make([]int32, len(t.fields))
		for k, _ := range t.fields {
			result = append(result, k)
		}
	}
	DEBUG("gamedata GetKeys table not exist:", tableName)
	return
}

func GetInt(tableName string, index int32, fieldName string) int64 {
	tmp := get(tableName, index, fieldName)
	if tmp == "" {
		return 0
	}

	result, _ := strconv.ParseInt(tmp, 10, 64)
	return int64(result)
}

func GetString(tableName string, index int32, fieldName string) string {
	return get(tableName, index, fieldName)
}

func GetFloat(tableName string, index int32, fieldName string) float64 {
	tmp := get(tableName, index, fieldName)
	if tmp == "" {
		return 0.0
	}

	result, _ := strconv.ParseFloat(tmp, 64)
	return result
}

func GetConfigId(tableName string, fields []string, values []interface{}) int32 {
	ok := false
	for _, id := range GetKeys(tableName) {
		for k := range fields {
			if get(tableName, id, fields[k]) == fmt.Sprint(values[k]) {
				ok = true
			} else {
				ok = false
				break
			}
		}
		if ok {
			return id
		}
	}
	return 0
}
