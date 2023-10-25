//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Apply) DeepCopyInto(out *Apply) {
	*out = *in
	out.FileRef = in.FileRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Apply.
func (in *Apply) DeepCopy() *Apply {
	if in == nil {
		return nil
	}
	out := new(Apply)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Assert) DeepCopyInto(out *Assert) {
	*out = *in
	out.FileRef = in.FileRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Assert.
func (in *Assert) DeepCopy() *Assert {
	if in == nil {
		return nil
	}
	out := new(Assert)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Configuration) DeepCopyInto(out *Configuration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Configuration.
func (in *Configuration) DeepCopy() *Configuration {
	if in == nil {
		return nil
	}
	out := new(Configuration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Configuration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationSpec) DeepCopyInto(out *ConfigurationSpec) {
	*out = *in
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(v1.Duration)
		**out = **in
	}
	if in.TestDirs != nil {
		in, out := &in.TestDirs, &out.TestDirs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RepeatCount != nil {
		in, out := &in.RepeatCount, &out.RepeatCount
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationSpec.
func (in *ConfigurationSpec) DeepCopy() *ConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(ConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Delete) DeepCopyInto(out *Delete) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Delete.
func (in *Delete) DeepCopy() *Delete {
	if in == nil {
		return nil
	}
	out := new(Delete)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Error) DeepCopyInto(out *Error) {
	*out = *in
	out.FileRef = in.FileRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Error.
func (in *Error) DeepCopy() *Error {
	if in == nil {
		return nil
	}
	out := new(Error)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileRef) DeepCopyInto(out *FileRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileRef.
func (in *FileRef) DeepCopy() *FileRef {
	if in == nil {
		return nil
	}
	out := new(FileRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Test) DeepCopyInto(out *Test) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Test.
func (in *Test) DeepCopy() *Test {
	if in == nil {
		return nil
	}
	out := new(Test)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Test) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestSpec) DeepCopyInto(out *TestSpec) {
	*out = *in
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(v1.Duration)
		**out = **in
	}
	if in.Steps != nil {
		in, out := &in.Steps, &out.Steps
		*out = make([]TestSpecStep, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SkipDelete != nil {
		in, out := &in.SkipDelete, &out.SkipDelete
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestSpec.
func (in *TestSpec) DeepCopy() *TestSpec {
	if in == nil {
		return nil
	}
	out := new(TestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestSpecStep) DeepCopyInto(out *TestSpecStep) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestSpecStep.
func (in *TestSpecStep) DeepCopy() *TestSpecStep {
	if in == nil {
		return nil
	}
	out := new(TestSpecStep)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestStep) DeepCopyInto(out *TestStep) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestStep.
func (in *TestStep) DeepCopy() *TestStep {
	if in == nil {
		return nil
	}
	out := new(TestStep)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TestStep) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestStepSpec) DeepCopyInto(out *TestStepSpec) {
	*out = *in
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(v1.Duration)
		**out = **in
	}
	if in.SkipDelete != nil {
		in, out := &in.SkipDelete, &out.SkipDelete
		*out = new(bool)
		**out = **in
	}
	if in.Assert != nil {
		in, out := &in.Assert, &out.Assert
		*out = make([]Assert, len(*in))
		copy(*out, *in)
	}
	if in.Apply != nil {
		in, out := &in.Apply, &out.Apply
		*out = make([]Apply, len(*in))
		copy(*out, *in)
	}
	if in.Error != nil {
		in, out := &in.Error, &out.Error
		*out = make([]Error, len(*in))
		copy(*out, *in)
	}
	if in.Delete != nil {
		in, out := &in.Delete, &out.Delete
		*out = make([]Delete, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestStepSpec.
func (in *TestStepSpec) DeepCopy() *TestStepSpec {
	if in == nil {
		return nil
	}
	out := new(TestStepSpec)
	in.DeepCopyInto(out)
	return out
}
