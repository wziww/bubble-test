package docker

import (
	"context"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/registry"
	"github.com/sirupsen/logrus"
)

// ImagesGet ...
func ImagesGet(ctx context.Context) []types.ImageSummary {
	client := New()
	if client != nil {
		defer release(client)
		images, err := client.ImageList(ctx, types.ImageListOptions{})
		if err != nil {
			logrus.Errorln(err)
		}
		return images
	}
	return nil
}

// ImagesSearch ...
func ImagesSearch(ctx context.Context) []registry.SearchResult {
	client := New()
	if client != nil {
		defer release(client)
		images, err := client.ImageSearch(ctx, "golang", types.ImageSearchOptions{
			Limit: 10,
		})
		if err != nil {
			logrus.Errorln(err)
		}
		return images
	}
	return nil
}

// ImagesPull ...
func ImagesPull(ctx context.Context, imageName string) io.ReadCloser {
	client := New()
	if client != nil {
		defer release(client)
		resp, err := client.ImagePull(ctx, imageName, types.ImagePullOptions{})
		if err != nil {
			logrus.Errorln(err)
			return nil
		}
		return resp
	}
	return nil
}
