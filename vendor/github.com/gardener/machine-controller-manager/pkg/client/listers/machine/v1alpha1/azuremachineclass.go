/*
Copyright (c) SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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
	v1alpha1 "github.com/gardener/machine-controller-manager/pkg/apis/machine/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AzureMachineClassLister helps list AzureMachineClasses.
// All objects returned here must be treated as read-only.
type AzureMachineClassLister interface {
	// List lists all AzureMachineClasses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AzureMachineClass, err error)
	// AzureMachineClasses returns an object that can list and get AzureMachineClasses.
	AzureMachineClasses(namespace string) AzureMachineClassNamespaceLister
	AzureMachineClassListerExpansion
}

// azureMachineClassLister implements the AzureMachineClassLister interface.
type azureMachineClassLister struct {
	indexer cache.Indexer
}

// NewAzureMachineClassLister returns a new AzureMachineClassLister.
func NewAzureMachineClassLister(indexer cache.Indexer) AzureMachineClassLister {
	return &azureMachineClassLister{indexer: indexer}
}

// List lists all AzureMachineClasses in the indexer.
func (s *azureMachineClassLister) List(selector labels.Selector) (ret []*v1alpha1.AzureMachineClass, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AzureMachineClass))
	})
	return ret, err
}

// AzureMachineClasses returns an object that can list and get AzureMachineClasses.
func (s *azureMachineClassLister) AzureMachineClasses(namespace string) AzureMachineClassNamespaceLister {
	return azureMachineClassNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AzureMachineClassNamespaceLister helps list and get AzureMachineClasses.
// All objects returned here must be treated as read-only.
type AzureMachineClassNamespaceLister interface {
	// List lists all AzureMachineClasses in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AzureMachineClass, err error)
	// Get retrieves the AzureMachineClass from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.AzureMachineClass, error)
	AzureMachineClassNamespaceListerExpansion
}

// azureMachineClassNamespaceLister implements the AzureMachineClassNamespaceLister
// interface.
type azureMachineClassNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AzureMachineClasses in the indexer for a given namespace.
func (s azureMachineClassNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AzureMachineClass, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AzureMachineClass))
	})
	return ret, err
}

// Get retrieves the AzureMachineClass from the indexer for a given namespace and name.
func (s azureMachineClassNamespaceLister) Get(name string) (*v1alpha1.AzureMachineClass, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("azuremachineclass"), name)
	}
	return obj.(*v1alpha1.AzureMachineClass), nil
}
