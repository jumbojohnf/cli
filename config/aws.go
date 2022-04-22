package config

type AWSConfig struct {
	Lambda  *LambdaConfig        `yaml:"lambda,omitempty"`
	Gateway *LambdaGatewayConfig `yaml:"gateway,omitempty"`
}

type LambdaConfig struct {
	RoleName string `yaml:"roleName"`
}

type LambdaGatewayConfig struct {
	SupergraphSDLUpdateInterval int    `yaml:"supergraphSDLUpdateInterval"`
	SupergraphSDLBucket         string `yaml:"supergraphSDLBucket"`
	SupergraphSDLKey            string `yaml:"supergraphSDLKey"`
}
