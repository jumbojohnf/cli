package aws

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/funcgql/cli/cliio"
	"github.com/pkg/errors"
)

//go:embed lambda-trust-policy.json
var trustPolicyData []byte

func (a *API) CreateLambdaRole() error {
	fmt.Println("👮‍♀️ Creating lambda execution role")

	policyFile, err := cliio.TempFile("policy")
	if err != nil {
		return err
	}
	if err := policyFile.WriteBytes(trustPolicyData); err != nil {
		return errors.Wrap(err, "failed to export lambda role policy content")
	}

	if output, err := a.execute(
		"iam", "create-role",
		"--role-name", a.cfg.Lambda.RoleName,
		"--assume-role-policy-document", fmt.Sprintf("file://%s", policyFile.AbsPath()),
	); err != nil {
		if output.ExitCode != 254 &&
			!strings.Contains(output.Combined, fmt.Sprintf("%s already exists", a.cfg.Lambda.RoleName)) {
			return errors.Wrapf(err, "failed to create lambda execution role %s", output.Combined)
		}
	}

	if output, err := a.execute(
		"iam", "attach-role-policy",
		"--role-name", a.cfg.Lambda.RoleName,
		"--policy-arn", "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole",
	); err != nil {
		return errors.Wrapf(err, "failed to attach lambda role policy %s", output.Combined)
	}

	return nil
}
