package service

import (
	"context"
	"github.com/apirator/apirator-backend/internal/errors"
	apiratorv1alpha1 "github.com/apirator/apirator-backend/pkg/apis/apirator.io/v1alpha1"
	"github.com/apirator/apirator-backend/pkg/generated/clientset/versioned/typed/apirator.io/v1alpha1"
	ms "github.com/apirator/apirator-backend/pkg/mock"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type MockService struct {
	repository v1alpha1.APIMockInterface
}

const (
	Ingress   = "ingress"
	Namespace = "namespace"
	CreatedOnTag = "CreatedOn"
	CreatedOnValue = "apirator-backend"
)

func (ms *MockService) Add(namespace string,mock *ms.Mock) (*ms.Mock, error) {
	api := toResource(namespace,mock)
	co := v1.CreateOptions{}
	k8sMock, err := ms.repository.Create(context.TODO(), api, co)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	return toDomain(*k8sMock), nil
}

func (ms *MockService) List(namespace string) ([]*ms.Mock, error) {
	lo := v1.ListOptions{}
	k8sMocks, err := ms.repository.List(context.TODO(), lo)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	return toDomainList(k8sMocks), nil
}

func NewMockService(repo v1alpha1.APIMockInterface) *MockService {
	return &MockService{repository: repo}
}

// =====================================================================================================================
// Internal functions
// =====================================================================================================================

func toResource(namespace string,mock *ms.Mock) *apiratorv1alpha1.APIMock {
	selector := make(map[string]string)
	selector[Ingress] = mock.Selector.Ingress
	selector[Namespace] = mock.Selector.Namespace
	sp := apiratorv1alpha1.APIMockSpec{
		Definition: mock.Definition,
		ServiceDefinition: apiratorv1alpha1.ServiceDefinition{
			Port:        mock.ServiceDefinition.Port,
			ServiceType: mock.ServiceDefinition.ServiceType,
		},
		Watch:    mock.Watch,
		Selector: selector,
		Host:     mock.Host,
	}
	api := &apiratorv1alpha1.APIMock{
		TypeMeta:   v1.TypeMeta{},
		ObjectMeta: v1.ObjectMeta{
			Namespace: namespace,
			Name: mock.Name,
			Annotations: map[string]string{
				CreatedOnTag : CreatedOnValue,
			},
		},
		Spec:       sp,
	}
	return api
}

func toDomainList(list *apiratorv1alpha1.APIMockList) []*ms.Mock {
	result := make([]*ms.Mock, 10)
	for _, api := range list.Items {
		result = append(result, toDomain(api))
	}
	return result
}

func toDomain(api apiratorv1alpha1.APIMock) *ms.Mock {
	return &ms.Mock{
		ServiceDefinition: ms.ServiceDefinition{
			Port:        api.Spec.ServiceDefinition.Port,
			ServiceType: api.Spec.ServiceDefinition.ServiceType,
		},
		Watch: api.Spec.Watch,
		Selector: ms.IngressSelector{
			Ingress:   api.Spec.Selector[Ingress],
			Namespace: api.Spec.Selector[Namespace],
		},
		Host:       api.Spec.Host,
		Definition: api.Spec.Definition,
	}

}
