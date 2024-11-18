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

// JellyfinGeneralCommandType This exists simply to identify a set of known commands.
type JellyfinGeneralCommandType string

// List of GeneralCommandType
const (
	MOVE_UP JellyfinGeneralCommandType = "MoveUp"
	MOVE_DOWN JellyfinGeneralCommandType = "MoveDown"
	MOVE_LEFT JellyfinGeneralCommandType = "MoveLeft"
	MOVE_RIGHT JellyfinGeneralCommandType = "MoveRight"
	PAGE_UP JellyfinGeneralCommandType = "PageUp"
	PAGE_DOWN JellyfinGeneralCommandType = "PageDown"
	PREVIOUS_LETTER JellyfinGeneralCommandType = "PreviousLetter"
	NEXT_LETTER JellyfinGeneralCommandType = "NextLetter"
	TOGGLE_OSD JellyfinGeneralCommandType = "ToggleOsd"
	TOGGLE_CONTEXT_MENU JellyfinGeneralCommandType = "ToggleContextMenu"
	SELECT JellyfinGeneralCommandType = "Select"
	BACK JellyfinGeneralCommandType = "Back"
	TAKE_SCREENSHOT JellyfinGeneralCommandType = "TakeScreenshot"
	SEND_KEY JellyfinGeneralCommandType = "SendKey"
	SEND_STRING JellyfinGeneralCommandType = "SendString"
	GO_HOME JellyfinGeneralCommandType = "GoHome"
	GO_TO_SETTINGS JellyfinGeneralCommandType = "GoToSettings"
	VOLUME_UP JellyfinGeneralCommandType = "VolumeUp"
	VOLUME_DOWN JellyfinGeneralCommandType = "VolumeDown"
	MUTE JellyfinGeneralCommandType = "Mute"
	UNMUTE JellyfinGeneralCommandType = "Unmute"
	TOGGLE_MUTE JellyfinGeneralCommandType = "ToggleMute"
	SET_VOLUME JellyfinGeneralCommandType = "SetVolume"
	SET_AUDIO_STREAM_INDEX JellyfinGeneralCommandType = "SetAudioStreamIndex"
	SET_SUBTITLE_STREAM_INDEX JellyfinGeneralCommandType = "SetSubtitleStreamIndex"
	TOGGLE_FULLSCREEN JellyfinGeneralCommandType = "ToggleFullscreen"
	DISPLAY_CONTENT JellyfinGeneralCommandType = "DisplayContent"
	GO_TO_SEARCH JellyfinGeneralCommandType = "GoToSearch"
	DISPLAY_MESSAGE JellyfinGeneralCommandType = "DisplayMessage"
	SET_REPEAT_MODE JellyfinGeneralCommandType = "SetRepeatMode"
	CHANNEL_UP JellyfinGeneralCommandType = "ChannelUp"
	CHANNEL_DOWN JellyfinGeneralCommandType = "ChannelDown"
	GUIDE JellyfinGeneralCommandType = "Guide"
	TOGGLE_STATS JellyfinGeneralCommandType = "ToggleStats"
	PLAY_MEDIA_SOURCE JellyfinGeneralCommandType = "PlayMediaSource"
	PLAY_TRAILERS JellyfinGeneralCommandType = "PlayTrailers"
	SET_SHUFFLE_QUEUE JellyfinGeneralCommandType = "SetShuffleQueue"
	PLAY_STATE JellyfinGeneralCommandType = "PlayState"
	PLAY_NEXT JellyfinGeneralCommandType = "PlayNext"
	TOGGLE_OSD_MENU JellyfinGeneralCommandType = "ToggleOsdMenu"
	PLAY JellyfinGeneralCommandType = "Play"
	SET_MAX_STREAMING_BITRATE JellyfinGeneralCommandType = "SetMaxStreamingBitrate"
	SET_PLAYBACK_ORDER JellyfinGeneralCommandType = "SetPlaybackOrder"
)

// All allowed values of JellyfinGeneralCommandType enum
var AllowedJellyfinGeneralCommandTypeEnumValues = []JellyfinGeneralCommandType{
	"MoveUp",
	"MoveDown",
	"MoveLeft",
	"MoveRight",
	"PageUp",
	"PageDown",
	"PreviousLetter",
	"NextLetter",
	"ToggleOsd",
	"ToggleContextMenu",
	"Select",
	"Back",
	"TakeScreenshot",
	"SendKey",
	"SendString",
	"GoHome",
	"GoToSettings",
	"VolumeUp",
	"VolumeDown",
	"Mute",
	"Unmute",
	"ToggleMute",
	"SetVolume",
	"SetAudioStreamIndex",
	"SetSubtitleStreamIndex",
	"ToggleFullscreen",
	"DisplayContent",
	"GoToSearch",
	"DisplayMessage",
	"SetRepeatMode",
	"ChannelUp",
	"ChannelDown",
	"Guide",
	"ToggleStats",
	"PlayMediaSource",
	"PlayTrailers",
	"SetShuffleQueue",
	"PlayState",
	"PlayNext",
	"ToggleOsdMenu",
	"Play",
	"SetMaxStreamingBitrate",
	"SetPlaybackOrder",
}

func (v *JellyfinGeneralCommandType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JellyfinGeneralCommandType(value)
	for _, existing := range AllowedJellyfinGeneralCommandTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JellyfinGeneralCommandType", value)
}

// NewJellyfinGeneralCommandTypeFromValue returns a pointer to a valid JellyfinGeneralCommandType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJellyfinGeneralCommandTypeFromValue(v string) (*JellyfinGeneralCommandType, error) {
	ev := JellyfinGeneralCommandType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JellyfinGeneralCommandType: valid values are %v", v, AllowedJellyfinGeneralCommandTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JellyfinGeneralCommandType) IsValid() bool {
	for _, existing := range AllowedJellyfinGeneralCommandTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GeneralCommandType value
func (v JellyfinGeneralCommandType) Ptr() *JellyfinGeneralCommandType {
	return &v
}

type NullableJellyfinGeneralCommandType struct {
	value *JellyfinGeneralCommandType
	isSet bool
}

func (v NullableJellyfinGeneralCommandType) Get() *JellyfinGeneralCommandType {
	return v.value
}

func (v *NullableJellyfinGeneralCommandType) Set(val *JellyfinGeneralCommandType) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinGeneralCommandType) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinGeneralCommandType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinGeneralCommandType(val *JellyfinGeneralCommandType) *NullableJellyfinGeneralCommandType {
	return &NullableJellyfinGeneralCommandType{value: val, isSet: true}
}

func (v NullableJellyfinGeneralCommandType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinGeneralCommandType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

