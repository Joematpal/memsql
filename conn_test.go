package memsql

import (
	"database/sql/driver"
	"reflect"
	"testing"
)

func TestConn_Query(t *testing.T) {
	type args struct {
		query string
		args  []driver.Value
	}
	tests := []struct {
		name    string
		c       *Conn
		args    args
		want    driver.Rows
		wantErr bool
	}{
		{
			name: "should pass",
			c:    &Conn{},
			args: args{
				query: "select * from foo",
				args:  []driver.Value{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Query(tt.args.query, tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Conn.Query() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Conn.Query() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConn_Exec(t *testing.T) {
	type args struct {
		query string
		args  []driver.Value
	}
	tests := []struct {
		name    string
		c       *Conn
		args    args
		want    driver.Result
		wantErr bool
	}{
		{
			name: "should pass",
			c:    &Conn{},
			args: args{
				query: "insert into foo (substance, mass) values ($1, $2)",
				args:  []driver.Value{},
			},
		},
		{
			name: "should pass",
			c:    &Conn{},
			args: args{
				query: "insert into foo (substance, mass) values (?, ?)",
				args:  []driver.Value{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Exec(tt.args.query, tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Conn.Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Conn.Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}
