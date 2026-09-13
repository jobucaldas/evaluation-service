package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	app := &App{}
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rr := httptest.NewRecorder()

	app.healthHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rr.Code)
	}
}

func TestEvaluationLogicDisabledFlag(t *testing.T) {
	app := &App{}
	info := &CombinedFlagInfo{
		Flag: &Flag{Name: "checkout", IsEnabled: false},
	}

	if app.runEvaluationLogic(info, "user-1") {
		t.Fatal("disabled flags must evaluate to false")
	}
}

func TestEvaluationLogicEnabledFlagWithoutRule(t *testing.T) {
	app := &App{}
	info := &CombinedFlagInfo{
		Flag: &Flag{Name: "checkout", IsEnabled: true},
	}

	if !app.runEvaluationLogic(info, "user-1") {
		t.Fatal("enabled flags without targeting rules should evaluate to true")
	}
}

func TestDeterministicBucketRange(t *testing.T) {
	bucket := getDeterministicBucket("user-1checkout")

	if bucket < 0 || bucket > 99 {
		t.Fatalf("expected bucket between 0 and 99, got %d", bucket)
	}
}
