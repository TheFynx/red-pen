# Diagrams

Part of the red-pen skill. The core rules in `SKILL.md` apply here too.

Use diagrams to show relationships, flows, and states that are harder to follow as text. Do not use a diagram when a numbered list or table is clearer.

## When to use a diagram

- Architecture: components and how they connect.
- Flows: data moving through a pipeline, request/response chains.
- State machines: lifecycle stages and transitions.
- Sequences: multi-service interactions in time order.
- Entity relationships: data models.
- Decision trees: branching logic.

## When not to use a diagram

- Linear steps. A numbered list is clearer.
- Fewer than three nodes. Write a sentence.
- Comparisons. Use a table.

## Mermaid

Mermaid is the default for docs hosted on platforms that support it. Use fenced code blocks with the `mermaid` language identifier.

**Platform support:**

- GitHub: Mermaid, GeoJSON, TopoJSON, ASCII STL. GitHub pins an older Mermaid version. Not all features from the Mermaid live editor work on GitHub.
- GitLab: Mermaid, PlantUML.
- Docusaurus, MkDocs, Obsidian, Notion: Mermaid via plugins or built-in. Feature support varies by version.

**Know your target renderer before writing the diagram.** A diagram that renders cleanly in the Mermaid live editor can break on GitHub, Notion, or a static site generator. Test in the actual platform.

**Declaration order controls layout.** The layout engine (dagre by default) places nodes in the order you declare them. Top to bottom for TD graphs, left to right for LR graphs. If you declare `A --> C` before `A --> B`, C renders above B and edges cross.

Rules:

- Declare nodes in the visual order you want them rendered.
- Declare edges in the same order.
- When possible, declare all nodes before drawing edges. This establishes positions first.
- If edges cross, reorder your source declarations. Do not fight the layout engine with invisible nodes or spacing hacks.
- Keep diagrams small. If a diagram needs horizontal scrolling or has more than 15 nodes, split it into multiple diagrams with prose connecting them.
- Label nodes with short names (two to four words). Put detail in the prose around the diagram.
- Use LR (left to right) for pipelines and data flows. Use TD (top down) for hierarchies and trees.
- Do not use features your target renderer does not support. ELK layout, some newer diagram types, and certain directives do not work on GitHub.
- Always include a prose description near the diagram. Some renderers fail silently. Screen readers cannot parse diagrams. The prose is the fallback.

**Diagram type quick reference:**

| Type | Mermaid keyword | Use for |
|------|-----------------|---------|
| Flowchart | `graph` or `flowchart` | Architecture, data flow, decision trees |
| Sequence | `sequenceDiagram` | API calls, multi-service interactions |
| State | `stateDiagram-v2` | Lifecycle, status machines |
| Entity relationship | `erDiagram` | Data models |
| Gantt | `gantt` | Timelines (consider a table instead) |
| Class | `classDiagram` | Object relationships, interfaces |

