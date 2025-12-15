# 🎮 How to Connect Claude to Your Neo4j Database

**A Super Simple Guide Anyone Can Follow!**

---

## 🤔 What Are We Doing?

Imagine you have a graph database (Neo4j) full of movies and actors. Right now, to get information from it, you need to:

1. Write code in Go ❌
2. Compile it ❌
3. Run it ❌
4. Wait for results ❌

**That's a lot of work!**

**After this tutorial**, you can just ask Claude questions in plain English:

- "Show me all movies" ✅
- "Who acted with Tom Hanks?" ✅
- "Find the shortest path between two actors" ✅

**And Claude will answer instantly!** 🚀

---

## 🎯 What You Need First

Before we start, make sure you have these ready:

### ✅ Checklist

- [ ] **Neo4j database running** (in Docker)
- [ ] **Claude Code installed** (the `claude` command works in terminal)
- [ ] **Python's uvx installed** (don't worry, it's probably already there!)
- [ ] **Terminal or Command Line** open

### 🧪 Quick Test

Open your terminal and try these commands:

```bash
# Test 1: Is Neo4j running?
docker ps | grep neo4j
```

**You should see:** Something like `neo4j` with "Up" status

```bash
# Test 2: Is Claude Code working?
claude --version
```

**You should see:** A version number like `2.0.35`

```bash
# Test 3: Is uvx installed?
uvx --version
```

**You should see:** A version number

**If all three tests passed, you're ready! Let's go! 🎉**

---

## 📚 Step-by-Step Installation

### Step 1: Go to Your Project Folder

Open terminal and type:

```bash
cd /Users/rupeshpanwar/Documents/PProject/AI-Infra
```

**Why?** This is where we'll install the MCP server.

**What's MCP?** Think of it as a "translator" that helps Claude talk to your Neo4j database.

---

### Step 2: Install the Neo4j MCP Server

Copy and paste this **exact** command:

```bash
claude mcp add --transport stdio neo4j-cypher --env NEO4J_URI=bolt://localhost:7687 --env NEO4J_USERNAME=neo4j --env "NEO4J_PASSWORD=Your@Password!@#" --env NEO4J_DATABASE=neo4j -- uvx mcp-neo4j-cypher@0.2.3
```

**What does this do?**

- `claude mcp add` = "Hey Claude, add a new tool!"
- `neo4j-cypher` = Name of the tool
- `--env` = Settings for connecting to your database
- `uvx mcp-neo4j-cypher@0.2.3` = Download and run this special program

**You should see:**

```
Added stdio MCP server neo4j-cypher with command: uvx mcp-neo4j-cypher@0.2.3 to local config
File modified: /Users/rupeshpanwar/.claude.json
```

**✅ Success!** The MCP server is installed!

---

### Step 3: Check If It's Working

Type this command:

```bash
claude mcp list
```

**You should see:**

```
Checking MCP server health...

neo4j-cypher: uvx mcp-neo4j-cypher@0.2.3 - ✓ Connected
```

**See the ✓ Connected?** That means it's working! 🎊

---

## 🎮 How to Use It

### Starting Claude

From your project folder, type:

```bash
claude
```

This starts Claude in interactive mode.

---

### Ask Questions!

Now you can ask Claude about your database in **plain English**:

#### Example 1: See All Movies

**You type:**

```
Show me all movies in the Neo4j database
```

**Claude will answer with:**

```
Found 10 movies:
1. Cloud Atlas (2012)
2. The Da Vinci Code (2006)
3. The Matrix Reloaded (2003)
... and more!
```

---

#### Example 2: Find Co-stars

**You type:**

```
Which actors worked with Tom Hanks?
```

**Claude will answer with:**

```
2 actors worked with Tom Hanks:
1. Gary Sinise (in Forrest Gump)
2. Robin Wright (in Forrest Gump)
```

---

#### Example 3: Get Database Info

**You type:**

```
What's the schema of my Neo4j database?
```

**Claude will answer with:**

```
Your database has:
- Node types: Movie, Person
- Relationships: ACTED_IN, DIRECTED
- Properties: name, title, released, tagline
```

---

## 🧠 Understanding What Happened

### Before (The Hard Way)

```
You → Write Go code → Compile → Run → Get answer
      (5 minutes of work)
```

### After (The Easy Way)

```
You → Ask Claude in English → Get instant answer
      (5 seconds!)
```

### What's the Magic?

The **MCP server** is like a bridge:

```
You (English question)
    ↓
Claude (understands you)
    ↓
MCP Server (translates to database language)
    ↓
Neo4j Database (finds the answer)
    ↓
MCP Server (translates back)
    ↓
Claude (tells you in English)
```

---

## 🔧 What Can You Do Now?

### 3 Cool Things the MCP Server Can Do

#### 1. **Get Database Schema**

Tool: `get-neo4j-schema`

Ask Claude:

```
Show me the structure of my database
```

This tells you what kind of data you have (movies, people, relationships).

---

#### 2. **Read Data (Safe Queries)**

Tool: `read-neo4j-cypher`

Ask Claude:

```
Find all movies from 1999
```

This only *reads* data - it won't change anything!

---

#### 3. **Write Data (Modify Database)**

Tool: `write-neo4j-cypher`

Ask Claude:

```
Add a new movie called "Star Wars" released in 1977
```

This can *create, update, or delete* data - be careful!

---

## 🎯 Fun Queries to Try

### Beginner Level 🟢

```
Show me all movies
```

```
How many actors are in the database?
```

```
What's in my database?
```

---

### Intermediate Level 🟡

```
Find all actors who worked with Keanu Reeves
```

```
Show me movies from the 1990s
```

```
List all directors and their movies
```

---

### Advanced Level 🔴

```
What's the shortest path between Tom Hanks and Keanu Reeves?
```

```
Which movie has the most actors?
```

```
Find actors who have worked together more than once
```

---

## 📖 What Did We Learn?

### ✅ You Can Now:

1. **Install MCP servers** using Claude CLI
2. **Connect Claude to Neo4j** database
3. **Ask questions in plain English** instead of writing code
4. **Get instant answers** from your database
5. **Troubleshoot problems** when things go wrong

---

## 🎓 The Magic Formula

Remember this command pattern for installing **any** MCP server:

```bash
claude mcp add --transport stdio [SERVER-NAME] --env KEY1=value1 --env KEY2=value2 -- [COMMAND]
```

**Example for Neo4j:**

```bash
claude mcp add --transport stdio neo4j-cypher --env NEO4J_URI=bolt://localhost:7687 --env NEO4J_USERNAME=neo4j --env "NEO4J_PASSWORD=Your@Password!@#" --env NEO4J_DATABASE=neo4j -- uvx mcp-neo4j-cypher@0.2.3
```

---

## 🔍 Important Files & Locations

### Configuration File

**Where:** `~/.claude.json`

This file stores all your MCP server settings. Claude reads this when it starts.

**To see it:**

```bash
cat ~/.claude.json | grep -A 15 neo4j-cypher
```

---

### Database Connection

**What:** Neo4j is running at `bolt://localhost:7687`

**Login:**

- Username: `neo4j`
- Password: `Your@Password!@#`
- Database: `neo4j`

**Browser UI:** http://localhost:7474

---

## 🚀 Next Adventures

### 1. Try More Queries

Experiment with different questions:

```
Show me the oldest movie
Which person acted in the most movies?
Find movies without a tagline
```

---

### 2. Learn Cypher

The database language is called **Cypher**. Here are some patterns:

**Find things:**

```cypher
MATCH (m:Movie)
RETURN m.title
```

**Find connections:**

```cypher
MATCH (p:Person)-[:ACTED_IN]->(m:Movie)
RETURN p.name, m.title
```

**Create new data:**

```cypher
CREATE (m:Movie {title: "New Movie", released: 2024})
```

---

### 3. Add More Data

Ask Claude:

```
Add a new actor named "John Doe" who acted in "The Matrix"
```

Claude will use the `write-neo4j-cypher` tool to do it!

---

### 4. Explore Other MCP Servers

Just like Neo4j, there are MCP servers for:

- PostgreSQL (Neon)
- GitHub
- Slack
- Google Drive
- And hundreds more!

---

## 📊 Quick Reference Card

### Commands You'll Use

```bash
# Start Claude
claude

# List MCP servers
claude mcp list

# Start Neo4j
docker start neo4j

# Check Neo4j status
docker ps | grep neo4j

# View Neo4j in browser
open http://localhost:7474
```

---

### Questions You Can Ask Claude

**Easy:**

- Show me all movies
- How many people are in the database?
- What's the newest movie?

**Medium:**

- Find actors who worked with Tom Hanks
- Show me movies from the 1990s
- List all directors

**Hard:**

- What's the shortest path between Actor A and Actor B?
- Which actor has worked with the most directors?
- Find actors who never worked together

---

## 🎉 Congratulations!

You've successfully:

✅ Installed the Neo4j MCP server
✅ Connected Claude to your database
✅ Asked questions in plain English
✅ Got instant answers without writing code!

**You're now a database wizard! 🧙‍♂️**

---

---

## 🎨 What Makes This Cool?

### Before MCP:

```
Person: I want to find all movies
↓
Writes Go code (30 lines)
↓
Compiles code
↓
Runs program
↓
Gets answer
Total time: 5-10 minutes
```

### After MCP:

```
Person: Show me all movies
↓
Claude uses MCP
↓
Gets answer
Total time: 3 seconds
```

**That's 100x faster!** ⚡

---

## 🌟 Fun Facts

**Did you know?**

1. **MCP** stands for "Model Context Protocol" - it's a way for AI to talk to other programs
2. **Neo4j** is a graph database - it stores data as connected nodes (like a social network!)
3. **Cypher** is named after the character from The Matrix (because Neo4j stores movie data!)
4. **uvx** automatically downloads and runs Python programs - no installation needed!
5. The **bolt** protocol (bolt://localhost:7687) is super fast - it's designed for graph queries!

---

## 📝 Summary in 3 Sentences

1. We installed a **translator** (MCP server) that helps Claude talk to Neo4j
2. Now you can ask questions in **plain English** instead of writing code
3. Claude gets answers **instantly** from your database

**That's it! You did it! 🎊**

---

**Happy querying! 🚀**

---

*Made with ❤️ for learners of all ages*

*Version 1.0 | Last updated: 2025*

## 🎓 Deep Dive: The Three MCP Tools

Let's understand each tool in detail:

### Tool #1: get-neo4j-schema 📊

**What it does:** Shows you the structure of your database

**When to use it:**
- First time exploring a new database
- Before writing queries (to know what's available)
- To understand relationships between data

**Example conversation:**

```
You: "What's the structure of my Neo4j database?"

Claude: "I'll check the schema for you..."
[Uses get-neo4j-schema tool]

Claude: "Your database has:

Node Labels:
- Movie (properties: title, released, tagline)
- Person (properties: name, born)

Relationships:
- Person -[:ACTED_IN]-> Movie
- Person -[:DIRECTED]-> Movie
```

**Why it's useful:** Prevents errors! You'll know exactly what data exists before querying.

---

### Tool #2: read-neo4j-cypher 📖

**What it does:** Reads data from database (safe - won't change anything!)

**When to use it:**
- Searching for information
- Getting lists or counts
- Exploring data
- Analyzing patterns

**Safety level:** 🟢 SAFE - Auto-approved, can't modify data

**Example conversation:**

```
You: "Show me the 5 oldest movies"

Claude: "I'll query that for you..."
[Uses read-neo4j-cypher tool]
[Executes: MATCH (m:Movie) RETURN m.title, m.released ORDER BY m.released LIMIT 5]

Claude: "Here are the 5 oldest movies:
1. The Matrix (1999)
2. The Green Mile (1999)
3. Snow Falling on Cedars (1999)
4. You've Got Mail (1998)
5. As Good as It Gets (1997)"
```

**Pro tip:** You can ask anything that starts with "show", "find", "list", "count", "what", "which", "who" - all read-only!

---

### Tool #3: write-neo4j-cypher ✍️

**What it does:** Creates, updates, or deletes data

**When to use it:**
- Adding new movies or actors
- Updating information
- Creating relationships
- Deleting data

**Safety level:** 🟡 REQUIRES APPROVAL - Claude will ask first!

**Example conversation:**

```
You: "Add a new movie called 'Star Wars' released in 1977"

Claude: "I'll create that movie. Here's the query I'll run:

CREATE (m:Movie {title: 'Star Wars', released: 1977})
RETURN m

⚠️  This will modify your database. Approve? (yes/no)"

You: "yes"

Claude: "✅ Movie added successfully!
Created: Star Wars (1977)"
```

**Important:** Claude always shows you the query BEFORE running it. Review it carefully!

---

## 💡 Real-World Query Examples

### Example Category 1: Basic Exploration 🔍

**Query 1: Count everything**

```
You: "How many movies and actors are in the database?"

Claude's query:
MATCH (m:Movie)
WITH count(m) as movieCount
MATCH (p:Person)
RETURN movieCount, count(p) as personCount

Result:
- 10 movies
- 7 people
```

---

**Query 2: List with details**

```
You: "Show me all movies with their release year and tagline"

Claude's query:
MATCH (m:Movie)
RETURN m.title, m.released, m.tagline
ORDER BY m.released DESC

Result:
Cloud Atlas (2012) - "Everything is connected"
The Da Vinci Code (2006) - null
...
```

---

### Example Category 2: Finding Relationships 🔗

**Query 3: Actor's filmography**

```
You: "What movies did Tom Hanks act in?"

Claude's query:
MATCH (p:Person {name: 'Tom Hanks'})-[:ACTED_IN]->(m:Movie)
RETURN m.title, m.released
ORDER BY m.released

Result:
- Cloud Atlas (2012)
- The Da Vinci Code (2006)
```

---

**Query 4: Movie's cast**

```
You: "Who acted in The Matrix?"

Claude's query:
MATCH (p:Person)-[:ACTED_IN]->(m:Movie {title: 'The Matrix'})
RETURN p.name

Result:
- Emil Eifrem
- Hugo Weaving
- Carrie-Anne Moss
- Keanu Reeves
```

---

**Query 5: Co-stars**

```
You: "Find actors who worked with Keanu Reeves"

Claude's query:
MATCH (keanu:Person {name: 'Keanu Reeves'})-[:ACTED_IN]->(m:Movie)<-[:ACTED_IN]-(costar:Person)
WHERE keanu <> costar
RETURN DISTINCT costar.name, collect(m.title) as movies

Result:
- Emil Eifrem (The Matrix)
- Hugo Weaving (The Matrix, Cloud Atlas)
- Carrie-Anne Moss (The Matrix)
```

---

### Example Category 3: Advanced Queries 🚀

**Query 6: Multi-talented people**

```
You: "Who both acted in AND directed movies?"

Claude's query:
MATCH (p:Person)-[:ACTED_IN]->(:Movie)
MATCH (p)-[:DIRECTED]->(:Movie)
RETURN DISTINCT p.name

Result:
- People who are both actors and directors
```

---

**Query 7: Shortest path**

```
You: "What's the shortest path between Tom Hanks and Keanu Reeves?"

Claude's query:
MATCH path = shortestPath(
  (tom:Person {name: 'Tom Hanks'})-[*]-(keanu:Person {name: 'Keanu Reeves'})
)
RETURN path

Result:
Tom Hanks → Cloud Atlas ← Hugo Weaving → The Matrix ← Keanu Reeves
(2 degrees of separation!)
```

---

**Query 8: Aggregations**

```
You: "Which person has the most movie connections?"

Claude's query:
MATCH (p:Person)-[:ACTED_IN|DIRECTED]-(m:Movie)
RETURN p.name, count(DISTINCT m) as movieCount
ORDER BY movieCount DESC
LIMIT 5

Result:
1. Hugo Weaving - 3 movies
2. Tom Hanks - 2 movies
...
```

---

## 🛠️ Modifying Data Safely

### Creating Data Examples

**Add a Movie:**

```
You: "Add a new movie 'Inception' released in 2010"

Claude shows:
CREATE (m:Movie {
  title: 'Inception',
  released: 2010,
  tagline: 'Your mind is the scene of the crime'
})
RETURN m

[Asks for approval] ✓
```

---

**Add a Person:**

```
You: "Create a new actor named Leonardo DiCaprio born in 1974"

Claude shows:
CREATE (p:Person {
  name: 'Leonardo DiCaprio',
  born: 1974
})
RETURN p

[Asks for approval] ✓
```

---

**Create Relationship:**

```
You: "Leonardo DiCaprio acted in Inception"

Claude shows:
MATCH (p:Person {name: 'Leonardo DiCaprio'})
MATCH (m:Movie {title: 'Inception'})
CREATE (p)-[:ACTED_IN]->(m)
RETURN p, m

[Asks for approval] ✓
```

---

### Updating Data Examples

**Update Movie Info:**

```
You: "Update The Matrix tagline to 'Free your mind'"

Claude shows:
MATCH (m:Movie {title: 'The Matrix'})
SET m.tagline = 'Free your mind'
RETURN m

[Asks for approval] ✓
```

---

**Update Person Info:**

```
You: "Tom Hanks was born in 1956"

Claude shows:
MATCH (p:Person {name: 'Tom Hanks'})
SET p.born = 1956
RETURN p

[Asks for approval] ✓
```

---

### Deleting Data Examples

**Remove Relationship:**

```
You: "Remove the ACTED_IN relationship between Tom Hanks and Cloud Atlas"

Claude shows:
MATCH (p:Person {name: 'Tom Hanks'})-[r:ACTED_IN]->(m:Movie {title: 'Cloud Atlas'})
DELETE r

[Asks for approval] ✓
```

---

**Delete Node:**

```
You: "Delete the movie Inception"

Claude shows:
MATCH (m:Movie {title: 'Inception'})
DETACH DELETE m

[Asks for approval] ✓

Note: DETACH deletes all relationships too!
```

---

## 🎯 Query Patterns Cheat Sheet

### Pattern 1: Find All Nodes

```cypher
MATCH (n:NodeLabel)
RETURN n
```

**Example:**
```
"Show me all movies" → MATCH (m:Movie) RETURN m
```

---

### Pattern 2: Find by Property

```cypher
MATCH (n:NodeLabel {property: 'value'})
RETURN n
```

**Example:**
```
"Find The Matrix" → MATCH (m:Movie {title: 'The Matrix'}) RETURN m
```

---

### Pattern 3: Find with Condition

```cypher
MATCH (n:NodeLabel)
WHERE n.property > value
RETURN n
```

**Example:**
```
"Movies after 2000" → MATCH (m:Movie) WHERE m.released > 2000 RETURN m
```

---

### Pattern 4: Find Relationships

```cypher
MATCH (a:LabelA)-[:RELATIONSHIP]->(b:LabelB)
RETURN a, b
```

**Example:**
```
"Who acted in what" → MATCH (p:Person)-[:ACTED_IN]->(m:Movie) RETURN p.name, m.title
```

---

### Pattern 5: Count Things

```cypher
MATCH (n:NodeLabel)
RETURN count(n)
```

**Example:**
```
"How many movies" → MATCH (m:Movie) RETURN count(m)
```

---

### Pattern 6: Create Node

```cypher
CREATE (n:NodeLabel {property: 'value'})
RETURN n
```

**Example:**
```
"Add a movie" → CREATE (m:Movie {title: 'New Movie', released: 2024}) RETURN m
```

---

### Pattern 7: Create Relationship

```cypher
MATCH (a), (b)
WHERE a.property = 'value1' AND b.property = 'value2'
CREATE (a)-[:RELATIONSHIP]->(b)
```

**Example:**
```
MATCH (p:Person {name: 'Actor'}), (m:Movie {title: 'Film'})
CREATE (p)-[:ACTED_IN]->(m)
```

---

### Pattern 8: Update Property

```cypher
MATCH (n:NodeLabel {property: 'value'})
SET n.property = 'new value'
RETURN n
```

**Example:**
```
"Update tagline" → MATCH (m:Movie {title: 'The Matrix'}) SET m.tagline = 'New tagline' RETURN m
```

---

## 🎮 Interactive Tutorial Session

Try this step-by-step tutorial in your Claude session:

### Level 1: Getting Started (5 minutes)

1. **Start Claude:**
   ```bash
   claude
   ```

2. **Check what's in the database:**
   ```
   What's in my Neo4j database?
   ```

3. **Count the data:**
   ```
   How many movies and people are there?
   ```

4. **See all movies:**
   ```
   List all movies with their release years
   ```

---

### Level 2: Exploring Connections (10 minutes)

5. **Pick a movie and find actors:**
   ```
   Who acted in The Matrix?
   ```

6. **Pick an actor and find their movies:**
   ```
   What movies did Keanu Reeves act in?
   ```

7. **Find co-stars:**
   ```
   Which actors worked with Tom Hanks?
   ```

8. **Find directors:**
   ```
   Show me all directors and their movies
   ```

---

### Level 3: Advanced Exploration (15 minutes)

9. **Find connections:**
   ```
   What's the shortest path between Tom Hanks and Keanu Reeves?
   ```

10. **Aggregate data:**
    ```
    Which actor appeared in the most movies?
    ```

11. **Filter results:**
    ```
    Show me movies from 1999 to 2005
    ```

12. **Complex query:**
    ```
    Find people who both acted in and directed movies
    ```

---

### Level 4: Making Changes (20 minutes)

13. **Add a movie:**
    ```
    Add a new movie called 'Inception' released in 2010
    ```

14. **Add an actor:**
    ```
    Create an actor named Leonardo DiCaprio
    ```

15. **Connect them:**
    ```
    Leonardo DiCaprio acted in Inception
    ```

16. **Verify:**
    ```
    Show me all movies Leonardo DiCaprio acted in
    ```

17. **Update info:**
    ```
    Set Inception's tagline to "Your mind is the scene of the crime"
    ```

18. **Final verification:**
    ```
    Show me the complete details for Inception
    ```

---

## 🔄 Common Workflows

### Workflow A: Exploring a New Database

```
Step 1: Get schema
"What's the structure of this database?"

Step 2: Count data
"How much data is here?"

Step 3: Sample data
"Show me 5 examples of each type"

Step 4: Check relationships
"What relationships exist between the data?"

Step 5: Test queries
"Find [specific example]"
```

---

### Workflow B: Adding New Data

```
Step 1: Check if exists
"Is there a movie called X?"

Step 2: Create it
"Add movie X released in YEAR"

Step 3: Verify
"Show me movie X"

Step 4: Add relationships
"Actor Y acted in movie X"

Step 5: Final check
"Show me all actors in movie X"
```

---

### Workflow C: Data Quality Check

```
Step 1: Find incomplete data
"Show me movies without a tagline"

Step 2: Count problems
"How many people don't have a birth year?"

Step 3: Fix issues
"Set the tagline for movie X to Y"

Step 4: Verify fix
"Show me movie X's details"

Step 5: Recount
"How many movies without taglines now?"
```

---

## 🧪 Experiment Ideas

### Experiment 1: Build a Movie Network

1. Add 3 new movies from different years
2. Add 2 new actors
3. Connect each actor to at least 2 movies
4. Find which actor has the most connections
5. Visualize the connections (ask Claude to describe them)

---

### Experiment 2: Time Travel Analysis

1. Count movies by decade
2. Find the oldest movie
3. Find the newest movie
4. See which decade had the most movies
5. Find actors who worked across multiple decades

---

### Experiment 3: Relationship Mapping

1. Pick an actor
2. Find all their movies
3. Find all co-stars
4. Map out the network
5. Find the most connected actor

---

## 📚 Learning Resources

### Understanding Cypher

Cypher is Neo4j's query language. Here are the basics:

**Nodes** (things):
```cypher
(m:Movie)           // A movie node
(p:Person)          // A person node
(n)                 // Any node
```

**Relationships** (connections):
```cypher
-[:ACTED_IN]->      // Acted in (direction matters)
-[:DIRECTED]->      // Directed
-[r]->              // Any relationship
```

**Properties** (details):
```cypher
{title: 'The Matrix'}      // Movie title
{name: 'Keanu Reeves'}     // Person name
{born: 1964}               // Birth year
```

**Patterns** (combinations):
```cypher
(p:Person)-[:ACTED_IN]->(m:Movie)
// Person acted in movie

(p)-[:ACTED_IN]->(m)<-[:DIRECTED]-(d)
// Person acted in movie, someone else directed it
```

---

### Query Building Blocks

**MATCH** - Find things
```cypher
MATCH (m:Movie)
```

**WHERE** - Filter
```cypher
WHERE m.released > 2000
```

**RETURN** - What to show
```cypher
RETURN m.title, m.released
```

**ORDER BY** - Sort results
```cypher
ORDER BY m.released DESC
```

**LIMIT** - Limit results
```cypher
LIMIT 10
```

**CREATE** - Make new things
```cypher
CREATE (m:Movie {title: 'New Movie'})
```

**SET** - Update properties
```cypher
SET m.tagline = 'New tagline'
```

**DELETE** - Remove things
```cypher
DELETE r  // Delete relationship
DETACH DELETE m  // Delete node and relationships
```

---

## 🎁 Bonus: Advanced Features

### Feature 1: Using Parameters

Instead of:
```
"Find movie The Matrix"
```

Claude can use parameters:
```cypher
MATCH (m:Movie {title: $movieTitle})
RETURN m
```

This is safer and faster!

---

### Feature 2: Multiple Conditions

```
"Find movies from the 1990s with good ratings"
```

Becomes:
```cypher
MATCH (m:Movie)
WHERE m.released >= 1990
  AND m.released < 2000
  AND m.rating > 7
RETURN m
```

---

### Feature 3: Aggregations

```
"How many movies per year?"
```

Becomes:
```cypher
MATCH (m:Movie)
RETURN m.released as year, count(m) as movieCount
ORDER BY year
```

---

### Feature 4: Collections

```
"Show me each actor with all their movies"
```

Becomes:
```cypher
MATCH (p:Person)-[:ACTED_IN]->(m:Movie)
RETURN p.name, collect(m.title) as movies
```

---

## 🎊 You're Now a Pro!
