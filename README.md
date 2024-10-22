# Notex

**Notex** is a minimalistic, markdown-based note-taking
app designed to boost productivity with features like
contextual intelligence,
version history, and live markdown preview.
Inspired by tools like Notion, Notex focuses on simplicity,
enabling users to take notes in markdown and organize them effectively.

## Features

- **Markdown Support**: Write and format notes using
  markdown syntax with real-time HTML preview.
- **Version Control**: Track changes to your notes with
  a version history feature, allowing users to compare and restore previous versions.

- **Contextual Intelligence**: Automatically detect and suggest
  actions based on the content of your notes
  (e.g., convert lists into tasks, recognize dates).

- **Interactive Blocks**: Add task lists, tables, and other
  block types directly in your notes for more dynamic content.

- **Dark Mode**: An optional dark mode for comfortable note-taking
  in low-light environments.
- **Collapsible Sidebar**: Organize your notes into sections and folders,
  easily accessible through a collapsible sidebar.

### Backend

- **Go** (Golang): API that handles user authentication,
  note CRUD operations, versioning, and markdown processing.
- **PostgreSQL**: Used for storing notes, users, and version history.
- **JWT Authentication**: To manage secure access to user data.

### Frontend

- **React**: For building the interactive user interface.
- **Tailwind CSS**: For rapid UI development with a focus on clean, modern design.
- **shadcn**: To provide a cohesive component library for consistent UI patterns.

### Prerequisites

- **Go** (version 1.23+)
- **Node.js** (version 16+)
- **PostgreSQL** (for database)

## License

This project is licensed under the GNU Affero General Public License v3.0.
See the [LICENSE](LICENSE) file for more information.
