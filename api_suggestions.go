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


// SuggestionsAPIService SuggestionsAPI service
type SuggestionsAPIService service

type ApiGetSuggestionsRequest struct {
	ctx context.Context
	ApiService *SuggestionsAPIService
	userId *string
	mediaType *[]MediaType
	type_ *[]BaseItemKind
	startIndex *int32
	limit *int32
	enableTotalRecordCount *bool
}

// The user id.
func (r ApiGetSuggestionsRequest) UserId(userId string) ApiGetSuggestionsRequest {
	r.userId = &userId
	return r
}

// The media types.
func (r ApiGetSuggestionsRequest) MediaType(mediaType []MediaType) ApiGetSuggestionsRequest {
	r.mediaType = &mediaType
	return r
}

// The type.
func (r ApiGetSuggestionsRequest) Type_(type_ []BaseItemKind) ApiGetSuggestionsRequest {
	r.type_ = &type_
	return r
}

// Optional. The start index.
func (r ApiGetSuggestionsRequest) StartIndex(startIndex int32) ApiGetSuggestionsRequest {
	r.startIndex = &startIndex
	return r
}

// Optional. The limit.
func (r ApiGetSuggestionsRequest) Limit(limit int32) ApiGetSuggestionsRequest {
	r.limit = &limit
	return r
}

// Whether to enable the total record count.
func (r ApiGetSuggestionsRequest) EnableTotalRecordCount(enableTotalRecordCount bool) ApiGetSuggestionsRequest {
	r.enableTotalRecordCount = &enableTotalRecordCount
	return r
}

func (r ApiGetSuggestionsRequest) Execute() (*BaseItemDtoQueryResult, *http.Response, error) {
	return r.ApiService.GetSuggestionsExecute(r)
}

/*
GetSuggestions Gets suggestions.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSuggestionsRequest
*/
func (a *SuggestionsAPIService) GetSuggestions(ctx context.Context) ApiGetSuggestionsRequest {
	return ApiGetSuggestionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BaseItemDtoQueryResult
func (a *SuggestionsAPIService) GetSuggestionsExecute(r ApiGetSuggestionsRequest) (*BaseItemDtoQueryResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BaseItemDtoQueryResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SuggestionsAPIService.GetSuggestions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Items/Suggestions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.userId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "userId", r.userId, "form", "")
	}
	if r.mediaType != nil {
		t := *r.mediaType
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "mediaType", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "mediaType", t, "form", "multi")
		}
	}
	if r.type_ != nil {
		t := *r.type_
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "type", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "type", t, "form", "multi")
		}
	}
	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startIndex", r.startIndex, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.enableTotalRecordCount != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "enableTotalRecordCount", r.enableTotalRecordCount, "form", "")
	} else {
		var defaultValue bool = false
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
