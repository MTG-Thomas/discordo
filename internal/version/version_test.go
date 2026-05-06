package version

import "testing"

func TestString(t *testing.T) {
	origVersion, origCommit, origDate, origDistribution := Version, Commit, Date, Distribution
	t.Cleanup(func() {
		Version, Commit, Date, Distribution = origVersion, origCommit, origDate, origDistribution
	})

	Version = "v1.2.3"
	Commit = "abc123"
	Date = "2026-05-06T14:00:00Z"
	Distribution = "mtg-fork"

	got := String()
	want := "discordo distribution=mtg-fork version=v1.2.3 commit=abc123 date=2026-05-06T14:00:00Z"
	if got != want {
		t.Fatalf("got = %q, want = %q", got, want)
	}
}
