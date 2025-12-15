#!/usr/bin/env python3
"""
Interactive Neo4j MCP Demo for Claude CLI

This script demonstrates how to use Claude CLI with Neo4j MCP tools
to build an interactive movie database application.
"""

import subprocess
import sys
import json


class Neo4jMCPDemo:
    """Interactive demo using Claude CLI and Neo4j MCP"""

    def __init__(self):
        self.examples = {
            '1': ('Get Database Schema', self.example_schema),
            '2': ('Count Movies and Actors', self.example_count),
            '3': ('List All Movies', self.example_list_movies),
            '4': ('Find Actors in The Matrix', self.example_matrix_actors),
            '5': ('Find Tom Hanks Movies', self.example_tom_hanks),
            '6': ('Find Co-Stars', self.example_costars),
            '7': ('Shortest Path Between Actors', self.example_shortest_path),
            '8': ('Database Analytics', self.example_analytics),
            '9': ('Add New Movie (Write)', self.example_add_movie),
            '10': ('Add Actor and Relationship (Write)', self.example_add_actor),
        }

    def run_claude_query(self, question):
        """Execute a Claude CLI query"""
        print(f"\n💬 Question: {question}")
        print("🤖 Claude is processing...")
        print("-" * 70)

        try:
            result = subprocess.run(
                ['claude', question],
                capture_output=True,
                text=True,
                timeout=30
            )

            if result.returncode == 0:
                print(result.stdout)
            else:
                print(f"❌ Error: {result.stderr}")

        except subprocess.TimeoutExpired:
            print("⏱️  Query timed out after 30 seconds")
        except FileNotFoundError:
            print("❌ Error: 'claude' command not found. Make sure Claude CLI is installed.")
        except Exception as e:
            print(f"❌ Error: {e}")

        print("-" * 70)

    def example_schema(self):
        """Example 1: Get database schema"""
        print("\n" + "="*70)
        print("EXAMPLE 1: Get Database Schema")
        print("="*70)
        print("\nThis uses the get_neo4j_schema tool to discover the database structure.")
        print("It shows node labels, properties, and relationships.")
        input("\nPress Enter to run...")

        self.run_claude_query(
            "What's the structure of my Neo4j database? "
            "Show all node labels, their properties, and relationships."
        )

    def example_count(self):
        """Example 2: Count movies and actors"""
        print("\n" + "="*70)
        print("EXAMPLE 2: Count Movies and Actors")
        print("="*70)
        print("\nThis uses read_neo4j_cypher to count nodes in the database.")
        input("\nPress Enter to run...")

        self.run_claude_query(
            "How many movies and people are in my Neo4j database? "
            "Give me exact counts for each."
        )

    def example_list_movies(self):
        """Example 3: List all movies"""
        print("\n" + "="*70)
        print("EXAMPLE 3: List All Movies")
        print("="*70)
        print("\nThis queries all movies with their details.")
        input("\nPress Enter to run...")

        self.run_claude_query(
            "List all movies in my Neo4j database with their title, year, and tagline. "
            "Order by year descending."
        )

    def example_matrix_actors(self):
        """Example 4: Find actors in The Matrix"""
        print("\n" + "="*70)
        print("EXAMPLE 4: Find Actors in The Matrix")
        print("="*70)
        print("\nThis demonstrates relationship traversal: Person-[:ACTED_IN]->Movie")
        input("\nPress Enter to run...")

        self.run_claude_query(
            "Who acted in The Matrix? List all actors' names."
        )

    def example_tom_hanks(self):
        """Example 5: Find Tom Hanks movies"""
        print("\n" + "="*70)
        print("EXAMPLE 5: Find Tom Hanks Movies")
        print("="*70)
        print("\nThis finds all movies for a specific actor.")
        input("\nPress Enter to run...")

        self.run_claude_query(
            "What movies did Tom Hanks act in? Show title and year, ordered by year."
        )

    def example_costars(self):
        """Example 6: Find co-stars"""
        print("\n" + "="*70)
        print("EXAMPLE 6: Find Co-Stars")
        print("="*70)
        print("\nThis demonstrates a more complex graph pattern.")
        input("\nPress Enter to run...")

        self.run_claude_query(
            "Find all actors who worked with Keanu Reeves. "
            "Show the actor names and which movies they shared."
        )

    def example_shortest_path(self):
        """Example 7: Shortest path between actors"""
        print("\n" + "="*70)
        print("EXAMPLE 7: Shortest Path Between Actors")
        print("="*70)
        print("\nThis uses graph algorithms to find connections.")
        input("\nPress Enter to run...")

        self.run_claude_query(
            "What's the shortest path between Tom Hanks and Keanu Reeves "
            "in my Neo4j database? Show the connection through movies and actors."
        )

    def example_analytics(self):
        """Example 8: Database analytics"""
        print("\n" + "="*70)
        print("EXAMPLE 8: Database Analytics")
        print("="*70)
        print("\nThis gets comprehensive statistics about the database.")
        input("\nPress Enter to run...")

        self.run_claude_query(
            "Give me complete analytics about my Neo4j movie database: "
            "total movies, total actors, total directors, year range, "
            "average rating if available."
        )

    def example_add_movie(self):
        """Example 9: Add new movie (write operation)"""
        print("\n" + "="*70)
        print("EXAMPLE 9: Add New Movie (WRITE OPERATION)")
        print("="*70)
        print("\n⚠️  WARNING: This will MODIFY your database!")
        print("Claude will show you the query and ask for approval.")
        confirm = input("\nContinue? (yes/no): ")

        if confirm.lower() != 'yes':
            print("Cancelled.")
            return

        self.run_claude_query(
            "Add a new movie to my Neo4j database: "
            "title='Inception', released=2010, "
            "tagline='Your mind is the scene of the crime'"
        )

    def example_add_actor(self):
        """Example 10: Add actor and relationship (write operation)"""
        print("\n" + "="*70)
        print("EXAMPLE 10: Add Actor and Relationship (WRITE OPERATION)")
        print("="*70)
        print("\n⚠️  WARNING: This will MODIFY your database!")
        print("Claude will create an actor and connect them to a movie.")
        confirm = input("\nContinue? (yes/no): ")

        if confirm.lower() != 'yes':
            print("Cancelled.")
            return

        print("\nStep 1: Create actor...")
        self.run_claude_query(
            "Add a new person to my Neo4j database: "
            "name='Leonardo DiCaprio', born=1974"
        )

        print("\nStep 2: Create relationship...")
        input("Press Enter to continue...")
        self.run_claude_query(
            "In my Neo4j database, create a relationship: "
            "Leonardo DiCaprio acted in Inception"
        )

        print("\nStep 3: Verify...")
        input("Press Enter to verify...")
        self.run_claude_query(
            "Show me all movies Leonardo DiCaprio acted in from my Neo4j database"
        )

    def show_menu(self):
        """Display the main menu"""
        print("\n" + "="*70)
        print("NEO4J MCP - INTERACTIVE DEMO")
        print("="*70)
        print("\nSelect an example to run:\n")

        for key in sorted(self.examples.keys(), key=int):
            title, _ = self.examples[key]
            print(f"  {key:2s}. {title}")

        print("\n   0. Exit")
        print("\n" + "="*70)

    def run(self):
        """Main application loop"""
        print("""
╔════════════════════════════════════════════════════════════════════╗
║                                                                    ║
║         Neo4j MCP Interactive Demo for Claude CLI                 ║
║                                                                    ║
║  Demonstrates the three MCP tools:                                ║
║    • get_neo4j_schema - Discover database structure               ║
║    • read_neo4j_cypher - Read data safely                         ║
║    • write_neo4j_cypher - Modify data with approval               ║
║                                                                    ║
╚════════════════════════════════════════════════════════════════════╝
        """)

        while True:
            self.show_menu()

            try:
                choice = input("\nEnter your choice: ").strip()

                if choice == '0':
                    print("\n👋 Goodbye!")
                    break

                if choice in self.examples:
                    _, func = self.examples[choice]
                    func()
                    input("\n⏎ Press Enter to continue...")
                else:
                    print("\n❌ Invalid choice. Please try again.")

            except KeyboardInterrupt:
                print("\n\n👋 Goodbye!")
                break
            except Exception as e:
                print(f"\n❌ Error: {e}")
                input("\n⏎ Press Enter to continue...")


def main():
    """Main entry point"""
    demo = Neo4jMCPDemo()
    demo.run()


if __name__ == "__main__":
    main()
