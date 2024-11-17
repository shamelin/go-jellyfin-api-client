/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package jellyfin-api-client

import (
	"encoding/json"
	"fmt"
)

// PlaystateCommand Enum PlaystateCommand.
type PlaystateCommand string

// List of PlaystateCommand
const (
	STOP PlaystateCommand = "Stop"
	PAUSE PlaystateCommand = "Pause"
	UNPAUSE PlaystateCommand = "Unpause"
	NEXT_TRACK PlaystateCommand = "NextTrack"
	PREVIOUS_TRACK PlaystateCommand = "PreviousTrack"
	SEEK PlaystateCommand = "Seek"
	REWIND PlaystateCommand = "Rewind"
	FAST_FORWARD PlaystateCommand = "FastForward"
	PLAY_PAUSE PlaystateCommand = "PlayPause"
)

// All allowed values of PlaystateCommand enum
var AllowedPlaystateCommandEnumValues = []PlaystateCommand{
	"Stop",
	"Pause",
	"Unpause",
	"NextTrack",
	"PreviousTrack",
	"Seek",
	"Rewind",
	"FastForward",
	"PlayPause",
}

func (v *PlaystateCommand) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PlaystateCommand(value)
	for _, existing := range AllowedPlaystateCommandEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PlaystateCommand", value)
}

// NewPlaystateCommandFromValue returns a pointer to a valid PlaystateCommand
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPlaystateCommandFromValue(v string) (*PlaystateCommand, error) {
	ev := PlaystateCommand(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PlaystateCommand: valid values are %v", v, AllowedPlaystateCommandEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PlaystateCommand) IsValid() bool {
	for _, existing := range AllowedPlaystateCommandEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PlaystateCommand value
func (v PlaystateCommand) Ptr() *PlaystateCommand {
	return &v
}

type NullablePlaystateCommand struct {
	value *PlaystateCommand
	isSet bool
}

func (v NullablePlaystateCommand) Get() *PlaystateCommand {
	return v.value
}

func (v *NullablePlaystateCommand) Set(val *PlaystateCommand) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaystateCommand) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaystateCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaystateCommand(val *PlaystateCommand) *NullablePlaystateCommand {
	return &NullablePlaystateCommand{value: val, isSet: true}
}

func (v NullablePlaystateCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaystateCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

