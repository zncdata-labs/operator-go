package persistence

import corev1 "k8s.io/api/core/v1"

type PersistenceSpec struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:default:=true
	Enable bool `json:"enable,omitempty"`
	// +kubebuilder:validation:Optional
	ExistingClaim *string `json:"existingClaim,omitempty"`
	// +kubebuilder:validation:Optional
	Annotations map[string]string `json:"annotations,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:default:="10Gi"
	StorageSize string `json:"storageSize,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:default:="ReadWriteOnce"
	AccessMode string `json:"accessMode,omitempty"`
	// +kubebuilder:validation:Optional
	StorageClassName *string `json:"storageClassName,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=Filesystem
	VolumeMode *corev1.PersistentVolumeMode `json:"volumeMode,omitempty"`
	// +kubebuilder:validation:Optional
	StorageClass *string `json:"storageClass,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:default={ReadWriteOnce}
	AccessModes []corev1.PersistentVolumeAccessMode `json:"accessModes,omitempty"`
	// +kubebuilder:default="10Gi"
	Size string `json:"size,omitempty"`

	// +kubebuilder:validation:Optional
	Labels map[string]string `json:"labels,omitempty"`
}

func (p *PersistenceSpec) Existing() bool {
	return p.ExistingClaim != nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistenceSpec) DeepCopyInto(out *PersistenceSpec) {
	*out = *in
	if in.StorageClass != nil {
		in, out := &in.StorageClass, &out.StorageClass
		*out = new(string)
		**out = **in
	}
	if in.AccessModes != nil {
		in, out := &in.AccessModes, &out.AccessModes
		*out = make([]corev1.PersistentVolumeAccessMode, len(*in))
		copy(*out, *in)
	}
	if in.ExistingClaim != nil {
		in, out := &in.ExistingClaim, &out.ExistingClaim
		*out = new(string)
		**out = **in
	}
	if in.VolumeMode != nil {
		in, out := &in.VolumeMode, &out.VolumeMode
		*out = new(corev1.PersistentVolumeMode)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistenceSpec.
func (in *PersistenceSpec) DeepCopy() *PersistenceSpec {
	if in == nil {
		return nil
	}
	out := new(PersistenceSpec)
	in.DeepCopyInto(out)
	return out
}
