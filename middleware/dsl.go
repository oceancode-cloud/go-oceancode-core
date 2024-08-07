package _middleware

import (
	"github.com/graphql-go/graphql"
	_response "go-oceancode-core/model/response"
	"go/types"
	"reflect"
	"strings"
)

func ExecuteQuery(query string, schema graphql.Schema) interface{} {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		return _response.ResultError()
	}
	return &result.Data
}

func mappingType(dataType interface{}) graphql.Type {
	isInt := dataType == types.Int || dataType == types.Int8 || dataType == types.Int16 || dataType == types.Int32 || dataType == types.Int64 ||
		dataType == types.Uint || dataType == types.Uint8 || dataType == types.Uint16 || dataType == types.Uint32 || dataType == types.Uint64 ||
		strings.EqualFold(dataType.(string), "int") ||
		strings.EqualFold(dataType.(string), "int8") ||
		strings.EqualFold(dataType.(string), "int16") ||
		strings.EqualFold(dataType.(string), "int32") ||
		strings.EqualFold(dataType.(string), "int64") ||
		strings.EqualFold(dataType.(string), "uint") ||
		strings.EqualFold(dataType.(string), "uint8") ||
		strings.EqualFold(dataType.(string), "uint16") ||
		strings.EqualFold(dataType.(string), "uint32") ||
		strings.EqualFold(dataType.(string), "uint64")

	if isInt {
		return graphql.Int
	}

	isFloat := dataType == types.Float32 || dataType == types.Float64 ||
		strings.EqualFold(dataType.(string), "float32") ||
		strings.EqualFold(dataType.(string), "float64")

	if isFloat {
		return graphql.Float
	}

	if dataType == types.Bool || strings.EqualFold(dataType.(string), "bool") {
		return graphql.Boolean
	}

	return graphql.String
}

func SetQueryField(fields interface{}, name string, inputKey string, input interface{}, result interface{}, resolverHandler func(params graphql.ResolveParams) (interface{}, error)) {
	gFields := fields.(graphql.Fields)
	gFields[name] = BuildQueryField(name, inputKey, input, result, resolverHandler)
}

func InitQuerySchema(handler func(interface{})) (graphql.Schema, error) {
	fields := graphql.Fields{}
	handler(fields)
	rootQuery := graphql.ObjectConfig{Name: "Query", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	return graphql.NewSchema(schemaConfig)
}

func BuildQueryField(name string, inputKey string, input interface{}, result interface{}, resolverHandler func(params graphql.ResolveParams) (interface{}, error)) (field *graphql.Field) {
	fieldsMapping := graphql.Fields{}
	if result != nil {
		values := reflect.ValueOf(result)
		types := reflect.TypeOf(result)

		for i := 0; i < values.NumField(); i++ {
			fieldName := types.Field(i).Name
			fieldType := types.Field(i).Type
			fieldName = strings.ToLower(fieldName[0:1]) + fieldName[1:]

			gType := mappingType(fieldType.Name())
			fieldsMapping[fieldName] = &graphql.Field{
				Type: gType,
			}
		}
	}

	var queryType = graphql.NewObject(
		graphql.ObjectConfig{
			Name:   name,
			Fields: fieldsMapping,
		})

	resolve := func(p graphql.ResolveParams) (interface{}, error) {
		return resolverHandler(p)
	}

	args := graphql.FieldConfigArgument{}

	if input != nil {
		resultValues := reflect.ValueOf(input)
		resultTypes := reflect.TypeOf(input)
		if !strings.EqualFold(inputKey, "") {
			args[inputKey] = &graphql.ArgumentConfig{
				Type: mappingType(input),
			}
		} else {
			for i := 0; i < resultValues.NumField(); i++ {
				fieldName := resultTypes.Field(i).Name
				fieldType := resultTypes.Field(i).Type
				fieldName = strings.ToLower(fieldName[0:1]) + fieldName[1:]
				args[fieldName] = &graphql.ArgumentConfig{
					Type: mappingType(fieldType.Name()),
				}
			}
		}
	}

	return &graphql.Field{
		Type:    queryType,
		Args:    args,
		Resolve: resolve,
	}
}
