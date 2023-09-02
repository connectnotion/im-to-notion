package version

import (
	"encoding/json"
	"fmt"
	"runtime"
)

// Version contains version information.
type Version struct {
	Major     string `json:"major"`
	Minor     string `json:"minor"`
	Patch     string `json:"patch"`
	GitCommit string `json:"git_commit"`
	BuildDate string `json:"build_date"`
	GoVersion string `json:"go_version"`
	Compiler  string `json:"compiler"`
	Platform  string `json:"platform"`
}

var (
	_major     string
	_minor     string
	_patch     string
	_buildDate string
	_gitCommit string

	// V contains the version information about the component.
	V Version
)

func init() {
	V = Version{
		Major:     _major,
		Minor:     _minor,
		Patch:     _patch,
		BuildDate: _buildDate,
		GoVersion: runtime.Version(),
		GitCommit: _gitCommit,
		Compiler:  runtime.Compiler,
		Platform:  fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}

// String shows the version info.
func (v Version) String() string {
	data, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(data)
}
