/*
Copyright 2020 The KubeSphere Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha2 "kubesphere.io/kubesphere/pkg/apis/tenant/v1alpha2"
)

// FakeWorkspaceTemplates implements WorkspaceTemplateInterface
type FakeWorkspaceTemplates struct {
	Fake *FakeTenantV1alpha2
}

var workspacetemplatesResource = schema.GroupVersionResource{Group: "tenant.kubesphere.io", Version: "v1alpha2", Resource: "workspacetemplates"}

var workspacetemplatesKind = schema.GroupVersionKind{Group: "tenant.kubesphere.io", Version: "v1alpha2", Kind: "WorkspaceTemplate"}

// Get takes name of the workspaceTemplate, and returns the corresponding workspaceTemplate object, and an error if there is any.
func (c *FakeWorkspaceTemplates) Get(name string, options v1.GetOptions) (result *v1alpha2.WorkspaceTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(workspacetemplatesResource, name), &v1alpha2.WorkspaceTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.WorkspaceTemplate), err
}

// List takes label and field selectors, and returns the list of WorkspaceTemplates that match those selectors.
func (c *FakeWorkspaceTemplates) List(opts v1.ListOptions) (result *v1alpha2.WorkspaceTemplateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(workspacetemplatesResource, workspacetemplatesKind, opts), &v1alpha2.WorkspaceTemplateList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.WorkspaceTemplateList{ListMeta: obj.(*v1alpha2.WorkspaceTemplateList).ListMeta}
	for _, item := range obj.(*v1alpha2.WorkspaceTemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested workspaceTemplates.
func (c *FakeWorkspaceTemplates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(workspacetemplatesResource, opts))
}

// Create takes the representation of a workspaceTemplate and creates it.  Returns the server's representation of the workspaceTemplate, and an error, if there is any.
func (c *FakeWorkspaceTemplates) Create(workspaceTemplate *v1alpha2.WorkspaceTemplate) (result *v1alpha2.WorkspaceTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(workspacetemplatesResource, workspaceTemplate), &v1alpha2.WorkspaceTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.WorkspaceTemplate), err
}

// Update takes the representation of a workspaceTemplate and updates it. Returns the server's representation of the workspaceTemplate, and an error, if there is any.
func (c *FakeWorkspaceTemplates) Update(workspaceTemplate *v1alpha2.WorkspaceTemplate) (result *v1alpha2.WorkspaceTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(workspacetemplatesResource, workspaceTemplate), &v1alpha2.WorkspaceTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.WorkspaceTemplate), err
}

// Delete takes name of the workspaceTemplate and deletes it. Returns an error if one occurs.
func (c *FakeWorkspaceTemplates) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(workspacetemplatesResource, name), &v1alpha2.WorkspaceTemplate{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeWorkspaceTemplates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(workspacetemplatesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha2.WorkspaceTemplateList{})
	return err
}

// Patch applies the patch and returns the patched workspaceTemplate.
func (c *FakeWorkspaceTemplates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.WorkspaceTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(workspacetemplatesResource, name, pt, data, subresources...), &v1alpha2.WorkspaceTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.WorkspaceTemplate), err
}