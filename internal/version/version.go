package version

import "fmt"

var (
	Version      = "dev"
	Commit       = "unknown"
	Date         = "unknown"
	Distribution = "mtg-fork"
)

func String() string {
	return fmt.Sprintf("discordo distribution=%s version=%s commit=%s date=%s", Distribution, Version, Commit, Date)
}
