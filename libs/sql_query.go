package libs

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InitMysql(c *Config) *sql.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=10s",
		c.MysqlUser, c.MysqlPassWord, c.MysqlHost, c.MysqlPort, c.MysqlName,
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

type Mysql struct {
	Sql string
}

func NewMySql() *Mysql {
	return &Mysql{}
}

func (m Mysql) GetRow(db *sql.DB) *sql.Rows {
	res, err := db.Query(m.Sql)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func (m Mysql) GetQuerySet(db *sql.DB) [][]string {
	res := m.GetRow(db)
	defer res.Close()
	columns, _ := res.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]sql.RawBytes, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	var results [][]string
	for res.Next() {
		err := res.Scan(scanArgs...)
		if err != nil {
			log.Fatal(err)
		}
		var row []string
		for _, v := range values {
			row = append(row, string(v))
		}
		results = append(results, row)
	}
	return results
}

func (m Mysql) GetColumnName(db *sql.DB) []string {
	row := m.GetRow(db)
	defer row.Close()
	columns, _ := row.Columns()
	return columns
}
