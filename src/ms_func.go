package src

import (
	"fmt"
	"strings"
)

type msCond struct {
	Suffix, Condition string
}

func ms_conds() []msCond {
	return []msCond{
		msCond{"_Eq", "="},
		msCond{"_NotEq", "!="},
		msCond{"_LT", "<"},
		msCond{"_LE", "<="},
		msCond{"_GT", ">"},
		msCond{"_GE", ">="},
	}

}

func ms_str_cond() []msCond {
	return []msCond{
		msCond{"_Eq", "="},
		msCond{"_NotEq", "!="},
	}

}

func ms_in() []msCond {
	return []msCond{
		msCond{"_In", " IN "},
		msCond{"_NotIn", " NOT IN "},
	}

}

func ms_gen_types() []string {
	return []string{"Deleter", "Updater", "Selector"}
}

func ms_col_name(f *Column) []string {
	return []string{}
}

func ms_to_slice(typ ...string) []string {
	return typ
}

func ms_col_comment_json(comment string) string {
	//fmt.Println( comment)
	arr := strings.Split(strings.ToLower(comment), ",")
	for _, s := range arr {
		//fmt.Println(s, comment)
		if "nojson" == strings.Trim(s, " ") {
			return "`json:\"-\"`"
		}
	}
	return ""
}

func ms_col_comment_raw(comment string) string {
	if strings.Trim(comment, " ") != "" {
		return "//" + comment
	}
	return ""
}

func ms_append_fieldnames(fields []*Column, slice string, ignoreNames ...string) string {
	ignore := map[string]bool{}
	for _, n := range ignoreNames {
		ignore[n] = true
	}

	str := ""
	i := 0
	for _, f := range fields {
		if ignore[f.ColumnName] {
			continue
		}
		str += fmt.Sprintf("%s = append(%s, row.%s) \n", slice, slice, f.ColumnNameCamel)

		i++
	}

	return str
}

func ms_question_mark(fields []*Column, ignoreNames ...string) string {
	ignore := map[string]bool{}
	for _, n := range ignoreNames {
		ignore[n] = true
	}

	n := 0
	//l:=len(fields) - len(ignore)
	for _, f := range fields {
		if ignore[f.ColumnName] {
			continue
		}
		n++
	}

	s := strings.Repeat("?,", n)

	return s[0 : len(s)-1]
}

// retype checks typ against known types, and prefixing
// ArgType.CustomTypePackage (if applicable).
func go_to_java_type(typ string) string {
	t := strings.ToLower(typ)
	switch t {
	case "int", "int64", "int32":
		return "int"
	case "string":
		return "String"
	case "float32":
		return "float"
	case "float64":
		return "double"
	case "binary", "varbinary", "tinyblob", "blob", "mediumblob", "longblob":
		return "byte[]"
	}
	return "UNKNOWN"
}

func go_datatype_to_defualt_go_type(typ string) string {
	t := strings.ToLower(typ)
	switch t {
	case "int", "int64", "int32":
		return "0"
	case "string":
		return `""`
	case "float32":
		return "float32(0)"
	case "float64":
		return "float64(0)"
	case "[]byte":
		return "[]byte{}"
	}
	return "UNKNOWN"
}

func ms_trigger_colmun(colmunGoType string) (columnName string) {
	if colmunGoType == "int" {
		return "TargetId"
	}
	return "TargetStr"
}
