package model

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	bookFieldNames          = builderx.RawFieldNames(&Book{})
	bookRows                = strings.Join(bookFieldNames, ",")
	bookRowsExpectAutoSet   = strings.Join(stringx.Remove(bookFieldNames, "`create_time`", "`update_time`"), ",")
	bookRowsWithPlaceHolder = strings.Join(stringx.Remove(bookFieldNames, "`book`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	BookModel interface {
		Insert(data Book) (sql.Result, error)
		FindOne(book string) (*Book, error)
		Update(data Book) error
		Delete(book string) error
	}

	defaultBookModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Book struct {
		Book  string `db:"book"`  // book name
		Price int64  `db:"price"` // book price
	}
)

func NewBookModel(conn sqlx.SqlConn) BookModel {
	return &defaultBookModel{
		conn:  conn,
		table: "`book`",
	}
}

func (m *defaultBookModel) Insert(data Book) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, bookRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Book, data.Price)
	return ret, err
}

func (m *defaultBookModel) FindOne(book string) (*Book, error) {
	query := fmt.Sprintf("select %s from %s where `book` = ? limit 1", bookRows, m.table)
	var resp Book
	err := m.conn.QueryRow(&resp, query, book)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultBookModel) Update(data Book) error {
	query := fmt.Sprintf("update %s set %s where `book` = ?", m.table, bookRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Price, data.Book)
	return err
}

func (m *defaultBookModel) Delete(book string) error {
	query := fmt.Sprintf("delete from %s where `book` = ?", m.table)
	_, err := m.conn.Exec(query, book)
	return err
}
