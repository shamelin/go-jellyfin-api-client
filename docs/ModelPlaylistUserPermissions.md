# ModelPlaylistUserPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** | Gets or sets the user id. | [optional] 
**CanEdit** | Pointer to **bool** | Gets or sets a value indicating whether the user has edit permissions. | [optional] 

## Methods

### NewModelPlaylistUserPermissions

`func NewModelPlaylistUserPermissions() *ModelPlaylistUserPermissions`

NewModelPlaylistUserPermissions instantiates a new ModelPlaylistUserPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelPlaylistUserPermissionsWithDefaults

`func NewModelPlaylistUserPermissionsWithDefaults() *ModelPlaylistUserPermissions`

NewModelPlaylistUserPermissionsWithDefaults instantiates a new ModelPlaylistUserPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *ModelPlaylistUserPermissions) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ModelPlaylistUserPermissions) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ModelPlaylistUserPermissions) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ModelPlaylistUserPermissions) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetCanEdit

`func (o *ModelPlaylistUserPermissions) GetCanEdit() bool`

GetCanEdit returns the CanEdit field if non-nil, zero value otherwise.

### GetCanEditOk

`func (o *ModelPlaylistUserPermissions) GetCanEditOk() (*bool, bool)`

GetCanEditOk returns a tuple with the CanEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEdit

`func (o *ModelPlaylistUserPermissions) SetCanEdit(v bool)`

SetCanEdit sets CanEdit field to given value.

### HasCanEdit

`func (o *ModelPlaylistUserPermissions) HasCanEdit() bool`

HasCanEdit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


