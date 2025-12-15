Here's what was created:

##   Implementation Summary

  main.go:1-19 - The complete implementation includes:

1. Driver Instance Creation - Creates a Neo4j driver with connection to neo4j://localhost:7687
2. Authentication - Uses basic auth with username "neo4j" (you'll need to update the password)
3. Connectivity Verification - Calls VerifyConnectivity() to ensure the connection works
4. First Query - Executes a query to return a greeting message and count all nodes
5. Database-Specific Query - Example of querying a specific database ("neo4j")
6. Proper Cleanup - Uses defer to close the driver when the function exits

  Next Steps

  Before running the application:

1. Update the password in main.go:13 - Replace "your-password" with your actual Neo4j password
2. Ensure Neo4j is running on localhost:7687

  Running the Application

 **cd /Users/rupeshpanwar/Documents/PProject/AI-Infra/hello-neo4j
  go run main.go**

  The application will:

- Verify the connection
- Print "Hello, Neo4j!"
- Display the total number of nodes in your database
- Show the count of nodes in the 'neo4j' database

  All the concepts from the guide have been implemented including driver lifecycle management, connectivity verification, and query execution.


# Installing the driver

In the** **[Cypher Fundamentals](https://graphacademy.neo4j.com/courses/cypher-fundamentals/) course, you learned how to query Neo4j using Cypher.

To run Cypher statements in a Go application, you’ll need the** **[Neo4j Go Driver](https://neo4j.com/docs/go-manual). The driver acts as a bridge between your Go code and Neo4j, handling connections to the database and the execution of Cypher queries.

## Setting up a Go Project

First, create a new Go project and initialize it as a module:

shell

Copy

```shell
mkdir hello-neo4j
cd hello-neo4j
go mod init graphacademy/hello
```

## Installing the Driver

To install the driver, use the** **`go get` command:

shell

Copy

```shell
go get github.com/neo4j/neo4j-go-driver/v5
```

## Creating a Driver Instance

go

Copy

```go
package main

import (
    "context"
    "fmt"
    "github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func main() {
    driver, err := neo4j.NewDriverWithContext(
        "neo4j://localhost:7687", // (1)
        neo4j.BasicAuth("neo4j", "your-password", ""), // (2)
    )
    if err != nil { panic(err) }
    defer driver.Close(context.Background()) // (3)
}
```

You start by importing the driver and creating an instance using the** **`neo4j.NewDriverWithContext()` function:

1. The connection string for your Neo4j database
2. Your Neo4j username and password
3. Always close the driver when done using the** **`defer` statement

Best Practice

Create** ****one** Driver instance and share it across your entire application.

## Verifying Connectivity

You can verify the connection is correct by calling the** **`VerifyConnectivity()`method.

go

Copy

```go
ctx := context.Background()
err = driver.VerifyConnectivity(ctx)
if err != nil {
    panic(err)
}
```

Verify Connectivity

The** **`VerifyConnectivity()` method will** **[return an error](https://neo4j.com/docs/go-manual/current/connect/) if the connection cannot be made.

## Running Your First Query

The** **`ExecuteQuery()` method executes a Cypher query and returns the results.

go

Copy

```go
ctx := context.Background()
result, err := neo4j.ExecuteQuery(ctx, driver, // (1)
    "RETURN COUNT {()} AS count",
    nil, // (2)
    neo4j.EagerResultTransformer, // (3)
)
if err != nil { panic(err) }

// Get the first record
first := result.Records[0] // (4)
// Print the count entry
count, _ := first.Get("count") // (5)
fmt.Println(count)
```

### What is happening here?

1. `neo4j.ExecuteQuery()` runs a Cypher query to get the count of all nodes in the database
2. `nil` means no parameters are passed to the query
3. `EagerResultTransformer` loads all results into memory
4. `Records` contains a slice of the rows returned
5. Values from the** **`RETURN` clause are accessed using the** **`Get()` method

## Specifying a database

You can specify a database to query using the** **`ExecuteQueryWithDatabase()`option.

go

Copy

```go
ctx := context.Background()
result, err := neo4j.ExecuteQuery(ctx, driver,
    "RETURN COUNT {()} AS count",
    nil,
    neo4j.EagerResultTransformer,
    neo4j.ExecuteQueryWithDatabase("neo4j"), // Query the 'neo4j' database
)
if err != nil { panic(err) }
```

## Full driver lifecycle

It is good practice to close the driver when you are finished with it. You can use the** **`defer` statement to automatically close the driver when the function exits.

You can also explicitly call** **`Close()` to release any resources held by the driver.

go

Copy

```go
driver, err := neo4j.NewDriverWithContext(
    "neo4j://localhost:7687",
    neo4j.BasicAuth("neo4j", "your-password", ""),
)
if err != nil {
    panic(err)
}
defer driver.Close(context.Background())
```

go

Copy

```go
ctx := context.Background()
driver.Close(ctx)
```


Check Neo4j Status

# If using Neo4j Desktop, ensure the database is started

# If using Docker:

  docker ps | grep neo4j

# Check if Neo4j is listening on port 7687:

  lsof -i :7687

  Start Neo4j

  Option 1: Using Docker
  docker run -d 
    --name neo4j
    -p 7474:7474 -p 7687:7687
    -e NEO4J_AUTH=neo4j/Your@Password!@#
    neo4j:latest

  Option 2: Using Neo4j Desktop

- Open Neo4j Desktop
- Start your database instance

  Option 3: Using Neo4j installed locally
  neo4j start

  Verify Neo4j is Running

  Once started, you should be able to access:

- Browser interface: http://localhost:7474
- Bolt protocol: bolt://localhost:7687

  Then run the application again:
  go run main.go


Bolt Protocol is Neo4j's proprietary binary protocol for client-server communication. Think of it as the "language" that Neo4j clients (like our Go application) use to talk to the Neo4j database server.

  Key Points

1. Binary Protocol

- Efficient, binary format (not text-based like HTTP)
- Optimized for database operations
- Lower overhead than REST APIs

2. Purpose

- Executes Cypher queries
- Manages transactions
- Streams results back to the client
- Handles authentication

3. Port 7687

- Default port for Bolt connections
- What the Neo4j driver connects to
- Similar to how MySQL uses port 3306, PostgreSQL uses 5432

  In Your Code

  driver, err := neo4j.NewDriverWithContext(
      "neo4j://localhost:7687",  // This uses Bolt protocol
      neo4j.BasicAuth("neo4j", "Your@Password!@#", ""),
  )

  The URI scheme neo4j:// tells the driver to use the Bolt protocol over port 7687.

  Comparison

| Protocol | Port | Purpose                                            |
| -------- | ---- | -------------------------------------------------- |
| Bolt     | 7687 | Fast binary protocol for queries (used by drivers) |
| HTTP     | 7474 | Browser interface, REST API, visualization         |

  The HTTP interface (port 7474) is for humans to browse and explore the database visually, while Bolt (port 7687) is for applications to interact with Neo4j programmatically - which is exactly what your Go
  application is doing!


1. Started Neo4j Database - Launched Neo4j using Docker with:
   - Port 7474 (HTTP interface)
   - Port 7687 (Bolt protocol)
   - Authentication matching your code: neo4j/Your@Password!@#
2. Verified Connection - The application now successfully:
   - Connects to Neo4j
   - Verifies connectivity
   - Executes queries
   - Returns results

  Output Confirmed

  Connection verified successfully!
  Hello, Neo4j!
  Total nodes in database: 0
  Nodes in 'neo4j' database: 0

  Access Points

- Browser UI: http://localhost:7474
- Bolt Protocol: bolt://localhost:7687

  Your Neo4j instance is now running and ready for development. The database is empty (0 nodes), which is expected for a fresh installation.



Bolt Protocol is Neo4j's proprietary binary protocol for client-server communication. Think of it as the "language" that Neo4j clients (like our Go application) use to talk to the Neo4j database server.

  Key Points

1. Binary Protocol

- Efficient, binary format (not text-based like HTTP)
- Optimized for database operations
- Lower overhead than REST APIs

2. Purpose

- Executes Cypher queries
- Manages transactions
- Streams results back to the client
- Handles authentication

3. Port 7687

- Default port for Bolt connections
- What the Neo4j driver connects to
- Similar to how MySQL uses port 3306, PostgreSQL uses 5432

  In Your Code

  driver, err := neo4j.NewDriverWithContext(
      "neo4j://localhost:7687",  // This uses Bolt protocol
      neo4j.BasicAuth("neo4j", "Your@Password!@#", ""),
  )

  The URI scheme neo4j:// tells the driver to use the Bolt protocol over port 7687.

  Comparison

| Protocol | Port | Purpose                                            |
| -------- | ---- | -------------------------------------------------- |
| Bolt     | 7687 | Fast binary protocol for queries (used by drivers) |
| HTTP     | 7474 | Browser interface, REST API, visualization         |

  The HTTP interface (port 7474) is for humans to browse and explore the database visually, while Bolt (port 7687) is for applications to interact with Neo4j programmatically - which is exactly what your Go
  application is doing!



This is a real, full-featured Neo4j database - just like SQLite or PostgreSQL, but with a key difference: it's a graph database instead of a relational database.

  What You're Running

  docker run -d --name neo4j neo4j:latest

  This is running:

- Neo4j Community Edition (the free, open-source version)
- Version 2025.10.1 (as shown in the logs)
- A complete database server with persistent storage

  Database Comparison

| Feature        | SQLite                  | PostgreSQL            | Neo4j                            |
| -------------- | ----------------------- | --------------------- | -------------------------------- |
| Type           | Relational (file-based) | Relational (server)   | Graph (server)                   |
| Data Model     | Tables, rows, columns   | Tables, rows, columns | Nodes, relationships, properties |
| Query Language | SQL                     | SQL                   | Cypher                           |
| Best For       | Single-user apps        | Traditional apps      | Connected data, relationships    |
| Storage        | Single file             | Database cluster      | Graph storage engine             |

  What Makes Neo4j Different

  Traditional Database (PostgreSQL/SQLite):
  SELECT u.name, f.name
  FROM users u
  JOIN friendships ON u.id = friendships.user_id
  JOIN users f ON f.id = friendships.friend_id
  WHERE u.name = 'Alice'

  Graph Database (Neo4j):
  MATCH (alice:User {name: 'Alice'})-[:FRIENDS_WITH]->(friend)
  RETURN friend.name

  Your Current Setup

  The Docker container includes:

- Database engine - Stores and queries graph data
- Data persistence - Data stored in /data inside the container
- Transaction support - ACID compliant
- Multi-database support - Can create multiple databases

  Is This Production-Ready?

  Current Setup:

- ✅ Real database
- ⚠️ Data will be lost if you remove the container (no volume mounted)
- ⚠️ Default configuration (not optimized)

  For Development: Perfect as-is

  For Production/Persistence: Mount a volume:
  docker run -d 
    --name neo4j
    -p 7474:7474 -p 7687:7687
    -v neo4j-data:/data
    -e NEO4J_AUTH=neo4j/Your@Password!@#
    neo4j:latest

  Check Your Database

  You can verify it's a real database by:

1. Browser UI: http://localhost:7474
2. Check container: docker exec -it neo4j neo4j-admin --version
3. Data directory: docker exec -it neo4j ls -la /data

  So yes - this is as real as PostgreSQL or MySQL, just optimized for graph data instead of tables
