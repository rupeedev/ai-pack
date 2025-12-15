package main

import (
	"context"
	"fmt"
	"time"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// ============================================================================
// TEMPORAL TYPES EXAMPLES
// ============================================================================

// Example 1: Writing Temporal Types
func writingTemporalTypes(ctx context.Context, driver neo4j.DriverWithContext) {
	fmt.Println("\n=== Example 1: Writing Temporal Types ===")

	// Create an event with different temporal types
	loc, _ := time.LoadLocation("America/New_York")
	eventTime := time.Date(2024, 5, 15, 14, 30, 0, 0, loc)

	cypher := `
		CREATE (e:Event {
			name: $name,
			startsAt: $datetime,
			createdAt: datetime($dtstring),
			updatedAt: datetime()
		})
		RETURN e
	`

	result, err := neo4j.ExecuteQuery(ctx, driver,
		cypher,
		map[string]any{
			"name":     "GraphConnect 2024",
			"datetime": eventTime,
			"dtstring": "2024-05-15T14:30:00-04:00",
		},
		neo4j.EagerResultTransformer,
	)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if len(result.Records) > 0 {
		event, _ := result.Records[0].Get("e")
		eventNode := event.(neo4j.Node)
		fmt.Printf("Created event: %s\n", eventNode.Props["name"])
		fmt.Printf("  Starts at: %v\n", eventNode.Props["startsAt"])
		fmt.Printf("  Created at: %v\n", eventNode.Props["createdAt"])
		fmt.Printf("  Updated at: %v\n", eventNode.Props["updatedAt"])
	}

	// Cleanup
	_, _ = neo4j.ExecuteQuery(ctx, driver,
		"MATCH (e:Event {name: 'GraphConnect 2024'}) DELETE e",
		nil,
		neo4j.EagerResultTransformer,
	)
}

// Example 2: Reading Temporal Types
func readingTemporalTypes(ctx context.Context, driver neo4j.DriverWithContext) {
	fmt.Println("\n=== Example 2: Reading Temporal Types ===")

	cypher := `
		RETURN
			date() as currentDate,
			time() as currentTime,
			datetime() as currentDateTime,
			localdatetime() as localDateTime,
			localtime() as localTime,
			toString(datetime()) as asString
	`

	result, err := neo4j.ExecuteQuery(ctx, driver,
		cypher,
		nil,
		neo4j.EagerResultTransformer,
	)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if len(result.Records) > 0 {
		record := result.Records[0]

		// Access different temporal types
		currentDate, _ := record.Get("currentDate")
		currentTime, _ := record.Get("currentTime")
		currentDateTime, _ := record.Get("currentDateTime")
		localDateTime, _ := record.Get("localDateTime")
		localTime, _ := record.Get("localTime")
		asString, _ := record.Get("asString")

		fmt.Println("Temporal types returned from Neo4j:")
		fmt.Printf("  Date: %v (type: %T)\n", currentDate, currentDate)
		fmt.Printf("  Time: %v (type: %T)\n", currentTime, currentTime)
		fmt.Printf("  DateTime: %v (type: %T)\n", currentDateTime, currentDateTime)
		fmt.Printf("  LocalDateTime: %v (type: %T)\n", localDateTime, localDateTime)
		fmt.Printf("  LocalTime: %v (type: %T)\n", localTime, localTime)
		fmt.Printf("  String: %v (type: %T)\n", asString, asString)
	}
}

// Example 3: Working with Durations
func workingWithDurations(ctx context.Context, driver neo4j.DriverWithContext) {
	fmt.Println("\n=== Example 3: Working with Durations ===")

	// Use UTC to avoid timezone issues with Neo4j
	startsAt := time.Now().UTC()

	// Create a Neo4j Duration (1 hour 30 minutes)
	eventLength := neo4j.DurationOf(
		0,   // months
		0,   // days
		int64((time.Hour + 30*time.Minute).Seconds()), // seconds
		0,   // nanos
	)

	// Calculate end time using Go's time.Duration
	endsAt := startsAt.Add(time.Hour + 30*time.Minute)

	cypher := `
		CREATE (e:Event {
			name: $name,
			startsAt: $startsAt,
			endsAt: $endsAt,
			duration: $eventLength,
			interval: duration('PT1H30M')
		})
		RETURN e
	`

	result, err := neo4j.ExecuteQuery(ctx, driver,
		cypher,
		map[string]any{
			"name":        "Workshop",
			"startsAt":    startsAt,
			"endsAt":      endsAt,
			"eventLength": eventLength,
		},
		neo4j.EagerResultTransformer,
	)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if len(result.Records) > 0 {
		event, _ := result.Records[0].Get("e")
		eventNode := event.(neo4j.Node)

		fmt.Printf("Event: %s\n", eventNode.Props["name"])
		fmt.Printf("  Duration: %v\n", eventNode.Props["duration"])
		fmt.Printf("  Interval: %v\n", eventNode.Props["interval"])

		// Calculate duration in Cypher
		durationQuery := `
			MATCH (e:Event {name: 'Workshop'})
			RETURN duration.between(e.startsAt, e.endsAt) AS calculatedDuration
		`

		durationResult, err := neo4j.ExecuteQuery(ctx, driver,
			durationQuery,
			nil,
			neo4j.EagerResultTransformer,
		)

		if err == nil && len(durationResult.Records) > 0 {
			calculated, _ := durationResult.Records[0].Get("calculatedDuration")
			fmt.Printf("  Calculated duration: %v\n", calculated)
		}
	}

	// Cleanup
	_, _ = neo4j.ExecuteQuery(ctx, driver,
		"MATCH (e:Event {name: 'Workshop'}) DELETE e",
		nil,
		neo4j.EagerResultTransformer,
	)
}

// Example 4: Date Arithmetic
func dateArithmetic(ctx context.Context, driver neo4j.DriverWithContext) {
	fmt.Println("\n=== Example 4: Date Arithmetic ===")

	cypher := `
		WITH date('2024-01-15') AS baseDate
		RETURN
			baseDate,
			baseDate + duration('P1M') AS plusOneMonth,
			baseDate + duration('P7D') AS plusOneWeek,
			baseDate - duration('P1Y') AS minusOneYear
	`

	result, err := neo4j.ExecuteQuery(ctx, driver,
		cypher,
		nil,
		neo4j.EagerResultTransformer,
	)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if len(result.Records) > 0 {
		record := result.Records[0]

		baseDate, _ := record.Get("baseDate")
		plusOneMonth, _ := record.Get("plusOneMonth")
		plusOneWeek, _ := record.Get("plusOneWeek")
		minusOneYear, _ := record.Get("minusOneYear")

		fmt.Println("Date arithmetic:")
		fmt.Printf("  Base date: %v\n", baseDate)
		fmt.Printf("  + 1 month: %v\n", plusOneMonth)
		fmt.Printf("  + 1 week: %v\n", plusOneWeek)
		fmt.Printf("  - 1 year: %v\n", minusOneYear)
	}
}

// ============================================================================
// SPATIAL TYPES EXAMPLES
// ============================================================================

// Example 5: Cartesian Points (2D)
func cartesianPoints2D(ctx context.Context, driver neo4j.DriverWithContext) {
	fmt.Println("\n=== Example 5: Cartesian Points (2D) ===")

	// Create 2D Cartesian points in Go
	point1 := neo4j.Point2D{
		X:            1.23,
		Y:            4.56,
		SpatialRefId: 7203, // Cartesian SRID
	}

	point2 := neo4j.Point2D{
		X:            2.34,
		Y:            5.67,
		SpatialRefId: 7203,
	}

	cypher := `
		CREATE (loc1:Location {name: $name1, position: $point1})
		CREATE (loc2:Location {name: $name2, position: $point2})
		RETURN loc1, loc2
	`

	result, err := neo4j.ExecuteQuery(ctx, driver,
		cypher,
		map[string]any{
			"name1":  "Point A",
			"point1": point1,
			"name2":  "Point B",
			"point2": point2,
		},
		neo4j.EagerResultTransformer,
	)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if len(result.Records) > 0 {
		loc1Data, _ := result.Records[0].Get("loc1")
		loc1 := loc1Data.(neo4j.Node)

		if pos, ok := loc1.Props["position"].(neo4j.Point2D); ok {
			fmt.Printf("Location: %s\n", loc1.Props["name"])
			fmt.Printf("  X: %.2f, Y: %.2f\n", pos.X, pos.Y)
			fmt.Printf("  SRID: %d (Cartesian)\n", pos.SpatialRefId)
		}
	}

	// Cleanup
	_, _ = neo4j.ExecuteQuery(ctx, driver,
		"MATCH (l:Location) WHERE l.name IN ['Point A', 'Point B'] DELETE l",
		nil,
		neo4j.EagerResultTransformer,
	)
}

// Example 6: Cartesian Points (3D)
func cartesianPoints3D(ctx context.Context, driver neo4j.DriverWithContext) {
	fmt.Println("\n=== Example 6: Cartesian Points (3D) ===")

	cypher := `
		RETURN point({x: 1.23, y: 4.56, z: 7.89}) AS point3D
	`

	result, err := neo4j.ExecuteQuery(ctx, driver,
		cypher,
		nil,
		neo4j.EagerResultTransformer,
	)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if len(result.Records) > 0 {
		point, _ := result.Records[0].Get("point3D")

		if p3d, ok := point.(neo4j.Point3D); ok {
			fmt.Println("3D Cartesian Point:")
			fmt.Printf("  X: %.2f, Y: %.2f, Z: %.2f\n", p3d.X, p3d.Y, p3d.Z)
			fmt.Printf("  SRID: %d (3D Cartesian)\n", p3d.SpatialRefId)
		}
	}
}

// Example 7: Geographic Points (WGS-84)
func geographicPoints(ctx context.Context, driver neo4j.DriverWithContext) {
	fmt.Println("\n=== Example 7: Geographic Points (WGS-84) ===")

	// Famous locations
	cypher := `
		CREATE (london:City {
			name: 'London',
			location: point({latitude: 51.509865, longitude: -0.118092})
		})
		CREATE (paris:City {
			name: 'Paris',
			location: point({latitude: 48.8566, longitude: 2.3522})
		})
		CREATE (ny:City {
			name: 'New York',
			location: point({latitude: 40.7128, longitude: -74.0060})
		})
		RETURN london, paris, ny
	`

	result, err := neo4j.ExecuteQuery(ctx, driver,
		cypher,
		nil,
		neo4j.EagerResultTransformer,
	)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if len(result.Records) > 0 {
		cities := []string{"london", "paris", "ny"}

		for _, cityKey := range cities {
			cityData, _ := result.Records[0].Get(cityKey)
			city := cityData.(neo4j.Node)

			if loc, ok := city.Props["location"].(neo4j.Point2D); ok {
				fmt.Printf("%s:\n", city.Props["name"])
				fmt.Printf("  Latitude: %.4f, Longitude: %.4f\n", loc.Y, loc.X)
				fmt.Printf("  SRID: %d (WGS-84)\n", loc.SpatialRefId)
			}
		}
	}

	// Cleanup
	_, _ = neo4j.ExecuteQuery(ctx, driver,
		"MATCH (c:City) WHERE c.name IN ['London', 'Paris', 'New York'] DELETE c",
		nil,
		neo4j.EagerResultTransformer,
	)
}

// Example 8: Geographic Points with Height (3D)
func geographicPoints3D(ctx context.Context, driver neo4j.DriverWithContext) {
	fmt.Println("\n=== Example 8: Geographic Points with Height (3D) ===")

	// The Shard in London with height
	shard := neo4j.Point3D{
		X:            -0.086500, // longitude
		Y:            51.504501, // latitude
		Z:            310,       // height in meters
		SpatialRefId: 4979,      // WGS-84 3D SRID
	}

	cypher := `
		CREATE (building:Building {
			name: $name,
			location: $location
		})
		RETURN building
	`

	result, err := neo4j.ExecuteQuery(ctx, driver,
		cypher,
		map[string]any{
			"name":     "The Shard",
			"location": shard,
		},
		neo4j.EagerResultTransformer,
	)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if len(result.Records) > 0 {
		buildingData, _ := result.Records[0].Get("building")
		building := buildingData.(neo4j.Node)

		if loc, ok := building.Props["location"].(neo4j.Point3D); ok {
			fmt.Printf("%s:\n", building.Props["name"])
			fmt.Printf("  Latitude: %.6f\n", loc.Y)
			fmt.Printf("  Longitude: %.6f\n", loc.X)
			fmt.Printf("  Height: %.0f meters\n", loc.Z)
			fmt.Printf("  SRID: %d (WGS-84 3D)\n", loc.SpatialRefId)
		}
	}

	// Cleanup
	_, _ = neo4j.ExecuteQuery(ctx, driver,
		"MATCH (b:Building {name: 'The Shard'}) DELETE b",
		nil,
		neo4j.EagerResultTransformer,
	)
}

// Example 9: Calculating Distances
func calculatingDistances(ctx context.Context, driver neo4j.DriverWithContext) {
	fmt.Println("\n=== Example 9: Calculating Distances ===")

	// Create points for distance calculation
	point1 := neo4j.Point2D{X: 1.23, Y: 4.56, SpatialRefId: 7203}
	point2 := neo4j.Point2D{X: 2.34, Y: 5.67, SpatialRefId: 7203}

	// Cartesian distance
	cypher := `
		RETURN point.distance($p1, $p2) AS cartesianDistance
	`

	result, err := neo4j.ExecuteQuery(ctx, driver,
		cypher,
		map[string]any{
			"p1": point1,
			"p2": point2,
		},
		neo4j.EagerResultTransformer,
	)

	if err == nil && len(result.Records) > 0 {
		distance, _ := result.Records[0].Get("cartesianDistance")
		fmt.Printf("Cartesian distance: %.2f units\n", distance.(float64))
	}

	// Geographic distance (real cities)
	geoQuery := `
		WITH
			point({latitude: 51.509865, longitude: -0.118092}) AS london,
			point({latitude: 48.8566, longitude: 2.3522}) AS paris
		RETURN point.distance(london, paris) AS distance
	`

	geoResult, err := neo4j.ExecuteQuery(ctx, driver,
		geoQuery,
		nil,
		neo4j.EagerResultTransformer,
	)

	if err == nil && len(geoResult.Records) > 0 {
		distance, _ := geoResult.Records[0].Get("distance")
		distanceMeters := distance.(float64)
		distanceKm := distanceMeters / 1000

		fmt.Printf("\nGeographic distance (London to Paris):\n")
		fmt.Printf("  %.0f meters\n", distanceMeters)
		fmt.Printf("  %.2f kilometers\n", distanceKm)
		fmt.Printf("  %.2f miles\n", distanceKm*0.621371)
	}
}

// Example 10: Finding Nearby Locations
func findingNearbyLocations(ctx context.Context, driver neo4j.DriverWithContext) {
	fmt.Println("\n=== Example 10: Finding Nearby Locations ===")

	// Create several locations
	setupCypher := `
		CREATE (sf:City {name: 'San Francisco', location: point({latitude: 37.7749, longitude: -122.4194})})
		CREATE (oak:City {name: 'Oakland', location: point({latitude: 37.8044, longitude: -122.2712})})
		CREATE (sj:City {name: 'San Jose', location: point({latitude: 37.3382, longitude: -121.8863})})
		CREATE (la:City {name: 'Los Angeles', location: point({latitude: 34.0522, longitude: -118.2437})})
	`

	_, err := neo4j.ExecuteQuery(ctx, driver, setupCypher, nil, neo4j.EagerResultTransformer)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Find cities within 50km of San Francisco
	searchCypher := `
		MATCH (sf:City {name: 'San Francisco'})
		MATCH (other:City)
		WHERE other <> sf
		WITH sf, other, point.distance(sf.location, other.location) AS distance
		WHERE distance < 50000
		RETURN other.name AS city, distance
		ORDER BY distance
	`

	result, err := neo4j.ExecuteQuery(ctx, driver,
		searchCypher,
		nil,
		neo4j.EagerResultTransformer,
	)

	if err == nil {
		fmt.Println("Cities within 50km of San Francisco:")
		for _, record := range result.Records {
			city, _ := record.Get("city")
			distance, _ := record.Get("distance")
			distanceKm := distance.(float64) / 1000
			fmt.Printf("  %s: %.2f km\n", city, distanceKm)
		}
	}

	// Cleanup
	_, _ = neo4j.ExecuteQuery(ctx, driver,
		"MATCH (c:City) WHERE c.name IN ['San Francisco', 'Oakland', 'San Jose', 'Los Angeles'] DELETE c",
		nil,
		neo4j.EagerResultTransformer,
	)
}

func runTemporalSpatialExamples() {
	// Setup driver
	driver, err := neo4j.NewDriverWithContext(
		"neo4j://localhost:7687",
		neo4j.BasicAuth("neo4j", "Your@Password!@#", ""),
	)
	if err != nil {
		panic(err)
	}
	defer driver.Close(context.Background())

	ctx := context.Background()

	// Verify connection
	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect: %v", err))
	}

	fmt.Println("Connected to Neo4j successfully!")
	fmt.Println("========================================")

	// Run temporal examples
	fmt.Println("\n** TEMPORAL TYPES **")
	writingTemporalTypes(ctx, driver)
	readingTemporalTypes(ctx, driver)
	workingWithDurations(ctx, driver)
	dateArithmetic(ctx, driver)

	// Run spatial examples
	fmt.Println("\n** SPATIAL TYPES **")
	cartesianPoints2D(ctx, driver)
	cartesianPoints3D(ctx, driver)
	geographicPoints(ctx, driver)
	geographicPoints3D(ctx, driver)
	calculatingDistances(ctx, driver)
	findingNearbyLocations(ctx, driver)

	fmt.Println("\n========================================")
	fmt.Println("All temporal and spatial examples completed!")
}
