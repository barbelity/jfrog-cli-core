package generic

import (
	"errors"
	"github.com/jfrog/jfrog-client-go/utils/errorutils"
	"github.com/jfrog/jfrog-client-go/utils/log"
)

type SetPropsCommand struct {
	PropsCommand
}

func NewSetPropsCommand() *SetPropsCommand {
	return &SetPropsCommand{}
}

func (setProps *SetPropsCommand) SetPropsCommand(command PropsCommand) *SetPropsCommand {
	setProps.PropsCommand = command
	return setProps
}

func (setProps *SetPropsCommand) CommandName() string {
	return "rt_set_properties"
}

func (setProps *SetPropsCommand) Run() error {
	rtDetails, err := setProps.RtDetails()
	if errorutils.CheckError(err) != nil {
		return err
	}
	servicesManager, err := createPropsServiceManager(setProps.threads, rtDetails)
	if err != nil {
		return err
	}

	var errorOccurred = false
	for i := 0; i < len(setProps.Spec().Files); i++ {
		propsParams, err := GetPropsParams(setProps.Spec().Get(i), setProps.props)
		if err != nil {
			errorOccurred = true
			log.Error(err)
			continue
		}

		partialSuccess, partialFailed, err := servicesManager.SetProps(propsParams)
		success := setProps.result.SuccessCount() + partialSuccess
		setProps.result.SetSuccessCount(success)
		failed := setProps.result.FailCount() + partialFailed
		setProps.result.SetFailCount(failed)
		if err != nil {
			errorOccurred = true
			log.Error(err)
			continue
		}
	}

	if errorOccurred {
		return errors.New("Set Properties finished with errors, please review the logs.")
	}
	return err
}
