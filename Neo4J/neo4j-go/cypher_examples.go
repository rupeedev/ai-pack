package main

import (
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// Example 1: Basic query execution with parameters
func executeWithParameters(ctx context.Context, driver neo4j.DriverWithContext) {
	fmt.Println("\n=== Example 1: Query with Parameters ===")

	cypher := `
		MATCH (a:Person {name: $name})-[r:ACTED_IN]->(m:Movie)
		RETURN a.name AS name, m.title AS title, r.role AS role
	`
	name := "Tom Hanks"

	result, err := neo4j.ExecuteQuery(ctx, driver,
		cypher,
		map[string]any{"name": name}, // Parameters prevent SQL injection
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Found %d records\n", len(result.Records))
	for _, record := range result.Records {
		actorName, _ := record.Get("name")
		movieTitle, _ := record.Get("title")
		role, _ := record.Get("role")
		fmt.Printf("  %s played %s in %s\n", actorName, role, movieTitle)
	}
}

// Example 2: Handling result metadata
func handleResultMetadata(ctx context.Context, driver neo4j.DriverWithContext) {
	fmt.Println("\n=== Example 2: Result Metadata ===")

	cypher := `
		MATCH (m:Movie)
		RETURN m.title AS title, m.released AS released
		ORDER BY m.released DESC
		LIMIT 5
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

	// Access result metadata
	fmt.Printf("Keys: %v\n", result.Keys)
	fmt.Printf("Number of records: %d\n", len(result.Records))

	// Summary information
	if result.Summary != nil {
		fmt.Printf("Statement type: %v\n", result.Summary.StatementType())
		if db := result.Summary.Database(); db != nil {
			fmt.Printf("Database: %s\n", db.Name())
		}
	}

	// Access records
	fmt.Println("\nRecent movies:")
	for _, record := range result.Records {
		title, _ := record.Get("title")
		released, _ := record.Get("released")
		fmt.Printf("  %s (%v)\n", title, released)
	}
}

// Example 3: Custom result transformer
func customTransformer(ctx context.Context, driver neo4j.DriverWithContext) {
	fmt.Println("\n=== Example 3: Custom Result Transformer ===")

	cypher := `
		MATCH (a:Person {name: $name})-[r:ACTED_IN]->(m:Movie)
		RETURN a.name AS name, m.title AS title, r.role AS role
	`

	// First get the result normally, then transform it
	result, err := neo4j.ExecuteQuery(ctx, driver,
		cypher,
		map[string]any{"name": "Tom Hanks"},
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"),
	)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Transform the results manually
	var roleDescriptions []string
	for _, record := range result.Records {
		name, _ := record.Get("name")
		title, _ := record.Get("title")
		role, _ := record.Get("role")
		roleDescriptions = append(roleDescriptions,
			fmt.Sprintf("%s played %s in %s", name, role, title))
	}

	fmt.Println("Roles (transformed):")
	for i, role := range roleDescriptions {
		fmt.Printf("  %d. %s\n", i+1, role)
	}
}

// Example 4: Specifying database
func querySpecificDatabase(ctx context.Context, driver neo4j.DriverWithContext) {
	fmt.Println("\n=== Example 4: Query Specific Database ===")

	cypher := "RETURN 'Hello from database!' AS greeting, COUNT {()} AS nodeCount"

	result, err := neo4j.ExecuteQuery(ctx, driver,
		cypher,
		nil,
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"), // Specify database
	)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if len(result.Records) > 0 {
		greeting, _ := result.Records[0].Get("greeting")
		count, _ := result.Records[0].Get("nodeCount")
		fmt.Printf("%s (Total nodes: %v)\n", greeting, count)
	}
}

// Example 5: Read vs Write routing
func readWriteRouting(ctx context.Context, driver neo4j.DriverWithContext) {
	fmt.Println("\n=== Example 5: Read/Write Routing ===")

	// Read query - distributes across cluster members
	readCypher := `
		MATCH (m:Movie)
		RETURN m.title AS title
		LIMIT 3
	`

	fmt.Println("Read query (optimized for cluster):")
	readResult, err := neo4j.ExecuteQuery(ctx, driver,
		readCypher,
		nil,
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"),
		neo4j.ExecuteQueryWithReadersRouting(), // Use read replicas
	)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for _, record := range readResult.Records {
		title, _ := record.Get("title")
		fmt.Printf("  - %s\n", title)
	}

	// Write query (default behavior - goes to leader)
	writeCypher := `
		CREATE (m:Movie {title: $title, released: $year})
		RETURN m.title AS title
	`

	fmt.Println("\nWrite query (goes to cluster leader):")
	writeResult, err := neo4j.ExecuteQuery(ctx, driver,
		writeCypher,
		map[string]any{
			"title": "Test Movie",
			"year":  2025,
		},
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"),
		// No routing specified = WRITE mode (default)
	)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if len(writeResult.Records) > 0 {
		title, _ := writeResult.Records[0].Get("title")
		fmt.Printf("  Created: %s\n", title)
	}

	// Cleanup - delete the test movie
	_, _ = neo4j.ExecuteQuery(ctx, driver,
		"MATCH (m:Movie {title: $title}) DELETE m",
		map[string]any{"title": "Test Movie"},
		neo4j.EagerResultTransformer,
	)
}

// Example 6: Accessing different data types
func accessDataTypes(ctx context.Context, driver neo4j.DriverWithContext) {
	fmt.Println("\n=== Example 6: Accessing Different Data Types ===")

	cypher := `
		MATCH (m:Movie)
		WHERE m.released > 2000
		RETURN m.title AS title,
		       m.released AS year,
		       m.tagline AS tagline
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

	for _, record := range result.Records {
		title, ok1 := record.Get("title")
		year, ok2 := record.Get("year")
		tagline, ok3 := record.Get("tagline")

		if ok1 && ok2 {
			fmt.Printf("\nMovie: %s (%v)\n", title, year)
			if ok3 && tagline != nil {
				fmt.Printf("  Tagline: %s\n", tagline)
			} else {
				fmt.Printf("  Tagline: (none)\n")
			}
		}
	}
}

// Example 7: Error handling
func errorHandling(ctx context.Context, driver neo4j.DriverWithContext) {
	fmt.Println("\n=== Example 7: Error Handling ===")

	// Intentional error - invalid Cypher
	invalidCypher := "INVALID CYPHER STATEMENT"

	result, err := neo4j.ExecuteQuery(ctx, driver,
		invalidCypher,
		nil,
		neo4j.EagerResultTransformer,
	)

	if err != nil {
		fmt.Printf("✓ Caught error as expected: %v\n", err)
	} else {
		fmt.Printf("Unexpected success: %v\n", result)
	}

	// Valid query
	validCypher := "RETURN 1 AS number"
	result, err = neo4j.ExecuteQuery(ctx, driver,
		validCypher,
		nil,
		neo4j.EagerResultTransformer,
	)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		if len(result.Records) > 0 {
			number, _ := result.Records[0].Get("number")
			fmt.Printf("✓ Valid query executed: %v\n", number)
		}
	}
}

func runCypherExamples() {
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
	executeWithParameters(ctx, driver)
	handleResultMetadata(ctx, driver)
	customTransformer(ctx, driver)
	querySpecificDatabase(ctx, driver)
	readWriteRouting(ctx, driver)
	accessDataTypes(ctx, driver)
	errorHandling(ctx, driver)

	fmt.Println("\n========================================")
	fmt.Println("All examples completed!")
}
