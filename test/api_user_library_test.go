/*
Jellyfin API

Testing UserLibraryAPIService

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

func Test_jellyfin_UserLibraryAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test UserLibraryAPIService DeleteUserItemRating", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var itemId string

		resp, httpRes, err := apiClient.UserLibraryAPI.DeleteUserItemRating(context.Background(), itemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserLibraryAPIService GetIntros", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var itemId string

		resp, httpRes, err := apiClient.UserLibraryAPI.GetIntros(context.Background(), itemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserLibraryAPIService GetItem", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var itemId string

		resp, httpRes, err := apiClient.UserLibraryAPI.GetItem(context.Background(), itemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserLibraryAPIService GetLatestMedia", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UserLibraryAPI.GetLatestMedia(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserLibraryAPIService GetLocalTrailers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var itemId string

		resp, httpRes, err := apiClient.UserLibraryAPI.GetLocalTrailers(context.Background(), itemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserLibraryAPIService GetRootFolder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UserLibraryAPI.GetRootFolder(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserLibraryAPIService GetSpecialFeatures", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var itemId string

		resp, httpRes, err := apiClient.UserLibraryAPI.GetSpecialFeatures(context.Background(), itemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserLibraryAPIService MarkFavoriteItem", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var itemId string

		resp, httpRes, err := apiClient.UserLibraryAPI.MarkFavoriteItem(context.Background(), itemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserLibraryAPIService UnmarkFavoriteItem", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var itemId string

		resp, httpRes, err := apiClient.UserLibraryAPI.UnmarkFavoriteItem(context.Background(), itemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserLibraryAPIService UpdateUserItemRating", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var itemId string

		resp, httpRes, err := apiClient.UserLibraryAPI.UpdateUserItemRating(context.Background(), itemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
