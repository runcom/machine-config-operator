package common

import (
	igntypes "github.com/coreos/ignition/config/v3_0/types"
)

// NewIgnConfig returns an empty ignition config with version set as 3.0.0
func NewIgnConfig() igntypes.Config {
	return igntypes.Config{
		Ignition: igntypes.Ignition{
			Version: "3.0.0",
		},
	}
}
