package ctrd

import (
	"github.com/alibaba/pouch/apis/types"
	"github.com/alibaba/pouch/daemon/containerio"

	specs "github.com/opencontainers/runtime-spec/specs-go"
)

// Container wraps container's info.
type Container struct {
	Info *types.ContainerInfo
	IO   *containerio.IO
	Spec *specs.Spec
}

// Process wraps exec process's info.
type Process struct {
	ContainerID string
	ExecID      string
	IO          *containerio.IO
	P           *specs.Process
}
