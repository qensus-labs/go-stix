package factory

import (
	"github.com/qensus-labs/go-stix/stix"
	"github.com/qensus-labs/go-stix/stix/sco"
)

func Process(
	builder *stix.Builder,
	pid int64,
	name string,
	commandLine string,
) (*sco.Process, error) {

	obj, err := sco.NewProcess(
		pid,
		name,
		commandLine,
	)

	if err != nil {
		return nil, err
	}

	builder.Add(obj)

	return obj, nil
}
