package conf

var Version = "development"
var VersionLong = "development+"
var Sha1 string
var BuildTime string

func init() {
	VersionLong = Version + "+" + Sha1
}
