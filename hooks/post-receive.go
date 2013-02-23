// Post Receive - runs deployment process
// - for production env push reference must be a tag
package main

import (
	"github.com/rochacon/git-hooks-to-run-fabric/misc"
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	// "strings"
)


func main () {
	// Remove GIT_DIR of OS environment, if present
	if os.Getenv("GIT_DIR") != "" {
		os.Setenv("GIT_DIR", "")
	}

	// Get environment based on bare path
	pwd, _ := os.Getwd()
	environment, err := util.GetEnvironmentFromPath(pwd)
	if err != nil {
		log.Fatalf("Unknown environment")
	}

	// Read values from stdin
	stdin := bufio.NewReader(os.Stdin)
	line, err := stdin.ReadString('\n')
	log.Println(line)
	if err != nil {
		log.Fatalf("Error reading from stdin")
	}
	// oldrev := line[0]
	// newrev := line[1]
	// git_reference := strings.TrimRight(line[2], "\n")

	// Parse git reference being pushed
	_, ref_name := util.ParseRef("refs/head/master") //git_reference)

	// Create a temporary directory
	tmpdir := "/tmp/blah"  // TODO make this dynamic
	defer util.CleanUp(tmpdir)
	os.Chdir(tmpdir)

	// Clone the repository
	git_clone := exec.Command("git", "clone", pwd, tmpdir)
	if out, err := git_clone.CombinedOutput(); err != nil {
		util.Abort(fmt.Sprintf("Error cloning the repo. Output:\n%s", out))
	}

	// Checkout the received reference
	git_checkout := exec.Command("git", "checkout", ref_name)
	if out, err := git_checkout.CombinedOutput(); err != nil {
		util.Abort(fmt.Sprintf("Error checking out the given reference. Output:\n%s", out))
	}

	// Run Fabric
	fmt.Println("--- Deploying %s to %s", ref_name, environment)
	fab := exec.Command("fab", environment, fmt.Sprintf("deploy:\"%s\"", ref_name))
	if out, err := fab.CombinedOutput(); err != nil {
		util.Abort(fmt.Sprintf("Error on deployment. Output:\n%s", out))
	}
}
