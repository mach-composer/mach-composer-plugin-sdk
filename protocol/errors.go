package protocol

import "fmt"

type PluginError struct {
	Message    string
	PluginName string
}

func (e *PluginError) Error() string {
	if e.PluginName != "" {
		return fmt.Sprintf("%s %s", e.PluginName, e.Message)
	}
	return e.Message
}

func wrapError(err error) *PluginError {
	if err == nil {
		return nil
	}
	return &PluginError{Message: err.Error()}
}

func unwrapError(name string, err *PluginError) error {
	if err == nil {
		return nil
	}
	err.PluginName = name
	return err
}
