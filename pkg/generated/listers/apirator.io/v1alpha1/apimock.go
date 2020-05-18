/*
Copyright The apirator.io Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/apirator/apirator-backend/pkg/apis/apirator.io/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// APIMockLister helps list APIMocks.
type APIMockLister interface {
	// List lists all APIMocks in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.APIMock, err error)
	// APIMocks returns an object that can list and get APIMocks.
	APIMocks(namespace string) APIMockNamespaceLister
	APIMockListerExpansion
}

// aPIMockLister implements the APIMockLister interface.
type aPIMockLister struct {
	indexer cache.Indexer
}

// NewAPIMockLister returns a new APIMockLister.
func NewAPIMockLister(indexer cache.Indexer) APIMockLister {
	return &aPIMockLister{indexer: indexer}
}

// List lists all APIMocks in the indexer.
func (s *aPIMockLister) List(selector labels.Selector) (ret []*v1alpha1.APIMock, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.APIMock))
	})
	return ret, err
}

// APIMocks returns an object that can list and get APIMocks.
func (s *aPIMockLister) APIMocks(namespace string) APIMockNamespaceLister {
	return aPIMockNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// APIMockNamespaceLister helps list and get APIMocks.
type APIMockNamespaceLister interface {
	// List lists all APIMocks in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.APIMock, err error)
	// Get retrieves the APIMock from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.APIMock, error)
	APIMockNamespaceListerExpansion
}

// aPIMockNamespaceLister implements the APIMockNamespaceLister
// interface.
type aPIMockNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all APIMocks in the indexer for a given namespace.
func (s aPIMockNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.APIMock, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.APIMock))
	})
	return ret, err
}

// Get retrieves the APIMock from the indexer for a given namespace and name.
func (s aPIMockNamespaceLister) Get(name string) (*v1alpha1.APIMock, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("apimock"), name)
	}
	return obj.(*v1alpha1.APIMock), nil
}