package config

import (
	"fmt"
	"os"
	"gopkg.in/yaml.v3"
)
// IMPORTANT:
// Types are defined only for internal purposes.
// Do not refer auto generated type names from outside.
// Because they will change as config schema changes.
type (
	autoGenA struct {
		MyKey    string `yaml:"my-key"`
		Password string `yaml:"password"`
	}
	autoGenB struct {
		Name      string `yaml:"name"`
		Namespace string `yaml:"namespace"`
	}
	autoGenC struct {
		Port       int    `yaml:"port"`
		Protocol   string `yaml:"protocol"`
		TargetPort int    `yaml:"targetPort"`
	}
	autoGenD struct {
		Number int `yaml:"number"`
	}
	autoGenE struct {
		Name string   `yaml:"name"`
		Port autoGenD `yaml:"port"`
	}
	autoGenF struct {
		Service autoGenE `yaml:"service"`
	}
	autoGenG struct {
		Backend  autoGenF `yaml:"backend"`
		Path     string   `yaml:"path"`
		PathType string   `yaml:"pathType"`
	}
	autoGenH struct {
		Paths []autoGenG `yaml:"paths"`
	}
	autoGenI struct {
		Host string   `yaml:"host"`
		Http autoGenH `yaml:"http"`
	}
	autoGenJ struct {
		App string `yaml:"app"`
	}
	autoGenK struct {
		MatchLabels autoGenJ `yaml:"matchLabels"`
	}
	autoGenL struct {
		Labels autoGenJ `yaml:"labels"`
	}
	autoGenM struct {
		Name string `yaml:"name"`
	}
	autoGenN struct {
		ConfigMapRef autoGenM `yaml:"configMapRef"`
		SecretRef    autoGenM `yaml:"secretRef"`
	}
	autoGenO struct {
		ContainerPort int `yaml:"containerPort"`
	}
	autoGenP struct {
		EnvFrom []autoGenN `yaml:"envFrom"`
		Image   string     `yaml:"image"`
		Name    string     `yaml:"name"`
		Ports   []autoGenO `yaml:"ports"`
	}
	autoGenQ struct {
		Containers []autoGenP `yaml:"containers"`
	}
	autoGenR struct {
		Metadata autoGenL `yaml:"metadata"`
		Spec     autoGenQ `yaml:"spec"`
	}
	autoGenS struct {
		Ports    []autoGenC `yaml:"ports"`
		Replicas int        `yaml:"replicas"`
		Rules    []autoGenI `yaml:"rules"`
		Selector autoGenK   `yaml:"selector"`
		Template autoGenR   `yaml:"template"`
	}
)

func (a autoGenA) Range() map[string]string {
	return map[string]string{"my-key": a.MyKey, "password": a.Password}
}
func (a autoGenB) Range() map[string]string {
	return map[string]string{"name": a.Name, "namespace": a.Namespace}
}
func (a autoGenD) Range() map[string]int {
	return map[string]int{"number": a.Number}
}
func (a autoGenF) Range() map[string]autoGenE {
	return map[string]autoGenE{"service": a.Service}
}
func (a autoGenJ) Range() map[string]string {
	return map[string]string{"app": a.App}
}
func (a autoGenK) Range() map[string]autoGenJ {
	return map[string]autoGenJ{"matchLabels": a.MatchLabels}
}
func (a autoGenL) Range() map[string]autoGenJ {
	return map[string]autoGenJ{"labels": a.Labels}
}
func (a autoGenM) Range() map[string]string {
	return map[string]string{"name": a.Name}
}
func (a autoGenN) Range() map[string]autoGenM {
	return map[string]autoGenM{"configMapRef": a.ConfigMapRef, "secretRef": a.SecretRef}
}
func (a autoGenO) Range() map[string]int {
	return map[string]int{"containerPort": a.ContainerPort}
}

type Config struct {
	ApiVersion string   `yaml:"apiVersion"`
	Data       autoGenA `yaml:"data"`
	Kind       string   `yaml:"kind"`
	Metadata   autoGenB `yaml:"metadata"`
	Spec       autoGenS `yaml:"spec"`
	Type       string   `yaml:"type"`
}

func ReadConfig(path string) (Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return Config{}, fmt.Errorf("opening config file: %w", err)
	}
	cfg := Config{}
	err = yaml.NewDecoder(f).Decode(&cfg)
	if err != nil {
		return Config{}, fmt.Errorf("decoding config file: %w", err)
	}
	return cfg, nil
}
