/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// ProfileConditionValue the model 'ProfileConditionValue'
type ProfileConditionValue string

// List of ProfileConditionValue
const (
	AUDIO_CHANNELS ProfileConditionValue = "AudioChannels"
	AUDIO_BITRATE ProfileConditionValue = "AudioBitrate"
	AUDIO_PROFILE ProfileConditionValue = "AudioProfile"
	WIDTH ProfileConditionValue = "Width"
	HEIGHT ProfileConditionValue = "Height"
	HAS64_BIT_OFFSETS ProfileConditionValue = "Has64BitOffsets"
	PACKET_LENGTH ProfileConditionValue = "PacketLength"
	VIDEO_BIT_DEPTH ProfileConditionValue = "VideoBitDepth"
	VIDEO_BITRATE ProfileConditionValue = "VideoBitrate"
	VIDEO_FRAMERATE ProfileConditionValue = "VideoFramerate"
	VIDEO_LEVEL ProfileConditionValue = "VideoLevel"
	VIDEO_PROFILE ProfileConditionValue = "VideoProfile"
	VIDEO_TIMESTAMP ProfileConditionValue = "VideoTimestamp"
	IS_ANAMORPHIC ProfileConditionValue = "IsAnamorphic"
	REF_FRAMES ProfileConditionValue = "RefFrames"
	NUM_AUDIO_STREAMS ProfileConditionValue = "NumAudioStreams"
	NUM_VIDEO_STREAMS ProfileConditionValue = "NumVideoStreams"
	IS_SECONDARY_AUDIO ProfileConditionValue = "IsSecondaryAudio"
	VIDEO_CODEC_TAG ProfileConditionValue = "VideoCodecTag"
	IS_AVC ProfileConditionValue = "IsAvc"
	IS_INTERLACED ProfileConditionValue = "IsInterlaced"
	AUDIO_SAMPLE_RATE ProfileConditionValue = "AudioSampleRate"
	AUDIO_BIT_DEPTH ProfileConditionValue = "AudioBitDepth"
	VIDEO_RANGE_TYPE ProfileConditionValue = "VideoRangeType"
)

// All allowed values of ProfileConditionValue enum
var AllowedProfileConditionValueEnumValues = []ProfileConditionValue{
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

func (v *ProfileConditionValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProfileConditionValue(value)
	for _, existing := range AllowedProfileConditionValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProfileConditionValue", value)
}

// NewProfileConditionValueFromValue returns a pointer to a valid ProfileConditionValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProfileConditionValueFromValue(v string) (*ProfileConditionValue, error) {
	ev := ProfileConditionValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProfileConditionValue: valid values are %v", v, AllowedProfileConditionValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProfileConditionValue) IsValid() bool {
	for _, existing := range AllowedProfileConditionValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProfileConditionValue value
func (v ProfileConditionValue) Ptr() *ProfileConditionValue {
	return &v
}

type NullableProfileConditionValue struct {
	value *ProfileConditionValue
	isSet bool
}

func (v NullableProfileConditionValue) Get() *ProfileConditionValue {
	return v.value
}

func (v *NullableProfileConditionValue) Set(val *ProfileConditionValue) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileConditionValue) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileConditionValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileConditionValue(val *ProfileConditionValue) *NullableProfileConditionValue {
	return &NullableProfileConditionValue{value: val, isSet: true}
}

func (v NullableProfileConditionValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileConditionValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

