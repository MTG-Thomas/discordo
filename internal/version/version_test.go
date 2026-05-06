package version

import "testing"

func TestString(t *testing.T) {
	origVersion, origCommit, origDate := Version, Commit, Date
	t.Cleanup(func() {
		Version, Commit, Date = origVersion, origCommit, origDate
	})

	Version = "v1.2.3"
	Commit = "abc123"
	Date = "2026-05-06T14:00:00Z"

	got := String()
	want := "discordo version=v1.2.3 commit=abc123 date=2026-05-06T14:00:00Z"
	if got != want {
		t.Fatalf("got = %q, want = %q", got, want)
	}
}
