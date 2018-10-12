package types

import "testing"

func TestSqsRequestsPrice(t *testing.T) {
	sqs := NewConfigFromFiles().Sqs

	if sqs.Price.Requests.Per != 1000000.0 {
		t.Errorf("Per was incorrect, got: %f, want: %f", sqs.Price.Requests.Per, 1000000.0)
	}

	switch {
	case sqs.Price.Requests.Standard <= 0:
		t.Errorf("Standard was incorrect, got: %f, want: %s", sqs.Price.Requests.Standard, "in (0,1)")
	case sqs.Price.Requests.Standard >= 1:
		t.Errorf("Standard was incorrect, got: %f, want: %s", sqs.Price.Requests.Standard, "in (0,1)")
	}

	switch {
	case sqs.Price.Requests.Fifo <= 0:
		t.Errorf("Fifo was incorrect, got: %f, want: %s", sqs.Price.Requests.Fifo, "in (0,1)")
	case sqs.Price.Requests.Fifo >= 1:
		t.Errorf("Fifo was incorrect, got: %f, want: %s", sqs.Price.Requests.Fifo, "in (0,1)")
	}
}
