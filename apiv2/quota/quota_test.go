//go:build !integration
// +build !integration

package quota

import (
	"context"
	"encoding/json"
	"strconv"
	"testing"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/goharbor/harbor/src/pkg/quota/types"
	"github.com/stretchr/testify/assert"

	"github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client/quota"
	"github.com/mittwald/goharbor-client/v4/apiv2/mocks"
	modelv2 "github.com/mittwald/goharbor-client/v4/apiv2/model"
	"github.com/mittwald/goharbor-client/v4/apiv2/pkg/common"
	unittesting "github.com/mittwald/goharbor-client/v4/apiv2/pkg/testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var (
	authInfo                       = runtimeclient.BasicAuth("foo", "bar")
	testStorageLimitPositive int64 = 1
	testStorageLimitNegative int64 = -1
	testStorageLimitNull     int64 = 0
	exampleQuotaID           int64 = 1
	exampleProjectID         int64 = 1
	ctx                            = context.Background()
)

func APIandMockClientsForTests() (*RESTClient, *unittesting.MockClients) {
	desiredMockClients := &unittesting.MockClients{
		Quota: mocks.MockQuotaClientService{},
	}

	v2Client := unittesting.BuildV2ClientWithMocks(desiredMockClients)

	cl := NewClient(v2Client, &unittesting.DefaultOpts, authInfo)

	return cl, desiredMockClients
}

func TestRESTClient_GetQuotaByProjectID_Unexpected(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	refID := strconv.Itoa(int(exampleProjectID))
	listParams := &quota.ListQuotasParams{
		PageSize:    &apiClient.Options.PageSize,
		ReferenceID: &refID,
		Sort:        &apiClient.Options.Sort,
		Context:     ctx,
	}

	mockClient.Quota.On("ListQuotas", listParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&quota.ListQuotasOK{}, nil)

	_, err := apiClient.GetQuotaByProjectID(ctx, exampleProjectID)
	require.Error(t, err)
	require.ErrorIs(t, err, &common.ErrQuotaRefNotFound{})

	mockClient.Quota.AssertExpectations(t)
}

func TestRESTClient_GetQuotaByProjectID(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	refID := strconv.Itoa(int(exampleProjectID))

	listParams := &quota.ListQuotasParams{
		PageSize:    &apiClient.Options.PageSize,
		ReferenceID: &refID,
		Sort:        &apiClient.Options.Sort,
		Context:     ctx,
	}

	mockClient.Quota.On("ListQuotas", listParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&quota.ListQuotasOK{
			Payload: []*modelv2.Quota{{
				Hard: types.ResourceList{
					types.ResourceStorage: 10,
				},
				ID: exampleProjectID,
				Ref: modelv2.QuotaRefObject(map[string]interface{}{
					"id": json.Number(strconv.Itoa(1)),
				}),
			}},
		}, nil)

	q, err := apiClient.GetQuotaByProjectID(ctx, exampleProjectID)

	assert.NoError(t, err)

	if assert.NotNil(t, q) {
		require.Equal(t, int64(10), q.Hard["storage"])
	}

	mockClient.Quota.AssertExpectations(t)
}

func TestRESTClient_UpdateStorageQuotaByProjectID(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	t.Run("PositiveLimit", func(t *testing.T) {
		updateParams := &quota.UpdateQuotaParams{
			Hard: &modelv2.QuotaUpdateReq{
				Hard: map[types.ResourceName]int64{
					types.ResourceStorage: testStorageLimitPositive,
				},
			},
			ID:      exampleQuotaID,
			Context: ctx,
		}

		mockClient.Quota.On("UpdateQuota", updateParams,
			mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
			Return(&quota.UpdateQuotaOK{}, nil)

		err := apiClient.UpdateStorageQuotaByProjectID(ctx, exampleProjectID, testStorageLimitPositive)

		assert.NoError(t, err)

		mockClient.Quota.AssertExpectations(t)
	})

	t.Run("NegativeLimit", func(t *testing.T) {
		updateParams := &quota.UpdateQuotaParams{
			Hard: &modelv2.QuotaUpdateReq{
				Hard: map[types.ResourceName]int64{
					types.ResourceStorage: testStorageLimitNegative,
				},
			},
			ID:      exampleQuotaID,
			Context: ctx,
		}

		mockClient.Quota.On("UpdateQuota", updateParams,
			mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
			Return(&quota.UpdateQuotaOK{}, nil)

		err := apiClient.UpdateStorageQuotaByProjectID(ctx, exampleProjectID, testStorageLimitNegative)

		assert.NoError(t, err)

		mockClient.Quota.AssertExpectations(t)
	})

	t.Run("NullLimit", func(t *testing.T) {
		updateParams := &quota.UpdateQuotaParams{
			Hard: &modelv2.QuotaUpdateReq{
				Hard: map[types.ResourceName]int64{
					types.ResourceStorage: testStorageLimitNegative,
				},
			},
			ID:      exampleQuotaID,
			Context: ctx,
		}

		mockClient.Quota.On("UpdateQuota", updateParams,
			mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
			Return(&quota.UpdateQuotaOK{}, nil)

		err := apiClient.UpdateStorageQuotaByProjectID(ctx, exampleProjectID, testStorageLimitNull)

		assert.NoError(t, err)

		mockClient.Quota.AssertExpectations(t)
	})
}
