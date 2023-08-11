/*
Pulp 3 API

Testing RemotesAptAPIService

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

func Test_pulpclient_RemotesAptAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RemotesAptAPIService RemotesDebAptCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RemotesAptAPI.RemotesDebAptCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesAptAPIService RemotesDebAptDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var debAptRemoteHref string

		resp, httpRes, err := apiClient.RemotesAptAPI.RemotesDebAptDelete(context.Background(), debAptRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesAptAPIService RemotesDebAptList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RemotesAptAPI.RemotesDebAptList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesAptAPIService RemotesDebAptPartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var debAptRemoteHref string

		resp, httpRes, err := apiClient.RemotesAptAPI.RemotesDebAptPartialUpdate(context.Background(), debAptRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesAptAPIService RemotesDebAptRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var debAptRemoteHref string

		resp, httpRes, err := apiClient.RemotesAptAPI.RemotesDebAptRead(context.Background(), debAptRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesAptAPIService RemotesDebAptUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var debAptRemoteHref string

		resp, httpRes, err := apiClient.RemotesAptAPI.RemotesDebAptUpdate(context.Background(), debAptRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}