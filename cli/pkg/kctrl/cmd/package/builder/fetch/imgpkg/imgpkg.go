package imgpkg

import (
	"github.com/cppforlife/go-cli-ui/ui"
	"github.com/vmware-tanzu/carvel-kapp-controller/cli/pkg/kctrl/cmd/package/builder/build"
	"github.com/vmware-tanzu/carvel-kapp-controller/cli/pkg/kctrl/cmd/package/builder/common"
	pkgui "github.com/vmware-tanzu/carvel-kapp-controller/cli/pkg/kctrl/cmd/package/builder/ui"
	"github.com/vmware-tanzu/carvel-kapp-controller/pkg/apis/kappctrl/v1alpha1"
)

type ImgpkgStep struct {
	pkgAuthoringUI pkgui.IPkgAuthoringUI
	pkgLocation    string
	pkgBuild       *build.PackageBuild
}

func NewImgPkgStep(ui pkgui.IPkgAuthoringUI, pkgLocation string, pkgBuild *build.PackageBuild) *ImgpkgStep {
	imgpkg := ImgpkgStep{
		pkgAuthoringUI: ui,
		pkgLocation:    pkgLocation,
		pkgBuild:       pkgBuild,
	}
	return &imgpkg
}

func (imgpkg ImgpkgStep) PreInteract() error {
	return nil
}

func (imgpkg *ImgpkgStep) Interact() error {
	var isImgpkgBundleCreated bool
	isImgpkgBundleCreated = false
	existingImgPkgBundleConf := imgpkg.pkgBuild.Spec.Pkg.Spec.Template.Spec.Fetch[0].ImgpkgBundle
	if existingImgPkgBundleConf == nil {
		imgpkg.pkgBuild.Spec.Pkg.Spec.Template.Spec.Fetch[0].ImgpkgBundle = &v1alpha1.AppFetchImgpkgBundle{}
	}

	if isImgpkgBundleCreated {
		//TODO Rohit should we add some information here.
		//imgpkg.ui.BeginLinef("")
		textOpts := ui.TextOpts{
			Label:        "Enter the imgpkg bundle url",
			Default:      "",
			ValidateFunc: nil,
		}
		image, err := imgpkg.pkgAuthoringUI.AskForText(textOpts)
		imgpkg.pkgBuild.Spec.Pkg.Spec.Template.Spec.Fetch[0].ImgpkgBundle.Image = image
		imgpkg.pkgBuild.WriteToFile(imgpkg.pkgLocation)
		if err != nil {
			return err
		}
	} else {
		createImgPkgStep := NewCreateImgPkgStep(imgpkg.pkgAuthoringUI, imgpkg.pkgLocation, imgpkg.pkgBuild)
		err := common.Run(createImgPkgStep)
		if err != nil {
			return err
		}
	}
	return nil
}

func (imgpkg ImgpkgStep) PostInteract() error {
	return nil
}