# 📋 LMS Rocket - Task Tracking

> File này được quản lý bởi AI workflows. Sử dụng `/task` để xem và quản lý tasks.
> 
> Last updated: 2026-03-25

---

## Legend
| Status | Emoji | Meaning |
|--------|-------|---------|
| `todo` | ⬜ | Chưa bắt đầu |
| `in-progress` | 🔵 | Đang làm |
| `review` | 🟡 | Đang review |
| `done` | ✅ | Hoàn thành |
| `blocked` | 🔴 | Bị block |

---

## Project Setup - 2026-03-25

### Objective
Thiết lập AI-first development workflow cho LMS Rocket project.

### Tasks
- ✅ **Task 1**: Setup AGENTS.md và workflow files `[M]`
  - Acceptance: Tất cả workflow files tồn tại và có nội dung đầy đủ
  - Files: AGENTS.md, .windsurf/workflows/*, .windsurf/skills/*

---

<!-- New features/tasks will be added below this line -->

## Complete Specification - 2026-03-25

### Objective
Hoàn thiện Specification.md với database schema chi tiết, API design đầy đủ, và system architecture.

### Tasks
- ✅ **Task 2**: Hoàn thiện Database Schema `[M]`
  - Acceptance: Có ERD diagram, data types, indexes, constraints, soft delete, timestamps
  - Files: Specification.md (section 5)

- ✅ **Task 3**: Chi tiết hóa API Design `[L]`
  - Acceptance: Tất cả endpoints với request/response schemas, auth per endpoint, error formats, pagination
  - Files: Specification.md (section 6)

- ✅ **Task 4**: Bổ sung System Architecture `[M]`
  - Acceptance: Folder structure Backend + Frontend, layer diagram, data flow
  - Files: Specification.md (section 2)

- ✅ **Task 5**: Bổ sung Security & Auth Strategy `[M]`
  - Acceptance: JWT strategy, validation rules, rate limiting, encryption approach
  - Files: Specification.md (new section)

---

## Backend Implementation Phase 1 - 2026-03-25

### Objective
Setup Golang backend project với Clean Architecture, database connection, và base structure.

### Tasks
- ✅ **Task 6**: Backend project scaffolding `[M]`
  - Acceptance: Go project structure theo spec, go.mod, main.go, Dockerfile
  - Files: lms-backend/cmd/api/main.go, go.mod, Dockerfile, docker-compose.yml, .env.example

- ⬜ **Task 7**: Database configuration & models `[M]`
  - Acceptance: GORM setup, connection, auto-migrate, domain models
  - Files: lms-backend/internal/config/database.go, internal/domain/*.go

- ⬜ **Task 8**: Authentication service `[L]`
  - Acceptance: Register, login, JWT generation, password hashing
  - Files: lms-backend/internal/service/auth_service.go, handler/auth_handler.go, middleware/auth.go

- ⬜ **Task 9**: User APIs `[M]`
  - Acceptance: Get profile, update profile endpoints
  - Files: lms-backend/internal/handler/user_handler.go, service/user_service.go
