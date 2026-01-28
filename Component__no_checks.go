//go:build no_runtime_type_checking

package flowconsole

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_Component) validateExecutesRequestParameters(action *string) error {
	return nil
}

func (c *jsiiProxy_Component) validateGetDataFromParameters(target Component, label *string) error {
	return nil
}

func (c *jsiiProxy_Component) validateSendsRequestParameters(target Component, label *string) error {
	return nil
}

func (c *jsiiProxy_Component) validateThenParameters(target Component) error {
	return nil
}

func (j *jsiiProxy_Component) validateSetBelongsToParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Component) validateSetRootParameters(val interface{}) error {
	return nil
}

func validateNewComponentParameters(args *ComponentArgs) error {
	return nil
}

