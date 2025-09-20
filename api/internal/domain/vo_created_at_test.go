package domain_test

import (
	"testing"
	"time"

	"github.com/iotassss/fushinsha-map-api/internal/domain"
)

func TestNewCreatedAt(t *testing.T) {
	now := time.Date(2023, 9, 19, 12, 34, 56, 0, time.UTC)
	ca, _ := domain.NewCreatedAt(now)
	if !ca.Time.Equal(now) {
		t.Errorf("NewCreatedAt: want %v, got %v", now, ca.Time)
	}
}

func TestCreatedAtString(t *testing.T) {
	now := time.Date(2023, 9, 19, 12, 34, 56, 0, time.UTC)
	ca, _ := domain.NewCreatedAt(now)
	want := now.Format(time.RFC3339)
	got := ca.String()
	if got != want {
		t.Errorf("CreatedAt.String: want %v, got %v", want, got)
	}
}
