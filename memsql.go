package memsql

import "database/sql"

func init() {
	sql.Register("memsql", &Driver{})
}
