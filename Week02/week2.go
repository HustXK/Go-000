package main

import (
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

type Date struct {
	Key     string    `json:"key"`
	Value   string    `json:"value"`
}

func  dao() error {
	var data Date
	var DB *sql.DB  //省略初始化
	err := sq.
		Select("key, value").
		From("dateTable").
		Where(sq.Eq{"key": data.Key}).
		RunWith(DB).
		QueryRow().
		Scan(&data.Key, &data.Value)
	if err == sql.ErrNoRows{
		return errors.New("Record does not exist")
	} else if err != nil {
		return err
	}
	return nil
}

func main() {
	err := dao()
	if err != nil {
		_ = fmt.Errorf("%+v\n", err)
	}
	//do some else
}

