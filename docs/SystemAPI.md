# \SystemAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEndpointInfo**](SystemAPI.md#GetEndpointInfo) | **Get** /System/Endpoint | Gets information about the request endpoint.
[**GetLogFile**](SystemAPI.md#GetLogFile) | **Get** /System/Logs/Log | Gets a log file.
[**GetPingSystem**](SystemAPI.md#GetPingSystem) | **Get** /System/Ping | Pings the system.
[**GetPublicSystemInfo**](SystemAPI.md#GetPublicSystemInfo) | **Get** /System/Info/Public | Gets public information about the server.
[**GetServerLogs**](SystemAPI.md#GetServerLogs) | **Get** /System/Logs | Gets a list of available server log files.
[**GetSystemInfo**](SystemAPI.md#GetSystemInfo) | **Get** /System/Info | Gets information about the server.
[**GetWakeOnLanInfo**](SystemAPI.md#GetWakeOnLanInfo) | **Get** /System/WakeOnLanInfo | Gets wake on lan information.
[**PostPingSystem**](SystemAPI.md#PostPingSystem) | **Post** /System/Ping | Pings the system.
[**RestartApplication**](SystemAPI.md#RestartApplication) | **Post** /System/Restart | Restarts the application.
[**ShutdownApplication**](SystemAPI.md#ShutdownApplication) | **Post** /System/Shutdown | Shuts down the application.



## GetEndpointInfo

> JellyfinJellyfinEndPointInfo GetEndpointInfo(ctx).Execute()

Gets information about the request endpoint.

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
	resp, r, err := apiClient.SystemAPI.GetEndpointInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.GetEndpointInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEndpointInfo`: JellyfinJellyfinEndPointInfo
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.GetEndpointInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointInfoRequest struct via the builder pattern


### Return type

[**JellyfinJellyfinEndPointInfo**](JellyfinEndPointInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogFile

> *os.File GetLogFile(ctx).Name(name).Execute()

Gets a log file.

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
	name := "name_example" // string | The name of the log file to get.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.GetLogFile(context.Background()).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.GetLogFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogFile`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.GetLogFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLogFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name of the log file to get. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPingSystem

> string GetPingSystem(ctx).Execute()

Pings the system.

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
	resp, r, err := apiClient.SystemAPI.GetPingSystem(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.GetPingSystem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPingSystem`: string
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.GetPingSystem`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPingSystemRequest struct via the builder pattern


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicSystemInfo

> JellyfinJellyfinPublicSystemInfo GetPublicSystemInfo(ctx).Execute()

Gets public information about the server.

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
	resp, r, err := apiClient.SystemAPI.GetPublicSystemInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.GetPublicSystemInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicSystemInfo`: JellyfinJellyfinPublicSystemInfo
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.GetPublicSystemInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicSystemInfoRequest struct via the builder pattern


### Return type

[**JellyfinJellyfinPublicSystemInfo**](JellyfinPublicSystemInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerLogs

> []JellyfinJellyfinLogFile GetServerLogs(ctx).Execute()

Gets a list of available server log files.

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
	resp, r, err := apiClient.SystemAPI.GetServerLogs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.GetServerLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerLogs`: []JellyfinJellyfinLogFile
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.GetServerLogs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerLogsRequest struct via the builder pattern


### Return type

[**[]JellyfinJellyfinLogFile**](JellyfinLogFile.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemInfo

> JellyfinJellyfinSystemInfo GetSystemInfo(ctx).Execute()

Gets information about the server.

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
	resp, r, err := apiClient.SystemAPI.GetSystemInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.GetSystemInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemInfo`: JellyfinJellyfinSystemInfo
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.GetSystemInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemInfoRequest struct via the builder pattern


### Return type

[**JellyfinJellyfinSystemInfo**](JellyfinSystemInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWakeOnLanInfo

> []JellyfinJellyfinWakeOnLanInfo GetWakeOnLanInfo(ctx).Execute()

Gets wake on lan information.

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
	resp, r, err := apiClient.SystemAPI.GetWakeOnLanInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.GetWakeOnLanInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWakeOnLanInfo`: []JellyfinJellyfinWakeOnLanInfo
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.GetWakeOnLanInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWakeOnLanInfoRequest struct via the builder pattern


### Return type

[**[]JellyfinJellyfinWakeOnLanInfo**](JellyfinWakeOnLanInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPingSystem

> string PostPingSystem(ctx).Execute()

Pings the system.

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
	resp, r, err := apiClient.SystemAPI.PostPingSystem(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.PostPingSystem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPingSystem`: string
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.PostPingSystem`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostPingSystemRequest struct via the builder pattern


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestartApplication

> RestartApplication(ctx).Execute()

Restarts the application.

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
	r, err := apiClient.SystemAPI.RestartApplication(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.RestartApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRestartApplicationRequest struct via the builder pattern


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


## ShutdownApplication

> ShutdownApplication(ctx).Execute()

Shuts down the application.

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
	r, err := apiClient.SystemAPI.ShutdownApplication(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.ShutdownApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiShutdownApplicationRequest struct via the builder pattern


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

