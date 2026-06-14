package presenter_test

import (
	"testing"
	"time"
	"trec/internal/adapter/presenter"

	"github.com/stretchr/testify/require"
)

func TestDurationString(t *testing.T) {
	require.Equal(t, "1:23:45", presenter.ToDurationString(1*time.Hour+23*time.Minute+45*time.Second))
}
