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

// checks if the JellyfinPlaybackProgressInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JellyfinPlaybackProgressInfo{}

// JellyfinPlaybackProgressInfo Class PlaybackProgressInfo.
type JellyfinPlaybackProgressInfo struct {
	// Gets or sets a value indicating whether this instance can seek.
	CanSeek *bool `json:"CanSeek,omitempty"`
	// Gets or sets the item.
	Item NullableJellyfinBaseItemDto `json:"Item,omitempty"`
	// Gets or sets the item identifier.
	ItemId *string `json:"ItemId,omitempty"`
	// Gets or sets the session id.
	SessionId NullableString `json:"SessionId,omitempty"`
	// Gets or sets the media version identifier.
	MediaSourceId NullableString `json:"MediaSourceId,omitempty"`
	// Gets or sets the index of the audio stream.
	AudioStreamIndex NullableInt32 `json:"AudioStreamIndex,omitempty"`
	// Gets or sets the index of the subtitle stream.
	SubtitleStreamIndex NullableInt32 `json:"SubtitleStreamIndex,omitempty"`
	// Gets or sets a value indicating whether this instance is paused.
	IsPaused *bool `json:"IsPaused,omitempty"`
	// Gets or sets a value indicating whether this instance is muted.
	IsMuted *bool `json:"IsMuted,omitempty"`
	// Gets or sets the position ticks.
	PositionTicks NullableInt64 `json:"PositionTicks,omitempty"`
	PlaybackStartTimeTicks NullableInt64 `json:"PlaybackStartTimeTicks,omitempty"`
	// Gets or sets the volume level.
	VolumeLevel NullableInt32 `json:"VolumeLevel,omitempty"`
	Brightness NullableInt32 `json:"Brightness,omitempty"`
	AspectRatio NullableString `json:"AspectRatio,omitempty"`
	// Gets or sets the play method.
	PlayMethod *JellyfinPlayMethod `json:"PlayMethod,omitempty"`
	// Gets or sets the live stream identifier.
	LiveStreamId NullableString `json:"LiveStreamId,omitempty"`
	// Gets or sets the play session identifier.
	PlaySessionId NullableString `json:"PlaySessionId,omitempty"`
	// Gets or sets the repeat mode.
	RepeatMode *JellyfinRepeatMode `json:"RepeatMode,omitempty"`
	// Gets or sets the playback order.
	PlaybackOrder *JellyfinPlaybackOrder `json:"PlaybackOrder,omitempty"`
	NowPlayingQueue []JellyfinQueueItem `json:"NowPlayingQueue,omitempty"`
	PlaylistItemId NullableString `json:"PlaylistItemId,omitempty"`
}

// NewJellyfinPlaybackProgressInfo instantiates a new JellyfinPlaybackProgressInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJellyfinPlaybackProgressInfo() *JellyfinPlaybackProgressInfo {
	this := JellyfinPlaybackProgressInfo{}
	return &this
}

// NewJellyfinPlaybackProgressInfoWithDefaults instantiates a new JellyfinPlaybackProgressInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJellyfinPlaybackProgressInfoWithDefaults() *JellyfinPlaybackProgressInfo {
	this := JellyfinPlaybackProgressInfo{}
	return &this
}

// GetCanSeek returns the CanSeek field value if set, zero value otherwise.
func (o *JellyfinPlaybackProgressInfo) GetCanSeek() bool {
	if o == nil || IsNil(o.CanSeek) {
		var ret bool
		return ret
	}
	return *o.CanSeek
}

// GetCanSeekOk returns a tuple with the CanSeek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinPlaybackProgressInfo) GetCanSeekOk() (*bool, bool) {
	if o == nil || IsNil(o.CanSeek) {
		return nil, false
	}
	return o.CanSeek, true
}

// HasCanSeek returns a boolean if a field has been set.
func (o *JellyfinPlaybackProgressInfo) HasCanSeek() bool {
	if o != nil && !IsNil(o.CanSeek) {
		return true
	}

	return false
}

// SetCanSeek gets a reference to the given bool and assigns it to the CanSeek field.
func (o *JellyfinPlaybackProgressInfo) SetCanSeek(v bool) {
	o.CanSeek = &v
}

// GetItem returns the Item field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinPlaybackProgressInfo) GetItem() JellyfinBaseItemDto {
	if o == nil || IsNil(o.Item.Get()) {
		var ret JellyfinBaseItemDto
		return ret
	}
	return *o.Item.Get()
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinPlaybackProgressInfo) GetItemOk() (*JellyfinBaseItemDto, bool) {
	if o == nil {
		return nil, false
	}
	return o.Item.Get(), o.Item.IsSet()
}

// HasItem returns a boolean if a field has been set.
func (o *JellyfinPlaybackProgressInfo) HasItem() bool {
	if o != nil && o.Item.IsSet() {
		return true
	}

	return false
}

// SetItem gets a reference to the given NullableJellyfinBaseItemDto and assigns it to the Item field.
func (o *JellyfinPlaybackProgressInfo) SetItem(v JellyfinBaseItemDto) {
	o.Item.Set(&v)
}
// SetItemNil sets the value for Item to be an explicit nil
func (o *JellyfinPlaybackProgressInfo) SetItemNil() {
	o.Item.Set(nil)
}

// UnsetItem ensures that no value is present for Item, not even an explicit nil
func (o *JellyfinPlaybackProgressInfo) UnsetItem() {
	o.Item.Unset()
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *JellyfinPlaybackProgressInfo) GetItemId() string {
	if o == nil || IsNil(o.ItemId) {
		var ret string
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinPlaybackProgressInfo) GetItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.ItemId) {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *JellyfinPlaybackProgressInfo) HasItemId() bool {
	if o != nil && !IsNil(o.ItemId) {
		return true
	}

	return false
}

// SetItemId gets a reference to the given string and assigns it to the ItemId field.
func (o *JellyfinPlaybackProgressInfo) SetItemId(v string) {
	o.ItemId = &v
}

// GetSessionId returns the SessionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinPlaybackProgressInfo) GetSessionId() string {
	if o == nil || IsNil(o.SessionId.Get()) {
		var ret string
		return ret
	}
	return *o.SessionId.Get()
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinPlaybackProgressInfo) GetSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SessionId.Get(), o.SessionId.IsSet()
}

// HasSessionId returns a boolean if a field has been set.
func (o *JellyfinPlaybackProgressInfo) HasSessionId() bool {
	if o != nil && o.SessionId.IsSet() {
		return true
	}

	return false
}

// SetSessionId gets a reference to the given NullableString and assigns it to the SessionId field.
func (o *JellyfinPlaybackProgressInfo) SetSessionId(v string) {
	o.SessionId.Set(&v)
}
// SetSessionIdNil sets the value for SessionId to be an explicit nil
func (o *JellyfinPlaybackProgressInfo) SetSessionIdNil() {
	o.SessionId.Set(nil)
}

// UnsetSessionId ensures that no value is present for SessionId, not even an explicit nil
func (o *JellyfinPlaybackProgressInfo) UnsetSessionId() {
	o.SessionId.Unset()
}

// GetMediaSourceId returns the MediaSourceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinPlaybackProgressInfo) GetMediaSourceId() string {
	if o == nil || IsNil(o.MediaSourceId.Get()) {
		var ret string
		return ret
	}
	return *o.MediaSourceId.Get()
}

// GetMediaSourceIdOk returns a tuple with the MediaSourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinPlaybackProgressInfo) GetMediaSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MediaSourceId.Get(), o.MediaSourceId.IsSet()
}

// HasMediaSourceId returns a boolean if a field has been set.
func (o *JellyfinPlaybackProgressInfo) HasMediaSourceId() bool {
	if o != nil && o.MediaSourceId.IsSet() {
		return true
	}

	return false
}

// SetMediaSourceId gets a reference to the given NullableString and assigns it to the MediaSourceId field.
func (o *JellyfinPlaybackProgressInfo) SetMediaSourceId(v string) {
	o.MediaSourceId.Set(&v)
}
// SetMediaSourceIdNil sets the value for MediaSourceId to be an explicit nil
func (o *JellyfinPlaybackProgressInfo) SetMediaSourceIdNil() {
	o.MediaSourceId.Set(nil)
}

// UnsetMediaSourceId ensures that no value is present for MediaSourceId, not even an explicit nil
func (o *JellyfinPlaybackProgressInfo) UnsetMediaSourceId() {
	o.MediaSourceId.Unset()
}

// GetAudioStreamIndex returns the AudioStreamIndex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinPlaybackProgressInfo) GetAudioStreamIndex() int32 {
	if o == nil || IsNil(o.AudioStreamIndex.Get()) {
		var ret int32
		return ret
	}
	return *o.AudioStreamIndex.Get()
}

// GetAudioStreamIndexOk returns a tuple with the AudioStreamIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinPlaybackProgressInfo) GetAudioStreamIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AudioStreamIndex.Get(), o.AudioStreamIndex.IsSet()
}

// HasAudioStreamIndex returns a boolean if a field has been set.
func (o *JellyfinPlaybackProgressInfo) HasAudioStreamIndex() bool {
	if o != nil && o.AudioStreamIndex.IsSet() {
		return true
	}

	return false
}

// SetAudioStreamIndex gets a reference to the given NullableInt32 and assigns it to the AudioStreamIndex field.
func (o *JellyfinPlaybackProgressInfo) SetAudioStreamIndex(v int32) {
	o.AudioStreamIndex.Set(&v)
}
// SetAudioStreamIndexNil sets the value for AudioStreamIndex to be an explicit nil
func (o *JellyfinPlaybackProgressInfo) SetAudioStreamIndexNil() {
	o.AudioStreamIndex.Set(nil)
}

// UnsetAudioStreamIndex ensures that no value is present for AudioStreamIndex, not even an explicit nil
func (o *JellyfinPlaybackProgressInfo) UnsetAudioStreamIndex() {
	o.AudioStreamIndex.Unset()
}

// GetSubtitleStreamIndex returns the SubtitleStreamIndex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinPlaybackProgressInfo) GetSubtitleStreamIndex() int32 {
	if o == nil || IsNil(o.SubtitleStreamIndex.Get()) {
		var ret int32
		return ret
	}
	return *o.SubtitleStreamIndex.Get()
}

// GetSubtitleStreamIndexOk returns a tuple with the SubtitleStreamIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinPlaybackProgressInfo) GetSubtitleStreamIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubtitleStreamIndex.Get(), o.SubtitleStreamIndex.IsSet()
}

// HasSubtitleStreamIndex returns a boolean if a field has been set.
func (o *JellyfinPlaybackProgressInfo) HasSubtitleStreamIndex() bool {
	if o != nil && o.SubtitleStreamIndex.IsSet() {
		return true
	}

	return false
}

// SetSubtitleStreamIndex gets a reference to the given NullableInt32 and assigns it to the SubtitleStreamIndex field.
func (o *JellyfinPlaybackProgressInfo) SetSubtitleStreamIndex(v int32) {
	o.SubtitleStreamIndex.Set(&v)
}
// SetSubtitleStreamIndexNil sets the value for SubtitleStreamIndex to be an explicit nil
func (o *JellyfinPlaybackProgressInfo) SetSubtitleStreamIndexNil() {
	o.SubtitleStreamIndex.Set(nil)
}

// UnsetSubtitleStreamIndex ensures that no value is present for SubtitleStreamIndex, not even an explicit nil
func (o *JellyfinPlaybackProgressInfo) UnsetSubtitleStreamIndex() {
	o.SubtitleStreamIndex.Unset()
}

// GetIsPaused returns the IsPaused field value if set, zero value otherwise.
func (o *JellyfinPlaybackProgressInfo) GetIsPaused() bool {
	if o == nil || IsNil(o.IsPaused) {
		var ret bool
		return ret
	}
	return *o.IsPaused
}

// GetIsPausedOk returns a tuple with the IsPaused field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinPlaybackProgressInfo) GetIsPausedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPaused) {
		return nil, false
	}
	return o.IsPaused, true
}

// HasIsPaused returns a boolean if a field has been set.
func (o *JellyfinPlaybackProgressInfo) HasIsPaused() bool {
	if o != nil && !IsNil(o.IsPaused) {
		return true
	}

	return false
}

// SetIsPaused gets a reference to the given bool and assigns it to the IsPaused field.
func (o *JellyfinPlaybackProgressInfo) SetIsPaused(v bool) {
	o.IsPaused = &v
}

// GetIsMuted returns the IsMuted field value if set, zero value otherwise.
func (o *JellyfinPlaybackProgressInfo) GetIsMuted() bool {
	if o == nil || IsNil(o.IsMuted) {
		var ret bool
		return ret
	}
	return *o.IsMuted
}

// GetIsMutedOk returns a tuple with the IsMuted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinPlaybackProgressInfo) GetIsMutedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMuted) {
		return nil, false
	}
	return o.IsMuted, true
}

// HasIsMuted returns a boolean if a field has been set.
func (o *JellyfinPlaybackProgressInfo) HasIsMuted() bool {
	if o != nil && !IsNil(o.IsMuted) {
		return true
	}

	return false
}

// SetIsMuted gets a reference to the given bool and assigns it to the IsMuted field.
func (o *JellyfinPlaybackProgressInfo) SetIsMuted(v bool) {
	o.IsMuted = &v
}

// GetPositionTicks returns the PositionTicks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinPlaybackProgressInfo) GetPositionTicks() int64 {
	if o == nil || IsNil(o.PositionTicks.Get()) {
		var ret int64
		return ret
	}
	return *o.PositionTicks.Get()
}

// GetPositionTicksOk returns a tuple with the PositionTicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinPlaybackProgressInfo) GetPositionTicksOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PositionTicks.Get(), o.PositionTicks.IsSet()
}

// HasPositionTicks returns a boolean if a field has been set.
func (o *JellyfinPlaybackProgressInfo) HasPositionTicks() bool {
	if o != nil && o.PositionTicks.IsSet() {
		return true
	}

	return false
}

// SetPositionTicks gets a reference to the given NullableInt64 and assigns it to the PositionTicks field.
func (o *JellyfinPlaybackProgressInfo) SetPositionTicks(v int64) {
	o.PositionTicks.Set(&v)
}
// SetPositionTicksNil sets the value for PositionTicks to be an explicit nil
func (o *JellyfinPlaybackProgressInfo) SetPositionTicksNil() {
	o.PositionTicks.Set(nil)
}

// UnsetPositionTicks ensures that no value is present for PositionTicks, not even an explicit nil
func (o *JellyfinPlaybackProgressInfo) UnsetPositionTicks() {
	o.PositionTicks.Unset()
}

// GetPlaybackStartTimeTicks returns the PlaybackStartTimeTicks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinPlaybackProgressInfo) GetPlaybackStartTimeTicks() int64 {
	if o == nil || IsNil(o.PlaybackStartTimeTicks.Get()) {
		var ret int64
		return ret
	}
	return *o.PlaybackStartTimeTicks.Get()
}

// GetPlaybackStartTimeTicksOk returns a tuple with the PlaybackStartTimeTicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinPlaybackProgressInfo) GetPlaybackStartTimeTicksOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlaybackStartTimeTicks.Get(), o.PlaybackStartTimeTicks.IsSet()
}

// HasPlaybackStartTimeTicks returns a boolean if a field has been set.
func (o *JellyfinPlaybackProgressInfo) HasPlaybackStartTimeTicks() bool {
	if o != nil && o.PlaybackStartTimeTicks.IsSet() {
		return true
	}

	return false
}

// SetPlaybackStartTimeTicks gets a reference to the given NullableInt64 and assigns it to the PlaybackStartTimeTicks field.
func (o *JellyfinPlaybackProgressInfo) SetPlaybackStartTimeTicks(v int64) {
	o.PlaybackStartTimeTicks.Set(&v)
}
// SetPlaybackStartTimeTicksNil sets the value for PlaybackStartTimeTicks to be an explicit nil
func (o *JellyfinPlaybackProgressInfo) SetPlaybackStartTimeTicksNil() {
	o.PlaybackStartTimeTicks.Set(nil)
}

// UnsetPlaybackStartTimeTicks ensures that no value is present for PlaybackStartTimeTicks, not even an explicit nil
func (o *JellyfinPlaybackProgressInfo) UnsetPlaybackStartTimeTicks() {
	o.PlaybackStartTimeTicks.Unset()
}

// GetVolumeLevel returns the VolumeLevel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinPlaybackProgressInfo) GetVolumeLevel() int32 {
	if o == nil || IsNil(o.VolumeLevel.Get()) {
		var ret int32
		return ret
	}
	return *o.VolumeLevel.Get()
}

// GetVolumeLevelOk returns a tuple with the VolumeLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinPlaybackProgressInfo) GetVolumeLevelOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.VolumeLevel.Get(), o.VolumeLevel.IsSet()
}

// HasVolumeLevel returns a boolean if a field has been set.
func (o *JellyfinPlaybackProgressInfo) HasVolumeLevel() bool {
	if o != nil && o.VolumeLevel.IsSet() {
		return true
	}

	return false
}

// SetVolumeLevel gets a reference to the given NullableInt32 and assigns it to the VolumeLevel field.
func (o *JellyfinPlaybackProgressInfo) SetVolumeLevel(v int32) {
	o.VolumeLevel.Set(&v)
}
// SetVolumeLevelNil sets the value for VolumeLevel to be an explicit nil
func (o *JellyfinPlaybackProgressInfo) SetVolumeLevelNil() {
	o.VolumeLevel.Set(nil)
}

// UnsetVolumeLevel ensures that no value is present for VolumeLevel, not even an explicit nil
func (o *JellyfinPlaybackProgressInfo) UnsetVolumeLevel() {
	o.VolumeLevel.Unset()
}

// GetBrightness returns the Brightness field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinPlaybackProgressInfo) GetBrightness() int32 {
	if o == nil || IsNil(o.Brightness.Get()) {
		var ret int32
		return ret
	}
	return *o.Brightness.Get()
}

// GetBrightnessOk returns a tuple with the Brightness field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinPlaybackProgressInfo) GetBrightnessOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Brightness.Get(), o.Brightness.IsSet()
}

// HasBrightness returns a boolean if a field has been set.
func (o *JellyfinPlaybackProgressInfo) HasBrightness() bool {
	if o != nil && o.Brightness.IsSet() {
		return true
	}

	return false
}

// SetBrightness gets a reference to the given NullableInt32 and assigns it to the Brightness field.
func (o *JellyfinPlaybackProgressInfo) SetBrightness(v int32) {
	o.Brightness.Set(&v)
}
// SetBrightnessNil sets the value for Brightness to be an explicit nil
func (o *JellyfinPlaybackProgressInfo) SetBrightnessNil() {
	o.Brightness.Set(nil)
}

// UnsetBrightness ensures that no value is present for Brightness, not even an explicit nil
func (o *JellyfinPlaybackProgressInfo) UnsetBrightness() {
	o.Brightness.Unset()
}

// GetAspectRatio returns the AspectRatio field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinPlaybackProgressInfo) GetAspectRatio() string {
	if o == nil || IsNil(o.AspectRatio.Get()) {
		var ret string
		return ret
	}
	return *o.AspectRatio.Get()
}

// GetAspectRatioOk returns a tuple with the AspectRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinPlaybackProgressInfo) GetAspectRatioOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AspectRatio.Get(), o.AspectRatio.IsSet()
}

// HasAspectRatio returns a boolean if a field has been set.
func (o *JellyfinPlaybackProgressInfo) HasAspectRatio() bool {
	if o != nil && o.AspectRatio.IsSet() {
		return true
	}

	return false
}

// SetAspectRatio gets a reference to the given NullableString and assigns it to the AspectRatio field.
func (o *JellyfinPlaybackProgressInfo) SetAspectRatio(v string) {
	o.AspectRatio.Set(&v)
}
// SetAspectRatioNil sets the value for AspectRatio to be an explicit nil
func (o *JellyfinPlaybackProgressInfo) SetAspectRatioNil() {
	o.AspectRatio.Set(nil)
}

// UnsetAspectRatio ensures that no value is present for AspectRatio, not even an explicit nil
func (o *JellyfinPlaybackProgressInfo) UnsetAspectRatio() {
	o.AspectRatio.Unset()
}

// GetPlayMethod returns the PlayMethod field value if set, zero value otherwise.
func (o *JellyfinPlaybackProgressInfo) GetPlayMethod() JellyfinPlayMethod {
	if o == nil || IsNil(o.PlayMethod) {
		var ret JellyfinPlayMethod
		return ret
	}
	return *o.PlayMethod
}

// GetPlayMethodOk returns a tuple with the PlayMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinPlaybackProgressInfo) GetPlayMethodOk() (*JellyfinPlayMethod, bool) {
	if o == nil || IsNil(o.PlayMethod) {
		return nil, false
	}
	return o.PlayMethod, true
}

// HasPlayMethod returns a boolean if a field has been set.
func (o *JellyfinPlaybackProgressInfo) HasPlayMethod() bool {
	if o != nil && !IsNil(o.PlayMethod) {
		return true
	}

	return false
}

// SetPlayMethod gets a reference to the given JellyfinPlayMethod and assigns it to the PlayMethod field.
func (o *JellyfinPlaybackProgressInfo) SetPlayMethod(v JellyfinPlayMethod) {
	o.PlayMethod = &v
}

// GetLiveStreamId returns the LiveStreamId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinPlaybackProgressInfo) GetLiveStreamId() string {
	if o == nil || IsNil(o.LiveStreamId.Get()) {
		var ret string
		return ret
	}
	return *o.LiveStreamId.Get()
}

// GetLiveStreamIdOk returns a tuple with the LiveStreamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinPlaybackProgressInfo) GetLiveStreamIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LiveStreamId.Get(), o.LiveStreamId.IsSet()
}

// HasLiveStreamId returns a boolean if a field has been set.
func (o *JellyfinPlaybackProgressInfo) HasLiveStreamId() bool {
	if o != nil && o.LiveStreamId.IsSet() {
		return true
	}

	return false
}

// SetLiveStreamId gets a reference to the given NullableString and assigns it to the LiveStreamId field.
func (o *JellyfinPlaybackProgressInfo) SetLiveStreamId(v string) {
	o.LiveStreamId.Set(&v)
}
// SetLiveStreamIdNil sets the value for LiveStreamId to be an explicit nil
func (o *JellyfinPlaybackProgressInfo) SetLiveStreamIdNil() {
	o.LiveStreamId.Set(nil)
}

// UnsetLiveStreamId ensures that no value is present for LiveStreamId, not even an explicit nil
func (o *JellyfinPlaybackProgressInfo) UnsetLiveStreamId() {
	o.LiveStreamId.Unset()
}

// GetPlaySessionId returns the PlaySessionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinPlaybackProgressInfo) GetPlaySessionId() string {
	if o == nil || IsNil(o.PlaySessionId.Get()) {
		var ret string
		return ret
	}
	return *o.PlaySessionId.Get()
}

// GetPlaySessionIdOk returns a tuple with the PlaySessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinPlaybackProgressInfo) GetPlaySessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlaySessionId.Get(), o.PlaySessionId.IsSet()
}

// HasPlaySessionId returns a boolean if a field has been set.
func (o *JellyfinPlaybackProgressInfo) HasPlaySessionId() bool {
	if o != nil && o.PlaySessionId.IsSet() {
		return true
	}

	return false
}

// SetPlaySessionId gets a reference to the given NullableString and assigns it to the PlaySessionId field.
func (o *JellyfinPlaybackProgressInfo) SetPlaySessionId(v string) {
	o.PlaySessionId.Set(&v)
}
// SetPlaySessionIdNil sets the value for PlaySessionId to be an explicit nil
func (o *JellyfinPlaybackProgressInfo) SetPlaySessionIdNil() {
	o.PlaySessionId.Set(nil)
}

// UnsetPlaySessionId ensures that no value is present for PlaySessionId, not even an explicit nil
func (o *JellyfinPlaybackProgressInfo) UnsetPlaySessionId() {
	o.PlaySessionId.Unset()
}

// GetRepeatMode returns the RepeatMode field value if set, zero value otherwise.
func (o *JellyfinPlaybackProgressInfo) GetRepeatMode() JellyfinRepeatMode {
	if o == nil || IsNil(o.RepeatMode) {
		var ret JellyfinRepeatMode
		return ret
	}
	return *o.RepeatMode
}

// GetRepeatModeOk returns a tuple with the RepeatMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinPlaybackProgressInfo) GetRepeatModeOk() (*JellyfinRepeatMode, bool) {
	if o == nil || IsNil(o.RepeatMode) {
		return nil, false
	}
	return o.RepeatMode, true
}

// HasRepeatMode returns a boolean if a field has been set.
func (o *JellyfinPlaybackProgressInfo) HasRepeatMode() bool {
	if o != nil && !IsNil(o.RepeatMode) {
		return true
	}

	return false
}

// SetRepeatMode gets a reference to the given JellyfinRepeatMode and assigns it to the RepeatMode field.
func (o *JellyfinPlaybackProgressInfo) SetRepeatMode(v JellyfinRepeatMode) {
	o.RepeatMode = &v
}

// GetPlaybackOrder returns the PlaybackOrder field value if set, zero value otherwise.
func (o *JellyfinPlaybackProgressInfo) GetPlaybackOrder() JellyfinPlaybackOrder {
	if o == nil || IsNil(o.PlaybackOrder) {
		var ret JellyfinPlaybackOrder
		return ret
	}
	return *o.PlaybackOrder
}

// GetPlaybackOrderOk returns a tuple with the PlaybackOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinPlaybackProgressInfo) GetPlaybackOrderOk() (*JellyfinPlaybackOrder, bool) {
	if o == nil || IsNil(o.PlaybackOrder) {
		return nil, false
	}
	return o.PlaybackOrder, true
}

// HasPlaybackOrder returns a boolean if a field has been set.
func (o *JellyfinPlaybackProgressInfo) HasPlaybackOrder() bool {
	if o != nil && !IsNil(o.PlaybackOrder) {
		return true
	}

	return false
}

// SetPlaybackOrder gets a reference to the given JellyfinPlaybackOrder and assigns it to the PlaybackOrder field.
func (o *JellyfinPlaybackProgressInfo) SetPlaybackOrder(v JellyfinPlaybackOrder) {
	o.PlaybackOrder = &v
}

// GetNowPlayingQueue returns the NowPlayingQueue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinPlaybackProgressInfo) GetNowPlayingQueue() []JellyfinQueueItem {
	if o == nil {
		var ret []JellyfinQueueItem
		return ret
	}
	return o.NowPlayingQueue
}

// GetNowPlayingQueueOk returns a tuple with the NowPlayingQueue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinPlaybackProgressInfo) GetNowPlayingQueueOk() ([]JellyfinQueueItem, bool) {
	if o == nil || IsNil(o.NowPlayingQueue) {
		return nil, false
	}
	return o.NowPlayingQueue, true
}

// HasNowPlayingQueue returns a boolean if a field has been set.
func (o *JellyfinPlaybackProgressInfo) HasNowPlayingQueue() bool {
	if o != nil && !IsNil(o.NowPlayingQueue) {
		return true
	}

	return false
}

// SetNowPlayingQueue gets a reference to the given []JellyfinQueueItem and assigns it to the NowPlayingQueue field.
func (o *JellyfinPlaybackProgressInfo) SetNowPlayingQueue(v []JellyfinQueueItem) {
	o.NowPlayingQueue = v
}

// GetPlaylistItemId returns the PlaylistItemId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinPlaybackProgressInfo) GetPlaylistItemId() string {
	if o == nil || IsNil(o.PlaylistItemId.Get()) {
		var ret string
		return ret
	}
	return *o.PlaylistItemId.Get()
}

// GetPlaylistItemIdOk returns a tuple with the PlaylistItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinPlaybackProgressInfo) GetPlaylistItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlaylistItemId.Get(), o.PlaylistItemId.IsSet()
}

// HasPlaylistItemId returns a boolean if a field has been set.
func (o *JellyfinPlaybackProgressInfo) HasPlaylistItemId() bool {
	if o != nil && o.PlaylistItemId.IsSet() {
		return true
	}

	return false
}

// SetPlaylistItemId gets a reference to the given NullableString and assigns it to the PlaylistItemId field.
func (o *JellyfinPlaybackProgressInfo) SetPlaylistItemId(v string) {
	o.PlaylistItemId.Set(&v)
}
// SetPlaylistItemIdNil sets the value for PlaylistItemId to be an explicit nil
func (o *JellyfinPlaybackProgressInfo) SetPlaylistItemIdNil() {
	o.PlaylistItemId.Set(nil)
}

// UnsetPlaylistItemId ensures that no value is present for PlaylistItemId, not even an explicit nil
func (o *JellyfinPlaybackProgressInfo) UnsetPlaylistItemId() {
	o.PlaylistItemId.Unset()
}

func (o JellyfinPlaybackProgressInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JellyfinPlaybackProgressInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CanSeek) {
		toSerialize["CanSeek"] = o.CanSeek
	}
	if o.Item.IsSet() {
		toSerialize["Item"] = o.Item.Get()
	}
	if !IsNil(o.ItemId) {
		toSerialize["ItemId"] = o.ItemId
	}
	if o.SessionId.IsSet() {
		toSerialize["SessionId"] = o.SessionId.Get()
	}
	if o.MediaSourceId.IsSet() {
		toSerialize["MediaSourceId"] = o.MediaSourceId.Get()
	}
	if o.AudioStreamIndex.IsSet() {
		toSerialize["AudioStreamIndex"] = o.AudioStreamIndex.Get()
	}
	if o.SubtitleStreamIndex.IsSet() {
		toSerialize["SubtitleStreamIndex"] = o.SubtitleStreamIndex.Get()
	}
	if !IsNil(o.IsPaused) {
		toSerialize["IsPaused"] = o.IsPaused
	}
	if !IsNil(o.IsMuted) {
		toSerialize["IsMuted"] = o.IsMuted
	}
	if o.PositionTicks.IsSet() {
		toSerialize["PositionTicks"] = o.PositionTicks.Get()
	}
	if o.PlaybackStartTimeTicks.IsSet() {
		toSerialize["PlaybackStartTimeTicks"] = o.PlaybackStartTimeTicks.Get()
	}
	if o.VolumeLevel.IsSet() {
		toSerialize["VolumeLevel"] = o.VolumeLevel.Get()
	}
	if o.Brightness.IsSet() {
		toSerialize["Brightness"] = o.Brightness.Get()
	}
	if o.AspectRatio.IsSet() {
		toSerialize["AspectRatio"] = o.AspectRatio.Get()
	}
	if !IsNil(o.PlayMethod) {
		toSerialize["PlayMethod"] = o.PlayMethod
	}
	if o.LiveStreamId.IsSet() {
		toSerialize["LiveStreamId"] = o.LiveStreamId.Get()
	}
	if o.PlaySessionId.IsSet() {
		toSerialize["PlaySessionId"] = o.PlaySessionId.Get()
	}
	if !IsNil(o.RepeatMode) {
		toSerialize["RepeatMode"] = o.RepeatMode
	}
	if !IsNil(o.PlaybackOrder) {
		toSerialize["PlaybackOrder"] = o.PlaybackOrder
	}
	if o.NowPlayingQueue != nil {
		toSerialize["NowPlayingQueue"] = o.NowPlayingQueue
	}
	if o.PlaylistItemId.IsSet() {
		toSerialize["PlaylistItemId"] = o.PlaylistItemId.Get()
	}
	return toSerialize, nil
}

type NullableJellyfinPlaybackProgressInfo struct {
	value *JellyfinPlaybackProgressInfo
	isSet bool
}

func (v NullableJellyfinPlaybackProgressInfo) Get() *JellyfinPlaybackProgressInfo {
	return v.value
}

func (v *NullableJellyfinPlaybackProgressInfo) Set(val *JellyfinPlaybackProgressInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinPlaybackProgressInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinPlaybackProgressInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinPlaybackProgressInfo(val *JellyfinPlaybackProgressInfo) *NullableJellyfinPlaybackProgressInfo {
	return &NullableJellyfinPlaybackProgressInfo{value: val, isSet: true}
}

func (v NullableJellyfinPlaybackProgressInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinPlaybackProgressInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

