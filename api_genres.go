/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package jellyfin

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"reflect"
)


// GenresAPIService GenresAPI service
type GenresAPIService service

type GenresAPIGetGenreRequest struct {
	ctx context.Context
	ApiService *GenresAPIService
	genreName string
	userId *string
}

// The user id.
func (r GenresAPIGetGenreRequest) UserId(userId string) GenresAPIGetGenreRequest {
	r.userId = &userId
	return r
}

func (r GenresAPIGetGenreRequest) Execute() (*JellyfinBaseItemDto, *http.Response, error) {
	return r.ApiService.GetGenreExecute(r)
}

/*
GetGenre Gets a genre, by name.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param genreName The genre name.
 @return GenresAPIGetGenreRequest
*/
func (a *GenresAPIService) GetGenre(ctx context.Context, genreName string) GenresAPIGetGenreRequest {
	return GenresAPIGetGenreRequest{
		ApiService: a,
		ctx: ctx,
		genreName: genreName,
	}
}

// Execute executes the request
//  @return JellyfinBaseItemDto
func (a *GenresAPIService) GetGenreExecute(r GenresAPIGetGenreRequest) (*JellyfinBaseItemDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *JellyfinBaseItemDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GenresAPIService.GetGenre")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Genres/{genreName}"
	localVarPath = strings.Replace(localVarPath, "{"+"genreName"+"}", url.PathEscape(parameterValueToString(r.genreName, "genreName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.userId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "userId", r.userId, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/json; profile=CamelCase", "application/json; profile=PascalCase"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["CustomAuthentication"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GenresAPIGetGenresRequest struct {
	ctx context.Context
	ApiService *GenresAPIService
	startIndex *int32
	limit *int32
	searchTerm *string
	parentId *string
	fields *[]JellyfinItemFields
	excludeItemTypes *[]JellyfinBaseItemKind
	includeItemTypes *[]JellyfinBaseItemKind
	isFavorite *bool
	imageTypeLimit *int32
	enableImageTypes *[]JellyfinImageType
	userId *string
	nameStartsWithOrGreater *string
	nameStartsWith *string
	nameLessThan *string
	sortBy *[]JellyfinItemSortBy
	sortOrder *[]JellyfinSortOrder
	enableImages *bool
	enableTotalRecordCount *bool
}

// Optional. The record index to start at. All items with a lower index will be dropped from the results.
func (r GenresAPIGetGenresRequest) StartIndex(startIndex int32) GenresAPIGetGenresRequest {
	r.startIndex = &startIndex
	return r
}

// Optional. The maximum number of records to return.
func (r GenresAPIGetGenresRequest) Limit(limit int32) GenresAPIGetGenresRequest {
	r.limit = &limit
	return r
}

// The search term.
func (r GenresAPIGetGenresRequest) SearchTerm(searchTerm string) GenresAPIGetGenresRequest {
	r.searchTerm = &searchTerm
	return r
}

// Specify this to localize the search to a specific item or folder. Omit to use the root.
func (r GenresAPIGetGenresRequest) ParentId(parentId string) GenresAPIGetGenresRequest {
	r.parentId = &parentId
	return r
}

// Optional. Specify additional fields of information to return in the output.
func (r GenresAPIGetGenresRequest) Fields(fields []JellyfinItemFields) GenresAPIGetGenresRequest {
	r.fields = &fields
	return r
}

// Optional. If specified, results will be filtered out based on item type. This allows multiple, comma delimited.
func (r GenresAPIGetGenresRequest) ExcludeItemTypes(excludeItemTypes []JellyfinBaseItemKind) GenresAPIGetGenresRequest {
	r.excludeItemTypes = &excludeItemTypes
	return r
}

// Optional. If specified, results will be filtered in based on item type. This allows multiple, comma delimited.
func (r GenresAPIGetGenresRequest) IncludeItemTypes(includeItemTypes []JellyfinBaseItemKind) GenresAPIGetGenresRequest {
	r.includeItemTypes = &includeItemTypes
	return r
}

// Optional filter by items that are marked as favorite, or not.
func (r GenresAPIGetGenresRequest) IsFavorite(isFavorite bool) GenresAPIGetGenresRequest {
	r.isFavorite = &isFavorite
	return r
}

// Optional, the max number of images to return, per image type.
func (r GenresAPIGetGenresRequest) ImageTypeLimit(imageTypeLimit int32) GenresAPIGetGenresRequest {
	r.imageTypeLimit = &imageTypeLimit
	return r
}

// Optional. The image types to include in the output.
func (r GenresAPIGetGenresRequest) EnableImageTypes(enableImageTypes []JellyfinImageType) GenresAPIGetGenresRequest {
	r.enableImageTypes = &enableImageTypes
	return r
}

// User id.
func (r GenresAPIGetGenresRequest) UserId(userId string) GenresAPIGetGenresRequest {
	r.userId = &userId
	return r
}

// Optional filter by items whose name is sorted equally or greater than a given input string.
func (r GenresAPIGetGenresRequest) NameStartsWithOrGreater(nameStartsWithOrGreater string) GenresAPIGetGenresRequest {
	r.nameStartsWithOrGreater = &nameStartsWithOrGreater
	return r
}

// Optional filter by items whose name is sorted equally than a given input string.
func (r GenresAPIGetGenresRequest) NameStartsWith(nameStartsWith string) GenresAPIGetGenresRequest {
	r.nameStartsWith = &nameStartsWith
	return r
}

// Optional filter by items whose name is equally or lesser than a given input string.
func (r GenresAPIGetGenresRequest) NameLessThan(nameLessThan string) GenresAPIGetGenresRequest {
	r.nameLessThan = &nameLessThan
	return r
}

// Optional. Specify one or more sort orders, comma delimited.
func (r GenresAPIGetGenresRequest) SortBy(sortBy []JellyfinItemSortBy) GenresAPIGetGenresRequest {
	r.sortBy = &sortBy
	return r
}

// Sort Order - Ascending,Descending.
func (r GenresAPIGetGenresRequest) SortOrder(sortOrder []JellyfinSortOrder) GenresAPIGetGenresRequest {
	r.sortOrder = &sortOrder
	return r
}

// Optional, include image information in output.
func (r GenresAPIGetGenresRequest) EnableImages(enableImages bool) GenresAPIGetGenresRequest {
	r.enableImages = &enableImages
	return r
}

// Optional. Include total record count.
func (r GenresAPIGetGenresRequest) EnableTotalRecordCount(enableTotalRecordCount bool) GenresAPIGetGenresRequest {
	r.enableTotalRecordCount = &enableTotalRecordCount
	return r
}

func (r GenresAPIGetGenresRequest) Execute() (*JellyfinBaseItemDtoQueryResult, *http.Response, error) {
	return r.ApiService.GetGenresExecute(r)
}

/*
GetGenres Gets all genres from a given item, folder, or the entire library.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return GenresAPIGetGenresRequest
*/
func (a *GenresAPIService) GetGenres(ctx context.Context) GenresAPIGetGenresRequest {
	return GenresAPIGetGenresRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return JellyfinBaseItemDtoQueryResult
func (a *GenresAPIService) GetGenresExecute(r GenresAPIGetGenresRequest) (*JellyfinBaseItemDtoQueryResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *JellyfinBaseItemDtoQueryResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GenresAPIService.GetGenres")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Genres"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startIndex", r.startIndex, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.searchTerm != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchTerm", r.searchTerm, "form", "")
	}
	if r.parentId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "parentId", r.parentId, "form", "")
	}
	if r.fields != nil {
		t := *r.fields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "fields", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "fields", t, "form", "multi")
		}
	}
	if r.excludeItemTypes != nil {
		t := *r.excludeItemTypes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "excludeItemTypes", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "excludeItemTypes", t, "form", "multi")
		}
	}
	if r.includeItemTypes != nil {
		t := *r.includeItemTypes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "includeItemTypes", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "includeItemTypes", t, "form", "multi")
		}
	}
	if r.isFavorite != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "isFavorite", r.isFavorite, "form", "")
	}
	if r.imageTypeLimit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "imageTypeLimit", r.imageTypeLimit, "form", "")
	}
	if r.enableImageTypes != nil {
		t := *r.enableImageTypes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "enableImageTypes", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "enableImageTypes", t, "form", "multi")
		}
	}
	if r.userId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "userId", r.userId, "form", "")
	}
	if r.nameStartsWithOrGreater != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nameStartsWithOrGreater", r.nameStartsWithOrGreater, "form", "")
	}
	if r.nameStartsWith != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nameStartsWith", r.nameStartsWith, "form", "")
	}
	if r.nameLessThan != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nameLessThan", r.nameLessThan, "form", "")
	}
	if r.sortBy != nil {
		t := *r.sortBy
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "sortBy", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "sortBy", t, "form", "multi")
		}
	}
	if r.sortOrder != nil {
		t := *r.sortOrder
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "sortOrder", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "sortOrder", t, "form", "multi")
		}
	}
	if r.enableImages != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "enableImages", r.enableImages, "form", "")
	} else {
		var defaultValue bool = true
		r.enableImages = &defaultValue
	}
	if r.enableTotalRecordCount != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "enableTotalRecordCount", r.enableTotalRecordCount, "form", "")
	} else {
		var defaultValue bool = true
		r.enableTotalRecordCount = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/json; profile=CamelCase", "application/json; profile=PascalCase"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["CustomAuthentication"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
