/*
Jellyfin API

Testing MediaSegmentsAPIService

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

func Test_jellyfin_MediaSegmentsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MediaSegmentsAPIService GetItemSegments", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var itemId string

		resp, httpRes, err := apiClient.MediaSegmentsAPI.GetItemSegments(context.Background(), itemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
