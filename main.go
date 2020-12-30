// Copyright (C) 2017 Battelle Memorial Institute
// Copyright (C) 2018 Joey Ma <majunjiev@gmail.com>
// All rights reserved.
//
// This software may be modified and distributed under the terms
// of the BSD-2 license.  See the LICENSE file for details.

package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/kenmoini/terraform-provider-ovirt/ovirt"
)

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary
	version string = "dev"

	// goreleaser can also pass the specific commit if you want
	// commit  string = ""
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: ovirt.Provider,
	})
}
