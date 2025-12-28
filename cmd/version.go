package cmd

import (
	"fmt"
	"os"
	"runtime"
	"text/tabwriter"

	goversion "github.com/caarlos0/go-version"
)

var name = `
                              __                                       _ __
   _________ ___  ____ ______/ /_      _________  ____ ___  ____ ___  (_) /_
  / ___/ __ ` + "`" + `__ \/ __ ` + "`" + `/ ___/ __/_____/ ___/ __ \/ __ ` + "`" + `__ \/ __ ` + "`" + `__ \/ / __/
 (__  ) / / / / / /_/ / /  / /_/_____/ /__/ /_/ / / / / / / / / / / / / /_
/____/_/ /_/ /_/\__,_/_/   \__/      \___/\____/_/ /_/ /_/_/ /_/ /_/_/\__/
`

var (
	GitCommit  = "unknown"
	GitVersion = "unknown"
	BuildDate  = "unknown"
	RepoURL    = "https://github.com/kevinliao852/smart-commit"
)

func Version(_ string) error {
	fmt.Printf("%s%s\n\n", name, RepoURL)

	type KV struct {
		K string
		V string
	}
	vi := goversion.GetVersionInfo()
	kvs := []KV{
		{K: "GitVersion", V: vi.GitVersion},
		{K: "GitCommit", V: vi.GitCommit},
		{K: "BuildDate", V: vi.BuildDate},
		{K: "GoVersion", V: runtime.Version()},
		{K: "Compiler", V: runtime.Compiler},
		{K: "Platform", V: fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)},
	}

	tw := tabwriter.NewWriter(os.Stdout, 1, 8, 1, ' ', 0)
	for _, kv := range kvs {
		if _, err := fmt.Fprintf(tw, "%s:\t%s\n", kv.K, kv.V); err != nil {
			return fmt.Errorf("failed to print: %s", err)
		}
	}
	if err := tw.Flush(); err != nil {
		return fmt.Errorf("failed to flush: %s", err)
	}
	os.Exit(0)
	return nil
}
