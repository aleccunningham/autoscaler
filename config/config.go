// Copyright 2018 Drone.IO Inc
// Use of this software is governed by the Business Source License
// that can be found in the LICENSE file.

package config

import "time"

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PipelineList is a list of Pipelines.
type PipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Autoscaler `json:"items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Autoscaler is a duke resource defining a CI lifecycle
type Autoscaler struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the ddesired behaviour of the pod terminator.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status
	// +optional
	Spec   AutoscalerSpec 	`json:"spec"`
	Status State        	`json:"state,omitempty"`
}

// AutoscalerSpec describes the specification for a Cloud Native Continous Delivery pipeline using Kubernetes as a build manager
type AutoscalerSpec struct {
	// Selector is how the target will be selected
	Selector map[string]string `json:"selector,omitempty"`
	Config Config `json:"config,omitempty"`
	Slack Slack `json:"slack,omitempty"`
	Logs Logs `json:"logs,omitempty"`
	Pool Pool `json:"pool,omitempty"`
	Server Server `json:"server,omitempty"`
	Agent Agent `json:"agent,omitempty"`
	HTTP HTTP `json:"http,omitempty"`
	TLS TLS `json:tls,omitempty"`
	Prometheus Prometheus `json:"prometheus,omitempty"`
	//TODO cloud providers
}

// Config stores the configuration settings.
type Config struct {
	License  string `json:"license,omitempty"` 
	Interval time.Duration `default:"5m" json:"interval,omitempty`

type Slack struct {
	Webhook string `json:"slackWebhook,omitempty"`
}

type Logs struct {
	Color  bool `json:"color,omitempty"`
	Debug  bool `default:"true" json:"debug,omitempty`
	Pretty bool `json:"pretty,omitempty"`
}

type Pool struct {
	Min    int           `default:"2" json:"mix,omitempty"`
	Max    int           `default:"4" json:"max,omitempty"`
	MinAge time.Duration `default:"55m" split_words:"true" json:"minAge,omitempty"`
}

type Server struct {
	Host  string `json:"host,omitempty"`
	Proto string `json:"proto,omitempty"`
	Token string `json:"token,omitempty"`
}

type Agent struct {
	Host        string `json:"host,omitempty"`
	Token       string `json:"token,omitempty"`
	Image       string `default:"drone/agent:0.8" json:"image,omitempty"`
	Concurrency int    `default:"2" json:"concurrency,omitempty"`
}

type HTTP struct {
	Host string `json:"host,omitempty"`
	Port string `default:":8080" json:"port,omitempty"`
	Root string `default:"/" json:"root,omitempty"`	
}

type TLS struct {
	Autocert bool `json:"autocert,omitempty"`
	Cert     string `json:"cert,omitempty"`
	Key      string `json:"key,omitempty"`
}

type Prometheus struct {
	AuthToken string `split_words:"true" json:"authToken,omitempty"`
}

type Database struct {
	Driver     string `default:"sqlite3" json:"driver,omitempty"`
	Datasource string `default:"database.sqlite?cache=shared&mode=rwc&_busy_timeout=9999999" json:"dataSource,omitempty"`
}

type Amazon struct {
	Image         string `json:"image,omitempty"`
	Instance      string `json:"instance,omitempty"`
	PrivateIP     bool `split_words:"true" json:"privateIP"`
	Region        string `json:"region,omitempty"`
	Retries       int `json:"retries,omitempty"`
	SSHKey        string `json:"SSHKey,omitempty"`
	SubnetID      string   `split_words:"true" json:"subnetID"`
	SecurityGroup []string `split_words:"true" json:"securityGroup"`
	Tags          map[string]string `json:"tags,omitempty"`
	UserData      string `envconfig:"DRONE_AMAZON_USERDATA" json:"userData"`
	UserDataFile  string `envconfig:"DRONE_AMAZON_USERDATA_FILE"`
}

type DigitalOcean struct {
	Token        string `json:"token,omitempty"`
	Image        string `json:"image,omitempty"`
	Region       string `json:"region,omitempty"`
	SSHKey       string `json:"SSHKey,omitempty"`	
	Size         string `json:"size,omitempty"`
	Tags         []string `json:"tags,omitempty"`
	UserData     string `envconfig:"DRONE_DIGITALOCEAN_USERDATA" json:"userData,omitempty"`
	UserDataFile string `envconfig:"DRONE_DIGITALOCEAN_USERDATA_FILE" json:"userDataFile,omitempty"`
}

type Google struct {
	MachineType  string            `envconfig:"DRONE_GOOGLE_MACHINE_TYPE" json:"machineType,omitempty"`
	MachineImage string            `envconfig:"DRONE_GOOGLE_MACHINE_IMAGE" json:"machineImage,omitempty"`
	Network      string            `envconfig:"DRONE_GOOGLE_NETWORK" json:"network,omitempty"`
	Labels       map[string]string `envconfig:"DRONE_GOOGLE_LABELS" json:"labels,omitempty"`
	Scopes       string            `envconfig:"DRONE_GOOGLE_SCOPES" json:"scopes,omitempty"`
	DiskSize     int64             `envconfig:"DRONE_GOOGLE_DISK_SIZE" json:"diskSize,omitempty"`
	DiskType     string            `envconfig:"DRONE_GOOGLE_DISK_TYPE" json:"diskType,omitempty"`
	Project      string            `envconfig:"DRONE_GOOGLE_PROJECT" json:"project,omitempty"`
	Tags         []string          `envconfig:"DRONE_GOOGLE_TAGS" json:"tags,omitempty"`
	UserData     string            `envconfig:"DRONE_GOOGLE_USERDATA" json:"userData,omitempty"`
	UserDataFile string            `envconfig:"DRONE_GOOGLE_USERDATA_FILE" json:"userDataFile,omitempty"`
	Zone         string            `envconfig:"DRONE_GOOGLE_ZONE" json:"zone,omitempty"`
}

type HetznerCloud struct {
	Datacenter   string `json:"dataCe ter,omitempty"`
	Image        string `json:"image,omitempty"`
	SSHKey       int `json:"SSHKey,omitempty"`
	Token        string `json:"token,omitempty"`
	Type         string `json:"type,omitempty"`
	UserData     string `envconfig:"DRONE_HETZNERCLOUD_USERDATA" json:"userData,omitempty"`
	UserDataFile string `envconfig:"DRONE_HETZNERCLOUD_USERDATA_FILE" json:"userDataFile,omitempty"`
}
