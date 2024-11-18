# ModelSessionInfoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlayState** | Pointer to [**NullableModelModelPlayerStateInfo**](ModelPlayerStateInfo.md) | Gets or sets the play state. | [optional] 
**AdditionalUsers** | Pointer to [**[]ModelModelSessionUserInfo**](ModelModelSessionUserInfo.md) | Gets or sets the additional users. | [optional] 
**Capabilities** | Pointer to [**NullableModelModelClientCapabilitiesDto**](ModelClientCapabilitiesDto.md) | Gets or sets the client capabilities. | [optional] 
**RemoteEndPoint** | Pointer to **NullableString** | Gets or sets the remote end point. | [optional] 
**PlayableMediaTypes** | Pointer to [**[]ModelModelMediaType**](ModelModelMediaType.md) | Gets or sets the playable media types. | [optional] 
**Id** | Pointer to **NullableString** | Gets or sets the id. | [optional] 
**UserId** | Pointer to **string** | Gets or sets the user id. | [optional] 
**UserName** | Pointer to **NullableString** | Gets or sets the username. | [optional] 
**Client** | Pointer to **NullableString** | Gets or sets the type of the client. | [optional] 
**LastActivityDate** | Pointer to **time.Time** | Gets or sets the last activity date. | [optional] 
**LastPlaybackCheckIn** | Pointer to **time.Time** | Gets or sets the last playback check in. | [optional] 
**LastPausedDate** | Pointer to **NullableTime** | Gets or sets the last paused date. | [optional] 
**DeviceName** | Pointer to **NullableString** | Gets or sets the name of the device. | [optional] 
**DeviceType** | Pointer to **NullableString** | Gets or sets the type of the device. | [optional] 
**NowPlayingItem** | Pointer to [**NullableModelModelBaseItemDto**](ModelBaseItemDto.md) | Gets or sets the now playing item. | [optional] 
**NowViewingItem** | Pointer to [**NullableModelModelBaseItemDto**](ModelBaseItemDto.md) | Gets or sets the now viewing item. | [optional] 
**DeviceId** | Pointer to **NullableString** | Gets or sets the device id. | [optional] 
**ApplicationVersion** | Pointer to **NullableString** | Gets or sets the application version. | [optional] 
**TranscodingInfo** | Pointer to [**NullableModelModelTranscodingInfo**](ModelTranscodingInfo.md) | Gets or sets the transcoding info. | [optional] 
**IsActive** | Pointer to **bool** | Gets or sets a value indicating whether this session is active. | [optional] 
**SupportsMediaControl** | Pointer to **bool** | Gets or sets a value indicating whether the session supports media control. | [optional] 
**SupportsRemoteControl** | Pointer to **bool** | Gets or sets a value indicating whether the session supports remote control. | [optional] 
**NowPlayingQueue** | Pointer to [**[]ModelModelQueueItem**](ModelModelQueueItem.md) | Gets or sets the now playing queue. | [optional] 
**NowPlayingQueueFullItems** | Pointer to [**[]ModelModelBaseItemDto**](ModelModelBaseItemDto.md) | Gets or sets the now playing queue full items. | [optional] 
**HasCustomDeviceName** | Pointer to **bool** | Gets or sets a value indicating whether the session has a custom device name. | [optional] 
**PlaylistItemId** | Pointer to **NullableString** | Gets or sets the playlist item id. | [optional] 
**ServerId** | Pointer to **NullableString** | Gets or sets the server id. | [optional] 
**UserPrimaryImageTag** | Pointer to **NullableString** | Gets or sets the user primary image tag. | [optional] 
**SupportedCommands** | Pointer to [**[]ModelModelGeneralCommandType**](ModelModelGeneralCommandType.md) | Gets or sets the supported commands. | [optional] 

## Methods

### NewModelSessionInfoDto

`func NewModelSessionInfoDto() *ModelSessionInfoDto`

NewModelSessionInfoDto instantiates a new ModelSessionInfoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelSessionInfoDtoWithDefaults

`func NewModelSessionInfoDtoWithDefaults() *ModelSessionInfoDto`

NewModelSessionInfoDtoWithDefaults instantiates a new ModelSessionInfoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlayState

`func (o *ModelSessionInfoDto) GetPlayState() ModelModelPlayerStateInfo`

GetPlayState returns the PlayState field if non-nil, zero value otherwise.

### GetPlayStateOk

`func (o *ModelSessionInfoDto) GetPlayStateOk() (*ModelModelPlayerStateInfo, bool)`

GetPlayStateOk returns a tuple with the PlayState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayState

`func (o *ModelSessionInfoDto) SetPlayState(v ModelModelPlayerStateInfo)`

SetPlayState sets PlayState field to given value.

### HasPlayState

`func (o *ModelSessionInfoDto) HasPlayState() bool`

HasPlayState returns a boolean if a field has been set.

### SetPlayStateNil

`func (o *ModelSessionInfoDto) SetPlayStateNil(b bool)`

 SetPlayStateNil sets the value for PlayState to be an explicit nil

### UnsetPlayState
`func (o *ModelSessionInfoDto) UnsetPlayState()`

UnsetPlayState ensures that no value is present for PlayState, not even an explicit nil
### GetAdditionalUsers

`func (o *ModelSessionInfoDto) GetAdditionalUsers() []ModelModelSessionUserInfo`

GetAdditionalUsers returns the AdditionalUsers field if non-nil, zero value otherwise.

### GetAdditionalUsersOk

`func (o *ModelSessionInfoDto) GetAdditionalUsersOk() (*[]ModelModelSessionUserInfo, bool)`

GetAdditionalUsersOk returns a tuple with the AdditionalUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalUsers

`func (o *ModelSessionInfoDto) SetAdditionalUsers(v []ModelModelSessionUserInfo)`

SetAdditionalUsers sets AdditionalUsers field to given value.

### HasAdditionalUsers

`func (o *ModelSessionInfoDto) HasAdditionalUsers() bool`

HasAdditionalUsers returns a boolean if a field has been set.

### SetAdditionalUsersNil

`func (o *ModelSessionInfoDto) SetAdditionalUsersNil(b bool)`

 SetAdditionalUsersNil sets the value for AdditionalUsers to be an explicit nil

### UnsetAdditionalUsers
`func (o *ModelSessionInfoDto) UnsetAdditionalUsers()`

UnsetAdditionalUsers ensures that no value is present for AdditionalUsers, not even an explicit nil
### GetCapabilities

`func (o *ModelSessionInfoDto) GetCapabilities() ModelModelClientCapabilitiesDto`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *ModelSessionInfoDto) GetCapabilitiesOk() (*ModelModelClientCapabilitiesDto, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *ModelSessionInfoDto) SetCapabilities(v ModelModelClientCapabilitiesDto)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *ModelSessionInfoDto) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### SetCapabilitiesNil

`func (o *ModelSessionInfoDto) SetCapabilitiesNil(b bool)`

 SetCapabilitiesNil sets the value for Capabilities to be an explicit nil

### UnsetCapabilities
`func (o *ModelSessionInfoDto) UnsetCapabilities()`

UnsetCapabilities ensures that no value is present for Capabilities, not even an explicit nil
### GetRemoteEndPoint

`func (o *ModelSessionInfoDto) GetRemoteEndPoint() string`

GetRemoteEndPoint returns the RemoteEndPoint field if non-nil, zero value otherwise.

### GetRemoteEndPointOk

`func (o *ModelSessionInfoDto) GetRemoteEndPointOk() (*string, bool)`

GetRemoteEndPointOk returns a tuple with the RemoteEndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteEndPoint

`func (o *ModelSessionInfoDto) SetRemoteEndPoint(v string)`

SetRemoteEndPoint sets RemoteEndPoint field to given value.

### HasRemoteEndPoint

`func (o *ModelSessionInfoDto) HasRemoteEndPoint() bool`

HasRemoteEndPoint returns a boolean if a field has been set.

### SetRemoteEndPointNil

`func (o *ModelSessionInfoDto) SetRemoteEndPointNil(b bool)`

 SetRemoteEndPointNil sets the value for RemoteEndPoint to be an explicit nil

### UnsetRemoteEndPoint
`func (o *ModelSessionInfoDto) UnsetRemoteEndPoint()`

UnsetRemoteEndPoint ensures that no value is present for RemoteEndPoint, not even an explicit nil
### GetPlayableMediaTypes

`func (o *ModelSessionInfoDto) GetPlayableMediaTypes() []ModelModelMediaType`

GetPlayableMediaTypes returns the PlayableMediaTypes field if non-nil, zero value otherwise.

### GetPlayableMediaTypesOk

`func (o *ModelSessionInfoDto) GetPlayableMediaTypesOk() (*[]ModelModelMediaType, bool)`

GetPlayableMediaTypesOk returns a tuple with the PlayableMediaTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayableMediaTypes

`func (o *ModelSessionInfoDto) SetPlayableMediaTypes(v []ModelModelMediaType)`

SetPlayableMediaTypes sets PlayableMediaTypes field to given value.

### HasPlayableMediaTypes

`func (o *ModelSessionInfoDto) HasPlayableMediaTypes() bool`

HasPlayableMediaTypes returns a boolean if a field has been set.

### GetId

`func (o *ModelSessionInfoDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelSessionInfoDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelSessionInfoDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelSessionInfoDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ModelSessionInfoDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ModelSessionInfoDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetUserId

`func (o *ModelSessionInfoDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ModelSessionInfoDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ModelSessionInfoDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ModelSessionInfoDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserName

`func (o *ModelSessionInfoDto) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ModelSessionInfoDto) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ModelSessionInfoDto) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *ModelSessionInfoDto) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *ModelSessionInfoDto) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *ModelSessionInfoDto) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetClient

`func (o *ModelSessionInfoDto) GetClient() string`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *ModelSessionInfoDto) GetClientOk() (*string, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *ModelSessionInfoDto) SetClient(v string)`

SetClient sets Client field to given value.

### HasClient

`func (o *ModelSessionInfoDto) HasClient() bool`

HasClient returns a boolean if a field has been set.

### SetClientNil

`func (o *ModelSessionInfoDto) SetClientNil(b bool)`

 SetClientNil sets the value for Client to be an explicit nil

### UnsetClient
`func (o *ModelSessionInfoDto) UnsetClient()`

UnsetClient ensures that no value is present for Client, not even an explicit nil
### GetLastActivityDate

`func (o *ModelSessionInfoDto) GetLastActivityDate() time.Time`

GetLastActivityDate returns the LastActivityDate field if non-nil, zero value otherwise.

### GetLastActivityDateOk

`func (o *ModelSessionInfoDto) GetLastActivityDateOk() (*time.Time, bool)`

GetLastActivityDateOk returns a tuple with the LastActivityDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivityDate

`func (o *ModelSessionInfoDto) SetLastActivityDate(v time.Time)`

SetLastActivityDate sets LastActivityDate field to given value.

### HasLastActivityDate

`func (o *ModelSessionInfoDto) HasLastActivityDate() bool`

HasLastActivityDate returns a boolean if a field has been set.

### GetLastPlaybackCheckIn

`func (o *ModelSessionInfoDto) GetLastPlaybackCheckIn() time.Time`

GetLastPlaybackCheckIn returns the LastPlaybackCheckIn field if non-nil, zero value otherwise.

### GetLastPlaybackCheckInOk

`func (o *ModelSessionInfoDto) GetLastPlaybackCheckInOk() (*time.Time, bool)`

GetLastPlaybackCheckInOk returns a tuple with the LastPlaybackCheckIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPlaybackCheckIn

`func (o *ModelSessionInfoDto) SetLastPlaybackCheckIn(v time.Time)`

SetLastPlaybackCheckIn sets LastPlaybackCheckIn field to given value.

### HasLastPlaybackCheckIn

`func (o *ModelSessionInfoDto) HasLastPlaybackCheckIn() bool`

HasLastPlaybackCheckIn returns a boolean if a field has been set.

### GetLastPausedDate

`func (o *ModelSessionInfoDto) GetLastPausedDate() time.Time`

GetLastPausedDate returns the LastPausedDate field if non-nil, zero value otherwise.

### GetLastPausedDateOk

`func (o *ModelSessionInfoDto) GetLastPausedDateOk() (*time.Time, bool)`

GetLastPausedDateOk returns a tuple with the LastPausedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPausedDate

`func (o *ModelSessionInfoDto) SetLastPausedDate(v time.Time)`

SetLastPausedDate sets LastPausedDate field to given value.

### HasLastPausedDate

`func (o *ModelSessionInfoDto) HasLastPausedDate() bool`

HasLastPausedDate returns a boolean if a field has been set.

### SetLastPausedDateNil

`func (o *ModelSessionInfoDto) SetLastPausedDateNil(b bool)`

 SetLastPausedDateNil sets the value for LastPausedDate to be an explicit nil

### UnsetLastPausedDate
`func (o *ModelSessionInfoDto) UnsetLastPausedDate()`

UnsetLastPausedDate ensures that no value is present for LastPausedDate, not even an explicit nil
### GetDeviceName

`func (o *ModelSessionInfoDto) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *ModelSessionInfoDto) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *ModelSessionInfoDto) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *ModelSessionInfoDto) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### SetDeviceNameNil

`func (o *ModelSessionInfoDto) SetDeviceNameNil(b bool)`

 SetDeviceNameNil sets the value for DeviceName to be an explicit nil

### UnsetDeviceName
`func (o *ModelSessionInfoDto) UnsetDeviceName()`

UnsetDeviceName ensures that no value is present for DeviceName, not even an explicit nil
### GetDeviceType

`func (o *ModelSessionInfoDto) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *ModelSessionInfoDto) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *ModelSessionInfoDto) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *ModelSessionInfoDto) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### SetDeviceTypeNil

`func (o *ModelSessionInfoDto) SetDeviceTypeNil(b bool)`

 SetDeviceTypeNil sets the value for DeviceType to be an explicit nil

### UnsetDeviceType
`func (o *ModelSessionInfoDto) UnsetDeviceType()`

UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
### GetNowPlayingItem

`func (o *ModelSessionInfoDto) GetNowPlayingItem() ModelModelBaseItemDto`

GetNowPlayingItem returns the NowPlayingItem field if non-nil, zero value otherwise.

### GetNowPlayingItemOk

`func (o *ModelSessionInfoDto) GetNowPlayingItemOk() (*ModelModelBaseItemDto, bool)`

GetNowPlayingItemOk returns a tuple with the NowPlayingItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNowPlayingItem

`func (o *ModelSessionInfoDto) SetNowPlayingItem(v ModelModelBaseItemDto)`

SetNowPlayingItem sets NowPlayingItem field to given value.

### HasNowPlayingItem

`func (o *ModelSessionInfoDto) HasNowPlayingItem() bool`

HasNowPlayingItem returns a boolean if a field has been set.

### SetNowPlayingItemNil

`func (o *ModelSessionInfoDto) SetNowPlayingItemNil(b bool)`

 SetNowPlayingItemNil sets the value for NowPlayingItem to be an explicit nil

### UnsetNowPlayingItem
`func (o *ModelSessionInfoDto) UnsetNowPlayingItem()`

UnsetNowPlayingItem ensures that no value is present for NowPlayingItem, not even an explicit nil
### GetNowViewingItem

`func (o *ModelSessionInfoDto) GetNowViewingItem() ModelModelBaseItemDto`

GetNowViewingItem returns the NowViewingItem field if non-nil, zero value otherwise.

### GetNowViewingItemOk

`func (o *ModelSessionInfoDto) GetNowViewingItemOk() (*ModelModelBaseItemDto, bool)`

GetNowViewingItemOk returns a tuple with the NowViewingItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNowViewingItem

`func (o *ModelSessionInfoDto) SetNowViewingItem(v ModelModelBaseItemDto)`

SetNowViewingItem sets NowViewingItem field to given value.

### HasNowViewingItem

`func (o *ModelSessionInfoDto) HasNowViewingItem() bool`

HasNowViewingItem returns a boolean if a field has been set.

### SetNowViewingItemNil

`func (o *ModelSessionInfoDto) SetNowViewingItemNil(b bool)`

 SetNowViewingItemNil sets the value for NowViewingItem to be an explicit nil

### UnsetNowViewingItem
`func (o *ModelSessionInfoDto) UnsetNowViewingItem()`

UnsetNowViewingItem ensures that no value is present for NowViewingItem, not even an explicit nil
### GetDeviceId

`func (o *ModelSessionInfoDto) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *ModelSessionInfoDto) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *ModelSessionInfoDto) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *ModelSessionInfoDto) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### SetDeviceIdNil

`func (o *ModelSessionInfoDto) SetDeviceIdNil(b bool)`

 SetDeviceIdNil sets the value for DeviceId to be an explicit nil

### UnsetDeviceId
`func (o *ModelSessionInfoDto) UnsetDeviceId()`

UnsetDeviceId ensures that no value is present for DeviceId, not even an explicit nil
### GetApplicationVersion

`func (o *ModelSessionInfoDto) GetApplicationVersion() string`

GetApplicationVersion returns the ApplicationVersion field if non-nil, zero value otherwise.

### GetApplicationVersionOk

`func (o *ModelSessionInfoDto) GetApplicationVersionOk() (*string, bool)`

GetApplicationVersionOk returns a tuple with the ApplicationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationVersion

`func (o *ModelSessionInfoDto) SetApplicationVersion(v string)`

SetApplicationVersion sets ApplicationVersion field to given value.

### HasApplicationVersion

`func (o *ModelSessionInfoDto) HasApplicationVersion() bool`

HasApplicationVersion returns a boolean if a field has been set.

### SetApplicationVersionNil

`func (o *ModelSessionInfoDto) SetApplicationVersionNil(b bool)`

 SetApplicationVersionNil sets the value for ApplicationVersion to be an explicit nil

### UnsetApplicationVersion
`func (o *ModelSessionInfoDto) UnsetApplicationVersion()`

UnsetApplicationVersion ensures that no value is present for ApplicationVersion, not even an explicit nil
### GetTranscodingInfo

`func (o *ModelSessionInfoDto) GetTranscodingInfo() ModelModelTranscodingInfo`

GetTranscodingInfo returns the TranscodingInfo field if non-nil, zero value otherwise.

### GetTranscodingInfoOk

`func (o *ModelSessionInfoDto) GetTranscodingInfoOk() (*ModelModelTranscodingInfo, bool)`

GetTranscodingInfoOk returns a tuple with the TranscodingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodingInfo

`func (o *ModelSessionInfoDto) SetTranscodingInfo(v ModelModelTranscodingInfo)`

SetTranscodingInfo sets TranscodingInfo field to given value.

### HasTranscodingInfo

`func (o *ModelSessionInfoDto) HasTranscodingInfo() bool`

HasTranscodingInfo returns a boolean if a field has been set.

### SetTranscodingInfoNil

`func (o *ModelSessionInfoDto) SetTranscodingInfoNil(b bool)`

 SetTranscodingInfoNil sets the value for TranscodingInfo to be an explicit nil

### UnsetTranscodingInfo
`func (o *ModelSessionInfoDto) UnsetTranscodingInfo()`

UnsetTranscodingInfo ensures that no value is present for TranscodingInfo, not even an explicit nil
### GetIsActive

`func (o *ModelSessionInfoDto) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ModelSessionInfoDto) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ModelSessionInfoDto) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *ModelSessionInfoDto) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetSupportsMediaControl

`func (o *ModelSessionInfoDto) GetSupportsMediaControl() bool`

GetSupportsMediaControl returns the SupportsMediaControl field if non-nil, zero value otherwise.

### GetSupportsMediaControlOk

`func (o *ModelSessionInfoDto) GetSupportsMediaControlOk() (*bool, bool)`

GetSupportsMediaControlOk returns a tuple with the SupportsMediaControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsMediaControl

`func (o *ModelSessionInfoDto) SetSupportsMediaControl(v bool)`

SetSupportsMediaControl sets SupportsMediaControl field to given value.

### HasSupportsMediaControl

`func (o *ModelSessionInfoDto) HasSupportsMediaControl() bool`

HasSupportsMediaControl returns a boolean if a field has been set.

### GetSupportsRemoteControl

`func (o *ModelSessionInfoDto) GetSupportsRemoteControl() bool`

GetSupportsRemoteControl returns the SupportsRemoteControl field if non-nil, zero value otherwise.

### GetSupportsRemoteControlOk

`func (o *ModelSessionInfoDto) GetSupportsRemoteControlOk() (*bool, bool)`

GetSupportsRemoteControlOk returns a tuple with the SupportsRemoteControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsRemoteControl

`func (o *ModelSessionInfoDto) SetSupportsRemoteControl(v bool)`

SetSupportsRemoteControl sets SupportsRemoteControl field to given value.

### HasSupportsRemoteControl

`func (o *ModelSessionInfoDto) HasSupportsRemoteControl() bool`

HasSupportsRemoteControl returns a boolean if a field has been set.

### GetNowPlayingQueue

`func (o *ModelSessionInfoDto) GetNowPlayingQueue() []ModelModelQueueItem`

GetNowPlayingQueue returns the NowPlayingQueue field if non-nil, zero value otherwise.

### GetNowPlayingQueueOk

`func (o *ModelSessionInfoDto) GetNowPlayingQueueOk() (*[]ModelModelQueueItem, bool)`

GetNowPlayingQueueOk returns a tuple with the NowPlayingQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNowPlayingQueue

`func (o *ModelSessionInfoDto) SetNowPlayingQueue(v []ModelModelQueueItem)`

SetNowPlayingQueue sets NowPlayingQueue field to given value.

### HasNowPlayingQueue

`func (o *ModelSessionInfoDto) HasNowPlayingQueue() bool`

HasNowPlayingQueue returns a boolean if a field has been set.

### SetNowPlayingQueueNil

`func (o *ModelSessionInfoDto) SetNowPlayingQueueNil(b bool)`

 SetNowPlayingQueueNil sets the value for NowPlayingQueue to be an explicit nil

### UnsetNowPlayingQueue
`func (o *ModelSessionInfoDto) UnsetNowPlayingQueue()`

UnsetNowPlayingQueue ensures that no value is present for NowPlayingQueue, not even an explicit nil
### GetNowPlayingQueueFullItems

`func (o *ModelSessionInfoDto) GetNowPlayingQueueFullItems() []ModelModelBaseItemDto`

GetNowPlayingQueueFullItems returns the NowPlayingQueueFullItems field if non-nil, zero value otherwise.

### GetNowPlayingQueueFullItemsOk

`func (o *ModelSessionInfoDto) GetNowPlayingQueueFullItemsOk() (*[]ModelModelBaseItemDto, bool)`

GetNowPlayingQueueFullItemsOk returns a tuple with the NowPlayingQueueFullItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNowPlayingQueueFullItems

`func (o *ModelSessionInfoDto) SetNowPlayingQueueFullItems(v []ModelModelBaseItemDto)`

SetNowPlayingQueueFullItems sets NowPlayingQueueFullItems field to given value.

### HasNowPlayingQueueFullItems

`func (o *ModelSessionInfoDto) HasNowPlayingQueueFullItems() bool`

HasNowPlayingQueueFullItems returns a boolean if a field has been set.

### SetNowPlayingQueueFullItemsNil

`func (o *ModelSessionInfoDto) SetNowPlayingQueueFullItemsNil(b bool)`

 SetNowPlayingQueueFullItemsNil sets the value for NowPlayingQueueFullItems to be an explicit nil

### UnsetNowPlayingQueueFullItems
`func (o *ModelSessionInfoDto) UnsetNowPlayingQueueFullItems()`

UnsetNowPlayingQueueFullItems ensures that no value is present for NowPlayingQueueFullItems, not even an explicit nil
### GetHasCustomDeviceName

`func (o *ModelSessionInfoDto) GetHasCustomDeviceName() bool`

GetHasCustomDeviceName returns the HasCustomDeviceName field if non-nil, zero value otherwise.

### GetHasCustomDeviceNameOk

`func (o *ModelSessionInfoDto) GetHasCustomDeviceNameOk() (*bool, bool)`

GetHasCustomDeviceNameOk returns a tuple with the HasCustomDeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCustomDeviceName

`func (o *ModelSessionInfoDto) SetHasCustomDeviceName(v bool)`

SetHasCustomDeviceName sets HasCustomDeviceName field to given value.

### HasHasCustomDeviceName

`func (o *ModelSessionInfoDto) HasHasCustomDeviceName() bool`

HasHasCustomDeviceName returns a boolean if a field has been set.

### GetPlaylistItemId

`func (o *ModelSessionInfoDto) GetPlaylistItemId() string`

GetPlaylistItemId returns the PlaylistItemId field if non-nil, zero value otherwise.

### GetPlaylistItemIdOk

`func (o *ModelSessionInfoDto) GetPlaylistItemIdOk() (*string, bool)`

GetPlaylistItemIdOk returns a tuple with the PlaylistItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistItemId

`func (o *ModelSessionInfoDto) SetPlaylistItemId(v string)`

SetPlaylistItemId sets PlaylistItemId field to given value.

### HasPlaylistItemId

`func (o *ModelSessionInfoDto) HasPlaylistItemId() bool`

HasPlaylistItemId returns a boolean if a field has been set.

### SetPlaylistItemIdNil

`func (o *ModelSessionInfoDto) SetPlaylistItemIdNil(b bool)`

 SetPlaylistItemIdNil sets the value for PlaylistItemId to be an explicit nil

### UnsetPlaylistItemId
`func (o *ModelSessionInfoDto) UnsetPlaylistItemId()`

UnsetPlaylistItemId ensures that no value is present for PlaylistItemId, not even an explicit nil
### GetServerId

`func (o *ModelSessionInfoDto) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *ModelSessionInfoDto) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *ModelSessionInfoDto) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *ModelSessionInfoDto) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### SetServerIdNil

`func (o *ModelSessionInfoDto) SetServerIdNil(b bool)`

 SetServerIdNil sets the value for ServerId to be an explicit nil

### UnsetServerId
`func (o *ModelSessionInfoDto) UnsetServerId()`

UnsetServerId ensures that no value is present for ServerId, not even an explicit nil
### GetUserPrimaryImageTag

`func (o *ModelSessionInfoDto) GetUserPrimaryImageTag() string`

GetUserPrimaryImageTag returns the UserPrimaryImageTag field if non-nil, zero value otherwise.

### GetUserPrimaryImageTagOk

`func (o *ModelSessionInfoDto) GetUserPrimaryImageTagOk() (*string, bool)`

GetUserPrimaryImageTagOk returns a tuple with the UserPrimaryImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrimaryImageTag

`func (o *ModelSessionInfoDto) SetUserPrimaryImageTag(v string)`

SetUserPrimaryImageTag sets UserPrimaryImageTag field to given value.

### HasUserPrimaryImageTag

`func (o *ModelSessionInfoDto) HasUserPrimaryImageTag() bool`

HasUserPrimaryImageTag returns a boolean if a field has been set.

### SetUserPrimaryImageTagNil

`func (o *ModelSessionInfoDto) SetUserPrimaryImageTagNil(b bool)`

 SetUserPrimaryImageTagNil sets the value for UserPrimaryImageTag to be an explicit nil

### UnsetUserPrimaryImageTag
`func (o *ModelSessionInfoDto) UnsetUserPrimaryImageTag()`

UnsetUserPrimaryImageTag ensures that no value is present for UserPrimaryImageTag, not even an explicit nil
### GetSupportedCommands

`func (o *ModelSessionInfoDto) GetSupportedCommands() []ModelModelGeneralCommandType`

GetSupportedCommands returns the SupportedCommands field if non-nil, zero value otherwise.

### GetSupportedCommandsOk

`func (o *ModelSessionInfoDto) GetSupportedCommandsOk() (*[]ModelModelGeneralCommandType, bool)`

GetSupportedCommandsOk returns a tuple with the SupportedCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedCommands

`func (o *ModelSessionInfoDto) SetSupportedCommands(v []ModelModelGeneralCommandType)`

SetSupportedCommands sets SupportedCommands field to given value.

### HasSupportedCommands

`func (o *ModelSessionInfoDto) HasSupportedCommands() bool`

HasSupportedCommands returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


