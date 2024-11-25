/*
Jellyfin API

Testing VideosAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package jellyfin

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/shamelin/go-jellyfin-api-client"
)

func Test_jellyfin_VideosAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test VideosAPIService DeleteAlternateSources", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var itemId string

		httpRes, err := apiClient.VideosAPI.DeleteAlternateSources(context.Background(), itemId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VideosAPIService GetAdditionalPart", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var itemId string

		resp, httpRes, err := apiClient.VideosAPI.GetAdditionalPart(context.Background(), itemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VideosAPIService GetVideoStream", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var itemId string

		resp, httpRes, err := apiClient.VideosAPI.GetVideoStream(context.Background(), itemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VideosAPIService GetVideoStreamByContainer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var itemId string
		var container string

		resp, httpRes, err := apiClient.VideosAPI.GetVideoStreamByContainer(context.Background(), itemId, container).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VideosAPIService HeadVideoStream", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var itemId string

		resp, httpRes, err := apiClient.VideosAPI.HeadVideoStream(context.Background(), itemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VideosAPIService HeadVideoStreamByContainer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var itemId string
		var container string

		resp, httpRes, err := apiClient.VideosAPI.HeadVideoStreamByContainer(context.Background(), itemId, container).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VideosAPIService MergeVersions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.VideosAPI.MergeVersions(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
