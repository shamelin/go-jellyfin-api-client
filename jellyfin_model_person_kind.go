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

// JellyfinPersonKind The person kind.
type JellyfinPersonKind string

// List of PersonKind
const (
	UNKNOWN JellyfinPersonKind = "Unknown"
	ACTOR JellyfinPersonKind = "Actor"
	DIRECTOR JellyfinPersonKind = "Director"
	COMPOSER JellyfinPersonKind = "Composer"
	WRITER JellyfinPersonKind = "Writer"
	GUEST_STAR JellyfinPersonKind = "GuestStar"
	PRODUCER JellyfinPersonKind = "Producer"
	CONDUCTOR JellyfinPersonKind = "Conductor"
	LYRICIST JellyfinPersonKind = "Lyricist"
	ARRANGER JellyfinPersonKind = "Arranger"
	ENGINEER JellyfinPersonKind = "Engineer"
	MIXER JellyfinPersonKind = "Mixer"
	REMIXER JellyfinPersonKind = "Remixer"
	CREATOR JellyfinPersonKind = "Creator"
	ARTIST JellyfinPersonKind = "Artist"
	ALBUM_ARTIST JellyfinPersonKind = "AlbumArtist"
	AUTHOR JellyfinPersonKind = "Author"
	ILLUSTRATOR JellyfinPersonKind = "Illustrator"
	PENCILLER JellyfinPersonKind = "Penciller"
	INKER JellyfinPersonKind = "Inker"
	COLORIST JellyfinPersonKind = "Colorist"
	LETTERER JellyfinPersonKind = "Letterer"
	COVER_ARTIST JellyfinPersonKind = "CoverArtist"
	EDITOR JellyfinPersonKind = "Editor"
	TRANSLATOR JellyfinPersonKind = "Translator"
)

// All allowed values of JellyfinPersonKind enum
var AllowedJellyfinPersonKindEnumValues = []JellyfinPersonKind{
	"Unknown",
	"Actor",
	"Director",
	"Composer",
	"Writer",
	"GuestStar",
	"Producer",
	"Conductor",
	"Lyricist",
	"Arranger",
	"Engineer",
	"Mixer",
	"Remixer",
	"Creator",
	"Artist",
	"AlbumArtist",
	"Author",
	"Illustrator",
	"Penciller",
	"Inker",
	"Colorist",
	"Letterer",
	"CoverArtist",
	"Editor",
	"Translator",
}

func (v *JellyfinPersonKind) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JellyfinPersonKind(value)
	for _, existing := range AllowedJellyfinPersonKindEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JellyfinPersonKind", value)
}

// NewJellyfinPersonKindFromValue returns a pointer to a valid JellyfinPersonKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJellyfinPersonKindFromValue(v string) (*JellyfinPersonKind, error) {
	ev := JellyfinPersonKind(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JellyfinPersonKind: valid values are %v", v, AllowedJellyfinPersonKindEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JellyfinPersonKind) IsValid() bool {
	for _, existing := range AllowedJellyfinPersonKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PersonKind value
func (v JellyfinPersonKind) Ptr() *JellyfinPersonKind {
	return &v
}

type NullableJellyfinPersonKind struct {
	value *JellyfinPersonKind
	isSet bool
}

func (v NullableJellyfinPersonKind) Get() *JellyfinPersonKind {
	return v.value
}

func (v *NullableJellyfinPersonKind) Set(val *JellyfinPersonKind) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinPersonKind) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinPersonKind) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinPersonKind(val *JellyfinPersonKind) *NullableJellyfinPersonKind {
	return &NullableJellyfinPersonKind{value: val, isSet: true}
}

func (v NullableJellyfinPersonKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinPersonKind) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

