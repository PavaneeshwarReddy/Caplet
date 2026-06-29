# Project Structure

This document describes the intended layout of the Caplet CLI project and what each folder is responsible for.

```text
caplet/
├── cmd/
│   └── caplet/
│       └── main.go                  # Application entry point
│
├── internal/
│   ├── cli/                         # Cobra command layer
│   │   ├── root.go                  # Root command setup
│   │   ├── doctor.go                # Health/check command
│   │   │
│   │   ├── skill/                   # Skill-related commands
│   │   │   ├── skill.go
│   │   │   ├── create.go
│   │   │   ├── list.go
│   │   │   ├── inspect.go
│   │   │   ├── validate.go
│   │   │   └── delete.go
│   │   │
│   │   ├── workspace/               # Workspace-related commands
│   │   │   ├── workspace.go
│   │   │   ├── create.go
│   │   │   ├── list.go
│   │   │   ├── use.go
│   │   │   └── delete.go
│   │   │
│   │   ├── vector/                  # Vector/index/search commands
│   │   │   ├── vector.go
│   │   │   ├── index.go
│   │   │   ├── search.go
│   │   │   └── rebuild.go
│   │   │
│   │   └── config/                  # Configuration commands
│   │       ├── config.go
│   │       ├── get.go
│   │       └── set.go
│   │
│   ├── skill/                       # Core skill domain logic
│   │   ├── service.go               # Skill use cases
│   │   ├── parser.go                # Parse skill definitions
│   │   ├── validator.go             # Validate skill structure
│   │   ├── loader.go                # Load skill files
│   │   └── repository.go            # Persistence for skills
│   │
│   ├── workspace/                   # Workspace management logic
│   │   ├── service.go               # Workspace use cases
│   │   ├── discovery.go             # Detect workspace context
│   │   └── repository.go            # Workspace storage access
│   │
│   ├── vector/                      # Vector search and indexing logic
│   │   ├── service.go               # Orchestrates vector features
│   │   ├── indexer.go               # Builds searchable indexes
│   │   ├── searcher.go              # Searches indexed content
│   │   └── embeddings.go            # Embedding generation
│   │
│   ├── agent/                       # Agent integration layer
│   │   ├── service.go               # Agent-related services
│   │   ├── protocol.go              # Communication protocol definitions
│   │   └── transport.go             # Transport implementation
│   │
│   ├── config/                      # Configuration handling
│   │   └── service.go               # Load/store app configuration
│   │
│   ├── storage/                     # File and cache storage abstractions
│   │   ├── filesystem.go            # File-system operations
│   │   └── cache.go                 # Cache implementation
│   │
│   ├── hooks/                       # Pre/post execution hooks
│   │   ├── pre.go                   # Pre-hooks
│   │   └── post.go                  # Post-hooks
│   │
│   └── util/                        # Shared helpers and utilities
│       ├── yaml.go                  # YAML utilities
│       ├── logger.go                # Logging helpers
│       └── path.go                  # Path helpers
│
├── docs/                            # Project documentation
├── examples/                        # Example usage and sample configs
├── testdata/                        # Test fixtures and sample data
├── go.mod                           # Go module definition
├── go.sum                           # Dependency checksums
└── README.md                        # Project overview and usage
```

## Folder responsibilities

- cmd/: Contains the runnable application entry points.
- internal/cli/: Holds Cobra commands and command-group organization.
- internal/skill/: Implements skill parsing, validation, loading, and storage.
- internal/workspace/: Manages workspace discovery and workspace-level operations.
- internal/vector/: Responsible for indexing and semantic/vector search capabilities.
- internal/agent/: Connects Caplet with agent-based workflows or protocols.
- internal/config/: Stores and loads project configuration.
- internal/storage/: Abstracts filesystem and caching behavior.
- internal/hooks/: Handles lifecycle hooks before and after operations.
- internal/util/: Provides reusable helper functions.
- docs/: Documentation files such as architecture and structure notes.
- examples/: Example projects or commands for users.
- testdata/: Files used in tests and local demonstrations.
