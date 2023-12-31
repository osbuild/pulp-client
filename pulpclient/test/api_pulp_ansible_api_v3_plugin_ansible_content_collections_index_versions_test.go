/*
Pulp 3 API

Testing PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexVersionsAPIService

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

func Test_pulpclient_PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexVersionsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexVersionsAPIService PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexVersionsDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var distroBasePath string
		var name string
		var namespace string
		var path string
		var version string

		resp, httpRes, err := apiClient.PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexVersionsAPI.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexVersionsDelete(context.Background(), distroBasePath, name, namespace, path, version).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexVersionsAPIService PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexVersionsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var distroBasePath string
		var name string
		var namespace string
		var path string

		resp, httpRes, err := apiClient.PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexVersionsAPI.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexVersionsList(context.Background(), distroBasePath, name, namespace, path).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexVersionsAPIService PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexVersionsRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var distroBasePath string
		var name string
		var namespace string
		var path string
		var version string

		resp, httpRes, err := apiClient.PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexVersionsAPI.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexVersionsRead(context.Background(), distroBasePath, name, namespace, path, version).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
