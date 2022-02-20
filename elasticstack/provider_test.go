package elasticstack_test

import (
	"github.com/vinaybabu16/terraform-provider-elasticstack/elasticstack/acctest"
	"testing"
)

func TestProvider(t *testing.T) {
	if err := acctest.Provider.InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}
