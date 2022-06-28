package wirebird

import (
	"github.com/harnyk/go-wirebird/internal/models"
	"github.com/harnyk/go-wirebird/internal/models/compat"
)

type Wirebird interface {
	BroadcastEvent(*models.LoggerEvent) error
	BroadcastLegacyEvent(*compat.SerializedLoggerEvent) error
}
