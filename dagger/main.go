// An example for building single platform and multi platform images

package main

import (
	"context"
	"fmt"
)

type Multi struct{}

// Build and push an image of a certain platform
func (m *Multi) BuildPush(ctx context.Context, platform string, image string, tag string) (string, error) {
	return dag.Container(ContainerOpts{Platform: Platform(platform)}).
		From("alpine:latest").
		WithExec([]string{"echo", platform}).
		Publish(ctx, fmt.Sprintf("ttl.sh/%v-%v:%v", image, platform, tag))
}

// Push a multi-platform image
func (m *Multi) MultiPush(ctx context.Context, image string, platforms []string, tag string) (string, error) {
	ctrs := []*Container{} //empty slice of Containers
	for _, platform := range platforms {
		ctrs = append(ctrs, dag.Container(ContainerOpts{Platform: Platform(platform)}).From(fmt.Sprintf("ttl.sh/%v-%v:%v", image, platform, tag)))
	}
	return dag.Container().Publish(ctx, fmt.Sprintf("ttl.sh/%v:%v", image, tag),
		ContainerPublishOpts{
			PlatformVariants: ctrs,
		})
}
