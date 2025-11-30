# Crowdsourced Life Maps 🌳

> **Git for Life Goals.**
> An open-source platform for creating, visualizing, and tracking complex life scenarios (relocation, education, health) as branching dependency graphs.

![Status](https://img.shields.io/badge/status-Work_in_Progress-yellow)
![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16-4169E1?logo=postgresql&logoColor=white)
![React](https://img.shields.io/badge/React-18-61DAFB?logo=react&logoColor=black)
![TypeScript](https://img.shields.io/badge/TypeScript-5-3178C6?logo=typescript&logoColor=white)

## 💡 Concept

People often get lost in life not because they lack skills, but because they lack a clear algorithm of actions. Articles get outdated, and simple checklists are too linear for complex reality.

**Life Maps** introduces an **"Infrastructure as Code"** approach to real life:

*   **Decision Trees:** Life isn't linear. Graph nodes have dependencies.
*   **Forks:** Copy someone else's path (e.g., "Relocation to Germany") and adapt it to your specific situation.
*   **Merge Requests:** Found a mistake in the document list? Propose a fix to the original guide.
*   **Analytics:** We track where people "drop off" to identify the hardest steps in any goal.

## 🛠 Tech Stack

### 🔙 Backend (API)
Built with a focus on high performance and complex hierarchical data structures.

*   **Language:** Go (Golang)
*   **Web Framework:** [Gin](https://github.com/gin-gonic/gin) — for high-performance REST API.
*   **Database:** PostgreSQL 16+
    *   **`ltree` extension:** For efficient storage and retrieval of hierarchical data (Materialized Path pattern).
    *   **`jsonb`:** For flexible node properties (documents, pricing, links, meta-data).
*   **Driver:** [pgx/v5](https://github.com/jackc/pgx) — low-level, high-performance driver.
*   **Architecture:** Clean Architecture (Core -> Repository -> Service -> Transport).

### 🎨 Frontend (Client)
A Single Page Application (SPA) focused on rich interactivity and visualization.

*   **Build Tool:** [Vite](https://vitejs.dev/) — for lightning-fast bundling.
*   **Language:** TypeScript.
*   **Graph Engine:** [React Flow](https://reactflow.dev/) — for rendering interactive nodes and handling drag-and-drop operations.
*   **State Management:** [TanStack Query (React Query)](https://tanstack.com/query/latest).
*   **UI Framework:** [Tailwind CSS](https://tailwindcss.com/) + [shadcn/ui](https://ui.shadcn.com/).

## 🚀 Key Features (MVP)

- [ ] **Auth:** Registration and JWT-based authentication.
- [ ] **Map Editor:** Visual graph editor for creating life scenarios.
- [ ] **Graph Logic:** Adding nodes, moving branches (automatic `ltree` path recalculation).
- [ ] **Forking:** Cloning a user's tree into a personal workspace.
- [ ] **Execution:** Tracking progress and marking steps as "Done".

## 💾 Data Model (Engineering)

The core of the project relies on the PostgreSQL `ltree` extension. This allows us to fetch entire branches and find descendants in a single SQL query without recursion, ensuring O(1) read complexity for path lookups.

**Example Path:** `relocation.usa.docs.visa`

```sql
-- Fetching the whole subtree efficiently
SELECT * FROM nodes WHERE path <@ 'relocation.usa';
```

## 📦 Getting Started

### Prerequisites
*   Go 1.22+
*   Node.js 18+ & npm/pnpm
*   Docker & Docker Compose

### 1. Backend Setup

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/xLightly/life-maps.git
    cd life-maps
    ```

2.  **Configure environment:**
    Create a `.env` file based on the example:
    ```bash
    cp .env.example .env
    ```

3.  **Start the Database:**
    ```bash
    docker-compose up -d db
    ```

4.  **Run Migrations:**
    ```bash
    go run cmd/migrator/main.go up
    ```

5.  **Run the Server:**
    ```bash
    go run cmd/app/main.go
    ```

### 2. Frontend Setup

1.  **Navigate to the frontend directory:**
    ```bash
    cd frontend
    ```

2.  **Install dependencies:**
    ```bash
    npm install
    ```

3.  **Configure Environment:**
    Create a `.env.local` file:
    ```env
    VITE_API_URL=http://localhost:8080/api/v1
    ```

4.  **Start the Development Server:**
    ```bash
    npm run dev
    ```

## 🗺 Roadmap

*   **v0.1:** Database schema implementation and basic Map CRUD API.
*   **v0.2:** Visual Graph Editor (Frontend) and Node manipulation logic.
*   **v0.3:** Forking logic and Merge Request system.
*   **v0.4:** **Link Warden Microservice** (Background worker to check for broken links in guides).
*   **v0.5:** "Bottleneck" analytics for user progress.

## 🤝 Contributing

The project is in the early stages of development. Ideas, issues, and Pull Requests are welcome!

1.  Fork it
2.  Create your feature branch (`git checkout -b feature/AmazingFeature`)
3.  Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4.  Push to the branch (`git push origin feature/AmazingFeature`)
5.  Open a Pull Request

## 📄 License

Distributed under the MIT License. See `LICENSE` for more information.
