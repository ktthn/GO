package main

import "testing"

// автоматичсекая генерация юнит тестов для функций в голэнде
func Test_equal(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"two slices are equal",
			args{
				a: []int{5, 7, 8},
				b: []int{5, 7, 8},
			},
			true},

		{"two slices are not equal",
			args{
				a: []int{5, 7, 8},
				b: []int{5, 9, 4},
			},
			false},

		{"two slices are not equal in length",
			args{
				a: []int{5, 7, 8},
				b: []int{5, 7, 8, 0},
			},
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equal(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("equal() = %v, want %v", got, tt.want)
			}
		})
	}
}
