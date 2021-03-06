package client

import (
	"fmt"
	"os"
	"testing"
)

func CreateResourceSet(rs ResourceSet) error {
	c, err := NewClient(GetConfigFromEnv())

	if err != nil {
		return err
	}
	_, err = c.CreateResourceSet(rs)
	return err
}

func CreateNetwork() error {
	c, err := NewClient(GetConfigFromEnv())

	if err != nil {
		return err
	}

	_, err = c.CreateNetwork(Network{
		NameLabel: testNetworkName,
		PoolId:    accTestPool.Id,
	})
	return err
}

var integrationTestPrefix string = "xenorchestra-client-"
var accTestPool Pool
var testTemplateName string

func TestMain(m *testing.M) {
	FindPoolForTests(&accTestPool)
	CreateNetwork()
	CreateResourceSet(testResourceSet)

	var found bool
	testTemplateName, found = os.LookupEnv("XOA_TEMPLATE")
	if !found {
		fmt.Println("The XOA_TEMPLATE environment variable must be set for the tests")
		os.Exit(-1)
	}

	code := m.Run()

	RemoveResourceSetsWithNamePrefix(integrationTestPrefix)("")
	RemoveNetworksWithNamePrefix(integrationTestPrefix)("")

	os.Exit(code)
}
