# A2-Back
GraphiQL interface and GraphQL endpoint running at http://localhost:3001/graphql

# Sample graphql query
```
query {
  address(cep:"09780090") {
    Bairro
    Cep
    Complemento
    Ddd
    Gia
    Ibge
    Localidade
    Logradouro
    Siafi
    Uf
  }
}
```