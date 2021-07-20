package sort

import (
	"sort"
)

// AscMsi map[string]interface{} sort asc
func AscMsi(msi map[string]interface{}) (key []string, val []interface{}) {
	if msi == nil {
		return
	}
	length := len(msi)
	key = make([]string, length, length)
	val = make([]interface{}, length, length)
	i := 0
	for k, _ := range msi {
		key[i] = k
		i++
	}
	sort.Strings(key)
	for k, v := range key {
		val[k] = msi[v]
	}
	return
}

// DescMsi map[string]interface{} sort desc
func DescMsi(msi map[string]interface{}) (key []string, val []interface{}) {
	if msi == nil {
		return
	}
	length := len(msi)
	key = make([]string, length, length)
	val = make([]interface{}, length, length)
	i := 0
	for k, _ := range msi {
		key[i] = k
		i++
	}
	sort.Strings(key)
	for x, y := 0, len(key)-1; x < y; x, y = x+1, y-1 {
		key[x], key[y] = key[y], key[x]
	}
	for k, v := range key {
		val[k] = msi[v]
	}
	return
}
