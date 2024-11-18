# ModelContainerProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**ModelModelDlnaProfileType**](ModelDlnaProfileType.md) | Gets or sets the MediaBrowser.Model.Dlna.DlnaProfileType which this container must meet. | [optional] 
**Conditions** | Pointer to [**[]ModelModelProfileCondition**](ModelModelProfileCondition.md) | Gets or sets the list of MediaBrowser.Model.Dlna.ProfileCondition which this container will be applied to. | [optional] 
**Container** | Pointer to **NullableString** | Gets or sets the container(s) which this container must meet. | [optional] 
**SubContainer** | Pointer to **NullableString** | Gets or sets the sub container(s) which this container must meet. | [optional] 

## Methods

### NewModelContainerProfile

`func NewModelContainerProfile() *ModelContainerProfile`

NewModelContainerProfile instantiates a new ModelContainerProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelContainerProfileWithDefaults

`func NewModelContainerProfileWithDefaults() *ModelContainerProfile`

NewModelContainerProfileWithDefaults instantiates a new ModelContainerProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ModelContainerProfile) GetType() ModelModelDlnaProfileType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelContainerProfile) GetTypeOk() (*ModelModelDlnaProfileType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelContainerProfile) SetType(v ModelModelDlnaProfileType)`

SetType sets Type field to given value.

### HasType

`func (o *ModelContainerProfile) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConditions

`func (o *ModelContainerProfile) GetConditions() []ModelModelProfileCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *ModelContainerProfile) GetConditionsOk() (*[]ModelModelProfileCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *ModelContainerProfile) SetConditions(v []ModelModelProfileCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *ModelContainerProfile) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetContainer

`func (o *ModelContainerProfile) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *ModelContainerProfile) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *ModelContainerProfile) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *ModelContainerProfile) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### SetContainerNil

`func (o *ModelContainerProfile) SetContainerNil(b bool)`

 SetContainerNil sets the value for Container to be an explicit nil

### UnsetContainer
`func (o *ModelContainerProfile) UnsetContainer()`

UnsetContainer ensures that no value is present for Container, not even an explicit nil
### GetSubContainer

`func (o *ModelContainerProfile) GetSubContainer() string`

GetSubContainer returns the SubContainer field if non-nil, zero value otherwise.

### GetSubContainerOk

`func (o *ModelContainerProfile) GetSubContainerOk() (*string, bool)`

GetSubContainerOk returns a tuple with the SubContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubContainer

`func (o *ModelContainerProfile) SetSubContainer(v string)`

SetSubContainer sets SubContainer field to given value.

### HasSubContainer

`func (o *ModelContainerProfile) HasSubContainer() bool`

HasSubContainer returns a boolean if a field has been set.

### SetSubContainerNil

`func (o *ModelContainerProfile) SetSubContainerNil(b bool)`

 SetSubContainerNil sets the value for SubContainer to be an explicit nil

### UnsetSubContainer
`func (o *ModelContainerProfile) UnsetSubContainer()`

UnsetSubContainer ensures that no value is present for SubContainer, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


