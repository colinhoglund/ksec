package models

import (
	"strings"
	"testing"

	testclient "k8s.io/client-go/kubernetes/fake"
)

var secretsClient *SecretsClient

func testErr(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestNewSecretsClient(t *testing.T) {
	secretsClient = NewSecretsClient("default", "testuser", testclient.NewSimpleClientset())
}

func TestCreate(t *testing.T) {
	_, err := secretsClient.Create("test-secret")
	testErr(err, t)
}

func TestList(t *testing.T) {
	secrets, err := secretsClient.List()
	testErr(err, t)

	if secrets.Items[0].Name != "test-secret" {
		t.Fatal(err.Error())
	}
}

func TestCreateWithData(t *testing.T) {
	dataArgs := []string{
		"key=value",
		"secret-key=secret-value",
		"ENV_VAR=~!@#$%^&*()_+",
		"DB_URL=mongodb://host1.example.com:27017,host2.example.com:27017/prod?replicaSet=prod",
	}
	data := make(map[string][]byte)

	for _, item := range dataArgs {
		split := strings.SplitN(item, "=", 2)
		if len(split) != 2 {
			t.Errorf("Data is not formatted correctly: %s", item)
		}
		data[split[0]] = []byte(split[1])
	}

	_, err := secretsClient.CreateWithData("test-secret-with-data", data)
	testErr(err, t)
}