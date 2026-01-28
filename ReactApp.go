package flowconsole

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/slackmaster9999/go-flowconsole/flowconsole/jsii"
)

type ReactApp interface {
	Component
	Badge() *string
	SetBadge(val *string)
	BelongsTo() interface{}
	SetBelongsTo(val interface{})
	Description() *string
	SetDescription(val *string)
	Framework() *string
	SetFramework(val *string)
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
	Url() *string
	SetUrl(val *string)
	ExecutesRequest(action *string) Component
	GetDataFrom(target Component, label *string, options ConnectionOptions) Component
	SendsRequest(target Component, label *string, options ConnectionOptions) Component
	Then(target Component) Component
}

// The jsii proxy struct for ReactApp
type jsiiProxy_ReactApp struct {
	jsiiProxy_Component
}

func (j *jsiiProxy_ReactApp) Badge() *string {
	var returns *string
	_jsii_.Get(
		j,
		"badge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactApp) BelongsTo() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"belongsTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactApp) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactApp) Framework() *string {
	var returns *string
	_jsii_.Get(
		j,
		"framework",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactApp) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactApp) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactApp) Root() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"root",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactApp) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactApp) Tone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactApp) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}


func NewReactApp(args *ComponentArgs) ReactApp {
	_init_.Initialize()

	if err := validateNewReactAppParameters(args); err != nil {
		panic(err)
	}
	j := jsiiProxy_ReactApp{}

	_jsii_.Create(
		"@flowconsole/sdk.ReactApp",
		[]interface{}{args},
		&j,
	)

	return &j
}

func NewReactApp_Override(r ReactApp, args *ComponentArgs) {
	_init_.Initialize()

	_jsii_.Create(
		"@flowconsole/sdk.ReactApp",
		[]interface{}{args},
		r,
	)
}

func (j *jsiiProxy_ReactApp)SetBadge(val *string) {
	_jsii_.Set(
		j,
		"badge",
		val,
	)
}

func (j *jsiiProxy_ReactApp)SetBelongsTo(val interface{}) {
	if err := j.validateSetBelongsToParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"belongsTo",
		val,
	)
}

func (j *jsiiProxy_ReactApp)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ReactApp)SetFramework(val *string) {
	_jsii_.Set(
		j,
		"framework",
		val,
	)
}

func (j *jsiiProxy_ReactApp)SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ReactApp)SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ReactApp)SetRoot(val interface{}) {
	if err := j.validateSetRootParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"root",
		val,
	)
}

func (j *jsiiProxy_ReactApp)SetTags(val *[]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ReactApp)SetTone(val *string) {
	_jsii_.Set(
		j,
		"tone",
		val,
	)
}

func (j *jsiiProxy_ReactApp)SetUrl(val *string) {
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (r *jsiiProxy_ReactApp) ExecutesRequest(action *string) Component {
	if err := r.validateExecutesRequestParameters(action); err != nil {
		panic(err)
	}
	var returns Component

	_jsii_.Invoke(
		r,
		"executesRequest",
		[]interface{}{action},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReactApp) GetDataFrom(target Component, label *string, options ConnectionOptions) Component {
	if err := r.validateGetDataFromParameters(target, label); err != nil {
		panic(err)
	}
	var returns Component

	_jsii_.Invoke(
		r,
		"getDataFrom",
		[]interface{}{target, label, options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReactApp) SendsRequest(target Component, label *string, options ConnectionOptions) Component {
	if err := r.validateSendsRequestParameters(target, label); err != nil {
		panic(err)
	}
	var returns Component

	_jsii_.Invoke(
		r,
		"sendsRequest",
		[]interface{}{target, label, options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReactApp) Then(target Component) Component {
	if err := r.validateThenParameters(target); err != nil {
		panic(err)
	}
	var returns Component

	_jsii_.Invoke(
		r,
		"then",
		[]interface{}{target},
		&returns,
	)

	return returns
}

