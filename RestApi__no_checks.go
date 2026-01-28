//go:build no_runtime_type_checking

package flowconsole

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RestApi) validateExecutesRequestParameters(action *string) error {
	return nil
}

func (r *jsiiProxy_RestApi) validateGetDataFromParameters(target Component, label *string) error {
	return nil
}

func (r *jsiiProxy_RestApi) validateSendsRequestParameters(target Component, label *string) error {
	return nil
}

func (r *jsiiProxy_RestApi) validateThenParameters(target Component) error {
	return nil
}

func (j *jsiiProxy_RestApi) validateSetBelongsToParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RestApi) validateSetRootParameters(val interface{}) error {
	return nil
}

func validateNewRestApiParameters(args *ComponentArgs) error {
	return nil
}

