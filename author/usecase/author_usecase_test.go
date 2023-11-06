package usecase

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/laughingstocK/go-crud/author"
	"github.com/laughingstocK/go-crud/models"
	"github.com/stretchr/testify/mock"
)

// MockMariadbAuthorRepo is a mock implementation of author.Repository for MariaDB.
type MockMariadbAuthorRepo struct {
	mock.Mock
}

// GetByID is a mocked method for getting an author by ID.
func (m *MockMariadbAuthorRepo) GetByID(ctx context.Context, id int64) (*models.Author, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*models.Author), args.Error(1)
}

// Create is a mocked method for getting an author by ID.
func (m *MockMariadbAuthorRepo) Create(ctx context.Context, name string) (*models.Author, error) {
	args := m.Called(ctx, name)
	return args.Get(0).(*models.Author), args.Error(1)
}

// MockGrpcAuthorRepo is a mock implementation of author.Repository for gRPC.
type MockGrpcAuthorRepo struct {
	mock.Mock
}

// GetByID is a mocked method for getting an author by ID.
func (m *MockGrpcAuthorRepo) GetByID(ctx context.Context, id int64) (*models.Author, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*models.Author), args.Error(1)
}

// GetByID is a mocked method for getting an author by ID.
func (m *MockGrpcAuthorRepo) Create(ctx context.Context, name string) (*models.Author, error) {
	args := m.Called(ctx, name)
	return args.Get(0).(*models.Author), args.Error(1)
}

func Test_authorUsecase_GetByID(t *testing.T) {
	type fields struct {
		mariadbAuthorRepo author.Repository
		grpcAuthorRepo    author.Repository
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	// Create mock implementations for both repositories
	mockMariadbRepo := &MockMariadbAuthorRepo{}
	mockGrpcRepo := &MockGrpcAuthorRepo{}

	// Inject the mock repositories into the authorUsecase
	usecase := &authorUsecase{
		mariadbAuthorRepo: mockMariadbRepo,
		grpcAuthorRepo:    mockGrpcRepo,
	}

	// Set up expectations for the mariadbRepo mock
	expectedID := int64(1)
	expectedTime := time.Now()
	expectedAuthor := &models.Author{ID: expectedID, Name: "John Doe", CreatedAt: expectedTime, UpdatedAt: expectedTime}
	mockMariadbRepo.On("GetByID", mock.Anything, expectedID).Return(expectedAuthor, nil)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Author
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Success",
			fields: fields{
				mariadbAuthorRepo: usecase.mariadbAuthorRepo,
				grpcAuthorRepo:    usecase.grpcAuthorRepo,
			},
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			want:    expectedAuthor,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &authorUsecase{
				mariadbAuthorRepo: tt.fields.mariadbAuthorRepo,
				grpcAuthorRepo:    tt.fields.grpcAuthorRepo,
			}
			got, err := a.GetByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("authorUsecase.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("authorUsecase.GetByID() = %v, want %v", got, tt.want)
			}
			mockMariadbRepo.AssertExpectations(t)
			mockGrpcRepo.AssertExpectations(t)
		})
	}
}
