package artifact

import (
	"context"

	"github.com/go-openapi/runtime"
	v2client "github.com/spotmaxtech/goharbor-client/v5/apiv2/internal/api/client"
	"github.com/spotmaxtech/goharbor-client/v5/apiv2/internal/api/client/artifact"
	"github.com/spotmaxtech/goharbor-client/v5/apiv2/model"
	"github.com/spotmaxtech/goharbor-client/v5/apiv2/pkg/config"
	"github.com/spotmaxtech/goharbor-client/v5/apiv2/pkg/errors"
)

// RESTClient is a subclient for handling repository related actions.
type RESTClient struct {
	// Options contains optional configuration when making API calls.
	Options *config.Options

	// The new client of the harbor v2 API
	V2Client *v2client.Harbor

	// AuthInfo contains the auth information that is provided on API calls.
	AuthInfo runtime.ClientAuthInfoWriter
}

func NewClient(v2Client *v2client.Harbor, opts *config.Options, authInfo runtime.ClientAuthInfoWriter) *RESTClient {
	return &RESTClient{
		Options:  opts,
		V2Client: v2Client,
		AuthInfo: authInfo,
	}
}

type Client interface {
	ListArtifacts(ctx context.Context, projectName string) ([]*model.Artifact, error)
}

func (c *RESTClient) ListArtifacts(ctx context.Context, projectName string, repositoryName string) ([]*model.Artifact, error) {
	params := &artifact.ListArtifactsParams{
		Page:           &c.Options.Page,
		PageSize:       &c.Options.PageSize,
		ProjectName:    projectName,
		RepositoryName: repositoryName,
		Q:              &c.Options.Query,
		Sort:           &c.Options.Sort,
		Context:        ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	resp, err := c.V2Client.Artifact.ListArtifacts(params, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerArtifactErrors(err)
	}

	if resp.Payload != nil {
		return resp.Payload, nil
	}

	return nil, &errors.ErrNotFound{}
}
