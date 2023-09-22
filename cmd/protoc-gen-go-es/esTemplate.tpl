{{$helperPkg := index .MapPackageName "helper"}}
{{$elasticPkg := index .MapPackageName "elastic"}}
{{$timePkg := index .MapPackageName "time"}}
{{$logPkg := index .MapPackageName "log"}}
{{$stringsPkg := index .MapPackageName "strings"}}
{{$ptypesPkg := index .MapPackageName "ptypes"}}
{{- range .MessageDescriptions}}
    {{- if ne .EsMapping nil }}
func (this *{{.MessageName}}) EsMapping(mappingEs map[string]interface{}) {
    {{- range .EsMapping.FieldDescriptions}}
        {{- if eq .EsType 1}}
    mappingEs["{{.TagName}}"] = map[string]interface{}{"type": "keyword",
        "fields": map[string]interface{}{
        "search":         map[string]string{"type": "text", "analyzer": "search", "search_analyzer": "search"},
        "search_reverse": map[string]string{"type": "text", "analyzer": "search_reverse", "search_analyzer": "search_reverse"}}}
        {{- else if eq .EsType 2}}
    mappingEs["{{.TagName}}"] = map[string]interface{}{"type": "keyword",
        "fields": map[string]interface{}{"search": map[string]string{"type": "text", "analyzer": "search", "search_analyzer": "search"}}}
        {{- else if eq .EsType 3}}
    mappingEs["{{.TagName}}"] = map[string]interface{}{"type": "keyword",
        "fields": map[string]interface{}{"search_reverse": map[string]string{"type": "text", "analyzer": "search_reverse", "search_analyzer": "search_reverse"}}}
        {{- else if eq .EsType 4}}
    mappingEs["{{.TagName}}"] = map[string]string{"type": "text"}
        {{- else if eq .EsType 5}}
    mappingEs["{{.TagName}}"] = map[string]interface{}{"type": "keyword", "fields": map[string]interface{}{
        search": map[string]string{"type": "text"}}}
        {{- else if eq .EsType 6}}
    if !{{$helperPkg}}IsNil(this.{{.FieldName}}){
        tmpMapingEs := map[string]interface{}{}
        {{.TypeName}} := &{{.TypeName}}{}
        this.{{.FieldName}} = append(this.{{.FieldName}}, {{.TypeName}})
        this.Get{{.FieldName}}()[0].EsMapping(tmpMapingEs)
        for key := range tmpMapingEs {
            mappingEs["{{.TagName}}."+key] = tmpMapingEs[key]
        }
    }
        {{- else if eq .EsType 7}}
    this.Get{{.FieldName}}().EsMapping({{$helperPkg}}MakeKeyEsMap({{$helperPkg}}MakeKeyEsMap(mappingEs, "{{.TagName}}"), "properties"))
        {{- else if eq .EsType 8}}
    mappingEs["{{.TagName}}"] = map[string]string{"type": "date"}
        {{- else if eq .EsType 9}}
    mappingEs["{{.TagName}}"] = map[string]string{"type": "boolean"}
        {{- else if eq .EsType 10}}
    mappingEs["{{.TagName}}"] = map[string]string{"type": "integer"}
        {{- else if eq .EsType 11}}
    mappingEs["{{.TagName}}"] = map[string]string{"type": "long"}
        {{- else if eq .EsType 12}}
    mappingEs["{{.TagName}}"] = map[string]string{"type": "float"}
        {{- else if eq .EsType 13}}
    mappingEs["{{.TagName}}"] = map[string]string{"type": "double"}
        {{- else if eq .EsType 14}}
    mappingEs["{{.TagName}}"] = map[string]string{"type": "keyword"}
        {{- end}}
    {{- end}}
}
    {{- end}}
    {{- if ne .BuildQuery nil}}
{{$initRangeQuery := .BuildQuery.InitRangeQuery}}
func (this *{{.MessageName}}) BuildQuery(query *{{$elasticPkg}}BoolQuery) *{{$elasticPkg}}BoolQuery {
    if query == nil {
        query = {{$elasticPkg}}NewBoolQuery()
    }
    rangeTmstampSearch := &{{$helperPkg}}MapRangeSearch{MapRangeTmStampSearch: map[string]*{{$helperPkg}}RangeTmstampSearch{}}
    disableRangeFilter, searchPhone := false, false
    {{- if $initRangeQuery}}
    r := &{{$helperPkg}}RangeQuery{
        MapQuery: map[string]*{{$elasticPkg}}RangeQuery{},
    }
    {{- end}}
    {{- range .BuildQuery.QueryDescriptions}}
        {{- if eq .QueryType 17}}
    if {{$helperPkg}}IsZero({{.VariableName}}) {
        query = query.MustNot({{$elasticPkg}}NewExistsQuery("{{.TagName}}"))
    } else {
        if !{{$helperPkg}}{{.CheckNilFunc}}({{.VariableName}}) && {{$helperPkg}}IsEnumAll({{.VariableName}}) {
            query = query.Must({{$elasticPkg}}NewTermQuery("{{.TagName}}", {{.VariableName}}))
        }
    }
        {{- else}}
    if !{{$helperPkg}}{{.CheckNilFunc}}({{.VariableName}}){
            {{- if eq .QueryType 1}}
        fields := {{$stringsPkg}}Split("{{.TagName}}", ";")
    	if len(fields) < 2 {
    	    if !disableRangeFilter && len(fmt.Sprintf("%v",{{.VariableName}})) >= 8 {
    	        disableRangeFilter = true
    	    }
    	    if "{{.TagName}}" == "userInfo.phoneNumber" {
    	        searchPhone = true
    	    }
    	    query = query.Must({{$elasticPkg}}NewMultiMatchQuery(
                {{.VariableName}}, "{{.TagName}}.search", "{{.TagName}}.search_reverse").MaxExpansions(1024).Slop(2).Type("phrase_prefix"))
    	} else {
    	    fieldsSearch := make([]string, 2*len(fields))
    	    for i, field := range fields {
    	        fieldsSearch[2*i] = field+ ".search"
    	        fieldsSearch[2*i+1] = field+".search_reverse"
    	    }
    	    query = query.Must({{$elasticPkg}}NewMultiMatchQuery(
                {{.VariableName}}, fieldsSearch...).MaxExpansions(1024).Slop(2).Type("phrase_prefix"))
    	}
            {{- else if eq .QueryType 2}}
        query = query.Must({{$elasticPkg}}NewMatchPhrasePrefixQuery(
            fmt.Sprintf("%s.search", "{{.TagName}}"),
        	{{.VariableName}}).MaxExpansions(1024).Slop(2))
            {{- else if eq .QueryType 3}}
        query = query.Must({{$elasticPkg}}NewMatchPhrasePrefixQuery(
            fmt.Sprintf("%s.search_reverse", "{{.TagName}}"),
            {{.VariableName}}).MaxExpansions(1024).Slop(2))
            {{- else if eq .QueryType 4}}
        s := fmt.Sprintf("%v", {{.VariableName}})
        if !{{$stringsPkg}}Contains(s, "*") {
            s = "*" + s + "*"
        }
        query = query.Must({{$elasticPkg}}NewWildcardQuery("{{.TagName}}", s))
            {{- else if eq .QueryType 5}}
        query = query.Must({{$elasticPkg}}NewWildcardQuery("{{.TagName}}", fmt.Sprintf("*%v", {{.VariableName}})))
            {{- else if eq .QueryType 6}}
        query = query.Must({{$elasticPkg}}NewWildcardQuery("{{.TagName}}", fmt.Sprintf("%v*", {{.VariableName}})))
            {{- else if eq .QueryType 7}}
        if !{{$helperPkg}}IsEnumAll({{.VariableName}}) {
        	query = query.Filter({{$elasticPkg}}NewTermQuery("{{.TagName}}", {{.VariableName}}))
        }
            {{- else if eq .QueryType 8}}
        query = query.Filter({{$elasticPkg}}NewTermsQuery("{{.TagName}}", {{$helperPkg}}DoubleSlice({{.VariableName}})...))
            {{- else if eq .QueryType 9}}
        query = query.Must({{$elasticPkg}}NewMatchQuery("{{.TagName}}", {{.VariableName}}))
            {{- else if eq .QueryType 10}}
        query = query.Must({{$elasticPkg}}NewMatchQuery("{{.TagName}}.search", {{.VariableName}}).MinimumShouldMatch("3<90%"))
            {{- else if eq .QueryType 11}}
        query = query.MustNot({{$elasticPkg}}NewTermQuery("{{.TagName}}", {{.VariableName}}))
            {{- else if eq .QueryType 12}}
        query = query.MustNot({{$elasticPkg}}NewTermsQuery("{{.TagName}}", {{$helperPkg}}DoubleSlice({{.VariableName}})...))
            {{- else if eq .QueryType 13}}
        if !rangeTmstampSearch.AddFrom("{{.TagName}}", {{.VariableName}}, true) {
            query = query.Must(r.NewRangeQuery("{{.TagName}}").Gte({{.VariableName}}))
        }
            {{- else if eq .QueryType 14}}
        if !rangeTmstampSearch.AddTo("{{.TagName}}", {{.VariableName}}, true) {
            query = query.Must(r.NewRangeQuery("{{.TagName}}").Lte({{.VariableName}}))
        }
            {{- else if eq .QueryType 15}}
        if !rangeTmstampSearch.AddFrom("{{.TagName}}", {{.VariableName}}, false) {
            query = query.Must(r.NewRangeQuery("{{.TagName}}").Gt({{.VariableName}}))
        }
            {{- else if eq .QueryType 16}}
        if !rangeTmstampSearch.AddTo("{{.TagName}}", {{.VariableName}}, false) {
            query = query.Must(r.NewRangeQuery("{{.TagName}}").Lt({{.VariableName}}))
        }
            {{- end}}
    }
        {{- end}}
    {{- end}}

    if !disableRangeFilter || searchPhone {
        for k, v := range rangeTmstampSearch.MapRangeTmStampSearch {
            f, t := v.From, v.To
            if f != 0 || t != 0 {
                rangeQuery := {{$elasticPkg}}NewRangeQuery(k)
				if f != 0 {
					rangeQuery.Gte(f)
				}
				if t != 0 {
					rangeQuery.Lte(t)
				}
				query = query.Filter(rangeQuery)
            } else {
                {{$logPkg}}Info("Invalid range query ", k)
            }
        }
    }
    source, _ := query.Source()
    a, _ := json.Marshal(source)
    {{$logPkg}}Info("query = ", string(a))

    return query
}
    {{- end}}
    {{- if ne .GetEsMap nil}}
func (this *{{.MessageName}}) GetEsMap(esMap *map[string]interface{}) {
    {{- range .GetEsMap.GetEsMapDescriptions}}
        {{- if eq .EsMapType 1}}
    this.Get{{.FieldName}}().GetEsMap({{$helperPkg}}MakeKeyMap(esMap, "{{.TagName}}"))
        {{- else if eq .EsMapType 2}}
    esMapArr := map[string][]interface{}{}
    for _, v := range this.Get{{.FieldName}}() {
        tpmEsMap := map[string]interface{}{}
        v.GetEsMap(&tpmEsMap)
        for key := range tpmEsMap {
            if _, ok := esMapArr[key]; ok {
                esMapArr[key] = append(esMapArr[key], tpmEsMap[key])
            } else {
                esMapArr[key] = []interface{}{tpmEsMap[key]}
            }
        }
    }
    for key := range esMapArr {
        (*esMap)["{{.TagName}}."+key] = esMapArr[key]
    }
        {{- else if eq .EsMapType 3}}
    if !{{$helperPkg}}IsNil({{.VariableName}}){
    	if ts, ok := {{$helperPkg}}CheckTimestampType({{.VariableName}}); ok {
    		if ts != nil {
    		    tm, err := {{$ptypesPkg}}Timestamp(ts)
    		    if err == nil {
    		        (*esMap)["{{.TagName}}"] = tm.UnixNano() / int64({{$timePkg}}Millisecond)
    		    }
    		}
    	}
    }
        {{- else if eq .EsMapType 4}}
    if !{{$helperPkg}}{{.CheckNilFunc}}({{.VariableName}}) {
        (*esMap)["{{.TagName}}"] = {{.VariableName}}
    }
        {{- else if eq .EsMapType 5}}
    if !{{$helperPkg}}IsNil({{.VariableName}}){
        if date, ok := {{$helperPkg}}CheckDateType({{.VariableName}}); ok {
            if date != nil {
                tm := {{.VariableName}}.AsTime()
                (*esMap)["{{.TagName}}"] = tm.UnixNano() / int64({{$timePkg}}Millisecond)
            }
        }
    }
        {{- end}}
    {{- end}}
}
    {{- end}}

{{- end}}