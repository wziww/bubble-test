package docker

import (
	"context"
	"errors"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/registry"
	"github.com/wziww/bubble-test/common/logger"
)

var (
	// ErrNoClient ...
	ErrNoClient = errors.New("Cant get client")
)

// List ...
func List(ctx context.Context) []types.ImageSummary {
	client := New()
	if client != nil {
		defer release(client)
		images, err := client.ImageList(ctx, types.ImageListOptions{})
		if err != nil {
			logger.Error(err.Error())
		}
		return images
	}
	return nil
}

// Search ...
func Search(ctx context.Context) []registry.SearchResult {
	client := New()
	if client != nil {
		defer release(client)
		images, err := client.ImageSearch(ctx, "golang", types.ImageSearchOptions{
			Limit: 10,
		})
		if err != nil {
			logger.Error(err.Error())
		}
		return images
	}
	return nil
}

// Pull ...
func Pull(ctx context.Context, imageName string) (io.ReadCloser, error) {
	client := New()
	if client != nil {
		defer release(client)
		resp, err := client.ImagePull(ctx, imageName, types.ImagePullOptions{})
		if err != nil {
			logger.Error(err.Error())
			return nil, err
		}
		return resp, nil
	}
	return nil, ErrNoClient
}
