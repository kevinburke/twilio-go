package token

import (
	"strings"
	"sync"
	"testing"
	"time"
)

func TestAccessTokenJWTConcurrent(t *testing.T) {
	t.Parallel()
	var wg sync.WaitGroup
	for range 20 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			at := New("AC123", "SK456", "secretkey", "test@example.com", time.Hour)
			at.AddGrant(NewConversationsGrant("IS123"))
			jwt, err := at.JWT()
			if err != nil {
				t.Errorf("JWT: %v", err)
				return
			}
			if parts := strings.Split(jwt, "."); len(parts) != 3 {
				t.Errorf("JWT had %d parts, want 3", len(parts))
			}
		}()
	}
	wg.Wait()
}
