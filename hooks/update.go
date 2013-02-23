// Update - push constraints
// - for production env push reference must be a tag
package main

import (
    "github.com/rochacon/git-hooks-to-run-fabric/misc/util"
    "os"
)

func main() {
    // Remove GIT_DIR of OS environment, if present
    if os.Getenv("GIT_DIR") != "" {
        os.Setenv("GIT_DIR", "")
    }

    // Get environment based on bare path
    environment, err := util.GetEnvironmentFromPath(os.Getwd())
    if err != nil {
        util.Abort(fmt."Unknown environment: %s", environment, nil)
    }

    // Parse git_ref being pushed
    ref_type, _ := util.ParseRef(os.Args[1])

    // Production only accepts tags
    if environment == "production" and ref_type != "tags" {
        util.Abort("Only tags can be deployed to production", nil)
    }
}
