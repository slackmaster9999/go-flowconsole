// FlowConsole SDK
package flowconsole

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@flowconsole/sdk.BackgroundJob",
		reflect.TypeOf((*BackgroundJob)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "badge", GoGetter: "Badge"},
			_jsii_.MemberProperty{JsiiProperty: "belongsTo", GoGetter: "BelongsTo"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "executesRequest", GoMethod: "ExecutesRequest"},
			_jsii_.MemberMethod{JsiiMethod: "getDataFrom", GoMethod: "GetDataFrom"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "root", GoGetter: "Root"},
			_jsii_.MemberProperty{JsiiProperty: "schedule", GoGetter: "Schedule"},
			_jsii_.MemberMethod{JsiiMethod: "sendsRequest", GoMethod: "SendsRequest"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "then", GoMethod: "Then"},
			_jsii_.MemberProperty{JsiiProperty: "tone", GoGetter: "Tone"},
		},
		func() interface{} {
			j := jsiiProxy_BackgroundJob{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Component)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@flowconsole/sdk.Component",
		reflect.TypeOf((*Component)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "badge", GoGetter: "Badge"},
			_jsii_.MemberProperty{JsiiProperty: "belongsTo", GoGetter: "BelongsTo"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "executesRequest", GoMethod: "ExecutesRequest"},
			_jsii_.MemberMethod{JsiiMethod: "getDataFrom", GoMethod: "GetDataFrom"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "root", GoGetter: "Root"},
			_jsii_.MemberMethod{JsiiMethod: "sendsRequest", GoMethod: "SendsRequest"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "then", GoMethod: "Then"},
			_jsii_.MemberProperty{JsiiProperty: "tone", GoGetter: "Tone"},
		},
		func() interface{} {
			return &jsiiProxy_Component{}
		},
	)
	_jsii_.RegisterStruct(
		"@flowconsole/sdk.ComponentArgs",
		reflect.TypeOf((*ComponentArgs)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@flowconsole/sdk.ComputerSystem",
		reflect.TypeOf((*ComputerSystem)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "badge", GoGetter: "Badge"},
			_jsii_.MemberProperty{JsiiProperty: "belongsTo", GoGetter: "BelongsTo"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "domain", GoGetter: "Domain"},
			_jsii_.MemberMethod{JsiiMethod: "executesRequest", GoMethod: "ExecutesRequest"},
			_jsii_.MemberMethod{JsiiMethod: "getDataFrom", GoMethod: "GetDataFrom"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "root", GoGetter: "Root"},
			_jsii_.MemberMethod{JsiiMethod: "sendsRequest", GoMethod: "SendsRequest"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "then", GoMethod: "Then"},
			_jsii_.MemberProperty{JsiiProperty: "tone", GoGetter: "Tone"},
		},
		func() interface{} {
			j := jsiiProxy_ComputerSystem{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Component)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@flowconsole/sdk.ConnectionOptions",
		reflect.TypeOf((*ConnectionOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "detail", GoGetter: "Detail"},
			_jsii_.MemberProperty{JsiiProperty: "icon", GoGetter: "Icon"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "muted", GoGetter: "Muted"},
		},
		func() interface{} {
			return &jsiiProxy_ConnectionOptions{}
		},
	)
	_jsii_.RegisterClass(
		"@flowconsole/sdk.Container",
		reflect.TypeOf((*Container)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "badge", GoGetter: "Badge"},
			_jsii_.MemberProperty{JsiiProperty: "belongsTo", GoGetter: "BelongsTo"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "executesRequest", GoMethod: "ExecutesRequest"},
			_jsii_.MemberMethod{JsiiMethod: "getDataFrom", GoMethod: "GetDataFrom"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "root", GoGetter: "Root"},
			_jsii_.MemberMethod{JsiiMethod: "sendsRequest", GoMethod: "SendsRequest"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "technology", GoGetter: "Technology"},
			_jsii_.MemberMethod{JsiiMethod: "then", GoMethod: "Then"},
			_jsii_.MemberProperty{JsiiProperty: "tone", GoGetter: "Tone"},
		},
		func() interface{} {
			j := jsiiProxy_Container{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Component)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@flowconsole/sdk.ExternalService",
		reflect.TypeOf((*ExternalService)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "badge", GoGetter: "Badge"},
			_jsii_.MemberProperty{JsiiProperty: "belongsTo", GoGetter: "BelongsTo"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "executesRequest", GoMethod: "ExecutesRequest"},
			_jsii_.MemberMethod{JsiiMethod: "getDataFrom", GoMethod: "GetDataFrom"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "root", GoGetter: "Root"},
			_jsii_.MemberMethod{JsiiMethod: "sendsRequest", GoMethod: "SendsRequest"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "then", GoMethod: "Then"},
			_jsii_.MemberProperty{JsiiProperty: "tone", GoGetter: "Tone"},
			_jsii_.MemberProperty{JsiiProperty: "vendor", GoGetter: "Vendor"},
		},
		func() interface{} {
			j := jsiiProxy_ExternalService{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Component)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@flowconsole/sdk.KafkaTopic",
		reflect.TypeOf((*KafkaTopic)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "badge", GoGetter: "Badge"},
			_jsii_.MemberProperty{JsiiProperty: "belongsTo", GoGetter: "BelongsTo"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "executesRequest", GoMethod: "ExecutesRequest"},
			_jsii_.MemberMethod{JsiiMethod: "getDataFrom", GoMethod: "GetDataFrom"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "partitionCount", GoGetter: "PartitionCount"},
			_jsii_.MemberProperty{JsiiProperty: "root", GoGetter: "Root"},
			_jsii_.MemberMethod{JsiiMethod: "sendsRequest", GoMethod: "SendsRequest"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "then", GoMethod: "Then"},
			_jsii_.MemberProperty{JsiiProperty: "tone", GoGetter: "Tone"},
		},
		func() interface{} {
			j := jsiiProxy_KafkaTopic{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Component)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@flowconsole/sdk.MessageQueue",
		reflect.TypeOf((*MessageQueue)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "badge", GoGetter: "Badge"},
			_jsii_.MemberProperty{JsiiProperty: "belongsTo", GoGetter: "BelongsTo"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "executesRequest", GoMethod: "ExecutesRequest"},
			_jsii_.MemberMethod{JsiiMethod: "getDataFrom", GoMethod: "GetDataFrom"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "root", GoGetter: "Root"},
			_jsii_.MemberMethod{JsiiMethod: "sendsRequest", GoMethod: "SendsRequest"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "then", GoMethod: "Then"},
			_jsii_.MemberProperty{JsiiProperty: "throughput", GoGetter: "Throughput"},
			_jsii_.MemberProperty{JsiiProperty: "tone", GoGetter: "Tone"},
		},
		func() interface{} {
			j := jsiiProxy_MessageQueue{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Component)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@flowconsole/sdk.Postgres",
		reflect.TypeOf((*Postgres)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "badge", GoGetter: "Badge"},
			_jsii_.MemberProperty{JsiiProperty: "belongsTo", GoGetter: "BelongsTo"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "executesRequest", GoMethod: "ExecutesRequest"},
			_jsii_.MemberMethod{JsiiMethod: "getDataFrom", GoMethod: "GetDataFrom"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "root", GoGetter: "Root"},
			_jsii_.MemberProperty{JsiiProperty: "schema", GoGetter: "Schema"},
			_jsii_.MemberMethod{JsiiMethod: "sendsRequest", GoMethod: "SendsRequest"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "then", GoMethod: "Then"},
			_jsii_.MemberProperty{JsiiProperty: "tone", GoGetter: "Tone"},
		},
		func() interface{} {
			j := jsiiProxy_Postgres{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Component)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@flowconsole/sdk.ReactApp",
		reflect.TypeOf((*ReactApp)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "badge", GoGetter: "Badge"},
			_jsii_.MemberProperty{JsiiProperty: "belongsTo", GoGetter: "BelongsTo"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "executesRequest", GoMethod: "ExecutesRequest"},
			_jsii_.MemberProperty{JsiiProperty: "framework", GoGetter: "Framework"},
			_jsii_.MemberMethod{JsiiMethod: "getDataFrom", GoMethod: "GetDataFrom"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "root", GoGetter: "Root"},
			_jsii_.MemberMethod{JsiiMethod: "sendsRequest", GoMethod: "SendsRequest"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "then", GoMethod: "Then"},
			_jsii_.MemberProperty{JsiiProperty: "tone", GoGetter: "Tone"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
		},
		func() interface{} {
			j := jsiiProxy_ReactApp{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Component)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@flowconsole/sdk.Redis",
		reflect.TypeOf((*Redis)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "badge", GoGetter: "Badge"},
			_jsii_.MemberProperty{JsiiProperty: "belongsTo", GoGetter: "BelongsTo"},
			_jsii_.MemberProperty{JsiiProperty: "cluster", GoGetter: "Cluster"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "executesRequest", GoMethod: "ExecutesRequest"},
			_jsii_.MemberMethod{JsiiMethod: "getDataFrom", GoMethod: "GetDataFrom"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "root", GoGetter: "Root"},
			_jsii_.MemberMethod{JsiiMethod: "sendsRequest", GoMethod: "SendsRequest"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "then", GoMethod: "Then"},
			_jsii_.MemberProperty{JsiiProperty: "tone", GoGetter: "Tone"},
		},
		func() interface{} {
			j := jsiiProxy_Redis{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Component)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@flowconsole/sdk.RestApi",
		reflect.TypeOf((*RestApi)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "badge", GoGetter: "Badge"},
			_jsii_.MemberProperty{JsiiProperty: "belongsTo", GoGetter: "BelongsTo"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "endpoint", GoGetter: "Endpoint"},
			_jsii_.MemberMethod{JsiiMethod: "executesRequest", GoMethod: "ExecutesRequest"},
			_jsii_.MemberMethod{JsiiMethod: "getDataFrom", GoMethod: "GetDataFrom"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "method", GoGetter: "Method"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "root", GoGetter: "Root"},
			_jsii_.MemberMethod{JsiiMethod: "sendsRequest", GoMethod: "SendsRequest"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "then", GoMethod: "Then"},
			_jsii_.MemberProperty{JsiiProperty: "tone", GoGetter: "Tone"},
		},
		func() interface{} {
			j := jsiiProxy_RestApi{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Component)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@flowconsole/sdk.User",
		reflect.TypeOf((*User)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "badge", GoGetter: "Badge"},
			_jsii_.MemberProperty{JsiiProperty: "belongsTo", GoGetter: "BelongsTo"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "executesRequest", GoMethod: "ExecutesRequest"},
			_jsii_.MemberMethod{JsiiMethod: "getDataFrom", GoMethod: "GetDataFrom"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "root", GoGetter: "Root"},
			_jsii_.MemberMethod{JsiiMethod: "sendsRequest", GoMethod: "SendsRequest"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "then", GoMethod: "Then"},
			_jsii_.MemberProperty{JsiiProperty: "tone", GoGetter: "Tone"},
		},
		func() interface{} {
			j := jsiiProxy_User{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Component)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@flowconsole/sdk.UserArgs",
		reflect.TypeOf((*UserArgs)(nil)).Elem(),
	)
}
