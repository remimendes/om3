package resdisk

import (
	"embed"

	"github.com/opensvc/om3/core/keywords"
	"github.com/opensvc/om3/core/manifest"
	"github.com/opensvc/om3/core/resource"
	"github.com/opensvc/om3/util/converters"
)

type (
	T struct {
		resource.T
		resource.SCSIPersistentReservation
		PromoteRW bool
	}
)

var (
	//go:embed text
	fs embed.FS

	KWPromoteRW = keywords.Keyword{
		Option:    "promote_rw",
		Attr:      "PromoteRW",
		Converter: converters.Bool,
		Text:      keywords.NewText(fs, "text/kw/promote_rw"),
	}

	BaseKeywords = append(
		manifest.SCSIPersistentReservationKeywords,
		KWPromoteRW,
	)
)
