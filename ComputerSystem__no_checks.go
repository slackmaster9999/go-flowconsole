//go:build no_runtime_type_checking

package flowconsole

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ComputerSystem) validateExecutesRequestParameters(action *string) error {
	return nil
}

func (c *jsiiProxy_ComputerSystem) validateGetDataFromParameters(target Component, label *string) error {
	return nil
}

func (c *jsiiProxy_ComputerSystem) validateSendsRequestParameters(target Component, label *string) error {
	return nil
}

func (c *jsiiProxy_ComputerSystem) validateThenParameters(target Component) error {
	return nil
}

func (j *jsiiProxy_ComputerSystem) validateSetBelongsToParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ComputerSystem) validateSetRootParameters(val interface{}) error {
	return nil
}

func validateNewComputerSystemParameters(args *ComponentArgs) error {
	return nil
}

