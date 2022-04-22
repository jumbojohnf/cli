package lambda

import (
	_ "embed"
	"path/filepath"

	"github.com/funcgql/cli/config"
	"github.com/funcgql/cli/functype"
	"github.com/funcgql/cli/gateway/template/index"
	"github.com/funcgql/cli/template"
	"github.com/pkg/errors"
)

func Export(gatewayDir string, cfg *config.LambdaGatewayConfig) error {
	if cfg == nil {
		return errors.New("missing lambda gateway configuration to export lambda index file")
	}

	content, err := render(cfg)
	if err != nil {
		return errors.Wrap(err, "failed to render lambda gateway index.ts")
	}

	dstPath := filepath.Join(gatewayDir, string(functype.Lambda), index.Filename)
	if _, err := template.Export(content, dstPath); err != nil {
		return err
	}
	return nil
}

//go:embed index.ts.template
var templateContent string

func render(cfg *config.LambdaGatewayConfig) (string, error) {
	type templateData struct {
		SupergraphSDLBucket         string
		SupergraphSDLKey            string
		SupergraphSDLUpdateInterval int
	}

	data := templateData{
		SupergraphSDLBucket:         cfg.SupergraphSDLBucket,
		SupergraphSDLKey:            cfg.SupergraphSDLKey,
		SupergraphSDLUpdateInterval: cfg.SupergraphSDLUpdateInterval,
	}

	return template.Render("index", templateContent, data)
}
