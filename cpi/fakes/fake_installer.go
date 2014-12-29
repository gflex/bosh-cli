package fakes

import (
	"fmt"

	bmcloud "github.com/cloudfoundry/bosh-micro-cli/cloud"
	bmmanifest "github.com/cloudfoundry/bosh-micro-cli/deployment/manifest"
	bmrel "github.com/cloudfoundry/bosh-micro-cli/release"
	bmtestutils "github.com/cloudfoundry/bosh-micro-cli/testutils"
)

type InstallInput struct {
	Deployment bmmanifest.CPIDeploymentManifest
	Release    bmrel.Release
	DirectorID string
}

type installOutput struct {
	cloud bmcloud.Cloud
	err   error
}

type ExtractInput struct {
	ReleaseTarballPath string
}

type extractCallback func(releaseTarballPath string) (bmrel.Release, error)

type FakeInstaller struct {
	InstallInputs   []InstallInput
	installBehavior map[string]installOutput

	ExtractInputs   []ExtractInput
	extractBehavior map[string]extractCallback
}

func NewFakeInstaller() *FakeInstaller {
	return &FakeInstaller{
		InstallInputs:   []InstallInput{},
		installBehavior: map[string]installOutput{},

		ExtractInputs:   []ExtractInput{},
		extractBehavior: map[string]extractCallback{},
	}
}

func (f *FakeInstaller) Extract(releaseTarballPath string) (bmrel.Release, error) {
	input := ExtractInput{
		ReleaseTarballPath: releaseTarballPath,
	}
	f.ExtractInputs = append(f.ExtractInputs, input)

	value, err := bmtestutils.MarshalToString(input)
	if err != nil {
		return nil, fmt.Errorf("Could not serialize input %#v", input)
	}

	callback, found := f.extractBehavior[value]
	if found {
		return callback(releaseTarballPath)
	}
	return nil, fmt.Errorf("Unsupported Extract Input: %s", value)
}

func (f *FakeInstaller) Install(deployment bmmanifest.CPIDeploymentManifest, release bmrel.Release, directorID string) (bmcloud.Cloud, error) {
	input := InstallInput{
		Deployment: deployment,
		Release:    release,
		DirectorID: directorID,
	}
	f.InstallInputs = append(f.InstallInputs, input)

	value, err := bmtestutils.MarshalToString(input)
	if err != nil {
		return nil, fmt.Errorf("Could not serialize input %#v", input)
	}

	output, found := f.installBehavior[value]
	if found {
		return output.cloud, output.err
	}
	return nil, fmt.Errorf("Unsupported Install Input: %s", value)
}

func (f *FakeInstaller) SetInstallBehavior(
	deployment bmmanifest.CPIDeploymentManifest,
	release bmrel.Release,
	directorID string,
	cloud bmcloud.Cloud,
	err error,
) error {
	input := InstallInput{
		Deployment: deployment,
		Release:    release,
		DirectorID: directorID,
	}

	value, err := bmtestutils.MarshalToString(input)
	if err != nil {
		return fmt.Errorf("Could not serialize input %#v", input)
	}
	f.installBehavior[value] = installOutput{cloud: cloud, err: err}
	return nil
}

func (f *FakeInstaller) SetExtractBehavior(releaseTarballPath string, callback extractCallback) error {
	input := ExtractInput{
		ReleaseTarballPath: releaseTarballPath,
	}

	value, err := bmtestutils.MarshalToString(input)
	if err != nil {
		return fmt.Errorf("Could not serialize input %#v", input)
	}
	f.extractBehavior[value] = callback
	return nil
}
