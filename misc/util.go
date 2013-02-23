package "github.com/rochacon/git-hooks-to-run-fabric/misc/util"

import (
	"fmt"
	"log"
	"os/exec"
)

func Abort(message string, clone string) {
	if clone != nil {
		// TODO delete clone folder
	}
	log.Fatalf(message)
}

// Return the reference type and name
// Example:
// 	refs/head/master -> branch, master
// 	refs/tags/1.0 -> tags, 1.0
func ParseRef(reference string) {
	ref := strings.Split(reference, "/")

	if len(ref) < 3 {
		log.Fatalf("Invalid reference")
	}

	// Otherwise, return the branch name
	return ref[1], ref[2]
}

// Based on a repository path, get the environment
// Example: /srv/git/staging/project.git
func GetEnviromentFromPath(path string) string, error {
	pathnames := strings.Split(path, "/")

	for env in range []string{"dev", "staging", "production"} {
        if env == pathnames[2] {
            return env, nil
        }
    }

	return pathnames[2], fmt.Errorf("")
}
