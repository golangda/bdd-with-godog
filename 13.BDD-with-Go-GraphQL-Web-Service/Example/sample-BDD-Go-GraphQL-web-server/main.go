package main

import (
    
    "net/http"
    "github.com/graphql-go/graphql"
    "github.com/graphql-go/handler"
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

var users []User
var userID = 1

var userType = graphql.NewObject(graphql.ObjectConfig{
    Name: "User",
    Fields: graphql.Fields{
        "id":    &graphql.Field{Type: graphql.Int},
        "name":  &graphql.Field{Type: graphql.String},
        "email": &graphql.Field{Type: graphql.String},
    },
})

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
    Name: "Query",
    Fields: graphql.Fields{
        "users": &graphql.Field{
            Type: graphql.NewList(userType),
            Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                return users, nil
            },
        },
    },
})

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
    Name: "Mutation",
    Fields: graphql.Fields{
        "createUser": &graphql.Field{
            Type: userType,
            Args: graphql.FieldConfigArgument{
                "name":  &graphql.ArgumentConfig{Type: graphql.String},
                "email": &graphql.ArgumentConfig{Type: graphql.String},
            },
            Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                user := User{
                    ID:    userID,
                    Name:  p.Args["name"].(string),
                    Email: p.Args["email"].(string),
                }
                users = append(users, user)
                userID++
                return user, nil
            },
        },
    },
})

func main() {
    schema, _ := graphql.NewSchema(graphql.SchemaConfig{
        Query:    rootQuery,
        Mutation: rootMutation,
    })

    h := handler.New(&handler.Config{
        Schema: &schema,
        Pretty: true,
    })

    http.Handle("/graphql", h)
    http.ListenAndServe(":8080", nil)
}