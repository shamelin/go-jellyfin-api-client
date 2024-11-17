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
	"reflect"
)


// UniversalAudioAPIService UniversalAudioAPI service
type UniversalAudioAPIService service

type ApiGetUniversalAudioStreamRequest struct {
	ctx context.Context
	ApiService *UniversalAudioAPIService
	itemId string
	container *[]string
	mediaSourceId *string
	deviceId *string
	userId *string
	audioCodec *string
	maxAudioChannels *int32
	transcodingAudioChannels *int32
	maxStreamingBitrate *int32
	audioBitRate *int32
	startTimeTicks *int64
	transcodingContainer *string
	transcodingProtocol *MediaStreamProtocol
	maxAudioSampleRate *int32
	maxAudioBitDepth *int32
	enableRemoteMedia *bool
	enableAudioVbrEncoding *bool
	breakOnNonKeyFrames *bool
	enableRedirection *bool
}

// Optional. The audio container.
func (r ApiGetUniversalAudioStreamRequest) Container(container []string) ApiGetUniversalAudioStreamRequest {
	r.container = &container
	return r
}

// The media version id, if playing an alternate version.
func (r ApiGetUniversalAudioStreamRequest) MediaSourceId(mediaSourceId string) ApiGetUniversalAudioStreamRequest {
	r.mediaSourceId = &mediaSourceId
	return r
}

// The device id of the client requesting. Used to stop encoding processes when needed.
func (r ApiGetUniversalAudioStreamRequest) DeviceId(deviceId string) ApiGetUniversalAudioStreamRequest {
	r.deviceId = &deviceId
	return r
}

// Optional. The user id.
func (r ApiGetUniversalAudioStreamRequest) UserId(userId string) ApiGetUniversalAudioStreamRequest {
	r.userId = &userId
	return r
}

// Optional. The audio codec to transcode to.
func (r ApiGetUniversalAudioStreamRequest) AudioCodec(audioCodec string) ApiGetUniversalAudioStreamRequest {
	r.audioCodec = &audioCodec
	return r
}

// Optional. The maximum number of audio channels.
func (r ApiGetUniversalAudioStreamRequest) MaxAudioChannels(maxAudioChannels int32) ApiGetUniversalAudioStreamRequest {
	r.maxAudioChannels = &maxAudioChannels
	return r
}

// Optional. The number of how many audio channels to transcode to.
func (r ApiGetUniversalAudioStreamRequest) TranscodingAudioChannels(transcodingAudioChannels int32) ApiGetUniversalAudioStreamRequest {
	r.transcodingAudioChannels = &transcodingAudioChannels
	return r
}

// Optional. The maximum streaming bitrate.
func (r ApiGetUniversalAudioStreamRequest) MaxStreamingBitrate(maxStreamingBitrate int32) ApiGetUniversalAudioStreamRequest {
	r.maxStreamingBitrate = &maxStreamingBitrate
	return r
}

// Optional. Specify an audio bitrate to encode to, e.g. 128000. If omitted this will be left to encoder defaults.
func (r ApiGetUniversalAudioStreamRequest) AudioBitRate(audioBitRate int32) ApiGetUniversalAudioStreamRequest {
	r.audioBitRate = &audioBitRate
	return r
}

// Optional. Specify a starting offset, in ticks. 1 tick &#x3D; 10000 ms.
func (r ApiGetUniversalAudioStreamRequest) StartTimeTicks(startTimeTicks int64) ApiGetUniversalAudioStreamRequest {
	r.startTimeTicks = &startTimeTicks
	return r
}

// Optional. The container to transcode to.
func (r ApiGetUniversalAudioStreamRequest) TranscodingContainer(transcodingContainer string) ApiGetUniversalAudioStreamRequest {
	r.transcodingContainer = &transcodingContainer
	return r
}

// Optional. The transcoding protocol.
func (r ApiGetUniversalAudioStreamRequest) TranscodingProtocol(transcodingProtocol MediaStreamProtocol) ApiGetUniversalAudioStreamRequest {
	r.transcodingProtocol = &transcodingProtocol
	return r
}

// Optional. The maximum audio sample rate.
func (r ApiGetUniversalAudioStreamRequest) MaxAudioSampleRate(maxAudioSampleRate int32) ApiGetUniversalAudioStreamRequest {
	r.maxAudioSampleRate = &maxAudioSampleRate
	return r
}

// Optional. The maximum audio bit depth.
func (r ApiGetUniversalAudioStreamRequest) MaxAudioBitDepth(maxAudioBitDepth int32) ApiGetUniversalAudioStreamRequest {
	r.maxAudioBitDepth = &maxAudioBitDepth
	return r
}

// Optional. Whether to enable remote media.
func (r ApiGetUniversalAudioStreamRequest) EnableRemoteMedia(enableRemoteMedia bool) ApiGetUniversalAudioStreamRequest {
	r.enableRemoteMedia = &enableRemoteMedia
	return r
}

// Optional. Whether to enable Audio Encoding.
func (r ApiGetUniversalAudioStreamRequest) EnableAudioVbrEncoding(enableAudioVbrEncoding bool) ApiGetUniversalAudioStreamRequest {
	r.enableAudioVbrEncoding = &enableAudioVbrEncoding
	return r
}

// Optional. Whether to break on non key frames.
func (r ApiGetUniversalAudioStreamRequest) BreakOnNonKeyFrames(breakOnNonKeyFrames bool) ApiGetUniversalAudioStreamRequest {
	r.breakOnNonKeyFrames = &breakOnNonKeyFrames
	return r
}

// Whether to enable redirection. Defaults to true.
func (r ApiGetUniversalAudioStreamRequest) EnableRedirection(enableRedirection bool) ApiGetUniversalAudioStreamRequest {
	r.enableRedirection = &enableRedirection
	return r
}

func (r ApiGetUniversalAudioStreamRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.GetUniversalAudioStreamExecute(r)
}

/*
GetUniversalAudioStream Gets an audio stream.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param itemId The item id.
 @return ApiGetUniversalAudioStreamRequest
*/
func (a *UniversalAudioAPIService) GetUniversalAudioStream(ctx context.Context, itemId string) ApiGetUniversalAudioStreamRequest {
	return ApiGetUniversalAudioStreamRequest{
		ApiService: a,
		ctx: ctx,
		itemId: itemId,
	}
}

// Execute executes the request
//  @return *os.File
func (a *UniversalAudioAPIService) GetUniversalAudioStreamExecute(r ApiGetUniversalAudioStreamRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UniversalAudioAPIService.GetUniversalAudioStream")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Audio/{itemId}/universal"
	localVarPath = strings.Replace(localVarPath, "{"+"itemId"+"}", url.PathEscape(parameterValueToString(r.itemId, "itemId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.container != nil {
		t := *r.container
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "container", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "container", t, "form", "multi")
		}
	}
	if r.mediaSourceId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mediaSourceId", r.mediaSourceId, "form", "")
	}
	if r.deviceId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "deviceId", r.deviceId, "form", "")
	}
	if r.userId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "userId", r.userId, "form", "")
	}
	if r.audioCodec != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "audioCodec", r.audioCodec, "form", "")
	}
	if r.maxAudioChannels != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maxAudioChannels", r.maxAudioChannels, "form", "")
	}
	if r.transcodingAudioChannels != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "transcodingAudioChannels", r.transcodingAudioChannels, "form", "")
	}
	if r.maxStreamingBitrate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maxStreamingBitrate", r.maxStreamingBitrate, "form", "")
	}
	if r.audioBitRate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "audioBitRate", r.audioBitRate, "form", "")
	}
	if r.startTimeTicks != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startTimeTicks", r.startTimeTicks, "form", "")
	}
	if r.transcodingContainer != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "transcodingContainer", r.transcodingContainer, "form", "")
	}
	if r.transcodingProtocol != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "transcodingProtocol", r.transcodingProtocol, "form", "")
	}
	if r.maxAudioSampleRate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maxAudioSampleRate", r.maxAudioSampleRate, "form", "")
	}
	if r.maxAudioBitDepth != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maxAudioBitDepth", r.maxAudioBitDepth, "form", "")
	}
	if r.enableRemoteMedia != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "enableRemoteMedia", r.enableRemoteMedia, "form", "")
	}
	if r.enableAudioVbrEncoding != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "enableAudioVbrEncoding", r.enableAudioVbrEncoding, "form", "")
	} else {
		var defaultValue bool = true
		r.enableAudioVbrEncoding = &defaultValue
	}
	if r.breakOnNonKeyFrames != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "breakOnNonKeyFrames", r.breakOnNonKeyFrames, "form", "")
	} else {
		var defaultValue bool = false
		r.breakOnNonKeyFrames = &defaultValue
	}
	if r.enableRedirection != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "enableRedirection", r.enableRedirection, "form", "")
	} else {
		var defaultValue bool = true
		r.enableRedirection = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"audio/*", "application/json", "application/json; profile=CamelCase", "application/json; profile=PascalCase"}

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
			var v ProblemDetails
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

type ApiHeadUniversalAudioStreamRequest struct {
	ctx context.Context
	ApiService *UniversalAudioAPIService
	itemId string
	container *[]string
	mediaSourceId *string
	deviceId *string
	userId *string
	audioCodec *string
	maxAudioChannels *int32
	transcodingAudioChannels *int32
	maxStreamingBitrate *int32
	audioBitRate *int32
	startTimeTicks *int64
	transcodingContainer *string
	transcodingProtocol *MediaStreamProtocol
	maxAudioSampleRate *int32
	maxAudioBitDepth *int32
	enableRemoteMedia *bool
	enableAudioVbrEncoding *bool
	breakOnNonKeyFrames *bool
	enableRedirection *bool
}

// Optional. The audio container.
func (r ApiHeadUniversalAudioStreamRequest) Container(container []string) ApiHeadUniversalAudioStreamRequest {
	r.container = &container
	return r
}

// The media version id, if playing an alternate version.
func (r ApiHeadUniversalAudioStreamRequest) MediaSourceId(mediaSourceId string) ApiHeadUniversalAudioStreamRequest {
	r.mediaSourceId = &mediaSourceId
	return r
}

// The device id of the client requesting. Used to stop encoding processes when needed.
func (r ApiHeadUniversalAudioStreamRequest) DeviceId(deviceId string) ApiHeadUniversalAudioStreamRequest {
	r.deviceId = &deviceId
	return r
}

// Optional. The user id.
func (r ApiHeadUniversalAudioStreamRequest) UserId(userId string) ApiHeadUniversalAudioStreamRequest {
	r.userId = &userId
	return r
}

// Optional. The audio codec to transcode to.
func (r ApiHeadUniversalAudioStreamRequest) AudioCodec(audioCodec string) ApiHeadUniversalAudioStreamRequest {
	r.audioCodec = &audioCodec
	return r
}

// Optional. The maximum number of audio channels.
func (r ApiHeadUniversalAudioStreamRequest) MaxAudioChannels(maxAudioChannels int32) ApiHeadUniversalAudioStreamRequest {
	r.maxAudioChannels = &maxAudioChannels
	return r
}

// Optional. The number of how many audio channels to transcode to.
func (r ApiHeadUniversalAudioStreamRequest) TranscodingAudioChannels(transcodingAudioChannels int32) ApiHeadUniversalAudioStreamRequest {
	r.transcodingAudioChannels = &transcodingAudioChannels
	return r
}

// Optional. The maximum streaming bitrate.
func (r ApiHeadUniversalAudioStreamRequest) MaxStreamingBitrate(maxStreamingBitrate int32) ApiHeadUniversalAudioStreamRequest {
	r.maxStreamingBitrate = &maxStreamingBitrate
	return r
}

// Optional. Specify an audio bitrate to encode to, e.g. 128000. If omitted this will be left to encoder defaults.
func (r ApiHeadUniversalAudioStreamRequest) AudioBitRate(audioBitRate int32) ApiHeadUniversalAudioStreamRequest {
	r.audioBitRate = &audioBitRate
	return r
}

// Optional. Specify a starting offset, in ticks. 1 tick &#x3D; 10000 ms.
func (r ApiHeadUniversalAudioStreamRequest) StartTimeTicks(startTimeTicks int64) ApiHeadUniversalAudioStreamRequest {
	r.startTimeTicks = &startTimeTicks
	return r
}

// Optional. The container to transcode to.
func (r ApiHeadUniversalAudioStreamRequest) TranscodingContainer(transcodingContainer string) ApiHeadUniversalAudioStreamRequest {
	r.transcodingContainer = &transcodingContainer
	return r
}

// Optional. The transcoding protocol.
func (r ApiHeadUniversalAudioStreamRequest) TranscodingProtocol(transcodingProtocol MediaStreamProtocol) ApiHeadUniversalAudioStreamRequest {
	r.transcodingProtocol = &transcodingProtocol
	return r
}

// Optional. The maximum audio sample rate.
func (r ApiHeadUniversalAudioStreamRequest) MaxAudioSampleRate(maxAudioSampleRate int32) ApiHeadUniversalAudioStreamRequest {
	r.maxAudioSampleRate = &maxAudioSampleRate
	return r
}

// Optional. The maximum audio bit depth.
func (r ApiHeadUniversalAudioStreamRequest) MaxAudioBitDepth(maxAudioBitDepth int32) ApiHeadUniversalAudioStreamRequest {
	r.maxAudioBitDepth = &maxAudioBitDepth
	return r
}

// Optional. Whether to enable remote media.
func (r ApiHeadUniversalAudioStreamRequest) EnableRemoteMedia(enableRemoteMedia bool) ApiHeadUniversalAudioStreamRequest {
	r.enableRemoteMedia = &enableRemoteMedia
	return r
}

// Optional. Whether to enable Audio Encoding.
func (r ApiHeadUniversalAudioStreamRequest) EnableAudioVbrEncoding(enableAudioVbrEncoding bool) ApiHeadUniversalAudioStreamRequest {
	r.enableAudioVbrEncoding = &enableAudioVbrEncoding
	return r
}

// Optional. Whether to break on non key frames.
func (r ApiHeadUniversalAudioStreamRequest) BreakOnNonKeyFrames(breakOnNonKeyFrames bool) ApiHeadUniversalAudioStreamRequest {
	r.breakOnNonKeyFrames = &breakOnNonKeyFrames
	return r
}

// Whether to enable redirection. Defaults to true.
func (r ApiHeadUniversalAudioStreamRequest) EnableRedirection(enableRedirection bool) ApiHeadUniversalAudioStreamRequest {
	r.enableRedirection = &enableRedirection
	return r
}

func (r ApiHeadUniversalAudioStreamRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.HeadUniversalAudioStreamExecute(r)
}

/*
HeadUniversalAudioStream Gets an audio stream.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param itemId The item id.
 @return ApiHeadUniversalAudioStreamRequest
*/
func (a *UniversalAudioAPIService) HeadUniversalAudioStream(ctx context.Context, itemId string) ApiHeadUniversalAudioStreamRequest {
	return ApiHeadUniversalAudioStreamRequest{
		ApiService: a,
		ctx: ctx,
		itemId: itemId,
	}
}

// Execute executes the request
//  @return *os.File
func (a *UniversalAudioAPIService) HeadUniversalAudioStreamExecute(r ApiHeadUniversalAudioStreamRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodHead
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UniversalAudioAPIService.HeadUniversalAudioStream")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Audio/{itemId}/universal"
	localVarPath = strings.Replace(localVarPath, "{"+"itemId"+"}", url.PathEscape(parameterValueToString(r.itemId, "itemId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.container != nil {
		t := *r.container
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "container", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "container", t, "form", "multi")
		}
	}
	if r.mediaSourceId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mediaSourceId", r.mediaSourceId, "form", "")
	}
	if r.deviceId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "deviceId", r.deviceId, "form", "")
	}
	if r.userId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "userId", r.userId, "form", "")
	}
	if r.audioCodec != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "audioCodec", r.audioCodec, "form", "")
	}
	if r.maxAudioChannels != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maxAudioChannels", r.maxAudioChannels, "form", "")
	}
	if r.transcodingAudioChannels != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "transcodingAudioChannels", r.transcodingAudioChannels, "form", "")
	}
	if r.maxStreamingBitrate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maxStreamingBitrate", r.maxStreamingBitrate, "form", "")
	}
	if r.audioBitRate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "audioBitRate", r.audioBitRate, "form", "")
	}
	if r.startTimeTicks != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startTimeTicks", r.startTimeTicks, "form", "")
	}
	if r.transcodingContainer != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "transcodingContainer", r.transcodingContainer, "form", "")
	}
	if r.transcodingProtocol != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "transcodingProtocol", r.transcodingProtocol, "form", "")
	}
	if r.maxAudioSampleRate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maxAudioSampleRate", r.maxAudioSampleRate, "form", "")
	}
	if r.maxAudioBitDepth != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maxAudioBitDepth", r.maxAudioBitDepth, "form", "")
	}
	if r.enableRemoteMedia != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "enableRemoteMedia", r.enableRemoteMedia, "form", "")
	}
	if r.enableAudioVbrEncoding != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "enableAudioVbrEncoding", r.enableAudioVbrEncoding, "form", "")
	} else {
		var defaultValue bool = true
		r.enableAudioVbrEncoding = &defaultValue
	}
	if r.breakOnNonKeyFrames != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "breakOnNonKeyFrames", r.breakOnNonKeyFrames, "form", "")
	} else {
		var defaultValue bool = false
		r.breakOnNonKeyFrames = &defaultValue
	}
	if r.enableRedirection != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "enableRedirection", r.enableRedirection, "form", "")
	} else {
		var defaultValue bool = true
		r.enableRedirection = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"audio/*", "application/json", "application/json; profile=CamelCase", "application/json; profile=PascalCase"}

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
			var v ProblemDetails
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
