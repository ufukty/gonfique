package config

type (
	Data struct {
		MyKey    string `yaml:"my-key"`
		Password string `yaml:"password"`
	}
	Metadata struct {
		Name      string `yaml:"name"`
		Namespace string `yaml:"namespace"`
	}
	Port struct {
		Port       int    `yaml:"port"`
		Protocol   string `yaml:"protocol"`
		TargetPort int    `yaml:"targetPort"`
	}
	ServicePort struct {
		Number int `yaml:"number"`
	}
)
