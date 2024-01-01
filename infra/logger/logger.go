package logger

import (
	"log/slog"
	"os"

	"github.com/kyong0612/fitness-saporter/infra/config"
)

func Init() {
	var slogHandler slog.Handler
	opt := slog.HandlerOptions{
		AddSource: true,
	}

	if config.Get().ENV == "local" {
		slogHandler = slog.NewTextHandler(os.Stdout, &opt)
	} else {
		opt.ReplaceAttr = func(groups []string, a slog.Attr) slog.Attr {
			switch a.Key {
			case slog.LevelKey:
				a = slog.Attr{
					Key:   "severity",
					Value: a.Value,
				}
			case slog.SourceKey:
				a = slog.Attr{
					Key:   "logging.googleapis.com/sourceLocation",
					Value: a.Value,
				}
			}
			return a
		}
		slogHandler = slog.NewJSONHandler(os.Stdout, &opt)
	}

	slog.SetDefault(slog.New(slogHandler))
}