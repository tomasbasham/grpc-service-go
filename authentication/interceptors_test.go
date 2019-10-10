package authentication_test

import (
	"testing"

	"github.com/tomasbasham/grpc-service-go/authentication"
)

func TestHello(t *testing.T) {
	if authentication.Hello() != "hello" {
		t.Errorf("oh nose!")
	}
}
