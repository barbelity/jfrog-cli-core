package generic

import (
	"github.com/jfrog/jfrog-cli-core/artifactory/spec"
	"github.com/jfrog/jfrog-cli-core/utils/config"
	"github.com/jfrog/jfrog-cli-core/utils/coreutils"
	"github.com/jfrog/jfrog-client-go/artifactory"
	"github.com/jfrog/jfrog-client-go/artifactory/services"
	clientConfig "github.com/jfrog/jfrog-client-go/config"
)

type PropsCommand struct {
	props   string
	threads int
	GenericCommand
}

func NewPropsCommand() *PropsCommand {
	return &PropsCommand{GenericCommand: *NewGenericCommand()}
}

func (pc *PropsCommand) Threads() int {
	return pc.threads
}

func (pc *PropsCommand) SetThreads(threads int) *PropsCommand {
	pc.threads = threads
	return pc
}

func (pc *PropsCommand) Props() string {
	return pc.props
}

func (pc *PropsCommand) SetProps(props string) *PropsCommand {
	pc.props = props
	return pc
}

func createPropsServiceManager(threads int, artDetails *config.ArtifactoryDetails) (artifactory.ArtifactoryServicesManager, error) {
	certsPath, err := coreutils.GetJfrogCertsDir()
	if err != nil {
		return nil, err
	}
	artAuth, err := artDetails.CreateArtAuthConfig()
	if err != nil {
		return nil, err
	}
	serviceConfig, err := clientConfig.NewConfigBuilder().
		SetServiceDetails(artAuth).
		SetCertificatesPath(certsPath).
		SetInsecureTls(artDetails.InsecureTls).
		SetThreads(threads).
		Build()

	return artifactory.New(&artAuth, serviceConfig)
}

func GetPropsParams(f *spec.File, properties string) (propsParams services.PropsParams, err error) {
	propsParams = services.NewPropsParams()
	propsParams.Properties = properties
	propsParams.ArtifactoryCommonParams = f.ToArtifactoryCommonParams()
	propsParams.Recursive, err = f.IsRecursive(true)
	if err != nil {
		return
	}
	propsParams.IncludeDirs, err = f.IsIncludeDirs(false)
	return
}
