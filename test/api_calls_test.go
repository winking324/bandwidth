/*
Bandwidth

Testing CallsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package bandwidth

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/winking324/bandwidth/pkg/bandwidth"
)

func Test_openapi_CallsAPIService(t *testing.T) {

	configuration := bandwidth.NewConfiguration()
	apiClient := bandwidth.NewAPIClient(configuration)

	t.Run("Test CallsAPIService CreateCall", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.CallsAPI.CreateCall(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CallsAPIService GetCallState", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var accountId string
		var callId string

		resp, httpRes, err := apiClient.CallsAPI.GetCallState(context.Background(), accountId, callId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CallsAPIService ListCalls", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.CallsAPI.ListCalls(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CallsAPIService UpdateCall", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var accountId string
		var callId string

		httpRes, err := apiClient.CallsAPI.UpdateCall(context.Background(), accountId, callId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CallsAPIService UpdateCallBxml", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var accountId string
		var callId string

		httpRes, err := apiClient.CallsAPI.UpdateCallBxml(context.Background(), accountId, callId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
