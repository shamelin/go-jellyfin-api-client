/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package jellyfin-api-client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"reflect"
)


// SearchAPIService SearchAPI service
type SearchAPIService service

type ApiGetSearchHintsRequest struct {
	ctx context.Context
	ApiService *SearchAPIService
	searchTerm *string
	startIndex *int32
	limit *int32
	userId *string
	includeItemTypes *[]BaseItemKind
	excludeItemTypes *[]BaseItemKind
	mediaTypes *[]MediaType
	parentId *string
	isMovie *bool
	isSeries *bool
	isNews *bool
	isKids *bool
	isSports *bool
	includePeople *bool
	includeMedia *bool
	includeGenres *bool
	includeStudios *bool
	includeArtists *bool
}

// The search term to filter on.
func (r ApiGetSearchHintsRequest) SearchTerm(searchTerm string) ApiGetSearchHintsRequest {
	r.searchTerm = &searchTerm
	return r
}

// Optional. The record index to start at. All items with a lower index will be dropped from the results.
func (r ApiGetSearchHintsRequest) StartIndex(startIndex int32) ApiGetSearchHintsRequest {
	r.startIndex = &startIndex
	return r
}

// Optional. The maximum number of records to return.
func (r ApiGetSearchHintsRequest) Limit(limit int32) ApiGetSearchHintsRequest {
	r.limit = &limit
	return r
}

// Optional. Supply a user id to search within a user&#39;s library or omit to search all.
func (r ApiGetSearchHintsRequest) UserId(userId string) ApiGetSearchHintsRequest {
	r.userId = &userId
	return r
}

// If specified, only results with the specified item types are returned. This allows multiple, comma delimited.
func (r ApiGetSearchHintsRequest) IncludeItemTypes(includeItemTypes []BaseItemKind) ApiGetSearchHintsRequest {
	r.includeItemTypes = &includeItemTypes
	return r
}

// If specified, results with these item types are filtered out. This allows multiple, comma delimited.
func (r ApiGetSearchHintsRequest) ExcludeItemTypes(excludeItemTypes []BaseItemKind) ApiGetSearchHintsRequest {
	r.excludeItemTypes = &excludeItemTypes
	return r
}

// If specified, only results with the specified media types are returned. This allows multiple, comma delimited.
func (r ApiGetSearchHintsRequest) MediaTypes(mediaTypes []MediaType) ApiGetSearchHintsRequest {
	r.mediaTypes = &mediaTypes
	return r
}

// If specified, only children of the parent are returned.
func (r ApiGetSearchHintsRequest) ParentId(parentId string) ApiGetSearchHintsRequest {
	r.parentId = &parentId
	return r
}

// Optional filter for movies.
func (r ApiGetSearchHintsRequest) IsMovie(isMovie bool) ApiGetSearchHintsRequest {
	r.isMovie = &isMovie
	return r
}

// Optional filter for series.
func (r ApiGetSearchHintsRequest) IsSeries(isSeries bool) ApiGetSearchHintsRequest {
	r.isSeries = &isSeries
	return r
}

// Optional filter for news.
func (r ApiGetSearchHintsRequest) IsNews(isNews bool) ApiGetSearchHintsRequest {
	r.isNews = &isNews
	return r
}

// Optional filter for kids.
func (r ApiGetSearchHintsRequest) IsKids(isKids bool) ApiGetSearchHintsRequest {
	r.isKids = &isKids
	return r
}

// Optional filter for sports.
func (r ApiGetSearchHintsRequest) IsSports(isSports bool) ApiGetSearchHintsRequest {
	r.isSports = &isSports
	return r
}

// Optional filter whether to include people.
func (r ApiGetSearchHintsRequest) IncludePeople(includePeople bool) ApiGetSearchHintsRequest {
	r.includePeople = &includePeople
	return r
}

// Optional filter whether to include media.
func (r ApiGetSearchHintsRequest) IncludeMedia(includeMedia bool) ApiGetSearchHintsRequest {
	r.includeMedia = &includeMedia
	return r
}

// Optional filter whether to include genres.
func (r ApiGetSearchHintsRequest) IncludeGenres(includeGenres bool) ApiGetSearchHintsRequest {
	r.includeGenres = &includeGenres
	return r
}

// Optional filter whether to include studios.
func (r ApiGetSearchHintsRequest) IncludeStudios(includeStudios bool) ApiGetSearchHintsRequest {
	r.includeStudios = &includeStudios
	return r
}

// Optional filter whether to include artists.
func (r ApiGetSearchHintsRequest) IncludeArtists(includeArtists bool) ApiGetSearchHintsRequest {
	r.includeArtists = &includeArtists
	return r
}

func (r ApiGetSearchHintsRequest) Execute() (*SearchHintResult, *http.Response, error) {
	return r.ApiService.GetSearchHintsExecute(r)
}

/*
GetSearchHints Gets the search hint result.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSearchHintsRequest
*/
func (a *SearchAPIService) GetSearchHints(ctx context.Context) ApiGetSearchHintsRequest {
	return ApiGetSearchHintsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SearchHintResult
func (a *SearchAPIService) GetSearchHintsExecute(r ApiGetSearchHintsRequest) (*SearchHintResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SearchHintResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SearchAPIService.GetSearchHints")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Search/Hints"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.searchTerm == nil {
		return localVarReturnValue, nil, reportError("searchTerm is required and must be specified")
	}

	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startIndex", r.startIndex, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.userId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "userId", r.userId, "form", "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "searchTerm", r.searchTerm, "form", "")
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
	if r.mediaTypes != nil {
		t := *r.mediaTypes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "mediaTypes", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "mediaTypes", t, "form", "multi")
		}
	}
	if r.parentId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "parentId", r.parentId, "form", "")
	}
	if r.isMovie != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "isMovie", r.isMovie, "form", "")
	}
	if r.isSeries != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "isSeries", r.isSeries, "form", "")
	}
	if r.isNews != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "isNews", r.isNews, "form", "")
	}
	if r.isKids != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "isKids", r.isKids, "form", "")
	}
	if r.isSports != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "isSports", r.isSports, "form", "")
	}
	if r.includePeople != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includePeople", r.includePeople, "form", "")
	} else {
		var defaultValue bool = true
		r.includePeople = &defaultValue
	}
	if r.includeMedia != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeMedia", r.includeMedia, "form", "")
	} else {
		var defaultValue bool = true
		r.includeMedia = &defaultValue
	}
	if r.includeGenres != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeGenres", r.includeGenres, "form", "")
	} else {
		var defaultValue bool = true
		r.includeGenres = &defaultValue
	}
	if r.includeStudios != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeStudios", r.includeStudios, "form", "")
	} else {
		var defaultValue bool = true
		r.includeStudios = &defaultValue
	}
	if r.includeArtists != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeArtists", r.includeArtists, "form", "")
	} else {
		var defaultValue bool = true
		r.includeArtists = &defaultValue
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
