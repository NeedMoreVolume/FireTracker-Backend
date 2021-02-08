package design

import 	. "goa.design/goa/v3/dsl"

var _ = Service("log", func() {
	Description("The fire service creates new fires, updates data for a fire, deletes fires, and gets/lists fires")

	Error("bad_request", ErrorResult, "Please ensure you are sending valid data.")
	Error("not_found", ErrorResult, "Sorry, we can't find that right now.")
	Error("unauthorized", ErrorResult, "Unauthorized.")
	Error("forbidden", ErrorResult, "Forbidden.")

	HTTP(func() {
		Path("/log")
	})

	Method("create", func() {
		Description("Add a log to existing fire")

		Payload(Log)

		Result(Log)
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
		Description("Get log and data friends")

		Payload(Log)

		Result(Log)
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
		Description("Update something about a log specifically")

		Payload(Log)

		Result(Log)
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
		Description("Delete some log specifically")

		Payload(Log)

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

		Payload(LogListPayload)

		Result(LogList)
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

	Files("/openapi.json", "./gen/http/openapi.json")
})

var LogFilters = Type("LogFilters", func() {
	Description("list filter payload for fires, with each field representing a valid filter key")
	Field(1, "name", ArrayOf(StringFilter))
	Field(2, "start", ArrayOf(TimeFilter))
	Field(3, "end", ArrayOf(TimeFilter))
})

var LogSort = Type("LogSorts", func() {
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

var LogSearch = Type("LogSearch", func() {
	Description("list search payload for products, with each field representing a valid search key")
	Field(1, "name", String)
	Field(2, "description", String)
})

var LogPagination = Type("LogPagination", func() {
	Description("list pagination for products")
	Field(1, "page", Int)
	Field(2, "limit", Int)
	Required("page", "limit")
})

var LogListPayload = Type("LogListPayload", func() {
	Description("List payload for product list")
	Field(1, "filters", LogFilters, "product filters to apply")
	Field(2, "search", LogSearch, "product search to apply")
	Field(3, "sort", LogSort, "product sort to apply")
	Field(4, "pagination", LogPagination, "product pagination to apply")
	Required("filters", "search", "sort", "pagination")
})

var Log = Type("Log", func() {
	Field(1, "id", Int, "id")
	Field(2, "createdAt", String, "name", func() {
		Format(FormatDateTime)
	})
	Field(3, "updatedAt", String, "name", func() {
		Format(FormatDateTime)
	})
	Field(5, "name", String, "name")
	Field(6, "size", String, "size of log", func() {
		Enum("S", "M", "L", "XL")
	})
	Field(7, "fireID", Int, "Fire identifier log belongs to")
	Field(8, "addedAt", String, "time log was added to fire", func() {
		Format(FormatDateTime)
	})
	// optional weather update when adding a log
	Field(9, "weather", Weather, "weather data at time log was added to fire")
})

var LogList = Type("LogList", func() {
	Field(1, "logs", ArrayOf(Log), "logs")
	Field(2, "pagination", Pagination, "pagination results")
})
