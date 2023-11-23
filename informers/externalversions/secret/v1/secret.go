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
// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	secretv1 "github.com/edgelive/def/apis/secret/v1"
	versioned "github.com/edgelive/def/clientset/versioned"
	internalinterfaces "github.com/edgelive/def/informers/externalversions/internalinterfaces"
	v1 "github.com/edgelive/def/listers/secret/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// SecretInformer provides access to a shared informer and lister for
// Secrets.
type SecretInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.SecretLister
}

type secretInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSecretInformer constructs a new informer for Secret type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSecretInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSecretInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSecretInformer constructs a new informer for Secret type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSecretInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SecretV1().Secrets(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SecretV1().Secrets(namespace).Watch(context.TODO(), options)
			},
		},
		&secretv1.Secret{},
		resyncPeriod,
		indexers,
	)
}

func (f *secretInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSecretInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *secretInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&secretv1.Secret{}, f.defaultInformer)
}

func (f *secretInformer) Lister() v1.SecretLister {
	return v1.NewSecretLister(f.Informer().GetIndexer())
}
