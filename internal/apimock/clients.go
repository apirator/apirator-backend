package apimock

import (
	"github.com/apirator/apirator-backend/internal/errors"
	"github.com/apirator/apirator-backend/pkg/generated/clientset/versioned/typed/apirator.io/v1alpha1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

func NewApiratorClient(config *rest.Config) (*v1alpha1.ApiratorV1alpha1Client, error) {
	ic, err := v1alpha1.NewForConfig(config)
	return ic, errors.Wrap(err)
}

func NewRestConfig() (*rest.Config, error) {
	if isRunModeLocal() {
		host := os.Getenv("KUBERNETES_SERVICE_HOST")
		configPath := os.Getenv("KUBERNETES_CONFIG_PATH")
		config, err := clientcmd.BuildConfigFromFlags(host, configPath)
		return config, errors.Wrap(err)
	}
	config, err := rest.InClusterConfig()
	return config, errors.Wrap(err)
}

func isRunModeLocal() bool {
	return os.Getenv("DEVELOPMENT") == "TRUE"
}
