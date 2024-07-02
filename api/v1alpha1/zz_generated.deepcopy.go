//go:build !ignore_autogenerated

/*
Copyright 2023.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlarmSubscriptionServerConfig) DeepCopyInto(out *AlarmSubscriptionServerConfig) {
	*out = *in
	out.ServerConfig = in.ServerConfig
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlarmSubscriptionServerConfig.
func (in *AlarmSubscriptionServerConfig) DeepCopy() *AlarmSubscriptionServerConfig {
	if in == nil {
		return nil
	}
	out := new(AlarmSubscriptionServerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterRequest) DeepCopyInto(out *ClusterRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterRequest.
func (in *ClusterRequest) DeepCopy() *ClusterRequest {
	if in == nil {
		return nil
	}
	out := new(ClusterRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterRequestList) DeepCopyInto(out *ClusterRequestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterRequestList.
func (in *ClusterRequestList) DeepCopy() *ClusterRequestList {
	if in == nil {
		return nil
	}
	out := new(ClusterRequestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterRequestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterRequestSpec) DeepCopyInto(out *ClusterRequestSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterRequestSpec.
func (in *ClusterRequestSpec) DeepCopy() *ClusterRequestSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterRequestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterRequestStatus) DeepCopyInto(out *ClusterRequestStatus) {
	*out = *in
	out.ClusterTemplateInputValidation = in.ClusterTemplateInputValidation
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterRequestStatus.
func (in *ClusterRequestStatus) DeepCopy() *ClusterRequestStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterRequestStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterTemplate) DeepCopyInto(out *ClusterTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterTemplate.
func (in *ClusterTemplate) DeepCopy() *ClusterTemplate {
	if in == nil {
		return nil
	}
	out := new(ClusterTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterTemplateInputValidation) DeepCopyInto(out *ClusterTemplateInputValidation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterTemplateInputValidation.
func (in *ClusterTemplateInputValidation) DeepCopy() *ClusterTemplateInputValidation {
	if in == nil {
		return nil
	}
	out := new(ClusterTemplateInputValidation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterTemplateList) DeepCopyInto(out *ClusterTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterTemplateList.
func (in *ClusterTemplateList) DeepCopy() *ClusterTemplateList {
	if in == nil {
		return nil
	}
	out := new(ClusterTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterTemplateSpec) DeepCopyInto(out *ClusterTemplateSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterTemplateSpec.
func (in *ClusterTemplateSpec) DeepCopy() *ClusterTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterTemplateStatus) DeepCopyInto(out *ClusterTemplateStatus) {
	*out = *in
	out.ClusterTemplateValidation = in.ClusterTemplateValidation
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterTemplateStatus.
func (in *ClusterTemplateStatus) DeepCopy() *ClusterTemplateStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterTemplateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterTemplateValidation) DeepCopyInto(out *ClusterTemplateValidation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterTemplateValidation.
func (in *ClusterTemplateValidation) DeepCopy() *ClusterTemplateValidation {
	if in == nil {
		return nil
	}
	out := new(ClusterTemplateValidation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentManager) DeepCopyInto(out *DeploymentManager) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentManager.
func (in *DeploymentManager) DeepCopy() *DeploymentManager {
	if in == nil {
		return nil
	}
	out := new(DeploymentManager)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeploymentManager) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentManagerList) DeepCopyInto(out *DeploymentManagerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DeploymentManager, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentManagerList.
func (in *DeploymentManagerList) DeepCopy() *DeploymentManagerList {
	if in == nil {
		return nil
	}
	out := new(DeploymentManagerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeploymentManagerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentManagerOrder) DeepCopyInto(out *DeploymentManagerOrder) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentManagerOrder.
func (in *DeploymentManagerOrder) DeepCopy() *DeploymentManagerOrder {
	if in == nil {
		return nil
	}
	out := new(DeploymentManagerOrder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeploymentManagerOrder) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentManagerOrderList) DeepCopyInto(out *DeploymentManagerOrderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DeploymentManagerOrder, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentManagerOrderList.
func (in *DeploymentManagerOrderList) DeepCopy() *DeploymentManagerOrderList {
	if in == nil {
		return nil
	}
	out := new(DeploymentManagerOrderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeploymentManagerOrderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentManagerOrderSpec) DeepCopyInto(out *DeploymentManagerOrderSpec) {
	*out = *in
	if in.Extensions != nil {
		in, out := &in.Extensions, &out.Extensions
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentManagerOrderSpec.
func (in *DeploymentManagerOrderSpec) DeepCopy() *DeploymentManagerOrderSpec {
	if in == nil {
		return nil
	}
	out := new(DeploymentManagerOrderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentManagerOrderStatus) DeepCopyInto(out *DeploymentManagerOrderStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentManagerOrderStatus.
func (in *DeploymentManagerOrderStatus) DeepCopy() *DeploymentManagerOrderStatus {
	if in == nil {
		return nil
	}
	out := new(DeploymentManagerOrderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentManagerServerConfig) DeepCopyInto(out *DeploymentManagerServerConfig) {
	*out = *in
	out.ServerConfig = in.ServerConfig
	if in.Extensions != nil {
		in, out := &in.Extensions, &out.Extensions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentManagerServerConfig.
func (in *DeploymentManagerServerConfig) DeepCopy() *DeploymentManagerServerConfig {
	if in == nil {
		return nil
	}
	out := new(DeploymentManagerServerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentManagerTemplate) DeepCopyInto(out *DeploymentManagerTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.NodeProfiles != nil {
		in, out := &in.NodeProfiles, &out.NodeProfiles
		*out = make([]DeploymentManagerTemplateNodeProfile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeSets != nil {
		in, out := &in.NodeSets, &out.NodeSets
		*out = make([]DeploymentManagerTemplateNodeSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Extensions != nil {
		in, out := &in.Extensions, &out.Extensions
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentManagerTemplate.
func (in *DeploymentManagerTemplate) DeepCopy() *DeploymentManagerTemplate {
	if in == nil {
		return nil
	}
	out := new(DeploymentManagerTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeploymentManagerTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentManagerTemplateList) DeepCopyInto(out *DeploymentManagerTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DeploymentManagerTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentManagerTemplateList.
func (in *DeploymentManagerTemplateList) DeepCopy() *DeploymentManagerTemplateList {
	if in == nil {
		return nil
	}
	out := new(DeploymentManagerTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeploymentManagerTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentManagerTemplateNodeProfile) DeepCopyInto(out *DeploymentManagerTemplateNodeProfile) {
	*out = *in
	if in.Extensions != nil {
		in, out := &in.Extensions, &out.Extensions
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentManagerTemplateNodeProfile.
func (in *DeploymentManagerTemplateNodeProfile) DeepCopy() *DeploymentManagerTemplateNodeProfile {
	if in == nil {
		return nil
	}
	out := new(DeploymentManagerTemplateNodeProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentManagerTemplateNodeSet) DeepCopyInto(out *DeploymentManagerTemplateNodeSet) {
	*out = *in
	if in.Extensions != nil {
		in, out := &in.Extensions, &out.Extensions
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentManagerTemplateNodeSet.
func (in *DeploymentManagerTemplateNodeSet) DeepCopy() *DeploymentManagerTemplateNodeSet {
	if in == nil {
		return nil
	}
	out := new(DeploymentManagerTemplateNodeSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentsStatus) DeepCopyInto(out *DeploymentsStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentsStatus.
func (in *DeploymentsStatus) DeepCopy() *DeploymentsStatus {
	if in == nil {
		return nil
	}
	out := new(DeploymentsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetadataServerConfig) DeepCopyInto(out *MetadataServerConfig) {
	*out = *in
	out.ServerConfig = in.ServerConfig
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetadataServerConfig.
func (in *MetadataServerConfig) DeepCopy() *MetadataServerConfig {
	if in == nil {
		return nil
	}
	out := new(MetadataServerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ORANO2IMS) DeepCopyInto(out *ORANO2IMS) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ORANO2IMS.
func (in *ORANO2IMS) DeepCopy() *ORANO2IMS {
	if in == nil {
		return nil
	}
	out := new(ORANO2IMS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ORANO2IMS) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ORANO2IMSList) DeepCopyInto(out *ORANO2IMSList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ORANO2IMS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ORANO2IMSList.
func (in *ORANO2IMSList) DeepCopy() *ORANO2IMSList {
	if in == nil {
		return nil
	}
	out := new(ORANO2IMSList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ORANO2IMSList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ORANO2IMSSpec) DeepCopyInto(out *ORANO2IMSSpec) {
	*out = *in
	out.MetadataServerConfig = in.MetadataServerConfig
	in.DeploymentManagerServerConfig.DeepCopyInto(&out.DeploymentManagerServerConfig)
	in.ResourceServerConfig.DeepCopyInto(&out.ResourceServerConfig)
	out.AlarmSubscriptionServerConfig = in.AlarmSubscriptionServerConfig
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ORANO2IMSSpec.
func (in *ORANO2IMSSpec) DeepCopy() *ORANO2IMSSpec {
	if in == nil {
		return nil
	}
	out := new(ORANO2IMSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ORANO2IMSStatus) DeepCopyInto(out *ORANO2IMSStatus) {
	*out = *in
	in.DeploymentsStatus.DeepCopyInto(&out.DeploymentsStatus)
	in.UsedServerConfig.DeepCopyInto(&out.UsedServerConfig)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ORANO2IMSStatus.
func (in *ORANO2IMSStatus) DeepCopy() *ORANO2IMSStatus {
	if in == nil {
		return nil
	}
	out := new(ORANO2IMSStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceServerConfig) DeepCopyInto(out *ResourceServerConfig) {
	*out = *in
	out.ServerConfig = in.ServerConfig
	if in.Extensions != nil {
		in, out := &in.Extensions, &out.Extensions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceServerConfig.
func (in *ResourceServerConfig) DeepCopy() *ResourceServerConfig {
	if in == nil {
		return nil
	}
	out := new(ResourceServerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerConfig) DeepCopyInto(out *ServerConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerConfig.
func (in *ServerConfig) DeepCopy() *ServerConfig {
	if in == nil {
		return nil
	}
	out := new(ServerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UsedServerConfig) DeepCopyInto(out *UsedServerConfig) {
	*out = *in
	if in.MetadataServerUsedConfig != nil {
		in, out := &in.MetadataServerUsedConfig, &out.MetadataServerUsedConfig
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ResourceServerUsedConfig != nil {
		in, out := &in.ResourceServerUsedConfig, &out.ResourceServerUsedConfig
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DeploymentManagerServerUsedConfig != nil {
		in, out := &in.DeploymentManagerServerUsedConfig, &out.DeploymentManagerServerUsedConfig
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UsedServerConfig.
func (in *UsedServerConfig) DeepCopy() *UsedServerConfig {
	if in == nil {
		return nil
	}
	out := new(UsedServerConfig)
	in.DeepCopyInto(out)
	return out
}
