package helpers

import "testing"

func TestGetTicketNumberFromString(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "Test", args: args{input: "feature/SPS-294-add-wordings-checkout"}, want: "SPS-294"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTicketNumberFromString(tt.args.input); got != tt.want {
				t.Errorf("GetTicketNumberFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}
