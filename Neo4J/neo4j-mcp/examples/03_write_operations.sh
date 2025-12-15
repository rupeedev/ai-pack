#!/bin/bash
# Write Operations with Neo4j MCP using Claude CLI
#
# ⚠️  WARNING: These operations will MODIFY your database!
# Claude will ask for approval before executing each write operation.

echo "✍️  Neo4j MCP - Write Operations Demo"
echo "======================================"
echo ""
echo "⚠️  WARNING: These operations will modify your database!"
echo "Claude will ask for approval before each operation."
echo ""
read -p "Continue? (yes/no): " confirm

if [ "$confirm" != "yes" ]; then
    echo "Cancelled."
    exit 0
fi

echo ""

# Example 1: Create a movie
echo "1️⃣  Creating a new movie: Inception..."
echo "---"
claude "Add a new movie to my Neo4j database: title='Inception', released=2010, tagline='Your mind is the scene of the crime'"
echo ""
echo "✅ Done"
echo ""

# Example 2: Create an actor
echo "2️⃣  Creating a new actor: Leonardo DiCaprio..."
echo "---"
claude "Add a new person to my Neo4j database: name='Leonardo DiCaprio', born=1974"
echo ""
echo "✅ Done"
echo ""

# Example 3: Create relationship
echo "3️⃣  Connecting Leonardo DiCaprio to Inception..."
echo "---"
claude "In my Neo4j database, create a relationship: Leonardo DiCaprio acted in Inception"
echo ""
echo "✅ Done"
echo ""

# Example 4: Verify the additions
echo "4️⃣  Verifying the new data..."
echo "---"
claude "Show me all movies Leonardo DiCaprio acted in from my Neo4j database"
echo ""
echo "✅ Done"
echo ""

# Example 5: Update a property
echo "5️⃣  Updating movie tagline..."
echo "---"
claude "Update the movie 'The Matrix' in my Neo4j database: set tagline to 'Free your mind'"
echo ""
echo "✅ Done"
echo ""

# Example 6: Verify the update
echo "6️⃣  Verifying the update..."
echo "---"
claude "Show me The Matrix movie details from my Neo4j database including the tagline"
echo ""
echo "✅ Done"
echo ""

echo "🎉 All write operations completed!"
echo ""
echo "💡 Tip: You can verify all changes by running:"
echo "   claude 'Show me all movies and actors in my Neo4j database'"
