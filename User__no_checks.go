//go:build no_runtime_type_checking

package flowconsole

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_User) validateExecutesRequestParameters(action *string) error {
	return nil
}

func (u *jsiiProxy_User) validateGetDataFromParameters(target Component, label *string) error {
	return nil
}

func (u *jsiiProxy_User) validateSendsRequestParameters(target Component, label *string) error {
	return nil
}

func (u *jsiiProxy_User) validateThenParameters(target Component) error {
	return nil
}

func (j *jsiiProxy_User) validateSetBelongsToParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_User) validateSetRootParameters(val interface{}) error {
	return nil
}

func validateNewUserParameters(args *UserArgs) error {
	return nil
}

