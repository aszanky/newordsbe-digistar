package repository

import (
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

// NewMock mocking database
func NewMock() (*sqlx.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	//convert sqlDB to sqlxDB
	sqlxDB := sqlx.NewDb(db, "postgres")

	return sqlxDB, mock
}

func TestNewRepository(t *testing.T) {
	db, mock := NewMock()
	defer mock.ExpectClose()

	type args struct {
		db *sqlx.DB
	}
	tests := []struct {
		name string
		args args
		want Repository
	}{
		{
			name: "success initialization",
			args: args{
				db,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRepository(tt.args.db); got == nil {
				t.Errorf("NewRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}
