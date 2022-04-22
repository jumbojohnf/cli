package config

type AWSConfig struct {
	Gateway *LambdaGatewayConfig `yaml:"gateway,omitempty"`
	Lambda  *LambdaConfig        `yaml:"lambda,omitempty"`
}

type LambdaConfig struct {
	RoleName string `yaml:"roleName"`
}

type LambdaGatewayConfig struct {
	SupergraphSDLBucket         string `yaml:"supergraphSDLBucket"`
	SupergraphSDLKey            string `yaml:"supergraphSDLKey"`
	SupergraphSDLUpdateInterval int    `yaml:"supergraphSDLUpdateInterval"`
}
