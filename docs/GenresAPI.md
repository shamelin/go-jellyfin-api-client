# \GenresAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGenre**](GenresAPI.md#GetGenre) | **Get** /Genres/{genreName} | Gets a genre, by name.
[**GetGenres**](GenresAPI.md#GetGenres) | **Get** /Genres | Gets all genres from a given item, folder, or the entire library.



## GetGenre

> JellyfinJellyfinBaseItemDto GetGenre(ctx, genreName).UserId(userId).Execute()

Gets a genre, by name.

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
	genreName := "genreName_example" // string | The genre name.
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The user id. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GenresAPI.GetGenre(context.Background(), genreName).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GenresAPI.GetGenre``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGenre`: JellyfinJellyfinBaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `GenresAPI.GetGenre`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**genreName** | **string** | The genre name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGenreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | The user id. | 

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


## GetGenres

> JellyfinJellyfinBaseItemDtoQueryResult GetGenres(ctx).StartIndex(startIndex).Limit(limit).SearchTerm(searchTerm).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).IsFavorite(isFavorite).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).UserId(userId).NameStartsWithOrGreater(nameStartsWithOrGreater).NameStartsWith(nameStartsWith).NameLessThan(nameLessThan).SortBy(sortBy).SortOrder(sortOrder).EnableImages(enableImages).EnableTotalRecordCount(enableTotalRecordCount).Execute()

Gets all genres from a given item, folder, or the entire library.

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
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return. (optional)
	searchTerm := "searchTerm_example" // string | The search term. (optional)
	parentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Specify this to localize the search to a specific item or folder. Omit to use the root. (optional)
	fields := []JellyfinJellyfinItemFields{"TODO"} // []JellyfinJellyfinItemFields | Optional. Specify additional fields of information to return in the output. (optional)
	excludeItemTypes := []JellyfinJellyfinBaseItemKind{"TODO"} // []JellyfinJellyfinBaseItemKind | Optional. If specified, results will be filtered out based on item type. This allows multiple, comma delimited. (optional)
	includeItemTypes := []JellyfinJellyfinBaseItemKind{"TODO"} // []JellyfinJellyfinBaseItemKind | Optional. If specified, results will be filtered in based on item type. This allows multiple, comma delimited. (optional)
	isFavorite := true // bool | Optional filter by items that are marked as favorite, or not. (optional)
	imageTypeLimit := int32(56) // int32 | Optional, the max number of images to return, per image type. (optional)
	enableImageTypes := []JellyfinJellyfinImageType{"TODO"} // []JellyfinJellyfinImageType | Optional. The image types to include in the output. (optional)
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User id. (optional)
	nameStartsWithOrGreater := "nameStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	nameStartsWith := "nameStartsWith_example" // string | Optional filter by items whose name is sorted equally than a given input string. (optional)
	nameLessThan := "nameLessThan_example" // string | Optional filter by items whose name is equally or lesser than a given input string. (optional)
	sortBy := []JellyfinJellyfinItemSortBy{"TODO"} // []JellyfinJellyfinItemSortBy | Optional. Specify one or more sort orders, comma delimited. (optional)
	sortOrder := []JellyfinJellyfinSortOrder{"TODO"} // []JellyfinJellyfinSortOrder | Sort Order - Ascending,Descending. (optional)
	enableImages := true // bool | Optional, include image information in output. (optional) (default to true)
	enableTotalRecordCount := true // bool | Optional. Include total record count. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GenresAPI.GetGenres(context.Background()).StartIndex(startIndex).Limit(limit).SearchTerm(searchTerm).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).IsFavorite(isFavorite).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).UserId(userId).NameStartsWithOrGreater(nameStartsWithOrGreater).NameStartsWith(nameStartsWith).NameLessThan(nameLessThan).SortBy(sortBy).SortOrder(sortOrder).EnableImages(enableImages).EnableTotalRecordCount(enableTotalRecordCount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GenresAPI.GetGenres``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGenres`: JellyfinJellyfinBaseItemDtoQueryResult
	fmt.Fprintf(os.Stdout, "Response from `GenresAPI.GetGenres`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGenresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return. | 
 **searchTerm** | **string** | The search term. | 
 **parentId** | **string** | Specify this to localize the search to a specific item or folder. Omit to use the root. | 
 **fields** | [**[]JellyfinJellyfinItemFields**](JellyfinItemFields.md) | Optional. Specify additional fields of information to return in the output. | 
 **excludeItemTypes** | [**[]JellyfinJellyfinBaseItemKind**](JellyfinBaseItemKind.md) | Optional. If specified, results will be filtered out based on item type. This allows multiple, comma delimited. | 
 **includeItemTypes** | [**[]JellyfinJellyfinBaseItemKind**](JellyfinBaseItemKind.md) | Optional. If specified, results will be filtered in based on item type. This allows multiple, comma delimited. | 
 **isFavorite** | **bool** | Optional filter by items that are marked as favorite, or not. | 
 **imageTypeLimit** | **int32** | Optional, the max number of images to return, per image type. | 
 **enableImageTypes** | [**[]JellyfinJellyfinImageType**](JellyfinImageType.md) | Optional. The image types to include in the output. | 
 **userId** | **string** | User id. | 
 **nameStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **nameStartsWith** | **string** | Optional filter by items whose name is sorted equally than a given input string. | 
 **nameLessThan** | **string** | Optional filter by items whose name is equally or lesser than a given input string. | 
 **sortBy** | [**[]JellyfinJellyfinItemSortBy**](JellyfinItemSortBy.md) | Optional. Specify one or more sort orders, comma delimited. | 
 **sortOrder** | [**[]JellyfinJellyfinSortOrder**](JellyfinSortOrder.md) | Sort Order - Ascending,Descending. | 
 **enableImages** | **bool** | Optional, include image information in output. | [default to true]
 **enableTotalRecordCount** | **bool** | Optional. Include total record count. | [default to true]

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

