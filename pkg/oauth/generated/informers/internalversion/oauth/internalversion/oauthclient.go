// This file was automatically generated by informer-gen

package internalversion

import (
	oauth "github.com/openshift/origin/pkg/oauth/apis/oauth"
	internalinterfaces "github.com/openshift/origin/pkg/oauth/generated/informers/internalversion/internalinterfaces"
	internalclientset "github.com/openshift/origin/pkg/oauth/generated/internalclientset"
	internalversion "github.com/openshift/origin/pkg/oauth/generated/listers/oauth/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// OAuthClientInformer provides access to a shared informer and lister for
// OAuthClients.
type OAuthClientInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.OAuthClientLister
}

type oAuthClientInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

// NewOAuthClientInformer constructs a new informer for OAuthClient type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewOAuthClientInformer(client internalclientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.Oauth().OAuthClients().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.Oauth().OAuthClients().Watch(options)
			},
		},
		&oauth.OAuthClient{},
		resyncPeriod,
		indexers,
	)
}

func defaultOAuthClientInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewOAuthClientInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
}

func (f *oAuthClientInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&oauth.OAuthClient{}, defaultOAuthClientInformer)
}

func (f *oAuthClientInformer) Lister() internalversion.OAuthClientLister {
	return internalversion.NewOAuthClientLister(f.Informer().GetIndexer())
}
