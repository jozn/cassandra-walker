package src

import (
	"github.com/gedex/inflector"
	"github.com/knq/snaker"
	"regexp"
	"strconv"
	"strings"
    "unicode"
)

//copy of "ms/xox/snaker"

// SnakeToCamel converts s to CamelCase.
func SnakeToCamel(s string) string {
	var r string

	if len(s) == 0 {
		return s
	}

	//ME: hack snake just for those of having "_"
	if strings.Index(s, "_") < 0 {
		return strings.ToUpper(s[:1]) + s[1:]
	}

	for _, w := range strings.Split(s, "_") {
		if w == "" {
			continue
		}

        r += strings.ToUpper(w[:1]) + strings.ToLower(w[1:])
		/*u := strings.ToUpper(w)
		if ok := commonInitialisms[u]; ok {//me not need we use: Id and Html
			r += u
		} else {
			r += strings.ToUpper(w[:1]) + strings.ToLower(w[1:])
		}*/
	}

	return r
}

// ToSnake convert the given string to snake case following the Golang format:
// acronyms are converted to lower-case and preceded by an underscore.
func ToSnake(in string) string {
    runes := []rune(in)
    length := len(runes)

    var out []rune
    for i := 0; i < length; i++ {
        if i > 0 && unicode.IsUpper(runes[i]) && ((i+1 < length && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) {
            out = append(out, '_')
        }
        out = append(out, unicode.ToLower(runes[i]))
    }

    return string(out)
}

// commonInitialisms is the set of commonInitialisms.
//
// taken from: github.com/golang/lint @ 206c0f0
var commonInitialisms = map[string]bool{
	"ACL":   true,
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LHS":   true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SQL":   true,
	"SSH":   true,
	"TCP":   true,
	"TLS":   true,
	"TTL":   true,
	"UDP":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
	"XMPP":  true,
	"XSRF":  true,
	"XSS":   true,
}

// sqlTypeToGoType parse a mysql type into a Go type based on the column
// definition.
func sqlTypeToGoType(sqlType string, nullable bool) (int, string, string) {
	precision := 0
	nilVal := "nil"
	unsigned := false

	// extract unsigned
	if strings.HasSuffix(sqlType, " unsigned") {
		unsigned = true
		sqlType = sqlType[:len(sqlType)-len(" unsigned")]
	}

	// extract precision
	sqlType, precision, _ = ParsePrecision(sqlType)

	var typ string

switchDT:
	switch sqlType {
	case "bit":
		nilVal = "0"
		if precision == 1 {
			nilVal = "false"
			typ = "bool"
			if nullable {
				nilVal = "sql.NullBool{}"
				typ = "sql.NullBool"
			}
			break switchDT
		} else if precision <= 8 {
			typ = "uint8"
		} else if precision <= 16 {
			typ = "uint16"
		} else if precision <= 32 {
			typ = "uint32"
		} else {
			typ = "uint64"
		}
		if nullable {
			nilVal = "sql.NullInt64{}"
			typ = "sql.NullInt64"
		}

	case "bool", "boolean":
		nilVal = "false"
		typ = "bool"
		if nullable {
			nilVal = "sql.NullBool{}"
			typ = "sql.NullBool"
		}

	case "char", "varchar", "tinytext", "text", "mediumtext", "longtext":
		nilVal = `""`
		typ = "string"
		if nullable {
			nilVal = "sql.NullString{}"
			typ = "sql.NullString"
		}

	case "tinyint", "smallint", "mediumint", "int", "integer", "bigint":
		nilVal = "0"
		typ = "int"
		if nullable {
			nilVal = "sql.NullInt64{}"
			typ = "sql.NullInt64"
		}

		/*case "tinyint", "smallint":
		      nilVal = "0"
		      typ = "int16"
		      if nullable {
		          nilVal = "sql.NullInt64{}"
		          typ = "sql.NullInt64"
		      }

		  case "mediumint", "int", "integer":
		      nilVal = "0"
		      typ = args.Int32Type
		      if nullable {
		          nilVal = "sql.NullInt64{}"
		          typ = "sql.NullInt64"
		      }

		  case "bigint":
		      nilVal = "0"
		      typ = "int64"
		      if nullable {
		          nilVal = "sql.NullInt64{}"
		          typ = "sql.NullInt64"
		      }
		*/
	case "float":
		nilVal = "0.0"
		typ = "float32"
		if nullable {
			nilVal = "sql.NullFloat64{}"
			typ = "sql.NullFloat64"
		}

	case "decimal", "double":
		nilVal = "0.0"
		typ = "float64"
		if nullable {
			nilVal = "sql.NullFloat64{}"
			typ = "sql.NullFloat64"
		}

	case "binary", "varbinary", "tinyblob", "blob", "mediumblob", "longblob":
		typ = "[]byte"

	case "timestamp", "datetime", "date", "time":
		nilVal = "time.Time{}"
		typ = "time.Time"
		if nullable {
			nilVal = "mysql.NullTime{}"
			typ = "mysql.NullTime"
		}

	default:
		/*if strings.HasPrefix(sqlType, args.Schema+".") {
		      // in the same schema, so chop off
		      typ = snaker.SnakeToCamelIdentifier(sqlType[len(args.Schema)+1:])
		      nilVal = typ + "(0)"
		  } else {
		      typ = snaker.SnakeToCamelIdentifier(sqlType)
		      nilVal = typ + "{}"
		  }*/
		typ = "UNKNOWN_sqlToGo"
	}

	// add 'u' as prefix to type if its unsigned
	// FIXME: this needs to be tested properly...
	if unsigned && IntRE.MatchString(typ) {
		typ = "u" + typ
	}

	return precision, nilVal, typ
}
var IntRE = regexp.MustCompile(`^int(32|64)?$`)


var PrecScaleRE = regexp.MustCompile(`\(([0-9]+)(\s*,[0-9]+)?\)$`)

// ParsePrecision extracts (precision[,scale]) strings from a data type and
// returns the data type without the string.
func ParsePrecision(dt string) (string, int, int) {
	var err error

	precision := -1
	scale := -1

	m := PrecScaleRE.FindStringSubmatchIndex(dt)
	if m != nil {
		// extract precision
		precision, err = strconv.Atoi(dt[m[2]:m[3]])
		if err != nil {
			panic("could not convert precision")
		}

		// extract scale
		if m[4] != -1 {
			scale, err = strconv.Atoi(dt[m[4]+1 : m[5]])
			if err != nil {
				panic("could not convert scale")
			}
		}

		// change dt
		dt = dt[:m[0]] + dt[m[1]:]
	}

	return dt, precision, scale
}

// SinguralizeIdentifier will singularize a identifier, returning it in
// CamelCase.
func SingularizeIdentifier(s string) string {
	if i := reverseIndexRune(s, '_'); i != -1 {
		s = s[:i] + "_" + inflector.Singularize(s[i+1:])
	} else {
		s = inflector.Singularize(s)
	}

	return snaker.SnakeToCamelIdentifier(s)
}

// reverseIndexRune finds the last rune r in s, returning -1 if not present.
func reverseIndexRune(s string, r rune) int {
	if s == "" {
		return -1
	}

	rs := []rune(s)
	for i := len(rs) - 1; i >= 0; i-- {
		if rs[i] == r {
			return i
		}
	}

	return -1
}
