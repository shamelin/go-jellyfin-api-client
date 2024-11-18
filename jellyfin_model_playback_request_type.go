/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package jellyfin

import (
	"encoding/json"
	"fmt"
)

// JellyfinPlaybackRequestType Enum PlaybackRequestType.
type JellyfinPlaybackRequestType string

// List of PlaybackRequestType
const (
	PLAY JellyfinPlaybackRequestType = "Play"
	SET_PLAYLIST_ITEM JellyfinPlaybackRequestType = "SetPlaylistItem"
	REMOVE_FROM_PLAYLIST JellyfinPlaybackRequestType = "RemoveFromPlaylist"
	MOVE_PLAYLIST_ITEM JellyfinPlaybackRequestType = "MovePlaylistItem"
	QUEUE JellyfinPlaybackRequestType = "Queue"
	UNPAUSE JellyfinPlaybackRequestType = "Unpause"
	PAUSE JellyfinPlaybackRequestType = "Pause"
	STOP JellyfinPlaybackRequestType = "Stop"
	SEEK JellyfinPlaybackRequestType = "Seek"
	BUFFER JellyfinPlaybackRequestType = "Buffer"
	READY JellyfinPlaybackRequestType = "Ready"
	NEXT_ITEM JellyfinPlaybackRequestType = "NextItem"
	PREVIOUS_ITEM JellyfinPlaybackRequestType = "PreviousItem"
	SET_REPEAT_MODE JellyfinPlaybackRequestType = "SetRepeatMode"
	SET_SHUFFLE_MODE JellyfinPlaybackRequestType = "SetShuffleMode"
	PING JellyfinPlaybackRequestType = "Ping"
	IGNORE_WAIT JellyfinPlaybackRequestType = "IgnoreWait"
)

// All allowed values of JellyfinPlaybackRequestType enum
var AllowedJellyfinPlaybackRequestTypeEnumValues = []JellyfinPlaybackRequestType{
	"Play",
	"SetPlaylistItem",
	"RemoveFromPlaylist",
	"MovePlaylistItem",
	"Queue",
	"Unpause",
	"Pause",
	"Stop",
	"Seek",
	"Buffer",
	"Ready",
	"NextItem",
	"PreviousItem",
	"SetRepeatMode",
	"SetShuffleMode",
	"Ping",
	"IgnoreWait",
}

func (v *JellyfinPlaybackRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JellyfinPlaybackRequestType(value)
	for _, existing := range AllowedJellyfinPlaybackRequestTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JellyfinPlaybackRequestType", value)
}

// NewJellyfinPlaybackRequestTypeFromValue returns a pointer to a valid JellyfinPlaybackRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJellyfinPlaybackRequestTypeFromValue(v string) (*JellyfinPlaybackRequestType, error) {
	ev := JellyfinPlaybackRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JellyfinPlaybackRequestType: valid values are %v", v, AllowedJellyfinPlaybackRequestTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JellyfinPlaybackRequestType) IsValid() bool {
	for _, existing := range AllowedJellyfinPlaybackRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PlaybackRequestType value
func (v JellyfinPlaybackRequestType) Ptr() *JellyfinPlaybackRequestType {
	return &v
}

type NullableJellyfinPlaybackRequestType struct {
	value *JellyfinPlaybackRequestType
	isSet bool
}

func (v NullableJellyfinPlaybackRequestType) Get() *JellyfinPlaybackRequestType {
	return v.value
}

func (v *NullableJellyfinPlaybackRequestType) Set(val *JellyfinPlaybackRequestType) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinPlaybackRequestType) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinPlaybackRequestType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinPlaybackRequestType(val *JellyfinPlaybackRequestType) *NullableJellyfinPlaybackRequestType {
	return &NullableJellyfinPlaybackRequestType{value: val, isSet: true}
}

func (v NullableJellyfinPlaybackRequestType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinPlaybackRequestType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

