package yaml

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
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

	err := LoadYamlViaTemplate("sprig_test.yml", out)
	if err != nil {
		t.Error("load yaml file failed.", err)
		t.Fail()
	}

	logrus.Infof("%v", out)

	obj, ok := out["rabbitmq"]
	if !ok {
		t.Error("rabbitmq config missed.")
		t.Fail()
	}

	if ok {
		if rabbitmq, ok := obj.(map[string]string); ok {
			if rabbitmq["host"] != expected {
				t.Errorf("host is not matched env, %s vs %s", rabbitmq["host"], expected)
				t.Fail()
			}
			if rabbitmq["password"] != "guest" {
				t.Errorf("base64 decode failed. %s vs guest", rabbitmq["password"])
				t.Fail()
			}
		}
	}

	if obj, ok := out["ace"]; ok {
		if ace, ok := obj.(string); ok {
			if ace != "guest" {
				t.Errorf("aes decode failed, %s vs guest", ace)
				t.Fail()
			}
		}
	}

	logrus.Infof("time: %v", out["timeFormat"])
	logrus.Infof("now: %v", out["now"])

}
