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
)


// RemoteImageAPIService RemoteImageAPI service
type RemoteImageAPIService service

type RemoteImageAPIDownloadRemoteImageRequest struct {
	ctx context.Context
	ApiService *RemoteImageAPIService
	itemId string
	type_ *JellyfinImageType
	imageUrl *string
}

// The image type.
func (r RemoteImageAPIDownloadRemoteImageRequest) Type_(type_ JellyfinImageType) RemoteImageAPIDownloadRemoteImageRequest {
	r.type_ = &type_
	return r
}

// The image url.
func (r RemoteImageAPIDownloadRemoteImageRequest) ImageUrl(imageUrl string) RemoteImageAPIDownloadRemoteImageRequest {
	r.imageUrl = &imageUrl
	return r
}

func (r RemoteImageAPIDownloadRemoteImageRequest) Execute() (*http.Response, error) {
	return r.ApiService.DownloadRemoteImageExecute(r)
}

/*
DownloadRemoteImage Downloads a remote image for an item.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param itemId Item Id.
 @return RemoteImageAPIDownloadRemoteImageRequest
*/
func (a *RemoteImageAPIService) DownloadRemoteImage(ctx context.Context, itemId string) RemoteImageAPIDownloadRemoteImageRequest {
	return RemoteImageAPIDownloadRemoteImageRequest{
		ApiService: a,
		ctx: ctx,
		itemId: itemId,
	}
}

// Execute executes the request
func (a *RemoteImageAPIService) DownloadRemoteImageExecute(r RemoteImageAPIDownloadRemoteImageRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemoteImageAPIService.DownloadRemoteImage")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Items/{itemId}/RemoteImages/Download"
	localVarPath = strings.Replace(localVarPath, "{"+"itemId"+"}", url.PathEscape(parameterValueToString(r.itemId, "itemId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.type_ == nil {
		return nil, reportError("type_ is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
	if r.imageUrl != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "imageUrl", r.imageUrl, "form", "")
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
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
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type RemoteImageAPIGetRemoteImageProvidersRequest struct {
	ctx context.Context
	ApiService *RemoteImageAPIService
	itemId string
}

func (r RemoteImageAPIGetRemoteImageProvidersRequest) Execute() ([]JellyfinImageProviderInfo, *http.Response, error) {
	return r.ApiService.GetRemoteImageProvidersExecute(r)
}

/*
GetRemoteImageProviders Gets available remote image providers for an item.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param itemId Item Id.
 @return RemoteImageAPIGetRemoteImageProvidersRequest
*/
func (a *RemoteImageAPIService) GetRemoteImageProviders(ctx context.Context, itemId string) RemoteImageAPIGetRemoteImageProvidersRequest {
	return RemoteImageAPIGetRemoteImageProvidersRequest{
		ApiService: a,
		ctx: ctx,
		itemId: itemId,
	}
}

// Execute executes the request
//  @return []JellyfinImageProviderInfo
func (a *RemoteImageAPIService) GetRemoteImageProvidersExecute(r RemoteImageAPIGetRemoteImageProvidersRequest) ([]JellyfinImageProviderInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []JellyfinImageProviderInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemoteImageAPIService.GetRemoteImageProviders")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Items/{itemId}/RemoteImages/Providers"
	localVarPath = strings.Replace(localVarPath, "{"+"itemId"+"}", url.PathEscape(parameterValueToString(r.itemId, "itemId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type RemoteImageAPIGetRemoteImagesRequest struct {
	ctx context.Context
	ApiService *RemoteImageAPIService
	itemId string
	type_ *JellyfinImageType
	startIndex *int32
	limit *int32
	providerName *string
	includeAllLanguages *bool
}

// The image type.
func (r RemoteImageAPIGetRemoteImagesRequest) Type_(type_ JellyfinImageType) RemoteImageAPIGetRemoteImagesRequest {
	r.type_ = &type_
	return r
}

// Optional. The record index to start at. All items with a lower index will be dropped from the results.
func (r RemoteImageAPIGetRemoteImagesRequest) StartIndex(startIndex int32) RemoteImageAPIGetRemoteImagesRequest {
	r.startIndex = &startIndex
	return r
}

// Optional. The maximum number of records to return.
func (r RemoteImageAPIGetRemoteImagesRequest) Limit(limit int32) RemoteImageAPIGetRemoteImagesRequest {
	r.limit = &limit
	return r
}

// Optional. The image provider to use.
func (r RemoteImageAPIGetRemoteImagesRequest) ProviderName(providerName string) RemoteImageAPIGetRemoteImagesRequest {
	r.providerName = &providerName
	return r
}

// Optional. Include all languages.
func (r RemoteImageAPIGetRemoteImagesRequest) IncludeAllLanguages(includeAllLanguages bool) RemoteImageAPIGetRemoteImagesRequest {
	r.includeAllLanguages = &includeAllLanguages
	return r
}

func (r RemoteImageAPIGetRemoteImagesRequest) Execute() (*JellyfinRemoteImageResult, *http.Response, error) {
	return r.ApiService.GetRemoteImagesExecute(r)
}

/*
GetRemoteImages Gets available remote images for an item.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param itemId Item Id.
 @return RemoteImageAPIGetRemoteImagesRequest
*/
func (a *RemoteImageAPIService) GetRemoteImages(ctx context.Context, itemId string) RemoteImageAPIGetRemoteImagesRequest {
	return RemoteImageAPIGetRemoteImagesRequest{
		ApiService: a,
		ctx: ctx,
		itemId: itemId,
	}
}

// Execute executes the request
//  @return JellyfinRemoteImageResult
func (a *RemoteImageAPIService) GetRemoteImagesExecute(r RemoteImageAPIGetRemoteImagesRequest) (*JellyfinRemoteImageResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *JellyfinRemoteImageResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemoteImageAPIService.GetRemoteImages")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Items/{itemId}/RemoteImages"
	localVarPath = strings.Replace(localVarPath, "{"+"itemId"+"}", url.PathEscape(parameterValueToString(r.itemId, "itemId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
	}
	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startIndex", r.startIndex, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.providerName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "providerName", r.providerName, "form", "")
	}
	if r.includeAllLanguages != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeAllLanguages", r.includeAllLanguages, "form", "")
	} else {
		var defaultValue bool = false
		r.includeAllLanguages = &defaultValue
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
