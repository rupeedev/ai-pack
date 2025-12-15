#!/usr/bin/env python3
"""
Movie Genre Finder

A command-line application that finds the top 5 movies in a given genre
ordered by IMDB rating.

Usage:
    python app.py

Then enter a genre name when prompted (e.g., Mystery, Action, Sci-Fi, Drama)
"""

import os
import sys
from neo4j import GraphDatabase
from dotenv import load_dotenv


class MovieGenreApp:
    """Application to find top movies by genre"""

    def __init__(self, uri, username, password):
        """Initialize Neo4j driver connection"""
        self.driver = GraphDatabase.driver(uri, auth=(username, password))

    def close(self):
        """Close the driver connection"""
        self.driver.close()

    def get_top_movies_by_genre(self, genre_name):
        """
        Query Neo4j for top 5 movies in the given genre.

        Args:
            genre_name (str): The genre to search for

        Returns:
            list: List of dictionaries containing movie title and rating
        """
        query = """
        MATCH (m:Movie)-[:IN_GENRE]->(g:Genre {name: $genreName})
        WHERE m.imdbRating IS NOT NULL
        RETURN m.title AS title, m.imdbRating AS rating
        ORDER BY m.imdbRating DESC
        LIMIT 5
        """

        with self.driver.session() as session:
            result = session.run(query, genreName=genre_name)
            return [record.data() for record in result]

    def display_movies(self, movies, genre):
        """
        Display the movies in a formatted table.

        Args:
            movies (list): List of movie dictionaries
            genre (str): The genre name
        """
        if not movies:
            print(f"\n❌ No movies found in the '{genre}' genre.")
            print("   Try: Mystery, Action, Sci-Fi, Drama, Thriller, Animation, Adventure")
            return

        print(f"\n🎬 Top {len(movies)} movies in '{genre}' genre:")
        print("=" * 60)
        print(f"{'Rank':<6} {'Title':<40} {'Rating':<10}")
        print("-" * 60)

        for i, movie in enumerate(movies, 1):
            title = movie['title']
            rating = movie['rating']

            # Add medal emojis for top 3
            if i == 1:
                rank = "🥇 1"
            elif i == 2:
                rank = "🥈 2"
            elif i == 3:
                rank = "🥉 3"
            else:
                rank = f"   {i}"

            print(f"{rank:<6} {title:<40} {rating:<10.1f}")

        print("=" * 60)


def main():
    """Main application entry point"""
    # Load environment variables from .env file
    load_dotenv()

    # Get Neo4j connection details from environment
    uri = os.getenv('NEO4J_URI')
    username = os.getenv('NEO4J_USERNAME')
    password = os.getenv('NEO4J_PASSWORD')

    # Validate environment variables
    if not all([uri, username, password]):
        print("❌ Error: Missing environment variables!")
        print("   Please ensure .env file contains:")
        print("   - NEO4J_URI")
        print("   - NEO4J_USERNAME")
        print("   - NEO4J_PASSWORD")
        sys.exit(1)

    # Create app instance
    app = MovieGenreApp(uri, username, password)

    try:
        # Welcome message
        print("\n" + "=" * 60)
        print("🎬  MOVIE GENRE FINDER")
        print("=" * 60)
        print("\nFind the top 5 movies in any genre!")
        print("\nAvailable genres: Mystery, Action, Sci-Fi, Drama,")
        print("                  Thriller, Animation, Adventure")
        print()

        # Get user input
        genre = input("Enter a movie genre: ").strip()

        if not genre:
            print("\n❌ Error: Genre cannot be empty!")
            sys.exit(1)

        # Query database
        movies = app.get_top_movies_by_genre(genre)

        # Display results
        app.display_movies(movies, genre)

    except KeyboardInterrupt:
        print("\n\n👋 Goodbye!")
        sys.exit(0)

    except Exception as e:
        print(f"\n❌ Error: {e}")
        sys.exit(1)

    finally:
        # Always close the driver connection
        app.close()


if __name__ == "__main__":
    main()
