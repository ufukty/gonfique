package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type ApiVersion string
type ContainerName string
type Name string
type Path struct {
	Backend struct {
		Service struct {
			Name string `yaml:"name"`
			Port struct {
				Number int `yaml:"number"`
			} `yaml:"port"`
		} `yaml:"service"`
	} `yaml:"backend"`
	Path     string `yaml:"path"`
	PathType string `yaml:"pathType"`
}
type Port struct {
	Port       int    `yaml:"port"`
	Protocol   string `yaml:"protocol"`
	TargetPort int    `yaml:"targetPort"`
}
type Rule struct {
	Host string `yaml:"host"`
	Http struct {
		Paths []Path `yaml:"paths"`
	} `yaml:"http"`
}
type SpecContainer struct {
	EnvFrom []struct {
		ConfigMapRef struct {
			Name string `yaml:"name"`
		} `yaml:"configMapRef"`
		SecretRef struct {
			Name string `yaml:"name"`
		} `yaml:"secretRef"`
	} `yaml:"envFrom"`
	Image string        `yaml:"image"`
	Name  ContainerName `yaml:"name"`
	Ports []struct {
		ContainerPort int `yaml:"containerPort"`
	} `yaml:"ports"`
}
type SpecContainers []SpecContainer
type Config struct {
	ApiVersion ApiVersion `yaml:"apiVersion"`
	Data       struct {
		MyKey    string `yaml:"my-key"`
		Password string `yaml:"password"`
	} `yaml:"data"`
	Kind     string `yaml:"kind"`
	Metadata struct {
		Name      Name   `yaml:"name"`
		Namespace string `yaml:"namespace"`
	} `yaml:"metadata"`
	Spec struct {
		Ports    []Port `yaml:"ports"`
		Replicas int    `yaml:"replicas"`
		Rules    []Rule `yaml:"rules"`
		Selector struct {
			MatchLabels struct {
				App string `yaml:"app"`
			} `yaml:"matchLabels"`
		} `yaml:"selector"`
		Template struct {
			Metadata struct {
				Labels struct {
					App string `yaml:"app"`
				} `yaml:"labels"`
			} `yaml:"metadata"`
			Spec struct {
				Containers SpecContainers `yaml:"containers"`
			} `yaml:"spec"`
		} `yaml:"template"`
	} `yaml:"spec"`
	Type string `yaml:"type"`
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
