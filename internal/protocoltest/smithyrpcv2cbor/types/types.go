// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

type ClientOptionalDefaults struct {
	Member *int32

	noSmithyDocumentSerde
}

type ComplexNestedErrorData struct {
	Foo *string

	noSmithyDocumentSerde
}

type Defaults struct {
	DefaultBlob []byte

	DefaultBoolean *bool

	DefaultByte *int8

	DefaultDouble *float64

	DefaultEnum TestEnum

	DefaultFloat *float32

	DefaultIntEnum TestIntEnum

	DefaultInteger *int32

	DefaultList []string

	DefaultLong *int64

	DefaultMap map[string]string

	DefaultShort *int16

	DefaultString *string

	DefaultTimestamp *time.Time

	EmptyBlob []byte

	EmptyString *string

	FalseBoolean bool

	ZeroByte int8

	ZeroDouble float64

	ZeroFloat float32

	ZeroInteger int32

	ZeroLong int64

	ZeroShort int16

	noSmithyDocumentSerde
}

type RecursiveShapesInputOutputNested1 struct {
	Foo *string

	Nested *RecursiveShapesInputOutputNested2

	noSmithyDocumentSerde
}

type RecursiveShapesInputOutputNested2 struct {
	Bar *string

	RecursiveMember *RecursiveShapesInputOutputNested1

	noSmithyDocumentSerde
}

type StructureListMember struct {
	A *string

	B *string

	noSmithyDocumentSerde
}

type GreetingStruct struct {
	Hi *string

	noSmithyDocumentSerde
}

// Describes one specific validation failure for an input member.
type ValidationExceptionField struct {

	// A detailed description of the validation failure.
	//
	// This member is required.
	Message *string

	// A JSONPointer expression to the structure member whose value failed to satisfy
	// the modeled constraints.
	//
	// This member is required.
	Path *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
