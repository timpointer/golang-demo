package time

import (
	"reflect"
	"testing"
	"time"
)

func TestGetListMonth(t *testing.T) {
	type args struct {
		start time.Time
		end   time.Time
	}
	arg := args{
		time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC),
		time.Date(2010, 2, 10, 23, 0, 0, 0, time.UTC),
	}
	arg2 := args{
		time.Date(2009, 5, 10, 23, 0, 0, 0, time.UTC),
		time.Date(2009, 7, 10, 23, 0, 0, 0, time.UTC),
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"first", arg, []string{"200911", "200912", "201001", "201002"}},
		{"second", arg2, []string{"200905", "200906", "200907"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetListMonth(tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetListMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}
