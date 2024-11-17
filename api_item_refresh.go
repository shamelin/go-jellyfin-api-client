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


// ItemRefreshAPIService ItemRefreshAPI service
type ItemRefreshAPIService service

type ApiRefreshItemRequest struct {
	ctx context.Context
	ApiService *ItemRefreshAPIService
	itemId string
	metadataRefreshMode *MetadataRefreshMode
	imageRefreshMode *MetadataRefreshMode
	replaceAllMetadata *bool
	replaceAllImages *bool
	regenerateTrickplay *bool
}

// (Optional) Specifies the metadata refresh mode.
func (r ApiRefreshItemRequest) MetadataRefreshMode(metadataRefreshMode MetadataRefreshMode) ApiRefreshItemRequest {
	r.metadataRefreshMode = &metadataRefreshMode
	return r
}

// (Optional) Specifies the image refresh mode.
func (r ApiRefreshItemRequest) ImageRefreshMode(imageRefreshMode MetadataRefreshMode) ApiRefreshItemRequest {
	r.imageRefreshMode = &imageRefreshMode
	return r
}

// (Optional) Determines if metadata should be replaced. Only applicable if mode is FullRefresh.
func (r ApiRefreshItemRequest) ReplaceAllMetadata(replaceAllMetadata bool) ApiRefreshItemRequest {
	r.replaceAllMetadata = &replaceAllMetadata
	return r
}

// (Optional) Determines if images should be replaced. Only applicable if mode is FullRefresh.
func (r ApiRefreshItemRequest) ReplaceAllImages(replaceAllImages bool) ApiRefreshItemRequest {
	r.replaceAllImages = &replaceAllImages
	return r
}

// (Optional) Determines if trickplay images should be replaced. Only applicable if mode is FullRefresh.
func (r ApiRefreshItemRequest) RegenerateTrickplay(regenerateTrickplay bool) ApiRefreshItemRequest {
	r.regenerateTrickplay = &regenerateTrickplay
	return r
}

func (r ApiRefreshItemRequest) Execute() (*http.Response, error) {
	return r.ApiService.RefreshItemExecute(r)
}

/*
RefreshItem Refreshes metadata for an item.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param itemId Item id.
 @return ApiRefreshItemRequest
*/
func (a *ItemRefreshAPIService) RefreshItem(ctx context.Context, itemId string) ApiRefreshItemRequest {
	return ApiRefreshItemRequest{
		ApiService: a,
		ctx: ctx,
		itemId: itemId,
	}
}

// Execute executes the request
func (a *ItemRefreshAPIService) RefreshItemExecute(r ApiRefreshItemRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ItemRefreshAPIService.RefreshItem")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Items/{itemId}/Refresh"
	localVarPath = strings.Replace(localVarPath, "{"+"itemId"+"}", url.PathEscape(parameterValueToString(r.itemId, "itemId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.metadataRefreshMode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "metadataRefreshMode", r.metadataRefreshMode, "form", "")
	} else {
		var defaultValue MetadataRefreshMode = "None"
		r.metadataRefreshMode = &defaultValue
	}
	if r.imageRefreshMode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "imageRefreshMode", r.imageRefreshMode, "form", "")
	} else {
		var defaultValue MetadataRefreshMode = "None"
		r.imageRefreshMode = &defaultValue
	}
	if r.replaceAllMetadata != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "replaceAllMetadata", r.replaceAllMetadata, "form", "")
	} else {
		var defaultValue bool = false
		r.replaceAllMetadata = &defaultValue
	}
	if r.replaceAllImages != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "replaceAllImages", r.replaceAllImages, "form", "")
	} else {
		var defaultValue bool = false
		r.replaceAllImages = &defaultValue
	}
	if r.regenerateTrickplay != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "regenerateTrickplay", r.regenerateTrickplay, "form", "")
	} else {
		var defaultValue bool = false
		r.regenerateTrickplay = &defaultValue
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
			var v ProblemDetails
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
