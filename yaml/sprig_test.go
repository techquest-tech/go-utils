package yaml_test

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/techquest-tech/go-utils/yaml"
)

// type TestObje struct {
// 	Rabbitmq struct{
// 		Host: string,

// 	},
// }

func TestLoadYamlSprig(t *testing.T) {
	expected := "10.13.x.xx"

	os.Setenv("AMQP_HOST", expected)

	out := map[string]interface{}{}

	err := yaml.LoadYamlViaTemplate("sprig_test.yml", out)
	assert.Nil(t, err)

	// logrus.Infof("%v", out)

	obj, ok := out["rabbitmq"]
	if !ok {
		t.Error("rabbitmq config missed.")
		t.Fail()
	}

	if ok {
		if rabbitmq, ok := obj.(map[string]string); ok {
			assert.Equal(t, rabbitmq["host"], expected)
			assert.Equal(t, rabbitmq["password"], "guest")
		}
	}

	if obj, ok := out["ace"]; ok {
		if ace, ok := obj.(string); ok {
			assert.Equal(t, ace, "guest")
		}
	}

	log.Printf("time: %v", out["timeFormat"])
	log.Printf("now: %v", out["now"])

}
