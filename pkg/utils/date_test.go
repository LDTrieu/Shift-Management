package utils

import "testing"

func TestDateStringToInt(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "test 1",
			args: args{date: "2021-07-15"},
			want: 20210715,
		},
		{
			name: "test 2",
			args: args{date: "2022-08-01"},
			want: 20220801,
		},
		{
			name: "error",
			args: args{date: "not-a-date-string"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DateStringToInt(tt.args.date); got != tt.want {
				t.Errorf("DateStringToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}