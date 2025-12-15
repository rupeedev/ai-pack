#!/bin/bash
# Advanced Neo4j MCP Queries using Claude CLI
#
# This script demonstrates complex graph queries and analytics

echo "🚀 Neo4j MCP - Advanced Queries Demo"
echo "====================================="
echo ""

# Example 1: Find co-stars
echo "1️⃣  Finding actors who worked with Keanu Reeves..."
echo "---"
claude "Find all actors who worked with Keanu Reeves in any movie. Show actor names and the movies they shared."
echo ""
echo "✅ Done"
echo ""

# Example 2: Shortest path
echo "2️⃣  Finding connection between Tom Hanks and Keanu Reeves..."
echo "---"
claude "What's the shortest path between Tom Hanks and Keanu Reeves in my Neo4j database? Show the connection through movies and actors."
echo ""
echo "✅ Done"
echo ""

# Example 3: Multi-talented people
echo "3️⃣  Finding people who both acted and directed..."
echo "---"
claude "Find people who both acted in AND directed movies in my Neo4j database."
echo ""
echo "✅ Done"
echo ""

# Example 4: Movies by decade
echo "4️⃣  Analyzing movies by decade..."
echo "---"
claude "Count movies by decade in my Neo4j database and show the results."
echo ""
echo "✅ Done"
echo ""

# Example 5: Most connected actor
echo "5️⃣  Finding the most connected actor..."
echo "---"
claude "Which actor has appeared in the most movies in my Neo4j database? Show top 5."
echo ""
echo "✅ Done"
echo ""

# Example 6: Database analytics
echo "6️⃣  Getting database statistics..."
echo "---"
claude "Give me complete analytics about my Neo4j movie database: total movies, actors, directors, year range, etc."
echo ""
echo "✅ Done"
echo ""

echo "🎉 All advanced queries completed!"
