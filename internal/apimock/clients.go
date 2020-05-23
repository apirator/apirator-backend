package apimock

import (
	"fmt"
	"github.com/apirator/apirator-backend/internal/errors"
	infra "github.com/apirator/apirator-backend/internal/logger"
	"github.com/apirator/apirator-backend/pkg/generated/clientset/versioned"
	"github.com/apirator/apirator-backend/pkg/generated/clientset/versioned/typed/apirator.io/v1alpha1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

func NewApiratorClient(config *rest.Config) (*versioned.Clientset, error) {
	ic, err := versioned.NewForConfig(config)
	return ic, errors.Wrap(err)
}

func NewMockClient(clientset *versioned.Clientset) v1alpha1.ApiratorV1alpha1Interface {
	return clientset.ApiratorV1alpha1()
}

func NewRestConfig() (*rest.Config, error) {
	infra.Logger.Info(fmt.Sprintf("Running in DEVELOPMENT mode %s", os.Getenv("DEVELOPMENT")))
	if isRunModeLocal() {
		host := os.Getenv("KUBERNETES_SERVICE_HOST")
		configPath := os.Getenv("KUBERNETES_CONFIG_PATH")
		config, err := clientcmd.BuildConfigFromFlags(host, configPath)
		return config, errors.Wrap(err)
	}
	infra.Logger.Info("Running in PRODUCTION mode. Start to create k8s config ")
	config, err := rest.InClusterConfig()
	if err == nil {
		infra.Logger.Info("Config created successfully")
	}
	return config, errors.WrapWithMessage(err, "Error to configure k8s client")
}

func isRunModeLocal() bool {
	return os.Getenv("DEVELOPMENT") == "TRUE"
}
