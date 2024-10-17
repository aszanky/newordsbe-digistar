package repository

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
)

func Test_repository_AddNewWords(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db, mock := NewMock()
	defer mock.ExpectClose()

	type args struct {
		word      string
		indonesia string
		notes     string
	}
	tests := []struct {
		name    string
		mockFn  func(a args)
		args    args
		wantErr bool
	}{
		{
			name: "Add Success",
			mockFn: func(a args) {
				mock.ExpectExec(regexp.QuoteMeta(QueryAddNewWords)).WithArgs(a.word, a.indonesia, a.notes).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			args: args{
				word:      "كثير",
				indonesia: "Banyak",
				notes:     "Isim jamid",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &repository{
				db: db,
			}
			if err := r.AddNewWords(tt.args.word, tt.args.indonesia, tt.args.notes); (err != nil) != tt.wantErr {
				t.Errorf("repository.AddNewWords() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
