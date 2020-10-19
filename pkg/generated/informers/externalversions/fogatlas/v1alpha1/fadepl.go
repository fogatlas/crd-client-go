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

// FADeplInformer provides access to a shared informer and lister for
// FADepls.
type FADeplInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.FADeplLister
}

type fADeplInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewFADeplInformer constructs a new informer for FADepl type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFADeplInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFADeplInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredFADeplInformer constructs a new informer for FADepl type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFADeplInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FogatlasV1alpha1().FADepls(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FogatlasV1alpha1().FADepls(namespace).Watch(options)
			},
		},
		&fogatlasv1alpha1.FADepl{},
		resyncPeriod,
		indexers,
	)
}

func (f *fADeplInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFADeplInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *fADeplInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&fogatlasv1alpha1.FADepl{}, f.defaultInformer)
}

func (f *fADeplInformer) Lister() v1alpha1.FADeplLister {
	return v1alpha1.NewFADeplLister(f.Informer().GetIndexer())
}
