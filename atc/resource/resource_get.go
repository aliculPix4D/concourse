package resource

import (
	"context"

	"github.com/concourse/concourse/atc/runtime"
)

func (resource *resource) Get(
	ctx context.Context,
	spec runtime.ProcessSpec,
	runnable runtime.Runnable,
) (runtime.VersionResult, error) {
	var vr runtime.VersionResult

	err := runnable.RunScript(
		ctx,
		spec.Path,
		[]string{spec.Dir},
		resource,
		&vr,
		spec.StderrWriter,
		true,
	)
	if err != nil {
		return runtime.VersionResult{}, err
	}

	return vr, nil
}
