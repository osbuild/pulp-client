/*
Pulp 3 API

Testing RepositoriesFileAPIService

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

func Test_pulpclient_RepositoriesFileAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RepositoriesFileAPIService RepositoriesFileFileAddRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesFileAPI.RepositoriesFileFileAddRole(context.Background(), fileFileRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesFileAPIService RepositoriesFileFileCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RepositoriesFileAPI.RepositoriesFileFileCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesFileAPIService RepositoriesFileFileDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesFileAPI.RepositoriesFileFileDelete(context.Background(), fileFileRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesFileAPIService RepositoriesFileFileList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RepositoriesFileAPI.RepositoriesFileFileList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesFileAPIService RepositoriesFileFileListRoles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesFileAPI.RepositoriesFileFileListRoles(context.Background(), fileFileRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesFileAPIService RepositoriesFileFileModify", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesFileAPI.RepositoriesFileFileModify(context.Background(), fileFileRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesFileAPIService RepositoriesFileFileMyPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesFileAPI.RepositoriesFileFileMyPermissions(context.Background(), fileFileRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesFileAPIService RepositoriesFileFilePartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesFileAPI.RepositoriesFileFilePartialUpdate(context.Background(), fileFileRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesFileAPIService RepositoriesFileFileRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesFileAPI.RepositoriesFileFileRead(context.Background(), fileFileRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesFileAPIService RepositoriesFileFileRemoveRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesFileAPI.RepositoriesFileFileRemoveRole(context.Background(), fileFileRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesFileAPIService RepositoriesFileFileSync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesFileAPI.RepositoriesFileFileSync(context.Background(), fileFileRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesFileAPIService RepositoriesFileFileUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesFileAPI.RepositoriesFileFileUpdate(context.Background(), fileFileRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
