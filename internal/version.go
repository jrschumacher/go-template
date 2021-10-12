package internal

import "github.com/jrschumacher/go-template/internal/conf"

type VersionStat struct {
	Version     string `json:"version"`
	VersionLong string `json:"versionLong"`
	BuildTime   string `json:"buildTime"`
}

func GetVersion() VersionStat {
	return VersionStat{
		Version:     conf.Version,
		VersionLong: conf.VersionLong,
		BuildTime:   conf.BuildTime,
	}
}
