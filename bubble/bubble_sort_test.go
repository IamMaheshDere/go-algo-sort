package bubble_test

import (
	"reflect"
	"testing"

	"github.com/IamMaheshDere/go-algo-sort/bubble"
)

func TestSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test case 1: sort array of 5 integers",
			args{
				arr: []int{3, 5, 1, 4, 2},
			},
			[]int{1, 2, 3, 4, 5},
		},
		{
			"test case 2: sort array of 6 integers having duplicate number",
			args{
				arr: []int{3, 5, 1, 4, 2, 2},
			},
			[]int{1, 2, 2, 3, 4, 5},
		},
		{
			"test case 3: sort already sorted integers",
			args{
				arr: []int{1, 2, 3, 4, 5},
			},
			[]int{1, 2, 3, 4, 5},
		},
		{
			"test case 4: sort empty array",
			args{
				arr: []int{},
			},
			[]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bubble.Sort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}
