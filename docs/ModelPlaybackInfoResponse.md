# ModelPlaybackInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MediaSources** | Pointer to [**[]ModelModelMediaSourceInfo**](ModelModelMediaSourceInfo.md) | Gets or sets the media sources. | [optional] 
**PlaySessionId** | Pointer to **NullableString** | Gets or sets the play session identifier. | [optional] 
**ErrorCode** | Pointer to [**NullableModelModelPlaybackErrorCode**](ModelPlaybackErrorCode.md) | Gets or sets the error code. | [optional] 

## Methods

### NewModelPlaybackInfoResponse

`func NewModelPlaybackInfoResponse() *ModelPlaybackInfoResponse`

NewModelPlaybackInfoResponse instantiates a new ModelPlaybackInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelPlaybackInfoResponseWithDefaults

`func NewModelPlaybackInfoResponseWithDefaults() *ModelPlaybackInfoResponse`

NewModelPlaybackInfoResponseWithDefaults instantiates a new ModelPlaybackInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMediaSources

`func (o *ModelPlaybackInfoResponse) GetMediaSources() []ModelModelMediaSourceInfo`

GetMediaSources returns the MediaSources field if non-nil, zero value otherwise.

### GetMediaSourcesOk

`func (o *ModelPlaybackInfoResponse) GetMediaSourcesOk() (*[]ModelModelMediaSourceInfo, bool)`

GetMediaSourcesOk returns a tuple with the MediaSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaSources

`func (o *ModelPlaybackInfoResponse) SetMediaSources(v []ModelModelMediaSourceInfo)`

SetMediaSources sets MediaSources field to given value.

### HasMediaSources

`func (o *ModelPlaybackInfoResponse) HasMediaSources() bool`

HasMediaSources returns a boolean if a field has been set.

### GetPlaySessionId

`func (o *ModelPlaybackInfoResponse) GetPlaySessionId() string`

GetPlaySessionId returns the PlaySessionId field if non-nil, zero value otherwise.

### GetPlaySessionIdOk

`func (o *ModelPlaybackInfoResponse) GetPlaySessionIdOk() (*string, bool)`

GetPlaySessionIdOk returns a tuple with the PlaySessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaySessionId

`func (o *ModelPlaybackInfoResponse) SetPlaySessionId(v string)`

SetPlaySessionId sets PlaySessionId field to given value.

### HasPlaySessionId

`func (o *ModelPlaybackInfoResponse) HasPlaySessionId() bool`

HasPlaySessionId returns a boolean if a field has been set.

### SetPlaySessionIdNil

`func (o *ModelPlaybackInfoResponse) SetPlaySessionIdNil(b bool)`

 SetPlaySessionIdNil sets the value for PlaySessionId to be an explicit nil

### UnsetPlaySessionId
`func (o *ModelPlaybackInfoResponse) UnsetPlaySessionId()`

UnsetPlaySessionId ensures that no value is present for PlaySessionId, not even an explicit nil
### GetErrorCode

`func (o *ModelPlaybackInfoResponse) GetErrorCode() ModelModelPlaybackErrorCode`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ModelPlaybackInfoResponse) GetErrorCodeOk() (*ModelModelPlaybackErrorCode, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ModelPlaybackInfoResponse) SetErrorCode(v ModelModelPlaybackErrorCode)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ModelPlaybackInfoResponse) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *ModelPlaybackInfoResponse) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *ModelPlaybackInfoResponse) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


