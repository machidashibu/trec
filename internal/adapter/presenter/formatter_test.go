package presenter_test

import (
	"testing"
	"time"
	"trec/internal/adapter/presenter"

	"github.com/stretchr/testify/require"
)

func TestDurationString(t *testing.T) {
	const d1h23m45s = 1*time.Hour + 23*time.Minute + 45*time.Second
	require.Equal(t, "1h23m45s", presenter.NewDurationFormatter("").String(d1h23m45s))
	require.Equal(t, "1h23m45s", presenter.NewDurationFormatter("#h#m#s").String(d1h23m45s))
	require.Equal(t, "1:23:45", presenter.NewDurationFormatter("#:#:#").String(d1h23m45s))
	require.Equal(t, "83m45s", presenter.NewDurationFormatter("#m#s").String(d1h23m45s))
	require.Equal(t, "83:45", presenter.NewDurationFormatter("#:#").String(d1h23m45s))
	require.Equal(t, "83'45''", presenter.NewDurationFormatter("#'#''").String(d1h23m45s))

	const d12m34s = 12*time.Minute + 34*time.Second
	require.Equal(t, "12m34s", presenter.NewDurationFormatter("").String(d12m34s))
	require.Equal(t, "0h12m34s", presenter.NewDurationFormatter("#h#m#s").String(d12m34s))
	require.Equal(t, "0:12:34", presenter.NewDurationFormatter("#:#:#").String(d12m34s))
	require.Equal(t, "12m34s", presenter.NewDurationFormatter("#m#s").String(d12m34s))
	require.Equal(t, "12:34", presenter.NewDurationFormatter("#:#").String(d12m34s))
	require.Equal(t, "12'34''", presenter.NewDurationFormatter("#'#''").String(d12m34s))

	const d12s = 12 * time.Second
	require.Equal(t, "12s", presenter.NewDurationFormatter("").String(d12s))
	require.Equal(t, "12s", presenter.NewDurationFormatter("#s").String(d12s))
	require.Equal(t, "12''", presenter.NewDurationFormatter("#''").String(d12s))
}
