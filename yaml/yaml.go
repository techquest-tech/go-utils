package yaml

import (
	"encoding/base64"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/techquest-tech/go-utils/str"
	"github.com/valyala/fasttemplate"
	"gopkg.in/yaml.v3"
)

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

func ExtendYamlPineline(w io.Writer, tag string) (int, error) {
	tags := strings.Split(tag, "|")
	prefix := strings.TrimSpace(tags[0])

	switch prefix {
	case "now":
		now := time.Now()
		strNow := now.Format(time.RFC3339)
		return w.Write([]byte(strNow))
	case "version":
		return w.Write([]byte(Version))
	}

	index := 1
	for {
		current := ""
		if index < len(tags) {
			current = strings.TrimSpace(tags[index])
		}
		switch current {
		case "now":
			now := time.Now()
			prefix = now.Format(prefix)
		case "env":
			envString := os.Getenv(prefix)
			prefix = envString
		case "upper":
			prefix = strings.ToUpper(prefix)
		case "lower":
			prefix = strings.ToLower(prefix)
		case "base64":
			decoded, err := base64.StdEncoding.DecodeString(prefix)
			if err != nil {
				logrus.Error("setting error, base64 decode failed.", err)
				return 0, err
			}
			prefix = string(decoded)
		case "aes":
			prefix = AesDecrypt(prefix, AesKey)
		default:
			return w.Write([]byte(prefix))
		}
		index = index + 1
	}

}

//ExtendYaml functions for fasttemplate
func ExtendYaml(w io.Writer, tag string) (int, error) {
	tags := strings.Split(tag, ".")
	prefix := strings.TrimSpace(tags[0])
	value := ""
	if len(tags) > 1 {
		value = strings.TrimSpace(tags[1])
	}

	switch prefix {
	case "env":
		envString := os.Getenv(value)
		return w.Write([]byte(envString))
	case "now":
		now := time.Now()
		strNow := now.Format(time.RFC3339)
		return w.Write([]byte(strNow))
	case "version":
		return w.Write([]byte(Version))
	case "base64":
		decoded, err := base64.StdEncoding.DecodeString(value)
		if err != nil {
			logrus.Error("setting error, base64 decode failed.", err)
			return 0, err
		}
		return w.Write([]byte(decoded))
	case "aes":
		decryptCode := AesDecrypt(value, AesKey)
		return w.Write([]byte(decryptCode))
	default:
		return w.Write([]byte(tag))
	}
}

//DecodeFile decode yaml file
func DecodeFile(file string) ([]byte, error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		logrus.WithField("file", file).Error("read file failed.", err)
		return nil, err
	}
	return DecodeBytes(content)
}

//DecodeBytes decode yaml
func DecodeBytes(content []byte) ([]byte, error) {
	template, err := fasttemplate.NewTemplate(string(content), StartTag, EndTag)
	if err != nil {
		logrus.Error("valid template file failed.", err)
		return nil, err
	}
	decoded := template.ExecuteFuncString(ExtendYamlPineline)
	return []byte(decoded), nil
}

// LoadYaml load yaml file, Unmarshal it to interface{}
func LoadYaml(file string, out interface{}) error {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		logrus.WithField("file", file).Error("read file failed.", err)
		return err
	}
	return LoadYamlBytes(content, out)
}

// LoadYamlBytes load yaml by bytes.
func LoadYamlBytes(content []byte, out interface{}) error {

	decoded, err := DecodeBytes(content)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal([]byte(decoded), out)
	if err != nil {
		logrus.Error("yaml unmarshal failed.", err)
		return err
	}

	return nil
}
