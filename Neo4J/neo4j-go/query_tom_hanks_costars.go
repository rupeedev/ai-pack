package main

import (
	"context"
	"fmt"
	"log"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func queryTomHanksCostars() {
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

	// Execute query to find all actors who worked with Tom Hanks
	result, err := neo4j.ExecuteQuery(ctx, driver,
		`MATCH (tom:Person {name: $name})-[:ACTED_IN]->(m:Movie)<-[:ACTED_IN]-(costar:Person)
		 WHERE tom <> costar
		 RETURN DISTINCT costar.name AS actor,
		                 m.title AS movie
		 ORDER BY actor`,
		map[string]any{"name": "Tom Hanks"},
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n=== Actors Who Worked with Tom Hanks ===\n")

	if len(result.Records) == 0 {
		fmt.Println("No co-stars found.")
		return
	}

	// Group by actor
	actorMovies := make(map[string][]string)
	for _, record := range result.Records {
		actor, _ := record.Get("actor")
		movie, _ := record.Get("movie")
		actorName := actor.(string)
		movieTitle := movie.(string)
		actorMovies[actorName] = append(actorMovies[actorName], movieTitle)
	}

	fmt.Printf("Found %d actors who worked with Tom Hanks:\n\n", len(actorMovies))

	i := 1
	for actor, movies := range actorMovies {
		fmt.Printf("%d. %s\n", i, actor)
		fmt.Printf("   Shared movies: %v\n\n", movies)
		i++
	}

	fmt.Printf("Total co-stars: %d\n", len(actorMovies))
	fmt.Printf("Total collaborations: %d\n", len(result.Records))
}
