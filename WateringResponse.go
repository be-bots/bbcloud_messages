package messages

import (
	"bbcloud_messages/part"
	"time"
)

// WateringData holds watering data
type WateringData struct {
	PreWateringWeight  int32
	PostWateringWeight int32
	TypeOfNossle       string
	FanStrength        int32
}

// WateringResponse is the response of a watering action
type WateringResponse struct {
	Meta          part.Meta
	ErrorStatus   part.ErrorStatus
	MessageID     string
	Type          string
	InstructionID string
	Timestamp     time.Time
	Source        part.Source
	WateringData  WateringData
	Datapoints    []part.Datapoint
}
