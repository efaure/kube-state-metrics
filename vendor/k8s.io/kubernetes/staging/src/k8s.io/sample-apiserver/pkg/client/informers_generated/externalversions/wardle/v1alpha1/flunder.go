/*
Copyright 2017 The Kubernetes Authors.

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

// This file was automatically generated by informer-gen

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	wardle_v1alpha1 "k8s.io/sample-apiserver/pkg/apis/wardle/v1alpha1"
	clientset "k8s.io/sample-apiserver/pkg/client/clientset_generated/clientset"
	internalinterfaces "k8s.io/sample-apiserver/pkg/client/informers_generated/externalversions/internalinterfaces"
	v1alpha1 "k8s.io/sample-apiserver/pkg/client/listers_generated/wardle/v1alpha1"
	time "time"
)

// FlunderInformer provides access to a shared informer and lister for
// Flunders.
type FlunderInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.FlunderLister
}

type flunderInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

func newFlunderInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	sharedIndexInformer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.WardleV1alpha1().Flunders(v1.NamespaceAll).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.WardleV1alpha1().Flunders(v1.NamespaceAll).Watch(options)
			},
		},
		&wardle_v1alpha1.Flunder{},
		resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	return sharedIndexInformer
}

func (f *flunderInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&wardle_v1alpha1.Flunder{}, newFlunderInformer)
}

func (f *flunderInformer) Lister() v1alpha1.FlunderLister {
	return v1alpha1.NewFlunderLister(f.Informer().GetIndexer())
}
