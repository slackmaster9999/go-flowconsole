//go:build !no_runtime_type_checking

package flowconsole

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (r *jsiiProxy_RestApi) validateExecutesRequestParameters(action *string) error {
	if action == nil {
		return fmt.Errorf("parameter action is required, but nil was provided")
	}

	return nil
}

func (r *jsiiProxy_RestApi) validateGetDataFromParameters(target Component, label *string) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	if label == nil {
		return fmt.Errorf("parameter label is required, but nil was provided")
	}

	return nil
}

func (r *jsiiProxy_RestApi) validateSendsRequestParameters(target Component, label *string) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	if label == nil {
		return fmt.Errorf("parameter label is required, but nil was provided")
	}

	return nil
}

func (r *jsiiProxy_RestApi) validateThenParameters(target Component) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_RestApi) validateSetBelongsToParameters(val interface{}) error {
	switch val.(type) {
	case Container:
		// ok
	case ComputerSystem:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: Container, ComputerSystem; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_RestApi) validateSetRootParameters(val interface{}) error {
	switch val.(type) {
	case Container:
		// ok
	case ComputerSystem:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: Container, ComputerSystem; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func validateNewRestApiParameters(args *ComponentArgs) error {
	if args == nil {
		return fmt.Errorf("parameter args is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(args, func() string { return "parameter args" }); err != nil {
		return err
	}

	return nil
}

