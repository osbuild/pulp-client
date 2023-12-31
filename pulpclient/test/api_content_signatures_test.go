/*
Pulp 3 API

Testing ContentSignaturesAPIService

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

func Test_pulpclient_ContentSignaturesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContentSignaturesAPIService ContentContainerSignaturesList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContentSignaturesAPI.ContentContainerSignaturesList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentSignaturesAPIService ContentContainerSignaturesRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerManifestSignatureHref string

		resp, httpRes, err := apiClient.ContentSignaturesAPI.ContentContainerSignaturesRead(context.Background(), containerManifestSignatureHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
