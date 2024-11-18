# JellyfinContainerProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**JellyfinJellyfinDlnaProfileType**](JellyfinDlnaProfileType.md) | Gets or sets the MediaBrowser.Model.Dlna.DlnaProfileType which this container must meet. | [optional] 
**Conditions** | Pointer to [**[]JellyfinJellyfinProfileCondition**](JellyfinJellyfinProfileCondition.md) | Gets or sets the list of MediaBrowser.Model.Dlna.ProfileCondition which this container will be applied to. | [optional] 
**Container** | Pointer to **NullableString** | Gets or sets the container(s) which this container must meet. | [optional] 
**SubContainer** | Pointer to **NullableString** | Gets or sets the sub container(s) which this container must meet. | [optional] 

## Methods

### NewJellyfinContainerProfile

`func NewJellyfinContainerProfile() *JellyfinContainerProfile`

NewJellyfinContainerProfile instantiates a new JellyfinContainerProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinContainerProfileWithDefaults

`func NewJellyfinContainerProfileWithDefaults() *JellyfinContainerProfile`

NewJellyfinContainerProfileWithDefaults instantiates a new JellyfinContainerProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *JellyfinContainerProfile) GetType() JellyfinJellyfinDlnaProfileType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JellyfinContainerProfile) GetTypeOk() (*JellyfinJellyfinDlnaProfileType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JellyfinContainerProfile) SetType(v JellyfinJellyfinDlnaProfileType)`

SetType sets Type field to given value.

### HasType

`func (o *JellyfinContainerProfile) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConditions

`func (o *JellyfinContainerProfile) GetConditions() []JellyfinJellyfinProfileCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *JellyfinContainerProfile) GetConditionsOk() (*[]JellyfinJellyfinProfileCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *JellyfinContainerProfile) SetConditions(v []JellyfinJellyfinProfileCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *JellyfinContainerProfile) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetContainer

`func (o *JellyfinContainerProfile) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *JellyfinContainerProfile) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *JellyfinContainerProfile) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *JellyfinContainerProfile) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### SetContainerNil

`func (o *JellyfinContainerProfile) SetContainerNil(b bool)`

 SetContainerNil sets the value for Container to be an explicit nil

### UnsetContainer
`func (o *JellyfinContainerProfile) UnsetContainer()`

UnsetContainer ensures that no value is present for Container, not even an explicit nil
### GetSubContainer

`func (o *JellyfinContainerProfile) GetSubContainer() string`

GetSubContainer returns the SubContainer field if non-nil, zero value otherwise.

### GetSubContainerOk

`func (o *JellyfinContainerProfile) GetSubContainerOk() (*string, bool)`

GetSubContainerOk returns a tuple with the SubContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubContainer

`func (o *JellyfinContainerProfile) SetSubContainer(v string)`

SetSubContainer sets SubContainer field to given value.

### HasSubContainer

`func (o *JellyfinContainerProfile) HasSubContainer() bool`

HasSubContainer returns a boolean if a field has been set.

### SetSubContainerNil

`func (o *JellyfinContainerProfile) SetSubContainerNil(b bool)`

 SetSubContainerNil sets the value for SubContainer to be an explicit nil

### UnsetSubContainer
`func (o *JellyfinContainerProfile) UnsetSubContainer()`

UnsetSubContainer ensures that no value is present for SubContainer, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


