//go:build no_runtime_type_checking

package flowconsole

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_Postgres) validateExecutesRequestParameters(action *string) error {
	return nil
}

func (p *jsiiProxy_Postgres) validateGetDataFromParameters(target Component, label *string) error {
	return nil
}

func (p *jsiiProxy_Postgres) validateSendsRequestParameters(target Component, label *string) error {
	return nil
}

func (p *jsiiProxy_Postgres) validateThenParameters(target Component) error {
	return nil
}

func (j *jsiiProxy_Postgres) validateSetBelongsToParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Postgres) validateSetRootParameters(val interface{}) error {
	return nil
}

func validateNewPostgresParameters(args *ComponentArgs) error {
	return nil
}

