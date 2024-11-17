/*
Jellyfin API

Testing CollectionAPIService

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

func Test_openapi_CollectionAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CollectionAPIService AddToCollection", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var collectionId string

		httpRes, err := apiClient.CollectionAPI.AddToCollection(context.Background(), collectionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CollectionAPIService CreateCollection", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CollectionAPI.CreateCollection(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CollectionAPIService RemoveFromCollection", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var collectionId string

		httpRes, err := apiClient.CollectionAPI.RemoveFromCollection(context.Background(), collectionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
