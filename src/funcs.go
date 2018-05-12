package src

import (
	"fmt"
	"ms/xox/snaker"
	"strconv"
	"strings"
	"text/template"
)

// NewTemplateFuncs returns a set of template funcs bound to the supplied args.
func NewTemplateFuncs() template.FuncMap {
	return template.FuncMap{
		//"colcount":       colcount,
		"colnames":      colnames,
		"colnamesquery": colnamesquery,
		//"colprefixnames": colprefixnames,
		"colvals":     colvals,
		"fieldnames":  fieldnames,
		"goparamlist": goparamlist,
		//"reniltype":      reniltype,
		//"retype":         retype,
		"shortname": shortname,
		//"convext":        convext,
		//"schema":         schemafn,
		"colname": colname,
		//"hascolumn":      hascolumn,
		//"hasfield":       hasfield,

		"toLower":              strings.ToLower,
		"toUpper":              strings.ToUpper,
		"ms_col_nanme":         ms_col_name,
		"ms_conds":             ms_conds,
		"ms_in":                ms_in,
		"ms_gen_types":         ms_gen_types,
		"ms_to_slice":          ms_to_slice,
		"ms_str_cond":          ms_str_cond,
		"ms_append_fieldnames": ms_append_fieldnames,
		"ms_question_mark":     ms_question_mark,
		"ms_col_comment_json":  ms_col_comment_json,
		"ms_col_comment_raw":   ms_col_comment_raw,
		"ms_trigger_colmun":    ms_trigger_colmun,
		//"to_java_type":                to_java_type,
		//"datatype_to_defualt_go_type": datatype_to_defualt_go_type,
	}
}

// retype checks typ against known types, and prefixing
// ArgType.CustomTypePackage (if applicable).
func retype(typ string) string {
	if strings.Contains(typ, ".") {
		return typ
	}

	prefix := ""
	for strings.HasPrefix(typ, "[]") {
		typ = typ[2:]
		prefix = prefix + "[]"
	}

	if _, ok := KnownTypeMap[typ]; !ok {
		/*pkg := CustomTypePackage
		if pkg != "" {
			pkg = pkg + "."
		}*/

		//return prefix + pkg + typ
		return prefix + typ
	}

	return prefix + typ
}

//// reniltype checks typ against known nil types (similar to retype), prefixing
//// ArgType.CustomTypePackage (if applicable).
//func reniltype(typ string) string {
//	if strings.Contains(typ, ".") {
//		return typ
//	}
//
//	if strings.HasSuffix(typ, "{}") {
//		if _, ok := c.KnownTypeMap[typ[:len(typ)-2]]; ok {
//			return typ
//		}
//
//		pkg := c.CustomTypePackage
//		if pkg != "" {
//			pkg = pkg + "."
//		}
//
//		return pkg + typ
//	}
//
//	return typ
//}
//
// shortname generates a safe Go identifier for typ. typ is first checked
// against ArgType.ShortNameTypeMap, and if not found, then the value is
// calculated and stored in the ShortNameTypeMap for future use.
//
// A shortname is the concatentation of the lowercase of the first character in
// the words comprising the name. For example, "MyCustomName" will have have
// the shortname of "mcn".
//
// If a generated shortname conflicts with a Go reserved name, then the
// corresponding value in goReservedNames map will be used.
//
// Generated shortnames that have conflicts with any scopeConflicts member will
// have ArgType.NameConflictSuffix appended.
//
// Note: recognized types for scopeConflicts are string, []*Field,
// []*QueryParam.
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

// colnames creates a list of the column names found in fields, excluding any
// Field with Name contained in ignoreNames.
//
// Used to present a comma separated list of column names, that can be used in
// a SELECT, or UPDATE, or other SQL clause requiring an list of identifiers
// (ie, "field_1, field_2, field_3, ...").
func colnames(columns []*Column, ignoreNames ...string) string {
	ignore := map[string]bool{}
	for _, n := range ignoreNames {
		ignore[n] = true
	}

	str := ""
	i := 0
	for _, f := range columns {
		if ignore[f.ColumnName] {
			continue
		}

		if i != 0 {
			str = str + ", "
		}
		str = str + colname(f)
		i++
	}

	return str
}

// colnamesquery creates a list of the column names in fields as a query and
// joined by sep, excluding any Field with Name contained in ignoreNames.
//
// Used to create a list of column names in a WHERE clause (ie, "field_1 = $1
// AND field_2 = $2 AND ...") or in an UPDATE clause (ie, "field = $1, field =
// $2, ...").
func colnamesquery(fields []*Column, sep string, ignoreNames ...string) string {
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

		if i != 0 {
			str = str + sep
		}
		str = str + colname(f) + " = " + "?" //c.Loader.NthParam(i)
		i++
	}

	return str
}

//
//// colprefixnames creates a list of the column names found in fields with the
//// supplied prefix, excluding any Field with Name contained in ignoreNames.
////
//// Used to present a comma separated list of column names with a prefix. Used in
//// a SELECT, or UPDATE (ie, "t.field_1, t.field_2, t.field_3, ...").
//func colprefixnames(fields []*Field, prefix string, ignoreNames ...string) string {
//	ignore := map[string]bool{}
//	for _, n := range ignoreNames {
//		ignore[n] = true
//	}
//
//	str := ""
//	i := 0
//	for _, f := range fields {
//		if ignore[f.Name] {
//			continue
//		}
//
//		if i != 0 {
//			str = str + ", "
//		}
//		str = str + prefix + "." + colname(f.Col)
//		i++
//	}
//
//	return str
//}
//
// colvals creates a list of value place holders for fields excluding any Field
// with Name contained in ignoreNames.
//
// Used to present a comma separated list of column place holders, used in a
// SELECT or UPDATE statement (ie, "$1, $2, $3 ...").
func colvals(fields []*Column, ignoreNames ...string) string {
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

		if i != 0 {
			str = str + ", "
		}
		str = str + "?" //c.Loader.NthParam(i)
		i++
	}

	return str
}

// fieldnames creates a list of field names from fields of the adding the
// provided prefix, and excluding any Field with Name contained in ignoreNames.
//
// Used to present a comma separated list of field names, ie in a Go statement
// (ie, "t.Field1, t.Field2, t.Field3 ...")
//me: used uper names
func fieldnames(fields []*Column, prefix string, ignoreNames ...string) string {
	ignore := map[string]bool{}
	for _, n := range ignoreNames {
		ignore[n] = true
	}

	str := ""
	i := 0
	for _, f := range fields {
		if ignore[f.ColumnNameCamel] {
			continue
		}

		if i != 0 {
			str = str + ", "
		}
		str = str + prefix + "." + f.ColumnNameCamel
		//fmt.Println(f.Name)
		i++
	}

	return str
}

//
//// colcount returns the 1-based count of fields, excluding any Field with Name
//// contained in ignoreNames.
////
//// Used to get the count of fields, and useful for specifying the last SQL
//// parameter.
//func colcount(fields []*Field, ignoreNames ...string) int {
//	ignore := map[string]bool{}
//	for _, n := range ignoreNames {
//		ignore[n] = true
//	}
//
//	i := 1
//	for _, f := range fields {
//		if ignore[f.Name] {
//			continue
//		}
//
//		i++
//	}
//	return i
//}
//
// goReservedNames is a map of of go reserved names to "safe" names.
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

// goparamlist converts a list of fields into their named Go parameters,
// skipping any Field with Name contained in ignoreNames. addType will cause
// the go Type to be added after each variable name. addPrefix will cause the
// returned string to be prefixed with ", " if the generated string is not
// empty.
//
// Any field name encountered will be checked against goReservedNames, and will
// have its name substituted by its corresponding looked up value.
//
// Used to present a comma separated list of Go variable names for use with as
// either a Go func parameter list, or in a call to another Go func.
// (ie, ", a, b, c, ..." or ", a T1, b T2, c T3, ...").
func goparamlist(fields []*Column, addPrefix bool, addType bool, ignoreNames ...string) string {
	ignore := map[string]bool{}
	for _, n := range ignoreNames {
		ignore[n] = true
	}

	i := 0
	vals := []string{}
	for _, f := range fields {
		if ignore[f.ColumnName] {
			continue
		}

		s := "v" + strconv.Itoa(i)
		if len(f.ColumnName) > 0 {
			n := strings.Split(snaker.CamelToSnake(f.ColumnName), "_")
			s = strings.ToLower(n[0]) + f.ColumnName[len(n[0]):]
		}

		// check go reserved names
		if r, ok := goReservedNames[strings.ToLower(s)]; ok {
			s = r
		}

		// add the go type
		if addType {
			s += " " + f.GoTypeOut //retype(f.SqlType)
		}

		// add to vals
		vals = append(vals, s)

		i++
	}

	// concat generated values
	str := strings.Join(vals, ", ")
	if addPrefix && str != "" {
		return ", " + str
	}

	return str
}

//// convext generates the Go conversion for f in order for it to be assignable
//// to t.
////
//// FIXME: this should be a better name, like "goconversion" or some such.
//func convext(prefix string, f *Field, t *Field) string {
//	expr := prefix + "." + f.Name
//	if f.Type == t.Type {
//		return expr
//	}
//
//	ft := f.Type
//	if strings.HasPrefix(ft, "sql.Null") {
//		expr = expr + "." + f.Type[8:]
//		ft = strings.ToLower(f.Type[8:])
//	}
//
//	if t.Type != ft {
//		expr = t.Type + "(" + expr + ")"
//	}
//
//	return expr
//}
//
//// schemafn takes a series of names and joins them with the schema name.
//func schemafn(s string, names ...string) string {
//	// escape table names
//	if c.EscapeTableNames {
//		for i, t := range names {
//			names[i] = c.Loader.Escape(TableEsc, t)
//		}
//	}
//
//	n := strings.Join(names, ".")
//
//	if s == "" && n == "" {
//		return ""
//	}
//
//	if s != "" && n != "" {
//		if c.EscapeSchemaName {
//			s = c.Loader.Escape(SchemaEsc, s)
//		}
//		s = s + "."
//	}
//
//	return s + n
//}
//
// colname returns the ColumnName of col, optionally escaping it if
// ArgType.EscapeColumnNames is toggled.
func colname(col *Column) string {
	if EscapeColumnNames {
		//return c.Loader.Escape(ColumnEsc, col.ColumnName)
		return fmt.Sprintf("`%s`", col.ColumnName)
	}

	return col.ColumnName
}

//
//// hascolumn takes a list of fields and determines if field with the specified
//// column name is in the list.
//func hascolumn(fields []*Field, name string) bool {
//	for _, f := range fields {
//		if f.Col.ColumnName == name {
//			return true
//		}
//	}
//
//	return false
//}
//
//// hasfield takes a list of fields and determines if field with the specified
//// field name is in the list.
//func hasfield(fields []*Field, name string) bool {
//	for _, f := range fields {
//		if f.Name == name {
//			return true
//		}
//	}
//
//	return false
//}
