package yaml

import (
	"os"
	"testing"
)

// type TestObje struct {
// 	Rabbitmq struct{
// 		Host: string,

// 	},
// }

func TestLoadYaml(t *testing.T) {
	expected := "192.168.1.3"

	os.Setenv("AMQP_HOST", expected)

	out := map[string]interface{}{}

	err := LoadYaml("yaml_test.yml", out)
	if err != nil {
		t.Error("load yaml file failed.", err)
		t.Fail()
	}

	if obj, ok := out["rabbitmq"]; ok {
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

}
