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
	internalinterfaces "github.com/fogatlas/crd-client-go/pkg/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// DynamicNodes returns a DynamicNodeInformer.
	DynamicNodes() DynamicNodeInformer
	// ExternalEndpoints returns a ExternalEndpointInformer.
	ExternalEndpoints() ExternalEndpointInformer
	// FADepls returns a FADeplInformer.
	FADepls() FADeplInformer
	// FedFAApps returns a FedFAAppInformer.
	FedFAApps() FedFAAppInformer
	// Links returns a LinkInformer.
	Links() LinkInformer
	// Regions returns a RegionInformer.
	Regions() RegionInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// DynamicNodes returns a DynamicNodeInformer.
func (v *version) DynamicNodes() DynamicNodeInformer {
	return &dynamicNodeInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ExternalEndpoints returns a ExternalEndpointInformer.
func (v *version) ExternalEndpoints() ExternalEndpointInformer {
	return &externalEndpointInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// FADepls returns a FADeplInformer.
func (v *version) FADepls() FADeplInformer {
	return &fADeplInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// FedFAApps returns a FedFAAppInformer.
func (v *version) FedFAApps() FedFAAppInformer {
	return &fedFAAppInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Links returns a LinkInformer.
func (v *version) Links() LinkInformer {
	return &linkInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Regions returns a RegionInformer.
func (v *version) Regions() RegionInformer {
	return &regionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
