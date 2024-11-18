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

// JellyfinProfileConditionValue the model 'JellyfinProfileConditionValue'
type JellyfinProfileConditionValue string

// List of ProfileConditionValue
const (
	AUDIO_CHANNELS JellyfinProfileConditionValue = "AudioChannels"
	AUDIO_BITRATE JellyfinProfileConditionValue = "AudioBitrate"
	AUDIO_PROFILE JellyfinProfileConditionValue = "AudioProfile"
	WIDTH JellyfinProfileConditionValue = "Width"
	HEIGHT JellyfinProfileConditionValue = "Height"
	HAS64_BIT_OFFSETS JellyfinProfileConditionValue = "Has64BitOffsets"
	PACKET_LENGTH JellyfinProfileConditionValue = "PacketLength"
	VIDEO_BIT_DEPTH JellyfinProfileConditionValue = "VideoBitDepth"
	VIDEO_BITRATE JellyfinProfileConditionValue = "VideoBitrate"
	VIDEO_FRAMERATE JellyfinProfileConditionValue = "VideoFramerate"
	VIDEO_LEVEL JellyfinProfileConditionValue = "VideoLevel"
	VIDEO_PROFILE JellyfinProfileConditionValue = "VideoProfile"
	VIDEO_TIMESTAMP JellyfinProfileConditionValue = "VideoTimestamp"
	IS_ANAMORPHIC JellyfinProfileConditionValue = "IsAnamorphic"
	REF_FRAMES JellyfinProfileConditionValue = "RefFrames"
	NUM_AUDIO_STREAMS JellyfinProfileConditionValue = "NumAudioStreams"
	NUM_VIDEO_STREAMS JellyfinProfileConditionValue = "NumVideoStreams"
	IS_SECONDARY_AUDIO JellyfinProfileConditionValue = "IsSecondaryAudio"
	VIDEO_CODEC_TAG JellyfinProfileConditionValue = "VideoCodecTag"
	IS_AVC JellyfinProfileConditionValue = "IsAvc"
	IS_INTERLACED JellyfinProfileConditionValue = "IsInterlaced"
	AUDIO_SAMPLE_RATE JellyfinProfileConditionValue = "AudioSampleRate"
	AUDIO_BIT_DEPTH JellyfinProfileConditionValue = "AudioBitDepth"
	VIDEO_RANGE_TYPE JellyfinProfileConditionValue = "VideoRangeType"
)

// All allowed values of JellyfinProfileConditionValue enum
var AllowedJellyfinProfileConditionValueEnumValues = []JellyfinProfileConditionValue{
	"AudioChannels",
	"AudioBitrate",
	"AudioProfile",
	"Width",
	"Height",
	"Has64BitOffsets",
	"PacketLength",
	"VideoBitDepth",
	"VideoBitrate",
	"VideoFramerate",
	"VideoLevel",
	"VideoProfile",
	"VideoTimestamp",
	"IsAnamorphic",
	"RefFrames",
	"NumAudioStreams",
	"NumVideoStreams",
	"IsSecondaryAudio",
	"VideoCodecTag",
	"IsAvc",
	"IsInterlaced",
	"AudioSampleRate",
	"AudioBitDepth",
	"VideoRangeType",
}

func (v *JellyfinProfileConditionValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JellyfinProfileConditionValue(value)
	for _, existing := range AllowedJellyfinProfileConditionValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JellyfinProfileConditionValue", value)
}

// NewJellyfinProfileConditionValueFromValue returns a pointer to a valid JellyfinProfileConditionValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJellyfinProfileConditionValueFromValue(v string) (*JellyfinProfileConditionValue, error) {
	ev := JellyfinProfileConditionValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JellyfinProfileConditionValue: valid values are %v", v, AllowedJellyfinProfileConditionValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JellyfinProfileConditionValue) IsValid() bool {
	for _, existing := range AllowedJellyfinProfileConditionValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProfileConditionValue value
func (v JellyfinProfileConditionValue) Ptr() *JellyfinProfileConditionValue {
	return &v
}

type NullableJellyfinProfileConditionValue struct {
	value *JellyfinProfileConditionValue
	isSet bool
}

func (v NullableJellyfinProfileConditionValue) Get() *JellyfinProfileConditionValue {
	return v.value
}

func (v *NullableJellyfinProfileConditionValue) Set(val *JellyfinProfileConditionValue) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinProfileConditionValue) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinProfileConditionValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinProfileConditionValue(val *JellyfinProfileConditionValue) *NullableJellyfinProfileConditionValue {
	return &NullableJellyfinProfileConditionValue{value: val, isSet: true}
}

func (v NullableJellyfinProfileConditionValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinProfileConditionValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
