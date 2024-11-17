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

// PersonKind The person kind.
type PersonKind string

// List of PersonKind
const (
	UNKNOWN PersonKind = "Unknown"
	ACTOR PersonKind = "Actor"
	DIRECTOR PersonKind = "Director"
	COMPOSER PersonKind = "Composer"
	WRITER PersonKind = "Writer"
	GUEST_STAR PersonKind = "GuestStar"
	PRODUCER PersonKind = "Producer"
	CONDUCTOR PersonKind = "Conductor"
	LYRICIST PersonKind = "Lyricist"
	ARRANGER PersonKind = "Arranger"
	ENGINEER PersonKind = "Engineer"
	MIXER PersonKind = "Mixer"
	REMIXER PersonKind = "Remixer"
	CREATOR PersonKind = "Creator"
	ARTIST PersonKind = "Artist"
	ALBUM_ARTIST PersonKind = "AlbumArtist"
	AUTHOR PersonKind = "Author"
	ILLUSTRATOR PersonKind = "Illustrator"
	PENCILLER PersonKind = "Penciller"
	INKER PersonKind = "Inker"
	COLORIST PersonKind = "Colorist"
	LETTERER PersonKind = "Letterer"
	COVER_ARTIST PersonKind = "CoverArtist"
	EDITOR PersonKind = "Editor"
	TRANSLATOR PersonKind = "Translator"
)

// All allowed values of PersonKind enum
var AllowedPersonKindEnumValues = []PersonKind{
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

func (v *PersonKind) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PersonKind(value)
	for _, existing := range AllowedPersonKindEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PersonKind", value)
}

// NewPersonKindFromValue returns a pointer to a valid PersonKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPersonKindFromValue(v string) (*PersonKind, error) {
	ev := PersonKind(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PersonKind: valid values are %v", v, AllowedPersonKindEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PersonKind) IsValid() bool {
	for _, existing := range AllowedPersonKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PersonKind value
func (v PersonKind) Ptr() *PersonKind {
	return &v
}

type NullablePersonKind struct {
	value *PersonKind
	isSet bool
}

func (v NullablePersonKind) Get() *PersonKind {
	return v.value
}

func (v *NullablePersonKind) Set(val *PersonKind) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonKind) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonKind) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonKind(val *PersonKind) *NullablePersonKind {
	return &NullablePersonKind{value: val, isSet: true}
}

func (v NullablePersonKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonKind) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

