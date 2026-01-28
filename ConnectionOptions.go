package flowconsole

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/slackmaster9999/go-flowconsole/flowconsole/jsii"
)

type ConnectionOptions interface {
	Detail() *string
	SetDetail(val *string)
	Icon() *string
	SetIcon(val *string)
	Kind() *string
	SetKind(val *string)
	Muted() *bool
	SetMuted(val *bool)
}

// The jsii proxy struct for ConnectionOptions
type jsiiProxy_ConnectionOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_ConnectionOptions) Detail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"detail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectionOptions) Icon() *string {
	var returns *string
	_jsii_.Get(
		j,
		"icon",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectionOptions) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectionOptions) Muted() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"muted",
		&returns,
	)
	return returns
}


func NewConnectionOptions() ConnectionOptions {
	_init_.Initialize()

	j := jsiiProxy_ConnectionOptions{}

	_jsii_.Create(
		"@flowconsole/sdk.ConnectionOptions",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewConnectionOptions_Override(c ConnectionOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@flowconsole/sdk.ConnectionOptions",
		nil, // no parameters
		c,
	)
}

func (j *jsiiProxy_ConnectionOptions)SetDetail(val *string) {
	_jsii_.Set(
		j,
		"detail",
		val,
	)
}

func (j *jsiiProxy_ConnectionOptions)SetIcon(val *string) {
	_jsii_.Set(
		j,
		"icon",
		val,
	)
}

func (j *jsiiProxy_ConnectionOptions)SetKind(val *string) {
	_jsii_.Set(
		j,
		"kind",
		val,
	)
}

func (j *jsiiProxy_ConnectionOptions)SetMuted(val *bool) {
	_jsii_.Set(
		j,
		"muted",
		val,
	)
}

