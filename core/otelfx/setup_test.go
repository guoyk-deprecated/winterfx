package otelfx

import (
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel"
	"testing"
)

func TestSetupOTEL(t *testing.T) {
	require.NoError(t, Setup())
	_ = otel.GetTracerProvider()
}
