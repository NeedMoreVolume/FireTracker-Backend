package design

import 	. "goa.design/goa/v3/dsl"

var _ = Service("fire", func() {
	Description("The fire service creates new fires, updates data for a fire, deletes fires, and gets/lists fires")

	Error("bad_request", ErrorResult, "Please ensure you are sending valid data.")
	Error("not_found", ErrorResult, "Sorry, we can't find that right now.")
	Error("unauthorized", ErrorResult, "Unauthorized.")
	Error("forbidden", ErrorResult, "Forbidden.")

	HTTP(func() {
		Path("/fire")
	})

	Method("create", func() {
		Description("Create a fire and optional payloads")

		Payload(Fire)

		Result(Fire)
		HTTP(func() {
			POST("/")
			Response(StatusOK)
			Response("bad_request", StatusBadRequest)
			Response(StatusInternalServerError)
		})
		GRPC(func() {
			Response(CodeOK)
			Response("bad_request", CodeInvalidArgument)
			Response(CodeInternal)
		})
	})

	Method("get", func() {
		Description("Get fire and data friends")

		Payload(func() {
			Field(1, "id", Int, "fire id")
		})

		Result(Fire)
		HTTP(func() {
			GET("/{id}")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
			Response(StatusInternalServerError)
		})
		GRPC(func() {
			Response(CodeOK)
			Response("not_found", StatusNotFound)
			Response(StatusInternalServerError)
		})
	})

	Method("update", func() {
		Description("Update something about a fire specifically")

		Payload(Fire)

		Result(Fire)
		HTTP(func() {
			PUT("/{id}")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
			Response(StatusInternalServerError)
		})
		GRPC(func() {
			Response(CodeOK)
			Response("not_found", CodeNotFound)
			Response(CodeInternal)
		})
	})

	Method("delete", func() {
		Description("Update something about a fire specifically")

		Payload(func() {
			Field(1, "id", Int, "fire id")
		})

		HTTP(func() {
			DELETE("/{id}")
			Response(StatusNoContent)
			Response("not_found", StatusNotFound)
			Response(StatusInternalServerError)
		})
		GRPC(func() {
			Response(CodeOK)
			Response("not_found", CodeNotFound)
			Response(CodeInternal)
		})
	})

	Method("list", func() {
		Description("List fires")

		Payload(FireListPayload)

		Result(FireList)
		HTTP(func() {
			POST("/list")
			Response(StatusOK)
			Response(StatusBadRequest)
			Response(StatusInternalServerError)
		})
		GRPC(func() {
			Response(CodeOK)
			Response(CodeInvalidArgument)
			Response(CodeInternal)
		})
	})

	Method("getWeatherForFire", func() {
		Description("Gets a list of weather for a fire")

		Payload(func() {
			Field(1, "id", Int, "fire id")
		})

		Result(WeatherList)
		HTTP(func() {
			GET("/{id}/weather")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
			Response(StatusInternalServerError)
		})
		GRPC(func() {
			Response(CodeOK)
			Response("not_found", CodeNotFound)
			Response(StatusInternalServerError)
		})
	})

	Method("getLogsForFire", func() {
		Description("Gets a list of logs for a fire")

		Payload(func() {
			Field(1, "id", Int, "fire id")
		})

		Result(LogList)
		HTTP(func() {
			GET("/{id}/logs")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
			Response(StatusInternalServerError)
		})
		GRPC(func() {
			Response(CodeOK)
			Response("not_found", CodeNotFound)
			Response(StatusInternalServerError)
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})

var FireFilters = Type("FireFilters", func() {
	Description("list filter payload for fires, with each field representing a valid filter key")
	Field(1, "name", ArrayOf(StringFilter))
	Field(2, "start", ArrayOf(TimeFilter))
	Field(3, "end", ArrayOf(TimeFilter))
})

var FireSort = Type("FireSorts", func() {
	Description("list sort payload for fires, with each field representing a valid sort key")
	Field(1, "id", String, func() {
		Enum("ASC, DESC")
	})
	Field(2, "start", String, func() {
		Enum("ASC, DESC")
	})
	Field(3, "end", String, func() {
		Enum("ASC, DESC")
	})
})

var FireSearch = Type("FireSearch", func() {
	Description("list search payload for products, with each field representing a valid search key")
	Field(1, "name", String)
	Field(2, "description", String)
})

var FirePagination = Type("FirePagination", func() {
	Description("list pagination for products")
	Field(1, "page", Int)
	Field(2, "limit", Int)
	Required("page", "limit")
})

var FireListPayload = Type("FireListPayload", func() {
	Description("List payload for product list")
	Field(1, "filters", FireFilters, "product filters to apply")
	Field(2, "search", FireSearch, "product search to apply")
	Field(3, "sort", FireSort, "product sort to apply")
	Field(4, "pagination", FirePagination, "product pagination to apply")
	Required("filters", "search", "sort", "pagination")
})

var Fire = Type("Fire", func() {
	Field(1, "id", Int, "id")
	Field(2, "createdAt", String, "name", func() {
		Format(FormatDateTime)
	})
	Field(3, "updatedAt", String, "name", func() {
		Format(FormatDateTime)
	})
	Field(4, "deletedAt", String, "name", func() {
		Format(FormatDateTime)
	})
	Field(5, "name", String, "name")
	Field(6, "description", String, "description")
	Field(7, "start", String, "start time of fire", func() {
		Format(FormatDateTime)
	})
	Field(8, "end", String, "end time of fire", func() {
		Format(FormatDateTime)
	})
})

var FireList = Type("FireList", func() {
	Field(1, "fires", ArrayOf(Fire), "list of fires")
	Field(2, "pagination", Pagination, "pagination results")
})
