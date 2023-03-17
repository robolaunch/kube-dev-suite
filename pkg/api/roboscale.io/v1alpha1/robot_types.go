package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func init() {
	SchemeBuilder.Register(&Robot{}, &RobotList{})
}

//+genclient
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="Distributions",type=string,JSONPath=`.spec.distributions`
//+kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

// Robot is the Schema for the robots API
type Robot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RobotSpec   `json:"spec,omitempty"`
	Status RobotStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// RobotList contains a list of Robot
type RobotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Robot `json:"items"`
}

// ********************************
// Robot types
// ********************************

// ROS distro selection. Allowed distros are Foxy and Galactic. It is aimed to support Humble, Melodic and Noetic in further versions.
// +kubebuilder:validation:Enum=foxy;galactic;noetic;melodic;humble
type ROSDistro string

const (
	// ROS Melodic Morenia
	ROSDistroMelodic ROSDistro = "melodic"
	// ROS Noetic Ninjemys
	ROSDistroNoetic ROSDistro = "noetic"
	// ROS 2 Foxy Fitzroy
	ROSDistroFoxy ROSDistro = "foxy"
	// ROS 2 Galactic Geochelone
	ROSDistroGalactic ROSDistro = "galactic"
	// ROS 2 Humble Hawksbill
	ROSDistroHumble ROSDistro = "humble"
)

// RMW implementation selection. Robot operator currently supports only FastRTPS. See https://docs.ros.org/en/foxy/How-To-Guides/Working-with-multiple-RMW-implementations.html.
// +kubebuilder:validation:Enum=rmw_fastrtps_cpp
type RMWImplementation string

const (
	// Cyclone DDS
	RMWImplementationCycloneDDS RMWImplementation = "rmw_cyclonedds_cpp"
	// FastRTPS
	RMWImplementationFastRTPS RMWImplementation = "rmw_fastrtps_cpp"
	// Connext
	RMWImplementationConnext RMWImplementation = "rmw_connext_cpp"
	// Gurum DDS
	RMWImplementationGurumDDS RMWImplementation = "rmw_gurumdds_cpp"
)

// Storage class configuration for a volume type.
type StorageClassConfig struct {
	// Storage class name
	Name string `json:"name,omitempty"`
	// PVC access mode
	AccessMode corev1.PersistentVolumeAccessMode `json:"accessMode,omitempty"`
}

// Robot's resource limitations.
type Storage struct {
	// Specifies how much storage will be allocated in total.
	// +kubebuilder:default=10000
	Amount int `json:"amount,omitempty"`
	// Storage class selection for robot's volumes.
	StorageClassConfig StorageClassConfig `json:"storageClassConfig,omitempty"`
}

type TLSSecretReference struct {
	// TLS secret object name.
	// +kubebuilder:validation:Required
	Name string `json:"name"`
	// TLS secret object namespace.
	// +kubebuilder:validation:Required
	Namespace string `json:"namespace"`
}

type RootDNSConfig struct {
	// DNS host.
	// +kubebuilder:validation:Required
	Host string `json:"host"`
}

// RobotSpec defines the desired state of Robot
type RobotSpec struct {
	// ROS distro to be used.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:MaxItems=2
	Distributions []ROSDistro `json:"distributions"`
	// Resource limitations of robot containers.
	Storage Storage `json:"storage,omitempty"`
	// Workspace manager template
	WorkspaceManagerTemplate WorkspaceManagerSpec `json:"workspaceManagerTemplate,omitempty"`
	// Build manager template for initial configuration
	BuildManagerTemplate BuildManagerSpec `json:"buildManagerTemplate,omitempty"`
	// Robot development suite template
	RobotDevSuiteTemplate RobotDevSuiteSpec `json:"robotDevSuiteTemplate,omitempty"`
	// Development enabled
	Development bool `json:"development,omitempty"`
	// Root DNS configuration.
	RootDNSConfig RootDNSConfig `json:"rootDNSConfig,omitempty"`
	// TLS secret reference.
	TLSSecretReference TLSSecretReference `json:"tlsSecretRef,omitempty"`
}

type VolumeStatus struct {
	Created                   bool   `json:"created,omitempty"`
	PersistentVolumeClaimName string `json:"persistentVolumeClaimName,omitempty"`
}

type VolumeStatuses struct {
	Var       VolumeStatus `json:"var,omitempty"`
	Etc       VolumeStatus `json:"etc,omitempty"`
	Usr       VolumeStatus `json:"usr,omitempty"`
	Opt       VolumeStatus `json:"opt,omitempty"`
	Workspace VolumeStatus `json:"workspace,omitempty"`
}

type RobotDevSuiteInstanceStatus struct {
	Created bool                `json:"created,omitempty"`
	Status  RobotDevSuiteStatus `json:"status,omitempty"`
}

type JobPhase string

const (
	JobActive    JobPhase = "Active"
	JobSucceeded JobPhase = "Succeeded"
	JobFailed    JobPhase = "Failed"
)

type LoaderJobStatus struct {
	Created bool     `json:"created,omitempty"`
	Phase   JobPhase `json:"phase,omitempty"`
}

type WorkspaceManagerInstanceStatus struct {
	Created bool                   `json:"created,omitempty"`
	Status  WorkspaceManagerStatus `json:"status,omitempty"`
}

type ManagerStatus struct {
	Name    string `json:"name,omitempty"`
	Created bool   `json:"created,omitempty"`
}

type AttachedBuildObject struct {
	Reference corev1.ObjectReference `json:"reference,omitempty"`
	Status    BuildManagerStatus     `json:"status,omitempty"`
}

type AttachedDevObject struct {
	Reference corev1.ObjectReference `json:"reference,omitempty"`
	Status    RobotDevSuiteStatus    `json:"status,omitempty"`
}

// RobotStatus defines the observed state of Robot
type RobotStatus struct {
	// Phase of robot
	Phase RobotPhase `json:"phase,omitempty"`
	// Image of robot
	Image string `json:"image,omitempty"`
	// Node name
	NodeName string `json:"nodeName,omitempty"`
	// Volume status
	VolumeStatuses VolumeStatuses `json:"volumeStatuses,omitempty"`
	// Robot development suite instance status
	RobotDevSuiteStatus RobotDevSuiteInstanceStatus `json:"robotDevSuiteStatus,omitempty"`
	// Loader job status that configures environment
	LoaderJobStatus LoaderJobStatus `json:"loaderJobStatus,omitempty"`
	// Workspace manager status
	WorkspaceManagerStatus WorkspaceManagerInstanceStatus `json:"workspaceManagerStatus,omitempty"`
	// Initial build manager creation status
	InitialBuildManagerStatus ManagerStatus `json:"initialBuildManagerStatus,omitempty"`
	// Initial launch manager creation status
	InitialLaunchManagerStatuses []ManagerStatus `json:"initialLaunchManagerStatuses,omitempty"`
	// Attached build object information
	AttachedBuildObject AttachedBuildObject `json:"attachedBuildObject,omitempty"`
	// Attached dev object information
	AttachedDevObjects []AttachedDevObject `json:"attachedDevObjects,omitempty"`
}

// ********************************
// RobotArtifact types
// ********************************
