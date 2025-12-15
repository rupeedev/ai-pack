package main

import (
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// Example 1: Basic Session Management
func basicSessionManagement(ctx context.Context, driver neo4j.DriverWithContext) {
	fmt.Println("\n=== Example 1: Basic Session Management ===")

	// Create a session
	session := driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx) // Always close session

	fmt.Println("✓ Session created successfully")
	fmt.Println("  Session automatically manages database connections")
	fmt.Println("  Use defer to ensure session is closed when done")
}

// Example 2: Using ExecuteRead for Read Queries
func executeReadExample(ctx context.Context, driver neo4j.DriverWithContext) {
	fmt.Println("\n=== Example 2: Using ExecuteRead ===")

	session := driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	// ExecuteRead optimizes for read operations (can use followers in cluster)
	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		cypher := `
			MATCH (m:Movie)
			RETURN m.title AS title, m.released AS released
			ORDER BY m.released DESC
			LIMIT 3
		`

		result, err := tx.Run(ctx, cypher, nil)
		if err != nil {
			return nil, err
		}

		// Collect results
		var movies []map[string]any
		for result.Next(ctx) {
			record := result.Record()
			movies = append(movies, map[string]any{
				"title":    record.Values[0],
				"released": record.Values[1],
			})
		}

		return movies, result.Err()
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Recent movies:")
	movies := result.([]map[string]any)
	for i, movie := range movies {
		fmt.Printf("  %d. %s (%v)\n", i+1, movie["title"], movie["released"])
	}
}

// Example 3: Using ExecuteWrite for Write Queries
func executeWriteExample(ctx context.Context, driver neo4j.DriverWithContext) {
	fmt.Println("\n=== Example 3: Using ExecuteWrite ===")

	session := driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	// ExecuteWrite optimizes for write operations (uses leader in cluster)
	result, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		cypher := `
			CREATE (p:Person {name: $name, role: $role})
			RETURN p
		`

		result, err := tx.Run(ctx, cypher, map[string]any{
			"name": "Test Person",
			"role": "Developer",
		})
		if err != nil {
			return nil, err
		}

		record, err := result.Single(ctx)
		if err != nil {
			return nil, err
		}

		node, _ := record.Get("p")
		return node, nil
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	person := result.(neo4j.Node)
	fmt.Printf("✓ Created person: %s (role: %s)\n",
		person.Props["name"],
		person.Props["role"])

	// Cleanup
	session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		return tx.Run(ctx,
			"MATCH (p:Person {name: 'Test Person'}) DELETE p",
			nil)
	})
}

// Example 4: Unit of Work Pattern
func unitOfWorkPattern(ctx context.Context, driver neo4j.DriverWithContext) {
	fmt.Println("\n=== Example 4: Unit of Work Pattern ===")

	session := driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	// Define a reusable transaction function
	createPerson := func(tx neo4j.ManagedTransaction, name string, age int64) (neo4j.Node, error) {
		result, err := tx.Run(ctx, `
			CREATE (p:Person {name: $name, age: $age})
			RETURN p
		`, map[string]any{"name": name, "age": age})

		if err != nil {
			return neo4j.Node{}, err
		}

		record, err := result.Single(ctx)
		if err != nil {
			return neo4j.Node{}, err
		}

		node, _ := record.Get("p")
		return node.(neo4j.Node), nil
	}

	// Use the transaction function
	result, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		return createPerson(tx, "Alice", 30)
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	person := result.(neo4j.Node)
	fmt.Printf("✓ Created person: %s (age: %v)\n",
		person.Props["name"],
		person.Props["age"])
	fmt.Println("  Unit of work pattern allows reusable transaction functions")

	// Cleanup
	session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		return tx.Run(ctx,
			"MATCH (p:Person {name: 'Alice'}) DELETE p",
			nil)
	})
}

// Example 5: Multiple Queries in One Transaction
func multipleQueriesInTransaction(ctx context.Context, driver neo4j.DriverWithContext) {
	fmt.Println("\n=== Example 5: Multiple Queries in One Transaction ===")

	session := driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	// Setup: Create two accounts
	session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		tx.Run(ctx, `
			CREATE (a1:Account {id: 'ACC001', balance: 1000.0})
			CREATE (a2:Account {id: 'ACC002', balance: 500.0})
		`, nil)
		return nil, nil
	})

	// Transfer funds (multiple queries, single transaction)
	transferFunds := func(tx neo4j.ManagedTransaction, fromAccount, toAccount string, amount float64) error {
		// Deduct from first account
		_, err := tx.Run(ctx,
			"MATCH (a:Account {id: $from}) SET a.balance = a.balance - $amount",
			map[string]any{"from": fromAccount, "amount": amount},
		)
		if err != nil {
			return err
		}

		// Add to second account
		_, err = tx.Run(ctx,
			"MATCH (a:Account {id: $to}) SET a.balance = a.balance + $amount",
			map[string]any{"to": toAccount, "amount": amount},
		)
		return err
	}

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		return nil, transferFunds(tx, "ACC001", "ACC002", 200.0)
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Verify balances
	result, _ := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := tx.Run(ctx, `
			MATCH (a:Account)
			WHERE a.id IN ['ACC001', 'ACC002']
			RETURN a.id AS id, a.balance AS balance
			ORDER BY a.id
		`, nil)
		if err != nil {
			return nil, err
		}

		var accounts []map[string]any
		for result.Next(ctx) {
			record := result.Record()
			accounts = append(accounts, map[string]any{
				"id":      record.Values[0],
				"balance": record.Values[1],
			})
		}
		return accounts, result.Err()
	})

	fmt.Println("✓ Funds transferred successfully")
	fmt.Println("Account balances after transfer:")
	accounts := result.([]map[string]any)
	for _, account := range accounts {
		fmt.Printf("  %s: $%.2f\n", account["id"], account["balance"])
	}
	fmt.Println("  Both operations completed or none (atomic transaction)")

	// Cleanup
	session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		return tx.Run(ctx,
			"MATCH (a:Account) WHERE a.id IN ['ACC001', 'ACC002'] DELETE a",
			nil)
	})
}

// Example 6: Transaction Rollback on Error
func transactionRollbackExample(ctx context.Context, driver neo4j.DriverWithContext) {
	fmt.Println("\n=== Example 6: Transaction Rollback on Error ===")

	session := driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	// Cleanup any existing "Rollback Test Actor" from previous runs
	session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		return tx.Run(ctx,
			"MATCH (p:Person {name: 'Rollback Test Actor'}) DELETE p",
			nil)
	})

	addActorToMovie := func(tx neo4j.ManagedTransaction, actorName string, movieTitle string, role string) error {
		// Create actor
		_, err := tx.Run(ctx, `
			CREATE (a:Person {name: $name})
		`, map[string]any{"name": actorName})
		if err != nil {
			return err // Transaction rolls back
		}

		fmt.Println("  Actor created in transaction...")

		// Try to create acting relationship
		// Using MERGE on the MATCH pattern which will fail if movie doesn't exist
		result, err := tx.Run(ctx, `
			MATCH (a:Person {name: $actorName})
			MATCH (m:Movie {title: $movieTitle})
			CREATE (a)-[:ACTED_IN {role: $role}]->(m)
			RETURN m
		`, map[string]any{
			"actorName":  actorName,
			"movieTitle": movieTitle,
			"role":       role,
		})

		if err != nil {
			return err // Transaction rolls back - actor creation is also undone!
		}

		// Check if relationship was actually created
		if !result.Next(ctx) {
			return fmt.Errorf("movie not found: %s", movieTitle)
		}

		return nil // Transaction commits
	}

	// Test with non-existent movie (will fail and rollback)
	fmt.Println("Attempting to add actor to non-existent movie...")
	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		return nil, addActorToMovie(tx, "Rollback Test Actor", "Non-Existent Movie", "Hero")
	})

	if err != nil {
		fmt.Printf("✓ Transaction rolled back as expected\n")
		fmt.Println("  Actor creation was undone (rollback)")
	} else {
		fmt.Println("✗ Transaction should have failed!")
	}

	// Verify actor was NOT created
	result, _ := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := tx.Run(ctx,
			"MATCH (p:Person {name: 'Rollback Test Actor'}) RETURN count(p) AS count",
			nil)
		if err != nil {
			return nil, err
		}
		record, err := result.Single(ctx)
		if err != nil {
			return nil, err
		}
		count, _ := record.Get("count")
		return count, nil
	})

	count := result.(int64)
	if count == 0 {
		fmt.Printf("✓ Verified: Actor was NOT created (count: %d)\n", count)
	} else {
		fmt.Printf("✗ Unexpected: Actor exists (count: %d)\n", count)
	}
}

// Example 7: Result Consumption and Summary
func resultConsumptionExample(ctx context.Context, driver neo4j.DriverWithContext) {
	fmt.Println("\n=== Example 7: Result Consumption and Summary ===")

	session := driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	// Execute a write query and get summary
	summary, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := tx.Run(ctx, `
			CREATE (p1:Person {name: 'Bob'}),
			       (p2:Person {name: 'Carol'}),
			       (p1)-[:KNOWS]->(p2)
			RETURN p1, p2
		`, nil)
		if err != nil {
			return nil, err
		}

		// Consume result and get summary
		return result.Consume(ctx)
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	summaryObj := summary.(neo4j.ResultSummary)
	fmt.Println("Transaction Summary:")
	fmt.Printf("  Nodes created: %d\n", summaryObj.Counters().NodesCreated())
	fmt.Printf("  Relationships created: %d\n", summaryObj.Counters().RelationshipsCreated())
	fmt.Printf("  Properties set: %d\n", summaryObj.Counters().PropertiesSet())
	fmt.Printf("  Results available after: %d ms\n", summaryObj.ResultAvailableAfter().Milliseconds())
	fmt.Printf("  Results consumed after: %d ms\n", summaryObj.ResultConsumedAfter().Milliseconds())
	fmt.Printf("  Statement type: %s\n", summaryObj.StatementType())

	// Cleanup
	session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		return tx.Run(ctx,
			"MATCH (p:Person) WHERE p.name IN ['Bob', 'Carol'] DETACH DELETE p",
			nil)
	})
}

// Example 8: Specifying a Database
func specifyingDatabaseExample(ctx context.Context, driver neo4j.DriverWithContext) {
	fmt.Println("\n=== Example 8: Specifying a Database ===")

	// Create session for specific database
	session := driver.NewSession(ctx, neo4j.SessionConfig{
		DatabaseName: "neo4j", // Specify database name
	})
	defer session.Close(ctx)

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := tx.Run(ctx,
			"RETURN COUNT {()} AS count",
			nil)
		if err != nil {
			return nil, err
		}
		record, err := result.Single(ctx)
		if err != nil {
			return nil, err
		}
		count, _ := record.Get("count")
		return count, nil
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	count := result.(int64)
	fmt.Printf("✓ Queried 'neo4j' database: %d nodes\n", count)
	fmt.Println("  Use SessionConfig.DatabaseName to specify database")
}

// Example 9: Transient Error Handling
func transientErrorHandling(ctx context.Context, driver neo4j.DriverWithContext) {
	fmt.Println("\n=== Example 9: Transient Error Handling ===")

	session := driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	fmt.Println("ExecuteRead/ExecuteWrite automatically retry on transient errors:")
	fmt.Println("  - Network issues")
	fmt.Println("  - Temporary database unavailability")
	fmt.Println("  - Cluster leader changes")
	fmt.Println("  - Lock timeouts")
	fmt.Println("")
	fmt.Println("✓ Your transaction functions are resilient by default!")
}

// Example 10: Best Practices Summary
func bestPracticesSummary(ctx context.Context, driver neo4j.DriverWithContext) {
	fmt.Println("\n=== Example 10: Best Practices Summary ===")
	fmt.Println("")
	fmt.Println("Transaction Management Best Practices:")
	fmt.Println("")
	fmt.Println("1. Always close sessions:")
	fmt.Println("   session := driver.NewSession(ctx, neo4j.SessionConfig{})")
	fmt.Println("   defer session.Close(ctx)")
	fmt.Println("")
	fmt.Println("2. Use ExecuteRead for read-only queries:")
	fmt.Println("   - Optimized for read operations")
	fmt.Println("   - Can use cluster followers")
	fmt.Println("")
	fmt.Println("3. Use ExecuteWrite for write queries:")
	fmt.Println("   - Routes to cluster leader")
	fmt.Println("   - Ensures consistency")
	fmt.Println("")
	fmt.Println("4. Keep transactions short:")
	fmt.Println("   - Transaction state is held in memory")
	fmt.Println("   - Break large operations into smaller transactions")
	fmt.Println("")
	fmt.Println("5. Use unit of work pattern:")
	fmt.Println("   - Create reusable transaction functions")
	fmt.Println("   - Pass parameters for flexibility")
	fmt.Println("")
	fmt.Println("6. Automatic rollback on error:")
	fmt.Println("   - Return error to rollback transaction")
	fmt.Println("   - All changes are undone atomically")
	fmt.Println("")
	fmt.Println("7. Automatic retry on transient errors:")
	fmt.Println("   - Network issues handled automatically")
	fmt.Println("   - No manual retry logic needed")
	fmt.Println("")
	fmt.Println("8. Consume results within transaction:")
	fmt.Println("   - Use result.Consume(ctx) for summary")
	fmt.Println("   - Process records immediately")
}

func runBestPracticesExamples() {
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
	basicSessionManagement(ctx, driver)
	executeReadExample(ctx, driver)
	executeWriteExample(ctx, driver)
	unitOfWorkPattern(ctx, driver)
	multipleQueriesInTransaction(ctx, driver)
	transactionRollbackExample(ctx, driver)
	resultConsumptionExample(ctx, driver)
	specifyingDatabaseExample(ctx, driver)
	transientErrorHandling(ctx, driver)
	bestPracticesSummary(ctx, driver)

	fmt.Println("\n========================================")
	fmt.Println("All best practices examples completed!")
}
