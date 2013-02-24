// Post Receive - runs deployment process
// - for production env push reference must be a tag
package main

import (
	"github.com/rochacon/git-hooks-to-run-fabric/misc"
	"bufio"
	"fmt"
	"log"
	"io"
	"os"
	"os/exec"
	"path"
	"strings"
)

func MkTempDir(reference string) (string, error) {
	folder := path.Join(os.TempDir(), reference)
	if err := os.MkdirAll(folder, 0700); err != nil {
		if os.IsExist(err) {
			os.RemoveAll(path.Join(folder, "*"))
		} else {
			return folder, fmt.Errorf("Impossible to create temporary folder")
		}
	}
	return folder, nil
}

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
	if err != nil {
		log.Fatalf("Error reading from stdin")
	}
	line_fields := strings.Split(line, " ")
	// old_revision := line[0]
	new_revision := line_fields[1]
	git_reference := strings.TrimRight(line_fields[2], "\n")

	// Parse git reference being pushed
	_, ref_name := util.ParseRef(git_reference)

	// Create a temporary directory
	tmpdir, err := MkTempDir(new_revision)
	if err != nil {
		util.Abort(fmt.Sprintf("%s", err))
	}
	defer util.CleanUp(tmpdir)
	if err := os.Chdir(tmpdir); err != nil {
		util.Abort(fmt.Sprintf("%s", err))
	}

	// Clone the repository
	git_clone := exec.Command("git", "clone", pwd, tmpdir)
	if out, err := git_clone.CombinedOutput(); err != nil {
		util.Abort(fmt.Sprintf("Error cloning the repo. Output:\n%s", out))
	}

	// Checkout the received reference
	os.Setenv("GIT_DIR", path.Join(tmpdir, ".git"))  // FIXME os.Chdir should be enough
	git_checkout := exec.Command("git", "checkout", ref_name)
	if out, err := git_checkout.CombinedOutput(); err != nil {
		util.Abort(fmt.Sprintf("Error checking out the given reference. Output:\n%s", out))
	}

	// Run Fabric
	log.Printf("--- Deploying %s to %s\n", ref_name, environment)
	// FIXME The command must stream its output, maybe os.StdoutPipe
	fab := exec.Command("fab", environment, fmt.Sprintf("deploy:'%s'", ref_name))
	stdout, err := fab.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := fab.Start(); err != nil {
		log.Fatal(err)
	}

	bufout := bufio.NewReader(stdout)
	for {
		line, err := bufout.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		log.Printf(line)
	}

	if err := fab.Wait(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Bye")
}
