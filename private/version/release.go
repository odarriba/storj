// Copyright (C) 2020 Storj Labs, Inc.
// See LICENSE for copying information.

package version

import _ "unsafe" // needed for go:linkname

//go:linkname buildTimestamp storj.io/private/version.buildTimestamp
var buildTimestamp string = "1645126121"

//go:linkname buildCommitHash storj.io/private/version.buildCommitHash
var buildCommitHash string = "993b8551046c17ab39ac7aaec66109262c1b41f9"

//go:linkname buildVersion storj.io/private/version.buildVersion
var buildVersion string = "v1.49.1-rc"

//go:linkname buildRelease storj.io/private/version.buildRelease
var buildRelease string = "true"

// ensure that linter understands that the variables are being used.
func init() { use(buildTimestamp, buildCommitHash, buildVersion, buildRelease) }

func use(...interface{}) {}
