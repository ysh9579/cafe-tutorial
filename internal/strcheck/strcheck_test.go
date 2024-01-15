package strcheck

import "testing"

func TestValidatePassword(t *testing.T) {
	type args struct {
		pwd string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "8자리 이하",
			args: args{
				pwd: "1234",
			},
			want: false,
		},
		{
			name: "한글이 포함",
			args: args{
				pwd: "1234123ㄱㅇ4",
			},
			want: false,
		},
		{
			name: "성공",
			args: args{
				pwd: "1234123Ab",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidatePassword(tt.args.pwd); got != tt.want {
				t.Errorf("ValidatePassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidatePhone(t *testing.T) {
	type args struct {
		phone string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "성공",
			args: args{
				phone: "010-1234-1234",
			},
			want: true,
		},
		{
			name: "-이 존재 하지 않음(1)",
			args: args{
				phone: "010-12341234",
			},
			want: false,
		},
		{
			name: "-이 존재 하지 않음(2)",
			args: args{
				phone: "0101234-1234",
			},
			want: false,
		},
		{
			name: "자리수가 부족",
			args: args{
				phone: "010-123-234",
			},
			want: false,
		},
		{
			name: "자리수가 넘침",
			args: args{
				phone: "010-12344-12344",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidatePhone(tt.args.phone); got != tt.want {
				t.Errorf("ValidatePhone() = %v, want %v", got, tt.want)
			}
		})
	}
}
