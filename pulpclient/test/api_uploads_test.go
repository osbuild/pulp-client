/*
Pulp 3 API

Testing UploadsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package pulpclient

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/osbuild/pulp-client/pulpclient"
)

func Test_pulpclient_UploadsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test UploadsAPIService UploadsAddRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uploadHref string

		resp, httpRes, err := apiClient.UploadsAPI.UploadsAddRole(context.Background(), uploadHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UploadsAPIService UploadsCommit", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uploadHref string

		resp, httpRes, err := apiClient.UploadsAPI.UploadsCommit(context.Background(), uploadHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UploadsAPIService UploadsCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UploadsAPI.UploadsCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UploadsAPIService UploadsDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uploadHref string

		httpRes, err := apiClient.UploadsAPI.UploadsDelete(context.Background(), uploadHref).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UploadsAPIService UploadsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UploadsAPI.UploadsList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UploadsAPIService UploadsListRoles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uploadHref string

		resp, httpRes, err := apiClient.UploadsAPI.UploadsListRoles(context.Background(), uploadHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UploadsAPIService UploadsMyPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uploadHref string

		resp, httpRes, err := apiClient.UploadsAPI.UploadsMyPermissions(context.Background(), uploadHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UploadsAPIService UploadsRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uploadHref string

		resp, httpRes, err := apiClient.UploadsAPI.UploadsRead(context.Background(), uploadHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UploadsAPIService UploadsRemoveRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uploadHref string

		resp, httpRes, err := apiClient.UploadsAPI.UploadsRemoveRole(context.Background(), uploadHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UploadsAPIService UploadsUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uploadHref string

		resp, httpRes, err := apiClient.UploadsAPI.UploadsUpdate(context.Background(), uploadHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}