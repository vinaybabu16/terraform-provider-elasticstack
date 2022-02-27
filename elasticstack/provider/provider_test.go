package provider_test

import (
	"testing"

	"github.com/vinaybabu16/terraform-provider-elasticstack/elasticstack/internal/acctest"
)

func TestProvider(t *testing.T) {
	if err := acctest.Provider.InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}
