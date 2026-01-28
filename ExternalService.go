package flowconsole

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/slackmaster9999/go-flowconsole/flowconsole/jsii"
)

type ExternalService interface {
	Component
	Badge() *string
	SetBadge(val *string)
	BelongsTo() interface{}
	SetBelongsTo(val interface{})
	Description() *string
	SetDescription(val *string)
	Id() *string
	SetId(val *string)
	Name() *string
	SetName(val *string)
	Root() interface{}
	SetRoot(val interface{})
	Tags() *[]*string
	SetTags(val *[]*string)
	Tone() *string
	SetTone(val *string)
	Vendor() *string
	SetVendor(val *string)
	ExecutesRequest(action *string) Component
	GetDataFrom(target Component, label *string, options ConnectionOptions) Component
	SendsRequest(target Component, label *string, options ConnectionOptions) Component
	Then(target Component) Component
}

// The jsii proxy struct for ExternalService
type jsiiProxy_ExternalService struct {
	jsiiProxy_Component
}

func (j *jsiiProxy_ExternalService) Badge() *string {
	var returns *string
	_jsii_.Get(
		j,
		"badge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalService) BelongsTo() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"belongsTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalService) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalService) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalService) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalService) Root() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"root",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalService) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalService) Tone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalService) Vendor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vendor",
		&returns,
	)
	return returns
}


func NewExternalService(args *ComponentArgs) ExternalService {
	_init_.Initialize()

	if err := validateNewExternalServiceParameters(args); err != nil {
		panic(err)
	}
	j := jsiiProxy_ExternalService{}

	_jsii_.Create(
		"@flowconsole/sdk.ExternalService",
		[]interface{}{args},
		&j,
	)

	return &j
}

func NewExternalService_Override(e ExternalService, args *ComponentArgs) {
	_init_.Initialize()

	_jsii_.Create(
		"@flowconsole/sdk.ExternalService",
		[]interface{}{args},
		e,
	)
}

func (j *jsiiProxy_ExternalService)SetBadge(val *string) {
	_jsii_.Set(
		j,
		"badge",
		val,
	)
}

func (j *jsiiProxy_ExternalService)SetBelongsTo(val interface{}) {
	if err := j.validateSetBelongsToParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"belongsTo",
		val,
	)
}

func (j *jsiiProxy_ExternalService)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ExternalService)SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ExternalService)SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ExternalService)SetRoot(val interface{}) {
	if err := j.validateSetRootParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"root",
		val,
	)
}

func (j *jsiiProxy_ExternalService)SetTags(val *[]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ExternalService)SetTone(val *string) {
	_jsii_.Set(
		j,
		"tone",
		val,
	)
}

func (j *jsiiProxy_ExternalService)SetVendor(val *string) {
	_jsii_.Set(
		j,
		"vendor",
		val,
	)
}

func (e *jsiiProxy_ExternalService) ExecutesRequest(action *string) Component {
	if err := e.validateExecutesRequestParameters(action); err != nil {
		panic(err)
	}
	var returns Component

	_jsii_.Invoke(
		e,
		"executesRequest",
		[]interface{}{action},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalService) GetDataFrom(target Component, label *string, options ConnectionOptions) Component {
	if err := e.validateGetDataFromParameters(target, label); err != nil {
		panic(err)
	}
	var returns Component

	_jsii_.Invoke(
		e,
		"getDataFrom",
		[]interface{}{target, label, options},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalService) SendsRequest(target Component, label *string, options ConnectionOptions) Component {
	if err := e.validateSendsRequestParameters(target, label); err != nil {
		panic(err)
	}
	var returns Component

	_jsii_.Invoke(
		e,
		"sendsRequest",
		[]interface{}{target, label, options},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalService) Then(target Component) Component {
	if err := e.validateThenParameters(target); err != nil {
		panic(err)
	}
	var returns Component

	_jsii_.Invoke(
		e,
		"then",
		[]interface{}{target},
		&returns,
	)

	return returns
}

