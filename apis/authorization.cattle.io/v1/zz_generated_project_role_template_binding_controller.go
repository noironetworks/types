package v1

import (
	"context"

	"github.com/rancher/norman/clientbase"
	"github.com/rancher/norman/controller"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

var (
	ProjectRoleTemplateBindingGroupVersionKind = schema.GroupVersionKind{
		Version: "v1",
		Group:   "authorization.cattle.io",
		Kind:    "ProjectRoleTemplateBinding",
	}
	ProjectRoleTemplateBindingResource = metav1.APIResource{
		Name:         "projectroletemplatebindings",
		SingularName: "projectroletemplatebinding",
		Namespaced:   false,
		Kind:         ProjectRoleTemplateBindingGroupVersionKind.Kind,
	}
)

type ProjectRoleTemplateBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectRoleTemplateBinding
}

type ProjectRoleTemplateBindingHandlerFunc func(key string, obj *ProjectRoleTemplateBinding) error

type ProjectRoleTemplateBindingController interface {
	Informer() cache.SharedIndexInformer
	AddHandler(handler ProjectRoleTemplateBindingHandlerFunc)
	Enqueue(namespace, name string)
	Start(ctx context.Context, threadiness int) error
}

type ProjectRoleTemplateBindingInterface interface {
	Create(*ProjectRoleTemplateBinding) (*ProjectRoleTemplateBinding, error)
	Get(name string, opts metav1.GetOptions) (*ProjectRoleTemplateBinding, error)
	Update(*ProjectRoleTemplateBinding) (*ProjectRoleTemplateBinding, error)
	Delete(name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*ProjectRoleTemplateBindingList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() ProjectRoleTemplateBindingController
}

type projectRoleTemplateBindingController struct {
	controller.GenericController
}

func (c *projectRoleTemplateBindingController) AddHandler(handler ProjectRoleTemplateBindingHandlerFunc) {
	c.GenericController.AddHandler(func(key string) error {
		obj, exists, err := c.Informer().GetStore().GetByKey(key)
		if err != nil {
			return err
		}
		if !exists {
			return handler(key, nil)
		}
		return handler(key, obj.(*ProjectRoleTemplateBinding))
	})
}

type projectRoleTemplateBindingFactory struct {
}

func (c projectRoleTemplateBindingFactory) Object() runtime.Object {
	return &ProjectRoleTemplateBinding{}
}

func (c projectRoleTemplateBindingFactory) List() runtime.Object {
	return &ProjectRoleTemplateBindingList{}
}

func (s *projectRoleTemplateBindingClient) Controller() ProjectRoleTemplateBindingController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.projectRoleTemplateBindingControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(ProjectRoleTemplateBindingGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &projectRoleTemplateBindingController{
		GenericController: genericController,
	}

	s.client.projectRoleTemplateBindingControllers[s.ns] = c

	return c
}

type projectRoleTemplateBindingClient struct {
	client       *Client
	ns           string
	objectClient *clientbase.ObjectClient
	controller   ProjectRoleTemplateBindingController
}

func (s *projectRoleTemplateBindingClient) Create(o *ProjectRoleTemplateBinding) (*ProjectRoleTemplateBinding, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*ProjectRoleTemplateBinding), err
}

func (s *projectRoleTemplateBindingClient) Get(name string, opts metav1.GetOptions) (*ProjectRoleTemplateBinding, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*ProjectRoleTemplateBinding), err
}

func (s *projectRoleTemplateBindingClient) Update(o *ProjectRoleTemplateBinding) (*ProjectRoleTemplateBinding, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*ProjectRoleTemplateBinding), err
}

func (s *projectRoleTemplateBindingClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *projectRoleTemplateBindingClient) List(opts metav1.ListOptions) (*ProjectRoleTemplateBindingList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*ProjectRoleTemplateBindingList), err
}

func (s *projectRoleTemplateBindingClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

func (s *projectRoleTemplateBindingClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}
