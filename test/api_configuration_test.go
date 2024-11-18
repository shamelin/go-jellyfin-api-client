/*
Jellyfin API

Testing ConfigurationAPIService

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

func Test_jellyfin_ConfigurationAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ConfigurationAPIService GetConfiguration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ConfigurationAPI.GetConfiguration(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigurationAPIService GetDefaultMetadataOptions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ConfigurationAPI.GetDefaultMetadataOptions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigurationAPIService GetNamedConfiguration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var key string

		resp, httpRes, err := apiClient.ConfigurationAPI.GetNamedConfiguration(context.Background(), key).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigurationAPIService UpdateConfiguration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ConfigurationAPI.UpdateConfiguration(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigurationAPIService UpdateNamedConfiguration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var key string

		httpRes, err := apiClient.ConfigurationAPI.UpdateNamedConfiguration(context.Background(), key).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
