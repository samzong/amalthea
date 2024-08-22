//go:build !ignore_autogenerated

/*
Copyright 2024.

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
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AmaltheaSession) DeepCopyInto(out *AmaltheaSession) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AmaltheaSession.
func (in *AmaltheaSession) DeepCopy() *AmaltheaSession {
	if in == nil {
		return nil
	}
	out := new(AmaltheaSession)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AmaltheaSession) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AmaltheaSessionCondition) DeepCopyInto(out *AmaltheaSessionCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AmaltheaSessionCondition.
func (in *AmaltheaSessionCondition) DeepCopy() *AmaltheaSessionCondition {
	if in == nil {
		return nil
	}
	out := new(AmaltheaSessionCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AmaltheaSessionList) DeepCopyInto(out *AmaltheaSessionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AmaltheaSession, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AmaltheaSessionList.
func (in *AmaltheaSessionList) DeepCopy() *AmaltheaSessionList {
	if in == nil {
		return nil
	}
	out := new(AmaltheaSessionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AmaltheaSessionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AmaltheaSessionSpec) DeepCopyInto(out *AmaltheaSessionSpec) {
	*out = *in
	in.Session.DeepCopyInto(&out.Session)
	if in.CodeRepositories != nil {
		in, out := &in.CodeRepositories, &out.CodeRepositories
		*out = make([]CodeRepository, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DataSources != nil {
		in, out := &in.DataSources, &out.DataSources
		*out = make([]DataSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Authentication != nil {
		in, out := &in.Authentication, &out.Authentication
		*out = new(Authentication)
		(*in).DeepCopyInto(*out)
	}
	out.Culling = in.Culling
	if in.ExtraContainers != nil {
		in, out := &in.ExtraContainers, &out.ExtraContainers
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExtraInitContainers != nil {
		in, out := &in.ExtraInitContainers, &out.ExtraInitContainers
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExtraVolumes != nil {
		in, out := &in.ExtraVolumes, &out.ExtraVolumes
		*out = make([]v1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = new(Ingress)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AmaltheaSessionSpec.
func (in *AmaltheaSessionSpec) DeepCopy() *AmaltheaSessionSpec {
	if in == nil {
		return nil
	}
	out := new(AmaltheaSessionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AmaltheaSessionStatus) DeepCopyInto(out *AmaltheaSessionStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]AmaltheaSessionCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.ContainerCounts = in.ContainerCounts
	out.InitContainerCounts = in.InitContainerCounts
	in.IdleSince.DeepCopyInto(&out.IdleSince)
	in.FailingSince.DeepCopyInto(&out.FailingSince)
	in.HibernatedSince.DeepCopyInto(&out.HibernatedSince)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AmaltheaSessionStatus.
func (in *AmaltheaSessionStatus) DeepCopy() *AmaltheaSessionStatus {
	if in == nil {
		return nil
	}
	out := new(AmaltheaSessionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Authentication) DeepCopyInto(out *Authentication) {
	*out = *in
	out.SecretRef = in.SecretRef
	if in.ExtraVolumeMounts != nil {
		in, out := &in.ExtraVolumeMounts, &out.ExtraVolumeMounts
		*out = make([]v1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Authentication.
func (in *Authentication) DeepCopy() *Authentication {
	if in == nil {
		return nil
	}
	out := new(Authentication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CodeRepository) DeepCopyInto(out *CodeRepository) {
	*out = *in
	if in.CloningConfigSecretRef != nil {
		in, out := &in.CloningConfigSecretRef, &out.CloningConfigSecretRef
		*out = new(SessionSecretRef)
		**out = **in
	}
	if in.ConfigSecretRef != nil {
		in, out := &in.ConfigSecretRef, &out.ConfigSecretRef
		*out = new(SessionSecretRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CodeRepository.
func (in *CodeRepository) DeepCopy() *CodeRepository {
	if in == nil {
		return nil
	}
	out := new(CodeRepository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerCounts) DeepCopyInto(out *ContainerCounts) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerCounts.
func (in *ContainerCounts) DeepCopy() *ContainerCounts {
	if in == nil {
		return nil
	}
	out := new(ContainerCounts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Culling) DeepCopyInto(out *Culling) {
	*out = *in
	out.MaxAge = in.MaxAge
	out.MaxIdleDuration = in.MaxIdleDuration
	out.MaxStartingDuration = in.MaxStartingDuration
	out.MaxFailedDuration = in.MaxFailedDuration
	out.MaxHibernatedDuration = in.MaxHibernatedDuration
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Culling.
func (in *Culling) DeepCopy() *Culling {
	if in == nil {
		return nil
	}
	out := new(Culling)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataSource) DeepCopyInto(out *DataSource) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(SessionSecretRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataSource.
func (in *DataSource) DeepCopy() *DataSource {
	if in == nil {
		return nil
	}
	out := new(DataSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ingress) DeepCopyInto(out *Ingress) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.IngressClassName != nil {
		in, out := &in.IngressClassName, &out.IngressClassName
		*out = new(string)
		**out = **in
	}
	if in.TLSSecretName != nil {
		in, out := &in.TLSSecretName, &out.TLSSecretName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ingress.
func (in *Ingress) DeepCopy() *Ingress {
	if in == nil {
		return nil
	}
	out := new(Ingress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Session) DeepCopyInto(out *Session) {
	*out = *in
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	in.Storage.DeepCopyInto(&out.Storage)
	if in.ShmSize != nil {
		in, out := &in.ShmSize, &out.ShmSize
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.ExtraVolumeMounts != nil {
		in, out := &in.ExtraVolumeMounts, &out.ExtraVolumeMounts
		*out = make([]v1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Session.
func (in *Session) DeepCopy() *Session {
	if in == nil {
		return nil
	}
	out := new(Session)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SessionSecretRef) DeepCopyInto(out *SessionSecretRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SessionSecretRef.
func (in *SessionSecretRef) DeepCopy() *SessionSecretRef {
	if in == nil {
		return nil
	}
	out := new(SessionSecretRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Storage) DeepCopyInto(out *Storage) {
	*out = *in
	if in.ClassName != nil {
		in, out := &in.ClassName, &out.ClassName
		*out = new(string)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		x := (*in).DeepCopy()
		*out = &x
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Storage.
func (in *Storage) DeepCopy() *Storage {
	if in == nil {
		return nil
	}
	out := new(Storage)
	in.DeepCopyInto(out)
	return out
}
