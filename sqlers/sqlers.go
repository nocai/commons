package sqlers

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type Sqler interface {
	GetSql() string
	GetCountSql() string
	GetPageSql(pageNo, pageSize int64) string

	GetArgs() []interface{}

	AppendSqlAndArgs(sql string, arg ...interface{}) Sqler
	AppendArgs(args ...interface{}) Sqler
}

type BaseSql struct {
	sql  bytes.Buffer
	args []interface{}
}

func (this BaseSql) GetSql() string {
	return this.sql.String()
}

func (this BaseSql) GetArgs() []interface{} {
	return this.args
}

func (this *BaseSql) AppendSqlAndArgs(sql string, arg ...interface{}) Sqler {
	if sql == "" {
		panic("The sql is blank")
	}
	if !strings.HasSuffix(this.sql.String(), "where 1 = 1") {
		this.sql.WriteString(" where 1 = 1")
	}
	if !strings.HasPrefix(sql, " ") {
		sql = " " + sql
	}
	this.sql.WriteString(sql)
	this.args = append(this.args, arg...)
	return this
}

func (this *BaseSql) AppendArgs(args ...interface{}) Sqler {
	this.args = append(this.args, args...)
	return this
}

func (this *BaseSql) GetCountSql() string {
	if this.sql.Len() > 0 {
		return "select count(1) from (" + this.sql.String() + ") as tttttttttttttttttttttt"
	}
	panic("sqler err")
}

func (this *BaseSql) GetPageSql(pageNo, pageSize int64) string {
	if pageNo > 0 && pageSize > 0 {
		startIndex := (pageNo - 1) * pageSize
		return this.GetSql() + " limit " + strconv.FormatInt(startIndex, 10) + ", " + strconv.FormatInt(pageSize, 10)
	}
	panic("sqler err")
}

func (this *BaseSql) String() string {
	return fmt.Sprintf(`
		The sql = %s
		The Args = %q`, this.sql.String(), this.args)
}

func New(sql string) Sqler {
	if len(sql) <= 0 {
		panic("The sql is blank")
	}

	var sqler BaseSql
	sqler.sql.WriteString(sql)
	return &sqler
}
