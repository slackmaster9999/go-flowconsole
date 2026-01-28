//go:build no_runtime_type_checking

package flowconsole

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_ReactApp) validateExecutesRequestParameters(action *string) error {
	return nil
}

func (r *jsiiProxy_ReactApp) validateGetDataFromParameters(target Component, label *string) error {
	return nil
}

func (r *jsiiProxy_ReactApp) validateSendsRequestParameters(target Component, label *string) error {
	return nil
}

func (r *jsiiProxy_ReactApp) validateThenParameters(target Component) error {
	return nil
}

func (j *jsiiProxy_ReactApp) validateSetBelongsToParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ReactApp) validateSetRootParameters(val interface{}) error {
	return nil
}

func validateNewReactAppParameters(args *ComponentArgs) error {
	return nil
}

