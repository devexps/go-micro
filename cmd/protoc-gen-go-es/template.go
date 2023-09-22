package main

import (
	"bytes"
	_ "embed"
	"text/template"
)

//go:embed esTemplate.tpl
var esTemplate string

type EsType int
type QueryType int
type EsMapType int

const (
	EsTypeNone               EsType = iota
	EsMatchPhrasePrefix      EsType = 1
	EsMatchPhrasePrefixLeft  EsType = 2
	EsMatchPhrasePrefixRight EsType = 3
	EsMatch                  EsType = 4
	EsMatchPrefix            EsType = 5
	RepeatedMessage          EsType = 6
	Message                  EsType = 7
	Date                     EsType = 8
	Boolean                  EsType = 9
	Int32                    EsType = 10
	Int64                    EsType = 11
	Float                    EsType = 12
	Double                   EsType = 13
	KeyWord                  EsType = 14
)

const (
	QueryTypeNone          QueryType = iota
	MatchPhrasePrefix      QueryType = 1
	MatchPhrasePrefixLeft  QueryType = 2
	MatchPhrasePrefixRight QueryType = 3
	Wildcard               QueryType = 4
	WildcardLeft           QueryType = 5
	WildcardRight          QueryType = 6
	Term                   QueryType = 7
	Terms                  QueryType = 8
	Match                  QueryType = 9
	MatchPrefix            QueryType = 10
	NotTerm                QueryType = 11
	NotTerms               QueryType = 12
	TmstampGte             QueryType = 13
	TmstampLte             QueryType = 14
	TmstampGt              QueryType = 15
	TmstampLt              QueryType = 16
	TermsZero              QueryType = 17
)

const (
	EsMapTypeNone  EsMapType = iota
	Struct         EsMapType = 1
	RepeatedStruct EsMapType = 2
	Timestamp      EsMapType = 3
	SimpleType     EsMapType = 4
	EsMapTypeDate  EsMapType = 5
)

type FieldDescription struct {
	TagName, TypeName, FieldName string
	EsType                       EsType
}

type QueryDescription struct {
	TagName, TypeName, VariableName string
	CheckNilFunc                    string
	QueryType                       QueryType
}

type esMapping struct {
	FieldDescriptions []*FieldDescription
}

type buildQuery struct {
	QueryDescriptions []*QueryDescription
	InitRangeQuery    bool
}

type getEsMap struct {
	GetEsMapDescriptions []*GetEsMapDescription
}

type GetEsMapDescription struct {
	TagName, VariableName, FieldName, CheckNilFunc string
	EsMapType                                      EsMapType
}

type messageDescription struct {
	EsMapping   *esMapping
	BuildQuery  *buildQuery
	GetEsMap    *getEsMap
	MessageName string
}

type esWrapper struct {
	MessageDescriptions []*messageDescription
	MapPackageName      map[string]string
}

func (e *esWrapper) execute() string {
	buf := new(bytes.Buffer)
	tmpl, err := template.New("es").Parse(esTemplate)
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(buf, e); err != nil {
		panic(err)
	}
	return buf.String()
}
