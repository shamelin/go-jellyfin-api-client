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

// checks if the JellyfinMediaAttachment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JellyfinMediaAttachment{}

// JellyfinMediaAttachment Class MediaAttachment.
type JellyfinMediaAttachment struct {
	// Gets or sets the codec.
	Codec NullableString `json:"Codec,omitempty"`
	// Gets or sets the codec tag.
	CodecTag NullableString `json:"CodecTag,omitempty"`
	// Gets or sets the comment.
	Comment NullableString `json:"Comment,omitempty"`
	// Gets or sets the index.
	Index *int32 `json:"Index,omitempty"`
	// Gets or sets the filename.
	FileName NullableString `json:"FileName,omitempty"`
	// Gets or sets the MIME type.
	MimeType NullableString `json:"MimeType,omitempty"`
	// Gets or sets the delivery URL.
	DeliveryUrl NullableString `json:"DeliveryUrl,omitempty"`
}

// NewJellyfinMediaAttachment instantiates a new JellyfinMediaAttachment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJellyfinMediaAttachment() *JellyfinMediaAttachment {
	this := JellyfinMediaAttachment{}
	return &this
}

// NewJellyfinMediaAttachmentWithDefaults instantiates a new JellyfinMediaAttachment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJellyfinMediaAttachmentWithDefaults() *JellyfinMediaAttachment {
	this := JellyfinMediaAttachment{}
	return &this
}

// GetCodec returns the Codec field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinMediaAttachment) GetCodec() string {
	if o == nil || IsNil(o.Codec.Get()) {
		var ret string
		return ret
	}
	return *o.Codec.Get()
}

// GetCodecOk returns a tuple with the Codec field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinMediaAttachment) GetCodecOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Codec.Get(), o.Codec.IsSet()
}

// HasCodec returns a boolean if a field has been set.
func (o *JellyfinMediaAttachment) HasCodec() bool {
	if o != nil && o.Codec.IsSet() {
		return true
	}

	return false
}

// SetCodec gets a reference to the given NullableString and assigns it to the Codec field.
func (o *JellyfinMediaAttachment) SetCodec(v string) {
	o.Codec.Set(&v)
}
// SetCodecNil sets the value for Codec to be an explicit nil
func (o *JellyfinMediaAttachment) SetCodecNil() {
	o.Codec.Set(nil)
}

// UnsetCodec ensures that no value is present for Codec, not even an explicit nil
func (o *JellyfinMediaAttachment) UnsetCodec() {
	o.Codec.Unset()
}

// GetCodecTag returns the CodecTag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinMediaAttachment) GetCodecTag() string {
	if o == nil || IsNil(o.CodecTag.Get()) {
		var ret string
		return ret
	}
	return *o.CodecTag.Get()
}

// GetCodecTagOk returns a tuple with the CodecTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinMediaAttachment) GetCodecTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CodecTag.Get(), o.CodecTag.IsSet()
}

// HasCodecTag returns a boolean if a field has been set.
func (o *JellyfinMediaAttachment) HasCodecTag() bool {
	if o != nil && o.CodecTag.IsSet() {
		return true
	}

	return false
}

// SetCodecTag gets a reference to the given NullableString and assigns it to the CodecTag field.
func (o *JellyfinMediaAttachment) SetCodecTag(v string) {
	o.CodecTag.Set(&v)
}
// SetCodecTagNil sets the value for CodecTag to be an explicit nil
func (o *JellyfinMediaAttachment) SetCodecTagNil() {
	o.CodecTag.Set(nil)
}

// UnsetCodecTag ensures that no value is present for CodecTag, not even an explicit nil
func (o *JellyfinMediaAttachment) UnsetCodecTag() {
	o.CodecTag.Unset()
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinMediaAttachment) GetComment() string {
	if o == nil || IsNil(o.Comment.Get()) {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinMediaAttachment) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *JellyfinMediaAttachment) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *JellyfinMediaAttachment) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *JellyfinMediaAttachment) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *JellyfinMediaAttachment) UnsetComment() {
	o.Comment.Unset()
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *JellyfinMediaAttachment) GetIndex() int32 {
	if o == nil || IsNil(o.Index) {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinMediaAttachment) GetIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *JellyfinMediaAttachment) HasIndex() bool {
	if o != nil && !IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *JellyfinMediaAttachment) SetIndex(v int32) {
	o.Index = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinMediaAttachment) GetFileName() string {
	if o == nil || IsNil(o.FileName.Get()) {
		var ret string
		return ret
	}
	return *o.FileName.Get()
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinMediaAttachment) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileName.Get(), o.FileName.IsSet()
}

// HasFileName returns a boolean if a field has been set.
func (o *JellyfinMediaAttachment) HasFileName() bool {
	if o != nil && o.FileName.IsSet() {
		return true
	}

	return false
}

// SetFileName gets a reference to the given NullableString and assigns it to the FileName field.
func (o *JellyfinMediaAttachment) SetFileName(v string) {
	o.FileName.Set(&v)
}
// SetFileNameNil sets the value for FileName to be an explicit nil
func (o *JellyfinMediaAttachment) SetFileNameNil() {
	o.FileName.Set(nil)
}

// UnsetFileName ensures that no value is present for FileName, not even an explicit nil
func (o *JellyfinMediaAttachment) UnsetFileName() {
	o.FileName.Unset()
}

// GetMimeType returns the MimeType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinMediaAttachment) GetMimeType() string {
	if o == nil || IsNil(o.MimeType.Get()) {
		var ret string
		return ret
	}
	return *o.MimeType.Get()
}

// GetMimeTypeOk returns a tuple with the MimeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinMediaAttachment) GetMimeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MimeType.Get(), o.MimeType.IsSet()
}

// HasMimeType returns a boolean if a field has been set.
func (o *JellyfinMediaAttachment) HasMimeType() bool {
	if o != nil && o.MimeType.IsSet() {
		return true
	}

	return false
}

// SetMimeType gets a reference to the given NullableString and assigns it to the MimeType field.
func (o *JellyfinMediaAttachment) SetMimeType(v string) {
	o.MimeType.Set(&v)
}
// SetMimeTypeNil sets the value for MimeType to be an explicit nil
func (o *JellyfinMediaAttachment) SetMimeTypeNil() {
	o.MimeType.Set(nil)
}

// UnsetMimeType ensures that no value is present for MimeType, not even an explicit nil
func (o *JellyfinMediaAttachment) UnsetMimeType() {
	o.MimeType.Unset()
}

// GetDeliveryUrl returns the DeliveryUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinMediaAttachment) GetDeliveryUrl() string {
	if o == nil || IsNil(o.DeliveryUrl.Get()) {
		var ret string
		return ret
	}
	return *o.DeliveryUrl.Get()
}

// GetDeliveryUrlOk returns a tuple with the DeliveryUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinMediaAttachment) GetDeliveryUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeliveryUrl.Get(), o.DeliveryUrl.IsSet()
}

// HasDeliveryUrl returns a boolean if a field has been set.
func (o *JellyfinMediaAttachment) HasDeliveryUrl() bool {
	if o != nil && o.DeliveryUrl.IsSet() {
		return true
	}

	return false
}

// SetDeliveryUrl gets a reference to the given NullableString and assigns it to the DeliveryUrl field.
func (o *JellyfinMediaAttachment) SetDeliveryUrl(v string) {
	o.DeliveryUrl.Set(&v)
}
// SetDeliveryUrlNil sets the value for DeliveryUrl to be an explicit nil
func (o *JellyfinMediaAttachment) SetDeliveryUrlNil() {
	o.DeliveryUrl.Set(nil)
}

// UnsetDeliveryUrl ensures that no value is present for DeliveryUrl, not even an explicit nil
func (o *JellyfinMediaAttachment) UnsetDeliveryUrl() {
	o.DeliveryUrl.Unset()
}

func (o JellyfinMediaAttachment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JellyfinMediaAttachment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Codec.IsSet() {
		toSerialize["Codec"] = o.Codec.Get()
	}
	if o.CodecTag.IsSet() {
		toSerialize["CodecTag"] = o.CodecTag.Get()
	}
	if o.Comment.IsSet() {
		toSerialize["Comment"] = o.Comment.Get()
	}
	if !IsNil(o.Index) {
		toSerialize["Index"] = o.Index
	}
	if o.FileName.IsSet() {
		toSerialize["FileName"] = o.FileName.Get()
	}
	if o.MimeType.IsSet() {
		toSerialize["MimeType"] = o.MimeType.Get()
	}
	if o.DeliveryUrl.IsSet() {
		toSerialize["DeliveryUrl"] = o.DeliveryUrl.Get()
	}
	return toSerialize, nil
}

type NullableJellyfinMediaAttachment struct {
	value *JellyfinMediaAttachment
	isSet bool
}

func (v NullableJellyfinMediaAttachment) Get() *JellyfinMediaAttachment {
	return v.value
}

func (v *NullableJellyfinMediaAttachment) Set(val *JellyfinMediaAttachment) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinMediaAttachment) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinMediaAttachment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinMediaAttachment(val *JellyfinMediaAttachment) *NullableJellyfinMediaAttachment {
	return &NullableJellyfinMediaAttachment{value: val, isSet: true}
}

func (v NullableJellyfinMediaAttachment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinMediaAttachment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


