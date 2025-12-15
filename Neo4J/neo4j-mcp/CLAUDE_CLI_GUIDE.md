# Neo4j MCP Tools with Claude CLI

Complete guide to using the Neo4j Cypher MCP Server with Claude CLI.

## What You Have Installed

You've already successfully installed the Neo4j MCP server:

```bash
Server: neo4j-cypher
Command: uvx mcp-neo4j-cypher@0.2.3
Status: ✓ Connected
```

This gives you access to **3 powerful tools**:

1. `mcp__neo4j-cypher__get_neo4j_schema` - Discover database structure
2. `mcp__neo4j-cypher__read_neo4j_cypher` - Read data (safe, auto-approved)
3. `mcp__neo4j-cypher__write_neo4j_cypher` - Modify data (requires approval)

---

## How to Use with Claude CLI

### Method 1: Interactive Session (Recommended)

**Start Claude:**
```bash
cd /Users/rupeshpanwar/Documents/PProject/AI-Infra
claude
```

**Ask questions in natural language:**

```
You: What's the structure of my Neo4j database?
```

Claude will automatically use `get_neo4j_schema` and show you:
- Node labels (Movie, Person, Genre, User)
- Properties for each node type
- Relationships between nodes

```
You: Who acted in The Matrix?
```

Claude will use `read_neo4j_cypher` with this query:
```cypher
MATCH (p:Person)-[:ACTED_IN]->(m:Movie {title: 'The Matrix'})
RETURN p.name AS actor
ORDER BY p.name
```

```
You: Add a new movie called 'Inception' released in 2010
```

Claude will:
1. Show you the CREATE query
2. Ask for approval
3. Use `write_neo4j_cypher` to execute it

---

### Method 2: One-Shot Commands

Execute single queries without entering interactive mode:

```bash
# Quick queries
claude "How many actors are in my Neo4j database?"

claude "Show me all movies from 1999"

claude "Find the shortest path between Tom Hanks and Keanu Reeves"
```

**Perfect for:**
- Scripts and automation
- Quick lookups
- CI/CD pipelines

---

### Method 3: Programmatic Access (Advanced)

Use Claude CLI in scripts:

```bash
#!/bin/bash
# query_neo4j.sh

QUERY="$1"

claude "Using Neo4j MCP, $QUERY" --output json
```

**Usage:**
```bash
./query_neo4j.sh "count all movies"
./query_neo4j.sh "list all actors"
```

---

## The Three Tools Explained

### Tool 1: get_neo4j_schema

**Purpose:** Understand your database structure

**When Claude Uses It:**
- "What's in my database?"
- "Show me the schema"
- "What node types exist?"
- "How are movies and actors related?"

**What It Returns:**
```json
[
  {
    "label": "Person",
    "attributes": {
      "name": "STRING indexed",
      "born": "DATE",
      "tmdbId": "STRING unique indexed"
    },
    "relationships": {
      "ACTED_IN": "Movie",
      "DIRECTED": "Movie"
    }
  }
]
```

**Example Conversation:**
```
You: Describe my Neo4j data model

Claude: I'll check the schema...
[Uses get_neo4j_schema]

Your database has:
• Movie nodes (properties: title, released, tagline, imdbRating, revenue)
• Person nodes (properties: name, born, tmdbId)
• Genre nodes (properties: name)

Relationships:
• Person -[:ACTED_IN]-> Movie
• Person -[:DIRECTED]-> Movie
• Movie -[:IN_GENRE]-> Genre
```

---

### Tool 2: read_neo4j_cypher

**Purpose:** Query data safely (read-only)

**When Claude Uses It:**
- "Show me..."
- "Find..."
- "List..."
- "Count..."
- "Which..."
- "Who..."

**Safety:** ✅ Auto-approved (cannot modify data)

**Example Queries:**

**Basic Count:**
```
You: How many movies are there?

Claude uses:
MATCH (m:Movie)
RETURN count(m) as movieCount

Result: 10 movies
```

**Filtered Search:**
```
You: Show me movies from 1999

Claude uses:
MATCH (m:Movie)
WHERE m.year = 1999
RETURN m.title, m.released
ORDER BY m.title

Results:
- The Matrix (1999)
- The Green Mile (1999)
- Snow Falling on Cedars (1999)
```

**Relationship Query:**
```
You: What movies did Tom Hanks act in?

Claude uses:
MATCH (p:Person {name: 'Tom Hanks'})-[:ACTED_IN]->(m:Movie)
RETURN m.title, m.released
ORDER BY m.released DESC

Results:
- Cloud Atlas (2012)
- The Da Vinci Code (2006)
```

**Complex Analysis:**
```
You: Find actors who worked with both Tom Hanks and Keanu Reeves

Claude uses:
MATCH (tom:Person {name: 'Tom Hanks'})-[:ACTED_IN]->(m1:Movie)<-[:ACTED_IN]-(other:Person)
MATCH (keanu:Person {name: 'Keanu Reeves'})-[:ACTED_IN]->(m2:Movie)<-[:ACTED_IN]-(other)
WHERE tom <> other AND keanu <> other
RETURN DISTINCT other.name

Result: Hugo Weaving (Cloud Atlas connects them)
```

---

### Tool 3: write_neo4j_cypher

**Purpose:** Modify database (create, update, delete)

**When Claude Uses It:**
- "Add..."
- "Create..."
- "Update..."
- "Delete..."
- "Set..."

**Safety:** ⚠️ Requires approval (shows query first)

**Example Operations:**

**Create Movie:**
```
You: Add a new movie 'Inception' released in 2010

Claude shows:
CREATE (m:Movie {
  title: 'Inception',
  released: 2010,
  tagline: 'Your mind is the scene of the crime'
})
RETURN m

⚠️  This will modify your database. Approve? (yes/no)

You: yes

Claude: ✅ Movie created successfully!
```

**Create Person:**
```
You: Add Leonardo DiCaprio born in 1974

Claude shows:
CREATE (p:Person {
  name: 'Leonardo DiCaprio',
  born: date('1974-11-11')
})
RETURN p

[Asks for approval]
```

**Create Relationship:**
```
You: Leonardo DiCaprio acted in Inception

Claude shows:
MATCH (p:Person {name: 'Leonardo DiCaprio'})
MATCH (m:Movie {title: 'Inception'})
CREATE (p)-[:ACTED_IN]->(m)
RETURN p.name, m.title

[Asks for approval]
```

**Update Property:**
```
You: Update The Matrix tagline to "Free your mind"

Claude shows:
MATCH (m:Movie {title: 'The Matrix'})
SET m.tagline = 'Free your mind'
RETURN m

[Asks for approval]
```

**Delete Data:**
```
You: Remove the Inception movie

Claude shows:
MATCH (m:Movie {title: 'Inception'})
DETACH DELETE m

⚠️  WARNING: This will delete the movie and all its relationships!
Approve? (yes/no)
```

---

## Real-World Examples

### Example 1: Movie Database Analytics

```bash
claude
```

```
You: Give me analytics about my movie database

Claude will:
1. Use get_neo4j_schema to understand structure
2. Use read_neo4j_cypher to run analytics:

MATCH (m:Movie)
OPTIONAL MATCH (p:Person)-[:ACTED_IN]->(m)
OPTIONAL MATCH (d:Person)-[:DIRECTED]->(m)
RETURN
  count(DISTINCT m) AS totalMovies,
  count(DISTINCT p) AS totalActors,
  count(DISTINCT d) AS totalDirectors,
  avg(m.year) AS averageYear,
  min(m.year) AS oldestMovie,
  max(m.year) AS newestMovie

Response:
Your database contains:
• 10 movies
• 7 actors
• 3 directors
• Average year: 2004
• Oldest: 1997
• Newest: 2012
```

---

### Example 2: Actor Network Analysis

```
You: Find the most connected actor in my database

Claude uses:
MATCH (p:Person)-[:ACTED_IN|DIRECTED]->(m:Movie)
RETURN p.name, count(DISTINCT m) as connections
ORDER BY connections DESC
LIMIT 5

Results:
1. Hugo Weaving - 3 movies
2. Tom Hanks - 2 movies
3. Keanu Reeves - 2 movies
...
```

```
You: Show me how Tom Hanks and Keanu Reeves are connected

Claude uses:
MATCH path = shortestPath(
  (tom:Person {name: 'Tom Hanks'})-[*]-(keanu:Person {name: 'Keanu Reeves'})
)
RETURN path

Result:
Tom Hanks → Cloud Atlas ← Hugo Weaving → The Matrix ← Keanu Reeves
(Connected through Hugo Weaving!)
```

---

### Example 3: Building a Movie Collection

**Step 1: Add movies**
```
You: Add these movies:
- Inception (2010)
- Interstellar (2014)
- The Dark Knight (2008)

Claude will create three separate CREATE queries and ask for approval.
```

**Step 2: Add actors**
```
You: Add actors:
- Leonardo DiCaprio
- Christian Bale
- Matthew McConaughey

Claude creates Person nodes.
```

**Step 3: Connect them**
```
You: Connect the actors to their movies:
- Leonardo DiCaprio acted in Inception
- Christian Bale acted in The Dark Knight
- Matthew McConaughey acted in Interstellar

Claude creates ACTED_IN relationships.
```

**Step 4: Verify**
```
You: Show me all movies and their actors

Claude uses:
MATCH (p:Person)-[:ACTED_IN]->(m:Movie)
RETURN m.title AS movie, collect(p.name) AS actors
ORDER BY m.title
```

---

## Advanced Usage Patterns

### Pattern 1: Conditional Queries

```
You: Show me movies from the 2000s with high ratings

Claude uses:
MATCH (m:Movie)
WHERE m.year >= 2000 AND m.year < 2010
  AND m.imdbRating >= 8.0
RETURN m.title, m.year, m.imdbRating
ORDER BY m.imdbRating DESC
```

---

### Pattern 2: Aggregations

```
You: Count movies by decade

Claude uses:
MATCH (m:Movie)
WHERE m.year IS NOT NULL
WITH (m.year / 10) * 10 AS decade, count(m) AS count
RETURN decade, count
ORDER BY decade DESC
```

---

### Pattern 3: Graph Patterns

```
You: Find actors who both acted and directed

Claude uses:
MATCH (p:Person)-[:ACTED_IN]->(:Movie)
MATCH (p)-[:DIRECTED]->(:Movie)
RETURN DISTINCT p.name
```

---

### Pattern 4: Data Quality Checks

```
You: Find movies without release dates

Claude uses:
MATCH (m:Movie)
WHERE m.released IS NULL OR m.year IS NULL
RETURN m.title
```

---

## Best Practices

### 1. Always Check Schema First

Before querying unfamiliar data:
```
You: What's the structure of this database?
```

This prevents errors from non-existent labels or properties.

---

### 2. Use Specific Queries

❌ Bad:
```
Show me everything
```

✅ Good:
```
Show me the 10 most recent movies
```

---

### 3. Review Write Operations

Always review the query Claude generates before approving:
- Check which nodes will be affected
- Verify property values
- Understand relationship directions

---

### 4. Use Parameters for Safety

Claude automatically uses parameters:
```cypher
MATCH (m:Movie {title: $title})  // Safe
# vs
MATCH (m:Movie {title: 'user input'})  // Risky
```

---

## Troubleshooting

### Issue: "MCP server not found"

**Check server status:**
```bash
claude mcp list
```

**Should show:**
```
neo4j-cypher: uvx mcp-neo4j-cypher@0.2.3 - ✓ Connected
```

**If not connected:**
```bash
# Restart Claude
claude mcp restart neo4j-cypher
```

---

### Issue: "Database connection failed"

**Check Neo4j is running:**
```bash
docker ps | grep neo4j
```

**Test connection:**
```bash
claude "Test Neo4j connection by counting movies"
```

---

### Issue: Query timeout

For long-running queries:
```
You: Run this query with a 30 second timeout: [your query]
```

---

## Quick Reference

### Common Questions

**Database exploration:**
- "What's in my database?"
- "Show me the schema"
- "Count all nodes"
- "List all node types"

**Finding data:**
- "Show me all [nodes]"
- "Find [specific item]"
- "Count [items]"
- "List [items] where [condition]"

**Relationships:**
- "Who [relationship] [item]?"
- "What [items] did [person] [relationship]?"
- "Find connections between [A] and [B]"
- "Show me the shortest path from [A] to [B]"

**Modifying data:**
- "Add a new [node] called [name]"
- "Create [node] with [properties]"
- "Update [node] set [property] to [value]"
- "Delete [node]"
- "Connect [A] to [B] with [relationship]"

---

## Example Session

Here's a complete example session:

```bash
$ cd /Users/rupeshpanwar/Documents/PProject/AI-Infra
$ claude

You: Hello! I want to explore my Neo4j movie database

Claude: I'll help you explore your Neo4j database. Let me start by checking what's in there.

[Uses get_neo4j_schema]

Your database contains:
- Movie nodes (10 total)
- Person nodes (7 total)
- Genre nodes (various)

Would you like me to show you some specific data?

You: Yes, show me all movies from 1999

[Uses read_neo4j_cypher]

Found 3 movies from 1999:
1. The Matrix
2. The Green Mile
3. Snow Falling on Cedars

You: Who acted in The Matrix?

[Uses read_neo4j_cypher]

4 actors in The Matrix:
1. Keanu Reeves
2. Laurence Fishburne
3. Carrie-Anne Moss
4. Hugo Weaving

You: Add a new movie called Inception from 2010

Claude: I'll create that movie. Here's the query:

CREATE (m:Movie {title: 'Inception', released: 2010})
RETURN m

Approve? (yes/no)

You: yes

[Uses write_neo4j_cypher]

✅ Movie created successfully!

You: Show me all movies now

[Uses read_neo4j_cypher]

Your database now has 11 movies, including the newly added Inception.
```

---

## Next Steps

1. **Try the interactive examples** in `practical_examples.py`
2. **Read the tutorial** in `TUTORIAL.md`
3. **Explore advanced patterns** in the GraphAcademy course
4. **Build your own queries** using natural language

---

## Summary

**You can now:**
- ✅ Query Neo4j using natural language
- ✅ Explore database structure automatically
- ✅ Read data safely without writing Cypher
- ✅ Modify data with approval workflow
- ✅ Analyze complex graph patterns
- ✅ Build movie networks interactively

**No more:**
- ❌ Writing Cypher manually
- ❌ Compiling Go code
- ❌ Managing database connections
- ❌ Debugging queries

Just ask Claude in plain English! 🚀
