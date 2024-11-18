# \ItemLookupAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplySearchCriteria**](ItemLookupAPI.md#ApplySearchCriteria) | **Post** /Items/RemoteSearch/Apply/{itemId} | Applies search criteria to an item and refreshes metadata.
[**GetBookRemoteSearchResults**](ItemLookupAPI.md#GetBookRemoteSearchResults) | **Post** /Items/RemoteSearch/Book | Get book remote search.
[**GetBoxSetRemoteSearchResults**](ItemLookupAPI.md#GetBoxSetRemoteSearchResults) | **Post** /Items/RemoteSearch/BoxSet | Get box set remote search.
[**GetExternalIdInfos**](ItemLookupAPI.md#GetExternalIdInfos) | **Get** /Items/{itemId}/ExternalIdInfos | Get the item&#39;s external id info.
[**GetMovieRemoteSearchResults**](ItemLookupAPI.md#GetMovieRemoteSearchResults) | **Post** /Items/RemoteSearch/Movie | Get movie remote search.
[**GetMusicAlbumRemoteSearchResults**](ItemLookupAPI.md#GetMusicAlbumRemoteSearchResults) | **Post** /Items/RemoteSearch/MusicAlbum | Get music album remote search.
[**GetMusicArtistRemoteSearchResults**](ItemLookupAPI.md#GetMusicArtistRemoteSearchResults) | **Post** /Items/RemoteSearch/MusicArtist | Get music artist remote search.
[**GetMusicVideoRemoteSearchResults**](ItemLookupAPI.md#GetMusicVideoRemoteSearchResults) | **Post** /Items/RemoteSearch/MusicVideo | Get music video remote search.
[**GetPersonRemoteSearchResults**](ItemLookupAPI.md#GetPersonRemoteSearchResults) | **Post** /Items/RemoteSearch/Person | Get person remote search.
[**GetSeriesRemoteSearchResults**](ItemLookupAPI.md#GetSeriesRemoteSearchResults) | **Post** /Items/RemoteSearch/Series | Get series remote search.
[**GetTrailerRemoteSearchResults**](ItemLookupAPI.md#GetTrailerRemoteSearchResults) | **Post** /Items/RemoteSearch/Trailer | Get trailer remote search.



## ApplySearchCriteria

> ApplySearchCriteria(ctx, itemId).JellyfinRemoteSearchResult(jellyfinRemoteSearchResult).ReplaceAllImages(replaceAllImages).Execute()

Applies search criteria to an item and refreshes metadata.

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
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Item id.
	jellyfinRemoteSearchResult := *openapiclient.NewJellyfinRemoteSearchResult() // JellyfinRemoteSearchResult | The remote search result.
	replaceAllImages := true // bool | Optional. Whether or not to replace all images. Default: True. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ItemLookupAPI.ApplySearchCriteria(context.Background(), itemId).JellyfinRemoteSearchResult(jellyfinRemoteSearchResult).ReplaceAllImages(replaceAllImages).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupAPI.ApplySearchCriteria``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | Item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplySearchCriteriaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jellyfinRemoteSearchResult** | [**JellyfinRemoteSearchResult**](JellyfinRemoteSearchResult.md) | The remote search result. | 
 **replaceAllImages** | **bool** | Optional. Whether or not to replace all images. Default: True. | [default to true]

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


## GetBookRemoteSearchResults

> []JellyfinJellyfinRemoteSearchResult GetBookRemoteSearchResults(ctx).JellyfinBookInfoRemoteSearchQuery(jellyfinBookInfoRemoteSearchQuery).Execute()

Get book remote search.

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
	jellyfinBookInfoRemoteSearchQuery := *openapiclient.NewJellyfinBookInfoRemoteSearchQuery() // JellyfinBookInfoRemoteSearchQuery | Remote search query.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemLookupAPI.GetBookRemoteSearchResults(context.Background()).JellyfinBookInfoRemoteSearchQuery(jellyfinBookInfoRemoteSearchQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupAPI.GetBookRemoteSearchResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBookRemoteSearchResults`: []JellyfinJellyfinRemoteSearchResult
	fmt.Fprintf(os.Stdout, "Response from `ItemLookupAPI.GetBookRemoteSearchResults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBookRemoteSearchResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jellyfinBookInfoRemoteSearchQuery** | [**JellyfinBookInfoRemoteSearchQuery**](JellyfinBookInfoRemoteSearchQuery.md) | Remote search query. | 

### Return type

[**[]JellyfinJellyfinRemoteSearchResult**](JellyfinRemoteSearchResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBoxSetRemoteSearchResults

> []JellyfinJellyfinRemoteSearchResult GetBoxSetRemoteSearchResults(ctx).JellyfinBoxSetInfoRemoteSearchQuery(jellyfinBoxSetInfoRemoteSearchQuery).Execute()

Get box set remote search.

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
	jellyfinBoxSetInfoRemoteSearchQuery := *openapiclient.NewJellyfinBoxSetInfoRemoteSearchQuery() // JellyfinBoxSetInfoRemoteSearchQuery | Remote search query.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemLookupAPI.GetBoxSetRemoteSearchResults(context.Background()).JellyfinBoxSetInfoRemoteSearchQuery(jellyfinBoxSetInfoRemoteSearchQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupAPI.GetBoxSetRemoteSearchResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBoxSetRemoteSearchResults`: []JellyfinJellyfinRemoteSearchResult
	fmt.Fprintf(os.Stdout, "Response from `ItemLookupAPI.GetBoxSetRemoteSearchResults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBoxSetRemoteSearchResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jellyfinBoxSetInfoRemoteSearchQuery** | [**JellyfinBoxSetInfoRemoteSearchQuery**](JellyfinBoxSetInfoRemoteSearchQuery.md) | Remote search query. | 

### Return type

[**[]JellyfinJellyfinRemoteSearchResult**](JellyfinRemoteSearchResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExternalIdInfos

> []JellyfinJellyfinExternalIdInfo GetExternalIdInfos(ctx, itemId).Execute()

Get the item's external id info.

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
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Item id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemLookupAPI.GetExternalIdInfos(context.Background(), itemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupAPI.GetExternalIdInfos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExternalIdInfos`: []JellyfinJellyfinExternalIdInfo
	fmt.Fprintf(os.Stdout, "Response from `ItemLookupAPI.GetExternalIdInfos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | Item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalIdInfosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]JellyfinJellyfinExternalIdInfo**](JellyfinExternalIdInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMovieRemoteSearchResults

> []JellyfinJellyfinRemoteSearchResult GetMovieRemoteSearchResults(ctx).JellyfinMovieInfoRemoteSearchQuery(jellyfinMovieInfoRemoteSearchQuery).Execute()

Get movie remote search.

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
	jellyfinMovieInfoRemoteSearchQuery := *openapiclient.NewJellyfinMovieInfoRemoteSearchQuery() // JellyfinMovieInfoRemoteSearchQuery | Remote search query.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemLookupAPI.GetMovieRemoteSearchResults(context.Background()).JellyfinMovieInfoRemoteSearchQuery(jellyfinMovieInfoRemoteSearchQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupAPI.GetMovieRemoteSearchResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMovieRemoteSearchResults`: []JellyfinJellyfinRemoteSearchResult
	fmt.Fprintf(os.Stdout, "Response from `ItemLookupAPI.GetMovieRemoteSearchResults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMovieRemoteSearchResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jellyfinMovieInfoRemoteSearchQuery** | [**JellyfinMovieInfoRemoteSearchQuery**](JellyfinMovieInfoRemoteSearchQuery.md) | Remote search query. | 

### Return type

[**[]JellyfinJellyfinRemoteSearchResult**](JellyfinRemoteSearchResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMusicAlbumRemoteSearchResults

> []JellyfinJellyfinRemoteSearchResult GetMusicAlbumRemoteSearchResults(ctx).JellyfinAlbumInfoRemoteSearchQuery(jellyfinAlbumInfoRemoteSearchQuery).Execute()

Get music album remote search.

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
	jellyfinAlbumInfoRemoteSearchQuery := *openapiclient.NewJellyfinAlbumInfoRemoteSearchQuery() // JellyfinAlbumInfoRemoteSearchQuery | Remote search query.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemLookupAPI.GetMusicAlbumRemoteSearchResults(context.Background()).JellyfinAlbumInfoRemoteSearchQuery(jellyfinAlbumInfoRemoteSearchQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupAPI.GetMusicAlbumRemoteSearchResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMusicAlbumRemoteSearchResults`: []JellyfinJellyfinRemoteSearchResult
	fmt.Fprintf(os.Stdout, "Response from `ItemLookupAPI.GetMusicAlbumRemoteSearchResults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMusicAlbumRemoteSearchResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jellyfinAlbumInfoRemoteSearchQuery** | [**JellyfinAlbumInfoRemoteSearchQuery**](JellyfinAlbumInfoRemoteSearchQuery.md) | Remote search query. | 

### Return type

[**[]JellyfinJellyfinRemoteSearchResult**](JellyfinRemoteSearchResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMusicArtistRemoteSearchResults

> []JellyfinJellyfinRemoteSearchResult GetMusicArtistRemoteSearchResults(ctx).JellyfinArtistInfoRemoteSearchQuery(jellyfinArtistInfoRemoteSearchQuery).Execute()

Get music artist remote search.

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
	jellyfinArtistInfoRemoteSearchQuery := *openapiclient.NewJellyfinArtistInfoRemoteSearchQuery() // JellyfinArtistInfoRemoteSearchQuery | Remote search query.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemLookupAPI.GetMusicArtistRemoteSearchResults(context.Background()).JellyfinArtistInfoRemoteSearchQuery(jellyfinArtistInfoRemoteSearchQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupAPI.GetMusicArtistRemoteSearchResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMusicArtistRemoteSearchResults`: []JellyfinJellyfinRemoteSearchResult
	fmt.Fprintf(os.Stdout, "Response from `ItemLookupAPI.GetMusicArtistRemoteSearchResults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMusicArtistRemoteSearchResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jellyfinArtistInfoRemoteSearchQuery** | [**JellyfinArtistInfoRemoteSearchQuery**](JellyfinArtistInfoRemoteSearchQuery.md) | Remote search query. | 

### Return type

[**[]JellyfinJellyfinRemoteSearchResult**](JellyfinRemoteSearchResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMusicVideoRemoteSearchResults

> []JellyfinJellyfinRemoteSearchResult GetMusicVideoRemoteSearchResults(ctx).JellyfinMusicVideoInfoRemoteSearchQuery(jellyfinMusicVideoInfoRemoteSearchQuery).Execute()

Get music video remote search.

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
	jellyfinMusicVideoInfoRemoteSearchQuery := *openapiclient.NewJellyfinMusicVideoInfoRemoteSearchQuery() // JellyfinMusicVideoInfoRemoteSearchQuery | Remote search query.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemLookupAPI.GetMusicVideoRemoteSearchResults(context.Background()).JellyfinMusicVideoInfoRemoteSearchQuery(jellyfinMusicVideoInfoRemoteSearchQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupAPI.GetMusicVideoRemoteSearchResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMusicVideoRemoteSearchResults`: []JellyfinJellyfinRemoteSearchResult
	fmt.Fprintf(os.Stdout, "Response from `ItemLookupAPI.GetMusicVideoRemoteSearchResults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMusicVideoRemoteSearchResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jellyfinMusicVideoInfoRemoteSearchQuery** | [**JellyfinMusicVideoInfoRemoteSearchQuery**](JellyfinMusicVideoInfoRemoteSearchQuery.md) | Remote search query. | 

### Return type

[**[]JellyfinJellyfinRemoteSearchResult**](JellyfinRemoteSearchResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPersonRemoteSearchResults

> []JellyfinJellyfinRemoteSearchResult GetPersonRemoteSearchResults(ctx).JellyfinPersonLookupInfoRemoteSearchQuery(jellyfinPersonLookupInfoRemoteSearchQuery).Execute()

Get person remote search.

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
	jellyfinPersonLookupInfoRemoteSearchQuery := *openapiclient.NewJellyfinPersonLookupInfoRemoteSearchQuery() // JellyfinPersonLookupInfoRemoteSearchQuery | Remote search query.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemLookupAPI.GetPersonRemoteSearchResults(context.Background()).JellyfinPersonLookupInfoRemoteSearchQuery(jellyfinPersonLookupInfoRemoteSearchQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupAPI.GetPersonRemoteSearchResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPersonRemoteSearchResults`: []JellyfinJellyfinRemoteSearchResult
	fmt.Fprintf(os.Stdout, "Response from `ItemLookupAPI.GetPersonRemoteSearchResults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPersonRemoteSearchResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jellyfinPersonLookupInfoRemoteSearchQuery** | [**JellyfinPersonLookupInfoRemoteSearchQuery**](JellyfinPersonLookupInfoRemoteSearchQuery.md) | Remote search query. | 

### Return type

[**[]JellyfinJellyfinRemoteSearchResult**](JellyfinRemoteSearchResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSeriesRemoteSearchResults

> []JellyfinJellyfinRemoteSearchResult GetSeriesRemoteSearchResults(ctx).JellyfinSeriesInfoRemoteSearchQuery(jellyfinSeriesInfoRemoteSearchQuery).Execute()

Get series remote search.

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
	jellyfinSeriesInfoRemoteSearchQuery := *openapiclient.NewJellyfinSeriesInfoRemoteSearchQuery() // JellyfinSeriesInfoRemoteSearchQuery | Remote search query.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemLookupAPI.GetSeriesRemoteSearchResults(context.Background()).JellyfinSeriesInfoRemoteSearchQuery(jellyfinSeriesInfoRemoteSearchQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupAPI.GetSeriesRemoteSearchResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSeriesRemoteSearchResults`: []JellyfinJellyfinRemoteSearchResult
	fmt.Fprintf(os.Stdout, "Response from `ItemLookupAPI.GetSeriesRemoteSearchResults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSeriesRemoteSearchResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jellyfinSeriesInfoRemoteSearchQuery** | [**JellyfinSeriesInfoRemoteSearchQuery**](JellyfinSeriesInfoRemoteSearchQuery.md) | Remote search query. | 

### Return type

[**[]JellyfinJellyfinRemoteSearchResult**](JellyfinRemoteSearchResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrailerRemoteSearchResults

> []JellyfinJellyfinRemoteSearchResult GetTrailerRemoteSearchResults(ctx).JellyfinTrailerInfoRemoteSearchQuery(jellyfinTrailerInfoRemoteSearchQuery).Execute()

Get trailer remote search.

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
	jellyfinTrailerInfoRemoteSearchQuery := *openapiclient.NewJellyfinTrailerInfoRemoteSearchQuery() // JellyfinTrailerInfoRemoteSearchQuery | Remote search query.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemLookupAPI.GetTrailerRemoteSearchResults(context.Background()).JellyfinTrailerInfoRemoteSearchQuery(jellyfinTrailerInfoRemoteSearchQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupAPI.GetTrailerRemoteSearchResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrailerRemoteSearchResults`: []JellyfinJellyfinRemoteSearchResult
	fmt.Fprintf(os.Stdout, "Response from `ItemLookupAPI.GetTrailerRemoteSearchResults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTrailerRemoteSearchResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jellyfinTrailerInfoRemoteSearchQuery** | [**JellyfinTrailerInfoRemoteSearchQuery**](JellyfinTrailerInfoRemoteSearchQuery.md) | Remote search query. | 

### Return type

[**[]JellyfinJellyfinRemoteSearchResult**](JellyfinRemoteSearchResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

