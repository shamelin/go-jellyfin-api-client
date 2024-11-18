# \UserAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthenticateUserByName**](UserAPI.md#AuthenticateUserByName) | **Post** /Users/AuthenticateByName | Authenticates a user by name.
[**AuthenticateWithQuickConnect**](UserAPI.md#AuthenticateWithQuickConnect) | **Post** /Users/AuthenticateWithQuickConnect | Authenticates a user with quick connect.
[**CreateUserByName**](UserAPI.md#CreateUserByName) | **Post** /Users/New | Creates a user.
[**DeleteUser**](UserAPI.md#DeleteUser) | **Delete** /Users/{userId} | Deletes a user.
[**ForgotPassword**](UserAPI.md#ForgotPassword) | **Post** /Users/ForgotPassword | Initiates the forgot password process for a local user.
[**ForgotPasswordPin**](UserAPI.md#ForgotPasswordPin) | **Post** /Users/ForgotPassword/Pin | Redeems a forgot password pin.
[**GetCurrentUser**](UserAPI.md#GetCurrentUser) | **Get** /Users/Me | Gets the user based on auth token.
[**GetPublicUsers**](UserAPI.md#GetPublicUsers) | **Get** /Users/Public | Gets a list of publicly visible users for display on a login screen.
[**GetUserById**](UserAPI.md#GetUserById) | **Get** /Users/{userId} | Gets a user by Id.
[**GetUsers**](UserAPI.md#GetUsers) | **Get** /Users | Gets a list of users.
[**UpdateUser**](UserAPI.md#UpdateUser) | **Post** /Users | Updates a user.
[**UpdateUserConfiguration**](UserAPI.md#UpdateUserConfiguration) | **Post** /Users/Configuration | Updates a user configuration.
[**UpdateUserPassword**](UserAPI.md#UpdateUserPassword) | **Post** /Users/Password | Updates a user&#39;s password.
[**UpdateUserPolicy**](UserAPI.md#UpdateUserPolicy) | **Post** /Users/{userId}/Policy | Updates a user policy.



## AuthenticateUserByName

> JellyfinJellyfinAuthenticationResult AuthenticateUserByName(ctx).JellyfinAuthenticateUserByName(jellyfinAuthenticateUserByName).Execute()

Authenticates a user by name.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/shamelin/go-jellyfin-api-client"
)

func main() {
	jellyfinAuthenticateUserByName := *openapiclient.NewJellyfinAuthenticateUserByName() // JellyfinAuthenticateUserByName | The M:Jellyfin.Api.Controllers.UserController.AuthenticateUserByName(Jellyfin.Api.Models.UserDtos.AuthenticateUserByName) request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.AuthenticateUserByName(context.Background()).JellyfinAuthenticateUserByName(jellyfinAuthenticateUserByName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.AuthenticateUserByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthenticateUserByName`: JellyfinJellyfinAuthenticationResult
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.AuthenticateUserByName`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticateUserByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jellyfinAuthenticateUserByName** | [**JellyfinAuthenticateUserByName**](JellyfinAuthenticateUserByName.md) | The M:Jellyfin.Api.Controllers.UserController.AuthenticateUserByName(Jellyfin.Api.Models.UserDtos.AuthenticateUserByName) request. | 

### Return type

[**JellyfinJellyfinAuthenticationResult**](JellyfinAuthenticationResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticateWithQuickConnect

> JellyfinJellyfinAuthenticationResult AuthenticateWithQuickConnect(ctx).JellyfinQuickConnectDto(jellyfinQuickConnectDto).Execute()

Authenticates a user with quick connect.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/shamelin/go-jellyfin-api-client"
)

func main() {
	jellyfinQuickConnectDto := *openapiclient.NewJellyfinQuickConnectDto("Secret_example") // JellyfinQuickConnectDto | The Jellyfin.Api.Models.UserDtos.QuickConnectDto request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.AuthenticateWithQuickConnect(context.Background()).JellyfinQuickConnectDto(jellyfinQuickConnectDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.AuthenticateWithQuickConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthenticateWithQuickConnect`: JellyfinJellyfinAuthenticationResult
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.AuthenticateWithQuickConnect`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticateWithQuickConnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jellyfinQuickConnectDto** | [**JellyfinQuickConnectDto**](JellyfinQuickConnectDto.md) | The Jellyfin.Api.Models.UserDtos.QuickConnectDto request. | 

### Return type

[**JellyfinJellyfinAuthenticationResult**](JellyfinAuthenticationResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUserByName

> JellyfinJellyfinUserDto CreateUserByName(ctx).JellyfinCreateUserByName(jellyfinCreateUserByName).Execute()

Creates a user.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/shamelin/go-jellyfin-api-client"
)

func main() {
	jellyfinCreateUserByName := *openapiclient.NewJellyfinCreateUserByName("Name_example") // JellyfinCreateUserByName | The create user by name request body.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.CreateUserByName(context.Background()).JellyfinCreateUserByName(jellyfinCreateUserByName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.CreateUserByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUserByName`: JellyfinJellyfinUserDto
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.CreateUserByName`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jellyfinCreateUserByName** | [**JellyfinCreateUserByName**](JellyfinCreateUserByName.md) | The create user by name request body. | 

### Return type

[**JellyfinJellyfinUserDto**](JellyfinUserDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> DeleteUser(ctx, userId).Execute()

Deletes a user.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/shamelin/go-jellyfin-api-client"
)

func main() {
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The user id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.DeleteUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.DeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The user id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ForgotPassword

> JellyfinJellyfinForgotPasswordResult ForgotPassword(ctx).JellyfinForgotPasswordDto(jellyfinForgotPasswordDto).Execute()

Initiates the forgot password process for a local user.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/shamelin/go-jellyfin-api-client"
)

func main() {
	jellyfinForgotPasswordDto := *openapiclient.NewJellyfinForgotPasswordDto("EnteredUsername_example") // JellyfinForgotPasswordDto | The forgot password request containing the entered username.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.ForgotPassword(context.Background()).JellyfinForgotPasswordDto(jellyfinForgotPasswordDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.ForgotPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ForgotPassword`: JellyfinJellyfinForgotPasswordResult
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.ForgotPassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiForgotPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jellyfinForgotPasswordDto** | [**JellyfinForgotPasswordDto**](JellyfinForgotPasswordDto.md) | The forgot password request containing the entered username. | 

### Return type

[**JellyfinJellyfinForgotPasswordResult**](JellyfinForgotPasswordResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ForgotPasswordPin

> JellyfinJellyfinPinRedeemResult ForgotPasswordPin(ctx).JellyfinForgotPasswordPinDto(jellyfinForgotPasswordPinDto).Execute()

Redeems a forgot password pin.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/shamelin/go-jellyfin-api-client"
)

func main() {
	jellyfinForgotPasswordPinDto := *openapiclient.NewJellyfinForgotPasswordPinDto("Pin_example") // JellyfinForgotPasswordPinDto | The forgot password pin request containing the entered pin.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.ForgotPasswordPin(context.Background()).JellyfinForgotPasswordPinDto(jellyfinForgotPasswordPinDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.ForgotPasswordPin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ForgotPasswordPin`: JellyfinJellyfinPinRedeemResult
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.ForgotPasswordPin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiForgotPasswordPinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jellyfinForgotPasswordPinDto** | [**JellyfinForgotPasswordPinDto**](JellyfinForgotPasswordPinDto.md) | The forgot password pin request containing the entered pin. | 

### Return type

[**JellyfinJellyfinPinRedeemResult**](JellyfinPinRedeemResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentUser

> JellyfinJellyfinUserDto GetCurrentUser(ctx).Execute()

Gets the user based on auth token.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/shamelin/go-jellyfin-api-client"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.GetCurrentUser(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetCurrentUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentUser`: JellyfinJellyfinUserDto
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetCurrentUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentUserRequest struct via the builder pattern


### Return type

[**JellyfinJellyfinUserDto**](JellyfinUserDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicUsers

> []JellyfinJellyfinUserDto GetPublicUsers(ctx).Execute()

Gets a list of publicly visible users for display on a login screen.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/shamelin/go-jellyfin-api-client"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.GetPublicUsers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetPublicUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicUsers`: []JellyfinJellyfinUserDto
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetPublicUsers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicUsersRequest struct via the builder pattern


### Return type

[**[]JellyfinJellyfinUserDto**](JellyfinUserDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserById

> JellyfinJellyfinUserDto GetUserById(ctx, userId).Execute()

Gets a user by Id.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/shamelin/go-jellyfin-api-client"
)

func main() {
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The user id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.GetUserById(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetUserById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserById`: JellyfinJellyfinUserDto
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetUserById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The user id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JellyfinJellyfinUserDto**](JellyfinUserDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsers

> []JellyfinJellyfinUserDto GetUsers(ctx).IsHidden(isHidden).IsDisabled(isDisabled).Execute()

Gets a list of users.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/shamelin/go-jellyfin-api-client"
)

func main() {
	isHidden := true // bool | Optional filter by IsHidden=true or false. (optional)
	isDisabled := true // bool | Optional filter by IsDisabled=true or false. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.GetUsers(context.Background()).IsHidden(isHidden).IsDisabled(isDisabled).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsers`: []JellyfinJellyfinUserDto
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isHidden** | **bool** | Optional filter by IsHidden&#x3D;true or false. | 
 **isDisabled** | **bool** | Optional filter by IsDisabled&#x3D;true or false. | 

### Return type

[**[]JellyfinJellyfinUserDto**](JellyfinUserDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> UpdateUser(ctx).JellyfinUserDto(jellyfinUserDto).UserId(userId).Execute()

Updates a user.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/shamelin/go-jellyfin-api-client"
)

func main() {
	jellyfinUserDto := *openapiclient.NewJellyfinUserDto() // JellyfinUserDto | The updated user model.
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The user id. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.UpdateUser(context.Background()).JellyfinUserDto(jellyfinUserDto).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UpdateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jellyfinUserDto** | [**JellyfinUserDto**](JellyfinUserDto.md) | The updated user model. | 
 **userId** | **string** | The user id. | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserConfiguration

> UpdateUserConfiguration(ctx).JellyfinUserConfiguration(jellyfinUserConfiguration).UserId(userId).Execute()

Updates a user configuration.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/shamelin/go-jellyfin-api-client"
)

func main() {
	jellyfinUserConfiguration := *openapiclient.NewJellyfinUserConfiguration() // JellyfinUserConfiguration | The new user configuration.
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The user id. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.UpdateUserConfiguration(context.Background()).JellyfinUserConfiguration(jellyfinUserConfiguration).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UpdateUserConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jellyfinUserConfiguration** | [**JellyfinUserConfiguration**](JellyfinUserConfiguration.md) | The new user configuration. | 
 **userId** | **string** | The user id. | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserPassword

> UpdateUserPassword(ctx).JellyfinUpdateUserPassword(jellyfinUpdateUserPassword).UserId(userId).Execute()

Updates a user's password.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/shamelin/go-jellyfin-api-client"
)

func main() {
	jellyfinUpdateUserPassword := *openapiclient.NewJellyfinUpdateUserPassword() // JellyfinUpdateUserPassword | The M:Jellyfin.Api.Controllers.UserController.UpdateUserPassword(System.Nullable{System.Guid},Jellyfin.Api.Models.UserDtos.UpdateUserPassword) request.
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The user id. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.UpdateUserPassword(context.Background()).JellyfinUpdateUserPassword(jellyfinUpdateUserPassword).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UpdateUserPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jellyfinUpdateUserPassword** | [**JellyfinUpdateUserPassword**](JellyfinUpdateUserPassword.md) | The M:Jellyfin.Api.Controllers.UserController.UpdateUserPassword(System.Nullable{System.Guid},Jellyfin.Api.Models.UserDtos.UpdateUserPassword) request. | 
 **userId** | **string** | The user id. | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserPolicy

> UpdateUserPolicy(ctx, userId).JellyfinUserPolicy(jellyfinUserPolicy).Execute()

Updates a user policy.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/shamelin/go-jellyfin-api-client"
)

func main() {
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The user id.
	jellyfinUserPolicy := *openapiclient.NewJellyfinUserPolicy("AuthenticationProviderId_example", "PasswordResetProviderId_example") // JellyfinUserPolicy | The new user policy.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.UpdateUserPolicy(context.Background(), userId).JellyfinUserPolicy(jellyfinUserPolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UpdateUserPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The user id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jellyfinUserPolicy** | [**JellyfinUserPolicy**](JellyfinUserPolicy.md) | The new user policy. | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

