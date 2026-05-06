package version

import "fmt"

var (
	Version = "dev"
	Commit  = "unknown"
	Date    = "unknown"
)

func String() string {
	return fmt.Sprintf("discordo version=%s commit=%s date=%s", Version, Commit, Date)
}
