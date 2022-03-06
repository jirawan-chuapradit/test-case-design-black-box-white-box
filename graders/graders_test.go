package main

import "testing"

func Test_grader_Test_Grade_A(t *testing.T) {
	type args struct {
		score int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "input below min value then expect not equal A",
			args: args{
				score: 79,
			},
			want: "B",
		},
		{
			name: "input min value then expect A",
			args: args{
				score: 80,
			},
			want: "A",
		},
		{
			name: "input above min value then expect A",
			args: args{
				score: 81,
			},
			want: "A",
		},
		{
			name: "input nominal value then expect A",
			args: args{
				score: 82,
			},
			want: "A",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := grader(tt.args.score); got != tt.want {
				t.Errorf("grader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_grader_Test_Grade_B(t *testing.T) {
	type args struct {
		score int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "input below min value then expect not equal B",
			args: args{
				score: 69,
			},
			want: "C",
		},
		{
			name: "input min value then expect B",
			args: args{
				score: 70,
			},
			want: "B",
		},
		{
			name: "input above min value then expect B",
			args: args{
				score: 71,
			},
			want: "B",
		},
		{
			name: "input nominal value then expect B",
			args: args{
				score: 72,
			},
			want: "B",
		},
		{
			name: "input below max value then expect B",
			args: args{
				score: 78,
			},
			want: "B",
		},
		{
			name: "input max value then expect B",
			args: args{
				score: 79,
			},
			want: "B",
		},
		{
			name: "input above max value then expect not B",
			args: args{
				score: 80,
			},
			want: "A",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := grader(tt.args.score); got != tt.want {
				t.Errorf("grader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_grader_Test_Grade_C(t *testing.T) {
	type args struct {
		score int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "input below min value then expect not equal D",
			args: args{
				score: 59,
			},
			want: "D",
		},
		{
			name: "input min value then expect C",
			args: args{
				score: 60,
			},
			want: "C",
		},
		{
			name: "input above min value then expect C",
			args: args{
				score: 61,
			},
			want: "C",
		},
		{
			name: "input nominal value then expect C",
			args: args{
				score: 62,
			},
			want: "C",
		},
		{
			name: "input below max value then expect C",
			args: args{
				score: 68,
			},
			want: "C",
		},
		{
			name: "input max value then expect C",
			args: args{
				score: 69,
			},
			want: "C",
		},
		{
			name: "input above max value then expect not C",
			args: args{
				score: 70,
			},
			want: "B",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := grader(tt.args.score); got != tt.want {
				t.Errorf("grader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_grader_Test_Grade_D(t *testing.T) {
	type args struct {
		score int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "input below min value then expect not equal D",
			args: args{
				score: 49,
			},
			want: "F",
		},
		{
			name: "input min value then expect D",
			args: args{
				score: 50,
			},
			want: "D",
		},
		{
			name: "input above min value then expect D",
			args: args{
				score: 51,
			},
			want: "D",
		},
		{
			name: "input nominal value then expect D",
			args: args{
				score: 52,
			},
			want: "D",
		},
		{
			name: "input below max value then expect D",
			args: args{
				score: 58,
			},
			want: "D",
		},
		{
			name: "input max value then expect D",
			args: args{
				score: 59,
			},
			want: "D",
		},
		{
			name: "input above max value then expect not D",
			args: args{
				score: 60,
			},
			want: "C",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := grader(tt.args.score); got != tt.want {
				t.Errorf("grader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_grader_Test_Grade_F(t *testing.T) {
	type args struct {
		score int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "input below min value then expect not equal F",
			args: args{
				score: -1,
			},
			want: "invalid score",
		},
		{
			name: "input min value then expect F",
			args: args{
				score: 0,
			},
			want: "F",
		},
		{
			name: "input above min value then expect F",
			args: args{
				score: 1,
			},
			want: "F",
		},
		{
			name: "input nominal value then expect F",
			args: args{
				score: 2,
			},
			want: "F",
		},
		{
			name: "input below max value then expect F",
			args: args{
				score: 48,
			},
			want: "F",
		},
		{
			name: "input max value then expect F",
			args: args{
				score: 49,
			},
			want: "F",
		},
		{
			name: "input above max value then expect not F",
			args: args{
				score: 50,
			},
			want: "D",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := grader(tt.args.score); got != tt.want {
				t.Errorf("grader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_grader_Test_Invalid_Score(t *testing.T) {
	type args struct {
		score int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "input nominal value then expect invalid score",
			args: args{
				score: -3,
			},
			want: "invalid score",
		},
		{
			name: "input below min value then expect invalid score",
			args: args{
				score: -2,
			},
			want: "invalid score",
		},
		{
			name: "input min value then expect invalid score",
			args: args{
				score: -1,
			},
			want: "invalid score",
		},
		{
			name: "input above min value then expect not equal invalid score",
			args: args{
				score: 0,
			},
			want: "F",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := grader(tt.args.score); got != tt.want {
				t.Errorf("grader() = %v, want %v", got, tt.want)
			}
		})
	}
}
