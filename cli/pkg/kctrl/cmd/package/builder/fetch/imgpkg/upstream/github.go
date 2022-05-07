package upstream

import (
	"github.com/cppforlife/go-cli-ui/ui"
	"github.com/vmware-tanzu/carvel-kapp-controller/cli/pkg/kctrl/cmd/package/builder/common"
	vendirconf "github.com/vmware-tanzu/carvel-vendir/pkg/vendir/config"
)

const (
	latestVersion = "latest"
)

type GithubStep struct {
	ui            ui.UI
	GithubRelease *vendirconf.DirectoryContentsGithubRelease `json:"githubRelease,omitempty"`
}

func NewGithubStep(ui ui.UI) *GithubStep {
	return &GithubStep{
		ui: ui,
	}
}

func (g *GithubStep) PreInteract() error {
	return nil
}

func (g *GithubStep) PostInteract() error {
	return nil
}

func (g *GithubStep) Interact() error {
	repoSlug, err := g.ui.AskForText("Enter slug for repository(org/repo)")
	if err != nil {
		return err
	}
	var releaseTag string
	var latest bool

	releaseVersion, err := g.getVersion()
	if releaseVersion == latestVersion {
		latest = true
	} else {
		releaseTag = releaseVersion
	}
	if err != nil {
		return err
	}
	//TODO Rohit getting the releaseTag even though it is empty bcoz we dont have omitEmpty in the json representation. Might be have to create PR on imgpkg
	directoryContentsGithubRelease := vendirconf.DirectoryContentsGithubRelease{
		Slug:                          repoSlug,
		Tag:                           releaseTag,
		Latest:                        latest,
		DisableAutoChecksumValidation: true,
	}
	g.GithubRelease = &directoryContentsGithubRelease
	return nil
}

func (g GithubStep) getVersion() (string, error) {
	var useLatestVersion bool
	input, err := g.ui.AskForText("Do you want to use the latest released version(y/n)")
	//TODO Rohit check when you press ctrl-C, does it generate an error

	for {
		if err != nil {
			return "", err
		}
		var isValidInput bool
		useLatestVersion, isValidInput = common.ValidateInputYesOrNo(input)
		if isValidInput {
			break
		} else {
			input, err = g.ui.AskForText("Invalid input. (must be 'y','n','Y','N')")
		}
	}

	if useLatestVersion {
		return latestVersion, nil
	} else {
		g.ui.BeginLinef("Ok. Then we have to mention the specific release tag which makes up the package configuration")
		releaseTag, err := g.ui.AskForText("Enter the release tag")
		if err != nil {
			return "", err
		}
		return releaseTag, nil
	}
	return "", nil
}

func (g *GithubStep) Run() error {
	err := g.PreInteract()
	if err != nil {
		return err
	}
	err = g.Interact()
	if err != nil {
		return err
	}
	err = g.PostInteract()
	if err != nil {
		return err
	}
	return nil
}