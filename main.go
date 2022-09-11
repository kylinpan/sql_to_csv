package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/sql_query/libs"
)

func WriteFile(quey_set [][]string) error {
	n := fmt.Sprintf("result-%s.csv", time.Now().Format("20060102150405"))
	f, err := os.Create(n)
	if err != nil {
		return err
	}
	defer f.Close()
	w := csv.NewWriter(f)
	defer w.Flush()
	for _, i := range quey_set {
		if err := w.Write(i); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	conf := libs.NewConfig()
	conf.LoadConfig("./config.json")
	db := libs.InitMysql(conf)
	defer db.Close()
	query := libs.NewMySql()
	query.Sql = conf.Query
	query_set := query.GetQuerySet(db)
	query_set = append([][]string{query.GetColumnName(db)}, query_set...)
	if err := WriteFile(query_set); err != nil {
		panic(err)
	}
}
