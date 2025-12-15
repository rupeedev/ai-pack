package main

import (
	"context"
	"fmt"
	"log"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func queryAllMovies() {
	// Create a driver instance
	driver, err := neo4j.NewDriverWithContext(
		"neo4j://localhost:7687",
		neo4j.BasicAuth("neo4j", "Your@Password!@#", ""),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer driver.Close(context.Background())

	ctx := context.Background()

	// Execute query to get all movies
	result, err := neo4j.ExecuteQuery(ctx, driver,
		`MATCH (m:Movie)
		 RETURN m.title AS title,
		        m.released AS released,
		        m.tagline AS tagline
		 ORDER BY m.released DESC`,
		nil,
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n=== All Movies in Neo4j Database ===\n")

	if len(result.Records) == 0 {
		fmt.Println("No movies found in the database.")
		return
	}

	fmt.Printf("Found %d movies:\n\n", len(result.Records))

	for i, record := range result.Records {
		title, _ := record.Get("title")
		released, _ := record.Get("released")
		tagline, _ := record.Get("tagline")

		fmt.Printf("%d. %s (%v)\n", i+1, title, released)
		if tagline != nil && tagline != "" {
			fmt.Printf("   Tagline: %s\n", tagline)
		}
		fmt.Println()
	}

	// Get total count and some stats
	fmt.Printf("Total movies: %d\n", len(result.Records))
}
