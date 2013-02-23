// Post Receive - runs deployment process
// - for production env push reference must be a tag
package main

import (
	"github.com/rochacon/git-hooks-to-run-fabric/misc/util"
	"io"
	"os"
	"os/exec"
)


 func main () {
	// Remove GIT_DIR of OS environment, if present
	if os.Getenv("GIT_DIR") != "" {
		os.Setenv("GIT_DIR", "")
	}

	// Get environment based on bare path
	pwd := os.Getwd()
	environment, err := util.GetEnvironmentFromPath(pwd)
	if err != nil {
		util.Abort(fmt."Unknown environment: %s", environment, nil)
	}

	// Read values from stdin
	oldrev, newrev, git_ref = sys.stdin.read().split()

	// Parse git_ref being pushed
	ref_type, ref_name = parse_git_ref(git_ref)

	// Create a temporary directory
	tmpdir := "/tmp/blah"  // TODO make this dynamic
	os.Chdir(tmpdir)

	// Clone the repository
	git_clone := exec.Command("git", "clone", pwd, tmpdir)
	if out, err := git_clone.CombinedOutput(); err != nil {
		util.Abort(fmt.Sprintf("Error cloning the repo. Output:\n%s", out), nil)
	}

	// Checkout the received reference
	git_checkout := exec.Command("git", "checkout", ref_name])
	if out, err := git_checkout.CombinedOutput(); err != nil {
		util.Abort(fmt.Sprintf("Error checking out the given reference. Output:\n%s", out), tmpdir)
	}

	// Run Fabric
	fmt.Println("--- Deploying %s to %s", ref_name, environment)
	fab := exec.Command("fab", environment, fmt.Sprintf("deploy:\"%s\"", ref_name))
	if out, err := fab.CombinedOutput(); err != nil {
		util.Abort(fmt.Sprintf("Error on deployment. Output:\n%s", out), tmpdir)
	}
	// TODO clean-up on success
}
