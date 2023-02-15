// k8s
package k8s

import (
	_init_ "example.com/cdk8s-Practice/imports/k8s/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"example.com/cdk8s-Practice/imports/k8s/internal"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// RuntimeClass defines a class of container runtime supported in the cluster.
//
// The RuntimeClass is used to determine which container runtime is used to run all containers in a pod. RuntimeClasses are manually defined by a user or cluster provisioner, and referenced in the PodSpec. The Kubelet is responsible for resolving the RuntimeClassName reference before running the pod.  For more details, see https://kubernetes.io/docs/concepts/containers/runtime-class/
type KubeRuntimeClass interface {
	cdk8s.ApiObject
	// The group portion of the API version (e.g. `authorization.k8s.io`).
	ApiGroup() *string
	// The object's API version (e.g. `authorization.k8s.io/v1`).
	ApiVersion() *string
	// The chart in which this object is defined.
	Chart() cdk8s.Chart
	// The object kind.
	Kind() *string
	// Metadata associated with this API object.
	Metadata() cdk8s.ApiObjectMetadataDefinition
	// The name of the API object.
	//
	// If a name is specified in `metadata.name` this will be the name returned.
	// Otherwise, a name will be generated by calling
	// `Chart.of(this).generatedObjectName(this)`, which by default uses the
	// construct path to generate a DNS-compatible name for the resource.
	Name() *string
	// The tree node.
	Node() constructs.Node
	// Create a dependency between this ApiObject and other constructs.
	//
	// These can be other ApiObjects, Charts, or custom.
	AddDependency(dependencies ...constructs.IConstruct)
	// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
	//
	// Example:
	//     kubePod.addJsonPatch(JsonPatch.replace('/spec/enableServiceLinks', true));
	//
	AddJsonPatch(ops ...cdk8s.JsonPatch)
	// Renders the object to Kubernetes JSON.
	ToJson() interface{}
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for KubeRuntimeClass
type jsiiProxy_KubeRuntimeClass struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_KubeRuntimeClass) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeRuntimeClass) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeRuntimeClass) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeRuntimeClass) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeRuntimeClass) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeRuntimeClass) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubeRuntimeClass) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Defines a "io.k8s.api.node.v1.RuntimeClass" API object.
func NewKubeRuntimeClass(scope constructs.Construct, id *string, props *KubeRuntimeClassProps) KubeRuntimeClass {
	_init_.Initialize()

	if err := validateNewKubeRuntimeClassParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubeRuntimeClass{}

	_jsii_.Create(
		"k8s.KubeRuntimeClass",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "io.k8s.api.node.v1.RuntimeClass" API object.
func NewKubeRuntimeClass_Override(k KubeRuntimeClass, scope constructs.Construct, id *string, props *KubeRuntimeClassProps) {
	_init_.Initialize()

	_jsii_.Create(
		"k8s.KubeRuntimeClass",
		[]interface{}{scope, id, props},
		k,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func KubeRuntimeClass_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKubeRuntimeClass_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"k8s.KubeRuntimeClass",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Renders a Kubernetes manifest for "io.k8s.api.node.v1.RuntimeClass".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func KubeRuntimeClass_Manifest(props *KubeRuntimeClassProps) interface{} {
	_init_.Initialize()

	if err := validateKubeRuntimeClass_ManifestParameters(props); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.StaticInvoke(
		"k8s.KubeRuntimeClass",
		"manifest",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns the `ApiObject` named `Resource` which is a child of the given construct.
//
// If `c` is an `ApiObject`, it is returned directly. Throws an
// exception if the construct does not have a child named `Default` _or_ if
// this child is not an `ApiObject`.
func KubeRuntimeClass_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	if err := validateKubeRuntimeClass_OfParameters(c); err != nil {
		panic(err)
	}
	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"k8s.KubeRuntimeClass",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func KubeRuntimeClass_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"k8s.KubeRuntimeClass",
		"GVK",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_KubeRuntimeClass) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		k,
		"addDependency",
		args,
	)
}

func (k *jsiiProxy_KubeRuntimeClass) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		k,
		"addJsonPatch",
		args,
	)
}

func (k *jsiiProxy_KubeRuntimeClass) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubeRuntimeClass) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

