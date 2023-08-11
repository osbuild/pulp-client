/*
Pulp 3 API

Testing DistributionsAnsibleAPIService

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

func Test_pulpclient_DistributionsAnsibleAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DistributionsAnsibleAPIService DistributionsAnsibleAnsibleAddRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ansibleAnsibleDistributionHref string

		resp, httpRes, err := apiClient.DistributionsAnsibleAPI.DistributionsAnsibleAnsibleAddRole(context.Background(), ansibleAnsibleDistributionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsAnsibleAPIService DistributionsAnsibleAnsibleCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DistributionsAnsibleAPI.DistributionsAnsibleAnsibleCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsAnsibleAPIService DistributionsAnsibleAnsibleDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ansibleAnsibleDistributionHref string

		resp, httpRes, err := apiClient.DistributionsAnsibleAPI.DistributionsAnsibleAnsibleDelete(context.Background(), ansibleAnsibleDistributionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsAnsibleAPIService DistributionsAnsibleAnsibleList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DistributionsAnsibleAPI.DistributionsAnsibleAnsibleList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsAnsibleAPIService DistributionsAnsibleAnsibleListRoles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ansibleAnsibleDistributionHref string

		resp, httpRes, err := apiClient.DistributionsAnsibleAPI.DistributionsAnsibleAnsibleListRoles(context.Background(), ansibleAnsibleDistributionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsAnsibleAPIService DistributionsAnsibleAnsibleMyPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ansibleAnsibleDistributionHref string

		resp, httpRes, err := apiClient.DistributionsAnsibleAPI.DistributionsAnsibleAnsibleMyPermissions(context.Background(), ansibleAnsibleDistributionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsAnsibleAPIService DistributionsAnsibleAnsiblePartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ansibleAnsibleDistributionHref string

		resp, httpRes, err := apiClient.DistributionsAnsibleAPI.DistributionsAnsibleAnsiblePartialUpdate(context.Background(), ansibleAnsibleDistributionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsAnsibleAPIService DistributionsAnsibleAnsibleRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ansibleAnsibleDistributionHref string

		resp, httpRes, err := apiClient.DistributionsAnsibleAPI.DistributionsAnsibleAnsibleRead(context.Background(), ansibleAnsibleDistributionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsAnsibleAPIService DistributionsAnsibleAnsibleRemoveRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ansibleAnsibleDistributionHref string

		resp, httpRes, err := apiClient.DistributionsAnsibleAPI.DistributionsAnsibleAnsibleRemoveRole(context.Background(), ansibleAnsibleDistributionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsAnsibleAPIService DistributionsAnsibleAnsibleUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ansibleAnsibleDistributionHref string

		resp, httpRes, err := apiClient.DistributionsAnsibleAPI.DistributionsAnsibleAnsibleUpdate(context.Background(), ansibleAnsibleDistributionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}