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
	"os"
)


// TrickplayAPIService TrickplayAPI service
type TrickplayAPIService service

type TrickplayAPIGetTrickplayHlsPlaylistRequest struct {
	ctx context.Context
	ApiService *TrickplayAPIService
	itemId string
	width int32
	mediaSourceId *string
}

// The media version id, if using an alternate version.
func (r TrickplayAPIGetTrickplayHlsPlaylistRequest) MediaSourceId(mediaSourceId string) TrickplayAPIGetTrickplayHlsPlaylistRequest {
	r.mediaSourceId = &mediaSourceId
	return r
}

func (r TrickplayAPIGetTrickplayHlsPlaylistRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.GetTrickplayHlsPlaylistExecute(r)
}

/*
GetTrickplayHlsPlaylist Gets an image tiles playlist for trickplay.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param itemId The item id.
 @param width The width of a single tile.
 @return TrickplayAPIGetTrickplayHlsPlaylistRequest
*/
func (a *TrickplayAPIService) GetTrickplayHlsPlaylist(ctx context.Context, itemId string, width int32) TrickplayAPIGetTrickplayHlsPlaylistRequest {
	return TrickplayAPIGetTrickplayHlsPlaylistRequest{
		ApiService: a,
		ctx: ctx,
		itemId: itemId,
		width: width,
	}
}

// Execute executes the request
//  @return *os.File
func (a *TrickplayAPIService) GetTrickplayHlsPlaylistExecute(r TrickplayAPIGetTrickplayHlsPlaylistRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TrickplayAPIService.GetTrickplayHlsPlaylist")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Videos/{itemId}/Trickplay/{width}/tiles.m3u8"
	localVarPath = strings.Replace(localVarPath, "{"+"itemId"+"}", url.PathEscape(parameterValueToString(r.itemId, "itemId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"width"+"}", url.PathEscape(parameterValueToString(r.width, "width")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.mediaSourceId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mediaSourceId", r.mediaSourceId, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/x-mpegURL", "application/json", "application/json; profile=CamelCase", "application/json; profile=PascalCase"}

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
		if localVarHTTPResponse.StatusCode == 404 {
			var v JellyfinProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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

type TrickplayAPIGetTrickplayTileImageRequest struct {
	ctx context.Context
	ApiService *TrickplayAPIService
	itemId string
	width int32
	index int32
	mediaSourceId *string
}

// The media version id, if using an alternate version.
func (r TrickplayAPIGetTrickplayTileImageRequest) MediaSourceId(mediaSourceId string) TrickplayAPIGetTrickplayTileImageRequest {
	r.mediaSourceId = &mediaSourceId
	return r
}

func (r TrickplayAPIGetTrickplayTileImageRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.GetTrickplayTileImageExecute(r)
}

/*
GetTrickplayTileImage Gets a trickplay tile image.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param itemId The item id.
 @param width The width of a single tile.
 @param index The index of the desired tile.
 @return TrickplayAPIGetTrickplayTileImageRequest
*/
func (a *TrickplayAPIService) GetTrickplayTileImage(ctx context.Context, itemId string, width int32, index int32) TrickplayAPIGetTrickplayTileImageRequest {
	return TrickplayAPIGetTrickplayTileImageRequest{
		ApiService: a,
		ctx: ctx,
		itemId: itemId,
		width: width,
		index: index,
	}
}

// Execute executes the request
//  @return *os.File
func (a *TrickplayAPIService) GetTrickplayTileImageExecute(r TrickplayAPIGetTrickplayTileImageRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TrickplayAPIService.GetTrickplayTileImage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Videos/{itemId}/Trickplay/{width}/{index}.jpg"
	localVarPath = strings.Replace(localVarPath, "{"+"itemId"+"}", url.PathEscape(parameterValueToString(r.itemId, "itemId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"width"+"}", url.PathEscape(parameterValueToString(r.width, "width")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", url.PathEscape(parameterValueToString(r.index, "index")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.mediaSourceId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mediaSourceId", r.mediaSourceId, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"image/*", "application/json", "application/json; profile=CamelCase", "application/json; profile=PascalCase"}

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
		if localVarHTTPResponse.StatusCode == 404 {
			var v JellyfinProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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
