/*
Pulp 3 API

Testing RemotesGemAPIService

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

func Test_pulpclient_RemotesGemAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RemotesGemAPIService RemotesGemGemCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RemotesGemAPI.RemotesGemGemCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesGemAPIService RemotesGemGemDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var gemGemRemoteHref string

		resp, httpRes, err := apiClient.RemotesGemAPI.RemotesGemGemDelete(context.Background(), gemGemRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesGemAPIService RemotesGemGemList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RemotesGemAPI.RemotesGemGemList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesGemAPIService RemotesGemGemPartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var gemGemRemoteHref string

		resp, httpRes, err := apiClient.RemotesGemAPI.RemotesGemGemPartialUpdate(context.Background(), gemGemRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesGemAPIService RemotesGemGemRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var gemGemRemoteHref string

		resp, httpRes, err := apiClient.RemotesGemAPI.RemotesGemGemRead(context.Background(), gemGemRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesGemAPIService RemotesGemGemUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var gemGemRemoteHref string

		resp, httpRes, err := apiClient.RemotesGemAPI.RemotesGemGemUpdate(context.Background(), gemGemRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}