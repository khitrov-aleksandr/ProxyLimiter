package base

import (
	"testing"

	"github.com/khitrov-aleksandr/proxyguard/test"
)

func TestCheck(t *testing.T) {
	mockRepo := test.NewMockRepository()
	rl := NewRateLimiter(mockRepo)

	key := "testKey"
	count := int64(5)
	timeout := 5

	result := rl.Allow(key, count, timeout)

	if result != true {
		t.Errorf("Expected true, got %t", result)
	}

	mockRepo.IncrValue = count - 1

	result = rl.Allow(key, count, timeout)

	if result != true {
		t.Errorf("Expected true, got %t", result)
	}

	mockRepo.IncrValue++

	result = rl.Allow(key, count, timeout)

	if result != false {
		t.Errorf("Expected false, got %t", result)
	}
}
