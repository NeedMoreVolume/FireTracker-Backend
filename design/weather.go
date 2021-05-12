package design

import . "goa.design/goa/v3/dsl"

var _ = Service("weather", func() {
	Description("The fire service creates new weather datas for fires, and gets/lists weather data")

	Error("bad_request", ErrorResult, "Please ensure you are sending valid data.")
	Error("not_found", ErrorResult, "Sorry, we can't find that right now.")
	Error("unauthorized", ErrorResult, "Unauthorized.")
	Error("forbidden", ErrorResult, "Forbidden.")

	HTTP(func() {
		Path("/weather")
	})

	Method("create", func() {
		Description("Create a weather observation and optional payloads")

		Payload(Weather)

		Result(Weather)
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
		Description("Get a specific piece of weather data")

		Payload(Weather)

		Result(Weather)
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
		Description("Update something about a weather observation specifically")

		Payload(Weather)

		Result(Weather)
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
		Description("Delete a weather observation")

		Payload(Weather)

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
		Description("List weather")

		Payload(WeatherListPayload)

		Result(WeatherList)
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

var Weather = Type("Weather", func() {
	Field(1, "id", Int, "id", func() {
		Example(1)
	})
	Field(2, "createdAt", String, "name", func() {
		Format(FormatDateTime)
		Example("2020-01-02T00:00:00Z")
	})
	Field(3, "fireID", Int, "ID of the fire", func() {
		Example(1)
	})
	Field(4, "logID", Int, "ID of the log", func() {
		Example(1)
	})
	Field(5, "temperature", Temperature, "temperature")
	Field(6, "humidity", Int32, "humidity level", func() {
		Example(1500)
	})
	Field(7, "dewPoint", Temperature, "dew point")
	Field(8, "wind", Wind, "wind data")
	Field(9, "weatherType", String, "type of weather", func() {
		Enum("Sunny", "Cloudy", "Raining", "Windy")
	})
	Field(10, "weatherTime", String, func() {
		Format(FormatDateTime)
		Example("2020-01-02T00:00:00Z")
	})
})

var Wind = Type("Wind", func() {
	Field(1, "speed", Int32, "wind speed", func() {
		Example(15)
	})
	Field(2, "direction", String, "wind direction", func() {
		Enum("S", "SE", "E", "NE", "N", "NW", "W", "SW")
	})
	Field(3, "unit", String, "measurement unit", func() {
		Enum("KPH", "MPH")
	})
})

var Temperature = Type("Temperature", func() {
	Field(1, "unit", String, "measurement unit", func() {
		Enum("K", "C", "F")
	})
	Field(2, "value", Int32, "temperature value", func() {
		Example(20)
	})
})

var WeatherFilters = Type("WeatherFilters", func() {
	Description("list weather payload for weather")
	Field(1, "time", ArrayOf(TimeFilter))
	Field(2, "windSpeed", ArrayOf(NumberFilter))
	Field(3, "temperature", ArrayOf(NumberFilter))
	Field(4, "humidity", ArrayOf(NumberFilter))
})

var WeatherSort = Type("WeatherSorts", func() {
	Description("list sort payload for products, with each field representing a valid sort key")
	Field(1, "id", String, func() {
		Enum("ASC, DESC")
	})
})

var WeatherSearch = Type("WeatherSearch", func() {
	Description("list search payload for products, with each field representing a valid search key")
	Field(1, "name", String, func() {
		Example("Windy Part")
	})
	Field(2, "description", String, func() {
		Example("Almost blew the fire out!")
	})
})

var WeatherPagination = Type("WeatherPagination", func() {
	Description("list pagination for products")
	Field(1, "page", Int, func() {
		Example(1)
	})
	Field(2, "limit", Int, func() {
		Example(10)
	})
	Required("page", "limit")
})

var WeatherListPayload = Type("WeatherListPayload", func() {
	Description("List payload for product list")
	Field(1, "filters", WeatherFilters, "product filters to apply")
	Field(2, "search", WeatherSearch, "product search to apply")
	Field(3, "sort", WeatherSort, "product sort to apply")
	Field(4, "pagination", WeatherPagination, "product pagination to apply")
	Required("filters", "search", "sort", "pagination")
})

var WeatherList = Type("WeatherList", func() {
	Description("List of weathers with pagin data")
	Field(1, "weathers", ArrayOf(Weather), "weather results")
	Field(2, "pagination", Pagination, "pagination info")
})
