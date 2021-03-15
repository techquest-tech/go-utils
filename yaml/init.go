package yaml

import "github.com/techquest-tech/go-utils/str"

//StartTag default start tag
var StartTag = "${"

//EndTag End tag
var EndTag = "}"

//Version version
var Version = "development"

func init() {
	str.ReplaceByEnv("APP_AES_KEY", &AesKey)
	str.ReplaceByEnv("YAML_START_TAG", &StartTag)
	str.ReplaceByEnv("YAML_END_TAG", &EndTag)
	str.ReplaceByEnv("APP_VERSION", &Version)
}
