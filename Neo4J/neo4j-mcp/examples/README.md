# Neo4j MCP Examples for Claude CLI

Practical examples demonstrating how to use the Neo4j Cypher MCP Server with Claude CLI.

## What's Included

### Shell Scripts

1. **`01_basic_queries.sh`** - Basic read-only queries
   - Get database schema
   - Count nodes
   - List movies
   - Find actors in movies
   - Find movies by actor

2. **`02_advanced_queries.sh`** - Advanced graph queries
   - Find co-stars
   - Shortest path between actors
   - Multi-talented people (actor-directors)
   - Movies by decade
   - Most connected actors
   - Database analytics

3. **`03_write_operations.sh`** - Data modification operations
   - Create movies
   - Create actors
   - Create relationships
   - Update properties
   - Verify changes

### Python Application

4. **`interactive_demo.py`** - Interactive menu-driven demo
   - 10 example queries
   - Step-by-step execution
   - Read and write operations
   - User confirmation for writes

---

## Prerequisites

Make sure you have:

1. **Neo4j MCP server installed and connected:**
   ```bash
   claude mcp list
   # Should show: neo4j-cypher: ... - ✓ Connected
   ```

2. **Neo4j database running:**
   ```bash
   docker ps | grep neo4j
   # Should show neo4j container running
   ```

3. **Claude CLI working:**
   ```bash
   claude --version
   ```

4. **Python 3.x** (for interactive demo):
   ```bash
   python3 --version
   ```

---

## Quick Start

### Run Basic Queries

```bash
cd /Users/rupeshpanwar/Documents/PProject/AI-Infra/neo4j-mcp/examples

# Make scripts executable
chmod +x *.sh

# Run basic queries
./01_basic_queries.sh
```

**Output:**
```
🎬 Neo4j MCP - Basic Queries Demo
==================================

1️⃣  Getting database schema...
---
[Claude shows database structure]
✅ Done

2️⃣  Counting movies and actors...
---
[Claude shows counts]
✅ Done
...
```

---

### Run Advanced Queries

```bash
./02_advanced_queries.sh
```

**This will:**
- Find actor connections
- Calculate shortest paths
- Analyze database statistics
- Show multi-talented people

---

### Run Write Operations

⚠️ **WARNING:** This modifies your database!

```bash
./03_write_operations.sh
```

**Follow the prompts:**
```
⚠️  WARNING: These operations will modify your database!
Claude will ask for approval before each operation.

Continue? (yes/no): yes

1️⃣  Creating a new movie: Inception...
---
[Claude shows CREATE query and asks for approval]
```

---

### Run Interactive Demo

```bash
python3 interactive_demo.py
```

**Interactive menu:**
```
╔════════════════════════════════════════════════════════════════════╗
║         Neo4j MCP Interactive Demo for Claude CLI                 ║
╚════════════════════════════════════════════════════════════════════╝

Select an example to run:

   1. Get Database Schema
   2. Count Movies and Actors
   3. List All Movies
   4. Find Actors in The Matrix
   5. Find Tom Hanks Movies
   6. Find Co-Stars
   7. Shortest Path Between Actors
   8. Database Analytics
   9. Add New Movie (Write)
  10. Add Actor and Relationship (Write)

   0. Exit
```

---

## Example Outputs

### Example 1: Get Schema

```bash
./01_basic_queries.sh
```

**Claude's response:**
```
Your Neo4j database has the following structure:

Node Labels:
• Movie
  - Properties: title (STRING), released (INTEGER), tagline (STRING),
                imdbRating (FLOAT), revenue (INTEGER)

• Person
  - Properties: name (STRING), born (DATE), tmdbId (STRING)

• Genre
  - Properties: name (STRING)

Relationships:
• (Person)-[:ACTED_IN]->(Movie)
• (Person)-[:DIRECTED]->(Movie)
• (Movie)-[:IN_GENRE]->(Genre)
```

---

### Example 2: Find Actors

**Query:**
```bash
claude "Who acted in The Matrix?"
```

**Response:**
```
4 actors appeared in The Matrix:

1. Carrie-Anne Moss
2. Hugo Weaving
3. Keanu Reeves
4. Laurence Fishburne
```

---

### Example 3: Shortest Path

**Query:**
```bash
claude "What's the shortest path between Tom Hanks and Keanu Reeves?"
```

**Response:**
```
I found a connection between Tom Hanks and Keanu Reeves:

Path: Tom Hanks → Cloud Atlas ← Hugo Weaving → The Matrix ← Keanu Reeves

Explanation:
1. Tom Hanks acted in Cloud Atlas
2. Hugo Weaving also acted in Cloud Atlas
3. Hugo Weaving acted in The Matrix
4. Keanu Reeves also acted in The Matrix

They are connected through Hugo Weaving with 2 degrees of separation.
```

---

## Understanding MCP Tool Usage

Each Claude query uses one of the three MCP tools:

### 🔍 `get_neo4j_schema`

**Triggered by:**
- "What's the structure..."
- "Show me the schema..."
- "What node types..."

**Example:**
```bash
claude "What's in my Neo4j database?"
```

---

### 📖 `read_neo4j_cypher`

**Triggered by:**
- "Show me..."
- "Find..."
- "List..."
- "Count..."
- "Who..."
- "What..."

**Example:**
```bash
claude "Show me all movies from 1999"
```

**Claude generates and runs:**
```cypher
MATCH (m:Movie)
WHERE m.year = 1999
RETURN m.title, m.released
ORDER BY m.title
```

---

### ✍️ `write_neo4j_cypher`

**Triggered by:**
- "Add..."
- "Create..."
- "Update..."
- "Delete..."
- "Set..."

**Example:**
```bash
claude "Add a new movie called Inception from 2010"
```

**Claude shows:**
```cypher
CREATE (m:Movie {
  title: 'Inception',
  released: 2010
})
RETURN m
```

**Then asks:** `Approve? (yes/no)`

---

## Customizing the Scripts

### Modify Queries

Edit the shell scripts to ask different questions:

```bash
# In 01_basic_queries.sh, change:
claude "List all movies..."

# To:
claude "Show me movies with IMDb rating above 8.0"
```

### Add New Examples

```bash
# Add to any script:
echo "My custom query..."
claude "Your custom question about Neo4j database"
```

---

## Common Use Cases

### Use Case 1: Database Exploration

```bash
# Quick exploration
claude "What's in my database?"
claude "How many nodes total?"
claude "Show me 5 random movies"
```

---

### Use Case 2: Data Analysis

```bash
# Analytics queries
claude "Count movies by decade"
claude "Which actor has the most movies?"
claude "What's the average movie year?"
```

---

### Use Case 3: Building Applications

```python
# In your Python app
import subprocess

def query_neo4j(question):
    result = subprocess.run(
        ['claude', f"From my Neo4j database, {question}"],
        capture_output=True,
        text=True
    )
    return result.stdout

# Use it
movies = query_neo4j("list all movies from 1999")
actors = query_neo4j("who acted in The Matrix?")
```

---

### Use Case 4: Data Migration

```bash
# Add new data
claude "Add these movies: Inception (2010), Interstellar (2014)"
claude "Add Leonardo DiCaprio as an actor"
claude "Connect Leonardo DiCaprio to Inception"

# Verify
claude "Show me all of Leonardo DiCaprio's movies"
```

---

## Best Practices

### 1. Always Verify Before Writing

Before write operations:
```bash
# Check if data exists
claude "Is there a movie called Inception?"

# If not, add it
claude "Add movie Inception from 2010"
```

---

### 2. Use Specific Language

❌ **Vague:**
```bash
claude "Show me stuff"
```

✅ **Specific:**
```bash
claude "Show me movies from the 1990s with their titles and years"
```

---

### 3. Review Write Queries

Always review the Cypher query Claude generates:
- Check property names
- Verify values
- Understand relationships

---

### 4. Test Queries First

Test complex queries with LIMIT:
```bash
# Test first
claude "Show me the first 5 results of..."

# Then run full query
claude "Show me all results of..."
```

---

## Troubleshooting

### Script Won't Run

```bash
# Make executable
chmod +x 01_basic_queries.sh

# Run with bash
bash 01_basic_queries.sh
```

---

### Claude Not Found

```bash
# Check installation
which claude

# If not found, install Claude CLI
# Follow installation instructions at claude.ai/code
```

---

### MCP Server Not Connected

```bash
# Check status
claude mcp list

# Restart if needed
claude mcp restart neo4j-cypher
```

---

### Neo4j Not Running

```bash
# Start Neo4j
docker start neo4j

# Verify
docker ps | grep neo4j
```

---

## Next Steps

1. **Run all three shell scripts** to see different query types
2. **Try the interactive demo** for hands-on learning
3. **Modify the scripts** with your own questions
4. **Build your own application** using the patterns shown
5. **Read the full guide** in `CLAUDE_CLI_GUIDE.md`

---

## Additional Resources

- **Main Guide:** `../CLAUDE_CLI_GUIDE.md`
- **Tutorial:** `../TUTORIAL.md`
- **Easy Setup:** `../EASY-SETUP-GUIDE.md`
- **Python Examples:** `../neo4j_mcp_demo.py`
- **Practical Examples:** `../practical_examples.py`

---

## Summary

These examples demonstrate:

✅ Using Claude CLI with Neo4j MCP
✅ Three types of MCP tools
✅ Read-only queries (safe)
✅ Write operations (approved)
✅ Graph algorithms
✅ Database analytics
✅ Natural language querying

**You can now query Neo4j using plain English!** 🚀
