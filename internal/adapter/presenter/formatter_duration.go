package presenter

import (
	"fmt"
	"log/slog"
	"strings"
	"time"
)

type DurationFormatter struct {
	format string
	printf string
}

func NewDurationFormatter(format string) *DurationFormatter {
	printf := strings.ReplaceAll(strings.Replace(format, "#", "%d", 1), "#", "%02d")
	slog.Debug("duration formatter", "frmat", printf)
	return &DurationFormatter{
		format: format,
		printf: printf,
	}
}

func (f *DurationFormatter) String(d time.Duration) string {
	switch f.format {
	case "#h#m#s", "#:#:#":
		hour := uint(d.Hours())
		min := uint(d.Minutes()) % 60
		sec := uint(d.Seconds()) % 60
		return fmt.Sprintf(f.printf, hour, min, sec)
	case "#m#s", "#:#", "#'#''":
		min := uint(d.Minutes())
		sec := uint(d.Seconds()) % 60
		return fmt.Sprintf(f.printf, min, sec)
	case "#s", "#''":
		sec := uint(d.Seconds())
		return fmt.Sprintf(f.printf, sec)
	default:
		return d.String()
	}
}
