package main

import (
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// setupSampleData creates sample Movie and Person nodes with relationships
func setupSampleData(ctx context.Context, driver neo4j.DriverWithContext) error {
	fmt.Println("\n=== Setting up sample data ===")

	// Clear existing data (optional - uncomment if needed)
	// clearCypher := "MATCH (n) DETACH DELETE n"
	// _, err := neo4j.ExecuteQuery(ctx, driver, clearCypher, nil, neo4j.EagerResultTransformer)
	// if err != nil {
	// 	return fmt.Errorf("failed to clear data: %w", err)
	// }

	// Create Movies
	moviesCypher := `
		MERGE (m1:Movie {title: "The Matrix", released: 1999, tagline: "Welcome to the Real World"})
		MERGE (m2:Movie {title: "The Matrix Reloaded", released: 2003, tagline: "Free your mind"})
		MERGE (m3:Movie {title: "The Matrix Revolutions", released: 2003})
		MERGE (m4:Movie {title: "Forrest Gump", released: 1994, tagline: "Life is like a box of chocolates"})
		MERGE (m5:Movie {title: "Toy Story", released: 1995, tagline: "To infinity and beyond!"})
		MERGE (m6:Movie {title: "Cast Away", released: 2000, tagline: "At the edge of the world, his journey begins"})
		MERGE (m7:Movie {title: "Apollo 13", released: 1995, tagline: "Houston, we have a problem"})
		MERGE (m8:Movie {title: "The Da Vinci Code", released: 2006})
		MERGE (m9:Movie {title: "Cloud Atlas", released: 2012, tagline: "Everything is connected"})
		MERGE (m10:Movie {title: "The Green Mile", released: 1999})
		RETURN count(*) AS moviesCreated
	`

	result, err := neo4j.ExecuteQuery(ctx, driver,
		moviesCypher,
		nil,
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return fmt.Errorf("failed to create movies: %w", err)
	}

	if len(result.Records) > 0 {
		count, _ := result.Records[0].Get("moviesCreated")
		fmt.Printf("✓ Movies created/verified: %v\n", count)
	}

	// Create People
	peopleCypher := `
		MERGE (p1:Person {name: "Tom Hanks", born: 1956})
		MERGE (p2:Person {name: "Keanu Reeves", born: 1964})
		MERGE (p3:Person {name: "Carrie-Anne Moss", born: 1967})
		MERGE (p4:Person {name: "Laurence Fishburne", born: 1961})
		MERGE (p5:Person {name: "Hugo Weaving", born: 1960})
		MERGE (p6:Person {name: "Lilly Wachowski", born: 1967})
		MERGE (p7:Person {name: "Lana Wachowski", born: 1965})
		MERGE (p8:Person {name: "Robert Zemeckis", born: 1951})
		MERGE (p9:Person {name: "Gary Sinise", born: 1955})
		MERGE (p10:Person {name: "Robin Wright", born: 1966})
		RETURN count(*) AS peopleCreated
	`

	result, err = neo4j.ExecuteQuery(ctx, driver,
		peopleCypher,
		nil,
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return fmt.Errorf("failed to create people: %w", err)
	}

	if len(result.Records) > 0 {
		count, _ := result.Records[0].Get("peopleCreated")
		fmt.Printf("✓ People created/verified: %v\n", count)
	}

	// Create relationships - Tom Hanks movies
	tomHanksRelsCypher := `
		MATCH (tom:Person {name: "Tom Hanks"})
		MATCH (forrest:Movie {title: "Forrest Gump"})
		MATCH (toy:Movie {title: "Toy Story"})
		MATCH (cast:Movie {title: "Cast Away"})
		MATCH (apollo:Movie {title: "Apollo 13"})
		MATCH (green:Movie {title: "The Green Mile"})
		MATCH (cloud:Movie {title: "Cloud Atlas"})
		MATCH (davinci:Movie {title: "The Da Vinci Code"})

		MERGE (tom)-[:ACTED_IN {role: "Forrest Gump"}]->(forrest)
		MERGE (tom)-[:ACTED_IN {role: "Woody"}]->(toy)
		MERGE (tom)-[:ACTED_IN {role: "Chuck Noland"}]->(cast)
		MERGE (tom)-[:ACTED_IN {role: "Jim Lovell"}]->(apollo)
		MERGE (tom)-[:ACTED_IN {role: "Paul Edgecomb"}]->(green)
		MERGE (tom)-[:ACTED_IN {role: "Zachry"}]->(cloud)
		MERGE (tom)-[:ACTED_IN {role: "Robert Langdon"}]->(davinci)

		RETURN count(*) AS relationshipsCreated
	`

	result, err = neo4j.ExecuteQuery(ctx, driver,
		tomHanksRelsCypher,
		nil,
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return fmt.Errorf("failed to create Tom Hanks relationships: %w", err)
	}

	if len(result.Records) > 0 {
		count, _ := result.Records[0].Get("relationshipsCreated")
		fmt.Printf("✓ Tom Hanks relationships created: %v\n", count)
	}

	// Create relationships - Matrix movies
	matrixRelsCypher := `
		MATCH (keanu:Person {name: "Keanu Reeves"})
		MATCH (carrie:Person {name: "Carrie-Anne Moss"})
		MATCH (laurence:Person {name: "Laurence Fishburne"})
		MATCH (hugo:Person {name: "Hugo Weaving"})
		MATCH (lilly:Person {name: "Lilly Wachowski"})
		MATCH (lana:Person {name: "Lana Wachowski"})

		MATCH (matrix1:Movie {title: "The Matrix"})
		MATCH (matrix2:Movie {title: "The Matrix Reloaded"})
		MATCH (matrix3:Movie {title: "The Matrix Revolutions"})

		MERGE (keanu)-[:ACTED_IN {role: "Neo"}]->(matrix1)
		MERGE (keanu)-[:ACTED_IN {role: "Neo"}]->(matrix2)
		MERGE (keanu)-[:ACTED_IN {role: "Neo"}]->(matrix3)

		MERGE (carrie)-[:ACTED_IN {role: "Trinity"}]->(matrix1)
		MERGE (carrie)-[:ACTED_IN {role: "Trinity"}]->(matrix2)
		MERGE (carrie)-[:ACTED_IN {role: "Trinity"}]->(matrix3)

		MERGE (laurence)-[:ACTED_IN {role: "Morpheus"}]->(matrix1)
		MERGE (laurence)-[:ACTED_IN {role: "Morpheus"}]->(matrix2)
		MERGE (laurence)-[:ACTED_IN {role: "Morpheus"}]->(matrix3)

		MERGE (hugo)-[:ACTED_IN {role: "Agent Smith"}]->(matrix1)
		MERGE (hugo)-[:ACTED_IN {role: "Agent Smith"}]->(matrix2)
		MERGE (hugo)-[:ACTED_IN {role: "Agent Smith"}]->(matrix3)

		MERGE (lilly)-[:DIRECTED]->(matrix1)
		MERGE (lana)-[:DIRECTED]->(matrix1)
		MERGE (lilly)-[:DIRECTED]->(matrix2)
		MERGE (lana)-[:DIRECTED]->(matrix2)
		MERGE (lilly)-[:DIRECTED]->(matrix3)
		MERGE (lana)-[:DIRECTED]->(matrix3)

		RETURN count(*) AS relationshipsCreated
	`

	result, err = neo4j.ExecuteQuery(ctx, driver,
		matrixRelsCypher,
		nil,
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return fmt.Errorf("failed to create Matrix relationships: %w", err)
	}

	if len(result.Records) > 0 {
		count, _ := result.Records[0].Get("relationshipsCreated")
		fmt.Printf("✓ Matrix relationships created: %v\n", count)
	}

	// Create relationships - Forrest Gump
	forrestGumpRelsCypher := `
		MATCH (zemeckis:Person {name: "Robert Zemeckis"})
		MATCH (gary:Person {name: "Gary Sinise"})
		MATCH (robin:Person {name: "Robin Wright"})
		MATCH (forrest:Movie {title: "Forrest Gump"})

		MERGE (zemeckis)-[:DIRECTED]->(forrest)
		MERGE (gary)-[:ACTED_IN {role: "Lieutenant Dan Taylor"}]->(forrest)
		MERGE (robin)-[:ACTED_IN {role: "Jenny Curran"}]->(forrest)

		RETURN count(*) AS relationshipsCreated
	`

	result, err = neo4j.ExecuteQuery(ctx, driver,
		forrestGumpRelsCypher,
		nil,
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return fmt.Errorf("failed to create Forrest Gump relationships: %w", err)
	}

	if len(result.Records) > 0 {
		count, _ := result.Records[0].Get("relationshipsCreated")
		fmt.Printf("✓ Forrest Gump relationships created: %v\n", count)
	}

	// Verify data
	verifyCypher := `
		MATCH (m:Movie)
		OPTIONAL MATCH (m)<-[r:ACTED_IN]-(p:Person)
		RETURN m.title AS movie, count(p) AS actors
		ORDER BY actors DESC
		LIMIT 5
	`

	result, err = neo4j.ExecuteQuery(ctx, driver,
		verifyCypher,
		nil,
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return fmt.Errorf("failed to verify data: %w", err)
	}

	fmt.Println("\nTop movies by actor count:")
	for _, record := range result.Records {
		movie, _ := record.Get("movie")
		actors, _ := record.Get("actors")
		fmt.Printf("  - %s: %v actors\n", movie, actors)
	}

	fmt.Println("\n✓ Sample data setup complete!")
	return nil
}
