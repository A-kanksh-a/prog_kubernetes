// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "k8s.io/cnat-client-go/cnat/pkg/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Ats returns a AtInformer.
	Ats() AtInformer
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

// Ats returns a AtInformer.
func (v *version) Ats() AtInformer {
	return &atInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
