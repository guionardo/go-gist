package logger

import "testing"

func TestConsoleColor_Paint(t *testing.T) {
	const text = "Test"

	tests := []struct {
		name  string
		color int
		want  string
	}{
		{
			name:  "Reset",
			color: Reset,
			want:  "\033[0mTest\033[0m",
		},
		{
			name:  "Red",
			color: Red,
			want:  "\033[31mTest\033[0m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Colors.Paint(text, tt.color); got != tt.want {
				t.Errorf("ConsoleColor.Paint() = %v, want %v", got, tt.want)
			}
		})
	}
}
