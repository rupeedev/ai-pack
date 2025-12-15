# Movie Genre Finder 🎬

A Python command-line application that finds the top 5 movies in any genre, ordered by IMDB rating.

Built following the **Neo4j MCP Tools** tutorial using the Model Context Protocol for schema discovery and query validation.

---

## Features

- ✅ Query Neo4j database by genre
- ✅ Returns top 5 movies ordered by IMDB rating
- ✅ Filters out movies without ratings
- ✅ Beautiful command-line interface
- ✅ Environment-based configuration
- ✅ Error handling and validation

---

## Prerequisites

- Python 3.7+
- Neo4j database running (local or remote)
- Neo4j Python driver

---

## Installation

1. **Clone or navigate to the project directory:**
   ```bash
   cd /Users/rupeshpanwar/Documents/PProject/AI-Infra/neo4j-mcp/movie_genre_app
   ```

2. **Install dependencies:**
   ```bash
   pip install -r requirements.txt
   ```

3. **Configure environment variables:**

   The `.env` file is already created with:
   ```bash
   NEO4J_URI=bolt://localhost:7687
   NEO4J_USERNAME=neo4j
   NEO4J_PASSWORD=Your@Password!@#
   ```

   Update these values if your Neo4j instance has different credentials.

---

## Usage

### Run the Application

```bash
python app.py
```

### Example Session

```
============================================================
🎬  MOVIE GENRE FINDER
============================================================

Find the top 5 movies in any genre!

Available genres: Mystery, Action, Sci-Fi, Drama,
                  Thriller, Animation, Adventure

Enter a movie genre: Mystery

🎬 Top 5 movies in 'Mystery' genre:
============================================================
Rank   Title                                    Rating
------------------------------------------------------------
🥇 1   The Dark Knight                          9.0
🥈 2   Inception                                8.8
🥉 3   The Green Mile                           8.6
   4   Cloud Atlas                              7.4
   5   The Da Vinci Code                        6.6
============================================================
```

---

## Available Genres

The database currently includes these genres:

- **Action** - High-energy movies with intense sequences
- **Sci-Fi** - Science fiction and futuristic themes
- **Drama** - Character-driven emotional stories
- **Thriller** - Suspenseful and tense narratives
- **Mystery** - Puzzles and investigations
- **Animation** - Animated films
- **Adventure** - Exciting journeys and quests

---

## How It Works

### 1. Schema Discovery (Using MCP)

The application was built using Neo4j MCP tools to discover the database schema:

```cypher
MATCH (m:Movie)-[:IN_GENRE]->(g:Genre)
RETURN m, g
```

**Schema:**
- `Movie` nodes with properties: `title`, `imdbRating`, `released`, `tagline`
- `Genre` nodes with property: `name`
- `IN_GENRE` relationship connecting Movies to Genres

### 2. Query Validation (Using MCP)

The Cypher query was tested with real data before implementation:

```cypher
MATCH (m:Movie)-[:IN_GENRE]->(g:Genre {name: $genreName})
WHERE m.imdbRating IS NOT NULL
RETURN m.title AS title, m.imdbRating AS rating
ORDER BY m.imdbRating DESC
LIMIT 5
```

**Validation Results (Mystery genre):**
- ✅ Returns exactly 5 movies
- ✅ Ordered by rating (highest first)
- ✅ Filters out null ratings
- ✅ Works with parameterized input

### 3. Python Implementation

The app uses the Neo4j Python driver to:
1. Load environment variables
2. Connect to Neo4j
3. Execute the validated Cypher query
4. Display formatted results
5. Handle errors gracefully

---

## Project Structure

```
movie_genre_app/
├── .env                # Environment variables (Neo4j credentials)
├── app.py              # Main application
├── requirements.txt    # Python dependencies
└── README.md          # This file
```

---

## Database Schema

### Nodes

**Movie:**
```
{
  title: STRING,
  imdbRating: FLOAT,
  released: INTEGER,
  tagline: STRING
}
```

**Genre:**
```
{
  name: STRING
}
```

### Relationships

```
(Movie)-[:IN_GENRE]->(Genre)
```

---

## Example Movies by Genre

### Mystery
1. The Dark Knight (9.0)
2. Inception (8.8)
3. The Green Mile (8.6)
4. Cloud Atlas (7.4)
5. The Da Vinci Code (6.6)

### Sci-Fi
1. The Dark Knight (9.0)
2. Inception (8.8)
3. The Matrix (8.7)
4. Interstellar (8.7)
5. Cloud Atlas (7.4)

### Action
1. The Dark Knight (9.0)
2. The Matrix (8.7)
3. V for Vendetta (8.2)
4. The Matrix Reloaded (7.2)
5. The Matrix Revolutions (6.7)

---

## Error Handling

The app handles:
- ✅ Missing environment variables
- ✅ Empty genre input
- ✅ No movies found in genre
- ✅ Database connection errors
- ✅ Keyboard interrupts (Ctrl+C)

---

## Development Notes

### Built with MCP Tools

This app was developed following the **MCP (Model Context Protocol)** approach:

1. **Schema Discovery** - Used `mcp__neo4j-cypher__read_neo4j_cypher` to explore database structure
2. **Query Testing** - Validated queries on real data before coding
3. **Implementation** - Created Python app with validated Cypher

### Why This Approach Works

- ✅ **Accurate Queries** - Schema discovered from real database
- ✅ **Tested First** - Queries validated before implementation
- ✅ **No Hallucinations** - AI doesn't invent non-existent labels/properties
- ✅ **Reliable Results** - Consistent output from database

---

## Troubleshooting

### "No movies found"

**Problem:** Genre doesn't exist or has no movies

**Solution:** Try these genres: Mystery, Action, Sci-Fi, Drama, Thriller, Animation, Adventure

---

### "Missing environment variables"

**Problem:** .env file not found or incomplete

**Solution:** Create `.env` file with:
```bash
NEO4J_URI=bolt://localhost:7687
NEO4J_USERNAME=neo4j
NEO4J_PASSWORD=Your@Password!@#
```

---

### "Connection refused"

**Problem:** Neo4j not running

**Solution:** Start Neo4j:
```bash
docker start neo4j
# or
docker-compose up -d
```

---

## Extending the App

### Add More Genres

```cypher
MERGE (g:Genre {name: 'Comedy'})
```

### Connect Movies to Genres

```cypher
MATCH (m:Movie {title: 'Some Movie'})
MATCH (g:Genre {name: 'Comedy'})
MERGE (m)-[:IN_GENRE]->(g)
```

### Change Result Limit

In `app.py`, modify the `LIMIT` clause:
```python
LIMIT 10  # Show top 10 instead of 5
```

---

## License

MIT License - Built for educational purposes following Neo4j MCP Tools tutorial.

---

## Credits

- Built with **Neo4j Python Driver**
- Follows **MCP (Model Context Protocol)** best practices
- Based on **GraphAcademy Neo4j MCP Tools** course

---

**Happy movie hunting! 🎬**
