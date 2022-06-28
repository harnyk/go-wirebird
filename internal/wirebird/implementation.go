package wirebird

import (
	"encoding/json"

	"github.com/harnyk/go-wirebird/internal/models"
	"github.com/harnyk/go-wirebird/internal/models/compat"
	"gopkg.in/olahol/melody.v1"
)

type wirebird struct {
	melody *melody.Melody
}

func New(melody *melody.Melody) Wirebird {
	return &wirebird{melody: melody}
}

func (w *wirebird) BroadcastEvent(event *models.LoggerEvent) error {
	broadcastEvent, err := json.Marshal(event)
	if err != nil {
		return err
	}
	return w.melody.Broadcast(broadcastEvent)
}

func (w *wirebird) BroadcastLegacyEvent(event *compat.SerializedLoggerEvent) error {
	newEvent, err := models.Upgrade(event)
	if err != nil {
		return err
	}
	return w.BroadcastEvent(newEvent)
}
