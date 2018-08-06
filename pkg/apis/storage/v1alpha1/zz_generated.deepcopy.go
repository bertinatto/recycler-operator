// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnterpriseRecycler) DeepCopyInto(out *EnterpriseRecycler) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnterpriseRecycler.
func (in *EnterpriseRecycler) DeepCopy() *EnterpriseRecycler {
	if in == nil {
		return nil
	}
	out := new(EnterpriseRecycler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnterpriseRecycler) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnterpriseRecyclerList) DeepCopyInto(out *EnterpriseRecyclerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EnterpriseRecycler, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnterpriseRecyclerList.
func (in *EnterpriseRecyclerList) DeepCopy() *EnterpriseRecyclerList {
	if in == nil {
		return nil
	}
	out := new(EnterpriseRecyclerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnterpriseRecyclerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnterpriseRecyclerSpec) DeepCopyInto(out *EnterpriseRecyclerSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnterpriseRecyclerSpec.
func (in *EnterpriseRecyclerSpec) DeepCopy() *EnterpriseRecyclerSpec {
	if in == nil {
		return nil
	}
	out := new(EnterpriseRecyclerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnterpriseRecyclerStatus) DeepCopyInto(out *EnterpriseRecyclerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnterpriseRecyclerStatus.
func (in *EnterpriseRecyclerStatus) DeepCopy() *EnterpriseRecyclerStatus {
	if in == nil {
		return nil
	}
	out := new(EnterpriseRecyclerStatus)
	in.DeepCopyInto(out)
	return out
}
