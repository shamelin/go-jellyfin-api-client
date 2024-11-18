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


// PersonsAPIService PersonsAPI service
type PersonsAPIService service

type ApiGetPersonRequest struct {
	ctx context.Context
	ApiService *PersonsAPIService
	name string
	userId *string
}

// Optional. Filter by user id, and attach user data.
func (r ApiGetPersonRequest) UserId(userId string) ApiGetPersonRequest {
	r.userId = &userId
	return r
}

func (r ApiGetPersonRequest) Execute() (*JellyfinBaseItemDto, *http.Response, error) {
	return r.ApiService.GetPersonExecute(r)
}

/*
GetPerson Get person by name.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name Person name.
 @return ApiGetPersonRequest
*/
func (a *PersonsAPIService) GetPerson(ctx context.Context, name string) ApiGetPersonRequest {
	return ApiGetPersonRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
	}
}

// Execute executes the request
//  @return JellyfinBaseItemDto
func (a *PersonsAPIService) GetPersonExecute(r ApiGetPersonRequest) (*JellyfinBaseItemDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *JellyfinBaseItemDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PersonsAPIService.GetPerson")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Persons/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)

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

type ApiGetPersonsRequest struct {
	ctx context.Context
	ApiService *PersonsAPIService
	limit *int32
	searchTerm *string
	fields *[]JellyfinItemFields
	filters *[]JellyfinItemFilter
	isFavorite *bool
	enableUserData *bool
	imageTypeLimit *int32
	enableImageTypes *[]JellyfinImageType
	excludePersonTypes *[]string
	personTypes *[]string
	appearsInItemId *string
	userId *string
	enableImages *bool
}

// Optional. The maximum number of records to return.
func (r ApiGetPersonsRequest) Limit(limit int32) ApiGetPersonsRequest {
	r.limit = &limit
	return r
}

// The search term.
func (r ApiGetPersonsRequest) SearchTerm(searchTerm string) ApiGetPersonsRequest {
	r.searchTerm = &searchTerm
	return r
}

// Optional. Specify additional fields of information to return in the output.
func (r ApiGetPersonsRequest) Fields(fields []JellyfinItemFields) ApiGetPersonsRequest {
	r.fields = &fields
	return r
}

// Optional. Specify additional filters to apply.
func (r ApiGetPersonsRequest) Filters(filters []JellyfinItemFilter) ApiGetPersonsRequest {
	r.filters = &filters
	return r
}

// Optional filter by items that are marked as favorite, or not. userId is required.
func (r ApiGetPersonsRequest) IsFavorite(isFavorite bool) ApiGetPersonsRequest {
	r.isFavorite = &isFavorite
	return r
}

// Optional, include user data.
func (r ApiGetPersonsRequest) EnableUserData(enableUserData bool) ApiGetPersonsRequest {
	r.enableUserData = &enableUserData
	return r
}

// Optional, the max number of images to return, per image type.
func (r ApiGetPersonsRequest) ImageTypeLimit(imageTypeLimit int32) ApiGetPersonsRequest {
	r.imageTypeLimit = &imageTypeLimit
	return r
}

// Optional. The image types to include in the output.
func (r ApiGetPersonsRequest) EnableImageTypes(enableImageTypes []JellyfinImageType) ApiGetPersonsRequest {
	r.enableImageTypes = &enableImageTypes
	return r
}

// Optional. If specified results will be filtered to exclude those containing the specified PersonType. Allows multiple, comma-delimited.
func (r ApiGetPersonsRequest) ExcludePersonTypes(excludePersonTypes []string) ApiGetPersonsRequest {
	r.excludePersonTypes = &excludePersonTypes
	return r
}

// Optional. If specified results will be filtered to include only those containing the specified PersonType. Allows multiple, comma-delimited.
func (r ApiGetPersonsRequest) PersonTypes(personTypes []string) ApiGetPersonsRequest {
	r.personTypes = &personTypes
	return r
}

// Optional. If specified, person results will be filtered on items related to said persons.
func (r ApiGetPersonsRequest) AppearsInItemId(appearsInItemId string) ApiGetPersonsRequest {
	r.appearsInItemId = &appearsInItemId
	return r
}

// User id.
func (r ApiGetPersonsRequest) UserId(userId string) ApiGetPersonsRequest {
	r.userId = &userId
	return r
}

// Optional, include image information in output.
func (r ApiGetPersonsRequest) EnableImages(enableImages bool) ApiGetPersonsRequest {
	r.enableImages = &enableImages
	return r
}

func (r ApiGetPersonsRequest) Execute() (*JellyfinBaseItemDtoQueryResult, *http.Response, error) {
	return r.ApiService.GetPersonsExecute(r)
}

/*
GetPersons Gets all persons.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetPersonsRequest
*/
func (a *PersonsAPIService) GetPersons(ctx context.Context) ApiGetPersonsRequest {
	return ApiGetPersonsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return JellyfinBaseItemDtoQueryResult
func (a *PersonsAPIService) GetPersonsExecute(r ApiGetPersonsRequest) (*JellyfinBaseItemDtoQueryResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *JellyfinBaseItemDtoQueryResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PersonsAPIService.GetPersons")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Persons"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.searchTerm != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchTerm", r.searchTerm, "form", "")
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
	if r.filters != nil {
		t := *r.filters
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "filters", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "filters", t, "form", "multi")
		}
	}
	if r.isFavorite != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "isFavorite", r.isFavorite, "form", "")
	}
	if r.enableUserData != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "enableUserData", r.enableUserData, "form", "")
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
	if r.excludePersonTypes != nil {
		t := *r.excludePersonTypes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "excludePersonTypes", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "excludePersonTypes", t, "form", "multi")
		}
	}
	if r.personTypes != nil {
		t := *r.personTypes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "personTypes", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "personTypes", t, "form", "multi")
		}
	}
	if r.appearsInItemId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "appearsInItemId", r.appearsInItemId, "form", "")
	}
	if r.userId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "userId", r.userId, "form", "")
	}
	if r.enableImages != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "enableImages", r.enableImages, "form", "")
	} else {
		var defaultValue bool = true
		r.enableImages = &defaultValue
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
