package flowconsole

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/slackmaster9999/go-flowconsole/flowconsole/jsii"
)

type ComputerSystem interface {
	Component
	Badge() *string
	SetBadge(val *string)
	BelongsTo() interface{}
	SetBelongsTo(val interface{})
	Description() *string
	SetDescription(val *string)
	Domain() *string
	SetDomain(val *string)
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
	ExecutesRequest(action *string) Component
	GetDataFrom(target Component, label *string, options ConnectionOptions) Component
	SendsRequest(target Component, label *string, options ConnectionOptions) Component
	Then(target Component) Component
}

// The jsii proxy struct for ComputerSystem
type jsiiProxy_ComputerSystem struct {
	jsiiProxy_Component
}

func (j *jsiiProxy_ComputerSystem) Badge() *string {
	var returns *string
	_jsii_.Get(
		j,
		"badge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputerSystem) BelongsTo() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"belongsTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputerSystem) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputerSystem) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputerSystem) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputerSystem) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputerSystem) Root() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"root",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputerSystem) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputerSystem) Tone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tone",
		&returns,
	)
	return returns
}


func NewComputerSystem(args *ComponentArgs) ComputerSystem {
	_init_.Initialize()

	if err := validateNewComputerSystemParameters(args); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputerSystem{}

	_jsii_.Create(
		"@flowconsole/sdk.ComputerSystem",
		[]interface{}{args},
		&j,
	)

	return &j
}

func NewComputerSystem_Override(c ComputerSystem, args *ComponentArgs) {
	_init_.Initialize()

	_jsii_.Create(
		"@flowconsole/sdk.ComputerSystem",
		[]interface{}{args},
		c,
	)
}

func (j *jsiiProxy_ComputerSystem)SetBadge(val *string) {
	_jsii_.Set(
		j,
		"badge",
		val,
	)
}

func (j *jsiiProxy_ComputerSystem)SetBelongsTo(val interface{}) {
	if err := j.validateSetBelongsToParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"belongsTo",
		val,
	)
}

func (j *jsiiProxy_ComputerSystem)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ComputerSystem)SetDomain(val *string) {
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_ComputerSystem)SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ComputerSystem)SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ComputerSystem)SetRoot(val interface{}) {
	if err := j.validateSetRootParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"root",
		val,
	)
}

func (j *jsiiProxy_ComputerSystem)SetTags(val *[]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ComputerSystem)SetTone(val *string) {
	_jsii_.Set(
		j,
		"tone",
		val,
	)
}

func (c *jsiiProxy_ComputerSystem) ExecutesRequest(action *string) Component {
	if err := c.validateExecutesRequestParameters(action); err != nil {
		panic(err)
	}
	var returns Component

	_jsii_.Invoke(
		c,
		"executesRequest",
		[]interface{}{action},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputerSystem) GetDataFrom(target Component, label *string, options ConnectionOptions) Component {
	if err := c.validateGetDataFromParameters(target, label); err != nil {
		panic(err)
	}
	var returns Component

	_jsii_.Invoke(
		c,
		"getDataFrom",
		[]interface{}{target, label, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputerSystem) SendsRequest(target Component, label *string, options ConnectionOptions) Component {
	if err := c.validateSendsRequestParameters(target, label); err != nil {
		panic(err)
	}
	var returns Component

	_jsii_.Invoke(
		c,
		"sendsRequest",
		[]interface{}{target, label, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputerSystem) Then(target Component) Component {
	if err := c.validateThenParameters(target); err != nil {
		panic(err)
	}
	var returns Component

	_jsii_.Invoke(
		c,
		"then",
		[]interface{}{target},
		&returns,
	)

	return returns
}

