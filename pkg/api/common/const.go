package common

// the orchestrators supported
const (
	// Mesos is the string constant for MESOS orchestrator type
	Mesos string = "Mesos"
	// DCOS is the string constant for DCOS orchestrator type and defaults to DCOS188
	DCOS string = "DCOS"
	// Swarm is the string constant for the Swarm orchestrator type
	Swarm string = "Swarm"
	// Kubernetes is the string constant for the Kubernetes orchestrator type
	Kubernetes string = "Kubernetes"
	// SwarmMode is the string constant for the Swarm Mode orchestrator type
	SwarmMode string = "SwarmMode"
)

// validation values
const (
	// MinAgentCount are the minimum number of agents per agent pool
	MinAgentCount = 1
	// MaxAgentCount are the maximum number of agents per agent pool
	MaxAgentCount = 100
	// MinPort specifies the minimum tcp port to open
	MinPort = 1
	// MaxPort specifies the maximum tcp port to open
	MaxPort = 65535
	// MaxDisks specifies the maximum attached disks to add to the cluster
	MaxDisks = 4
	// MinDiskSizeGB specifies the minimum attached disk size
	MinDiskSizeGB = 1
	// MaxDiskSizeGB specifies the maximum attached disk size
	MaxDiskSizeGB = 1023
	// MinIPAddressCount specifies the minimum number of IP addresses per network interface
	MinIPAddressCount = 1
	// MaxIPAddressCount specifies the maximum number of IP addresses per network interface
	MaxIPAddressCount = 256
)

// Availability profiles
const (
	// AvailabilitySet means that the vms are in an availability set
	AvailabilitySet = "AvailabilitySet"
	// VirtualMachineScaleSets means that the vms are in a virtual machine scaleset
	VirtualMachineScaleSets = "VirtualMachineScaleSets"
)

// storage profiles
const (
	// StorageAccount means that the nodes use raw storage accounts for their os and attached volumes
	StorageAccount = "StorageAccount"
	// ManagedDisks means that the nodes use managed disks for their os and attached volumes
	ManagedDisks = "ManagedDisks"
)

const (
	// KubernetesVersion1Dot8Dot2 is the major.minor.patch string for 1.8.2 versions of kubernetes
	KubernetesVersion1Dot8Dot2 string = "1.8.2"
	// KubernetesVersion1Dot7Dot9 is the major.minor.patch string for 1.7.9 versions of kubernetes
	KubernetesVersion1Dot7Dot9 string = "1.7.9"
	// KubernetesVersion1Dot6Dot11 is the major.minor.patch string for 1.6.11 versions of kubernetes
	KubernetesVersion1Dot6Dot11 string = "1.6.11"
	// KubernetesVersion1Dot5Dot8 is the major.minor.patch string for 1.5.8 versions of kubernetes
	KubernetesVersion1Dot5Dot8 string = "1.5.8"
	// KubernetesDefaultVersion is the default major.minor.patch version for kubernetes
	KubernetesDefaultVersion string = KubernetesVersion1Dot7Dot9
)

// AllKubernetesSupportedVersions maintain a list of available k8s versions in acs-engine
var AllKubernetesSupportedVersions = []string{
	KubernetesVersion1Dot8Dot2,
	KubernetesVersion1Dot7Dot9,
	KubernetesVersion1Dot6Dot11,
	KubernetesVersion1Dot5Dot8,
}

const (
	// DCOSVersion1Dot10Dot0 is the major.minor.patch string for 1.10.0 versions of DCOS
	DCOSVersion1Dot10Dot0 string = "1.10.0"
	// DCOSVersion1Dot9Dot0 is the major.minor.patch string for 1.9.0 versions of DCOS
	DCOSVersion1Dot9Dot0 string = "1.9.0"
	// DCOSVersion1Dot8Dot8 is the major.minor.patch string for 1.8.8 versions of DCOS
	DCOSVersion1Dot8Dot8 string = "1.8.8"
	// DCOSDefaultVersion is the default major.minor.patch version for DCOS
	DCOSDefaultVersion string = DCOSVersion1Dot9Dot0
)

// AllDCOSSupportedVersions maintain a list of available dcos versions in acs-engine
var AllDCOSSupportedVersions = []string{
	DCOSVersion1Dot10Dot0,
	DCOSVersion1Dot9Dot0,
	DCOSVersion1Dot8Dot8,
}
