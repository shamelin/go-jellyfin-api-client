# ModelCodecProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**ModelModelCodecType**](ModelCodecType.md) | Gets or sets the MediaBrowser.Model.Dlna.CodecType which this container must meet. | [optional] 
**Conditions** | Pointer to [**[]ModelModelProfileCondition**](ModelModelProfileCondition.md) | Gets or sets the list of MediaBrowser.Model.Dlna.ProfileCondition which this profile must meet. | [optional] 
**ApplyConditions** | Pointer to [**[]ModelModelProfileCondition**](ModelModelProfileCondition.md) | Gets or sets the list of MediaBrowser.Model.Dlna.ProfileCondition to apply if this profile is met. | [optional] 
**Codec** | Pointer to **NullableString** | Gets or sets the codec(s) that this profile applies to. | [optional] 
**Container** | Pointer to **NullableString** | Gets or sets the container(s) which this profile will be applied to. | [optional] 
**SubContainer** | Pointer to **NullableString** | Gets or sets the sub-container(s) which this profile will be applied to. | [optional] 

## Methods

### NewModelCodecProfile

`func NewModelCodecProfile() *ModelCodecProfile`

NewModelCodecProfile instantiates a new ModelCodecProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelCodecProfileWithDefaults

`func NewModelCodecProfileWithDefaults() *ModelCodecProfile`

NewModelCodecProfileWithDefaults instantiates a new ModelCodecProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ModelCodecProfile) GetType() ModelModelCodecType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelCodecProfile) GetTypeOk() (*ModelModelCodecType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelCodecProfile) SetType(v ModelModelCodecType)`

SetType sets Type field to given value.

### HasType

`func (o *ModelCodecProfile) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConditions

`func (o *ModelCodecProfile) GetConditions() []ModelModelProfileCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *ModelCodecProfile) GetConditionsOk() (*[]ModelModelProfileCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *ModelCodecProfile) SetConditions(v []ModelModelProfileCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *ModelCodecProfile) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetApplyConditions

`func (o *ModelCodecProfile) GetApplyConditions() []ModelModelProfileCondition`

GetApplyConditions returns the ApplyConditions field if non-nil, zero value otherwise.

### GetApplyConditionsOk

`func (o *ModelCodecProfile) GetApplyConditionsOk() (*[]ModelModelProfileCondition, bool)`

GetApplyConditionsOk returns a tuple with the ApplyConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyConditions

`func (o *ModelCodecProfile) SetApplyConditions(v []ModelModelProfileCondition)`

SetApplyConditions sets ApplyConditions field to given value.

### HasApplyConditions

`func (o *ModelCodecProfile) HasApplyConditions() bool`

HasApplyConditions returns a boolean if a field has been set.

### GetCodec

`func (o *ModelCodecProfile) GetCodec() string`

GetCodec returns the Codec field if non-nil, zero value otherwise.

### GetCodecOk

`func (o *ModelCodecProfile) GetCodecOk() (*string, bool)`

GetCodecOk returns a tuple with the Codec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodec

`func (o *ModelCodecProfile) SetCodec(v string)`

SetCodec sets Codec field to given value.

### HasCodec

`func (o *ModelCodecProfile) HasCodec() bool`

HasCodec returns a boolean if a field has been set.

### SetCodecNil

`func (o *ModelCodecProfile) SetCodecNil(b bool)`

 SetCodecNil sets the value for Codec to be an explicit nil

### UnsetCodec
`func (o *ModelCodecProfile) UnsetCodec()`

UnsetCodec ensures that no value is present for Codec, not even an explicit nil
### GetContainer

`func (o *ModelCodecProfile) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *ModelCodecProfile) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *ModelCodecProfile) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *ModelCodecProfile) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### SetContainerNil

`func (o *ModelCodecProfile) SetContainerNil(b bool)`

 SetContainerNil sets the value for Container to be an explicit nil

### UnsetContainer
`func (o *ModelCodecProfile) UnsetContainer()`

UnsetContainer ensures that no value is present for Container, not even an explicit nil
### GetSubContainer

`func (o *ModelCodecProfile) GetSubContainer() string`

GetSubContainer returns the SubContainer field if non-nil, zero value otherwise.

### GetSubContainerOk

`func (o *ModelCodecProfile) GetSubContainerOk() (*string, bool)`

GetSubContainerOk returns a tuple with the SubContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubContainer

`func (o *ModelCodecProfile) SetSubContainer(v string)`

SetSubContainer sets SubContainer field to given value.

### HasSubContainer

`func (o *ModelCodecProfile) HasSubContainer() bool`

HasSubContainer returns a boolean if a field has been set.

### SetSubContainerNil

`func (o *ModelCodecProfile) SetSubContainerNil(b bool)`

 SetSubContainerNil sets the value for SubContainer to be an explicit nil

### UnsetSubContainer
`func (o *ModelCodecProfile) UnsetSubContainer()`

UnsetSubContainer ensures that no value is present for SubContainer, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


