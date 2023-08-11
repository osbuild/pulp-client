/*
Pulp 3 API

Testing DistributionsGemAPIService

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

func Test_pulpclient_DistributionsGemAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DistributionsGemAPIService DistributionsGemGemCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DistributionsGemAPI.DistributionsGemGemCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsGemAPIService DistributionsGemGemDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var gemGemDistributionHref string

		resp, httpRes, err := apiClient.DistributionsGemAPI.DistributionsGemGemDelete(context.Background(), gemGemDistributionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsGemAPIService DistributionsGemGemList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DistributionsGemAPI.DistributionsGemGemList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsGemAPIService DistributionsGemGemPartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var gemGemDistributionHref string

		resp, httpRes, err := apiClient.DistributionsGemAPI.DistributionsGemGemPartialUpdate(context.Background(), gemGemDistributionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsGemAPIService DistributionsGemGemRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var gemGemDistributionHref string

		resp, httpRes, err := apiClient.DistributionsGemAPI.DistributionsGemGemRead(context.Background(), gemGemDistributionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsGemAPIService DistributionsGemGemUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var gemGemDistributionHref string

		resp, httpRes, err := apiClient.DistributionsGemAPI.DistributionsGemGemUpdate(context.Background(), gemGemDistributionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}