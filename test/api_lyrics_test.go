/*
Jellyfin API

Testing LyricsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_LyricsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test LyricsAPIService DeleteLyrics", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var itemId string

		httpRes, err := apiClient.LyricsAPI.DeleteLyrics(context.Background(), itemId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LyricsAPIService DownloadRemoteLyrics", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var itemId string
		var lyricId string

		resp, httpRes, err := apiClient.LyricsAPI.DownloadRemoteLyrics(context.Background(), itemId, lyricId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LyricsAPIService GetLyrics", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var itemId string

		resp, httpRes, err := apiClient.LyricsAPI.GetLyrics(context.Background(), itemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LyricsAPIService GetRemoteLyrics", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var lyricId string

		resp, httpRes, err := apiClient.LyricsAPI.GetRemoteLyrics(context.Background(), lyricId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LyricsAPIService SearchRemoteLyrics", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var itemId string

		resp, httpRes, err := apiClient.LyricsAPI.SearchRemoteLyrics(context.Background(), itemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LyricsAPIService UploadLyrics", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var itemId string

		resp, httpRes, err := apiClient.LyricsAPI.UploadLyrics(context.Background(), itemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
