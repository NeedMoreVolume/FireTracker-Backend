// Code generated by goa v3.2.5, DO NOT EDIT.
//
// tracker gRPC client CLI support package
//
// Command:
// $ goa gen github.com/NeedMoreVolume/FireTracker/design

package cli

import (
	"flag"
	"fmt"
	"os"

	firec "github.com/NeedMoreVolume/FireTracker/gen/grpc/fire/client"
	logc "github.com/NeedMoreVolume/FireTracker/gen/grpc/log/client"
	weatherc "github.com/NeedMoreVolume/FireTracker/gen/grpc/weather/client"
	goa "goa.design/goa/v3/pkg"
	grpc "google.golang.org/grpc"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `log (create|get|update|delete|list)
fire (create|get|update|delete|list|get-weather-for-fire|get-logs-for-fire)
weather (create|get|update|delete|list)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` log create --message '{
      "addedAt": "1977-05-06T14:52:07Z",
      "createdAt": "1994-05-07T16:07:26Z",
      "fireID": 6736437899940715186,
      "id": 1913927608535220508,
      "name": "Illo ut et excepturi et adipisci.",
      "size": "S",
      "updatedAt": "1981-01-10T06:48:05Z",
      "weather": {
         "createdAt": "2013-03-31T00:00:31Z",
         "dewPoint": {
            "unit": "F",
            "value": 727998182
         },
         "fireID": 5021708757207724489,
         "humidity": 1961871777,
         "id": 4545438651844921453,
         "logID": 1167531932544065094,
         "temperature": {
            "unit": "F",
            "value": 727998182
         },
         "weatherTime": "2003-07-28T19:16:39Z",
         "weatherType": "Windy",
         "wind": {
            "direction": "SE",
            "speed": 740748338,
            "unit": "MPH"
         }
      }
   }'` + "\n" +
		os.Args[0] + ` fire create --message '{
      "createdAt": "2001-07-26T15:42:42Z",
      "deletedAt": "2008-01-23T13:39:15Z",
      "description": "Vero dolorum enim et aut.",
      "end": "2002-03-03T18:26:05Z",
      "id": 5117576016712403101,
      "name": "Aliquid odio asperiores totam dignissimos.",
      "start": "1987-09-07T04:18:11Z",
      "updatedAt": "1985-03-21T02:39:44Z"
   }'` + "\n" +
		os.Args[0] + ` weather create --message '{
      "createdAt": "1997-03-09T21:43:11Z",
      "dewPoint": {
         "unit": "F",
         "value": 727998182
      },
      "fireID": 82313619316836653,
      "humidity": 1148712326,
      "id": 8860917851139109311,
      "logID": 1879967797057239286,
      "temperature": {
         "unit": "F",
         "value": 727998182
      },
      "weatherTime": "1982-04-16T02:42:32Z",
      "weatherType": "Cloudy",
      "wind": {
         "direction": "SE",
         "speed": 740748338,
         "unit": "MPH"
      }
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(cc *grpc.ClientConn, opts ...grpc.CallOption) (goa.Endpoint, interface{}, error) {
	var (
		logFlags = flag.NewFlagSet("log", flag.ContinueOnError)

		logCreateFlags       = flag.NewFlagSet("create", flag.ExitOnError)
		logCreateMessageFlag = logCreateFlags.String("message", "", "")

		logGetFlags       = flag.NewFlagSet("get", flag.ExitOnError)
		logGetMessageFlag = logGetFlags.String("message", "", "")

		logUpdateFlags       = flag.NewFlagSet("update", flag.ExitOnError)
		logUpdateMessageFlag = logUpdateFlags.String("message", "", "")

		logDeleteFlags       = flag.NewFlagSet("delete", flag.ExitOnError)
		logDeleteMessageFlag = logDeleteFlags.String("message", "", "")

		logListFlags       = flag.NewFlagSet("list", flag.ExitOnError)
		logListMessageFlag = logListFlags.String("message", "", "")

		fireFlags = flag.NewFlagSet("fire", flag.ContinueOnError)

		fireCreateFlags       = flag.NewFlagSet("create", flag.ExitOnError)
		fireCreateMessageFlag = fireCreateFlags.String("message", "", "")

		fireGetFlags       = flag.NewFlagSet("get", flag.ExitOnError)
		fireGetMessageFlag = fireGetFlags.String("message", "", "")

		fireUpdateFlags       = flag.NewFlagSet("update", flag.ExitOnError)
		fireUpdateMessageFlag = fireUpdateFlags.String("message", "", "")

		fireDeleteFlags       = flag.NewFlagSet("delete", flag.ExitOnError)
		fireDeleteMessageFlag = fireDeleteFlags.String("message", "", "")

		fireListFlags       = flag.NewFlagSet("list", flag.ExitOnError)
		fireListMessageFlag = fireListFlags.String("message", "", "")

		fireGetWeatherForFireFlags       = flag.NewFlagSet("get-weather-for-fire", flag.ExitOnError)
		fireGetWeatherForFireMessageFlag = fireGetWeatherForFireFlags.String("message", "", "")

		fireGetLogsForFireFlags       = flag.NewFlagSet("get-logs-for-fire", flag.ExitOnError)
		fireGetLogsForFireMessageFlag = fireGetLogsForFireFlags.String("message", "", "")

		weatherFlags = flag.NewFlagSet("weather", flag.ContinueOnError)

		weatherCreateFlags       = flag.NewFlagSet("create", flag.ExitOnError)
		weatherCreateMessageFlag = weatherCreateFlags.String("message", "", "")

		weatherGetFlags       = flag.NewFlagSet("get", flag.ExitOnError)
		weatherGetMessageFlag = weatherGetFlags.String("message", "", "")

		weatherUpdateFlags       = flag.NewFlagSet("update", flag.ExitOnError)
		weatherUpdateMessageFlag = weatherUpdateFlags.String("message", "", "")

		weatherDeleteFlags       = flag.NewFlagSet("delete", flag.ExitOnError)
		weatherDeleteMessageFlag = weatherDeleteFlags.String("message", "", "")

		weatherListFlags       = flag.NewFlagSet("list", flag.ExitOnError)
		weatherListMessageFlag = weatherListFlags.String("message", "", "")
	)
	logFlags.Usage = logUsage
	logCreateFlags.Usage = logCreateUsage
	logGetFlags.Usage = logGetUsage
	logUpdateFlags.Usage = logUpdateUsage
	logDeleteFlags.Usage = logDeleteUsage
	logListFlags.Usage = logListUsage

	fireFlags.Usage = fireUsage
	fireCreateFlags.Usage = fireCreateUsage
	fireGetFlags.Usage = fireGetUsage
	fireUpdateFlags.Usage = fireUpdateUsage
	fireDeleteFlags.Usage = fireDeleteUsage
	fireListFlags.Usage = fireListUsage
	fireGetWeatherForFireFlags.Usage = fireGetWeatherForFireUsage
	fireGetLogsForFireFlags.Usage = fireGetLogsForFireUsage

	weatherFlags.Usage = weatherUsage
	weatherCreateFlags.Usage = weatherCreateUsage
	weatherGetFlags.Usage = weatherGetUsage
	weatherUpdateFlags.Usage = weatherUpdateUsage
	weatherDeleteFlags.Usage = weatherDeleteUsage
	weatherListFlags.Usage = weatherListUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "log":
			svcf = logFlags
		case "fire":
			svcf = fireFlags
		case "weather":
			svcf = weatherFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "log":
			switch epn {
			case "create":
				epf = logCreateFlags

			case "get":
				epf = logGetFlags

			case "update":
				epf = logUpdateFlags

			case "delete":
				epf = logDeleteFlags

			case "list":
				epf = logListFlags

			}

		case "fire":
			switch epn {
			case "create":
				epf = fireCreateFlags

			case "get":
				epf = fireGetFlags

			case "update":
				epf = fireUpdateFlags

			case "delete":
				epf = fireDeleteFlags

			case "list":
				epf = fireListFlags

			case "get-weather-for-fire":
				epf = fireGetWeatherForFireFlags

			case "get-logs-for-fire":
				epf = fireGetLogsForFireFlags

			}

		case "weather":
			switch epn {
			case "create":
				epf = weatherCreateFlags

			case "get":
				epf = weatherGetFlags

			case "update":
				epf = weatherUpdateFlags

			case "delete":
				epf = weatherDeleteFlags

			case "list":
				epf = weatherListFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "log":
			c := logc.NewClient(cc, opts...)
			switch epn {
			case "create":
				endpoint = c.Create()
				data, err = logc.BuildCreatePayload(*logCreateMessageFlag)
			case "get":
				endpoint = c.Get()
				data, err = logc.BuildGetPayload(*logGetMessageFlag)
			case "update":
				endpoint = c.Update()
				data, err = logc.BuildUpdatePayload(*logUpdateMessageFlag)
			case "delete":
				endpoint = c.Delete()
				data, err = logc.BuildDeletePayload(*logDeleteMessageFlag)
			case "list":
				endpoint = c.List()
				data, err = logc.BuildListPayload(*logListMessageFlag)
			}
		case "fire":
			c := firec.NewClient(cc, opts...)
			switch epn {
			case "create":
				endpoint = c.Create()
				data, err = firec.BuildCreatePayload(*fireCreateMessageFlag)
			case "get":
				endpoint = c.Get()
				data, err = firec.BuildGetPayload(*fireGetMessageFlag)
			case "update":
				endpoint = c.Update()
				data, err = firec.BuildUpdatePayload(*fireUpdateMessageFlag)
			case "delete":
				endpoint = c.Delete()
				data, err = firec.BuildDeletePayload(*fireDeleteMessageFlag)
			case "list":
				endpoint = c.List()
				data, err = firec.BuildListPayload(*fireListMessageFlag)
			case "get-weather-for-fire":
				endpoint = c.GetWeatherForFire()
				data, err = firec.BuildGetWeatherForFirePayload(*fireGetWeatherForFireMessageFlag)
			case "get-logs-for-fire":
				endpoint = c.GetLogsForFire()
				data, err = firec.BuildGetLogsForFirePayload(*fireGetLogsForFireMessageFlag)
			}
		case "weather":
			c := weatherc.NewClient(cc, opts...)
			switch epn {
			case "create":
				endpoint = c.Create()
				data, err = weatherc.BuildCreatePayload(*weatherCreateMessageFlag)
			case "get":
				endpoint = c.Get()
				data, err = weatherc.BuildGetPayload(*weatherGetMessageFlag)
			case "update":
				endpoint = c.Update()
				data, err = weatherc.BuildUpdatePayload(*weatherUpdateMessageFlag)
			case "delete":
				endpoint = c.Delete()
				data, err = weatherc.BuildDeletePayload(*weatherDeleteMessageFlag)
			case "list":
				endpoint = c.List()
				data, err = weatherc.BuildListPayload(*weatherListMessageFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// logUsage displays the usage of the log command and its subcommands.
func logUsage() {
	fmt.Fprintf(os.Stderr, `The fire service creates new fires, updates data for a fire, deletes fires, and gets/lists fires
Usage:
    %s [globalflags] log COMMAND [flags]

COMMAND:
    create: Add a log to existing fire
    get: Get log and data friends
    update: Update something about a log specifically
    delete: Delete some log specifically
    list: List fires

Additional help:
    %s log COMMAND --help
`, os.Args[0], os.Args[0])
}
func logCreateUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] log create -message JSON

Add a log to existing fire
    -message JSON: 

Example:
    `+os.Args[0]+` log create --message '{
      "addedAt": "1977-05-06T14:52:07Z",
      "createdAt": "1994-05-07T16:07:26Z",
      "fireID": 6736437899940715186,
      "id": 1913927608535220508,
      "name": "Illo ut et excepturi et adipisci.",
      "size": "S",
      "updatedAt": "1981-01-10T06:48:05Z",
      "weather": {
         "createdAt": "2013-03-31T00:00:31Z",
         "dewPoint": {
            "unit": "F",
            "value": 727998182
         },
         "fireID": 5021708757207724489,
         "humidity": 1961871777,
         "id": 4545438651844921453,
         "logID": 1167531932544065094,
         "temperature": {
            "unit": "F",
            "value": 727998182
         },
         "weatherTime": "2003-07-28T19:16:39Z",
         "weatherType": "Windy",
         "wind": {
            "direction": "SE",
            "speed": 740748338,
            "unit": "MPH"
         }
      }
   }'
`, os.Args[0])
}

func logGetUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] log get -message JSON

Get log and data friends
    -message JSON: 

Example:
    `+os.Args[0]+` log get --message '{
      "addedAt": "1993-04-24T08:45:18Z",
      "createdAt": "1998-06-02T03:20:26Z",
      "fireID": 6408389337853557776,
      "id": 4532697486990614527,
      "name": "Qui dignissimos laborum.",
      "size": "S",
      "updatedAt": "1987-07-27T11:51:34Z",
      "weather": {
         "createdAt": "2013-03-31T00:00:31Z",
         "dewPoint": {
            "unit": "F",
            "value": 727998182
         },
         "fireID": 5021708757207724489,
         "humidity": 1961871777,
         "id": 4545438651844921453,
         "logID": 1167531932544065094,
         "temperature": {
            "unit": "F",
            "value": 727998182
         },
         "weatherTime": "2003-07-28T19:16:39Z",
         "weatherType": "Windy",
         "wind": {
            "direction": "SE",
            "speed": 740748338,
            "unit": "MPH"
         }
      }
   }'
`, os.Args[0])
}

func logUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] log update -message JSON

Update something about a log specifically
    -message JSON: 

Example:
    `+os.Args[0]+` log update --message '{
      "addedAt": "1972-02-16T20:27:18Z",
      "createdAt": "2012-09-13T17:23:27Z",
      "fireID": 2482404165864863284,
      "id": 5531877719792856122,
      "name": "Rem et asperiores et qui qui quibusdam.",
      "size": "M",
      "updatedAt": "1992-08-03T04:42:16Z",
      "weather": {
         "createdAt": "2013-03-31T00:00:31Z",
         "dewPoint": {
            "unit": "F",
            "value": 727998182
         },
         "fireID": 5021708757207724489,
         "humidity": 1961871777,
         "id": 4545438651844921453,
         "logID": 1167531932544065094,
         "temperature": {
            "unit": "F",
            "value": 727998182
         },
         "weatherTime": "2003-07-28T19:16:39Z",
         "weatherType": "Windy",
         "wind": {
            "direction": "SE",
            "speed": 740748338,
            "unit": "MPH"
         }
      }
   }'
`, os.Args[0])
}

func logDeleteUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] log delete -message JSON

Delete some log specifically
    -message JSON: 

Example:
    `+os.Args[0]+` log delete --message '{
      "addedAt": "1994-10-18T19:42:57Z",
      "createdAt": "1980-09-14T10:41:37Z",
      "fireID": 8736365977478103614,
      "id": 7218447836258405858,
      "name": "Eos soluta.",
      "size": "L",
      "updatedAt": "2003-01-06T12:49:53Z",
      "weather": {
         "createdAt": "2013-03-31T00:00:31Z",
         "dewPoint": {
            "unit": "F",
            "value": 727998182
         },
         "fireID": 5021708757207724489,
         "humidity": 1961871777,
         "id": 4545438651844921453,
         "logID": 1167531932544065094,
         "temperature": {
            "unit": "F",
            "value": 727998182
         },
         "weatherTime": "2003-07-28T19:16:39Z",
         "weatherType": "Windy",
         "wind": {
            "direction": "SE",
            "speed": 740748338,
            "unit": "MPH"
         }
      }
   }'
`, os.Args[0])
}

func logListUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] log list -message JSON

List fires
    -message JSON: 

Example:
    `+os.Args[0]+` log list --message '{
      "filters": {
         "end": [
            {
               "key": "1996-06-12T08:40:44Z",
               "operator": "\u003e=",
               "value": 5573547598936979342
            },
            {
               "key": "1996-06-12T08:40:44Z",
               "operator": "\u003e=",
               "value": 5573547598936979342
            },
            {
               "key": "1996-06-12T08:40:44Z",
               "operator": "\u003e=",
               "value": 5573547598936979342
            }
         ],
         "name": [
            {
               "key": "Qui hic.",
               "operator": "!=",
               "value": "Quia sint eligendi eum architecto sint est."
            },
            {
               "key": "Qui hic.",
               "operator": "!=",
               "value": "Quia sint eligendi eum architecto sint est."
            },
            {
               "key": "Qui hic.",
               "operator": "!=",
               "value": "Quia sint eligendi eum architecto sint est."
            },
            {
               "key": "Qui hic.",
               "operator": "!=",
               "value": "Quia sint eligendi eum architecto sint est."
            }
         ],
         "start": [
            {
               "key": "1996-06-12T08:40:44Z",
               "operator": "\u003e=",
               "value": 5573547598936979342
            },
            {
               "key": "1996-06-12T08:40:44Z",
               "operator": "\u003e=",
               "value": 5573547598936979342
            }
         ]
      },
      "pagination": {
         "limit": 2313523494388906239,
         "page": 7128205722982720020
      },
      "search": {
         "description": "Exercitationem animi.",
         "name": "Et ipsa nihil officia accusantium beatae doloremque."
      },
      "sort": {
         "end": "ASC, DESC",
         "id": "ASC, DESC",
         "start": "ASC, DESC"
      }
   }'
`, os.Args[0])
}

// fireUsage displays the usage of the fire command and its subcommands.
func fireUsage() {
	fmt.Fprintf(os.Stderr, `The fire service creates new fires, updates data for a fire, deletes fires, and gets/lists fires
Usage:
    %s [globalflags] fire COMMAND [flags]

COMMAND:
    create: Create a fire and optional payloads
    get: Get fire and data friends
    update: Update something about a fire specifically
    delete: Update something about a fire specifically
    list: List fires
    get-weather-for-fire: Gets a list of weather for a fire
    get-logs-for-fire: Gets a list of logs for a fire

Additional help:
    %s fire COMMAND --help
`, os.Args[0], os.Args[0])
}
func fireCreateUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] fire create -message JSON

Create a fire and optional payloads
    -message JSON: 

Example:
    `+os.Args[0]+` fire create --message '{
      "createdAt": "2001-07-26T15:42:42Z",
      "deletedAt": "2008-01-23T13:39:15Z",
      "description": "Vero dolorum enim et aut.",
      "end": "2002-03-03T18:26:05Z",
      "id": 5117576016712403101,
      "name": "Aliquid odio asperiores totam dignissimos.",
      "start": "1987-09-07T04:18:11Z",
      "updatedAt": "1985-03-21T02:39:44Z"
   }'
`, os.Args[0])
}

func fireGetUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] fire get -message JSON

Get fire and data friends
    -message JSON: 

Example:
    `+os.Args[0]+` fire get --message '{
      "id": 3471621047790542047
   }'
`, os.Args[0])
}

func fireUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] fire update -message JSON

Update something about a fire specifically
    -message JSON: 

Example:
    `+os.Args[0]+` fire update --message '{
      "createdAt": "1979-12-25T16:28:11Z",
      "deletedAt": "2009-02-13T05:16:56Z",
      "description": "Consequuntur ea ipsum deleniti modi.",
      "end": "1976-08-28T20:58:47Z",
      "id": 5069143882329837951,
      "name": "Dolor quis sit.",
      "start": "1988-03-22T03:56:20Z",
      "updatedAt": "2006-06-29T13:18:00Z"
   }'
`, os.Args[0])
}

func fireDeleteUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] fire delete -message JSON

Update something about a fire specifically
    -message JSON: 

Example:
    `+os.Args[0]+` fire delete --message '{
      "id": 1211853009398384327
   }'
`, os.Args[0])
}

func fireListUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] fire list -message JSON

List fires
    -message JSON: 

Example:
    `+os.Args[0]+` fire list --message '{
      "filters": {
         "end": [
            {
               "key": "1996-06-12T08:40:44Z",
               "operator": "\u003e=",
               "value": 5573547598936979342
            },
            {
               "key": "1996-06-12T08:40:44Z",
               "operator": "\u003e=",
               "value": 5573547598936979342
            },
            {
               "key": "1996-06-12T08:40:44Z",
               "operator": "\u003e=",
               "value": 5573547598936979342
            },
            {
               "key": "1996-06-12T08:40:44Z",
               "operator": "\u003e=",
               "value": 5573547598936979342
            }
         ],
         "name": [
            {
               "key": "Qui hic.",
               "operator": "!=",
               "value": "Quia sint eligendi eum architecto sint est."
            },
            {
               "key": "Qui hic.",
               "operator": "!=",
               "value": "Quia sint eligendi eum architecto sint est."
            },
            {
               "key": "Qui hic.",
               "operator": "!=",
               "value": "Quia sint eligendi eum architecto sint est."
            }
         ],
         "start": [
            {
               "key": "1996-06-12T08:40:44Z",
               "operator": "\u003e=",
               "value": 5573547598936979342
            },
            {
               "key": "1996-06-12T08:40:44Z",
               "operator": "\u003e=",
               "value": 5573547598936979342
            }
         ]
      },
      "pagination": {
         "limit": 7720515675719359881,
         "page": 5747059457513862389
      },
      "search": {
         "description": "Rem quaerat vel vero velit et.",
         "name": "Quo reiciendis omnis ut."
      },
      "sort": {
         "end": "ASC, DESC",
         "id": "ASC, DESC",
         "start": "ASC, DESC"
      }
   }'
`, os.Args[0])
}

func fireGetWeatherForFireUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] fire get-weather-for-fire -message JSON

Gets a list of weather for a fire
    -message JSON: 

Example:
    `+os.Args[0]+` fire get-weather-for-fire --message '{
      "id": 141765740622424101
   }'
`, os.Args[0])
}

func fireGetLogsForFireUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] fire get-logs-for-fire -message JSON

Gets a list of logs for a fire
    -message JSON: 

Example:
    `+os.Args[0]+` fire get-logs-for-fire --message '{
      "id": 5391076212118225793
   }'
`, os.Args[0])
}

// weatherUsage displays the usage of the weather command and its subcommands.
func weatherUsage() {
	fmt.Fprintf(os.Stderr, `The fire service creates new weather datas for fires, and gets/lists weather data
Usage:
    %s [globalflags] weather COMMAND [flags]

COMMAND:
    create: Create a weather observation and optional payloads
    get: Get a specific piece of weather data
    update: Update something about a weather observation specifically
    delete: Delete a weather observation
    list: List weather

Additional help:
    %s weather COMMAND --help
`, os.Args[0], os.Args[0])
}
func weatherCreateUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] weather create -message JSON

Create a weather observation and optional payloads
    -message JSON: 

Example:
    `+os.Args[0]+` weather create --message '{
      "createdAt": "1997-03-09T21:43:11Z",
      "dewPoint": {
         "unit": "F",
         "value": 727998182
      },
      "fireID": 82313619316836653,
      "humidity": 1148712326,
      "id": 8860917851139109311,
      "logID": 1879967797057239286,
      "temperature": {
         "unit": "F",
         "value": 727998182
      },
      "weatherTime": "1982-04-16T02:42:32Z",
      "weatherType": "Cloudy",
      "wind": {
         "direction": "SE",
         "speed": 740748338,
         "unit": "MPH"
      }
   }'
`, os.Args[0])
}

func weatherGetUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] weather get -message JSON

Get a specific piece of weather data
    -message JSON: 

Example:
    `+os.Args[0]+` weather get --message '{
      "createdAt": "1986-06-12T00:41:34Z",
      "dewPoint": {
         "unit": "F",
         "value": 727998182
      },
      "fireID": 4697907119565893333,
      "humidity": 745260731,
      "id": 5931137490381952062,
      "logID": 5468251973968172269,
      "temperature": {
         "unit": "F",
         "value": 727998182
      },
      "weatherTime": "2010-04-01T11:38:38Z",
      "weatherType": "Sunny",
      "wind": {
         "direction": "SE",
         "speed": 740748338,
         "unit": "MPH"
      }
   }'
`, os.Args[0])
}

func weatherUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] weather update -message JSON

Update something about a weather observation specifically
    -message JSON: 

Example:
    `+os.Args[0]+` weather update --message '{
      "createdAt": "2006-08-09T07:44:27Z",
      "dewPoint": {
         "unit": "F",
         "value": 727998182
      },
      "fireID": 3927858731322961439,
      "humidity": 1092504710,
      "id": 1656434714664168450,
      "logID": 6351804479751839719,
      "temperature": {
         "unit": "F",
         "value": 727998182
      },
      "weatherTime": "1974-08-12T23:30:42Z",
      "weatherType": "Sunny",
      "wind": {
         "direction": "SE",
         "speed": 740748338,
         "unit": "MPH"
      }
   }'
`, os.Args[0])
}

func weatherDeleteUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] weather delete -message JSON

Delete a weather observation
    -message JSON: 

Example:
    `+os.Args[0]+` weather delete --message '{
      "createdAt": "1998-06-20T02:24:02Z",
      "dewPoint": {
         "unit": "F",
         "value": 727998182
      },
      "fireID": 5087279065898837603,
      "humidity": 1206286630,
      "id": 6616655541483071643,
      "logID": 4857834743276712703,
      "temperature": {
         "unit": "F",
         "value": 727998182
      },
      "weatherTime": "1973-05-08T21:40:35Z",
      "weatherType": "Cloudy",
      "wind": {
         "direction": "SE",
         "speed": 740748338,
         "unit": "MPH"
      }
   }'
`, os.Args[0])
}

func weatherListUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] weather list -message JSON

List weather
    -message JSON: 

Example:
    `+os.Args[0]+` weather list --message '{
      "filters": {
         "humidity": [
            {
               "key": "Saepe inventore voluptates voluptas.",
               "operator": "\u003c=",
               "value": 8441511991902083154
            },
            {
               "key": "Saepe inventore voluptates voluptas.",
               "operator": "\u003c=",
               "value": 8441511991902083154
            }
         ],
         "temperature": [
            {
               "key": "Saepe inventore voluptates voluptas.",
               "operator": "\u003c=",
               "value": 8441511991902083154
            },
            {
               "key": "Saepe inventore voluptates voluptas.",
               "operator": "\u003c=",
               "value": 8441511991902083154
            },
            {
               "key": "Saepe inventore voluptates voluptas.",
               "operator": "\u003c=",
               "value": 8441511991902083154
            }
         ],
         "time": [
            {
               "key": "1996-06-12T08:40:44Z",
               "operator": "\u003e=",
               "value": 5573547598936979342
            },
            {
               "key": "1996-06-12T08:40:44Z",
               "operator": "\u003e=",
               "value": 5573547598936979342
            },
            {
               "key": "1996-06-12T08:40:44Z",
               "operator": "\u003e=",
               "value": 5573547598936979342
            }
         ],
         "windSpeed": [
            {
               "key": "Saepe inventore voluptates voluptas.",
               "operator": "\u003c=",
               "value": 8441511991902083154
            },
            {
               "key": "Saepe inventore voluptates voluptas.",
               "operator": "\u003c=",
               "value": 8441511991902083154
            }
         ]
      },
      "pagination": {
         "limit": 558397952000332283,
         "page": 2106131844815631837
      },
      "search": {
         "description": "Praesentium dolores.",
         "name": "Nemo sint facilis omnis at."
      },
      "sort": {
         "id": "ASC, DESC"
      }
   }'
`, os.Args[0])
}
