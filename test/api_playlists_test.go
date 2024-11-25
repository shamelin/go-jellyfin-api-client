/*
Jellyfin API

Testing PlaylistsAPIService

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

func Test_jellyfin_PlaylistsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PlaylistsAPIService AddItemToPlaylist", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var playlistId string

		httpRes, err := apiClient.PlaylistsAPI.AddItemToPlaylist(context.Background(), playlistId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlaylistsAPIService CreatePlaylist", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PlaylistsAPI.CreatePlaylist(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlaylistsAPIService GetPlaylist", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var playlistId string

		resp, httpRes, err := apiClient.PlaylistsAPI.GetPlaylist(context.Background(), playlistId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlaylistsAPIService GetPlaylistItems", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var playlistId string

		resp, httpRes, err := apiClient.PlaylistsAPI.GetPlaylistItems(context.Background(), playlistId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlaylistsAPIService GetPlaylistUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var playlistId string
		var userId string

		resp, httpRes, err := apiClient.PlaylistsAPI.GetPlaylistUser(context.Background(), playlistId, userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlaylistsAPIService GetPlaylistUsers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var playlistId string

		resp, httpRes, err := apiClient.PlaylistsAPI.GetPlaylistUsers(context.Background(), playlistId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlaylistsAPIService MoveItem", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var playlistId string
		var itemId string
		var newIndex int32

		httpRes, err := apiClient.PlaylistsAPI.MoveItem(context.Background(), playlistId, itemId, newIndex).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlaylistsAPIService RemoveItemFromPlaylist", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var playlistId string

		httpRes, err := apiClient.PlaylistsAPI.RemoveItemFromPlaylist(context.Background(), playlistId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlaylistsAPIService RemoveUserFromPlaylist", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var playlistId string
		var userId string

		httpRes, err := apiClient.PlaylistsAPI.RemoveUserFromPlaylist(context.Background(), playlistId, userId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlaylistsAPIService UpdatePlaylist", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var playlistId string

		httpRes, err := apiClient.PlaylistsAPI.UpdatePlaylist(context.Background(), playlistId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlaylistsAPIService UpdatePlaylistUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var playlistId string
		var userId string

		httpRes, err := apiClient.PlaylistsAPI.UpdatePlaylistUser(context.Background(), playlistId, userId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
