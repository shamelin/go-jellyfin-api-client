# \SyncPlayAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SyncPlayBuffering**](SyncPlayAPI.md#SyncPlayBuffering) | **Post** /SyncPlay/Buffering | Notify SyncPlay group that member is buffering.
[**SyncPlayCreateGroup**](SyncPlayAPI.md#SyncPlayCreateGroup) | **Post** /SyncPlay/New | Create a new SyncPlay group.
[**SyncPlayGetGroups**](SyncPlayAPI.md#SyncPlayGetGroups) | **Get** /SyncPlay/List | Gets all SyncPlay groups.
[**SyncPlayJoinGroup**](SyncPlayAPI.md#SyncPlayJoinGroup) | **Post** /SyncPlay/Join | Join an existing SyncPlay group.
[**SyncPlayLeaveGroup**](SyncPlayAPI.md#SyncPlayLeaveGroup) | **Post** /SyncPlay/Leave | Leave the joined SyncPlay group.
[**SyncPlayMovePlaylistItem**](SyncPlayAPI.md#SyncPlayMovePlaylistItem) | **Post** /SyncPlay/MovePlaylistItem | Request to move an item in the playlist in SyncPlay group.
[**SyncPlayNextItem**](SyncPlayAPI.md#SyncPlayNextItem) | **Post** /SyncPlay/NextItem | Request next item in SyncPlay group.
[**SyncPlayPause**](SyncPlayAPI.md#SyncPlayPause) | **Post** /SyncPlay/Pause | Request pause in SyncPlay group.
[**SyncPlayPing**](SyncPlayAPI.md#SyncPlayPing) | **Post** /SyncPlay/Ping | Update session ping.
[**SyncPlayPreviousItem**](SyncPlayAPI.md#SyncPlayPreviousItem) | **Post** /SyncPlay/PreviousItem | Request previous item in SyncPlay group.
[**SyncPlayQueue**](SyncPlayAPI.md#SyncPlayQueue) | **Post** /SyncPlay/Queue | Request to queue items to the playlist of a SyncPlay group.
[**SyncPlayReady**](SyncPlayAPI.md#SyncPlayReady) | **Post** /SyncPlay/Ready | Notify SyncPlay group that member is ready for playback.
[**SyncPlayRemoveFromPlaylist**](SyncPlayAPI.md#SyncPlayRemoveFromPlaylist) | **Post** /SyncPlay/RemoveFromPlaylist | Request to remove items from the playlist in SyncPlay group.
[**SyncPlaySeek**](SyncPlayAPI.md#SyncPlaySeek) | **Post** /SyncPlay/Seek | Request seek in SyncPlay group.
[**SyncPlaySetIgnoreWait**](SyncPlayAPI.md#SyncPlaySetIgnoreWait) | **Post** /SyncPlay/SetIgnoreWait | Request SyncPlay group to ignore member during group-wait.
[**SyncPlaySetNewQueue**](SyncPlayAPI.md#SyncPlaySetNewQueue) | **Post** /SyncPlay/SetNewQueue | Request to set new playlist in SyncPlay group.
[**SyncPlaySetPlaylistItem**](SyncPlayAPI.md#SyncPlaySetPlaylistItem) | **Post** /SyncPlay/SetPlaylistItem | Request to change playlist item in SyncPlay group.
[**SyncPlaySetRepeatMode**](SyncPlayAPI.md#SyncPlaySetRepeatMode) | **Post** /SyncPlay/SetRepeatMode | Request to set repeat mode in SyncPlay group.
[**SyncPlaySetShuffleMode**](SyncPlayAPI.md#SyncPlaySetShuffleMode) | **Post** /SyncPlay/SetShuffleMode | Request to set shuffle mode in SyncPlay group.
[**SyncPlayStop**](SyncPlayAPI.md#SyncPlayStop) | **Post** /SyncPlay/Stop | Request stop in SyncPlay group.
[**SyncPlayUnpause**](SyncPlayAPI.md#SyncPlayUnpause) | **Post** /SyncPlay/Unpause | Request unpause in SyncPlay group.



## SyncPlayBuffering

> SyncPlayBuffering(ctx).JellyfinBufferRequestDto(jellyfinBufferRequestDto).Execute()

Notify SyncPlay group that member is buffering.

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
	jellyfinBufferRequestDto := *openapiclient.NewJellyfinBufferRequestDto() // JellyfinBufferRequestDto | The player status.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncPlayAPI.SyncPlayBuffering(context.Background()).JellyfinBufferRequestDto(jellyfinBufferRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncPlayAPI.SyncPlayBuffering``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncPlayBufferingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jellyfinBufferRequestDto** | [**JellyfinBufferRequestDto**](JellyfinBufferRequestDto.md) | The player status. | 

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


## SyncPlayCreateGroup

> SyncPlayCreateGroup(ctx).JellyfinNewGroupRequestDto(jellyfinNewGroupRequestDto).Execute()

Create a new SyncPlay group.

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
	jellyfinNewGroupRequestDto := *openapiclient.NewJellyfinNewGroupRequestDto() // JellyfinNewGroupRequestDto | The settings of the new group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncPlayAPI.SyncPlayCreateGroup(context.Background()).JellyfinNewGroupRequestDto(jellyfinNewGroupRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncPlayAPI.SyncPlayCreateGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncPlayCreateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jellyfinNewGroupRequestDto** | [**JellyfinNewGroupRequestDto**](JellyfinNewGroupRequestDto.md) | The settings of the new group. | 

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


## SyncPlayGetGroups

> []JellyfinJellyfinGroupInfoDto SyncPlayGetGroups(ctx).Execute()

Gets all SyncPlay groups.

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
	resp, r, err := apiClient.SyncPlayAPI.SyncPlayGetGroups(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncPlayAPI.SyncPlayGetGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncPlayGetGroups`: []JellyfinJellyfinGroupInfoDto
	fmt.Fprintf(os.Stdout, "Response from `SyncPlayAPI.SyncPlayGetGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSyncPlayGetGroupsRequest struct via the builder pattern


### Return type

[**[]JellyfinJellyfinGroupInfoDto**](JellyfinGroupInfoDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncPlayJoinGroup

> SyncPlayJoinGroup(ctx).JellyfinJoinGroupRequestDto(jellyfinJoinGroupRequestDto).Execute()

Join an existing SyncPlay group.

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
	jellyfinJoinGroupRequestDto := *openapiclient.NewJellyfinJoinGroupRequestDto() // JellyfinJoinGroupRequestDto | The group to join.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncPlayAPI.SyncPlayJoinGroup(context.Background()).JellyfinJoinGroupRequestDto(jellyfinJoinGroupRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncPlayAPI.SyncPlayJoinGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncPlayJoinGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jellyfinJoinGroupRequestDto** | [**JellyfinJoinGroupRequestDto**](JellyfinJoinGroupRequestDto.md) | The group to join. | 

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


## SyncPlayLeaveGroup

> SyncPlayLeaveGroup(ctx).Execute()

Leave the joined SyncPlay group.

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
	r, err := apiClient.SyncPlayAPI.SyncPlayLeaveGroup(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncPlayAPI.SyncPlayLeaveGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSyncPlayLeaveGroupRequest struct via the builder pattern


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


## SyncPlayMovePlaylistItem

> SyncPlayMovePlaylistItem(ctx).JellyfinMovePlaylistItemRequestDto(jellyfinMovePlaylistItemRequestDto).Execute()

Request to move an item in the playlist in SyncPlay group.

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
	jellyfinMovePlaylistItemRequestDto := *openapiclient.NewJellyfinMovePlaylistItemRequestDto() // JellyfinMovePlaylistItemRequestDto | The new position for the item.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncPlayAPI.SyncPlayMovePlaylistItem(context.Background()).JellyfinMovePlaylistItemRequestDto(jellyfinMovePlaylistItemRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncPlayAPI.SyncPlayMovePlaylistItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncPlayMovePlaylistItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jellyfinMovePlaylistItemRequestDto** | [**JellyfinMovePlaylistItemRequestDto**](JellyfinMovePlaylistItemRequestDto.md) | The new position for the item. | 

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


## SyncPlayNextItem

> SyncPlayNextItem(ctx).JellyfinNextItemRequestDto(jellyfinNextItemRequestDto).Execute()

Request next item in SyncPlay group.

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
	jellyfinNextItemRequestDto := *openapiclient.NewJellyfinNextItemRequestDto() // JellyfinNextItemRequestDto | The current item information.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncPlayAPI.SyncPlayNextItem(context.Background()).JellyfinNextItemRequestDto(jellyfinNextItemRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncPlayAPI.SyncPlayNextItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncPlayNextItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jellyfinNextItemRequestDto** | [**JellyfinNextItemRequestDto**](JellyfinNextItemRequestDto.md) | The current item information. | 

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


## SyncPlayPause

> SyncPlayPause(ctx).Execute()

Request pause in SyncPlay group.

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
	r, err := apiClient.SyncPlayAPI.SyncPlayPause(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncPlayAPI.SyncPlayPause``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSyncPlayPauseRequest struct via the builder pattern


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


## SyncPlayPing

> SyncPlayPing(ctx).JellyfinPingRequestDto(jellyfinPingRequestDto).Execute()

Update session ping.

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
	jellyfinPingRequestDto := *openapiclient.NewJellyfinPingRequestDto() // JellyfinPingRequestDto | The new ping.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncPlayAPI.SyncPlayPing(context.Background()).JellyfinPingRequestDto(jellyfinPingRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncPlayAPI.SyncPlayPing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncPlayPingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jellyfinPingRequestDto** | [**JellyfinPingRequestDto**](JellyfinPingRequestDto.md) | The new ping. | 

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


## SyncPlayPreviousItem

> SyncPlayPreviousItem(ctx).JellyfinPreviousItemRequestDto(jellyfinPreviousItemRequestDto).Execute()

Request previous item in SyncPlay group.

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
	jellyfinPreviousItemRequestDto := *openapiclient.NewJellyfinPreviousItemRequestDto() // JellyfinPreviousItemRequestDto | The current item information.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncPlayAPI.SyncPlayPreviousItem(context.Background()).JellyfinPreviousItemRequestDto(jellyfinPreviousItemRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncPlayAPI.SyncPlayPreviousItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncPlayPreviousItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jellyfinPreviousItemRequestDto** | [**JellyfinPreviousItemRequestDto**](JellyfinPreviousItemRequestDto.md) | The current item information. | 

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


## SyncPlayQueue

> SyncPlayQueue(ctx).JellyfinQueueRequestDto(jellyfinQueueRequestDto).Execute()

Request to queue items to the playlist of a SyncPlay group.

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
	jellyfinQueueRequestDto := *openapiclient.NewJellyfinQueueRequestDto() // JellyfinQueueRequestDto | The items to add.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncPlayAPI.SyncPlayQueue(context.Background()).JellyfinQueueRequestDto(jellyfinQueueRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncPlayAPI.SyncPlayQueue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncPlayQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jellyfinQueueRequestDto** | [**JellyfinQueueRequestDto**](JellyfinQueueRequestDto.md) | The items to add. | 

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


## SyncPlayReady

> SyncPlayReady(ctx).JellyfinReadyRequestDto(jellyfinReadyRequestDto).Execute()

Notify SyncPlay group that member is ready for playback.

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
	jellyfinReadyRequestDto := *openapiclient.NewJellyfinReadyRequestDto() // JellyfinReadyRequestDto | The player status.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncPlayAPI.SyncPlayReady(context.Background()).JellyfinReadyRequestDto(jellyfinReadyRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncPlayAPI.SyncPlayReady``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncPlayReadyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jellyfinReadyRequestDto** | [**JellyfinReadyRequestDto**](JellyfinReadyRequestDto.md) | The player status. | 

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


## SyncPlayRemoveFromPlaylist

> SyncPlayRemoveFromPlaylist(ctx).JellyfinRemoveFromPlaylistRequestDto(jellyfinRemoveFromPlaylistRequestDto).Execute()

Request to remove items from the playlist in SyncPlay group.

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
	jellyfinRemoveFromPlaylistRequestDto := *openapiclient.NewJellyfinRemoveFromPlaylistRequestDto() // JellyfinRemoveFromPlaylistRequestDto | The items to remove.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncPlayAPI.SyncPlayRemoveFromPlaylist(context.Background()).JellyfinRemoveFromPlaylistRequestDto(jellyfinRemoveFromPlaylistRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncPlayAPI.SyncPlayRemoveFromPlaylist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncPlayRemoveFromPlaylistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jellyfinRemoveFromPlaylistRequestDto** | [**JellyfinRemoveFromPlaylistRequestDto**](JellyfinRemoveFromPlaylistRequestDto.md) | The items to remove. | 

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


## SyncPlaySeek

> SyncPlaySeek(ctx).JellyfinSeekRequestDto(jellyfinSeekRequestDto).Execute()

Request seek in SyncPlay group.

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
	jellyfinSeekRequestDto := *openapiclient.NewJellyfinSeekRequestDto() // JellyfinSeekRequestDto | The new playback position.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncPlayAPI.SyncPlaySeek(context.Background()).JellyfinSeekRequestDto(jellyfinSeekRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncPlayAPI.SyncPlaySeek``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncPlaySeekRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jellyfinSeekRequestDto** | [**JellyfinSeekRequestDto**](JellyfinSeekRequestDto.md) | The new playback position. | 

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


## SyncPlaySetIgnoreWait

> SyncPlaySetIgnoreWait(ctx).JellyfinIgnoreWaitRequestDto(jellyfinIgnoreWaitRequestDto).Execute()

Request SyncPlay group to ignore member during group-wait.

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
	jellyfinIgnoreWaitRequestDto := *openapiclient.NewJellyfinIgnoreWaitRequestDto() // JellyfinIgnoreWaitRequestDto | The settings to set.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncPlayAPI.SyncPlaySetIgnoreWait(context.Background()).JellyfinIgnoreWaitRequestDto(jellyfinIgnoreWaitRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncPlayAPI.SyncPlaySetIgnoreWait``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncPlaySetIgnoreWaitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jellyfinIgnoreWaitRequestDto** | [**JellyfinIgnoreWaitRequestDto**](JellyfinIgnoreWaitRequestDto.md) | The settings to set. | 

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


## SyncPlaySetNewQueue

> SyncPlaySetNewQueue(ctx).JellyfinPlayRequestDto(jellyfinPlayRequestDto).Execute()

Request to set new playlist in SyncPlay group.

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
	jellyfinPlayRequestDto := *openapiclient.NewJellyfinPlayRequestDto() // JellyfinPlayRequestDto | The new playlist to play in the group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncPlayAPI.SyncPlaySetNewQueue(context.Background()).JellyfinPlayRequestDto(jellyfinPlayRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncPlayAPI.SyncPlaySetNewQueue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncPlaySetNewQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jellyfinPlayRequestDto** | [**JellyfinPlayRequestDto**](JellyfinPlayRequestDto.md) | The new playlist to play in the group. | 

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


## SyncPlaySetPlaylistItem

> SyncPlaySetPlaylistItem(ctx).JellyfinSetPlaylistItemRequestDto(jellyfinSetPlaylistItemRequestDto).Execute()

Request to change playlist item in SyncPlay group.

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
	jellyfinSetPlaylistItemRequestDto := *openapiclient.NewJellyfinSetPlaylistItemRequestDto() // JellyfinSetPlaylistItemRequestDto | The new item to play.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncPlayAPI.SyncPlaySetPlaylistItem(context.Background()).JellyfinSetPlaylistItemRequestDto(jellyfinSetPlaylistItemRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncPlayAPI.SyncPlaySetPlaylistItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncPlaySetPlaylistItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jellyfinSetPlaylistItemRequestDto** | [**JellyfinSetPlaylistItemRequestDto**](JellyfinSetPlaylistItemRequestDto.md) | The new item to play. | 

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


## SyncPlaySetRepeatMode

> SyncPlaySetRepeatMode(ctx).JellyfinSetRepeatModeRequestDto(jellyfinSetRepeatModeRequestDto).Execute()

Request to set repeat mode in SyncPlay group.

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
	jellyfinSetRepeatModeRequestDto := *openapiclient.NewJellyfinSetRepeatModeRequestDto() // JellyfinSetRepeatModeRequestDto | The new repeat mode.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncPlayAPI.SyncPlaySetRepeatMode(context.Background()).JellyfinSetRepeatModeRequestDto(jellyfinSetRepeatModeRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncPlayAPI.SyncPlaySetRepeatMode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncPlaySetRepeatModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jellyfinSetRepeatModeRequestDto** | [**JellyfinSetRepeatModeRequestDto**](JellyfinSetRepeatModeRequestDto.md) | The new repeat mode. | 

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


## SyncPlaySetShuffleMode

> SyncPlaySetShuffleMode(ctx).JellyfinSetShuffleModeRequestDto(jellyfinSetShuffleModeRequestDto).Execute()

Request to set shuffle mode in SyncPlay group.

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
	jellyfinSetShuffleModeRequestDto := *openapiclient.NewJellyfinSetShuffleModeRequestDto() // JellyfinSetShuffleModeRequestDto | The new shuffle mode.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncPlayAPI.SyncPlaySetShuffleMode(context.Background()).JellyfinSetShuffleModeRequestDto(jellyfinSetShuffleModeRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncPlayAPI.SyncPlaySetShuffleMode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncPlaySetShuffleModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jellyfinSetShuffleModeRequestDto** | [**JellyfinSetShuffleModeRequestDto**](JellyfinSetShuffleModeRequestDto.md) | The new shuffle mode. | 

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


## SyncPlayStop

> SyncPlayStop(ctx).Execute()

Request stop in SyncPlay group.

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
	r, err := apiClient.SyncPlayAPI.SyncPlayStop(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncPlayAPI.SyncPlayStop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSyncPlayStopRequest struct via the builder pattern


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


## SyncPlayUnpause

> SyncPlayUnpause(ctx).Execute()

Request unpause in SyncPlay group.

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
	r, err := apiClient.SyncPlayAPI.SyncPlayUnpause(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncPlayAPI.SyncPlayUnpause``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSyncPlayUnpauseRequest struct via the builder pattern


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

