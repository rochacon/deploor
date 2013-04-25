// Update - push constraints
// - for production env push reference must be a tag
package main

import (
	"fmt"
	"github.com/rochacon/deploor/misc"
	"os"
)

func main() {
	// Remove GIT_DIR of OS environment, if present
	if os.Getenv("GIT_DIR") != "" {
		os.Setenv("GIT_DIR", "")
	}

	// Get environment based on bare path
	pwd, _ := os.Getwd()
	environment, err := util.GetEnvironmentFromPath(pwd)
	if err != nil {
		util.Abort(fmt.Sprintf("Unknown environment: %s", environment))
	}

	// Parse git_ref being pushed
	ref_type, _ := util.ParseRef(os.Args[1])

	// Production only accepts tags
	if environment == "production" && ref_type != "tags" {
		util.Abort("Only tags can be deployed to production")
	}
}
