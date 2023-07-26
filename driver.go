package memsql

import (
	"database/sql/driver"
)

type Driver struct {
}

// implements sql.Driver interface
// Open - open file for reading, return the connection
func (d *Driver) Open(name string) (driver.Conn, error) {

	return &Conn{}, nil
}
