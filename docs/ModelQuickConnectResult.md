# ModelQuickConnectResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authenticated** | Pointer to **bool** | Gets or sets a value indicating whether this request is authorized. | [optional] 
**Secret** | Pointer to **string** | Gets the secret value used to uniquely identify this request. Can be used to retrieve authentication information. | [optional] 
**Code** | Pointer to **string** | Gets the user facing code used so the user can quickly differentiate this request from others. | [optional] 
**DeviceId** | Pointer to **string** | Gets the requesting device id. | [optional] 
**DeviceName** | Pointer to **string** | Gets the requesting device name. | [optional] 
**AppName** | Pointer to **string** | Gets the requesting app name. | [optional] 
**AppVersion** | Pointer to **string** | Gets the requesting app version. | [optional] 
**DateAdded** | Pointer to **time.Time** | Gets or sets the DateTime that this request was created. | [optional] 

## Methods

### NewModelQuickConnectResult

`func NewModelQuickConnectResult() *ModelQuickConnectResult`

NewModelQuickConnectResult instantiates a new ModelQuickConnectResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelQuickConnectResultWithDefaults

`func NewModelQuickConnectResultWithDefaults() *ModelQuickConnectResult`

NewModelQuickConnectResultWithDefaults instantiates a new ModelQuickConnectResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticated

`func (o *ModelQuickConnectResult) GetAuthenticated() bool`

GetAuthenticated returns the Authenticated field if non-nil, zero value otherwise.

### GetAuthenticatedOk

`func (o *ModelQuickConnectResult) GetAuthenticatedOk() (*bool, bool)`

GetAuthenticatedOk returns a tuple with the Authenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticated

`func (o *ModelQuickConnectResult) SetAuthenticated(v bool)`

SetAuthenticated sets Authenticated field to given value.

### HasAuthenticated

`func (o *ModelQuickConnectResult) HasAuthenticated() bool`

HasAuthenticated returns a boolean if a field has been set.

### GetSecret

`func (o *ModelQuickConnectResult) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ModelQuickConnectResult) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ModelQuickConnectResult) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *ModelQuickConnectResult) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetCode

`func (o *ModelQuickConnectResult) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ModelQuickConnectResult) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ModelQuickConnectResult) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ModelQuickConnectResult) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDeviceId

`func (o *ModelQuickConnectResult) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *ModelQuickConnectResult) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *ModelQuickConnectResult) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *ModelQuickConnectResult) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDeviceName

`func (o *ModelQuickConnectResult) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *ModelQuickConnectResult) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *ModelQuickConnectResult) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *ModelQuickConnectResult) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetAppName

`func (o *ModelQuickConnectResult) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *ModelQuickConnectResult) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *ModelQuickConnectResult) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *ModelQuickConnectResult) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### GetAppVersion

`func (o *ModelQuickConnectResult) GetAppVersion() string`

GetAppVersion returns the AppVersion field if non-nil, zero value otherwise.

### GetAppVersionOk

`func (o *ModelQuickConnectResult) GetAppVersionOk() (*string, bool)`

GetAppVersionOk returns a tuple with the AppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVersion

`func (o *ModelQuickConnectResult) SetAppVersion(v string)`

SetAppVersion sets AppVersion field to given value.

### HasAppVersion

`func (o *ModelQuickConnectResult) HasAppVersion() bool`

HasAppVersion returns a boolean if a field has been set.

### GetDateAdded

`func (o *ModelQuickConnectResult) GetDateAdded() time.Time`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *ModelQuickConnectResult) GetDateAddedOk() (*time.Time, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *ModelQuickConnectResult) SetDateAdded(v time.Time)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *ModelQuickConnectResult) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


