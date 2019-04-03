// +build !ignore_autogenerated

/*
Copyright 2019 kubeflow.org.

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
// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanarySpec) DeepCopyInto(out *CanarySpec) {
	*out = *in
	in.DefaultSpec.DeepCopyInto(&out.DefaultSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanarySpec.
func (in *CanarySpec) DeepCopy() *CanarySpec {
	if in == nil {
		return nil
	}
	out := new(CanarySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomSpec) DeepCopyInto(out *CustomSpec) {
	*out = *in
	in.Container.DeepCopyInto(&out.Container)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomSpec.
func (in *CustomSpec) DeepCopy() *CustomSpec {
	if in == nil {
		return nil
	}
	out := new(CustomSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultSpec) DeepCopyInto(out *DefaultSpec) {
	*out = *in
	if in.Custom != nil {
		in, out := &in.Custom, &out.Custom
		*out = new(CustomSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Tensorflow != nil {
		in, out := &in.Tensorflow, &out.Tensorflow
		*out = new(TensorflowSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.XGBoost != nil {
		in, out := &in.XGBoost, &out.XGBoost
		*out = new(XGBoostSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ScikitLearn != nil {
		in, out := &in.ScikitLearn, &out.ScikitLearn
		*out = new(ScikitLearnSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultSpec.
func (in *DefaultSpec) DeepCopy() *DefaultSpec {
	if in == nil {
		return nil
	}
	out := new(DefaultSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KFService) DeepCopyInto(out *KFService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KFService.
func (in *KFService) DeepCopy() *KFService {
	if in == nil {
		return nil
	}
	out := new(KFService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KFService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KFServiceList) DeepCopyInto(out *KFServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KFService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KFServiceList.
func (in *KFServiceList) DeepCopy() *KFServiceList {
	if in == nil {
		return nil
	}
	out := new(KFServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KFServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KFServiceSpec) DeepCopyInto(out *KFServiceSpec) {
	*out = *in
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new(DefaultSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KFServiceSpec.
func (in *KFServiceSpec) DeepCopy() *KFServiceSpec {
	if in == nil {
		return nil
	}
	out := new(KFServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KFServiceStatus) DeepCopyInto(out *KFServiceStatus) {
	*out = *in
	in.Conditions.DeepCopyInto(&out.Conditions)
	out.URI = in.URI
	out.Default = in.Default
	out.Canary = in.Canary
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KFServiceStatus.
func (in *KFServiceStatus) DeepCopy() *KFServiceStatus {
	if in == nil {
		return nil
	}
	out := new(KFServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScikitLearnSpec) DeepCopyInto(out *ScikitLearnSpec) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScikitLearnSpec.
func (in *ScikitLearnSpec) DeepCopy() *ScikitLearnSpec {
	if in == nil {
		return nil
	}
	out := new(ScikitLearnSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusConditionsSpec) DeepCopyInto(out *StatusConditionsSpec) {
	*out = *in
	in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusConditionsSpec.
func (in *StatusConditionsSpec) DeepCopy() *StatusConditionsSpec {
	if in == nil {
		return nil
	}
	out := new(StatusConditionsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusConfigurationSpec) DeepCopyInto(out *StatusConfigurationSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusConfigurationSpec.
func (in *StatusConfigurationSpec) DeepCopy() *StatusConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(StatusConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TensorflowSpec) DeepCopyInto(out *TensorflowSpec) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TensorflowSpec.
func (in *TensorflowSpec) DeepCopy() *TensorflowSpec {
	if in == nil {
		return nil
	}
	out := new(TensorflowSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *URISpec) DeepCopyInto(out *URISpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new URISpec.
func (in *URISpec) DeepCopy() *URISpec {
	if in == nil {
		return nil
	}
	out := new(URISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XGBoostSpec) DeepCopyInto(out *XGBoostSpec) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XGBoostSpec.
func (in *XGBoostSpec) DeepCopy() *XGBoostSpec {
	if in == nil {
		return nil
	}
	out := new(XGBoostSpec)
	in.DeepCopyInto(out)
	return out
}