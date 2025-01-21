/*
Bandwidth

Testing MessagesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	openapiclient "github.com/winking324/bandwidth"
)

func Test_openapi_MessagesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MessagesAPIService CreateMessage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.MessagesAPI.CreateMessage(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MessagesAPIService ListMessages", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.MessagesAPI.ListMessages(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
