package src

import "strings"

var notMakeTableType = []string{}//"user"}

func skipTableModel(table string) bool {
	t := strings.ToLower(table)
	for _, ent := range notMakeTableType {
		if ent == t {
			return true
		}
	}
	return false
}


func needTriggerTable(table string) bool {
	t := strings.ToLower(table)
	for _, ent := range triggerNeededArr {
		if ent == t {
			return true
		}
	}
	return false
}

// KnownTypeMap is the collection of known Go types.
var KnownTypeMap = map[string]bool{
	"bool":        true,
	"string":      true,
	"byte":        true,
	"rune":        true,
	"int":         true,
	"int16":       true,
	"int32":       true,
	"int64":       true,
	"uint":        true,
	"uint8":       true,
	"uint16":      true,
	"uint32":      true,
	"uint64":      true,
	"float32":     true,
	"float64":     true,
	"Slice":       true,
	"StringSlice": true,
}

// ShortNameTypeMap is the collection of Go style short names for types, mainly
// used for use with declaring a func receiver on a type.
var ShortNameTypeMap = map[string]string{
	"bool":        "b",
	"string":      "s",
	"byte":        "b",
	"rune":        "r",
	"int":         "i",
	"int16":       "i",
	"int32":       "i",
	"int64":       "i",
	"uint":        "u",
	"uint8":       "u",
	"uint16":      "u",
	"uint32":      "u",
	"uint64":      "u",
	"float32":     "f",
	"float64":     "f",
	"Slice":       "s",
	"StringSlice": "ss",
}
