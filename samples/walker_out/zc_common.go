package xc

import (
	"strings"

	"bytes"
	"fmt"
	"log"
)

//////////////// Logs Funcs ////////////////////
// log funcs are variable func, so you can 'inject' your owen custom logs. ex set x.XCLogErr = func(err error) {} for not logging

var XCLogErr = func(err error) {
	if err != nil {
		log.Println(err)
	}
}

var XCLog = func(str ...interface{}) {
	if len(str) > 0 {
		log.Println("CQL: ", str)
	}
}

//////////// End of Logs /////////////////

//////////////// Constants //////////////////

type whereClause struct {
	condition string
	args      []interface{}
}

func whereClusesToSql(wheres []whereClause, whereSep string) (string, []interface{}) {
	var wheresArr []string
	for _, w := range wheres {
		wheresArr = append(wheresArr, w.condition)
	}
	wheresStr := strings.Join(wheresArr, whereSep)

	var args []interface{}
	for _, w := range wheres {
		args = append(args, w.args...)
	}
	return wheresStr, args
}

func sqlManyDollars(colSize, repeat int, isMysql bool) string {
	//isMysql = true
	if isMysql {
		s := strings.Repeat("?,", colSize)
		s = "(" + s[0:len(s)-1] + "),"
		insVals_ := strings.Repeat(s, repeat)

		return insVals_[0 : len(insVals_)-1]
	}

	buff := bytes.NewBufferString("")
	cnt := 1
	for i := 0; i < repeat; i++ {
		buff.WriteString("(")
		for j := 0; j < colSize; j++ {
			buff.WriteString(fmt.Sprintf("$%d", cnt))
			if j+1 != colSize {
				buff.WriteString(",")
			}
			cnt++
		}
		buff.WriteString(")")
		if i+1 != repeat {
			buff.WriteString(",")
		}
	}
	return buff.String()
}

func dbQuestionForSqlIn(size int) string {
	if size < 1 {
		return ""
	}

	if size == 1 {
		return "?"
	}

	s := strings.Repeat("?,", size)
	s = s[0 : len(s)-1] //remove last ','
	return s
}

//////////////// End of Constants ///////////////
