package schemas

import (
	"log"
	"main/services"

	"github.com/graphql-go/graphql"
)

// GenerateCEPSchema generates the generic CEP schema
func GenerateCEPSchema() (graphql.Schema, error) {
	fields := graphql.Fields{
		"address": &graphql.Field{
			Type: graphql.NewObject(
				graphql.ObjectConfig{
					Name: "Address",
					Fields: graphql.Fields{
						"Cep": &graphql.Field{
							Type: graphql.String,
						},
						"Logradouro": &graphql.Field{
							Type: graphql.String,
						},
						"Complemento": &graphql.Field{
							Type: graphql.String,
						},
						"Bairro": &graphql.Field{
							Type: graphql.String,
						},
						"Localidade": &graphql.Field{
							Type: graphql.String,
						},
						"Uf": &graphql.Field{
							Type: graphql.String,
						},
						"Ibge": &graphql.Field{
							Type: graphql.String,
						},
						"Gia": &graphql.Field{
							Type: graphql.String,
						},
						"Ddd": &graphql.Field{
							Type: graphql.String,
						},
						"Siafi": &graphql.Field{
							Type: graphql.String,
						},
					},
				},
			),
			Args: graphql.FieldConfigArgument{
				"cep": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				cep := p.Args["cep"]
				cepStr := cep.(string)
				return services.FetchDataByCep(cepStr)
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
		log.Fatalf("Failed to create new schema, error: %v", err)
	}

	return schema, err
}
