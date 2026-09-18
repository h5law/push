package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	remotes, err := git("remote")
	if err != nil {
		fatal(err)
	}

	seen := make(map[string]struct{})

	for _, remote := range lines(remotes) {
		urls, err := git("remote", "get-url", "--push", "--all", remote)
		if err != nil {
			fatal(err)
		}

		for _, url := range lines(urls) {
			if _, ok := seen[url]; ok {
				continue
			}
			seen[url] = struct{}{}

			fmt.Printf("→ %s\n", url)

			cmd := exec.Command("git", "push", url)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Stdin = os.Stdin

			if err := cmd.Run(); err != nil {
				fmt.Fprintf(os.Stderr, "push failed: %s: %v\n", url, err)
				os.Exit(1)
			}
		}
	}
}

func git(args ...string) (string, error) {
	out, err := exec.Command("git", args...).Output()
	return strings.TrimSpace(string(out)), err
}

func lines(s string) []string {
	var result []string
	for line := range strings.SplitSeq(s, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		result = append(result, line)
	}
	return result
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, "push:", err)
	os.Exit(1)
}
