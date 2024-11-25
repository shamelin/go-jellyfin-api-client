# \LiveTvAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddListingProvider**](LiveTvAPI.md#AddListingProvider) | **Post** /LiveTv/ListingProviders | Adds a listings provider.
[**AddTunerHost**](LiveTvAPI.md#AddTunerHost) | **Post** /LiveTv/TunerHosts | Adds a tuner host.
[**CancelSeriesTimer**](LiveTvAPI.md#CancelSeriesTimer) | **Delete** /LiveTv/SeriesTimers/{timerId} | Cancels a live tv series timer.
[**CancelTimer**](LiveTvAPI.md#CancelTimer) | **Delete** /LiveTv/Timers/{timerId} | Cancels a live tv timer.
[**CreateSeriesTimer**](LiveTvAPI.md#CreateSeriesTimer) | **Post** /LiveTv/SeriesTimers | Creates a live tv series timer.
[**CreateTimer**](LiveTvAPI.md#CreateTimer) | **Post** /LiveTv/Timers | Creates a live tv timer.
[**DeleteListingProvider**](LiveTvAPI.md#DeleteListingProvider) | **Delete** /LiveTv/ListingProviders | Delete listing provider.
[**DeleteRecording**](LiveTvAPI.md#DeleteRecording) | **Delete** /LiveTv/Recordings/{recordingId} | Deletes a live tv recording.
[**DeleteTunerHost**](LiveTvAPI.md#DeleteTunerHost) | **Delete** /LiveTv/TunerHosts | Deletes a tuner host.
[**DiscoverTuners**](LiveTvAPI.md#DiscoverTuners) | **Get** /LiveTv/Tuners/Discover | Discover tuners.
[**DiscvoverTuners**](LiveTvAPI.md#DiscvoverTuners) | **Get** /LiveTv/Tuners/Discvover | Discover tuners.
[**GetChannel**](LiveTvAPI.md#GetChannel) | **Get** /LiveTv/Channels/{channelId} | Gets a live tv channel.
[**GetChannelMappingOptions**](LiveTvAPI.md#GetChannelMappingOptions) | **Get** /LiveTv/ChannelMappingOptions | Get channel mapping options.
[**GetDefaultListingProvider**](LiveTvAPI.md#GetDefaultListingProvider) | **Get** /LiveTv/ListingProviders/Default | Gets default listings provider info.
[**GetDefaultTimer**](LiveTvAPI.md#GetDefaultTimer) | **Get** /LiveTv/Timers/Defaults | Gets the default values for a new timer.
[**GetGuideInfo**](LiveTvAPI.md#GetGuideInfo) | **Get** /LiveTv/GuideInfo | Get guid info.
[**GetLineups**](LiveTvAPI.md#GetLineups) | **Get** /LiveTv/ListingProviders/Lineups | Gets available lineups.
[**GetLiveRecordingFile**](LiveTvAPI.md#GetLiveRecordingFile) | **Get** /LiveTv/LiveRecordings/{recordingId}/stream | Gets a live tv recording stream.
[**GetLiveStreamFile**](LiveTvAPI.md#GetLiveStreamFile) | **Get** /LiveTv/LiveStreamFiles/{streamId}/stream.{container} | Gets a live tv channel stream.
[**GetLiveTvChannels**](LiveTvAPI.md#GetLiveTvChannels) | **Get** /LiveTv/Channels | Gets available live tv channels.
[**GetLiveTvInfo**](LiveTvAPI.md#GetLiveTvInfo) | **Get** /LiveTv/Info | Gets available live tv services.
[**GetLiveTvPrograms**](LiveTvAPI.md#GetLiveTvPrograms) | **Get** /LiveTv/Programs | Gets available live tv epgs.
[**GetProgram**](LiveTvAPI.md#GetProgram) | **Get** /LiveTv/Programs/{programId} | Gets a live tv program.
[**GetPrograms**](LiveTvAPI.md#GetPrograms) | **Post** /LiveTv/Programs | Gets available live tv epgs.
[**GetRecommendedPrograms**](LiveTvAPI.md#GetRecommendedPrograms) | **Get** /LiveTv/Programs/Recommended | Gets recommended live tv epgs.
[**GetRecording**](LiveTvAPI.md#GetRecording) | **Get** /LiveTv/Recordings/{recordingId} | Gets a live tv recording.
[**GetRecordingFolders**](LiveTvAPI.md#GetRecordingFolders) | **Get** /LiveTv/Recordings/Folders | Gets recording folders.
[**GetRecordingGroup**](LiveTvAPI.md#GetRecordingGroup) | **Get** /LiveTv/Recordings/Groups/{groupId} | Get recording group.
[**GetRecordingGroups**](LiveTvAPI.md#GetRecordingGroups) | **Get** /LiveTv/Recordings/Groups | Gets live tv recording groups.
[**GetRecordings**](LiveTvAPI.md#GetRecordings) | **Get** /LiveTv/Recordings | Gets live tv recordings.
[**GetRecordingsSeries**](LiveTvAPI.md#GetRecordingsSeries) | **Get** /LiveTv/Recordings/Series | Gets live tv recording series.
[**GetSchedulesDirectCountries**](LiveTvAPI.md#GetSchedulesDirectCountries) | **Get** /LiveTv/ListingProviders/SchedulesDirect/Countries | Gets available countries.
[**GetSeriesTimer**](LiveTvAPI.md#GetSeriesTimer) | **Get** /LiveTv/SeriesTimers/{timerId} | Gets a live tv series timer.
[**GetSeriesTimers**](LiveTvAPI.md#GetSeriesTimers) | **Get** /LiveTv/SeriesTimers | Gets live tv series timers.
[**GetTimer**](LiveTvAPI.md#GetTimer) | **Get** /LiveTv/Timers/{timerId} | Gets a timer.
[**GetTimers**](LiveTvAPI.md#GetTimers) | **Get** /LiveTv/Timers | Gets the live tv timers.
[**GetTunerHostTypes**](LiveTvAPI.md#GetTunerHostTypes) | **Get** /LiveTv/TunerHosts/Types | Get tuner host types.
[**ResetTuner**](LiveTvAPI.md#ResetTuner) | **Post** /LiveTv/Tuners/{tunerId}/Reset | Resets a tv tuner.
[**SetChannelMapping**](LiveTvAPI.md#SetChannelMapping) | **Post** /LiveTv/ChannelMappings | Set channel mappings.
[**UpdateSeriesTimer**](LiveTvAPI.md#UpdateSeriesTimer) | **Post** /LiveTv/SeriesTimers/{timerId} | Updates a live tv series timer.
[**UpdateTimer**](LiveTvAPI.md#UpdateTimer) | **Post** /LiveTv/Timers/{timerId} | Updates a live tv timer.



## AddListingProvider

> JellyfinJellyfinListingsProviderInfo AddListingProvider(ctx).Pw(pw).ValidateListings(validateListings).ValidateLogin(validateLogin).JellyfinListingsProviderInfo(jellyfinListingsProviderInfo).Execute()

Adds a listings provider.

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
	pw := "pw_example" // string | Password. (optional)
	validateListings := true // bool | Validate listings. (optional) (default to false)
	validateLogin := true // bool | Validate login. (optional) (default to false)
	jellyfinListingsProviderInfo := *openapiclient.NewJellyfinListingsProviderInfo() // JellyfinListingsProviderInfo | New listings info. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvAPI.AddListingProvider(context.Background()).Pw(pw).ValidateListings(validateListings).ValidateLogin(validateLogin).JellyfinListingsProviderInfo(jellyfinListingsProviderInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvAPI.AddListingProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddListingProvider`: JellyfinJellyfinListingsProviderInfo
	fmt.Fprintf(os.Stdout, "Response from `LiveTvAPI.AddListingProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddListingProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pw** | **string** | Password. | 
 **validateListings** | **bool** | Validate listings. | [default to false]
 **validateLogin** | **bool** | Validate login. | [default to false]
 **jellyfinListingsProviderInfo** | [**JellyfinListingsProviderInfo**](JellyfinListingsProviderInfo.md) | New listings info. | 

### Return type

[**JellyfinJellyfinListingsProviderInfo**](JellyfinListingsProviderInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddTunerHost

> JellyfinJellyfinTunerHostInfo AddTunerHost(ctx).JellyfinTunerHostInfo(jellyfinTunerHostInfo).Execute()

Adds a tuner host.

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
	jellyfinTunerHostInfo := *openapiclient.NewJellyfinTunerHostInfo() // JellyfinTunerHostInfo | New tuner host. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvAPI.AddTunerHost(context.Background()).JellyfinTunerHostInfo(jellyfinTunerHostInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvAPI.AddTunerHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddTunerHost`: JellyfinJellyfinTunerHostInfo
	fmt.Fprintf(os.Stdout, "Response from `LiveTvAPI.AddTunerHost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddTunerHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jellyfinTunerHostInfo** | [**JellyfinTunerHostInfo**](JellyfinTunerHostInfo.md) | New tuner host. | 

### Return type

[**JellyfinJellyfinTunerHostInfo**](JellyfinTunerHostInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelSeriesTimer

> CancelSeriesTimer(ctx, timerId).Execute()

Cancels a live tv series timer.

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
	timerId := "timerId_example" // string | Timer id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveTvAPI.CancelSeriesTimer(context.Background(), timerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvAPI.CancelSeriesTimer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**timerId** | **string** | Timer id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelSeriesTimerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelTimer

> CancelTimer(ctx, timerId).Execute()

Cancels a live tv timer.

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
	timerId := "timerId_example" // string | Timer id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveTvAPI.CancelTimer(context.Background(), timerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvAPI.CancelTimer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**timerId** | **string** | Timer id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelTimerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSeriesTimer

> CreateSeriesTimer(ctx).JellyfinSeriesTimerInfoDto(jellyfinSeriesTimerInfoDto).Execute()

Creates a live tv series timer.

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
	jellyfinSeriesTimerInfoDto := *openapiclient.NewJellyfinSeriesTimerInfoDto() // JellyfinSeriesTimerInfoDto | New series timer info. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveTvAPI.CreateSeriesTimer(context.Background()).JellyfinSeriesTimerInfoDto(jellyfinSeriesTimerInfoDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvAPI.CreateSeriesTimer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSeriesTimerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jellyfinSeriesTimerInfoDto** | [**JellyfinSeriesTimerInfoDto**](JellyfinSeriesTimerInfoDto.md) | New series timer info. | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTimer

> CreateTimer(ctx).JellyfinTimerInfoDto(jellyfinTimerInfoDto).Execute()

Creates a live tv timer.

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
	jellyfinTimerInfoDto := *openapiclient.NewJellyfinTimerInfoDto() // JellyfinTimerInfoDto | New timer info. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveTvAPI.CreateTimer(context.Background()).JellyfinTimerInfoDto(jellyfinTimerInfoDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvAPI.CreateTimer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTimerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jellyfinTimerInfoDto** | [**JellyfinTimerInfoDto**](JellyfinTimerInfoDto.md) | New timer info. | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteListingProvider

> DeleteListingProvider(ctx).Id(id).Execute()

Delete listing provider.

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
	id := "id_example" // string | Listing provider id. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveTvAPI.DeleteListingProvider(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvAPI.DeleteListingProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteListingProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Listing provider id. | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRecording

> DeleteRecording(ctx, recordingId).Execute()

Deletes a live tv recording.

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
	recordingId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Recording id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveTvAPI.DeleteRecording(context.Background(), recordingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvAPI.DeleteRecording``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recordingId** | **string** | Recording id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRecordingRequest struct via the builder pattern


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


## DeleteTunerHost

> DeleteTunerHost(ctx).Id(id).Execute()

Deletes a tuner host.

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
	id := "id_example" // string | Tuner host id. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveTvAPI.DeleteTunerHost(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvAPI.DeleteTunerHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTunerHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Tuner host id. | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscoverTuners

> []JellyfinJellyfinTunerHostInfo DiscoverTuners(ctx).NewDevicesOnly(newDevicesOnly).Execute()

Discover tuners.

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
	newDevicesOnly := true // bool | Only discover new tuners. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvAPI.DiscoverTuners(context.Background()).NewDevicesOnly(newDevicesOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvAPI.DiscoverTuners``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscoverTuners`: []JellyfinJellyfinTunerHostInfo
	fmt.Fprintf(os.Stdout, "Response from `LiveTvAPI.DiscoverTuners`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscoverTunersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newDevicesOnly** | **bool** | Only discover new tuners. | [default to false]

### Return type

[**[]JellyfinJellyfinTunerHostInfo**](JellyfinTunerHostInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscvoverTuners

> []JellyfinJellyfinTunerHostInfo DiscvoverTuners(ctx).NewDevicesOnly(newDevicesOnly).Execute()

Discover tuners.

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
	newDevicesOnly := true // bool | Only discover new tuners. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvAPI.DiscvoverTuners(context.Background()).NewDevicesOnly(newDevicesOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvAPI.DiscvoverTuners``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscvoverTuners`: []JellyfinJellyfinTunerHostInfo
	fmt.Fprintf(os.Stdout, "Response from `LiveTvAPI.DiscvoverTuners`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscvoverTunersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newDevicesOnly** | **bool** | Only discover new tuners. | [default to false]

### Return type

[**[]JellyfinJellyfinTunerHostInfo**](JellyfinTunerHostInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannel

> JellyfinJellyfinBaseItemDto GetChannel(ctx, channelId).UserId(userId).Execute()

Gets a live tv channel.

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
	channelId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Channel id.
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional. Attach user data. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvAPI.GetChannel(context.Background(), channelId).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvAPI.GetChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChannel`: JellyfinJellyfinBaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `LiveTvAPI.GetChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** | Channel id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | Optional. Attach user data. | 

### Return type

[**JellyfinJellyfinBaseItemDto**](JellyfinBaseItemDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannelMappingOptions

> JellyfinJellyfinChannelMappingOptionsDto GetChannelMappingOptions(ctx).ProviderId(providerId).Execute()

Get channel mapping options.

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
	providerId := "providerId_example" // string | Provider id. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvAPI.GetChannelMappingOptions(context.Background()).ProviderId(providerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvAPI.GetChannelMappingOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChannelMappingOptions`: JellyfinJellyfinChannelMappingOptionsDto
	fmt.Fprintf(os.Stdout, "Response from `LiveTvAPI.GetChannelMappingOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelMappingOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerId** | **string** | Provider id. | 

### Return type

[**JellyfinJellyfinChannelMappingOptionsDto**](JellyfinChannelMappingOptionsDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefaultListingProvider

> JellyfinJellyfinListingsProviderInfo GetDefaultListingProvider(ctx).Execute()

Gets default listings provider info.

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
	resp, r, err := apiClient.LiveTvAPI.GetDefaultListingProvider(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvAPI.GetDefaultListingProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDefaultListingProvider`: JellyfinJellyfinListingsProviderInfo
	fmt.Fprintf(os.Stdout, "Response from `LiveTvAPI.GetDefaultListingProvider`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultListingProviderRequest struct via the builder pattern


### Return type

[**JellyfinJellyfinListingsProviderInfo**](JellyfinListingsProviderInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefaultTimer

> JellyfinJellyfinSeriesTimerInfoDto GetDefaultTimer(ctx).ProgramId(programId).Execute()

Gets the default values for a new timer.

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
	programId := "programId_example" // string | Optional. To attach default values based on a program. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvAPI.GetDefaultTimer(context.Background()).ProgramId(programId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvAPI.GetDefaultTimer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDefaultTimer`: JellyfinJellyfinSeriesTimerInfoDto
	fmt.Fprintf(os.Stdout, "Response from `LiveTvAPI.GetDefaultTimer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultTimerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **programId** | **string** | Optional. To attach default values based on a program. | 

### Return type

[**JellyfinJellyfinSeriesTimerInfoDto**](JellyfinSeriesTimerInfoDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGuideInfo

> JellyfinJellyfinGuideInfo GetGuideInfo(ctx).Execute()

Get guid info.

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
	resp, r, err := apiClient.LiveTvAPI.GetGuideInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvAPI.GetGuideInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGuideInfo`: JellyfinJellyfinGuideInfo
	fmt.Fprintf(os.Stdout, "Response from `LiveTvAPI.GetGuideInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGuideInfoRequest struct via the builder pattern


### Return type

[**JellyfinJellyfinGuideInfo**](JellyfinGuideInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLineups

> []JellyfinJellyfinNameIdPair GetLineups(ctx).Id(id).Type_(type_).Location(location).Country(country).Execute()

Gets available lineups.

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
	id := "id_example" // string | Provider id. (optional)
	type_ := "type__example" // string | Provider type. (optional)
	location := "location_example" // string | Location. (optional)
	country := "country_example" // string | Country. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvAPI.GetLineups(context.Background()).Id(id).Type_(type_).Location(location).Country(country).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvAPI.GetLineups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLineups`: []JellyfinJellyfinNameIdPair
	fmt.Fprintf(os.Stdout, "Response from `LiveTvAPI.GetLineups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLineupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Provider id. | 
 **type_** | **string** | Provider type. | 
 **location** | **string** | Location. | 
 **country** | **string** | Country. | 

### Return type

[**[]JellyfinJellyfinNameIdPair**](JellyfinNameIdPair.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLiveRecordingFile

> *os.File GetLiveRecordingFile(ctx, recordingId).Execute()

Gets a live tv recording stream.

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
	recordingId := "recordingId_example" // string | Recording id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvAPI.GetLiveRecordingFile(context.Background(), recordingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvAPI.GetLiveRecordingFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLiveRecordingFile`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `LiveTvAPI.GetLiveRecordingFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recordingId** | **string** | Recording id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLiveRecordingFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: video/*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLiveStreamFile

> *os.File GetLiveStreamFile(ctx, streamId, container).Execute()

Gets a live tv channel stream.

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
	streamId := "streamId_example" // string | Stream id.
	container := "container_example" // string | Container type.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvAPI.GetLiveStreamFile(context.Background(), streamId, container).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvAPI.GetLiveStreamFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLiveStreamFile`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `LiveTvAPI.GetLiveStreamFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **string** | Stream id. | 
**container** | **string** | Container type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLiveStreamFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: video/*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLiveTvChannels

> JellyfinJellyfinBaseItemDtoQueryResult GetLiveTvChannels(ctx).Type_(type_).UserId(userId).StartIndex(startIndex).IsMovie(isMovie).IsSeries(isSeries).IsNews(isNews).IsKids(isKids).IsSports(isSports).Limit(limit).IsFavorite(isFavorite).IsLiked(isLiked).IsDisliked(isDisliked).EnableImages(enableImages).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Fields(fields).EnableUserData(enableUserData).SortBy(sortBy).SortOrder(sortOrder).EnableFavoriteSorting(enableFavoriteSorting).AddCurrentProgram(addCurrentProgram).Execute()

Gets available live tv channels.

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
	type_ := "type__example" // JellyfinJellyfinChannelType | Optional. Filter by channel type. (optional)
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional. Filter by user and attach user data. (optional)
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	isMovie := true // bool | Optional. Filter for movies. (optional)
	isSeries := true // bool | Optional. Filter for series. (optional)
	isNews := true // bool | Optional. Filter for news. (optional)
	isKids := true // bool | Optional. Filter for kids. (optional)
	isSports := true // bool | Optional. Filter for sports. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return. (optional)
	isFavorite := true // bool | Optional. Filter by channels that are favorites, or not. (optional)
	isLiked := true // bool | Optional. Filter by channels that are liked, or not. (optional)
	isDisliked := true // bool | Optional. Filter by channels that are disliked, or not. (optional)
	enableImages := true // bool | Optional. Include image information in output. (optional)
	imageTypeLimit := int32(56) // int32 | Optional. The max number of images to return, per image type. (optional)
	enableImageTypes := []JellyfinJellyfinImageType{"TODO"} // []JellyfinJellyfinImageType | \"Optional. The image types to include in the output. (optional)
	fields := []JellyfinJellyfinItemFields{"TODO"} // []JellyfinJellyfinItemFields | Optional. Specify additional fields of information to return in the output. (optional)
	enableUserData := true // bool | Optional. Include user data. (optional)
	sortBy := []JellyfinJellyfinItemSortBy{"TODO"} // []JellyfinJellyfinItemSortBy | Optional. Key to sort by. (optional)
	sortOrder := "sortOrder_example" // JellyfinJellyfinSortOrder | Optional. Sort order. (optional)
	enableFavoriteSorting := true // bool | Optional. Incorporate favorite and like status into channel sorting. (optional) (default to false)
	addCurrentProgram := true // bool | Optional. Adds current program info to each channel. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvAPI.GetLiveTvChannels(context.Background()).Type_(type_).UserId(userId).StartIndex(startIndex).IsMovie(isMovie).IsSeries(isSeries).IsNews(isNews).IsKids(isKids).IsSports(isSports).Limit(limit).IsFavorite(isFavorite).IsLiked(isLiked).IsDisliked(isDisliked).EnableImages(enableImages).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Fields(fields).EnableUserData(enableUserData).SortBy(sortBy).SortOrder(sortOrder).EnableFavoriteSorting(enableFavoriteSorting).AddCurrentProgram(addCurrentProgram).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvAPI.GetLiveTvChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLiveTvChannels`: JellyfinJellyfinBaseItemDtoQueryResult
	fmt.Fprintf(os.Stdout, "Response from `LiveTvAPI.GetLiveTvChannels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLiveTvChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **JellyfinJellyfinChannelType** | Optional. Filter by channel type. | 
 **userId** | **string** | Optional. Filter by user and attach user data. | 
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **isMovie** | **bool** | Optional. Filter for movies. | 
 **isSeries** | **bool** | Optional. Filter for series. | 
 **isNews** | **bool** | Optional. Filter for news. | 
 **isKids** | **bool** | Optional. Filter for kids. | 
 **isSports** | **bool** | Optional. Filter for sports. | 
 **limit** | **int32** | Optional. The maximum number of records to return. | 
 **isFavorite** | **bool** | Optional. Filter by channels that are favorites, or not. | 
 **isLiked** | **bool** | Optional. Filter by channels that are liked, or not. | 
 **isDisliked** | **bool** | Optional. Filter by channels that are disliked, or not. | 
 **enableImages** | **bool** | Optional. Include image information in output. | 
 **imageTypeLimit** | **int32** | Optional. The max number of images to return, per image type. | 
 **enableImageTypes** | [**[]JellyfinJellyfinImageType**](JellyfinImageType.md) | \&quot;Optional. The image types to include in the output. | 
 **fields** | [**[]JellyfinJellyfinItemFields**](JellyfinItemFields.md) | Optional. Specify additional fields of information to return in the output. | 
 **enableUserData** | **bool** | Optional. Include user data. | 
 **sortBy** | [**[]JellyfinJellyfinItemSortBy**](JellyfinItemSortBy.md) | Optional. Key to sort by. | 
 **sortOrder** | **JellyfinJellyfinSortOrder** | Optional. Sort order. | 
 **enableFavoriteSorting** | **bool** | Optional. Incorporate favorite and like status into channel sorting. | [default to false]
 **addCurrentProgram** | **bool** | Optional. Adds current program info to each channel. | [default to true]

### Return type

[**JellyfinJellyfinBaseItemDtoQueryResult**](JellyfinBaseItemDtoQueryResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLiveTvInfo

> JellyfinJellyfinLiveTvInfo GetLiveTvInfo(ctx).Execute()

Gets available live tv services.

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
	resp, r, err := apiClient.LiveTvAPI.GetLiveTvInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvAPI.GetLiveTvInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLiveTvInfo`: JellyfinJellyfinLiveTvInfo
	fmt.Fprintf(os.Stdout, "Response from `LiveTvAPI.GetLiveTvInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLiveTvInfoRequest struct via the builder pattern


### Return type

[**JellyfinJellyfinLiveTvInfo**](JellyfinLiveTvInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLiveTvPrograms

> JellyfinJellyfinBaseItemDtoQueryResult GetLiveTvPrograms(ctx).ChannelIds(channelIds).UserId(userId).MinStartDate(minStartDate).HasAired(hasAired).IsAiring(isAiring).MaxStartDate(maxStartDate).MinEndDate(minEndDate).MaxEndDate(maxEndDate).IsMovie(isMovie).IsSeries(isSeries).IsNews(isNews).IsKids(isKids).IsSports(isSports).StartIndex(startIndex).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Genres(genres).GenreIds(genreIds).EnableImages(enableImages).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).EnableUserData(enableUserData).SeriesTimerId(seriesTimerId).LibrarySeriesId(librarySeriesId).Fields(fields).EnableTotalRecordCount(enableTotalRecordCount).Execute()

Gets available live tv epgs.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/shamelin/go-jellyfin-api-client"
)

func main() {
	channelIds := []string{"Inner_example"} // []string | The channels to return guide information for. (optional)
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional. Filter by user id. (optional)
	minStartDate := time.Now() // time.Time | Optional. The minimum premiere start date. (optional)
	hasAired := true // bool | Optional. Filter by programs that have completed airing, or not. (optional)
	isAiring := true // bool | Optional. Filter by programs that are currently airing, or not. (optional)
	maxStartDate := time.Now() // time.Time | Optional. The maximum premiere start date. (optional)
	minEndDate := time.Now() // time.Time | Optional. The minimum premiere end date. (optional)
	maxEndDate := time.Now() // time.Time | Optional. The maximum premiere end date. (optional)
	isMovie := true // bool | Optional. Filter for movies. (optional)
	isSeries := true // bool | Optional. Filter for series. (optional)
	isNews := true // bool | Optional. Filter for news. (optional)
	isKids := true // bool | Optional. Filter for kids. (optional)
	isSports := true // bool | Optional. Filter for sports. (optional)
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return. (optional)
	sortBy := []JellyfinJellyfinItemSortBy{"TODO"} // []JellyfinJellyfinItemSortBy | Optional. Specify one or more sort orders, comma delimited. Options: Name, StartDate. (optional)
	sortOrder := []JellyfinJellyfinSortOrder{"TODO"} // []JellyfinJellyfinSortOrder | Sort Order - Ascending,Descending. (optional)
	genres := []string{"Inner_example"} // []string | The genres to return guide information for. (optional)
	genreIds := []string{"Inner_example"} // []string | The genre ids to return guide information for. (optional)
	enableImages := true // bool | Optional. Include image information in output. (optional)
	imageTypeLimit := int32(56) // int32 | Optional. The max number of images to return, per image type. (optional)
	enableImageTypes := []JellyfinJellyfinImageType{"TODO"} // []JellyfinJellyfinImageType | Optional. The image types to include in the output. (optional)
	enableUserData := true // bool | Optional. Include user data. (optional)
	seriesTimerId := "seriesTimerId_example" // string | Optional. Filter by series timer id. (optional)
	librarySeriesId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional. Filter by library series id. (optional)
	fields := []JellyfinJellyfinItemFields{"TODO"} // []JellyfinJellyfinItemFields | Optional. Specify additional fields of information to return in the output. (optional)
	enableTotalRecordCount := true // bool | Retrieve total record count. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvAPI.GetLiveTvPrograms(context.Background()).ChannelIds(channelIds).UserId(userId).MinStartDate(minStartDate).HasAired(hasAired).IsAiring(isAiring).MaxStartDate(maxStartDate).MinEndDate(minEndDate).MaxEndDate(maxEndDate).IsMovie(isMovie).IsSeries(isSeries).IsNews(isNews).IsKids(isKids).IsSports(isSports).StartIndex(startIndex).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Genres(genres).GenreIds(genreIds).EnableImages(enableImages).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).EnableUserData(enableUserData).SeriesTimerId(seriesTimerId).LibrarySeriesId(librarySeriesId).Fields(fields).EnableTotalRecordCount(enableTotalRecordCount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvAPI.GetLiveTvPrograms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLiveTvPrograms`: JellyfinJellyfinBaseItemDtoQueryResult
	fmt.Fprintf(os.Stdout, "Response from `LiveTvAPI.GetLiveTvPrograms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLiveTvProgramsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelIds** | **[]string** | The channels to return guide information for. | 
 **userId** | **string** | Optional. Filter by user id. | 
 **minStartDate** | **time.Time** | Optional. The minimum premiere start date. | 
 **hasAired** | **bool** | Optional. Filter by programs that have completed airing, or not. | 
 **isAiring** | **bool** | Optional. Filter by programs that are currently airing, or not. | 
 **maxStartDate** | **time.Time** | Optional. The maximum premiere start date. | 
 **minEndDate** | **time.Time** | Optional. The minimum premiere end date. | 
 **maxEndDate** | **time.Time** | Optional. The maximum premiere end date. | 
 **isMovie** | **bool** | Optional. Filter for movies. | 
 **isSeries** | **bool** | Optional. Filter for series. | 
 **isNews** | **bool** | Optional. Filter for news. | 
 **isKids** | **bool** | Optional. Filter for kids. | 
 **isSports** | **bool** | Optional. Filter for sports. | 
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return. | 
 **sortBy** | [**[]JellyfinJellyfinItemSortBy**](JellyfinItemSortBy.md) | Optional. Specify one or more sort orders, comma delimited. Options: Name, StartDate. | 
 **sortOrder** | [**[]JellyfinJellyfinSortOrder**](JellyfinSortOrder.md) | Sort Order - Ascending,Descending. | 
 **genres** | **[]string** | The genres to return guide information for. | 
 **genreIds** | **[]string** | The genre ids to return guide information for. | 
 **enableImages** | **bool** | Optional. Include image information in output. | 
 **imageTypeLimit** | **int32** | Optional. The max number of images to return, per image type. | 
 **enableImageTypes** | [**[]JellyfinJellyfinImageType**](JellyfinImageType.md) | Optional. The image types to include in the output. | 
 **enableUserData** | **bool** | Optional. Include user data. | 
 **seriesTimerId** | **string** | Optional. Filter by series timer id. | 
 **librarySeriesId** | **string** | Optional. Filter by library series id. | 
 **fields** | [**[]JellyfinJellyfinItemFields**](JellyfinItemFields.md) | Optional. Specify additional fields of information to return in the output. | 
 **enableTotalRecordCount** | **bool** | Retrieve total record count. | [default to true]

### Return type

[**JellyfinJellyfinBaseItemDtoQueryResult**](JellyfinBaseItemDtoQueryResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProgram

> JellyfinJellyfinBaseItemDto GetProgram(ctx, programId).UserId(userId).Execute()

Gets a live tv program.

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
	programId := "programId_example" // string | Program id.
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional. Attach user data. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvAPI.GetProgram(context.Background(), programId).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvAPI.GetProgram``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProgram`: JellyfinJellyfinBaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `LiveTvAPI.GetProgram`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**programId** | **string** | Program id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProgramRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | Optional. Attach user data. | 

### Return type

[**JellyfinJellyfinBaseItemDto**](JellyfinBaseItemDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrograms

> JellyfinJellyfinBaseItemDtoQueryResult GetPrograms(ctx).JellyfinGetProgramsDto(jellyfinGetProgramsDto).Execute()

Gets available live tv epgs.

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
	jellyfinGetProgramsDto := *openapiclient.NewJellyfinGetProgramsDto() // JellyfinGetProgramsDto | Request body. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvAPI.GetPrograms(context.Background()).JellyfinGetProgramsDto(jellyfinGetProgramsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvAPI.GetPrograms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrograms`: JellyfinJellyfinBaseItemDtoQueryResult
	fmt.Fprintf(os.Stdout, "Response from `LiveTvAPI.GetPrograms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProgramsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jellyfinGetProgramsDto** | [**JellyfinGetProgramsDto**](JellyfinGetProgramsDto.md) | Request body. | 

### Return type

[**JellyfinJellyfinBaseItemDtoQueryResult**](JellyfinBaseItemDtoQueryResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecommendedPrograms

> JellyfinJellyfinBaseItemDtoQueryResult GetRecommendedPrograms(ctx).UserId(userId).Limit(limit).IsAiring(isAiring).HasAired(hasAired).IsSeries(isSeries).IsMovie(isMovie).IsNews(isNews).IsKids(isKids).IsSports(isSports).EnableImages(enableImages).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).GenreIds(genreIds).Fields(fields).EnableUserData(enableUserData).EnableTotalRecordCount(enableTotalRecordCount).Execute()

Gets recommended live tv epgs.

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional. filter by user id. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return. (optional)
	isAiring := true // bool | Optional. Filter by programs that are currently airing, or not. (optional)
	hasAired := true // bool | Optional. Filter by programs that have completed airing, or not. (optional)
	isSeries := true // bool | Optional. Filter for series. (optional)
	isMovie := true // bool | Optional. Filter for movies. (optional)
	isNews := true // bool | Optional. Filter for news. (optional)
	isKids := true // bool | Optional. Filter for kids. (optional)
	isSports := true // bool | Optional. Filter for sports. (optional)
	enableImages := true // bool | Optional. Include image information in output. (optional)
	imageTypeLimit := int32(56) // int32 | Optional. The max number of images to return, per image type. (optional)
	enableImageTypes := []JellyfinJellyfinImageType{"TODO"} // []JellyfinJellyfinImageType | Optional. The image types to include in the output. (optional)
	genreIds := []string{"Inner_example"} // []string | The genres to return guide information for. (optional)
	fields := []JellyfinJellyfinItemFields{"TODO"} // []JellyfinJellyfinItemFields | Optional. Specify additional fields of information to return in the output. (optional)
	enableUserData := true // bool | Optional. include user data. (optional)
	enableTotalRecordCount := true // bool | Retrieve total record count. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvAPI.GetRecommendedPrograms(context.Background()).UserId(userId).Limit(limit).IsAiring(isAiring).HasAired(hasAired).IsSeries(isSeries).IsMovie(isMovie).IsNews(isNews).IsKids(isKids).IsSports(isSports).EnableImages(enableImages).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).GenreIds(genreIds).Fields(fields).EnableUserData(enableUserData).EnableTotalRecordCount(enableTotalRecordCount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvAPI.GetRecommendedPrograms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRecommendedPrograms`: JellyfinJellyfinBaseItemDtoQueryResult
	fmt.Fprintf(os.Stdout, "Response from `LiveTvAPI.GetRecommendedPrograms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRecommendedProgramsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | Optional. filter by user id. | 
 **limit** | **int32** | Optional. The maximum number of records to return. | 
 **isAiring** | **bool** | Optional. Filter by programs that are currently airing, or not. | 
 **hasAired** | **bool** | Optional. Filter by programs that have completed airing, or not. | 
 **isSeries** | **bool** | Optional. Filter for series. | 
 **isMovie** | **bool** | Optional. Filter for movies. | 
 **isNews** | **bool** | Optional. Filter for news. | 
 **isKids** | **bool** | Optional. Filter for kids. | 
 **isSports** | **bool** | Optional. Filter for sports. | 
 **enableImages** | **bool** | Optional. Include image information in output. | 
 **imageTypeLimit** | **int32** | Optional. The max number of images to return, per image type. | 
 **enableImageTypes** | [**[]JellyfinJellyfinImageType**](JellyfinImageType.md) | Optional. The image types to include in the output. | 
 **genreIds** | **[]string** | The genres to return guide information for. | 
 **fields** | [**[]JellyfinJellyfinItemFields**](JellyfinItemFields.md) | Optional. Specify additional fields of information to return in the output. | 
 **enableUserData** | **bool** | Optional. include user data. | 
 **enableTotalRecordCount** | **bool** | Retrieve total record count. | [default to true]

### Return type

[**JellyfinJellyfinBaseItemDtoQueryResult**](JellyfinBaseItemDtoQueryResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecording

> JellyfinJellyfinBaseItemDto GetRecording(ctx, recordingId).UserId(userId).Execute()

Gets a live tv recording.

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
	recordingId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Recording id.
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional. Attach user data. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvAPI.GetRecording(context.Background(), recordingId).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvAPI.GetRecording``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRecording`: JellyfinJellyfinBaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `LiveTvAPI.GetRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recordingId** | **string** | Recording id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecordingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | Optional. Attach user data. | 

### Return type

[**JellyfinJellyfinBaseItemDto**](JellyfinBaseItemDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecordingFolders

> JellyfinJellyfinBaseItemDtoQueryResult GetRecordingFolders(ctx).UserId(userId).Execute()

Gets recording folders.

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional. Filter by user and attach user data. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvAPI.GetRecordingFolders(context.Background()).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvAPI.GetRecordingFolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRecordingFolders`: JellyfinJellyfinBaseItemDtoQueryResult
	fmt.Fprintf(os.Stdout, "Response from `LiveTvAPI.GetRecordingFolders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRecordingFoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | Optional. Filter by user and attach user data. | 

### Return type

[**JellyfinJellyfinBaseItemDtoQueryResult**](JellyfinBaseItemDtoQueryResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecordingGroup

> GetRecordingGroup(ctx, groupId).Execute()

Get recording group.

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
	groupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Group id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveTvAPI.GetRecordingGroup(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvAPI.GetRecordingGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecordingGroupRequest struct via the builder pattern


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


## GetRecordingGroups

> JellyfinJellyfinBaseItemDtoQueryResult GetRecordingGroups(ctx).UserId(userId).Execute()

Gets live tv recording groups.

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional. Filter by user and attach user data. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvAPI.GetRecordingGroups(context.Background()).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvAPI.GetRecordingGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRecordingGroups`: JellyfinJellyfinBaseItemDtoQueryResult
	fmt.Fprintf(os.Stdout, "Response from `LiveTvAPI.GetRecordingGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRecordingGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | Optional. Filter by user and attach user data. | 

### Return type

[**JellyfinJellyfinBaseItemDtoQueryResult**](JellyfinBaseItemDtoQueryResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecordings

> JellyfinJellyfinBaseItemDtoQueryResult GetRecordings(ctx).ChannelId(channelId).UserId(userId).StartIndex(startIndex).Limit(limit).Status(status).IsInProgress(isInProgress).SeriesTimerId(seriesTimerId).EnableImages(enableImages).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Fields(fields).EnableUserData(enableUserData).IsMovie(isMovie).IsSeries(isSeries).IsKids(isKids).IsSports(isSports).IsNews(isNews).IsLibraryItem(isLibraryItem).EnableTotalRecordCount(enableTotalRecordCount).Execute()

Gets live tv recordings.

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
	channelId := "channelId_example" // string | Optional. Filter by channel id. (optional)
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional. Filter by user and attach user data. (optional)
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return. (optional)
	status := "status_example" // JellyfinJellyfinRecordingStatus | Optional. Filter by recording status. (optional)
	isInProgress := true // bool | Optional. Filter by recordings that are in progress, or not. (optional)
	seriesTimerId := "seriesTimerId_example" // string | Optional. Filter by recordings belonging to a series timer. (optional)
	enableImages := true // bool | Optional. Include image information in output. (optional)
	imageTypeLimit := int32(56) // int32 | Optional. The max number of images to return, per image type. (optional)
	enableImageTypes := []JellyfinJellyfinImageType{"TODO"} // []JellyfinJellyfinImageType | Optional. The image types to include in the output. (optional)
	fields := []JellyfinJellyfinItemFields{"TODO"} // []JellyfinJellyfinItemFields | Optional. Specify additional fields of information to return in the output. (optional)
	enableUserData := true // bool | Optional. Include user data. (optional)
	isMovie := true // bool | Optional. Filter for movies. (optional)
	isSeries := true // bool | Optional. Filter for series. (optional)
	isKids := true // bool | Optional. Filter for kids. (optional)
	isSports := true // bool | Optional. Filter for sports. (optional)
	isNews := true // bool | Optional. Filter for news. (optional)
	isLibraryItem := true // bool | Optional. Filter for is library item. (optional)
	enableTotalRecordCount := true // bool | Optional. Return total record count. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvAPI.GetRecordings(context.Background()).ChannelId(channelId).UserId(userId).StartIndex(startIndex).Limit(limit).Status(status).IsInProgress(isInProgress).SeriesTimerId(seriesTimerId).EnableImages(enableImages).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Fields(fields).EnableUserData(enableUserData).IsMovie(isMovie).IsSeries(isSeries).IsKids(isKids).IsSports(isSports).IsNews(isNews).IsLibraryItem(isLibraryItem).EnableTotalRecordCount(enableTotalRecordCount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvAPI.GetRecordings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRecordings`: JellyfinJellyfinBaseItemDtoQueryResult
	fmt.Fprintf(os.Stdout, "Response from `LiveTvAPI.GetRecordings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRecordingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelId** | **string** | Optional. Filter by channel id. | 
 **userId** | **string** | Optional. Filter by user and attach user data. | 
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return. | 
 **status** | **JellyfinJellyfinRecordingStatus** | Optional. Filter by recording status. | 
 **isInProgress** | **bool** | Optional. Filter by recordings that are in progress, or not. | 
 **seriesTimerId** | **string** | Optional. Filter by recordings belonging to a series timer. | 
 **enableImages** | **bool** | Optional. Include image information in output. | 
 **imageTypeLimit** | **int32** | Optional. The max number of images to return, per image type. | 
 **enableImageTypes** | [**[]JellyfinJellyfinImageType**](JellyfinImageType.md) | Optional. The image types to include in the output. | 
 **fields** | [**[]JellyfinJellyfinItemFields**](JellyfinItemFields.md) | Optional. Specify additional fields of information to return in the output. | 
 **enableUserData** | **bool** | Optional. Include user data. | 
 **isMovie** | **bool** | Optional. Filter for movies. | 
 **isSeries** | **bool** | Optional. Filter for series. | 
 **isKids** | **bool** | Optional. Filter for kids. | 
 **isSports** | **bool** | Optional. Filter for sports. | 
 **isNews** | **bool** | Optional. Filter for news. | 
 **isLibraryItem** | **bool** | Optional. Filter for is library item. | 
 **enableTotalRecordCount** | **bool** | Optional. Return total record count. | [default to true]

### Return type

[**JellyfinJellyfinBaseItemDtoQueryResult**](JellyfinBaseItemDtoQueryResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecordingsSeries

> JellyfinJellyfinBaseItemDtoQueryResult GetRecordingsSeries(ctx).ChannelId(channelId).UserId(userId).GroupId(groupId).StartIndex(startIndex).Limit(limit).Status(status).IsInProgress(isInProgress).SeriesTimerId(seriesTimerId).EnableImages(enableImages).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Fields(fields).EnableUserData(enableUserData).EnableTotalRecordCount(enableTotalRecordCount).Execute()

Gets live tv recording series.

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
	channelId := "channelId_example" // string | Optional. Filter by channel id. (optional)
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional. Filter by user and attach user data. (optional)
	groupId := "groupId_example" // string | Optional. Filter by recording group. (optional)
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return. (optional)
	status := "status_example" // JellyfinJellyfinRecordingStatus | Optional. Filter by recording status. (optional)
	isInProgress := true // bool | Optional. Filter by recordings that are in progress, or not. (optional)
	seriesTimerId := "seriesTimerId_example" // string | Optional. Filter by recordings belonging to a series timer. (optional)
	enableImages := true // bool | Optional. Include image information in output. (optional)
	imageTypeLimit := int32(56) // int32 | Optional. The max number of images to return, per image type. (optional)
	enableImageTypes := []JellyfinJellyfinImageType{"TODO"} // []JellyfinJellyfinImageType | Optional. The image types to include in the output. (optional)
	fields := []JellyfinJellyfinItemFields{"TODO"} // []JellyfinJellyfinItemFields | Optional. Specify additional fields of information to return in the output. (optional)
	enableUserData := true // bool | Optional. Include user data. (optional)
	enableTotalRecordCount := true // bool | Optional. Return total record count. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvAPI.GetRecordingsSeries(context.Background()).ChannelId(channelId).UserId(userId).GroupId(groupId).StartIndex(startIndex).Limit(limit).Status(status).IsInProgress(isInProgress).SeriesTimerId(seriesTimerId).EnableImages(enableImages).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Fields(fields).EnableUserData(enableUserData).EnableTotalRecordCount(enableTotalRecordCount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvAPI.GetRecordingsSeries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRecordingsSeries`: JellyfinJellyfinBaseItemDtoQueryResult
	fmt.Fprintf(os.Stdout, "Response from `LiveTvAPI.GetRecordingsSeries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRecordingsSeriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelId** | **string** | Optional. Filter by channel id. | 
 **userId** | **string** | Optional. Filter by user and attach user data. | 
 **groupId** | **string** | Optional. Filter by recording group. | 
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return. | 
 **status** | **JellyfinJellyfinRecordingStatus** | Optional. Filter by recording status. | 
 **isInProgress** | **bool** | Optional. Filter by recordings that are in progress, or not. | 
 **seriesTimerId** | **string** | Optional. Filter by recordings belonging to a series timer. | 
 **enableImages** | **bool** | Optional. Include image information in output. | 
 **imageTypeLimit** | **int32** | Optional. The max number of images to return, per image type. | 
 **enableImageTypes** | [**[]JellyfinJellyfinImageType**](JellyfinImageType.md) | Optional. The image types to include in the output. | 
 **fields** | [**[]JellyfinJellyfinItemFields**](JellyfinItemFields.md) | Optional. Specify additional fields of information to return in the output. | 
 **enableUserData** | **bool** | Optional. Include user data. | 
 **enableTotalRecordCount** | **bool** | Optional. Return total record count. | [default to true]

### Return type

[**JellyfinJellyfinBaseItemDtoQueryResult**](JellyfinBaseItemDtoQueryResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSchedulesDirectCountries

> *os.File GetSchedulesDirectCountries(ctx).Execute()

Gets available countries.

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
	resp, r, err := apiClient.LiveTvAPI.GetSchedulesDirectCountries(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvAPI.GetSchedulesDirectCountries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSchedulesDirectCountries`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `LiveTvAPI.GetSchedulesDirectCountries`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSchedulesDirectCountriesRequest struct via the builder pattern


### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSeriesTimer

> JellyfinJellyfinSeriesTimerInfoDto GetSeriesTimer(ctx, timerId).Execute()

Gets a live tv series timer.

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
	timerId := "timerId_example" // string | Timer id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvAPI.GetSeriesTimer(context.Background(), timerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvAPI.GetSeriesTimer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSeriesTimer`: JellyfinJellyfinSeriesTimerInfoDto
	fmt.Fprintf(os.Stdout, "Response from `LiveTvAPI.GetSeriesTimer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**timerId** | **string** | Timer id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSeriesTimerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JellyfinJellyfinSeriesTimerInfoDto**](JellyfinSeriesTimerInfoDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSeriesTimers

> JellyfinJellyfinSeriesTimerInfoDtoQueryResult GetSeriesTimers(ctx).SortBy(sortBy).SortOrder(sortOrder).Execute()

Gets live tv series timers.

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
	sortBy := "sortBy_example" // string | Optional. Sort by SortName or Priority. (optional)
	sortOrder := "sortOrder_example" // JellyfinJellyfinSortOrder | Optional. Sort in Ascending or Descending order. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvAPI.GetSeriesTimers(context.Background()).SortBy(sortBy).SortOrder(sortOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvAPI.GetSeriesTimers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSeriesTimers`: JellyfinJellyfinSeriesTimerInfoDtoQueryResult
	fmt.Fprintf(os.Stdout, "Response from `LiveTvAPI.GetSeriesTimers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSeriesTimersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sortBy** | **string** | Optional. Sort by SortName or Priority. | 
 **sortOrder** | **JellyfinJellyfinSortOrder** | Optional. Sort in Ascending or Descending order. | 

### Return type

[**JellyfinJellyfinSeriesTimerInfoDtoQueryResult**](JellyfinSeriesTimerInfoDtoQueryResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTimer

> JellyfinJellyfinTimerInfoDto GetTimer(ctx, timerId).Execute()

Gets a timer.

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
	timerId := "timerId_example" // string | Timer id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvAPI.GetTimer(context.Background(), timerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvAPI.GetTimer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTimer`: JellyfinJellyfinTimerInfoDto
	fmt.Fprintf(os.Stdout, "Response from `LiveTvAPI.GetTimer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**timerId** | **string** | Timer id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTimerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JellyfinJellyfinTimerInfoDto**](JellyfinTimerInfoDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTimers

> JellyfinJellyfinTimerInfoDtoQueryResult GetTimers(ctx).ChannelId(channelId).SeriesTimerId(seriesTimerId).IsActive(isActive).IsScheduled(isScheduled).Execute()

Gets the live tv timers.

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
	channelId := "channelId_example" // string | Optional. Filter by channel id. (optional)
	seriesTimerId := "seriesTimerId_example" // string | Optional. Filter by timers belonging to a series timer. (optional)
	isActive := true // bool | Optional. Filter by timers that are active. (optional)
	isScheduled := true // bool | Optional. Filter by timers that are scheduled. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvAPI.GetTimers(context.Background()).ChannelId(channelId).SeriesTimerId(seriesTimerId).IsActive(isActive).IsScheduled(isScheduled).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvAPI.GetTimers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTimers`: JellyfinJellyfinTimerInfoDtoQueryResult
	fmt.Fprintf(os.Stdout, "Response from `LiveTvAPI.GetTimers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTimersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelId** | **string** | Optional. Filter by channel id. | 
 **seriesTimerId** | **string** | Optional. Filter by timers belonging to a series timer. | 
 **isActive** | **bool** | Optional. Filter by timers that are active. | 
 **isScheduled** | **bool** | Optional. Filter by timers that are scheduled. | 

### Return type

[**JellyfinJellyfinTimerInfoDtoQueryResult**](JellyfinTimerInfoDtoQueryResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTunerHostTypes

> []JellyfinJellyfinNameIdPair GetTunerHostTypes(ctx).Execute()

Get tuner host types.

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
	resp, r, err := apiClient.LiveTvAPI.GetTunerHostTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvAPI.GetTunerHostTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTunerHostTypes`: []JellyfinJellyfinNameIdPair
	fmt.Fprintf(os.Stdout, "Response from `LiveTvAPI.GetTunerHostTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTunerHostTypesRequest struct via the builder pattern


### Return type

[**[]JellyfinJellyfinNameIdPair**](JellyfinNameIdPair.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetTuner

> ResetTuner(ctx, tunerId).Execute()

Resets a tv tuner.

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
	tunerId := "tunerId_example" // string | Tuner id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveTvAPI.ResetTuner(context.Background(), tunerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvAPI.ResetTuner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tunerId** | **string** | Tuner id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetTunerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetChannelMapping

> JellyfinJellyfinTunerChannelMapping SetChannelMapping(ctx).JellyfinSetChannelMappingDto(jellyfinSetChannelMappingDto).Execute()

Set channel mappings.

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
	jellyfinSetChannelMappingDto := *openapiclient.NewJellyfinSetChannelMappingDto("ProviderId_example", "TunerChannelId_example", "ProviderChannelId_example") // JellyfinSetChannelMappingDto | The set channel mapping dto.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvAPI.SetChannelMapping(context.Background()).JellyfinSetChannelMappingDto(jellyfinSetChannelMappingDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvAPI.SetChannelMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetChannelMapping`: JellyfinJellyfinTunerChannelMapping
	fmt.Fprintf(os.Stdout, "Response from `LiveTvAPI.SetChannelMapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetChannelMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jellyfinSetChannelMappingDto** | [**JellyfinSetChannelMappingDto**](JellyfinSetChannelMappingDto.md) | The set channel mapping dto. | 

### Return type

[**JellyfinJellyfinTunerChannelMapping**](JellyfinTunerChannelMapping.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSeriesTimer

> UpdateSeriesTimer(ctx, timerId).JellyfinSeriesTimerInfoDto(jellyfinSeriesTimerInfoDto).Execute()

Updates a live tv series timer.

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
	timerId := "timerId_example" // string | Timer id.
	jellyfinSeriesTimerInfoDto := *openapiclient.NewJellyfinSeriesTimerInfoDto() // JellyfinSeriesTimerInfoDto | New series timer info. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveTvAPI.UpdateSeriesTimer(context.Background(), timerId).JellyfinSeriesTimerInfoDto(jellyfinSeriesTimerInfoDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvAPI.UpdateSeriesTimer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**timerId** | **string** | Timer id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSeriesTimerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jellyfinSeriesTimerInfoDto** | [**JellyfinSeriesTimerInfoDto**](JellyfinSeriesTimerInfoDto.md) | New series timer info. | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTimer

> UpdateTimer(ctx, timerId).JellyfinTimerInfoDto(jellyfinTimerInfoDto).Execute()

Updates a live tv timer.

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
	timerId := "timerId_example" // string | Timer id.
	jellyfinTimerInfoDto := *openapiclient.NewJellyfinTimerInfoDto() // JellyfinTimerInfoDto | New timer info. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveTvAPI.UpdateTimer(context.Background(), timerId).JellyfinTimerInfoDto(jellyfinTimerInfoDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvAPI.UpdateTimer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**timerId** | **string** | Timer id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTimerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jellyfinTimerInfoDto** | [**JellyfinTimerInfoDto**](JellyfinTimerInfoDto.md) | New timer info. | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

