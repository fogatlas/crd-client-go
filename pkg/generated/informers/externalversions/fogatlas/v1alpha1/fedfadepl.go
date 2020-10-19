/*
Copyright FBK.

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

package v1alpha1

import (
	time "time"

	fogatlasv1alpha1 "github.com/fogatlas/crd-client-go/pkg/apis/fogatlas/v1alpha1"
	versioned "github.com/fogatlas/crd-client-go/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/fogatlas/crd-client-go/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/fogatlas/crd-client-go/pkg/generated/listers/fogatlas/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// FedFADeplInformer provides access to a shared informer and lister for
// FedFADepls.
type FedFADeplInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.FedFADeplLister
}

type fedFADeplInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewFedFADeplInformer constructs a new informer for FedFADepl type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFedFADeplInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFedFADeplInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredFedFADeplInformer constructs a new informer for FedFADepl type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFedFADeplInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FogatlasV1alpha1().FedFADepls(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FogatlasV1alpha1().FedFADepls(namespace).Watch(options)
			},
		},
		&fogatlasv1alpha1.FedFADepl{},
		resyncPeriod,
		indexers,
	)
}

func (f *fedFADeplInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFedFADeplInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *fedFADeplInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&fogatlasv1alpha1.FedFADepl{}, f.defaultInformer)
}

func (f *fedFADeplInformer) Lister() v1alpha1.FedFADeplLister {
	return v1alpha1.NewFedFADeplLister(f.Informer().GetIndexer())
}
