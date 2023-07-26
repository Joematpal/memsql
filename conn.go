package memsql

import (
	"context"
	"database/sql/driver"
	"fmt"
)

var _ driver.ConnPrepareContext = &Conn{}

var _ driver.ConnBeginTx = &Conn{}

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
	return nil, fmt.Errorf("Query method not implemented")
}
