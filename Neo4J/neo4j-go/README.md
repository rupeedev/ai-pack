# Neo4j Go Tutorial

Learn to build apps with Neo4j (a graph database) using Go.

## What is Neo4j?

Neo4j stores data as **nodes** (things) and **relationships** (connections between things).

Example:

```
Tom Hanks --[acted in {role: "Woody"}]--> Toy Story
```

Instead of tables like regular databases, you get a network of connected data.

**Why use a graph database?**

- Finding connections is super fast (friends of friends, recommendations)
- Data looks like the real world (people know people, actors act in movies)
- Easy to understand (just draw circles and arrows!)

## What You'll Learn

This tutorial teaches you:

1. ✅ How to connect Go to Neo4j
2. ✅ How to run queries (finding data)
3. ✅ How to handle different data types
4. ✅ How to work with nodes, relationships, and paths
5. ✅ How to handle dates, times, and locations
6. ✅ How to manage transactions for production apps

## What You Need

- Go installed (version 1.24.4+)
- Docker installed
- A terminal/command line
- 15 minutes of your time

## Quick Start (5 Steps)

### Step 1: Start the Database

Open your terminal and run:

```bash
docker run -d \
  --name neo4j \
  -p 7474:7474 -p 7687:7687 \
  -e NEO4J_AUTH=neo4j/Your@Password!@# \
  neo4j:latest
```

**What this does:** Downloads and starts Neo4j in a container.

**Wait 10 seconds**, then check it's running:

```bash
docker ps | grep neo4j
```

You should see "neo4j" in the output.

### Step 2: View the Database

Open your browser: http://localhost:7474

- Username: `neo4j`
- Password: `Your@Password!@#`

You'll see an empty database. That's normal!

### Step 3: Run the Code

In this folder, run:

```bash
go run .
```

You should see:

```
✓ Connected to Neo4j successfully!
Hello, Neo4j!
Total nodes in database: 0
```

**What happened:** Your Go program connected to Neo4j and ran a simple query.

### Step 4: Add Sample Data

Run:

```bash
go run . setup
```

This creates:

- **10 movies**: The Matrix, Forrest Gump, Toy Story, etc.
- **10 people**: Tom Hanks, Keanu Reeves, directors
- **Connections**: Who acted in what, who directed what

You'll see:

```
✓ Movies created/verified: 10
✓ People created/verified: 10
✓ Tom Hanks relationships created: 7
```

Now go back to http://localhost:7474 and type:

```cypher
MATCH (n) RETURN n LIMIT 25
```

Click the play button. You'll see a **graph visualization**!

### Step 5: Run Examples

Try these commands:

```bash
# Query examples (7 examples)
go run . examples

# Result handling examples (7 examples)
go run . results

# Temporal & spatial examples (10 examples)
go run . temporal

# Transaction management examples (10 examples)
go run . transactions
```

Watch the output to see what each example does.

## All Available Commands

| Command                   | What It Does                    | When to Use                     |
| ------------------------- | ------------------------------- | ------------------------------- |
| `go run .`              | Test connection                 | First time setup                |
| `go run . setup`        | Add sample movies/actors        | Before running examples         |
| `go run . examples`     | Run query examples              | Learn how to write queries      |
| `go run . results`      | Run result handling examples    | Learn data types                |
| `go run . temporal`     | Run temporal & spatial examples | Learn dates, times, coordinates |
| `go run . transactions` | Run transaction examples        | Learn production best practices |
| `go run . help`         | Show help menu                  | When you forget commands        |

## Understanding the Examples

### Query Examples (`go run . examples`)

**Example 1: Find Tom Hanks Movies**

```
Tom Hanks played Woody in Toy Story
Tom Hanks played Forrest Gump in Forrest Gump
...
```

**What it teaches:** How to search with parameters (safe from hackers)

**Example 2: Get Movie Details**

```
Keys: [title released]
Recent movies:
  Cloud Atlas (2012)
  The Da Vinci Code (2006)
```

**What it teaches:** How to access result metadata and summaries

**Example 3: Format Results**

```
1. Tom Hanks played Woody in Toy Story
2. Tom Hanks played Forrest Gump in Forrest Gump
```

**What it teaches:** How to transform raw data into readable text

**Example 4: Query a Specific Database**

```
Hello from database! (Total nodes: 20)
```

**What it teaches:** Neo4j can have multiple databases

**Example 5: Read vs Write**

```
Read query (optimized for cluster):
  - The Matrix
Write query (goes to cluster leader):
  Created: Test Movie
```

**What it teaches:** Different routing for reading vs changing data

**Example 6: Handle Different Data**

```
Movie: The Matrix Reloaded (2003)
  Tagline: Free your mind

Movie: The Da Vinci Code (2006)
  Tagline: (none)
```

**What it teaches:** Some data might be missing - handle it gracefully

**Example 7: Handle Errors**

```
✓ Caught error as expected: Neo.ClientError.Statement.SyntaxError
✓ Valid query executed: 1
```

**What it teaches:** How to catch and handle errors properly

### Result Handling Examples (`go run . results`)

**Example 1: Working with Nodes**

```
Element ID: 4:ddcaa0e1-899f-4e05-9843-2946a94dd0c1:4
Labels: [Movie]
Title: Toy Story
Released: 1995
Tagline: To infinity and beyond!
```

**What it teaches:** Nodes have IDs, labels, and properties

**Example 2: Working with Relationships**

```
Relationship Details:
  Type: ACTED_IN
  Role: Woody

Tom Hanks -[ACTED_IN]-> Toy Story
```

**What it teaches:** Relationships connect nodes and have properties

**Example 3: Working with Paths**

```
Nodes in path:
  0. Person: Hugo Weaving
  1. Movie: The Matrix

Relationships in path:
  0. ACTED_IN (role: Agent Smith)
```

**What it teaches:** Paths are sequences of connected nodes

**Example 4: Working with Multiple Paths**

```
1. Hugo Weaving played Agent Smith in The Matrix
2. Laurence Fishburne played Morpheus in The Matrix
3. Carrie-Anne Moss played Trinity in The Matrix
```

**What it teaches:** How to process many paths efficiently

**Example 5: Working with Dates**

```
Born (Date): 1990-05-15
Born (formatted): May 15, 1990
Created At: 2025-12-09 06:42:10
```

**What it teaches:** How to convert Neo4j dates to Go time

**Example 6: Working with Locations**

```
Office: (lat: 37.7749, lng: -122.4194)
Home: (lat: 37.7849, lng: -122.4094)
Distance: 1418.91 meters
```

**What it teaches:** Store coordinates and calculate distances

**Example 7: Type Mapping**

```
Neo4j Type → Go Type:
  Integer → int64 (42)
  Float → float64 (3.14)
  String → string (Hello)
  Date → dbtype.Date (2025-01-01)
```

**What it teaches:** How Neo4j types map to Go types

### Temporal & Spatial Examples (`go run . temporal`)

**Example 1: Writing Temporal Types**

```
Created event: GraphConnect 2024
  Starts at: 2024-05-15 14:30:00
  Created at: 2024-05-15 14:30:00
```

**What it teaches:** Create events with dates and times

**Example 2: Reading Temporal Types**

```
Date: 2025-12-09 (type: dbtype.Date)
Time: 07:49:02.889Z (type: dbtype.Time)
DateTime: 2025-12-09 07:49:02.889 +0000 (type: time.Time)
```

**What it teaches:** All temporal type variations Neo4j supports

**Example 3: Working with Durations**

```
Event: Workshop
  Duration: P0M0DT5400S
  Interval: P0M0DT5400S
```

**What it teaches:** Create time periods (1 hour 30 minutes)

**Example 4: Date Arithmetic**

```
Base date: 2024-01-15
  + 1 month: 2024-02-15
  + 1 week: 2024-01-22
  - 1 year: 2023-01-15
```

**What it teaches:** Add and subtract time periods

**Example 5-6: Cartesian Points (2D & 3D)**

```
2D: X: 1.23, Y: 4.56, SRID: 7203
3D: X: 1.23, Y: 4.56, Z: 7.89, SRID: 9157
```

**What it teaches:** Store coordinates in a flat coordinate system

**Example 7-8: Geographic Points (2D & 3D)**

```
London: Lat: 51.5099, Lng: -0.1181 (SRID: 4326)
The Shard: Lat: 51.504501, Lng: -0.086500, Height: 310m (SRID: 4979)
```

**What it teaches:** Store real-world locations with latitude/longitude

**Example 9: Calculating Distances**

```
London to Paris: 343.82 kilometers
```

**What it teaches:** Calculate distance between two points

**Example 10: Finding Nearby Locations**

```
Cities within 50km of San Francisco:
  Oakland: 13.44 km
```

**What it teaches:** Find locations within a radius

### Transaction Management Examples (`go run . transactions`)

**Example 1: Basic Session Management**

```
✓ Session created successfully
  Session automatically manages database connections
  Use defer to ensure session is closed when done
```

**What it teaches:** Create and manage database sessions

**Example 2: Using ExecuteRead**

```
Recent movies:
  1. Cloud Atlas (2012)
  2. The Da Vinci Code (2006)
```

**What it teaches:** Read-only queries (optimized for clusters)

**Example 3: Using ExecuteWrite**

```
✓ Created person: Test Person (role: Developer)
```

**What it teaches:** Write queries (routed to cluster leader)

**Example 4: Unit of Work Pattern**

```
✓ Created person: Alice (age: 30)
  Unit of work pattern allows reusable transaction functions
```

**What it teaches:** Create reusable transaction functions

**Example 5: Multiple Queries in One Transaction**

```
✓ Funds transferred successfully
Account balances after transfer:
  ACC001: $800.00
  ACC002: $700.00
  Both operations completed or none (atomic transaction)
```

**What it teaches:** All operations succeed together or all fail (atomic)

**Example 6: Transaction Rollback on Error**

```
Attempting to add actor to non-existent movie...
  Actor created in transaction...
✓ Transaction rolled back as expected
✓ Verified: Actor was NOT created (count: 0)
```

**What it teaches:** Errors undo ALL changes automatically

**Example 7: Result Consumption and Summary**

```
Transaction Summary:
  Nodes created: 2
  Relationships created: 1
  Results available after: 21 ms
```

**What it teaches:** Get transaction statistics and timing

**Example 8: Specifying a Database**

```
✓ Queried 'neo4j' database: 21 nodes
  Use SessionConfig.DatabaseName to specify database
```

**What it teaches:** Work with specific databases in multi-database setup

**Example 9: Transient Error Handling**

```
ExecuteRead/ExecuteWrite automatically retry on transient errors:
  - Network issues
  - Cluster leader changes
```

**What it teaches:** Automatic retry for temporary failures

**Example 10: Best Practices Summary**

- Always close sessions with defer
- Use ExecuteRead for queries, ExecuteWrite for changes
- Keep transactions short (memory is limited)
- Errors trigger automatic rollback
- Transient errors retry automatically

## File Structure

```
hello-neo4j/
├── main.go              # Starting point (command router)
├── cypher_examples.go   # 7 query examples
├── result_handling.go   # 7 result handling examples
├── temporal_spatial.go  # 10 temporal & spatial examples
├── best_practices.go    # 10 transaction management examples
├── setup_data.go        # Creates sample data
├── go.mod               # Dependencies
└── README.md            # This file
```

## Try It Yourself

### In the Browser (http://localhost:7474)

**Query 1: Find All Movies**

```cypher
MATCH (m:Movie)
RETURN m.title, m.released
ORDER BY m.released DESC
```

**Query 2: Find Who Acted in The Matrix**

```cypher
MATCH (p:Person)-[:ACTED_IN]->(m:Movie {title: "The Matrix"})
RETURN p.name, m.title
```

**Query 3: Find Directors**

```cypher
MATCH (d:Person)-[:DIRECTED]->(m:Movie)
RETURN d.name AS director, m.title AS movie
```

**Query 4: Find Actors Who Worked Together**

```cypher
MATCH (p1:Person)-[:ACTED_IN]->(m:Movie)<-[:ACTED_IN]-(p2:Person)
WHERE p1.name = "Tom Hanks" AND p1 <> p2
RETURN DISTINCT p2.name AS costar, m.title AS movie
```

**Query 5: Find Paths Between Actors**

```cypher
MATCH path = shortestPath(
  (keanu:Person {name: "Keanu Reeves"})-[*]-(tom:Person {name: "Tom Hanks"})
)
RETURN path
```

### In Your Code

**Modify Example 1:**
Change `cypher_examples.go` line 17 to search for a different actor:

```go
name := "Keanu Reeves"  // Instead of "Tom Hanks"
```

**Add Your Own Movie:**
Modify `setup_data.go` to add your favorite movie:

```go
MERGE (m:Movie {title: "Your Favorite Movie", released: 2024})
```

## Understanding the Code

### Basic Query Structure

```go
result, err := neo4j.ExecuteQuery(ctx, driver,
    "MATCH (p:Person) RETURN p.name",  // The Cypher query
    nil,                                // Parameters (none here)
    neo4j.EagerResultTransformer,       // Load all results
)
```

**The 4 parts:**

1. **Context** - For timeouts and cancellation
2. **Query** - The Cypher statement
3. **Parameters** - Safe way to pass values
4. **Transformer** - How to get results back

### Using Parameters (Important!)

**Safe way (always do this):**

```go
params := map[string]any{"name": "Tom Hanks"}
cypher := "MATCH (p:Person {name: $name}) RETURN p"
```

**Unsafe way (never do this):**

```go
cypher := "MATCH (p:Person {name: '" + name + "'}) RETURN p"
```

**Why?** Hackers can break the unsafe version with special characters.

### Reading Results

```go
for _, record := range result.Records {
    name, _ := record.Get("name")
    fmt.Println(name)
}
```

Each "record" is one row of results (like a row in a table).

### Working with Nodes

```go
node := record.Get("movie").(neo4j.Node)

// Access properties
title := node.Props["title"]           // "Toy Story"
labels := node.Labels                    // ["Movie"]
id := node.ElementId                     // Unique ID
```

### Working with Relationships

```go
rel := record.Get("actedIn").(neo4j.Relationship)

// Access relationship data
relType := rel.Type                      // "ACTED_IN"
role := rel.Props["role"]                // "Woody"
startId := rel.StartElementId            // Person node ID
endId := rel.EndElementId                // Movie node ID
```

### Working with Paths

```go
path := record.Get("path").(neo4j.Path)

// Iterate through the path
for _, node := range path.Nodes {
    fmt.Println(node.Props["name"])
}

for _, rel := range path.Relationships {
    fmt.Println(rel.Type)
}
```

### Working with Sessions and Transactions

**What's the difference?**

- **ExecuteQuery** - Quick one-time queries (what we used before)
- **Sessions** - For production apps that need transaction control

**When to use sessions:**

- Multiple related queries that must succeed together
- Better control over when data is saved
- Production applications

**Creating a session:**

```go
session := driver.NewSession(ctx, neo4j.SessionConfig{})
defer session.Close(ctx)  // Always close when done!
```

**ExecuteRead (for reading data):**

```go
result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
    result, err := tx.Run(ctx,
        "MATCH (m:Movie) RETURN m.title LIMIT 5",
        nil)
    // Process results...
    return movies, nil
})
```

- Faster in clusters (can use any server)
- Use for SELECT-like queries

**ExecuteWrite (for changing data):**

```go
result, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
    result, err := tx.Run(ctx,
        "CREATE (p:Person {name: $name})",
        map[string]any{"name": "Alice"})
    // Process results...
    return person, nil
})
```

- Goes to the main server in clusters
- Use for CREATE, UPDATE, DELETE queries

**Why use transactions?**

1. **Atomic operations** - All queries succeed or all fail together
2. **No partial updates** - Your data stays consistent
3. **Automatic retry** - Network issues handled automatically
4. **Automatic rollback** - Errors undo all changes

**Example: Transfer money (atomic)**

```go
// Both queries must succeed or both fail
session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
    // Subtract from account 1
    tx.Run(ctx, "MATCH (a:Account {id: $id}) SET a.balance = a.balance - $amount", ...)

    // Add to account 2
    tx.Run(ctx, "MATCH (a:Account {id: $id}) SET a.balance = a.balance + $amount", ...)

    // If any error: both are undone!
    return nil, nil
})
```

**Transaction best practices:**

- ✅ Always close sessions with `defer`
- ✅ Keep transactions short (memory is limited)
- ✅ Use ExecuteRead for queries, ExecuteWrite for changes
- ✅ Return error to rollback transaction
- ✅ Let Neo4j handle retries automatically

## Data Types Reference

### Simple Types

| Neo4j Type  | Go Type     | Example     |
| ----------- | ----------- | ----------- |
| `null`    | `nil`     | `nil`     |
| `Boolean` | `bool`    | `true`    |
| `Integer` | `int64`   | `42`      |
| `Float`   | `float64` | `3.14`    |
| `String`  | `string`  | `"Hello"` |

### Collection Types

| Neo4j Type | Go Type                    | Example            |
| ---------- | -------------------------- | ------------------ |
| `List`   | `[]interface{}`          | `[1, 2, 3]`      |
| `Map`    | `map[string]interface{}` | `{key: "value"}` |

### Graph Types

| Neo4j Type       | Go Type                | What It Represents            |
| ---------------- | ---------------------- | ----------------------------- |
| `Node`         | `neo4j.Node`         | A thing (person, movie)       |
| `Relationship` | `neo4j.Relationship` | A connection between things   |
| `Path`         | `neo4j.Path`         | A sequence of connected nodes |

### Temporal Types

| Neo4j Type        | Go Type                 | Example                  |
| ----------------- | ----------------------- | ------------------------ |
| `Date`          | `neo4j.Date`          | `2025-01-01`           |
| `DateTime`      | `time.Time`           | `2025-01-01T12:00:00Z` |
| `LocalDateTime` | `neo4j.LocalDateTime` | `2025-01-01T12:00:00`  |
| `Time`          | `neo4j.OffsetTime`    | `12:00:00Z`            |
| `LocalTime`     | `neo4j.LocalTime`     | `12:00:00`             |
| `Duration`      | `neo4j.Duration`      | `P1M10DT12H30M`        |

**Working with dates:**

```go
// Create a date
date := neo4j.Date{Year: 2025, Month: 1, Day: 15}

// Convert Neo4j Date to Go time.Time
goTime := date.Time()

// Use in query
neo4j.ExecuteQuery(ctx, driver,
    "CREATE (e:Event {date: $date})",
    map[string]any{"date": time.Now()})
```

### Spatial Types

| Neo4j Type                | Go Type           | Example                            | SRID     |
| ------------------------- | ----------------- | ---------------------------------- | -------- |
| `Point` (Cartesian 2D)  | `neo4j.Point2D` | `{X: 1.0, Y: 2.0}`               | `7203` |
| `Point` (Cartesian 3D)  | `neo4j.Point3D` | `{X: 1.0, Y: 2.0, Z: 3.0}`       | `9157` |
| `Point` (Geographic 2D) | `neo4j.Point2D` | `{X: -122.41, Y: 37.77}`         | `4326` |
| `Point` (Geographic 3D) | `neo4j.Point3D` | `{X: -122.41, Y: 37.77, Z: 100}` | `4979` |

**Working with locations:**

```go
// Create geographic point (longitude, latitude)
london := neo4j.Point2D{
    X: -0.118092,      // longitude
    Y: 51.509865,      // latitude
    SpatialRefId: 4326, // WGS-84 (Earth coordinates)
}

// Calculate distance in Cypher
neo4j.ExecuteQuery(ctx, driver,
    "RETURN point.distance($p1, $p2) AS distance",
    map[string]any{"p1": london, "p2": paris})
```

## Common Problems & Solutions

### Problem: "Connection refused"

```
Error: ConnectivityError: dial tcp [::1]:7687: connect: connection refused
```

**Solution:** Neo4j isn't running

```bash
docker ps | grep neo4j        # Check if running
docker start neo4j            # Start if stopped
```

### Problem: "Authentication failed"

```
Error: authentication failed
```

**Solution:** Password doesn't match

- Check password in `main.go` (line 14)
- Check password in Docker command
- They must be the same!

### Problem: "Found 0 records"

```
Found 0 records
```

**Solution:** No data in database

```bash
go run . setup    # Add sample data
```

### Problem: "Type assertion failed"

```
panic: interface conversion: interface {} is nil
```

**Solution:** Check if value exists before using it

```go
if value, ok := node.Props["title"]; ok {
    fmt.Println(value)  // Safe
}
```

## What's Next?

### Beginner Projects

1. **Movie Search App** - Search movies by title or actor
2. **Friend Finder** - Build a simple social network
3. **Recommendation Engine** - Suggest movies based on actors

### Intermediate Projects

1. **Path Finder** - Find shortest path between two actors
2. **Graph Visualizer** - Draw the graph in a web page
3. **Data Importer** - Load data from CSV files (use transactions!)
4. **Banking App** - Transfer funds atomically between accounts
5. **Event Scheduler** - Store events with dates, times, locations

### Advanced Topics

1. **Indexes** - Make queries faster with CREATE INDEX
2. **Constraints** - Ensure data is valid with CONSTRAINT
3. **Full-Text Search** - Search text in properties
4. **Clustering** - Deploy Neo4j in high-availability mode
5. **Performance Tuning** - Optimize queries with EXPLAIN and PROFILE

## Learn More

### Official Resources

- [Neo4j Browser Guide](https://neo4j.com/docs/browser-manual/) - How to use the web interface
- [Cypher Basics](https://neo4j.com/docs/getting-started/cypher-intro/) - Query language tutorial
- [Go Driver Docs](https://neo4j.com/docs/go-manual) - Complete reference
- [GraphAcademy](https://graphacademy.neo4j.com/) - Free online courses

### Practice

- [Cypher Tutorial](https://neo4j.com/graphacademy/training-cypher-40/) - Interactive lessons
- [Graph Examples](https://neo4j.com/developer/example-data/) - More datasets to try

## Managing Neo4j

### Start/Stop Database

```bash
# Start
docker start neo4j

# Stop
docker stop neo4j

# Check status
docker ps | grep neo4j
```

### View Logs

```bash
docker logs neo4j
```

### Access Shell

```bash
docker exec -it neo4j bash
```

### Backup Data

```bash
docker exec neo4j neo4j-admin dump --database=neo4j --to=/backups/neo4j.dump
```

### Remove Everything

```bash
docker stop neo4j
docker rm neo4j
```

To start fresh, just run the Docker command from Step 1 again!

## Tips & Best Practices


Key concepts covered:

* **Transaction Functions** -** **`ExecuteRead()` and** **`ExecuteWrite()` for managing transactions
* **Unit of Work Pattern** - Grouping related operations into a single transaction
* **Automatic Rollback** - Any error causes the entire transaction to be rolled back
* **ACID Compliance** - All operations succeed together or fail together
* **Result Consumption** - Processing results within the transaction function

You should use transaction functions for read and write operations when you want to start consuming results as soon as they are available, and when you need to ensure data consistency across multiple operations.

In the next lesson, you will take a quiz to test your knowledge of using transactions.

### Writing Queries

- ✅ Use parameters for user input
- ✅ Start with small LIMIT to test
- ✅ Use EXPLAIN to see query plan
- ✅ Create indexes for common searches

### Go Code

- ✅ Reuse the driver (don't create new ones)
- ✅ Always close the driver with `defer`
- ✅ Handle errors properly
- ✅ Type assert graph types before using

### Transactions

- ✅ Use ExecuteQuery for simple one-off queries
- ✅ Use sessions for production apps
- ✅ ExecuteRead for queries, ExecuteWrite for changes
- ✅ Always close sessions with `defer`
- ✅ Keep transactions short
- ✅ Return error to rollback automatically

### Performance

- ✅ Use indexes on commonly searched properties
- ✅ Limit results when testing
- ✅ Use parameters (they're cached)
- ✅ Close transactions quickly

## Need Help?

- Check the [Neo4j Community Forum](https://community.neo4j.com/)
- Ask questions on [Stack Overflow](https://stackoverflow.com/questions/tagged/neo4j)
- Read the [Go Driver Issues](https://github.com/neo4j/neo4j-go-driver/issues)

## Summary

You now know:

- ✅ How to start Neo4j
- ✅ How to connect from Go
- ✅ How to run queries safely with parameters
- ✅ How to work with nodes, relationships, and paths
- ✅ How to handle dates, times, durations (temporal types)
- ✅ How to work with locations and distances (spatial types)
- ✅ How to manage transactions for production apps
- ✅ When to use ExecuteQuery vs Sessions
- ✅ How atomic transactions prevent data corruption
- ✅ Best practices for robust Neo4j applications

**You've completed 34 working examples!**

- 7 query execution examples
- 7 result handling examples
- 10 temporal & spatial examples
- 10 transaction management examples

**Next:** Build something cool! 🚀

## License

This is educational code based on Neo4j GraphAcademy examples.
