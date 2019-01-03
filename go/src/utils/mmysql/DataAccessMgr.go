package mmysql

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 安装mysql驱动：go get github.com/go-sql-driver/mysql

// 连接sql：
// db, err := sql.Open("mysql", "user:password@tcp(localhost:5555)/dbname?charset=utf8")

type DataSource struct {
	Username string
	Password string
	Ip       string
	Port     int
	Dbname   string
}

func NewDataSource(username string, password string, ip string, port int, dbname string) DataSource {
	return DataSource{Username: username, Password: password, Ip: ip, Port: port, Dbname: dbname}
}

// 默认端口 3306
func NewDataSourceDef(username string, password string, ip string, dbname string) DataSource {
	return DataSource{Username: username, Password: password, Ip: ip, Port: 3306, Dbname: dbname}
}

// parseTime=true 可以解析为time.Time类型，否则则解析为[]uint8类型（DateAccessMgr会转为string类型）
// 完全版：format := "%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local&timeout=%s&readTimeout=%s&writeTimeout=%s"
func (d DataSource) ConnectString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local", d.Username, d.Password, d.Ip, d.Port, d.Dbname)
}

type DataAccessMgr struct {
	DataSource DataSource
	Db         *sql.DB
}

func Open(dataSource DataSource) (*DataAccessMgr, error) {
	driverName := "mysql"
	db, err := sql.Open(driverName, dataSource.ConnectString())
	if err != nil {
		return nil, err
	}
	return &DataAccessMgr{Db: db, DataSource: dataSource}, nil
}

// 打印结果时候使用
func (client DataAccessMgr) Show(obj interface{}) {
	switch obj.(type) {
	case []map[string]interface{}:
		list, _ := obj.([]map[string]interface{})
		fmt.Println("client.show.list.size=", len(list))
		for _, item := range list {
			fmt.Println(item)
		}
	case []interface{}:
		list, _ := obj.([]interface{})
		fmt.Println("client.show.list.size=", len(list))
		for _, item := range list {
			fmt.Println(item)
		}
	default:
		fmt.Println(obj)
	}
	fmt.Println("------------------------------------------------------------")
}

// 获取 列名称  len(columns)为查询列数
func (client DataAccessMgr) QueryColumns(sql string, args ...interface{}) ([]string, error) {
	rows, err := client.Db.Query(sql, args...)

	if err != nil {
		return nil, err
	}

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	return columns, nil
}

func (client DataAccessMgr) QueryAllMap(sql string, args ...interface{}) ([]map[string]interface{}, error) {
	rows, err := client.Db.Query(sql, args...)

	if err != nil {
		return nil, err
	}

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	result := []map[string]interface{}{}

	for rows.Next() {
		//将行数据保存到record字典
		_ = rows.Scan(scanArgs...)
		record := make(map[string]interface{})
		for i, col := range values {
			if col != nil {
				switch col.(type) {
				case []uint8:
					record[columns[i]] = string(col.([]byte))
				default:
					record[columns[i]] = col
				}
			}
		}
		result = append(result, record)

	}
	return result, nil
}

// 如果返回结果不是1，返回error
func (client DataAccessMgr) QueryMap(sql string, args ...interface{}) (map[string]interface{}, error) {
	rows, err := client.Db.Query(sql, args...)

	if err != nil {
		return nil, err
	}

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	result := []map[string]interface{}{}

	for rows.Next() {
		//将行数据保存到record字典
		_ = rows.Scan(scanArgs...)
		record := make(map[string]interface{})
		for i, col := range values {
			if col != nil {
				switch col.(type) {
				case []uint8:
					record[columns[i]] = string(col.([]byte))
				default:
					record[columns[i]] = col
				}
			}
		}
		result = append(result, record)
		if len(result) > 1 {
			return nil, errors.New("expect 1 row,actual 2 more!")
		}

	}

	if len(result) == 0 {
		return nil, errors.New("expect 1 row,actual 0!")
	}
	return result[0], nil
}

func (client DataAccessMgr) QueryAllValue(sql string, args ...interface{}) ([]interface{}, error) {
	rows, err := client.Db.Query(sql, args...)

	if err != nil {
		return nil, err
	}

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	if len(columns) != 1 {
		return nil, errors.New(fmt.Sprintf("expect 1 column, actual %d", len(columns)))
	}

	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	result := []interface{}{}

	for rows.Next() {
		//将行数据保存到record字典
		_ = rows.Scan(scanArgs...)
		record := []interface{}{}
		for _, col := range values {
			if col != nil {
				switch col.(type) {
				case []uint8:
					record = append(record, string(col.([]byte)))
				default:
					record = append(record, col)
				}
			}
		}
		result = append(result, record)

	}
	return result, nil
}

func (client DataAccessMgr) QueryValue(sql string, args ...interface{}) (interface{}, error) {
	rows, err := client.Db.Query(sql, args...)

	if err != nil {
		return nil, err
	}

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	if len(columns) != 1 {
		return nil, errors.New(fmt.Sprintf("expect 1 column, actual %d", len(columns)))
	}

	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	result := []interface{}{}

	for rows.Next() {
		//将行数据保存到record字典
		_ = rows.Scan(scanArgs...)
		record := []interface{}{}
		for _, col := range values {
			if col != nil {
				switch col.(type) {
				case []uint8:
					record = append(record, string(col.([]byte)))
				default:
					record = append(record, col)
				}
			}
		}
		result = append(result, record)
		if len(result) > 1 {
			return nil, errors.New("expect 1 row,actual 2 more!")
		}

	}

	if len(result) == 0 {
		return nil, errors.New("expect 1 row,actual 0!")
	}
	return result[0], nil
}

func (client DataAccessMgr) Insert(sql string, args ...interface{}) (int64, error) {
	stmt, err := client.Db.Prepare(sql)

	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(args...) // 时间可以传time类型，也可以传string类型

	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (client DataAccessMgr) Exec(sql string, args ...interface{}) (int64, error) {
	ret, err := client.Db.Exec(sql, args...)

	if err != nil {
		return 0, err
	}

	return ret.RowsAffected()
}

func (client DataAccessMgr) Close() error {
	return client.Db.Close()
}
