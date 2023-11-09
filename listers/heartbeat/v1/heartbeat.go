/*
Copyright The Kubernetes Authors.

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

package v1

import (
	v1 "github.com/edgelive/pkg/apis/heartbeat/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// HeartbeatLister helps list Heartbeats.
// All objects returned here must be treated as read-only.
type HeartbeatLister interface {
	// List lists all Heartbeats in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Heartbeat, err error)
	// Heartbeats returns an object that can list and get Heartbeats.
	Heartbeats(namespace string) HeartbeatNamespaceLister
	HeartbeatListerExpansion
}

// heartbeatLister implements the HeartbeatLister interface.
type heartbeatLister struct {
	indexer cache.Indexer
}

// NewHeartbeatLister returns a new HeartbeatLister.
func NewHeartbeatLister(indexer cache.Indexer) HeartbeatLister {
	return &heartbeatLister{indexer: indexer}
}

// List lists all Heartbeats in the indexer.
func (s *heartbeatLister) List(selector labels.Selector) (ret []*v1.Heartbeat, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Heartbeat))
	})
	return ret, err
}

// Heartbeats returns an object that can list and get Heartbeats.
func (s *heartbeatLister) Heartbeats(namespace string) HeartbeatNamespaceLister {
	return heartbeatNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// HeartbeatNamespaceLister helps list and get Heartbeats.
// All objects returned here must be treated as read-only.
type HeartbeatNamespaceLister interface {
	// List lists all Heartbeats in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Heartbeat, err error)
	// Get retrieves the Heartbeat from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Heartbeat, error)
	HeartbeatNamespaceListerExpansion
}

// heartbeatNamespaceLister implements the HeartbeatNamespaceLister
// interface.
type heartbeatNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Heartbeats in the indexer for a given namespace.
func (s heartbeatNamespaceLister) List(selector labels.Selector) (ret []*v1.Heartbeat, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Heartbeat))
	})
	return ret, err
}

// Get retrieves the Heartbeat from the indexer for a given namespace and name.
func (s heartbeatNamespaceLister) Get(name string) (*v1.Heartbeat, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("heartbeat"), name)
	}
	return obj.(*v1.Heartbeat), nil
}