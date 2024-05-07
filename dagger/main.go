// A generated module for Multi functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

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

// Push a multi-arch image
func (m *Multi) MultiPush(ctx context.Context, image string, platforms []string, tag string) (string, error) {
	ctrs := []*Container{} //empty slice of Containers
	for _, platform := range platforms {
		platform := platform
		ctrs = append(ctrs, dag.Container().From(fmt.Sprintf("ttl.sh/%v-%v:%v", image, platform, tag)))
	}
	return dag.Container().Publish(ctx, fmt.Sprintf("ttl.sh/%v:%v", image, tag),
		ContainerPublishOpts{
			PlatformVariants: ctrs,
		})
}
