package flowconsole

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/slackmaster9999/go-flowconsole/flowconsole/jsii"
)

type Postgres interface {
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
	Schema() *string
	SetSchema(val *string)
	Tags() *[]*string
	SetTags(val *[]*string)
	Tone() *string
	SetTone(val *string)
	ExecutesRequest(action *string) Component
	GetDataFrom(target Component, label *string, options ConnectionOptions) Component
	SendsRequest(target Component, label *string, options ConnectionOptions) Component
	Then(target Component) Component
}

// The jsii proxy struct for Postgres
type jsiiProxy_Postgres struct {
	jsiiProxy_Component
}

func (j *jsiiProxy_Postgres) Badge() *string {
	var returns *string
	_jsii_.Get(
		j,
		"badge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Postgres) BelongsTo() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"belongsTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Postgres) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Postgres) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Postgres) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Postgres) Root() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"root",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Postgres) Schema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Postgres) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Postgres) Tone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tone",
		&returns,
	)
	return returns
}


func NewPostgres(args *ComponentArgs) Postgres {
	_init_.Initialize()

	if err := validateNewPostgresParameters(args); err != nil {
		panic(err)
	}
	j := jsiiProxy_Postgres{}

	_jsii_.Create(
		"@flowconsole/sdk.Postgres",
		[]interface{}{args},
		&j,
	)

	return &j
}

func NewPostgres_Override(p Postgres, args *ComponentArgs) {
	_init_.Initialize()

	_jsii_.Create(
		"@flowconsole/sdk.Postgres",
		[]interface{}{args},
		p,
	)
}

func (j *jsiiProxy_Postgres)SetBadge(val *string) {
	_jsii_.Set(
		j,
		"badge",
		val,
	)
}

func (j *jsiiProxy_Postgres)SetBelongsTo(val interface{}) {
	if err := j.validateSetBelongsToParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"belongsTo",
		val,
	)
}

func (j *jsiiProxy_Postgres)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Postgres)SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Postgres)SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Postgres)SetRoot(val interface{}) {
	if err := j.validateSetRootParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"root",
		val,
	)
}

func (j *jsiiProxy_Postgres)SetSchema(val *string) {
	_jsii_.Set(
		j,
		"schema",
		val,
	)
}

func (j *jsiiProxy_Postgres)SetTags(val *[]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_Postgres)SetTone(val *string) {
	_jsii_.Set(
		j,
		"tone",
		val,
	)
}

func (p *jsiiProxy_Postgres) ExecutesRequest(action *string) Component {
	if err := p.validateExecutesRequestParameters(action); err != nil {
		panic(err)
	}
	var returns Component

	_jsii_.Invoke(
		p,
		"executesRequest",
		[]interface{}{action},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Postgres) GetDataFrom(target Component, label *string, options ConnectionOptions) Component {
	if err := p.validateGetDataFromParameters(target, label); err != nil {
		panic(err)
	}
	var returns Component

	_jsii_.Invoke(
		p,
		"getDataFrom",
		[]interface{}{target, label, options},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Postgres) SendsRequest(target Component, label *string, options ConnectionOptions) Component {
	if err := p.validateSendsRequestParameters(target, label); err != nil {
		panic(err)
	}
	var returns Component

	_jsii_.Invoke(
		p,
		"sendsRequest",
		[]interface{}{target, label, options},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Postgres) Then(target Component) Component {
	if err := p.validateThenParameters(target); err != nil {
		panic(err)
	}
	var returns Component

	_jsii_.Invoke(
		p,
		"then",
		[]interface{}{target},
		&returns,
	)

	return returns
}

