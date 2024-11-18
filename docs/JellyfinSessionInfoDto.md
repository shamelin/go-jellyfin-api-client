# JellyfinSessionInfoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlayState** | Pointer to [**NullableJellyfinJellyfinPlayerStateInfo**](JellyfinPlayerStateInfo.md) | Gets or sets the play state. | [optional] 
**AdditionalUsers** | Pointer to [**[]JellyfinJellyfinSessionUserInfo**](JellyfinJellyfinSessionUserInfo.md) | Gets or sets the additional users. | [optional] 
**Capabilities** | Pointer to [**NullableJellyfinJellyfinClientCapabilitiesDto**](JellyfinClientCapabilitiesDto.md) | Gets or sets the client capabilities. | [optional] 
**RemoteEndPoint** | Pointer to **NullableString** | Gets or sets the remote end point. | [optional] 
**PlayableMediaTypes** | Pointer to [**[]JellyfinJellyfinMediaType**](JellyfinJellyfinMediaType.md) | Gets or sets the playable media types. | [optional] 
**Id** | Pointer to **NullableString** | Gets or sets the id. | [optional] 
**UserId** | Pointer to **string** | Gets or sets the user id. | [optional] 
**UserName** | Pointer to **NullableString** | Gets or sets the username. | [optional] 
**Client** | Pointer to **NullableString** | Gets or sets the type of the client. | [optional] 
**LastActivityDate** | Pointer to **time.Time** | Gets or sets the last activity date. | [optional] 
**LastPlaybackCheckIn** | Pointer to **time.Time** | Gets or sets the last playback check in. | [optional] 
**LastPausedDate** | Pointer to **NullableTime** | Gets or sets the last paused date. | [optional] 
**DeviceName** | Pointer to **NullableString** | Gets or sets the name of the device. | [optional] 
**DeviceType** | Pointer to **NullableString** | Gets or sets the type of the device. | [optional] 
**NowPlayingItem** | Pointer to [**NullableJellyfinJellyfinBaseItemDto**](JellyfinBaseItemDto.md) | Gets or sets the now playing item. | [optional] 
**NowViewingItem** | Pointer to [**NullableJellyfinJellyfinBaseItemDto**](JellyfinBaseItemDto.md) | Gets or sets the now viewing item. | [optional] 
**DeviceId** | Pointer to **NullableString** | Gets or sets the device id. | [optional] 
**ApplicationVersion** | Pointer to **NullableString** | Gets or sets the application version. | [optional] 
**TranscodingInfo** | Pointer to [**NullableJellyfinJellyfinTranscodingInfo**](JellyfinTranscodingInfo.md) | Gets or sets the transcoding info. | [optional] 
**IsActive** | Pointer to **bool** | Gets or sets a value indicating whether this session is active. | [optional] 
**SupportsMediaControl** | Pointer to **bool** | Gets or sets a value indicating whether the session supports media control. | [optional] 
**SupportsRemoteControl** | Pointer to **bool** | Gets or sets a value indicating whether the session supports remote control. | [optional] 
**NowPlayingQueue** | Pointer to [**[]JellyfinJellyfinQueueItem**](JellyfinJellyfinQueueItem.md) | Gets or sets the now playing queue. | [optional] 
**NowPlayingQueueFullItems** | Pointer to [**[]JellyfinJellyfinBaseItemDto**](JellyfinJellyfinBaseItemDto.md) | Gets or sets the now playing queue full items. | [optional] 
**HasCustomDeviceName** | Pointer to **bool** | Gets or sets a value indicating whether the session has a custom device name. | [optional] 
**PlaylistItemId** | Pointer to **NullableString** | Gets or sets the playlist item id. | [optional] 
**ServerId** | Pointer to **NullableString** | Gets or sets the server id. | [optional] 
**UserPrimaryImageTag** | Pointer to **NullableString** | Gets or sets the user primary image tag. | [optional] 
**SupportedCommands** | Pointer to [**[]JellyfinJellyfinGeneralCommandType**](JellyfinJellyfinGeneralCommandType.md) | Gets or sets the supported commands. | [optional] 

## Methods

### NewJellyfinSessionInfoDto

`func NewJellyfinSessionInfoDto() *JellyfinSessionInfoDto`

NewJellyfinSessionInfoDto instantiates a new JellyfinSessionInfoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinSessionInfoDtoWithDefaults

`func NewJellyfinSessionInfoDtoWithDefaults() *JellyfinSessionInfoDto`

NewJellyfinSessionInfoDtoWithDefaults instantiates a new JellyfinSessionInfoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlayState

`func (o *JellyfinSessionInfoDto) GetPlayState() JellyfinJellyfinPlayerStateInfo`

GetPlayState returns the PlayState field if non-nil, zero value otherwise.

### GetPlayStateOk

`func (o *JellyfinSessionInfoDto) GetPlayStateOk() (*JellyfinJellyfinPlayerStateInfo, bool)`

GetPlayStateOk returns a tuple with the PlayState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayState

`func (o *JellyfinSessionInfoDto) SetPlayState(v JellyfinJellyfinPlayerStateInfo)`

SetPlayState sets PlayState field to given value.

### HasPlayState

`func (o *JellyfinSessionInfoDto) HasPlayState() bool`

HasPlayState returns a boolean if a field has been set.

### SetPlayStateNil

`func (o *JellyfinSessionInfoDto) SetPlayStateNil(b bool)`

 SetPlayStateNil sets the value for PlayState to be an explicit nil

### UnsetPlayState
`func (o *JellyfinSessionInfoDto) UnsetPlayState()`

UnsetPlayState ensures that no value is present for PlayState, not even an explicit nil
### GetAdditionalUsers

`func (o *JellyfinSessionInfoDto) GetAdditionalUsers() []JellyfinJellyfinSessionUserInfo`

GetAdditionalUsers returns the AdditionalUsers field if non-nil, zero value otherwise.

### GetAdditionalUsersOk

`func (o *JellyfinSessionInfoDto) GetAdditionalUsersOk() (*[]JellyfinJellyfinSessionUserInfo, bool)`

GetAdditionalUsersOk returns a tuple with the AdditionalUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalUsers

`func (o *JellyfinSessionInfoDto) SetAdditionalUsers(v []JellyfinJellyfinSessionUserInfo)`

SetAdditionalUsers sets AdditionalUsers field to given value.

### HasAdditionalUsers

`func (o *JellyfinSessionInfoDto) HasAdditionalUsers() bool`

HasAdditionalUsers returns a boolean if a field has been set.

### SetAdditionalUsersNil

`func (o *JellyfinSessionInfoDto) SetAdditionalUsersNil(b bool)`

 SetAdditionalUsersNil sets the value for AdditionalUsers to be an explicit nil

### UnsetAdditionalUsers
`func (o *JellyfinSessionInfoDto) UnsetAdditionalUsers()`

UnsetAdditionalUsers ensures that no value is present for AdditionalUsers, not even an explicit nil
### GetCapabilities

`func (o *JellyfinSessionInfoDto) GetCapabilities() JellyfinJellyfinClientCapabilitiesDto`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *JellyfinSessionInfoDto) GetCapabilitiesOk() (*JellyfinJellyfinClientCapabilitiesDto, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *JellyfinSessionInfoDto) SetCapabilities(v JellyfinJellyfinClientCapabilitiesDto)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *JellyfinSessionInfoDto) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### SetCapabilitiesNil

`func (o *JellyfinSessionInfoDto) SetCapabilitiesNil(b bool)`

 SetCapabilitiesNil sets the value for Capabilities to be an explicit nil

### UnsetCapabilities
`func (o *JellyfinSessionInfoDto) UnsetCapabilities()`

UnsetCapabilities ensures that no value is present for Capabilities, not even an explicit nil
### GetRemoteEndPoint

`func (o *JellyfinSessionInfoDto) GetRemoteEndPoint() string`

GetRemoteEndPoint returns the RemoteEndPoint field if non-nil, zero value otherwise.

### GetRemoteEndPointOk

`func (o *JellyfinSessionInfoDto) GetRemoteEndPointOk() (*string, bool)`

GetRemoteEndPointOk returns a tuple with the RemoteEndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteEndPoint

`func (o *JellyfinSessionInfoDto) SetRemoteEndPoint(v string)`

SetRemoteEndPoint sets RemoteEndPoint field to given value.

### HasRemoteEndPoint

`func (o *JellyfinSessionInfoDto) HasRemoteEndPoint() bool`

HasRemoteEndPoint returns a boolean if a field has been set.

### SetRemoteEndPointNil

`func (o *JellyfinSessionInfoDto) SetRemoteEndPointNil(b bool)`

 SetRemoteEndPointNil sets the value for RemoteEndPoint to be an explicit nil

### UnsetRemoteEndPoint
`func (o *JellyfinSessionInfoDto) UnsetRemoteEndPoint()`

UnsetRemoteEndPoint ensures that no value is present for RemoteEndPoint, not even an explicit nil
### GetPlayableMediaTypes

`func (o *JellyfinSessionInfoDto) GetPlayableMediaTypes() []JellyfinJellyfinMediaType`

GetPlayableMediaTypes returns the PlayableMediaTypes field if non-nil, zero value otherwise.

### GetPlayableMediaTypesOk

`func (o *JellyfinSessionInfoDto) GetPlayableMediaTypesOk() (*[]JellyfinJellyfinMediaType, bool)`

GetPlayableMediaTypesOk returns a tuple with the PlayableMediaTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayableMediaTypes

`func (o *JellyfinSessionInfoDto) SetPlayableMediaTypes(v []JellyfinJellyfinMediaType)`

SetPlayableMediaTypes sets PlayableMediaTypes field to given value.

### HasPlayableMediaTypes

`func (o *JellyfinSessionInfoDto) HasPlayableMediaTypes() bool`

HasPlayableMediaTypes returns a boolean if a field has been set.

### GetId

`func (o *JellyfinSessionInfoDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JellyfinSessionInfoDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JellyfinSessionInfoDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JellyfinSessionInfoDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *JellyfinSessionInfoDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *JellyfinSessionInfoDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetUserId

`func (o *JellyfinSessionInfoDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *JellyfinSessionInfoDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *JellyfinSessionInfoDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *JellyfinSessionInfoDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserName

`func (o *JellyfinSessionInfoDto) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *JellyfinSessionInfoDto) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *JellyfinSessionInfoDto) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *JellyfinSessionInfoDto) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *JellyfinSessionInfoDto) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *JellyfinSessionInfoDto) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetClient

`func (o *JellyfinSessionInfoDto) GetClient() string`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *JellyfinSessionInfoDto) GetClientOk() (*string, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *JellyfinSessionInfoDto) SetClient(v string)`

SetClient sets Client field to given value.

### HasClient

`func (o *JellyfinSessionInfoDto) HasClient() bool`

HasClient returns a boolean if a field has been set.

### SetClientNil

`func (o *JellyfinSessionInfoDto) SetClientNil(b bool)`

 SetClientNil sets the value for Client to be an explicit nil

### UnsetClient
`func (o *JellyfinSessionInfoDto) UnsetClient()`

UnsetClient ensures that no value is present for Client, not even an explicit nil
### GetLastActivityDate

`func (o *JellyfinSessionInfoDto) GetLastActivityDate() time.Time`

GetLastActivityDate returns the LastActivityDate field if non-nil, zero value otherwise.

### GetLastActivityDateOk

`func (o *JellyfinSessionInfoDto) GetLastActivityDateOk() (*time.Time, bool)`

GetLastActivityDateOk returns a tuple with the LastActivityDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivityDate

`func (o *JellyfinSessionInfoDto) SetLastActivityDate(v time.Time)`

SetLastActivityDate sets LastActivityDate field to given value.

### HasLastActivityDate

`func (o *JellyfinSessionInfoDto) HasLastActivityDate() bool`

HasLastActivityDate returns a boolean if a field has been set.

### GetLastPlaybackCheckIn

`func (o *JellyfinSessionInfoDto) GetLastPlaybackCheckIn() time.Time`

GetLastPlaybackCheckIn returns the LastPlaybackCheckIn field if non-nil, zero value otherwise.

### GetLastPlaybackCheckInOk

`func (o *JellyfinSessionInfoDto) GetLastPlaybackCheckInOk() (*time.Time, bool)`

GetLastPlaybackCheckInOk returns a tuple with the LastPlaybackCheckIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPlaybackCheckIn

`func (o *JellyfinSessionInfoDto) SetLastPlaybackCheckIn(v time.Time)`

SetLastPlaybackCheckIn sets LastPlaybackCheckIn field to given value.

### HasLastPlaybackCheckIn

`func (o *JellyfinSessionInfoDto) HasLastPlaybackCheckIn() bool`

HasLastPlaybackCheckIn returns a boolean if a field has been set.

### GetLastPausedDate

`func (o *JellyfinSessionInfoDto) GetLastPausedDate() time.Time`

GetLastPausedDate returns the LastPausedDate field if non-nil, zero value otherwise.

### GetLastPausedDateOk

`func (o *JellyfinSessionInfoDto) GetLastPausedDateOk() (*time.Time, bool)`

GetLastPausedDateOk returns a tuple with the LastPausedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPausedDate

`func (o *JellyfinSessionInfoDto) SetLastPausedDate(v time.Time)`

SetLastPausedDate sets LastPausedDate field to given value.

### HasLastPausedDate

`func (o *JellyfinSessionInfoDto) HasLastPausedDate() bool`

HasLastPausedDate returns a boolean if a field has been set.

### SetLastPausedDateNil

`func (o *JellyfinSessionInfoDto) SetLastPausedDateNil(b bool)`

 SetLastPausedDateNil sets the value for LastPausedDate to be an explicit nil

### UnsetLastPausedDate
`func (o *JellyfinSessionInfoDto) UnsetLastPausedDate()`

UnsetLastPausedDate ensures that no value is present for LastPausedDate, not even an explicit nil
### GetDeviceName

`func (o *JellyfinSessionInfoDto) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *JellyfinSessionInfoDto) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *JellyfinSessionInfoDto) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *JellyfinSessionInfoDto) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### SetDeviceNameNil

`func (o *JellyfinSessionInfoDto) SetDeviceNameNil(b bool)`

 SetDeviceNameNil sets the value for DeviceName to be an explicit nil

### UnsetDeviceName
`func (o *JellyfinSessionInfoDto) UnsetDeviceName()`

UnsetDeviceName ensures that no value is present for DeviceName, not even an explicit nil
### GetDeviceType

`func (o *JellyfinSessionInfoDto) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *JellyfinSessionInfoDto) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *JellyfinSessionInfoDto) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *JellyfinSessionInfoDto) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### SetDeviceTypeNil

`func (o *JellyfinSessionInfoDto) SetDeviceTypeNil(b bool)`

 SetDeviceTypeNil sets the value for DeviceType to be an explicit nil

### UnsetDeviceType
`func (o *JellyfinSessionInfoDto) UnsetDeviceType()`

UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
### GetNowPlayingItem

`func (o *JellyfinSessionInfoDto) GetNowPlayingItem() JellyfinJellyfinBaseItemDto`

GetNowPlayingItem returns the NowPlayingItem field if non-nil, zero value otherwise.

### GetNowPlayingItemOk

`func (o *JellyfinSessionInfoDto) GetNowPlayingItemOk() (*JellyfinJellyfinBaseItemDto, bool)`

GetNowPlayingItemOk returns a tuple with the NowPlayingItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNowPlayingItem

`func (o *JellyfinSessionInfoDto) SetNowPlayingItem(v JellyfinJellyfinBaseItemDto)`

SetNowPlayingItem sets NowPlayingItem field to given value.

### HasNowPlayingItem

`func (o *JellyfinSessionInfoDto) HasNowPlayingItem() bool`

HasNowPlayingItem returns a boolean if a field has been set.

### SetNowPlayingItemNil

`func (o *JellyfinSessionInfoDto) SetNowPlayingItemNil(b bool)`

 SetNowPlayingItemNil sets the value for NowPlayingItem to be an explicit nil

### UnsetNowPlayingItem
`func (o *JellyfinSessionInfoDto) UnsetNowPlayingItem()`

UnsetNowPlayingItem ensures that no value is present for NowPlayingItem, not even an explicit nil
### GetNowViewingItem

`func (o *JellyfinSessionInfoDto) GetNowViewingItem() JellyfinJellyfinBaseItemDto`

GetNowViewingItem returns the NowViewingItem field if non-nil, zero value otherwise.

### GetNowViewingItemOk

`func (o *JellyfinSessionInfoDto) GetNowViewingItemOk() (*JellyfinJellyfinBaseItemDto, bool)`

GetNowViewingItemOk returns a tuple with the NowViewingItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNowViewingItem

`func (o *JellyfinSessionInfoDto) SetNowViewingItem(v JellyfinJellyfinBaseItemDto)`

SetNowViewingItem sets NowViewingItem field to given value.

### HasNowViewingItem

`func (o *JellyfinSessionInfoDto) HasNowViewingItem() bool`

HasNowViewingItem returns a boolean if a field has been set.

### SetNowViewingItemNil

`func (o *JellyfinSessionInfoDto) SetNowViewingItemNil(b bool)`

 SetNowViewingItemNil sets the value for NowViewingItem to be an explicit nil

### UnsetNowViewingItem
`func (o *JellyfinSessionInfoDto) UnsetNowViewingItem()`

UnsetNowViewingItem ensures that no value is present for NowViewingItem, not even an explicit nil
### GetDeviceId

`func (o *JellyfinSessionInfoDto) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *JellyfinSessionInfoDto) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *JellyfinSessionInfoDto) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *JellyfinSessionInfoDto) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### SetDeviceIdNil

`func (o *JellyfinSessionInfoDto) SetDeviceIdNil(b bool)`

 SetDeviceIdNil sets the value for DeviceId to be an explicit nil

### UnsetDeviceId
`func (o *JellyfinSessionInfoDto) UnsetDeviceId()`

UnsetDeviceId ensures that no value is present for DeviceId, not even an explicit nil
### GetApplicationVersion

`func (o *JellyfinSessionInfoDto) GetApplicationVersion() string`

GetApplicationVersion returns the ApplicationVersion field if non-nil, zero value otherwise.

### GetApplicationVersionOk

`func (o *JellyfinSessionInfoDto) GetApplicationVersionOk() (*string, bool)`

GetApplicationVersionOk returns a tuple with the ApplicationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationVersion

`func (o *JellyfinSessionInfoDto) SetApplicationVersion(v string)`

SetApplicationVersion sets ApplicationVersion field to given value.

### HasApplicationVersion

`func (o *JellyfinSessionInfoDto) HasApplicationVersion() bool`

HasApplicationVersion returns a boolean if a field has been set.

### SetApplicationVersionNil

`func (o *JellyfinSessionInfoDto) SetApplicationVersionNil(b bool)`

 SetApplicationVersionNil sets the value for ApplicationVersion to be an explicit nil

### UnsetApplicationVersion
`func (o *JellyfinSessionInfoDto) UnsetApplicationVersion()`

UnsetApplicationVersion ensures that no value is present for ApplicationVersion, not even an explicit nil
### GetTranscodingInfo

`func (o *JellyfinSessionInfoDto) GetTranscodingInfo() JellyfinJellyfinTranscodingInfo`

GetTranscodingInfo returns the TranscodingInfo field if non-nil, zero value otherwise.

### GetTranscodingInfoOk

`func (o *JellyfinSessionInfoDto) GetTranscodingInfoOk() (*JellyfinJellyfinTranscodingInfo, bool)`

GetTranscodingInfoOk returns a tuple with the TranscodingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodingInfo

`func (o *JellyfinSessionInfoDto) SetTranscodingInfo(v JellyfinJellyfinTranscodingInfo)`

SetTranscodingInfo sets TranscodingInfo field to given value.

### HasTranscodingInfo

`func (o *JellyfinSessionInfoDto) HasTranscodingInfo() bool`

HasTranscodingInfo returns a boolean if a field has been set.

### SetTranscodingInfoNil

`func (o *JellyfinSessionInfoDto) SetTranscodingInfoNil(b bool)`

 SetTranscodingInfoNil sets the value for TranscodingInfo to be an explicit nil

### UnsetTranscodingInfo
`func (o *JellyfinSessionInfoDto) UnsetTranscodingInfo()`

UnsetTranscodingInfo ensures that no value is present for TranscodingInfo, not even an explicit nil
### GetIsActive

`func (o *JellyfinSessionInfoDto) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *JellyfinSessionInfoDto) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *JellyfinSessionInfoDto) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *JellyfinSessionInfoDto) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetSupportsMediaControl

`func (o *JellyfinSessionInfoDto) GetSupportsMediaControl() bool`

GetSupportsMediaControl returns the SupportsMediaControl field if non-nil, zero value otherwise.

### GetSupportsMediaControlOk

`func (o *JellyfinSessionInfoDto) GetSupportsMediaControlOk() (*bool, bool)`

GetSupportsMediaControlOk returns a tuple with the SupportsMediaControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsMediaControl

`func (o *JellyfinSessionInfoDto) SetSupportsMediaControl(v bool)`

SetSupportsMediaControl sets SupportsMediaControl field to given value.

### HasSupportsMediaControl

`func (o *JellyfinSessionInfoDto) HasSupportsMediaControl() bool`

HasSupportsMediaControl returns a boolean if a field has been set.

### GetSupportsRemoteControl

`func (o *JellyfinSessionInfoDto) GetSupportsRemoteControl() bool`

GetSupportsRemoteControl returns the SupportsRemoteControl field if non-nil, zero value otherwise.

### GetSupportsRemoteControlOk

`func (o *JellyfinSessionInfoDto) GetSupportsRemoteControlOk() (*bool, bool)`

GetSupportsRemoteControlOk returns a tuple with the SupportsRemoteControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsRemoteControl

`func (o *JellyfinSessionInfoDto) SetSupportsRemoteControl(v bool)`

SetSupportsRemoteControl sets SupportsRemoteControl field to given value.

### HasSupportsRemoteControl

`func (o *JellyfinSessionInfoDto) HasSupportsRemoteControl() bool`

HasSupportsRemoteControl returns a boolean if a field has been set.

### GetNowPlayingQueue

`func (o *JellyfinSessionInfoDto) GetNowPlayingQueue() []JellyfinJellyfinQueueItem`

GetNowPlayingQueue returns the NowPlayingQueue field if non-nil, zero value otherwise.

### GetNowPlayingQueueOk

`func (o *JellyfinSessionInfoDto) GetNowPlayingQueueOk() (*[]JellyfinJellyfinQueueItem, bool)`

GetNowPlayingQueueOk returns a tuple with the NowPlayingQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNowPlayingQueue

`func (o *JellyfinSessionInfoDto) SetNowPlayingQueue(v []JellyfinJellyfinQueueItem)`

SetNowPlayingQueue sets NowPlayingQueue field to given value.

### HasNowPlayingQueue

`func (o *JellyfinSessionInfoDto) HasNowPlayingQueue() bool`

HasNowPlayingQueue returns a boolean if a field has been set.

### SetNowPlayingQueueNil

`func (o *JellyfinSessionInfoDto) SetNowPlayingQueueNil(b bool)`

 SetNowPlayingQueueNil sets the value for NowPlayingQueue to be an explicit nil

### UnsetNowPlayingQueue
`func (o *JellyfinSessionInfoDto) UnsetNowPlayingQueue()`

UnsetNowPlayingQueue ensures that no value is present for NowPlayingQueue, not even an explicit nil
### GetNowPlayingQueueFullItems

`func (o *JellyfinSessionInfoDto) GetNowPlayingQueueFullItems() []JellyfinJellyfinBaseItemDto`

GetNowPlayingQueueFullItems returns the NowPlayingQueueFullItems field if non-nil, zero value otherwise.

### GetNowPlayingQueueFullItemsOk

`func (o *JellyfinSessionInfoDto) GetNowPlayingQueueFullItemsOk() (*[]JellyfinJellyfinBaseItemDto, bool)`

GetNowPlayingQueueFullItemsOk returns a tuple with the NowPlayingQueueFullItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNowPlayingQueueFullItems

`func (o *JellyfinSessionInfoDto) SetNowPlayingQueueFullItems(v []JellyfinJellyfinBaseItemDto)`

SetNowPlayingQueueFullItems sets NowPlayingQueueFullItems field to given value.

### HasNowPlayingQueueFullItems

`func (o *JellyfinSessionInfoDto) HasNowPlayingQueueFullItems() bool`

HasNowPlayingQueueFullItems returns a boolean if a field has been set.

### SetNowPlayingQueueFullItemsNil

`func (o *JellyfinSessionInfoDto) SetNowPlayingQueueFullItemsNil(b bool)`

 SetNowPlayingQueueFullItemsNil sets the value for NowPlayingQueueFullItems to be an explicit nil

### UnsetNowPlayingQueueFullItems
`func (o *JellyfinSessionInfoDto) UnsetNowPlayingQueueFullItems()`

UnsetNowPlayingQueueFullItems ensures that no value is present for NowPlayingQueueFullItems, not even an explicit nil
### GetHasCustomDeviceName

`func (o *JellyfinSessionInfoDto) GetHasCustomDeviceName() bool`

GetHasCustomDeviceName returns the HasCustomDeviceName field if non-nil, zero value otherwise.

### GetHasCustomDeviceNameOk

`func (o *JellyfinSessionInfoDto) GetHasCustomDeviceNameOk() (*bool, bool)`

GetHasCustomDeviceNameOk returns a tuple with the HasCustomDeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCustomDeviceName

`func (o *JellyfinSessionInfoDto) SetHasCustomDeviceName(v bool)`

SetHasCustomDeviceName sets HasCustomDeviceName field to given value.

### HasHasCustomDeviceName

`func (o *JellyfinSessionInfoDto) HasHasCustomDeviceName() bool`

HasHasCustomDeviceName returns a boolean if a field has been set.

### GetPlaylistItemId

`func (o *JellyfinSessionInfoDto) GetPlaylistItemId() string`

GetPlaylistItemId returns the PlaylistItemId field if non-nil, zero value otherwise.

### GetPlaylistItemIdOk

`func (o *JellyfinSessionInfoDto) GetPlaylistItemIdOk() (*string, bool)`

GetPlaylistItemIdOk returns a tuple with the PlaylistItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistItemId

`func (o *JellyfinSessionInfoDto) SetPlaylistItemId(v string)`

SetPlaylistItemId sets PlaylistItemId field to given value.

### HasPlaylistItemId

`func (o *JellyfinSessionInfoDto) HasPlaylistItemId() bool`

HasPlaylistItemId returns a boolean if a field has been set.

### SetPlaylistItemIdNil

`func (o *JellyfinSessionInfoDto) SetPlaylistItemIdNil(b bool)`

 SetPlaylistItemIdNil sets the value for PlaylistItemId to be an explicit nil

### UnsetPlaylistItemId
`func (o *JellyfinSessionInfoDto) UnsetPlaylistItemId()`

UnsetPlaylistItemId ensures that no value is present for PlaylistItemId, not even an explicit nil
### GetServerId

`func (o *JellyfinSessionInfoDto) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *JellyfinSessionInfoDto) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *JellyfinSessionInfoDto) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *JellyfinSessionInfoDto) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### SetServerIdNil

`func (o *JellyfinSessionInfoDto) SetServerIdNil(b bool)`

 SetServerIdNil sets the value for ServerId to be an explicit nil

### UnsetServerId
`func (o *JellyfinSessionInfoDto) UnsetServerId()`

UnsetServerId ensures that no value is present for ServerId, not even an explicit nil
### GetUserPrimaryImageTag

`func (o *JellyfinSessionInfoDto) GetUserPrimaryImageTag() string`

GetUserPrimaryImageTag returns the UserPrimaryImageTag field if non-nil, zero value otherwise.

### GetUserPrimaryImageTagOk

`func (o *JellyfinSessionInfoDto) GetUserPrimaryImageTagOk() (*string, bool)`

GetUserPrimaryImageTagOk returns a tuple with the UserPrimaryImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrimaryImageTag

`func (o *JellyfinSessionInfoDto) SetUserPrimaryImageTag(v string)`

SetUserPrimaryImageTag sets UserPrimaryImageTag field to given value.

### HasUserPrimaryImageTag

`func (o *JellyfinSessionInfoDto) HasUserPrimaryImageTag() bool`

HasUserPrimaryImageTag returns a boolean if a field has been set.

### SetUserPrimaryImageTagNil

`func (o *JellyfinSessionInfoDto) SetUserPrimaryImageTagNil(b bool)`

 SetUserPrimaryImageTagNil sets the value for UserPrimaryImageTag to be an explicit nil

### UnsetUserPrimaryImageTag
`func (o *JellyfinSessionInfoDto) UnsetUserPrimaryImageTag()`

UnsetUserPrimaryImageTag ensures that no value is present for UserPrimaryImageTag, not even an explicit nil
### GetSupportedCommands

`func (o *JellyfinSessionInfoDto) GetSupportedCommands() []JellyfinJellyfinGeneralCommandType`

GetSupportedCommands returns the SupportedCommands field if non-nil, zero value otherwise.

### GetSupportedCommandsOk

`func (o *JellyfinSessionInfoDto) GetSupportedCommandsOk() (*[]JellyfinJellyfinGeneralCommandType, bool)`

GetSupportedCommandsOk returns a tuple with the SupportedCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedCommands

`func (o *JellyfinSessionInfoDto) SetSupportedCommands(v []JellyfinJellyfinGeneralCommandType)`

SetSupportedCommands sets SupportedCommands field to given value.

### HasSupportedCommands

`func (o *JellyfinSessionInfoDto) HasSupportedCommands() bool`

HasSupportedCommands returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


