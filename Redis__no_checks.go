//go:build no_runtime_type_checking

package flowconsole

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_Redis) validateExecutesRequestParameters(action *string) error {
	return nil
}

func (r *jsiiProxy_Redis) validateGetDataFromParameters(target Component, label *string) error {
	return nil
}

func (r *jsiiProxy_Redis) validateSendsRequestParameters(target Component, label *string) error {
	return nil
}

func (r *jsiiProxy_Redis) validateThenParameters(target Component) error {
	return nil
}

func (j *jsiiProxy_Redis) validateSetBelongsToParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Redis) validateSetRootParameters(val interface{}) error {
	return nil
}

func validateNewRedisParameters(args *ComponentArgs) error {
	return nil
}

