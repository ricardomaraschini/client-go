// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	operatorv1 "github.com/openshift/api/operator/v1"
	versioned "github.com/openshift/client-go/operator/clientset/versioned"
	internalinterfaces "github.com/openshift/client-go/operator/informers/externalversions/internalinterfaces"
	v1 "github.com/openshift/client-go/operator/listers/operator/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// DNSInformer provides access to a shared informer and lister for
// DNSs.
type DNSInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.DNSLister
}

type dNSInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewDNSInformer constructs a new informer for DNS type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDNSInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDNSInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredDNSInformer constructs a new informer for DNS type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDNSInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperatorV1().DNSs().List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperatorV1().DNSs().Watch(options)
			},
		},
		&operatorv1.DNS{},
		resyncPeriod,
		indexers,
	)
}

func (f *dNSInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDNSInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *dNSInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&operatorv1.DNS{}, f.defaultInformer)
}

func (f *dNSInformer) Lister() v1.DNSLister {
	return v1.NewDNSLister(f.Informer().GetIndexer())
}
