package notes_test

import (
	"context"
	"testing"

	model "github.com/ZBox-Notes/ZBox/backend/models/generated_model"

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
