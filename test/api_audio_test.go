/*
Jellyfin API

Testing AudioAPIService

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

func Test_jellyfin_AudioAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AudioAPIService GetAudioStream", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var itemId string

		resp, httpRes, err := apiClient.AudioAPI.GetAudioStream(context.Background(), itemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AudioAPIService GetAudioStreamByContainer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var itemId string
		var container string

		resp, httpRes, err := apiClient.AudioAPI.GetAudioStreamByContainer(context.Background(), itemId, container).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AudioAPIService HeadAudioStream", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var itemId string

		resp, httpRes, err := apiClient.AudioAPI.HeadAudioStream(context.Background(), itemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AudioAPIService HeadAudioStreamByContainer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var itemId string
		var container string

		resp, httpRes, err := apiClient.AudioAPI.HeadAudioStreamByContainer(context.Background(), itemId, container).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
