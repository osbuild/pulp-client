/*
Pulp 3 API

Testing PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIService

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

func Test_pulpclient_PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIService PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var distroBasePath string
		var name string
		var namespace string
		var path string

		resp, httpRes, err := apiClient.PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPI.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexDelete(context.Background(), distroBasePath, name, namespace, path).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIService PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var distroBasePath string
		var path string

		resp, httpRes, err := apiClient.PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPI.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexList(context.Background(), distroBasePath, path).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIService PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var distroBasePath string
		var name string
		var namespace string
		var path string

		resp, httpRes, err := apiClient.PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPI.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexRead(context.Background(), distroBasePath, name, namespace, path).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIService PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var distroBasePath string
		var name string
		var namespace string
		var path string

		resp, httpRes, err := apiClient.PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPI.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexUpdate(context.Background(), distroBasePath, name, namespace, path).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}