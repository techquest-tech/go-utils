package yaml

import (
	"bytes"
	"html/template"
	"os"

	"github.com/Masterminds/sprig/v3"
	"gopkg.in/yaml.v3"
)

type YamlTemplate struct {
	Version string
	Yaml    string
}

func Parse(file string) ([]byte, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return []byte{}, err
	}

	tmpl := template.New("yaml").Funcs(sprig.FuncMap()).Funcs(template.FuncMap{
		"aesdec": AESDecrypt,
		"aes":    AESEncrypt,
		"version": func() string {
			return Version
		},
	})

	if StartTag != "" {
		tmpl.Delims(StartTag, EndTag)
	}

	tmpl, err = tmpl.Parse(string(content))

	if err != nil {
		return nil, err
	}

	y := &YamlTemplate{
		Version: Version,
	}
	out := &bytes.Buffer{}
	err = tmpl.Execute(out, y)
	if err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

//LoadYamlViaTemplate Load yaml with template engine.
func LoadYamlViaTemplate(file string, out interface{}) error {

	result, err := Parse(file)

	if err != nil {
		return err
	}

	// logrus.Infof("%s", result)

	err = yaml.Unmarshal(result, out)
	if err != nil {
		return err
	}

	return nil
}
