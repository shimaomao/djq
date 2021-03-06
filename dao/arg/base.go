package arg

import (
	"errors"
	"mimi/djq/model"
	"mimi/djq/util"
	"strings"
)

var ErrUpdateObjectEmpty = errors.New("dao: updateObject is empty")

const (
	//SELECT 列名称 FROM 表名称
	SelectSql = "select {columnNames} from {tableName} {conditions};"
	CountSql  = "select count(*) from {tableName} {conditions};"
	//DELETE FROM 表名称 WHERE 列名称 = 值
	DeleteSql = "delete from {tableName} {conditions};"
	//INSERT INTO table_name (列1, 列2,...) VALUES (值1, 值2,....)
	InsertSql = "insert into {tableName} ({columnNames}) values ({columnValues});"
	//"UPDATE 表名称 SET 列名称 = 新值 WHERE 列名称 = 某值"
	UpdateSql = "update {tableName} set {collumnNameValues} {conditions};"
)

func bindColumnNames(sql, columnNames string) string {
	return strings.Replace(sql, "{columnNames}", columnNames, -1)
}

func bindColumnValues(sql, columnValues string) string {
	return strings.Replace(sql, "{columnValues}", columnValues, -1)
}

func bindColumnNameValues(sql, collumnNameValues string) string {
	return strings.Replace(sql, "{collumnNameValues}", collumnNameValues, -1)
}

func bindTableName(sql, tableName string) string {
	return strings.Replace(sql, "{tableName}", tableName, -1)
}

func bindConditions(sql, conditions string) string {
	return strings.Replace(sql, "{conditions}", conditions, -1)
}

type BaseArgInterface interface {
	getCountConditions() (string, []interface{})
	GetTargetPage() int
	SetTargetPage(int)
	GetPageSize() int
	SetPageSize(int)
	GetIdEqual() string
	SetIdEqual(idEqual string)
	GetIdsIn() []string
	SetIdsIn(idsIn []string)
	GetIncludeDeleted() bool
	SetIncludeDeleted(includeDeleted bool)
	GetUpdateObject() interface{}
	SetUpdateObject(updateObject interface{})
	GetUpdateNames() []string
	SetUpdateNames(updateNames []string)
	GetOrderBy() string
	GetModelInstance() model.BaseModelInterface
	GetDisplayNames() []string
}

func getBaseSql(arg BaseArgInterface, sql string) string {
	return bindTableName(sql, arg.GetModelInstance().GetTableName())
}

func BuildFindSql(arg BaseArgInterface) (string, []interface{}, []string) {
	sql := getBaseSql(arg, SelectSql)

	columnNames := getShowColumnNames(arg)
	sql = bindColumnNames(sql, strings.Join(columnNames, ","))

	conditionStr, params := getFindConditions(arg)
	sql = bindConditions(sql, conditionStr)

	return sql, params, columnNames
}

func BuildCountSql(arg BaseArgInterface) (string, []interface{}) {
	sql := getBaseSql(arg, CountSql)

	conditionStr, params := getCountConditions(arg)
	sql = bindConditions(sql, conditionStr)

	return sql, params
}

func BuildInsertSql(arg BaseArgInterface) (string, []string) {
	sql := getBaseSql(arg, InsertSql)

	columnNames := getAllColumnNames(arg)
	sql = bindColumnNames(sql, strings.Join(columnNames, ","))

	columnValues := getColumnValuePlaceHolder(columnNames)
	sql = bindColumnValues(sql, strings.Join(columnValues, ","))

	return sql, columnNames
}

func BuildUpdateSql(arg BaseArgInterface) (string, []interface{}) {
	if arg.GetUpdateObject() == nil {
		panic(ErrUpdateObjectEmpty)
	}
	sql := getBaseSql(arg, UpdateSql)

	columnNameValues, paramValues := getColumnNameValues(arg)
	sql = bindColumnNameValues(sql, strings.Join(columnNameValues, ","))

	conditionStr, paramConditions := getCountConditions(arg)
	sql = bindConditions(sql, conditionStr)

	return sql, append(paramValues, paramConditions...)
}

func BuildDeleteSql(arg BaseArgInterface) (string, []interface{}) {
	sql := getBaseSql(arg, DeleteSql)

	conditionStr, params := getCountConditions(arg)
	sql = bindConditions(sql, conditionStr)

	return sql, params
}

func BuildLogicalDeleteSql(arg BaseArgInterface) (string, []interface{}) {
	sql := getBaseSql(arg, UpdateSql)

	sql = bindColumnNameValues(sql, " del_flag = true")

	conditionStr, paramConditions := getCountConditions(arg)
	sql = bindConditions(sql, conditionStr)

	return sql, paramConditions
}
func getAllColumnNames(arg BaseArgInterface) []string {
	return arg.GetModelInstance().GetDBNames()
}

func getColumnValuePlaceHolder(columnNames []string) []string {
	len := len(columnNames)
	placeholder := make([]string, 0, len)
	for i := 0; i < len; i++ {
		placeholder = append(placeholder, "?")
	}
	return placeholder
}

func getShowColumnNames(arg BaseArgInterface) []string {
	if arg.GetDisplayNames() == nil || len(arg.GetDisplayNames()) == 0 {
		return getAllColumnNames(arg)
	}
	size := len(arg.GetDisplayNames())
	s := make([]string, size, size)
	modelObj := arg.GetModelInstance()
	for i, v := range arg.GetDisplayNames() {
		s[i] = modelObj.GetDBFromMapName(v)
	}
	if len(s) == 0 {
		return getAllColumnNames(arg)
	}
	return s
}

func getColumnNameValues(arg BaseArgInterface) ([]string, []interface{}) {
	if arg.GetUpdateNames() == nil || len(arg.GetUpdateNames()) == 0 {
		arg.SetUpdateNames(arg.GetModelInstance().GetMapNames())
	}
	size := len(arg.GetUpdateNames())
	s := make([]string, size, size)
	params := make([]interface{}, size, size)
	modelObj := arg.GetModelInstance()
	for i, v := range arg.GetUpdateNames() {
		s[i] = modelObj.GetDBFromMapName(v) + " = ? "
		params[i] = arg.GetUpdateObject().(model.BaseModelInterface).GetValue4Map(v)
	}
	return s, params
}

func getFindConditions(arg BaseArgInterface) (string, []interface{}) {
	sql, params := getCountConditions(arg)
	if arg.GetOrderBy() != "" {
		sql += " order by " + arg.GetOrderBy()
	}
	if arg.GetPageSize() > 0 {
		sql += " limit ?,?"
		start := util.ComputePageStart(arg.GetTargetPage(), arg.GetPageSize())
		params = append(params, start, arg.GetPageSize())
	}
	return sql, params
}

func getCountConditions(arg BaseArgInterface) (string, []interface{}) {
	sql, params := arg.getCountConditions()
	if arg.GetIdEqual() != "" {
		if sql != "" {
			sql += " and"
		} else {
			sql += " where"
		}
		sql += " id = ?"
		params = append(params, arg.GetIdEqual())
	}
	if arg.GetIdsIn() != nil && len(arg.GetIdsIn()) != 0 {
		if sql != "" {
			sql += " and"
		} else {
			sql += " where"
		}
		sql += " id in ("
		for i, id := range arg.GetIdsIn() {
			if i != 0 {
				sql += ","
			}
			sql += "?"
			params = append(params, id)
		}
		sql += ")"
	}
	if !arg.GetIncludeDeleted() {
		if sql != "" {
			sql += " and"
		} else {
			sql += " where"
		}
		sql += " del_flag = false"
	}
	return sql, params
}
