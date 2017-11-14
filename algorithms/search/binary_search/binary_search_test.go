package binary_search

import "testing"

func Test_binary_search(t *testing.T) {
	list := []int{1, 3, 5, 7, 9}
	type args struct {
		list []int
		item int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"find", args{list, 5}, 2},
		{"find first", args{list, 1}, 0},
		{"find last", args{list, 9}, 4},
		{"not find", args{list, 4}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binary_search(tt.args.list, tt.args.item); got != tt.want {
				t.Errorf("binary_search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_binary_search(b *testing.B) {
	list := []int{1, 3, 5, 7, 9}
	for i := 0; i < b.N; i++ {
		binary_search(list, 5)
	}
}
