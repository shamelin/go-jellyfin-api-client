/*
Jellyfin API

Testing UserAPIService

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

func Test_jellyfin_UserAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test UserAPIService AuthenticateUserByName", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UserAPI.AuthenticateUserByName(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAPIService AuthenticateWithQuickConnect", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UserAPI.AuthenticateWithQuickConnect(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAPIService CreateUserByName", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UserAPI.CreateUserByName(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAPIService DeleteUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId string

		httpRes, err := apiClient.UserAPI.DeleteUser(context.Background(), userId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAPIService ForgotPassword", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UserAPI.ForgotPassword(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAPIService ForgotPasswordPin", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UserAPI.ForgotPasswordPin(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAPIService GetCurrentUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UserAPI.GetCurrentUser(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAPIService GetPublicUsers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UserAPI.GetPublicUsers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAPIService GetUserById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId string

		resp, httpRes, err := apiClient.UserAPI.GetUserById(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAPIService GetUsers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UserAPI.GetUsers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAPIService UpdateUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.UserAPI.UpdateUser(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAPIService UpdateUserConfiguration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.UserAPI.UpdateUserConfiguration(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAPIService UpdateUserPassword", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.UserAPI.UpdateUserPassword(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAPIService UpdateUserPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId string

		httpRes, err := apiClient.UserAPI.UpdateUserPolicy(context.Background(), userId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
