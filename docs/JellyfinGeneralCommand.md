# JellyfinGeneralCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to [**JellyfinJellyfinGeneralCommandType**](JellyfinGeneralCommandType.md) | This exists simply to identify a set of known commands. | [optional] 
**ControllingUserId** | Pointer to **string** |  | [optional] 
**Arguments** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewJellyfinGeneralCommand

`func NewJellyfinGeneralCommand() *JellyfinGeneralCommand`

NewJellyfinGeneralCommand instantiates a new JellyfinGeneralCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinGeneralCommandWithDefaults

`func NewJellyfinGeneralCommandWithDefaults() *JellyfinGeneralCommand`

NewJellyfinGeneralCommandWithDefaults instantiates a new JellyfinGeneralCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *JellyfinGeneralCommand) GetName() JellyfinJellyfinGeneralCommandType`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JellyfinGeneralCommand) GetNameOk() (*JellyfinJellyfinGeneralCommandType, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JellyfinGeneralCommand) SetName(v JellyfinJellyfinGeneralCommandType)`

SetName sets Name field to given value.

### HasName

`func (o *JellyfinGeneralCommand) HasName() bool`

HasName returns a boolean if a field has been set.

### GetControllingUserId

`func (o *JellyfinGeneralCommand) GetControllingUserId() string`

GetControllingUserId returns the ControllingUserId field if non-nil, zero value otherwise.

### GetControllingUserIdOk

`func (o *JellyfinGeneralCommand) GetControllingUserIdOk() (*string, bool)`

GetControllingUserIdOk returns a tuple with the ControllingUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllingUserId

`func (o *JellyfinGeneralCommand) SetControllingUserId(v string)`

SetControllingUserId sets ControllingUserId field to given value.

### HasControllingUserId

`func (o *JellyfinGeneralCommand) HasControllingUserId() bool`

HasControllingUserId returns a boolean if a field has been set.

### GetArguments

`func (o *JellyfinGeneralCommand) GetArguments() map[string]string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *JellyfinGeneralCommand) GetArgumentsOk() (*map[string]string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *JellyfinGeneralCommand) SetArguments(v map[string]string)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *JellyfinGeneralCommand) HasArguments() bool`

HasArguments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


