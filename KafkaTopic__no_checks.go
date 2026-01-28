//go:build no_runtime_type_checking

package flowconsole

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KafkaTopic) validateExecutesRequestParameters(action *string) error {
	return nil
}

func (k *jsiiProxy_KafkaTopic) validateGetDataFromParameters(target Component, label *string) error {
	return nil
}

func (k *jsiiProxy_KafkaTopic) validateSendsRequestParameters(target Component, label *string) error {
	return nil
}

func (k *jsiiProxy_KafkaTopic) validateThenParameters(target Component) error {
	return nil
}

func (j *jsiiProxy_KafkaTopic) validateSetBelongsToParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_KafkaTopic) validateSetRootParameters(val interface{}) error {
	return nil
}

func validateNewKafkaTopicParameters(args *ComponentArgs) error {
	return nil
}

