package memsql

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"reflect"
	"testing"
)

func TestDriver_Open(t *testing.T) {
	type args struct {
		name       string
		dataSource string
	}

	tests := []struct {
		name    string
		d       *Driver
		args    args
		want    driver.Conn
		wantErr bool
	}{
		{
			name: "should pass",
			args: args{
				name: "memsql",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sql.Open(tt.args.name, tt.args.dataSource)
			if (err != nil) != tt.wantErr {
				t.Errorf("Driver.Open() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Printf("%#v\n", got)
			got.Ping()
			// got.QueryRow()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Driver.Open() = %v, want %v", got, tt.want)
			}
		})
	}
}
