package version

// Version components
const (
	Maj = "0"
	Min = "25"
	Fix = "1"
)

var (
	// Version is the current version of Tendermint
	// Must be a string because scripts like dist.sh read this file.
	Version = "0.25.1-rc0-iris"

	// GitCommit is the current HEAD set using ldflags.
	GitCommit string
)

// ABCIVersion is the version of the ABCI library
const ABCIVersion = "0.14.0"

func init() {
	if GitCommit != "" {
		Version += "-" + GitCommit
	}
}
