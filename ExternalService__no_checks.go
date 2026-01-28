//go:build no_runtime_type_checking

package flowconsole

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_ExternalService) validateExecutesRequestParameters(action *string) error {
	return nil
}

func (e *jsiiProxy_ExternalService) validateGetDataFromParameters(target Component, label *string) error {
	return nil
}

func (e *jsiiProxy_ExternalService) validateSendsRequestParameters(target Component, label *string) error {
	return nil
}

func (e *jsiiProxy_ExternalService) validateThenParameters(target Component) error {
	return nil
}

func (j *jsiiProxy_ExternalService) validateSetBelongsToParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ExternalService) validateSetRootParameters(val interface{}) error {
	return nil
}

func validateNewExternalServiceParameters(args *ComponentArgs) error {
	return nil
}

