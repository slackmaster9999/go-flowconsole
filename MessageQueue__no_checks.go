//go:build no_runtime_type_checking

package flowconsole

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MessageQueue) validateExecutesRequestParameters(action *string) error {
	return nil
}

func (m *jsiiProxy_MessageQueue) validateGetDataFromParameters(target Component, label *string) error {
	return nil
}

func (m *jsiiProxy_MessageQueue) validateSendsRequestParameters(target Component, label *string) error {
	return nil
}

func (m *jsiiProxy_MessageQueue) validateThenParameters(target Component) error {
	return nil
}

func (j *jsiiProxy_MessageQueue) validateSetBelongsToParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MessageQueue) validateSetRootParameters(val interface{}) error {
	return nil
}

func validateNewMessageQueueParameters(args *ComponentArgs) error {
	return nil
}

