package repository

import (
	"context"
	"database/sql"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/laughingstocK/go-crud/models"
)

func Test_mariadbAuthorRepo_GetByID(t *testing.T) {
	// Create a test database connection (you might want to use a test database).
	// db, err := sql.Open("mysql", "admin:password@tcp(localhost:3306)/gocrud")
	// if err != nil {
	// 	t.Fatalf("Failed to connect to the database: %v", err)
	// }
	// defer db.Close()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock DB: %v", err)
	}
	defer db.Close()
	time := time.Now()
	expectedID := 1
	expectedName := "John Doe"
	expectedCreatedAt := time
	expectedUpdatedAt := time
	rows := sqlmock.NewRows([]string{"id", "name", "createdAt", "updatedAt"}).
		AddRow(expectedID, expectedName, expectedCreatedAt, expectedUpdatedAt)

	// Expect the query to be executed and return the mock rows
	mock.ExpectQuery("SELECT id, name, createdAt, updatedAt FROM Author WHERE id = ?").
		WithArgs(expectedID).
		WillReturnRows(rows)

	notFoundID := int64(0)
	mock.ExpectQuery("SELECT id, name, createdAt, updatedAt FROM Author WHERE id = ?").
		WithArgs(notFoundID).
		WillReturnError(sql.ErrNoRows)

	mock.ExpectQuery("SELECT id, name, createdAt, updatedAt FROM Author WHERE id = ?").
		WithArgs(expectedID).
		WillReturnError(errors.New("connection refused"))

	type fields struct {
		Conn *sql.DB
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Author
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"Success",
			fields{
				Conn: db,
			},
			args{
				ctx: context.Background(),
				id:  1,
			},
			&models.Author{
				ID:        int64(expectedID),
				Name:      expectedName,
				CreatedAt: expectedCreatedAt,
				UpdatedAt: expectedUpdatedAt,
			},
			false,
		},
		{
			"Not Found",
			fields{
				Conn: db,
			},
			args{
				ctx: context.Background(),
				id:  notFoundID,
			},
			nil,
			true,
		},
		{
			"Connection refused",
			fields{
				Conn: db,
			},
			args{
				ctx: context.Background(),
				id:  int64(expectedID),
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mariadbAuthorRepo{
				Conn: tt.fields.Conn,
			}

			got, err := m.GetByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("mariadbAuthorRepo.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mariadbAuthorRepo.GetByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
