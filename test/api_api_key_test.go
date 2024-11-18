/*
Jellyfin API

Testing ApiKeyAPIService

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

func Test_jellyfin_ApiKeyAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ApiKeyAPIService CreateKey", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ApiKeyAPI.CreateKey(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApiKeyAPIService GetKeys", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ApiKeyAPI.GetKeys(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApiKeyAPIService RevokeKey", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var key string

		httpRes, err := apiClient.ApiKeyAPI.RevokeKey(context.Background(), key).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
