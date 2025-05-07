package version

import (
	"encoding/json"
	"fmt"
	"runtime"

	"github.com/gosuri/uitable"
)

var (
	gitVersion   = "v0.0.0-master+$Format:%H$"
	gitCommit    = "$Format:%H$"
	gitTreeState = ""
	buildDate    = "1970-01-01T00:00:00Z"
)

// Info contains versioning information for the FastGo project.
type Info struct {
	GitVersion   string `json:"gitVersion"`
	GitCommit    string `json:"gitCommit"`
	GitTreeState string `json:"gitTreeState"`
	BuildDate    string `json:"buildDate"`
	GoVersion    string `json:"goVersion"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
}

// String return info as a version string
func (info Info) String() string {
	return info.GitVersion
}

// ToJSON return the json string of version infomation
func (info Info) ToJSON() string {
	s, _ := json.Marshal(info)
	return string(s)
}

// Text encodes the version information into UTF-8 text and return the result.
func (info Info) Text() string {
	table := uitable.New()
	table.RightAlign(0)
	table.MaxColWidth = 80
	table.Separator = " "
	table.AddRow("GitVersion: ", info.GitVersion)
	table.AddRow("GitCommit: ", info.GitCommit)
	table.AddRow("GitTreeState: ", info.GitTreeState)
	table.AddRow("BuildDate: ", info.BuildDate)
	table.AddRow("GoVersion: ", info.GoVersion)
	table.AddRow("Compiler: ", info.Compiler)
	table.AddRow("Platform: ", info.Platform)
	return table.String()
}

func Get() Info {
	return Info{
		GitVersion:   gitVersion,
		GitCommit:    gitCommit,
		GitTreeState: gitTreeState,
		BuildDate:    buildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
