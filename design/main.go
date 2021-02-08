package design

import (
	"fmt"
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

var JWTAuth = JWTSecurity("jwt", func() {
	Description("Secures endpoint by requiring a valid JWT token")
})

var stageHttpHost = "test.tmf.com"
var stageGrpcHost = "test.tmf.com"
var prodHttpHost = "trackmyfire.com"
var prodGrpcHost = "trackmyfire.com"

var _ = API("fireTracker", func() {
	Title("Fire Tracker Service")
	Description("Service for fire tracking")
	cors.Origin("*", func() {
		cors.Headers("X-Authorization", "X-Time", "X-Api-Version",
			"Content-Type", "Origin", "Authorization")
		cors.Methods("GET", "POST", "OPTIONS")
		cors.Expose("Content-Type", "Origin")
		cors.MaxAge(100)
		cors.Credentials()
	})
	Version("1.0")
	TermsOfService("terms")
	Contact(func() {
		Name("support")
		Email("support@firetracker.com")
		URL("firetracker.com")
	}) //contact info
	License(func() {
		Name("GPT-3")
		URL("licence url")
	}) // licensce info
	Docs(func() {
		Description("doc description")
		URL("license URL")
	}) //documentation links
	Server("tracker", func() {
		Host("development", func() {
			URI("http://localhost:8080")
			URI("grpc://localhost:9080")
		})
		Host("stage", func() {
			URI(fmt.Sprintf("https://%s", stageHttpHost))
			URI(fmt.Sprintf("grpc://%s", stageGrpcHost))
		})
		Host("production", func() {
			URI(fmt.Sprintf("https://%s", prodHttpHost))
			URI(fmt.Sprintf("grpc://%s", prodGrpcHost))
		})
	})
})

var StringFilter = Type("StringFilter", func() {
	Field(1, "key", String, "filter key")
	Field(2, "operator", String, "operator value", func() {
		Enum("=", "!=")
	})
	Field(3, "value", String, "filter value")
	Required("key", "operator", "value")
})

var NumberFilter = Type("NumberFilter", func() {
	Field(1, "key", String, "filter key")
	Field(2, "operator", String, "operator value", func() {
		Enum("=", ">", ">=", "<=", "<")
	})
	Field(3, "value", Int, "filter value")
	Required("key", "operator", "value")
})

var TimeFilter = Type("TimeFilter", func() {
	Field(1, "key", String, "filter key", func() {
		Format(FormatDateTime)
	})
	Field(2, "operator", String, "operator value", func() {
		Enum("=", ">", ">=", "<=", "<")
	})
	Field(3, "value", Int, "filter value")
	Required("key", "operator", "value")
})

var Pagination = Type("Pagination", func() {
	Field(1, "total", Int, "count of the things")
	Field(2, "page", Int, "page number")
	Field(3, "limit", Int, "max number of things")
})
