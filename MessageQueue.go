package flowconsole

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/slackmaster9999/go-flowconsole/flowconsole/jsii"
)

type MessageQueue interface {
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
	Throughput() *string
	SetThroughput(val *string)
	Tone() *string
	SetTone(val *string)
	ExecutesRequest(action *string) Component
	GetDataFrom(target Component, label *string, options ConnectionOptions) Component
	SendsRequest(target Component, label *string, options ConnectionOptions) Component
	Then(target Component) Component
}

// The jsii proxy struct for MessageQueue
type jsiiProxy_MessageQueue struct {
	jsiiProxy_Component
}

func (j *jsiiProxy_MessageQueue) Badge() *string {
	var returns *string
	_jsii_.Get(
		j,
		"badge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MessageQueue) BelongsTo() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"belongsTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MessageQueue) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MessageQueue) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MessageQueue) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MessageQueue) Root() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"root",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MessageQueue) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MessageQueue) Throughput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"throughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MessageQueue) Tone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tone",
		&returns,
	)
	return returns
}


func NewMessageQueue(args *ComponentArgs) MessageQueue {
	_init_.Initialize()

	if err := validateNewMessageQueueParameters(args); err != nil {
		panic(err)
	}
	j := jsiiProxy_MessageQueue{}

	_jsii_.Create(
		"@flowconsole/sdk.MessageQueue",
		[]interface{}{args},
		&j,
	)

	return &j
}

func NewMessageQueue_Override(m MessageQueue, args *ComponentArgs) {
	_init_.Initialize()

	_jsii_.Create(
		"@flowconsole/sdk.MessageQueue",
		[]interface{}{args},
		m,
	)
}

func (j *jsiiProxy_MessageQueue)SetBadge(val *string) {
	_jsii_.Set(
		j,
		"badge",
		val,
	)
}

func (j *jsiiProxy_MessageQueue)SetBelongsTo(val interface{}) {
	if err := j.validateSetBelongsToParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"belongsTo",
		val,
	)
}

func (j *jsiiProxy_MessageQueue)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_MessageQueue)SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MessageQueue)SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MessageQueue)SetRoot(val interface{}) {
	if err := j.validateSetRootParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"root",
		val,
	)
}

func (j *jsiiProxy_MessageQueue)SetTags(val *[]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_MessageQueue)SetThroughput(val *string) {
	_jsii_.Set(
		j,
		"throughput",
		val,
	)
}

func (j *jsiiProxy_MessageQueue)SetTone(val *string) {
	_jsii_.Set(
		j,
		"tone",
		val,
	)
}

func (m *jsiiProxy_MessageQueue) ExecutesRequest(action *string) Component {
	if err := m.validateExecutesRequestParameters(action); err != nil {
		panic(err)
	}
	var returns Component

	_jsii_.Invoke(
		m,
		"executesRequest",
		[]interface{}{action},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MessageQueue) GetDataFrom(target Component, label *string, options ConnectionOptions) Component {
	if err := m.validateGetDataFromParameters(target, label); err != nil {
		panic(err)
	}
	var returns Component

	_jsii_.Invoke(
		m,
		"getDataFrom",
		[]interface{}{target, label, options},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MessageQueue) SendsRequest(target Component, label *string, options ConnectionOptions) Component {
	if err := m.validateSendsRequestParameters(target, label); err != nil {
		panic(err)
	}
	var returns Component

	_jsii_.Invoke(
		m,
		"sendsRequest",
		[]interface{}{target, label, options},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MessageQueue) Then(target Component) Component {
	if err := m.validateThenParameters(target); err != nil {
		panic(err)
	}
	var returns Component

	_jsii_.Invoke(
		m,
		"then",
		[]interface{}{target},
		&returns,
	)

	return returns
}

