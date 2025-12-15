#!/bin/bash
# Basic Neo4j MCP Queries using Claude CLI
#
# This script demonstrates basic read-only queries using the Neo4j MCP server
# All queries are safe and won't modify your database

echo "🎬 Neo4j MCP - Basic Queries Demo"
echo "=================================="
echo ""

# Example 1: Get database schema
echo "1️⃣  Getting database schema..."
echo "---"
claude "What's the structure of my Neo4j database? Show node labels and relationships."
echo ""
echo "✅ Done"
echo ""

# Example 2: Count nodes
echo "2️⃣  Counting movies and actors..."
echo "---"
claude "How many movies and people are in my Neo4j database?"
echo ""
echo "✅ Done"
echo ""

# Example 3: List movies
echo "3️⃣  Listing all movies..."
echo "---"
claude "List all movies in my Neo4j database with their release years, ordered by year."
echo ""
echo "✅ Done"
echo ""

# Example 4: Find actors in a movie
echo "4️⃣  Finding actors in The Matrix..."
echo "---"
claude "Who acted in The Matrix? List all actors."
echo ""
echo "✅ Done"
echo ""

# Example 5: Find movies by actor
echo "5️⃣  Finding Tom Hanks movies..."
echo "---"
claude "What movies did Tom Hanks act in? Show title and year."
echo ""
echo "✅ Done"
echo ""

echo "🎉 All basic queries completed!"
