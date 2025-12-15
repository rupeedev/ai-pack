package main

import (
	"context"
	"fmt"
	"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func basicExample(ctx context.Context, driver neo4j.DriverWithContext) {
	fmt.Println("\n=== Basic Neo4j Connection Example ===")

	// Run first query - Count all nodes
	result, err := neo4j.ExecuteQuery(ctx, driver,
		"RETURN 'Hello, Neo4j!' AS message, COUNT {()} AS count",
		nil,                          // No parameters
		neo4j.EagerResultTransformer, // Load all results into memory
	)
	if err != nil {
		panic(err)
	}

	// Get the first record
	if len(result.Records) > 0 {
		first := result.Records[0]

		// Print the message
		message, _ := first.Get("message")
		fmt.Println(message)

		// Print the count of nodes
		count, _ := first.Get("count")
		fmt.Printf("Total nodes in database: %v\n", count)
	}

	// Example: Querying a specific database
	result2, err := neo4j.ExecuteQuery(ctx, driver,
		"RETURN COUNT {()} AS count",
		nil,
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"), // Query the 'neo4j' database
	)
	if err != nil {
		panic(err)
	}

	if len(result2.Records) > 0 {
		count, _ := result2.Records[0].Get("count")
		fmt.Printf("Nodes in 'neo4j' database: %v\n", count)
	}
}

func printUsage() {
	fmt.Println("Neo4j Go Examples")
	fmt.Println("=================")
	fmt.Println("\nUsage: go run . [command]")
	fmt.Println("\nCommands:")
	fmt.Println("  basic        - Run basic connection example (default)")
	fmt.Println("  setup        - Setup sample data (Movies & Actors)")
	fmt.Println("  examples     - Run all Cypher execution examples")
	fmt.Println("  results      - Run result handling examples (Nodes, Paths, etc.)")
	fmt.Println("  temporal     - Run temporal & spatial data examples (Dates, Points)")
	fmt.Println("  transactions - Run transaction management examples (Sessions, Units of Work)")
	fmt.Println("  movies       - Query and display all movies in the database")
	fmt.Println("  help         - Show this help message")
	fmt.Println("\nExamples:")
	fmt.Println("  go run .                 # Run basic example")
	fmt.Println("  go run . setup           # Setup sample data")
	fmt.Println("  go run . examples        # Run all Cypher examples")
	fmt.Println("  go run . results         # Run result handling examples")
	fmt.Println("  go run . temporal        # Run temporal & spatial examples")
	fmt.Println("  go run . transactions    # Run transaction management examples")
}

func main() {
	// Parse command line arguments
	command := "basic"
	if len(os.Args) > 1 {
		command = os.Args[1]
	}

	if command == "help" {
		printUsage()
		return
	}

	// Create a driver instance
	driver, err := neo4j.NewDriverWithContext(
		"neo4j://localhost:7687",                         // Connection string
		neo4j.BasicAuth("neo4j", "Your@Password!@#", ""), // Authentication
	)
	if err != nil {
		panic(err)
	}
	defer driver.Close(context.Background()) // Always close the driver when done

	// Verify connectivity
	ctx := context.Background()
	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Neo4j: %v\n\nMake sure Neo4j is running:\n  docker ps | grep neo4j\n", err))
	}
	fmt.Println("✓ Connected to Neo4j successfully!")

	// Run the appropriate command
	switch command {
	case "basic":
		basicExample(ctx, driver)

	case "setup":
		err := setupSampleData(ctx, driver)
		if err != nil {
			panic(err)
		}
		fmt.Println("\nTip: Now run 'go run . examples' to see the data in action!")

	case "examples":
		runCypherExamples()

	case "results":
		runResultHandlingExamples()

	case "temporal":
		runTemporalSpatialExamples()

	case "transactions":
		runBestPracticesExamples()

	case "movies":
		queryAllMovies()

	case "costars":
		queryTomHanksCostars()

	default:
		fmt.Printf("Unknown command: %s\n\n", command)
		printUsage()
		os.Exit(1)
	}

	fmt.Println("\n✓ Done!")
}
