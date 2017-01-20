package repository

import (
	"fmt"
	"github.com/hapoon/gendao/dao"
	"github.com/jinzhu/gorm"
	"github.com/hapoon/gendao/log"
)

type Columns struct {
	DB   *gorm.DB
	Data []Column
}

type Column struct {
	Name string
	Data []dao.Column
}

func (rcv *Columns) New(db *gorm.DB) {
	rcv.DB = db
}

func (rcv *Columns) FindAll(table string) error {
	query := fmt.Sprintf("SHOW COLUMNS FROM %s", table)
	rows, err := rcv.DB.Raw(query).Rows()
	if err != nil {
		return err
	}
	defer rows.Close()
	log.Debugf("table: %v\n",table)
	col := Column{Name:table}
	for rows.Next() {
		column := dao.Column{}
		rows.Scan(&column.Field, &column.Type, &column.Null, &column.Key, &column.Default, &column.Extra)
		log.Debugf("column: %+v\n", column)
		col.Data = append(col.Data, column)
	}
	rcv.Data = append(rcv.Data, col)
	return nil
}
