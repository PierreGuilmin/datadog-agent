package sampler

import (
	"testing"

	"github.com/DataDog/datadog-agent/pkg/trace/pb"
	"github.com/stretchr/testify/assert"
)

func testSpan() *pb.Span {
	return &pb.Span{
		Duration: 10000000,
		Error:    0,
		Resource: "GET /some/raclette",
		Service:  "django",
		Name:     "django.controller",
		SpanID:   42,
		Start:    1448466874000000000,
		TraceID:  424242,
		Meta: map[string]string{
			"user": "leo",
			"pool": "fondue",
		},
		Metrics: map[string]float64{
			"cheese_weight": 100000.0,
		},
		ParentID: 1111,
		Type:     "http",
	}
}

func TestSpanString(t *testing.T) {
	assert := assert.New(t)
	assert.NotEqual("", testSpan().String())
}

func TestSpanWeight(t *testing.T) {
	assert := assert.New(t)

	span := testSpan()
	assert.Equal(1.0, Weight(span))

	span.Metrics[KeySamplingRateGlobal] = -1.0
	assert.Equal(1.0, Weight(span))

	span.Metrics[KeySamplingRateGlobal] = 0.0
	assert.Equal(1.0, Weight(span))

	span.Metrics[KeySamplingRateGlobal] = 0.25
	assert.Equal(4.0, Weight(span))

	span.Metrics[KeySamplingRateGlobal] = 1.0
	assert.Equal(1.0, Weight(span))

	span.Metrics[KeySamplingRateGlobal] = 1.5
	assert.Equal(1.0, Weight(span))
}

func TestSpanWeightNil(t *testing.T) {
	assert := assert.New(t)

	var span *pb.Span

	assert.Equal(1.0, Weight(span), "Weight should be callable on nil and return a default value")
}
