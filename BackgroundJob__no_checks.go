//go:build no_runtime_type_checking

package flowconsole

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BackgroundJob) validateExecutesRequestParameters(action *string) error {
	return nil
}

func (b *jsiiProxy_BackgroundJob) validateGetDataFromParameters(target Component, label *string) error {
	return nil
}

func (b *jsiiProxy_BackgroundJob) validateSendsRequestParameters(target Component, label *string) error {
	return nil
}

func (b *jsiiProxy_BackgroundJob) validateThenParameters(target Component) error {
	return nil
}

func (j *jsiiProxy_BackgroundJob) validateSetBelongsToParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BackgroundJob) validateSetRootParameters(val interface{}) error {
	return nil
}

func validateNewBackgroundJobParameters(args *ComponentArgs) error {
	return nil
}

