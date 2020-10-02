package generic

import (
	"errors"
	"github.com/jfrog/jfrog-client-go/utils/errorutils"
	"github.com/jfrog/jfrog-client-go/utils/log"
)

type DeletePropsCommand struct {
	PropsCommand
}

func NewDeletePropsCommand() *DeletePropsCommand {
	return &DeletePropsCommand{}
}

func (deleteProps *DeletePropsCommand) DeletePropsCommand(command PropsCommand) *DeletePropsCommand {
	deleteProps.PropsCommand = command
	return deleteProps
}

func (deleteProps *DeletePropsCommand) CommandName() string {
	return "rt_delete_properties"
}

func (deleteProps *DeletePropsCommand) Run() error {
	rtDetails, err := deleteProps.RtDetails()
	if errorutils.CheckError(err) != nil {
		return err
	}
	servicesManager, err := createPropsServiceManager(deleteProps.threads, rtDetails)
	if err != nil {
		return err
	}

	var errorOccurred = false
	for i := 0; i < len(deleteProps.Spec().Files); i++ {
		propsParams, err := GetPropsParams(deleteProps.Spec().Get(i), deleteProps.props)
		if err != nil {
			errorOccurred = true
			log.Error(err)
			continue
		}

		partialSuccess, partialFailed, err := servicesManager.DeleteProps(propsParams)
		success := deleteProps.result.SuccessCount() + partialSuccess
		deleteProps.result.SetSuccessCount(success)
		failed := deleteProps.result.FailCount() + partialFailed
		deleteProps.result.SetFailCount(failed)
		if err != nil {
			errorOccurred = true
			log.Error(err)
			continue
		}
	}

	if errorOccurred {
		return errors.New("Delete Properties finished with errors, please review the logs.")
	}
	return err
}
