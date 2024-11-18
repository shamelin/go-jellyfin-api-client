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
	JELLYFINGENERALCOMMANDTYPE_MOVE_UP JellyfinGeneralCommandType = "MoveUp"
	JELLYFINGENERALCOMMANDTYPE_MOVE_DOWN JellyfinGeneralCommandType = "MoveDown"
	JELLYFINGENERALCOMMANDTYPE_MOVE_LEFT JellyfinGeneralCommandType = "MoveLeft"
	JELLYFINGENERALCOMMANDTYPE_MOVE_RIGHT JellyfinGeneralCommandType = "MoveRight"
	JELLYFINGENERALCOMMANDTYPE_PAGE_UP JellyfinGeneralCommandType = "PageUp"
	JELLYFINGENERALCOMMANDTYPE_PAGE_DOWN JellyfinGeneralCommandType = "PageDown"
	JELLYFINGENERALCOMMANDTYPE_PREVIOUS_LETTER JellyfinGeneralCommandType = "PreviousLetter"
	JELLYFINGENERALCOMMANDTYPE_NEXT_LETTER JellyfinGeneralCommandType = "NextLetter"
	JELLYFINGENERALCOMMANDTYPE_TOGGLE_OSD JellyfinGeneralCommandType = "ToggleOsd"
	JELLYFINGENERALCOMMANDTYPE_TOGGLE_CONTEXT_MENU JellyfinGeneralCommandType = "ToggleContextMenu"
	JELLYFINGENERALCOMMANDTYPE_SELECT JellyfinGeneralCommandType = "Select"
	JELLYFINGENERALCOMMANDTYPE_BACK JellyfinGeneralCommandType = "Back"
	JELLYFINGENERALCOMMANDTYPE_TAKE_SCREENSHOT JellyfinGeneralCommandType = "TakeScreenshot"
	JELLYFINGENERALCOMMANDTYPE_SEND_KEY JellyfinGeneralCommandType = "SendKey"
	JELLYFINGENERALCOMMANDTYPE_SEND_STRING JellyfinGeneralCommandType = "SendString"
	JELLYFINGENERALCOMMANDTYPE_GO_HOME JellyfinGeneralCommandType = "GoHome"
	JELLYFINGENERALCOMMANDTYPE_GO_TO_SETTINGS JellyfinGeneralCommandType = "GoToSettings"
	JELLYFINGENERALCOMMANDTYPE_VOLUME_UP JellyfinGeneralCommandType = "VolumeUp"
	JELLYFINGENERALCOMMANDTYPE_VOLUME_DOWN JellyfinGeneralCommandType = "VolumeDown"
	JELLYFINGENERALCOMMANDTYPE_MUTE JellyfinGeneralCommandType = "Mute"
	JELLYFINGENERALCOMMANDTYPE_UNMUTE JellyfinGeneralCommandType = "Unmute"
	JELLYFINGENERALCOMMANDTYPE_TOGGLE_MUTE JellyfinGeneralCommandType = "ToggleMute"
	JELLYFINGENERALCOMMANDTYPE_SET_VOLUME JellyfinGeneralCommandType = "SetVolume"
	JELLYFINGENERALCOMMANDTYPE_SET_AUDIO_STREAM_INDEX JellyfinGeneralCommandType = "SetAudioStreamIndex"
	JELLYFINGENERALCOMMANDTYPE_SET_SUBTITLE_STREAM_INDEX JellyfinGeneralCommandType = "SetSubtitleStreamIndex"
	JELLYFINGENERALCOMMANDTYPE_TOGGLE_FULLSCREEN JellyfinGeneralCommandType = "ToggleFullscreen"
	JELLYFINGENERALCOMMANDTYPE_DISPLAY_CONTENT JellyfinGeneralCommandType = "DisplayContent"
	JELLYFINGENERALCOMMANDTYPE_GO_TO_SEARCH JellyfinGeneralCommandType = "GoToSearch"
	JELLYFINGENERALCOMMANDTYPE_DISPLAY_MESSAGE JellyfinGeneralCommandType = "DisplayMessage"
	JELLYFINGENERALCOMMANDTYPE_SET_REPEAT_MODE JellyfinGeneralCommandType = "SetRepeatMode"
	JELLYFINGENERALCOMMANDTYPE_CHANNEL_UP JellyfinGeneralCommandType = "ChannelUp"
	JELLYFINGENERALCOMMANDTYPE_CHANNEL_DOWN JellyfinGeneralCommandType = "ChannelDown"
	JELLYFINGENERALCOMMANDTYPE_GUIDE JellyfinGeneralCommandType = "Guide"
	JELLYFINGENERALCOMMANDTYPE_TOGGLE_STATS JellyfinGeneralCommandType = "ToggleStats"
	JELLYFINGENERALCOMMANDTYPE_PLAY_MEDIA_SOURCE JellyfinGeneralCommandType = "PlayMediaSource"
	JELLYFINGENERALCOMMANDTYPE_PLAY_TRAILERS JellyfinGeneralCommandType = "PlayTrailers"
	JELLYFINGENERALCOMMANDTYPE_SET_SHUFFLE_QUEUE JellyfinGeneralCommandType = "SetShuffleQueue"
	JELLYFINGENERALCOMMANDTYPE_PLAY_STATE JellyfinGeneralCommandType = "PlayState"
	JELLYFINGENERALCOMMANDTYPE_PLAY_NEXT JellyfinGeneralCommandType = "PlayNext"
	JELLYFINGENERALCOMMANDTYPE_TOGGLE_OSD_MENU JellyfinGeneralCommandType = "ToggleOsdMenu"
	JELLYFINGENERALCOMMANDTYPE_PLAY JellyfinGeneralCommandType = "Play"
	JELLYFINGENERALCOMMANDTYPE_SET_MAX_STREAMING_BITRATE JellyfinGeneralCommandType = "SetMaxStreamingBitrate"
	JELLYFINGENERALCOMMANDTYPE_SET_PLAYBACK_ORDER JellyfinGeneralCommandType = "SetPlaybackOrder"
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

