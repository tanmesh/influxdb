package context_test

import (
	"context"
	"testing"

	"github.com/influxdata/influxdb/v2"
	icontext "github.com/influxdata/influxdb/v2/context"
	"github.com/influxdata/influxdb/v2/kit/platform"
)

func TestGetAuthorizer(t *testing.T) {
	ctx := context.Background()
	ctx = icontext.SetAuthorizer(ctx, &influxdb.Authorization{
		ID: 1234,
	})
	got, err := icontext.GetAuthorizer(ctx)
	if err != nil {
		t.Errorf("unexpected error while retrieving token: %v", err)
	}

	if want := platform.ID(1234); got.Identifier() != want {
		t.Errorf("GetToken() want %s, got %s", want, got)
	}
}

func TestGetToken(t *testing.T) {
	ctx := context.Background()
	ctx = icontext.SetAuthorizer(ctx, &influxdb.Authorization{
		Token: "howdy",
	})
	got, err := icontext.GetToken(ctx)
	if err != nil {
		t.Errorf("unexpected error while retrieving token: %v", err)
	}

	if want := "howdy"; got != want {
		t.Errorf("GetToken() want %s, got %s", want, got)
	}
}

func TestGetUserID(t *testing.T) {
	ctx := context.Background()
	ctx = icontext.SetAuthorizer(ctx, &influxdb.Authorization{
		UserID: 5678,
	})
	got, err := icontext.GetUserID(ctx)
	if err != nil {
		t.Errorf("unexpected error while retrieving user ID: %v", err)
	}

	if want := platform.ID(5678); got != want {
		t.Errorf("GetUserID() want %s, got %s", want, got)
	}
}
