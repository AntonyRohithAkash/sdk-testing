package version

//go:generate go run gen.go -version=$4

func GetVersion() string {
	return Version
}
