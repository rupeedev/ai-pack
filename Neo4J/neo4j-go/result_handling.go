package main

import (
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// Example 1: Working with Nodes
func handleNodes(ctx context.Context, driver neo4j.DriverWithContext) {
	fmt.Println("\n=== Example 1: Working with Nodes ===")

	cypher := `
		MATCH (m:Movie {title: $title})
		RETURN m
	`

	result, err := neo4j.ExecuteQuery(ctx, driver,
		cypher,
		map[string]any{"title": "Toy Story"},
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if len(result.Records) == 0 {
		fmt.Println("No results found. Run 'go run . setup' first.")
		return
	}

	for _, record := range result.Records {
		node, ok := record.Get("m")
		if !ok {
			continue
		}

		movieNode := node.(neo4j.Node)

		// Access node properties
		fmt.Printf("Element ID: %s\n", movieNode.ElementId)
		fmt.Printf("Labels: %v\n", movieNode.Labels)
		fmt.Printf("All Properties: %v\n", movieNode.Props)

		// Access specific properties
		if title, ok := movieNode.Props["title"]; ok {
			fmt.Printf("Title: %s\n", title)
		}
		if released, ok := movieNode.Props["released"]; ok {
			fmt.Printf("Released: %v\n", released)
		}
		if tagline, ok := movieNode.Props["tagline"]; ok {
			fmt.Printf("Tagline: %s\n", tagline)
		}
	}
}

// Example 2: Working with Relationships
func handleRelationships(ctx context.Context, driver neo4j.DriverWithContext) {
	fmt.Println("\n=== Example 2: Working with Relationships ===")

	cypher := `
		MATCH (p:Person)-[r:ACTED_IN]->(m:Movie {title: $title})
		RETURN p, r, m
	`

	result, err := neo4j.ExecuteQuery(ctx, driver,
		cypher,
		map[string]any{"title": "Toy Story"},
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if len(result.Records) == 0 {
		fmt.Println("No results found. Run 'go run . setup' first.")
		return
	}

	for _, record := range result.Records {
		// Get person node
		personNode, _ := record.Get("p")
		person := personNode.(neo4j.Node)

		// Get relationship
		relData, _ := record.Get("r")
		relationship := relData.(neo4j.Relationship)

		// Get movie node
		movieNode, _ := record.Get("m")
		movie := movieNode.(neo4j.Node)

		fmt.Printf("\nRelationship Details:\n")
		fmt.Printf("  Element ID: %s\n", relationship.ElementId)
		fmt.Printf("  Type: %s\n", relationship.Type)
		fmt.Printf("  Start Node ID: %s\n", relationship.StartElementId)
		fmt.Printf("  End Node ID: %s\n", relationship.EndElementId)

		// Access relationship properties
		if role, ok := relationship.Props["role"]; ok {
			fmt.Printf("  Role: %s\n", role)
		}

		// Show the connection
		personName := person.Props["name"]
		movieTitle := movie.Props["title"]
		fmt.Printf("\n  %s -[%s]-> %s\n", personName, relationship.Type, movieTitle)
	}
}

// Example 3: Working with Paths
func handlePaths(ctx context.Context, driver neo4j.DriverWithContext) {
	fmt.Println("\n=== Example 3: Working with Paths ===")

	cypher := `
		MATCH path = (p:Person)-[r:ACTED_IN]->(m:Movie {title: $title})
		RETURN path
		LIMIT 1
	`

	result, err := neo4j.ExecuteQuery(ctx, driver,
		cypher,
		map[string]any{"title": "The Matrix"},
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if len(result.Records) == 0 {
		fmt.Println("No results found. Run 'go run . setup' first.")
		return
	}

	for _, record := range result.Records {
		pathData, _ := record.Get("path")
		path := pathData.(neo4j.Path)

		fmt.Printf("Path length: %d\n", len(path.Relationships))
		fmt.Printf("Number of nodes: %d\n", len(path.Nodes))

		// Iterate over nodes
		fmt.Println("\nNodes in path:")
		for i, node := range path.Nodes {
			name := node.Props["name"]
			title := node.Props["title"]
			if name != nil {
				fmt.Printf("  %d. Person: %s\n", i, name)
			} else if title != nil {
				fmt.Printf("  %d. Movie: %s\n", i, title)
			}
		}

		// Iterate over relationships
		fmt.Println("\nRelationships in path:")
		for i, rel := range path.Relationships {
			role := rel.Props["role"]
			fmt.Printf("  %d. %s (role: %s)\n", i, rel.Type, role)
		}
	}
}

// Example 4: Working with Multiple Paths
func handleMultiplePaths(ctx context.Context, driver neo4j.DriverWithContext) {
	fmt.Println("\n=== Example 4: Working with Multiple Paths ===")

	cypher := `
		MATCH path = (p:Person)-[r:ACTED_IN]->(m:Movie)
		WHERE m.title STARTS WITH 'The Matrix'
		RETURN path
		LIMIT 3
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

	fmt.Printf("Found %d paths\n", len(result.Records))

	for idx, record := range result.Records {
		pathData, _ := record.Get("path")
		path := pathData.(neo4j.Path)

		// Extract person and movie from path
		person := path.Nodes[0]
		movie := path.Nodes[1]
		relationship := path.Relationships[0]

		personName := person.Props["name"]
		movieTitle := movie.Props["title"]
		role := relationship.Props["role"]

		fmt.Printf("%d. %s played %s in %s\n",
			idx+1, personName, role, movieTitle)
	}
}

// Example 5: Working with Temporal Types
func handleTemporalTypes(ctx context.Context, driver neo4j.DriverWithContext) {
	fmt.Println("\n=== Example 5: Working with Temporal Types ===")

	// Create a person with birth date
	cypher := `
		MERGE (p:Person {name: "Test Person"})
		SET p.born = date('1990-05-15'),
		    p.createdAt = datetime(),
		    p.lastUpdated = localdatetime()
		RETURN p
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
		personData, _ := result.Records[0].Get("p")
		person := personData.(neo4j.Node)

		fmt.Println("Temporal properties:")

		if born, ok := person.Props["born"]; ok {
			// Neo4j Date type
			if date, ok := born.(neo4j.Date); ok {
				fmt.Printf("  Born (Date): %s\n", date)

				// Convert to Go time.Time
				goTime := date.Time()
				fmt.Printf("  Born (time.Time): %s\n", goTime.Format("2006-01-02"))
				fmt.Printf("  Born (formatted): %s\n", goTime.Format("January 2, 2006"))
			}
		}

		if createdAt, ok := person.Props["createdAt"]; ok {
			fmt.Printf("  Created At: %v\n", createdAt)
		}

		if lastUpdated, ok := person.Props["lastUpdated"]; ok {
			fmt.Printf("  Last Updated: %v\n", lastUpdated)
		}
	}

	// Cleanup
	_, _ = neo4j.ExecuteQuery(ctx, driver,
		"MATCH (p:Person {name: 'Test Person'}) DELETE p",
		nil,
		neo4j.EagerResultTransformer,
	)
}

// Example 6: Working with Spatial Types
func handleSpatialTypes(ctx context.Context, driver neo4j.DriverWithContext) {
	fmt.Println("\n=== Example 6: Working with Spatial Types ===")

	// Create locations with spatial data
	cypher := `
		MERGE (office:Location {name: "Office"})
		SET office.location = point({latitude: 37.7749, longitude: -122.4194})
		MERGE (home:Location {name: "Home"})
		SET home.location = point({latitude: 37.7849, longitude: -122.4094})
		RETURN office, home
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
		officeData, _ := result.Records[0].Get("office")
		office := officeData.(neo4j.Node)

		homeData, _ := result.Records[0].Get("home")
		home := homeData.(neo4j.Node)

		fmt.Println("Spatial properties:")

		if location, ok := office.Props["location"]; ok {
			if point, ok := location.(neo4j.Point2D); ok {
				fmt.Printf("  Office: (lat: %.4f, lng: %.4f)\n",
					point.Y, point.X)
			}
		}

		if location, ok := home.Props["location"]; ok {
			if point, ok := location.(neo4j.Point2D); ok {
				fmt.Printf("  Home: (lat: %.4f, lng: %.4f)\n",
					point.Y, point.X)
			}
		}
	}

	// Calculate distance
	distanceCypher := `
		MATCH (o:Location {name: "Office"}), (h:Location {name: "Home"})
		RETURN point.distance(o.location, h.location) AS distance
	`

	distResult, err := neo4j.ExecuteQuery(ctx, driver,
		distanceCypher,
		nil,
		neo4j.EagerResultTransformer,
	)

	if err == nil && len(distResult.Records) > 0 {
		distance, _ := distResult.Records[0].Get("distance")
		fmt.Printf("\nDistance between Office and Home: %.2f meters\n", distance)
	}

	// Cleanup
	_, _ = neo4j.ExecuteQuery(ctx, driver,
		"MATCH (l:Location) WHERE l.name IN ['Office', 'Home'] DELETE l",
		nil,
		neo4j.EagerResultTransformer,
	)
}

// Example 7: Type Mapping Summary
func showTypeMappings(ctx context.Context, driver neo4j.DriverWithContext) {
	fmt.Println("\n=== Example 7: Type Mapping Summary ===")

	cypher := `
		RETURN
			null AS nullValue,
			true AS boolValue,
			42 AS intValue,
			3.14 AS floatValue,
			'Hello' AS stringValue,
			[1, 2, 3] AS listValue,
			{key: 'value'} AS mapValue,
			date('2025-01-01') AS dateValue,
			point({latitude: 37.7749, longitude: -122.4194}) AS pointValue
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

		fmt.Println("Neo4j Type → Go Type:")

		nullVal, _ := record.Get("nullValue")
		fmt.Printf("  null → %T (%v)\n", nullVal, nullVal)

		boolVal, _ := record.Get("boolValue")
		fmt.Printf("  Boolean → %T (%v)\n", boolVal, boolVal)

		intVal, _ := record.Get("intValue")
		fmt.Printf("  Integer → %T (%v)\n", intVal, intVal)

		floatVal, _ := record.Get("floatValue")
		fmt.Printf("  Float → %T (%v)\n", floatVal, floatVal)

		stringVal, _ := record.Get("stringValue")
		fmt.Printf("  String → %T (%v)\n", stringVal, stringVal)

		listVal, _ := record.Get("listValue")
		fmt.Printf("  List → %T (%v)\n", listVal, listVal)

		mapVal, _ := record.Get("mapValue")
		fmt.Printf("  Map → %T (%v)\n", mapVal, mapVal)

		dateVal, _ := record.Get("dateValue")
		fmt.Printf("  Date → %T (%v)\n", dateVal, dateVal)

		pointVal, _ := record.Get("pointValue")
		fmt.Printf("  Point → %T (%v)\n", pointVal, pointVal)
	}
}

func runResultHandlingExamples() {
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

	// Run all examples
	handleNodes(ctx, driver)
	handleRelationships(ctx, driver)
	handlePaths(ctx, driver)
	handleMultiplePaths(ctx, driver)
	handleTemporalTypes(ctx, driver)
	handleSpatialTypes(ctx, driver)
	showTypeMappings(ctx, driver)

	fmt.Println("\n========================================")
	fmt.Println("All result handling examples completed!")
}
