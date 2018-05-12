package cwalker

import (
    "strings"
)

func cqlTypesToGoType(sqlType string) (typ, def string) {
    switch strings.ToLower(sqlType) {
    case "string","uuid","text","varchar":
        typ = "string"
        def =`""`
    case "bool":
        typ = "bool"
    case "int","serial","bigint":
        typ = "int"
        def =`0`
    case "json":
        typ = "string"
        def =`""`
    case "bytes","blob":
        typ = "[]byte"
        def =`[]byte{}`
    case "date","time","timestamp":
        typ = "time.Time"
        def =`time.Time.Now()`
    case "decimal":
        typ = "float64"
        def =`0`
    case "float":
        typ = "float32"
        def =`0`

    default:
        typ = "UNKNOWN_sqlToGo__" + typ
        def =`""`
    }
    return
}

