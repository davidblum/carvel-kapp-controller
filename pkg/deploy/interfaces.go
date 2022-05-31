// Copyright 2020 VMware, Inc.
// SPDX-License-Identifier: Apache-2.0

package deploy

import (
	"github.com/vmware-tanzu/carvel-kapp-controller/pkg/exec"
)

type Deploy interface {
	Deploy(tplOutput string, appName string, startedApplyingFunc func(),
		changedFunc func(exec.CmdRunResult)) exec.CmdRunResult

	Delete(appName string, startedApplyingFunc func(),
		changedFunc func(exec.CmdRunResult)) exec.CmdRunResult

	Inspect(appName string) exec.CmdRunResult
}

type GenericOpts struct {
	Name      string
	Namespace string
}

type ProcessedGenericOpts struct {
	Name      string
	Namespace string

	Kubeconfig                    *KubeconfigRestricted
	DangerousUsePodServiceAccount bool
}
