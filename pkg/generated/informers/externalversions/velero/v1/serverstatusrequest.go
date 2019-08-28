/*
Copyright the Velero contributors.

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
	time "time"

	velerov1 "github.com/heptio/velero/pkg/apis/velero/v1"
	versioned "github.com/heptio/velero/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/heptio/velero/pkg/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/heptio/velero/pkg/generated/listers/velero/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ServerStatusRequestInformer provides access to a shared informer and lister for
// ServerStatusRequests.
type ServerStatusRequestInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ServerStatusRequestLister
}

type serverStatusRequestInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewServerStatusRequestInformer constructs a new informer for ServerStatusRequest type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewServerStatusRequestInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredServerStatusRequestInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredServerStatusRequestInformer constructs a new informer for ServerStatusRequest type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredServerStatusRequestInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.VeleroV1().ServerStatusRequests(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.VeleroV1().ServerStatusRequests(namespace).Watch(options)
			},
		},
		&velerov1.ServerStatusRequest{},
		resyncPeriod,
		indexers,
	)
}

func (f *serverStatusRequestInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredServerStatusRequestInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *serverStatusRequestInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&velerov1.ServerStatusRequest{}, f.defaultInformer)
}

func (f *serverStatusRequestInformer) Lister() v1.ServerStatusRequestLister {
	return v1.NewServerStatusRequestLister(f.Informer().GetIndexer())
}
