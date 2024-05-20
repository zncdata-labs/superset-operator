//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2024 zncdata-labs.

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

import ()

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CPUResource) DeepCopyInto(out *CPUResource) {
	*out = *in
	out.Max = in.Max.DeepCopy()
	out.Min = in.Min.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CPUResource.
func (in *CPUResource) DeepCopy() *CPUResource {
	if in == nil {
		return nil
	}
	out := new(CPUResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterOperationSpec) DeepCopyInto(out *ClusterOperationSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterOperationSpec.
func (in *ClusterOperationSpec) DeepCopy() *ClusterOperationSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterOperationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerLoggingSpec) DeepCopyInto(out *ContainerLoggingSpec) {
	*out = *in
	if in.SparkHistory != nil {
		in, out := &in.SparkHistory, &out.SparkHistory
		*out = new(LoggingConfigSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerLoggingSpec.
func (in *ContainerLoggingSpec) DeepCopy() *ContainerLoggingSpec {
	if in == nil {
		return nil
	}
	out := new(ContainerLoggingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogLevelSpec) DeepCopyInto(out *LogLevelSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogLevelSpec.
func (in *LogLevelSpec) DeepCopy() *LogLevelSpec {
	if in == nil {
		return nil
	}
	out := new(LogLevelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingConfigSpec) DeepCopyInto(out *LoggingConfigSpec) {
	*out = *in
	if in.Loggers != nil {
		in, out := &in.Loggers, &out.Loggers
		*out = make(map[string]*LogLevelSpec, len(*in))
		for key, val := range *in {
			var outVal *LogLevelSpec
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(LogLevelSpec)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Console != nil {
		in, out := &in.Console, &out.Console
		*out = new(LogLevelSpec)
		**out = **in
	}
	if in.File != nil {
		in, out := &in.File, &out.File
		*out = new(LogLevelSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingConfigSpec.
func (in *LoggingConfigSpec) DeepCopy() *LoggingConfigSpec {
	if in == nil {
		return nil
	}
	out := new(LoggingConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemoryResource) DeepCopyInto(out *MemoryResource) {
	*out = *in
	out.Limit = in.Limit.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemoryResource.
func (in *MemoryResource) DeepCopy() *MemoryResource {
	if in == nil {
		return nil
	}
	out := new(MemoryResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodDisruptionBudgetSpec) DeepCopyInto(out *PodDisruptionBudgetSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodDisruptionBudgetSpec.
func (in *PodDisruptionBudgetSpec) DeepCopy() *PodDisruptionBudgetSpec {
	if in == nil {
		return nil
	}
	out := new(PodDisruptionBudgetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcesSpec) DeepCopyInto(out *ResourcesSpec) {
	*out = *in
	if in.CPU != nil {
		in, out := &in.CPU, &out.CPU
		*out = new(CPUResource)
		(*in).DeepCopyInto(*out)
	}
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		*out = new(MemoryResource)
		(*in).DeepCopyInto(*out)
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(StorageResource)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcesSpec.
func (in *ResourcesSpec) DeepCopy() *ResourcesSpec {
	if in == nil {
		return nil
	}
	out := new(ResourcesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageResource) DeepCopyInto(out *StorageResource) {
	*out = *in
	out.Capacity = in.Capacity.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageResource.
func (in *StorageResource) DeepCopy() *StorageResource {
	if in == nil {
		return nil
	}
	out := new(StorageResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageResourceSpec) DeepCopyInto(out *StorageResourceSpec) {
	*out = *in
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = new(StorageResource)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageResourceSpec.
func (in *StorageResourceSpec) DeepCopy() *StorageResourceSpec {
	if in == nil {
		return nil
	}
	out := new(StorageResourceSpec)
	in.DeepCopyInto(out)
	return out
}
