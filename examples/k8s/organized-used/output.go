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
		Name string      `yaml:"name"`
		Port ServicePort `yaml:"port"`
	}
	autoGenB struct {
		Service autoGenA `yaml:"service"`
	}
	autoGenC struct {
		Backend  autoGenB `yaml:"backend"`
		Path     string   `yaml:"path"`
		PathType string   `yaml:"pathType"`
	}
	autoGenD struct {
		Paths []autoGenC `yaml:"paths"`
	}
	autoGenE struct {
		Host string   `yaml:"host"`
		Http autoGenD `yaml:"http"`
	}
	autoGenF struct {
		App string `yaml:"app"`
	}
	autoGenG struct {
		MatchLabels autoGenF `yaml:"matchLabels"`
	}
	autoGenH struct {
		Labels autoGenF `yaml:"labels"`
	}
	autoGenI struct {
		Name string `yaml:"name"`
	}
	autoGenJ struct {
		ConfigMapRef autoGenI `yaml:"configMapRef"`
		SecretRef    autoGenI `yaml:"secretRef"`
	}
	autoGenK struct {
		ContainerPort int `yaml:"containerPort"`
	}
	autoGenL struct {
		EnvFrom []autoGenJ `yaml:"envFrom"`
		Image   string     `yaml:"image"`
		Name    string     `yaml:"name"`
		Ports   []autoGenK `yaml:"ports"`
	}
	autoGenM struct {
		Containers []autoGenL `yaml:"containers"`
	}
	autoGenN struct {
		Metadata autoGenH `yaml:"metadata"`
		Spec     autoGenM `yaml:"spec"`
	}
	autoGenO struct {
		Ports    []Port     `yaml:"ports"`
		Replicas int        `yaml:"replicas"`
		Rules    []autoGenE `yaml:"rules"`
		Selector autoGenG   `yaml:"selector"`
		Template autoGenN   `yaml:"template"`
	}
)

func (c autoGenB) Range() map[string]autoGenA {
	return map[string]autoGenA{"service": c.Service}
}
func (c autoGenF) Range() map[string]string {
	return map[string]string{"app": c.App}
}
func (c autoGenG) Range() map[string]autoGenF {
	return map[string]autoGenF{"matchLabels": c.MatchLabels}
}
func (c autoGenH) Range() map[string]autoGenF {
	return map[string]autoGenF{"labels": c.Labels}
}
func (c autoGenI) Range() map[string]string {
	return map[string]string{"name": c.Name}
}
func (c autoGenJ) Range() map[string]autoGenI {
	return map[string]autoGenI{"configMapRef": c.ConfigMapRef, "secretRef": c.SecretRef}
}
func (c autoGenK) Range() map[string]int {
	return map[string]int{"containerPort": c.ContainerPort}
}

type Config struct {
	ApiVersion string   `yaml:"apiVersion"`
	Data       Data     `yaml:"data"`
	Kind       string   `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata"`
	Spec       autoGenO `yaml:"spec"`
	Type       string   `yaml:"type"`
}

func ReadConfig(path string) (Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return Config{}, fmt.Errorf("opening config file: %w", err)
	}
	defer file.Close()
	c := Config{}
	err = yaml.NewDecoder(file).Decode(&c)
	if err != nil {
		return Config{}, fmt.Errorf("decoding config file: %w", err)
	}
	return c, nil
}
