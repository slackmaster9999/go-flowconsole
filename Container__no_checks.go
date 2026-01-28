//go:build no_runtime_type_checking

package flowconsole

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_Container) validateExecutesRequestParameters(action *string) error {
	return nil
}

func (c *jsiiProxy_Container) validateGetDataFromParameters(target Component, label *string) error {
	return nil
}

func (c *jsiiProxy_Container) validateSendsRequestParameters(target Component, label *string) error {
	return nil
}

func (c *jsiiProxy_Container) validateThenParameters(target Component) error {
	return nil
}

func (j *jsiiProxy_Container) validateSetBelongsToParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Container) validateSetRootParameters(val interface{}) error {
	return nil
}

func validateNewContainerParameters(args *ComponentArgs) error {
	return nil
}

