/*
Copyright 2022.

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
	"context"
	time "time"

	versioned "github.com/robolaunch/kube-dev-suite/client/generated/clientset/versioned"
	internalinterfaces "github.com/robolaunch/kube-dev-suite/client/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/robolaunch/kube-dev-suite/client/generated/listers/roboscale.io/v1alpha1"
	roboscaleiov1alpha1 "github.com/robolaunch/kube-dev-suite/pkg/api/roboscale.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ROSBridgeInformer provides access to a shared informer and lister for
// ROSBridges.
type ROSBridgeInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ROSBridgeLister
}

type rOSBridgeInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewROSBridgeInformer constructs a new informer for ROSBridge type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewROSBridgeInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredROSBridgeInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredROSBridgeInformer constructs a new informer for ROSBridge type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredROSBridgeInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RoboscaleV1alpha1().ROSBridges(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RoboscaleV1alpha1().ROSBridges(namespace).Watch(context.TODO(), options)
			},
		},
		&roboscaleiov1alpha1.ROSBridge{},
		resyncPeriod,
		indexers,
	)
}

func (f *rOSBridgeInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredROSBridgeInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *rOSBridgeInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&roboscaleiov1alpha1.ROSBridge{}, f.defaultInformer)
}

func (f *rOSBridgeInformer) Lister() v1alpha1.ROSBridgeLister {
	return v1alpha1.NewROSBridgeLister(f.Informer().GetIndexer())
}
