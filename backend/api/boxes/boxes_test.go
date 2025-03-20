package boxes_test

import (
	model "backend/models/generated_model"
	"context"
	"testing"

	"github.com/pashagolub/pgxmock/v4"
)

func TestNewService(t *testing.T) {
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mock.Close(context.Background())
	model.New(mock)
	t.Skip("Skipping test")
}
