//go:build integration

package registry

import (
	"context"
	"net/url"
	"testing"

	integrationtest "github.com/spotmaxtech/goharbor-client/v5/apiv1/testing"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/spotmaxtech/goharbor-client/v5/apiv1/internal/api/client"
	model "github.com/spotmaxtech/goharbor-client/v5/apiv1/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	u, _          = url.Parse(integrationtest.Host)
	swaggerClient = client.New(runtimeclient.New(u.Host, u.Path, []string{u.Scheme}), strfmt.Default)
	authInfo      = runtimeclient.BasicAuth(integrationtest.User, integrationtest.Password)
)

func TestAPIRegistryNew(t *testing.T) {
	name := "test-registry"
	registryType := "harbor"
	url := "http://registry-docker-registry:5000/"
	credential := model.RegistryCredential{
		AccessKey:    integrationtest.User,
		AccessSecret: integrationtest.Password,
		Type:         "basic",
	}

	ctx := context.Background()
	c := NewClient(swaggerClient, authInfo)

	r, err := c.NewRegistry(ctx, name, registryType, url, &credential, false)
	defer c.DeleteRegistry(ctx, r)

	require.NoError(t, err)
	assert.Equal(t, name, r.Name)
}

func TestAPIRegistryGet(t *testing.T) {
	name := "test-registry"
	registryType := "harbor"
	url := "http://registry-docker-registry:5000/"
	credential := model.RegistryCredential{
		AccessKey:    integrationtest.User,
		AccessSecret: integrationtest.Password,
		Type:         "basic",
	}

	ctx := context.Background()
	c := NewClient(swaggerClient, authInfo)

	r, err := c.NewRegistry(ctx, name, registryType, url, &credential, false)
	require.NoError(t, err)
	defer c.DeleteRegistry(ctx, r)

	p2, err := c.GetRegistry(ctx, name)
	require.NoError(t, err)
	assert.Equal(t, r, p2)
}

func TestAPIRegistryDelete(t *testing.T) {
	name := "test-registry"
	registryType := "harbor"
	url := "http://registry-docker-registry:5000/"
	credential := model.RegistryCredential{
		AccessKey:    integrationtest.User,
		AccessSecret: integrationtest.Password,
		Type:         "basic",
	}

	ctx := context.Background()
	c := NewClient(swaggerClient, authInfo)

	r, err := c.NewRegistry(ctx, name, registryType, url, &credential, false)
	require.NoError(t, err)

	err = c.DeleteRegistry(ctx, r)
	require.NoError(t, err)

	r, err = c.GetRegistry(ctx, name)
	if assert.Error(t, err) {
		assert.IsType(t, &ErrRegistryNotFound{}, err)
	}
}
