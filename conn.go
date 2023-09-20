package memsql

import (
	"context"
	"database/sql/driver"
	"fmt"
	"log"
	"regexp"

	"github.com/auxten/postgresql-parser/pkg/sql/parser"
	"github.com/auxten/postgresql-parser/pkg/sql/sem/tree"
	"github.com/auxten/postgresql-parser/pkg/walk"
)

type FactoryConn interface {
	driver.Conn
	driver.ConnBeginTx
	driver.ConnPrepareContext
	// I know this is deprecated it is here for compatability
	driver.Execer
	driver.ExecerContext
	driver.QueryerContext
	driver.Pinger
	// driver.Validator
}

var _ FactoryConn = &Conn{}

type Conn struct {
}

func (c *Conn) BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error) {
	return nil, fmt.Errorf("BeginTx method not implemented")
}

func (c *Conn) Begin() (driver.Tx, error) {
	return nil, fmt.Errorf("Begin method not implemented")
}

func (c *Conn) Close() error {
	return fmt.Errorf("close method not implemented")
}

// implements driver.Conn interface
// Prepare - not implemented ...
func (c *Conn) Prepare(query string) (driver.Stmt, error) {
	return nil, fmt.Errorf("Prepare method not implemented")
}

// implements driver.Conn interface
// PrepareContext - not implemented ...
func (c *Conn) PrepareContext(ctx context.Context, query string) (driver.Stmt, error) {
	return nil, fmt.Errorf("PrepareContext method not implemented")
}

// implements driver.Tx interface
// Rollback - not implemented ...
func (c *Conn) Rollback() error {
	return fmt.Errorf("Rollback method not implemented")
}

// Queryer interface
func (c *Conn) Query(query string, args []driver.Value) (driver.Rows, error) {
	stmts, err := parser.Parse(query)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("stmts: %#v\n", stmts)
	w := &walk.AstWalker{
		Fn: func(ctx interface{}, node interface{}) (stop bool) {
			log.Printf("node type %T", node)
			return false
		},
	}
	_, _ = w.Walk(stmts, nil)
	return nil, fmt.Errorf("Query method not implemented")
}

var odbcParamsRE = regexp.MustCompile(`\?`)

func (c *Conn) Exec(query string, args []driver.Value) (driver.Result, error) {
	i := 0
	query = odbcParamsRE.ReplaceAllStringFunc(query, func(s string) string {
		i++
		return fmt.Sprintf("$%d", i)
	})

	stmts, err := parser.Parse(query)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("stmts: %#v\n", stmts)
	w := &walk.AstWalker{
		Fn: func(ctx interface{}, node interface{}) (stop bool) {
			log.Printf("node type %T", node)
			switch v := node.(type) {
			case *tree.Insert:
				fmt.Println(v.Columns)
				fmt.Printf("rows select type: %T %+v\n", v.Rows.Select, v.Rows.Select)
				fmt.Printf("rows select type: %T %+v\n", v.Rows.Select.(*tree.ValuesClause), v.Rows.Select.(*tree.ValuesClause).Rows)
			}
			return false
		},
	}
	_, err = w.Walk(stmts, nil)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (c *Conn) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	return nil, nil
}

func (c *Conn) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	return nil, nil
}

func (c *Conn) Ping(ctx context.Context) error {
	return nil
}
