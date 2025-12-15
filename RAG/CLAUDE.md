# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a **multi-week learning project** building a production-grade RAG (Retrieval-Augmented Generation) system called "arXiv Paper Curator". The project is organized into weekly subdirectories (week-1 through week-7, plus main), each representing a progressive stage of RAG system development.

**Important:** Each week directory is a standalone version of the project at that stage. The `main` directory contains the most complete implementation (Week 7).

## Project Structure

```
RAG/
├── main/                    # Latest complete implementation (Week 7)
├── week-1/                  # Infrastructure setup
├── week-2/                  # Data ingestion pipeline
├── week-3/                  # BM25 keyword search
├── week-4/                  # Hybrid search (BM25 + vectors)
├── week-5/                  # Complete RAG with LLM
├── week-6/                  # Monitoring and caching
├── week-7/                  # Agentic RAG with Telegram bot
└── References/              # Documentation and references
```

### Core Application Structure (within each week directory)

```
src/
├── routers/              # API endpoints
│   ├── ping.py          # Health checks
│   ├── hybrid_search.py # Search endpoints
│   ├── ask.py           # RAG query endpoints
│   └── agentic_ask.py   # Agentic RAG endpoints (Week 7)
├── services/            # Business logic layer
│   ├── arxiv/           # arXiv API integration
│   ├── pdf_parser/      # PDF parsing with Docling
│   ├── opensearch/      # Search engine client
│   ├── embeddings/      # Jina AI embeddings
│   ├── indexing/        # Document chunking and indexing
│   ├── ollama/          # Local LLM client
│   ├── langfuse/        # Observability tracing
│   ├── cache/           # Redis caching
│   ├── telegram/        # Telegram bot (Week 7)
│   └── agents/          # LangGraph agents (Week 7)
├── models/              # SQLAlchemy database models
├── schemas/             # Pydantic request/response schemas
├── repositories/        # Database access layer
├── config.py            # Environment configuration
├── database.py          # Database connection
└── main.py              # FastAPI application entry
```

## Development Commands

### Using Makefile (Recommended)

```bash
make help         # Show all available commands
make start        # Start all services (docker compose up --build -d)
make stop         # Stop all services
make status       # Show service status
make logs         # Show service logs
make health       # Check all services health

make setup        # Install dependencies (uv sync)
make format       # Format code (ruff format)
make lint         # Lint and type check (ruff check + mypy)
make test         # Run tests (pytest)
make test-cov     # Run tests with coverage

make clean        # Clean up everything (docker compose down -v)
```

### Direct Commands

```bash
# Service management
docker compose up --build -d    # Start all services
docker compose ps               # Check status
docker compose logs -f          # View logs
docker compose down             # Stop services

# Development
uv sync                        # Install dependencies
uv run pytest                  # Run tests
uv run ruff format            # Format code
uv run ruff check --fix       # Lint code
uv run mypy src/              # Type check

# Run Gradio interface (Week 5+)
uv run python gradio_launcher.py

# Run Jupyter notebooks
uv run jupyter notebook notebooks/weekX/
```

## Architecture Patterns

### Service Layer Architecture

The project follows a layered architecture:

1. **Routers** (`src/routers/`) - FastAPI endpoints, request validation
2. **Services** (`src/services/`) - Business logic, external integrations
3. **Repositories** (`src/repositories/`) - Database access
4. **Models** (`src/models/`) - SQLAlchemy ORM models
5. **Schemas** (`src/schemas/`) - Pydantic validation models

### Key Service Components

**OpenSearch Service** (`src/services/opensearch/`):

- Handles BM25 keyword search and vector search
- Index management with proper mappings
- Hybrid search using RRF (Reciprocal Rank Fusion)

**Embeddings Service** (`src/services/embeddings/`):

- Jina AI integration for semantic embeddings
- Fallback strategies for production resilience
- Batch processing support

**Indexing Service** (`src/services/indexing/`):

- Section-based document chunking
- Overlap strategies for context preservation
- Metadata extraction from scientific papers

**Ollama Service** (`src/services/ollama/`):

- Local LLM integration (llama3.2:3b-instruct-fp16)
- Streaming response support
- Optimized prompts in `prompts/rag_system.txt`

**Langfuse Service** (`src/services/langfuse/`):

- End-to-end RAG pipeline tracing
- Performance monitoring and cost tracking
- Custom metadata for search and generation steps

**Cache Service** (`src/services/cache/`):

- Redis-based caching with graceful fallback
- Exact-match caching for identical queries
- TTL management (1 hour default)

**Agents Service** (`src/services/agents/`) - Week 7:

- LangGraph workflow orchestration
- Decision nodes: guardrail, retrieve, grade, rewrite, generate
- Adaptive retrieval with multi-attempt fallback

**Telegram Service** (`src/services/telegram/`) - Week 7:

- Async bot integration
- Command handlers and message processing
- Integration with agentic RAG workflow

### Configuration Management

All configuration is managed through environment variables using `pydantic-settings`:

- `.env.example` provides template with defaults
- `.env` for local development (gitignored)
- `src/config.py` defines typed settings with validation

Key configuration sections:

- Database (PostgreSQL)
- OpenSearch (connection, auth, indices)
- Jina AI (embeddings API)
- Ollama (LLM model, endpoints)
- Langfuse (monitoring keys)
- Redis (caching)
- Telegram (bot token)

## Technology Stack

- **FastAPI** - Async REST API framework
- **PostgreSQL 16** - Paper metadata storage
- **OpenSearch 2.19** - Hybrid search engine
- **Apache Airflow 3.0** - Workflow orchestration
- **Ollama** - Local LLM serving
- **Redis** - High-performance caching
- **Langfuse** - RAG observability
- **Jina AI** - Embedding generation
- **LangGraph** - Agent orchestration (Week 7)
- **python-telegram-bot** - Telegram integration (Week 7)
- **UV** - Fast Python package manager
- **Ruff** - Python linter and formatter
- **MyPy** - Static type checking
- **Pytest** - Testing framework

## Testing

Tests are located in `tests/` directory:

- Integration tests use testcontainers for isolated environments
- Run tests with `make test` or `uv run pytest`
- Coverage reports with `make test-cov`

## Service Access Points

| Service    | URL                        | Purpose                |
| ---------- | -------------------------- | ---------------------- |
| FastAPI    | http://localhost:8000      | Main API               |
| API Docs   | http://localhost:8000/docs | Interactive Swagger UI |
| Gradio     | http://localhost:7861      | Chat interface         |
| Langfuse   | http://localhost:3000      | Monitoring dashboard   |
| Airflow    | http://localhost:8080      | Workflow management    |
| OpenSearch | http://localhost:5601      | Search dashboards      |

## Working with Different Weeks

When switching between weeks:

1. Navigate to the specific week directory
2. Check that week's README.md for specific setup
3. Run `uv sync` to install dependencies
4. Copy `.env.example` to `.env` and configure
5. Start services with `make start` or `docker compose up --build -d`
6. Refer to the corresponding blog post for detailed explanations

Each week builds incrementally:

- Week 1: Infrastructure only
- Week 2: + Data ingestion
- Week 3: + BM25 search
- Week 4: + Hybrid search with embeddings
- Week 5: + RAG with LLM and Gradio UI
- Week 6: + Monitoring and caching
- Week 7: + Agentic RAG with Telegram bot

## Common Development Patterns

### Adding New API Endpoints

1. Create schema in `src/schemas/`
2. Implement service logic in `src/services/`
3. Add router in `src/routers/`
4. Register router in `src/main.py`

### Database Changes

1. Modify models in `src/models/`
2. Create Alembic migration: `alembic revision --autogenerate -m "description"`
3. Apply migration: `alembic upgrade head`

### Adding New Search Features

1. Update OpenSearch mappings in `src/services/opensearch/service.py`
2. Modify search queries in search service methods
3. Update schemas to reflect new parameters/responses
4. Add endpoint in `src/routers/hybrid_search.py`

## Monitoring and Observability

- **Langfuse**: Access at http://localhost:3000 for trace visualization
- **Health checks**: `/health` endpoint provides service status
- **Logs**: `make logs` or `docker compose logs -f [service]`
- **Metrics**: Langfuse dashboard shows latency, costs, and performance

## Code Generation Guidelines

When generating or modifying code in this repository:

1. **Maintain weekly integrity**: Don't mix week-specific features across directories
2. **Follow the layered architecture**: Router → Service → Repository → Model
3. **Use dependency injection**: Services are injected via FastAPI dependencies
4. **Type everything**: Use Pydantic for validation, type hints for functions
5. **Handle errors gracefully**: Use custom exceptions in `src/exceptions.py`
6. **Follow async patterns**: Use `async/await` for I/O operations
7. **Document with docstrings**: Especially for complex service methods
8. **Environment-driven config**: Never hardcode configuration values

## Documentation Files

DON'T create summary document if its not asked by the user.

### Core Context Files

1. /docs/CURRENT_STATE.txt - Current project status, active services, and next steps
2. /docs/SESSION_HANDOFF.txt - Session summary and context for continuity

### Reference Files (Read as needed):

3. /docs/Journal.txt - Historical session records and project timeline
4. /docs/development-workflow.txt - Complete development workflow guide for all weeks
5. /docs/rag-architecture-overview.txt - Detailed RAG system architecture evolution across weeks
6. /docs/troubleshooting-guide.txt - Common issues and solutions for all services

### Security-Specific Files:

7. /docs/security/security-best-practices.txt - Security guidelines and best practices

## Session Management

Read /docs/CURRENT_STATE.txt and /docs/SESSION_HANDOFF.txt at the start of each new session to understand the current project state and previous work.

Update /docs/SESSION_HANDOFF.txt at the end of each session with:

- Accomplishments
- Current state
- Next steps
- Important context for next session

Update /docs/Journal.txt periodically with major milestones and decisions.
