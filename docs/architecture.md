# Architecture

This section describes the individual data elements used within the repository, the overall data flow, and the design decisions that shape the architecture.

## Caplet.yaml
This is the file structure that CLI tool use to store the skills. The default format is .yaml will be extended to multiple formats.
### Data Elements

Each skill is defined in a `SKILL.md` file. This file is consumed by AI agents (such as Claude, ChatGPT, Copilot, Kimi, etc.) to understand and execute the skill. How these files are discovered and supplied to agents is explained in the **Data Flow** section.

#### Headers

The header fields are always supplied to the agent to help it identify and select the most appropriate skill.

* `name`: A unique identifier for the skill within the workspace.
* `description`: A concise, high-level summary of the skill. Keep this short, as it is primarily used during skill selection.

#### Body

The body is supplied only after the skill becomes active. It contains the information and resources required by the agent to execute the skill.

* `shared`: Contains resources that can be reused across one or more operations in the execution chain.

  * `scripts`: Executable scripts that the agent can invoke as part of an operation.
  * `reference`: Reference materials such as documentation, PDFs, Markdown files, or design documents that operations can consult when additional context is needed.
  * `assets`: Supporting files or resources required to complete an operation, such as templates, configuration files, datasets, or sample inputs.

* `chain`: A chain is an ordered sequence of operations that an agent performs to accomplish a specific task.

  * `operations`: Represents an implementation task that the agent is expected to perform.

  * `name`: A unique name for the operation within the current skill.

  * `description`: A brief explanation of the operation's objective.

  * `depends_on`: A list of operations that must complete before this operation can begin.

  * `inputs`: The resources or artifacts required to execute the operation. These may include outputs from previous operations, scripts, reference documents, assets, or user-provided inputs.

  * `outputs`: The artifacts produced by the operation that can be consumed by subsequent operations.

  * `tools`: The tools the agent is allowed to use while performing the operation (for example: `git`, `npm`, `docker`, `kubectl`).

  * `permissions`: The permissions required by the operation. If additional permissions are needed beyond those granted at the skill level, the agent should request user approval.

  * `resources`: References to shared resources (scripts, documentation, assets, templates, etc.) that are relevant to this operation.

  * `guidelines`: A set of instructions, coding standards, architectural principles, or best practices that the agent should follow while executing the operation.

  * `constraints`: Rules describing what the agent must not do (for example, "Do not modify public APIs" or "Do not delete existing files").

  * `validation`: Criteria used to verify that the operation completed successfully. This may include running tests, linting, or checking for expected output files.

  * `retry`: Specifies whether the agent may retry the operation and, if so, how many attempts are allowed.

  * `examples`: Example inputs, outputs, or workflows that demonstrate how the operation should be executed.

  * `failure_handling`: Instructions describing how the agent should respond if the operation fails, including cleanup or recovery steps.

  * `notes`: Additional information that may help the agent but is not required for execution.


#### Hooks

Hooks define actions or skills that should be executed before or after the main execution chain.

* `pre`: Actions or skills that must be executed before the main chain begins.
* `post`: Actions or skills that should be executed after the main chain completes, typically for cleanup, validation, or follow-up tasks.


### Data Flow

* The complete Caplet manual is provided to the agent at the beginning of the interaction.

* An agent can request the list of skills available within a workspace. Only the skill headers (such as `name` and `description`) are returned to help the agent identify relevant skills.

* Once the agent selects one or more skills, it can request their full definitions. Caplet then returns the complete contents of the requested `SKILL.md` files.

* As an optimization, the agent may provide the user's query when requesting available skills. Caplet can use vector-based similarity search to return only the most relevant skills instead of the entire skill catalog. This behavior is optional and falls back to returning all skill headers if semantic search is unavailable.

* An agent may request only specific sections or fields of a skill (for example, `shared.references`, `chain.operations`, or `hooks`) instead of loading the entire `SKILL.md`. This minimizes context usage and allows the agent to retrieve only the information required for the current task.


### Design Principles

The following design principles guide the architecture and implementation of Caplet. They are intended to keep the system modular, extensible, efficient, and easy to integrate with both current and future AI agents.

* **Agent Agnostic** – Skills are designed to be consumable by any compatible AI agent without being tied to a specific provider or model.

* **Modularity** – Each skill represents a single capability and should remain self-contained and independently maintainable.

* **Composability** – Skills can build upon other skills, enabling complex workflows to be assembled from reusable components.

* **Lazy Loading** – Only the information required by the agent is loaded, minimizing context usage and improving efficiency.

* **Agent Autonomy** – Skills define objectives and constraints, while the agent determines the most appropriate execution strategy.

* **Reusability** – Shared resources such as scripts, references, assets, and skills should be reusable across multiple skills whenever possible.

* **Extensibility** – The skill specification is designed to evolve over time without breaking existing skills or integrations.

* **Context Efficiency** – Agents should be able to retrieve only the specific skills or sections they require, reducing unnecessary context consumption.

