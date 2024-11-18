# ModelDirectPlayProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Container** | Pointer to **string** | Gets or sets the container. | [optional] 
**AudioCodec** | Pointer to **NullableString** | Gets or sets the audio codec. | [optional] 
**VideoCodec** | Pointer to **NullableString** | Gets or sets the video codec. | [optional] 
**Type** | Pointer to [**ModelModelDlnaProfileType**](ModelDlnaProfileType.md) | Gets or sets the Dlna profile type. | [optional] 

## Methods

### NewModelDirectPlayProfile

`func NewModelDirectPlayProfile() *ModelDirectPlayProfile`

NewModelDirectPlayProfile instantiates a new ModelDirectPlayProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelDirectPlayProfileWithDefaults

`func NewModelDirectPlayProfileWithDefaults() *ModelDirectPlayProfile`

NewModelDirectPlayProfileWithDefaults instantiates a new ModelDirectPlayProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainer

`func (o *ModelDirectPlayProfile) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *ModelDirectPlayProfile) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *ModelDirectPlayProfile) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *ModelDirectPlayProfile) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetAudioCodec

`func (o *ModelDirectPlayProfile) GetAudioCodec() string`

GetAudioCodec returns the AudioCodec field if non-nil, zero value otherwise.

### GetAudioCodecOk

`func (o *ModelDirectPlayProfile) GetAudioCodecOk() (*string, bool)`

GetAudioCodecOk returns a tuple with the AudioCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioCodec

`func (o *ModelDirectPlayProfile) SetAudioCodec(v string)`

SetAudioCodec sets AudioCodec field to given value.

### HasAudioCodec

`func (o *ModelDirectPlayProfile) HasAudioCodec() bool`

HasAudioCodec returns a boolean if a field has been set.

### SetAudioCodecNil

`func (o *ModelDirectPlayProfile) SetAudioCodecNil(b bool)`

 SetAudioCodecNil sets the value for AudioCodec to be an explicit nil

### UnsetAudioCodec
`func (o *ModelDirectPlayProfile) UnsetAudioCodec()`

UnsetAudioCodec ensures that no value is present for AudioCodec, not even an explicit nil
### GetVideoCodec

`func (o *ModelDirectPlayProfile) GetVideoCodec() string`

GetVideoCodec returns the VideoCodec field if non-nil, zero value otherwise.

### GetVideoCodecOk

`func (o *ModelDirectPlayProfile) GetVideoCodecOk() (*string, bool)`

GetVideoCodecOk returns a tuple with the VideoCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoCodec

`func (o *ModelDirectPlayProfile) SetVideoCodec(v string)`

SetVideoCodec sets VideoCodec field to given value.

### HasVideoCodec

`func (o *ModelDirectPlayProfile) HasVideoCodec() bool`

HasVideoCodec returns a boolean if a field has been set.

### SetVideoCodecNil

`func (o *ModelDirectPlayProfile) SetVideoCodecNil(b bool)`

 SetVideoCodecNil sets the value for VideoCodec to be an explicit nil

### UnsetVideoCodec
`func (o *ModelDirectPlayProfile) UnsetVideoCodec()`

UnsetVideoCodec ensures that no value is present for VideoCodec, not even an explicit nil
### GetType

`func (o *ModelDirectPlayProfile) GetType() ModelModelDlnaProfileType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelDirectPlayProfile) GetTypeOk() (*ModelModelDlnaProfileType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelDirectPlayProfile) SetType(v ModelModelDlnaProfileType)`

SetType sets Type field to given value.

### HasType

`func (o *ModelDirectPlayProfile) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


