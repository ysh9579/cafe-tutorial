package service

import (
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"hello-cafe/internal/api"
	"hello-cafe/internal/db"
	"hello-cafe/model"
	"hello-cafe/model/request"
	"hello-cafe/repository"
)

func Test_itemService_Create(t *testing.T) {
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
		item request.CreateItem
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "필수 정보 미입력",
			fields: fields{
				repo: repo,
			},
			args: args{
				item: request.CreateItem{
					AdminSeq:    0,
					Category:    new(request.ItemCategory),
					Barcode:     new(string),
					Price:       new(int64),
					Cost:        new(int64),
					Name:        new(string),
					Description: new(string),
					ExpireDT:    new(time.Time),
					Size:        new(request.ItemSize),
				},
			},
			wantErr: true,
		},
		{
			name: "필수 정보 미입력",
			fields: fields{
				repo: repo,
			},
			args: args{
				item: request.CreateItem{
					AdminSeq:    1,
					Category:    nil,
					Barcode:     new(string),
					Price:       new(int64),
					Cost:        new(int64),
					Name:        new(string),
					Description: new(string),
					ExpireDT:    new(time.Time),
					Size:        new(request.ItemSize),
				},
			},
			wantErr: true,
		},
		{
			name: "필수 정보 미입력",
			fields: fields{
				repo: repo,
			},
			args: args{
				item: request.CreateItem{
					AdminSeq:    1,
					Category:    new(request.ItemCategory),
					Barcode:     nil,
					Price:       new(int64),
					Cost:        new(int64),
					Name:        new(string),
					Description: new(string),
					ExpireDT:    new(time.Time),
					Size:        new(request.ItemSize),
				},
			},
			wantErr: true,
		},
		{
			name: "필수 정보 미입력",
			fields: fields{
				repo: repo,
			},
			args: args{
				item: request.CreateItem{
					AdminSeq:    1,
					Category:    new(request.ItemCategory),
					Barcode:     new(string),
					Price:       nil,
					Cost:        new(int64),
					Name:        new(string),
					Description: new(string),
					ExpireDT:    new(time.Time),
					Size:        new(request.ItemSize),
				},
			},
			wantErr: true,
		},
		{
			name: "필수 정보 미입력",
			fields: fields{
				repo: repo,
			},
			args: args{
				item: request.CreateItem{
					AdminSeq:    1,
					Category:    new(request.ItemCategory),
					Barcode:     new(string),
					Price:       new(int64),
					Cost:        nil,
					Name:        new(string),
					Description: new(string),
					ExpireDT:    new(time.Time),
					Size:        new(request.ItemSize),
				},
			},
			wantErr: true,
		},
		{
			name: "필수 정보 미입력",
			fields: fields{
				repo: repo,
			},
			args: args{
				item: request.CreateItem{
					AdminSeq:    1,
					Category:    new(request.ItemCategory),
					Barcode:     new(string),
					Price:       new(int64),
					Cost:        new(int64),
					Name:        nil,
					Description: new(string),
					ExpireDT:    new(time.Time),
					Size:        new(request.ItemSize),
				},
			},
			wantErr: true,
		},
		{
			name: "필수 정보 미입력",
			fields: fields{
				repo: repo,
			},
			args: args{
				item: request.CreateItem{
					AdminSeq:    1,
					Category:    new(request.ItemCategory),
					Barcode:     new(string),
					Price:       new(int64),
					Cost:        new(int64),
					Name:        new(string),
					Description: nil,
					ExpireDT:    new(time.Time),
					Size:        new(request.ItemSize),
				},
			},
			wantErr: true,
		},
		{
			name: "필수 정보 미입력",
			fields: fields{
				repo: repo,
			},
			args: args{
				item: request.CreateItem{
					AdminSeq:    1,
					Category:    new(request.ItemCategory),
					Barcode:     new(string),
					Price:       new(int64),
					Cost:        new(int64),
					Name:        new(string),
					Description: new(string),
					ExpireDT:    nil,
					Size:        new(request.ItemSize),
				},
			},
			wantErr: true,
		},
		{
			name: "필수 정보 미입력",
			fields: fields{
				repo: repo,
			},
			args: args{
				item: request.CreateItem{
					AdminSeq:    1,
					Category:    new(request.ItemCategory),
					Barcode:     new(string),
					Price:       new(int64),
					Cost:        new(int64),
					Name:        new(string),
					Description: new(string),
					ExpireDT:    new(time.Time),
					Size:        nil,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &itemService{
				repo: tt.fields.repo,
			}
			if err := s.Create(tt.args.item); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_itemService_Delete(t *testing.T) {
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
		itemSeq int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "삭제 실패(item_seq 누락)",
			fields: fields{
				repo: repo,
			},
			args: args{
				itemSeq: 0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &itemService{
				repo: tt.fields.repo,
			}
			if err := s.Delete(tt.args.itemSeq); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_itemService_Find(t *testing.T) {
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
		adminSeq    int64
		lastItemSeq int64
		limit       int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.Items
		wantErr bool
	}{
		{
			name: "조회 실패(관리자 정보 누락)",
			fields: fields{
				repo: repo,
			},
			args: args{
				adminSeq: 0,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &itemService{
				repo: tt.fields.repo,
			}
			got, err := s.Find(tt.args.adminSeq, tt.args.lastItemSeq, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_itemService_Get(t *testing.T) {
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
		itemSeq int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Item
		wantErr bool
	}{
		{
			name: "조회 실패(item_seq 누락)",
			fields: fields{
				repo: repo,
			},
			args: args{
				itemSeq: 0,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &itemService{
				repo: tt.fields.repo,
			}
			got, err := s.Get(tt.args.itemSeq)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_itemService_Search(t *testing.T) {
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
		adminSeq int64
		text     string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.Items
		wantErr bool
	}{
		{
			name: "검색 실패 ( 관리자 정보 누락 )",
			fields: fields{
				repo: repo,
			},
			args: args{
				adminSeq: 0,
				text:     "ㅋㅍ",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "검색 실패 ( 검색어 누락 )",
			fields: fields{
				repo: repo,
			},
			args: args{
				adminSeq: 1,
				text:     "",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &itemService{
				repo: tt.fields.repo,
			}
			got, err := s.Search(tt.args.adminSeq, tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Search() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_itemService_Update(t *testing.T) {
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
		item request.UpdateItem
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "item_seq  미입력",
			fields: fields{
				repo: repo,
			},
			args: args{
				item: request.UpdateItem{
					ItemSeq:     0,
					Category:    new(request.ItemCategory),
					Barcode:     new(string),
					Price:       new(int64),
					Cost:        new(int64),
					Name:        new(string),
					Description: new(string),
					ExpireDT:    new(time.Time),
					Size:        new(request.ItemSize),
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &itemService{
				repo: tt.fields.repo,
			}
			if err := s.Update(tt.args.item); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
