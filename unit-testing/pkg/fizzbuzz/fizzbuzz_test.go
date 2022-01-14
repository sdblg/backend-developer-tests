package fizzbuzz

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	type args struct {
		total  int64
		fizzAt int64
		buzzAt int64
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Default",
			args: args{
				total:  16,
				fizzAt: 3,
				buzzAt: 5,
			},
			want: []string{
				"1",
				"2",
				"Fizz",
				"4",
				"Buzz",
				"Fizz",
				"7",
				"8",
				"Fizz",
				"Buzz",
				"11",
				"Fizz",
				"13",
				"14",
				"FizzBuzz",
				"16",
			},
		},
		{
			name: "Edge Case - fizzAt equals 0",
			args: args{
				total:  16,
				fizzAt: 0,
				buzzAt: 5,
			},
			want: []string{},
		},
		{
			name: "Edge Case - buzzAt equals 0",
			args: args{
				total:  16,
				fizzAt: 3,
				buzzAt: 0,
			},
			want: []string{},
		},
		{
			name: "Edge Case - fizzAt and buzzAt equal 0",
			args: args{
				total:  16,
				fizzAt: 0,
				buzzAt: 0,
			},
			want: []string{},
		},
		{
			name: "Edge Case - total 0",
			args: args{
				total:  -1,
				fizzAt: 3,
				buzzAt: 5,
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FizzBuzz(tt.args.total, tt.args.fizzAt, tt.args.buzzAt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FizzBuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}
