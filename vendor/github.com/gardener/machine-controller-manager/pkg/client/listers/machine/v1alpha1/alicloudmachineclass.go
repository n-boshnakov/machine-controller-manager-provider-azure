/*
Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

// AlicloudMachineClassLister helps list AlicloudMachineClasses.
type AlicloudMachineClassLister interface {
	// List lists all AlicloudMachineClasses in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AlicloudMachineClass, err error)
	// AlicloudMachineClasses returns an object that can list and get AlicloudMachineClasses.
	AlicloudMachineClasses(namespace string) AlicloudMachineClassNamespaceLister
	AlicloudMachineClassListerExpansion
}

// alicloudMachineClassLister implements the AlicloudMachineClassLister interface.
type alicloudMachineClassLister struct {
	indexer cache.Indexer
}

// NewAlicloudMachineClassLister returns a new AlicloudMachineClassLister.
func NewAlicloudMachineClassLister(indexer cache.Indexer) AlicloudMachineClassLister {
	return &alicloudMachineClassLister{indexer: indexer}
}

// List lists all AlicloudMachineClasses in the indexer.
func (s *alicloudMachineClassLister) List(selector labels.Selector) (ret []*v1alpha1.AlicloudMachineClass, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AlicloudMachineClass))
	})
	return ret, err
}

// AlicloudMachineClasses returns an object that can list and get AlicloudMachineClasses.
func (s *alicloudMachineClassLister) AlicloudMachineClasses(namespace string) AlicloudMachineClassNamespaceLister {
	return alicloudMachineClassNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AlicloudMachineClassNamespaceLister helps list and get AlicloudMachineClasses.
type AlicloudMachineClassNamespaceLister interface {
	// List lists all AlicloudMachineClasses in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AlicloudMachineClass, err error)
	// Get retrieves the AlicloudMachineClass from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AlicloudMachineClass, error)
	AlicloudMachineClassNamespaceListerExpansion
}

// alicloudMachineClassNamespaceLister implements the AlicloudMachineClassNamespaceLister
// interface.
type alicloudMachineClassNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AlicloudMachineClasses in the indexer for a given namespace.
func (s alicloudMachineClassNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AlicloudMachineClass, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AlicloudMachineClass))
	})
	return ret, err
}

// Get retrieves the AlicloudMachineClass from the indexer for a given namespace and name.
func (s alicloudMachineClassNamespaceLister) Get(name string) (*v1alpha1.AlicloudMachineClass, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("alicloudmachineclass"), name)
	}
	return obj.(*v1alpha1.AlicloudMachineClass), nil
}