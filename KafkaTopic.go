package flowconsole

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/slackmaster9999/go-flowconsole/flowconsole/jsii"
)

type KafkaTopic interface {
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
	PartitionCount() *float64
	SetPartitionCount(val *float64)
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

// The jsii proxy struct for KafkaTopic
type jsiiProxy_KafkaTopic struct {
	jsiiProxy_Component
}

func (j *jsiiProxy_KafkaTopic) Badge() *string {
	var returns *string
	_jsii_.Get(
		j,
		"badge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaTopic) BelongsTo() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"belongsTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaTopic) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaTopic) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaTopic) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaTopic) PartitionCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"partitionCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaTopic) Root() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"root",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaTopic) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaTopic) Tone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tone",
		&returns,
	)
	return returns
}


func NewKafkaTopic(args *ComponentArgs) KafkaTopic {
	_init_.Initialize()

	if err := validateNewKafkaTopicParameters(args); err != nil {
		panic(err)
	}
	j := jsiiProxy_KafkaTopic{}

	_jsii_.Create(
		"@flowconsole/sdk.KafkaTopic",
		[]interface{}{args},
		&j,
	)

	return &j
}

func NewKafkaTopic_Override(k KafkaTopic, args *ComponentArgs) {
	_init_.Initialize()

	_jsii_.Create(
		"@flowconsole/sdk.KafkaTopic",
		[]interface{}{args},
		k,
	)
}

func (j *jsiiProxy_KafkaTopic)SetBadge(val *string) {
	_jsii_.Set(
		j,
		"badge",
		val,
	)
}

func (j *jsiiProxy_KafkaTopic)SetBelongsTo(val interface{}) {
	if err := j.validateSetBelongsToParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"belongsTo",
		val,
	)
}

func (j *jsiiProxy_KafkaTopic)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_KafkaTopic)SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_KafkaTopic)SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_KafkaTopic)SetPartitionCount(val *float64) {
	_jsii_.Set(
		j,
		"partitionCount",
		val,
	)
}

func (j *jsiiProxy_KafkaTopic)SetRoot(val interface{}) {
	if err := j.validateSetRootParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"root",
		val,
	)
}

func (j *jsiiProxy_KafkaTopic)SetTags(val *[]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_KafkaTopic)SetTone(val *string) {
	_jsii_.Set(
		j,
		"tone",
		val,
	)
}

func (k *jsiiProxy_KafkaTopic) ExecutesRequest(action *string) Component {
	if err := k.validateExecutesRequestParameters(action); err != nil {
		panic(err)
	}
	var returns Component

	_jsii_.Invoke(
		k,
		"executesRequest",
		[]interface{}{action},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KafkaTopic) GetDataFrom(target Component, label *string, options ConnectionOptions) Component {
	if err := k.validateGetDataFromParameters(target, label); err != nil {
		panic(err)
	}
	var returns Component

	_jsii_.Invoke(
		k,
		"getDataFrom",
		[]interface{}{target, label, options},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KafkaTopic) SendsRequest(target Component, label *string, options ConnectionOptions) Component {
	if err := k.validateSendsRequestParameters(target, label); err != nil {
		panic(err)
	}
	var returns Component

	_jsii_.Invoke(
		k,
		"sendsRequest",
		[]interface{}{target, label, options},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KafkaTopic) Then(target Component) Component {
	if err := k.validateThenParameters(target); err != nil {
		panic(err)
	}
	var returns Component

	_jsii_.Invoke(
		k,
		"then",
		[]interface{}{target},
		&returns,
	)

	return returns
}

