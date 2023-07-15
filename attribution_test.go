package weatherkit

import (
	"testing"
)

func TestAttributionRequestFullUrlGeneration(t *testing.T) {
	req := AttributionRequest{
		Language: "en",
	}

	want := BaseUrl + "/attribution/en"
	have := req.url()

	if want != have {
		t.Errorf("want: %s, have: %s", want, have)
	}
}
