package templ

var RepositoryTest = `package repository

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_{{.DomainCamelCase}}Repository_Foo(t *testing.T) {
	type fields struct {
		db sqlmock.Sqlmock
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		mock    func(f fields)
		args    args
		wantErr error
	}{
		{
			name: "Success to retrieve query result",
			mock: func(f fields) {
				f.db.ExpectQuery("SELECT 1").
					WillReturnRows(
						sqlmock.NewRows([]string{"column"}).AddRow(1),
					)
			},
			args: args{
				ctx: context.TODO(),
			},
			wantErr: nil,
		},
		{
			name: "Error when retrieve query result",
			mock: func(f fields) {
				f.db.ExpectQuery("SELECT 1").
					WillReturnError(assert.AnError)
			},
			args: args{
				ctx: context.TODO(),
			},
			wantErr: Err{{.VarErr1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a gorm stub database connection", err)
			}

			f := fields{
				db: mock,
			}
			tt.mock(f)

			{{.DomainShort}} := &{{.DomainCamelCase}}Repository{
				db: gormDB,
			}

			got := {{.DomainShort}}.Foo(tt.args.ctx)
			assert.ErrorIs(t, got, tt.wantErr)
		})
	}
}
`
