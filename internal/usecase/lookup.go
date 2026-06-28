package usecase

import (
	"context"
	"log/slog"
	"trec/internal/core/logger"
	"trec/internal/domain"
)

type lookupOptions interface {
	Style() domain.LookupStyle
	Filter() domain.Filter
	Order() []domain.Order
}

type Lookup struct {
	simple    *LookupSimple
	full      *LookupFull
	collapsed *LookupCollapsed
}

func NewLookup(simple *LookupSimple, full *LookupFull, collapsed *LookupCollapsed) *Lookup {
	return &Lookup{
		simple:    simple,
		full:      full,
		collapsed: collapsed,
	}
}

func (uc *Lookup) Lookup(ctx context.Context, opts lookupOptions) error {
	slog.Debug("Execute lookup", "opts", opts)

	switch opts.Style() {
	case domain.LookupSimple:
		err := uc.simple.Lookup(ctx, opts)
		if err != nil {
			return logger.Error("Lookup", "simple lookup error", err)
		}
	case domain.LookupFull:
		err := uc.full.Lookup(ctx, opts)
		if err != nil {
			return logger.Error("Lookup", "full lookup error", err)
		}
	case domain.LookupCollapse:
		err := uc.collapsed.Lookup(ctx, opts)
		if err != nil {
			return logger.Error("Lookup", "collapsed lookup error", err)
		}
	default:
		return logger.Error("Lookup", "lookup format error", domain.ErrorUnknownFormat, "style", opts.Style())
	}

	slog.Debug("Finished lookup")
	return nil
}
