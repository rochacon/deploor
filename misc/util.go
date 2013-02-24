package util

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func CleanUp(clone string) {
	if clone != "" {
		os.RemoveAll(clone)
	}
}

func Abort(message string) {
	log.Fatalf(message)
	os.Exit(1)
}

// Return the reference type and name
// Example:
// 	refs/head/master -> branch, master
// 	refs/tags/1.0 -> tags, 1.0
func ParseRef(reference string) (string, string) {
	ref := strings.Split(reference, "/")

	if len(ref) < 3 {
		log.Fatalf("Invalid reference")
	}

	// Otherwise, return the branch name
	return ref[1], ref[2]
}

// Based on a repository path, get the environment
// Example: /srv/git/staging/project.git
func GetEnvironmentFromPath(path string) (string, error) {
	pathnames := strings.Split(path, "/")
	depth := len(pathnames) - 2

	for _, env := range []string{"dev", "staging", "production"} {
		if env == pathnames[depth] {
			return env, nil
		}
	}

	return pathnames[depth], fmt.Errorf("No environment found")
}
