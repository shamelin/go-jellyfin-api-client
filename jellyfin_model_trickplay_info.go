/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package jellyfin

import (
	"encoding/json"
)

// checks if the JellyfinTrickplayInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JellyfinTrickplayInfo{}

// JellyfinTrickplayInfo An entity representing the metadata for a group of trickplay tiles.
type JellyfinTrickplayInfo struct {
	// Gets or sets width of an individual thumbnail.
	Width *int32 `json:"Width,omitempty"`
	// Gets or sets height of an individual thumbnail.
	Height *int32 `json:"Height,omitempty"`
	// Gets or sets amount of thumbnails per row.
	TileWidth *int32 `json:"TileWidth,omitempty"`
	// Gets or sets amount of thumbnails per column.
	TileHeight *int32 `json:"TileHeight,omitempty"`
	// Gets or sets total amount of non-black thumbnails.
	ThumbnailCount *int32 `json:"ThumbnailCount,omitempty"`
	// Gets or sets interval in milliseconds between each trickplay thumbnail.
	Interval *int32 `json:"Interval,omitempty"`
	// Gets or sets peak bandwith usage in bits per second.
	Bandwidth *int32 `json:"Bandwidth,omitempty"`
}

// NewJellyfinTrickplayInfo instantiates a new JellyfinTrickplayInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJellyfinTrickplayInfo() *JellyfinTrickplayInfo {
	this := JellyfinTrickplayInfo{}
	return &this
}

// NewJellyfinTrickplayInfoWithDefaults instantiates a new JellyfinTrickplayInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJellyfinTrickplayInfoWithDefaults() *JellyfinTrickplayInfo {
	this := JellyfinTrickplayInfo{}
	return &this
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *JellyfinTrickplayInfo) GetWidth() int32 {
	if o == nil || IsNil(o.Width) {
		var ret int32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinTrickplayInfo) GetWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *JellyfinTrickplayInfo) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given int32 and assigns it to the Width field.
func (o *JellyfinTrickplayInfo) SetWidth(v int32) {
	o.Width = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *JellyfinTrickplayInfo) GetHeight() int32 {
	if o == nil || IsNil(o.Height) {
		var ret int32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinTrickplayInfo) GetHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *JellyfinTrickplayInfo) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
func (o *JellyfinTrickplayInfo) SetHeight(v int32) {
	o.Height = &v
}

// GetTileWidth returns the TileWidth field value if set, zero value otherwise.
func (o *JellyfinTrickplayInfo) GetTileWidth() int32 {
	if o == nil || IsNil(o.TileWidth) {
		var ret int32
		return ret
	}
	return *o.TileWidth
}

// GetTileWidthOk returns a tuple with the TileWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinTrickplayInfo) GetTileWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.TileWidth) {
		return nil, false
	}
	return o.TileWidth, true
}

// HasTileWidth returns a boolean if a field has been set.
func (o *JellyfinTrickplayInfo) HasTileWidth() bool {
	if o != nil && !IsNil(o.TileWidth) {
		return true
	}

	return false
}

// SetTileWidth gets a reference to the given int32 and assigns it to the TileWidth field.
func (o *JellyfinTrickplayInfo) SetTileWidth(v int32) {
	o.TileWidth = &v
}

// GetTileHeight returns the TileHeight field value if set, zero value otherwise.
func (o *JellyfinTrickplayInfo) GetTileHeight() int32 {
	if o == nil || IsNil(o.TileHeight) {
		var ret int32
		return ret
	}
	return *o.TileHeight
}

// GetTileHeightOk returns a tuple with the TileHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinTrickplayInfo) GetTileHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.TileHeight) {
		return nil, false
	}
	return o.TileHeight, true
}

// HasTileHeight returns a boolean if a field has been set.
func (o *JellyfinTrickplayInfo) HasTileHeight() bool {
	if o != nil && !IsNil(o.TileHeight) {
		return true
	}

	return false
}

// SetTileHeight gets a reference to the given int32 and assigns it to the TileHeight field.
func (o *JellyfinTrickplayInfo) SetTileHeight(v int32) {
	o.TileHeight = &v
}

// GetThumbnailCount returns the ThumbnailCount field value if set, zero value otherwise.
func (o *JellyfinTrickplayInfo) GetThumbnailCount() int32 {
	if o == nil || IsNil(o.ThumbnailCount) {
		var ret int32
		return ret
	}
	return *o.ThumbnailCount
}

// GetThumbnailCountOk returns a tuple with the ThumbnailCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinTrickplayInfo) GetThumbnailCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ThumbnailCount) {
		return nil, false
	}
	return o.ThumbnailCount, true
}

// HasThumbnailCount returns a boolean if a field has been set.
func (o *JellyfinTrickplayInfo) HasThumbnailCount() bool {
	if o != nil && !IsNil(o.ThumbnailCount) {
		return true
	}

	return false
}

// SetThumbnailCount gets a reference to the given int32 and assigns it to the ThumbnailCount field.
func (o *JellyfinTrickplayInfo) SetThumbnailCount(v int32) {
	o.ThumbnailCount = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *JellyfinTrickplayInfo) GetInterval() int32 {
	if o == nil || IsNil(o.Interval) {
		var ret int32
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinTrickplayInfo) GetIntervalOk() (*int32, bool) {
	if o == nil || IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *JellyfinTrickplayInfo) HasInterval() bool {
	if o != nil && !IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given int32 and assigns it to the Interval field.
func (o *JellyfinTrickplayInfo) SetInterval(v int32) {
	o.Interval = &v
}

// GetBandwidth returns the Bandwidth field value if set, zero value otherwise.
func (o *JellyfinTrickplayInfo) GetBandwidth() int32 {
	if o == nil || IsNil(o.Bandwidth) {
		var ret int32
		return ret
	}
	return *o.Bandwidth
}

// GetBandwidthOk returns a tuple with the Bandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinTrickplayInfo) GetBandwidthOk() (*int32, bool) {
	if o == nil || IsNil(o.Bandwidth) {
		return nil, false
	}
	return o.Bandwidth, true
}

// HasBandwidth returns a boolean if a field has been set.
func (o *JellyfinTrickplayInfo) HasBandwidth() bool {
	if o != nil && !IsNil(o.Bandwidth) {
		return true
	}

	return false
}

// SetBandwidth gets a reference to the given int32 and assigns it to the Bandwidth field.
func (o *JellyfinTrickplayInfo) SetBandwidth(v int32) {
	o.Bandwidth = &v
}

func (o JellyfinTrickplayInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JellyfinTrickplayInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Width) {
		toSerialize["Width"] = o.Width
	}
	if !IsNil(o.Height) {
		toSerialize["Height"] = o.Height
	}
	if !IsNil(o.TileWidth) {
		toSerialize["TileWidth"] = o.TileWidth
	}
	if !IsNil(o.TileHeight) {
		toSerialize["TileHeight"] = o.TileHeight
	}
	if !IsNil(o.ThumbnailCount) {
		toSerialize["ThumbnailCount"] = o.ThumbnailCount
	}
	if !IsNil(o.Interval) {
		toSerialize["Interval"] = o.Interval
	}
	if !IsNil(o.Bandwidth) {
		toSerialize["Bandwidth"] = o.Bandwidth
	}
	return toSerialize, nil
}

type NullableJellyfinTrickplayInfo struct {
	value *JellyfinTrickplayInfo
	isSet bool
}

func (v NullableJellyfinTrickplayInfo) Get() *JellyfinTrickplayInfo {
	return v.value
}

func (v *NullableJellyfinTrickplayInfo) Set(val *JellyfinTrickplayInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinTrickplayInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinTrickplayInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinTrickplayInfo(val *JellyfinTrickplayInfo) *NullableJellyfinTrickplayInfo {
	return &NullableJellyfinTrickplayInfo{value: val, isSet: true}
}

func (v NullableJellyfinTrickplayInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinTrickplayInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


