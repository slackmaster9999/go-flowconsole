package flowconsole

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/slackmaster9999/go-flowconsole/flowconsole/jsii"
)

type Redis interface {
	Component
	Badge() *string
	SetBadge(val *string)
	BelongsTo() interface{}
	SetBelongsTo(val interface{})
	Cluster() *string
	SetCluster(val *string)
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
	ExecutesRequest(action *string) Component
	GetDataFrom(target Component, label *string, options ConnectionOptions) Component
	SendsRequest(target Component, label *string, options ConnectionOptions) Component
	Then(target Component) Component
}

// The jsii proxy struct for Redis
type jsiiProxy_Redis struct {
	jsiiProxy_Component
}

func (j *jsiiProxy_Redis) Badge() *string {
	var returns *string
	_jsii_.Get(
		j,
		"badge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Redis) BelongsTo() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"belongsTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Redis) Cluster() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Redis) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Redis) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Redis) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Redis) Root() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"root",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Redis) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Redis) Tone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tone",
		&returns,
	)
	return returns
}


func NewRedis(args *ComponentArgs) Redis {
	_init_.Initialize()

	if err := validateNewRedisParameters(args); err != nil {
		panic(err)
	}
	j := jsiiProxy_Redis{}

	_jsii_.Create(
		"@flowconsole/sdk.Redis",
		[]interface{}{args},
		&j,
	)

	return &j
}

func NewRedis_Override(r Redis, args *ComponentArgs) {
	_init_.Initialize()

	_jsii_.Create(
		"@flowconsole/sdk.Redis",
		[]interface{}{args},
		r,
	)
}

func (j *jsiiProxy_Redis)SetBadge(val *string) {
	_jsii_.Set(
		j,
		"badge",
		val,
	)
}

func (j *jsiiProxy_Redis)SetBelongsTo(val interface{}) {
	if err := j.validateSetBelongsToParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"belongsTo",
		val,
	)
}

func (j *jsiiProxy_Redis)SetCluster(val *string) {
	_jsii_.Set(
		j,
		"cluster",
		val,
	)
}

func (j *jsiiProxy_Redis)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Redis)SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Redis)SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Redis)SetRoot(val interface{}) {
	if err := j.validateSetRootParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"root",
		val,
	)
}

func (j *jsiiProxy_Redis)SetTags(val *[]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_Redis)SetTone(val *string) {
	_jsii_.Set(
		j,
		"tone",
		val,
	)
}

func (r *jsiiProxy_Redis) ExecutesRequest(action *string) Component {
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

func (r *jsiiProxy_Redis) GetDataFrom(target Component, label *string, options ConnectionOptions) Component {
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

func (r *jsiiProxy_Redis) SendsRequest(target Component, label *string, options ConnectionOptions) Component {
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

func (r *jsiiProxy_Redis) Then(target Component) Component {
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

