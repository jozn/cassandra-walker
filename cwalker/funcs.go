package cwalker

import (
    "text/template"
    "strings"
    "github.com/jozn/xox/snaker"
)



func NewTemplateFuncs() template.FuncMap {
    return template.FuncMap{
      ////"colcount":       colcount,
      //  "colnames":      colnames,
      //  "colnamesquery": colnamesquery,
      //  //"colprefixnames": colprefixnames,
      //  "colvals":     colvals,
      //  "fieldnames":  fieldnames,
      //  "goparamlist": goparamlist,
      //  //"reniltype":      reniltype,
      //  //"retype":         retype,
      //  "shortname": shortname,
      //  //"convext":        convext,
      //  //"schema":         schemafn,
      //  "colname": colname,
      //  //"hascolumn":      hascolumn,
      //  //"hasfield":       hasfield,

        //"toLower":              strings.ToLower,
        //"toUpper":              strings.ToUpper,
        //"ms_col_nanme":         ms_col_name,
        //"ms_conds":             ms_conds,
        //"ms_in":                ms_in,
        //"ms_gen_types":         ms_gen_types,
        "ms_to_slice":          ms_to_slice,
        //"ms_str_cond":          ms_str_cond,
        //"ms_append_fieldnames": ms_append_fieldnames,
        //"ms_question_mark":     ms_question_mark,
        //"ms_col_comment_json":  ms_col_comment_json,
        //"ms_col_comment_raw":   ms_col_comment_raw,
        //"ms_trigger_colmun":    ms_trigger_colmun,
        //"to_java_type":                to_java_type,
        //"datatype_to_defualt_go_type": datatype_to_defualt_go_type,
    }
}

func shortname(typ string, scopeConflicts ...interface{}) string {
    var v string
    var ok bool

    // check short name map
    if v, ok = ShortNameTypeMap[typ]; !ok {
        // calc the short name
        u := []string{}
        for _, s := range strings.Split(strings.ToLower(snaker.CamelToSnake(typ)), "_") {
            if len(s) > 0 && s != "id" {
                u = append(u, s[:1])
            }
        }
        v = strings.Join(u, "")

        // check go reserved names
        if n, ok := goReservedNames[v]; ok {
            v = n
        }

        // store back to short name map
        ShortNameTypeMap[typ] = v
    }

    // initial conflicts are the default imported packages from
    // xo_package.go.tpl
    conflicts := map[string]bool{
        "sql":     true,
        "driver":  true,
        "csv":     true,
        "errors":  true,
        "fmt":     true,
        "regexp":  true,
        "strings": true,
        "time":    true,
    }

    // add scopeConflicts to conflicts
    for _, c := range scopeConflicts {
        switch k := c.(type) {
        case string:
            conflicts[k] = true

        case []*Column:
            for _, f := range k {
                conflicts[f.ColumnName] = true
            }
            /*case []*QueryParam:
            for _, f := range k {
                conflicts[f.Name] = true
            }*/

        default:
            panic("not implemented")
        }
    }

    // append suffix if conflict exists
    if _, ok := conflicts[v]; ok {
        v = v + "_sufix" //NameConflictSuffix
    }

    return v
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

var goReservedNames = map[string]string{
    "break":       "brk",
    "case":        "cs",
    "chan":        "chn",
    "const":       "cnst",
    "continue":    "cnt",
    "default":     "def",
    "defer":       "dfr",
    "else":        "els",
    "fallthrough": "flthrough",
    "for":         "fr",
    "func":        "fn",
    "go":          "goVal",
    "goto":        "gt",
    "if":          "ifVal",
    "import":      "imp",
    "interface":   "iface",
    "map":         "mp",
    "package":     "pkg",
    "range":       "rnge",
    "return":      "ret",
    "select":      "slct",
    "struct":      "strct",
    "switch":      "swtch",
    "type":        "typ",
    "var":         "vr",

    // go types
    "error":      "e",
    "bool":       "b",
    "string":     "str",
    "byte":       "byt",
    "rune":       "r",
    "uintptr":    "uptr",
    "int":        "i",
    "int8":       "i8",
    "int16":      "i16",
    "int32":      "i32",
    "int64":      "i64",
    "uint":       "u",
    "uint8":      "u8",
    "uint16":     "u16",
    "uint32":     "u32",
    "uint64":     "u64",
    "float32":    "z",
    "float64":    "f",
    "complex64":  "c",
    "complex128": "c128",
}
