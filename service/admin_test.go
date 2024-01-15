package service

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"hello-cafe/internal/api"
	"hello-cafe/internal/db"
	"hello-cafe/repository"
)

func Test_adminService_SignIn(t *testing.T) {
	repo, err := repository.NewRepository()
	require.NoError(t, err)

	err = os.Setenv("CONFIG_PATH", "/Users/shyu/workspace/hello-cafe/configs/config.yml")
	require.NoError(t, err)

	cfg, err := api.Config()
	require.NoError(t, err)

	err = db.Connect(cfg.DB)
	require.NoError(t, err)

	type fields struct {
		repo repository.Repository
	}
	type args struct {
		phone    string
		password string
	}
	tests := []struct {
		name            string
		fields          fields
		args            args
		wantAccessToken string
		wantErr         bool
	}{
		{
			name: "로그인 성공",
			fields: fields{
				repo: repo,
			},
			args: args{
				phone:    "010-1234-1111",
				password: "12341234",
			},
			wantErr: false,
		},
		{
			name: "로그인 실패(핸드폰번호)",
			fields: fields{
				repo: repo,
			},
			args: args{
				phone:    "010-1234-1111123",
				password: "12341234",
			},
			wantErr: true,
		},
		{
			name: "로그인 실패(비밀번호)",
			fields: fields{
				repo: repo,
			},
			args: args{
				phone:    "010-1234-1111",
				password: "12341234123",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &adminService{
				repo: tt.fields.repo,
			}
			_, err := s.SignIn(tt.args.phone, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignIn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_adminService_SignOut(t *testing.T) {
	repo, err := repository.NewRepository()
	require.NoError(t, err)

	err = os.Setenv("CONFIG_PATH", "/Users/shyu/workspace/hello-cafe/configs/config.yml")
	require.NoError(t, err)

	cfg, err := api.Config()
	require.NoError(t, err)

	err = db.Connect(cfg.DB)
	require.NoError(t, err)

	type fields struct {
		repo repository.Repository
	}
	type args struct {
		phone string
		token string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "로그아웃 실패",
			fields: fields{
				repo: repo,
			},
			args: args{
				phone: "010-1234-1234",
				token: "invalid token",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &adminService{
				repo: tt.fields.repo,
			}
			if err := s.SignOut(tt.args.phone, tt.args.token); (err != nil) != tt.wantErr {
				t.Errorf("SignOut() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_adminService_SignUp(t *testing.T) {
	repo, err := repository.NewRepository()
	require.NoError(t, err)

	err = os.Setenv("CONFIG_PATH", "/Users/shyu/workspace/hello-cafe/configs/config.yml")
	require.NoError(t, err)

	cfg, err := api.Config()
	require.NoError(t, err)

	err = db.Connect(cfg.DB)
	require.NoError(t, err)

	type fields struct {
		repo repository.Repository
	}
	type args struct {
		phone    string
		password string
		name     string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "회원 가입 실패(핸드폰 번호 미입력)",
			fields: fields{
				repo: repo,
			},
			args: args{
				phone:    "",
				password: "1234",
				name:     "홍길동",
			},
			wantErr: true,
		},
		{
			name: "회원 가입 실패(비밀 번호 미입력)",
			fields: fields{
				repo: repo,
			},
			args: args{
				phone:    "010-1234-1234",
				password: "",
				name:     "홍길동",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &adminService{
				repo: tt.fields.repo,
			}
			if err := s.SignUp(tt.args.phone, tt.args.password, tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("SignUp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
