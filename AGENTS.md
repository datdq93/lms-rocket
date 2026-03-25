# LMS Rocket - AI Development Agents Guide

## 🚀 Project Overview

**LMS Rocket** là hệ thống Learning Management System được phát triển theo hướng **AI-first**.

## 📋 Development Philosophy

### AI-First Principles
1. **AI là developer chính** - Mọi code đều được AI generate, human review và approve
2. **Workflow-driven** - Mọi thao tác đều đi qua workflow chuẩn hóa
3. **Task-based** - Mọi thay đổi đều gắn với task cụ thể trong `tasks.md`
4. **Auto-review** - Code luôn được review tự động trước khi commit
5. **Iterative** - Phát triển theo vòng lặp nhỏ: plan → implement → review → ship

## 🔄 Development Pipeline

```
User Request → /plan → /implement → /review → /commit → GitHub PR → /gh-review → Done
```

### Workflow Commands
| Command | Mô tả |
|---------|--------|
| `/plan` | Phân tích yêu cầu → tạo plan → chia task |
| `/task` | Xem/quản lý danh sách task |
| `/implement` | Implement task hiện tại hoặc task được chỉ định |
| `/review` | Review code, auto-fix issues trước khi commit |
| `/commit` | Generate commit message → add → commit → push |
| `/ship` | Pipeline đầy đủ: review → commit → push |
| `/gh-review` | Đọc GitHub PR comments → fix → push |

## 📁 Project Structure Convention

```
lms-rocket/
├── AGENTS.md                    # AI development guide (this file)
├── tasks.md                     # Task tracking file
├── .windsurf/
│   ├── workflows/               # Workflow definitions
│   │   ├── plan.md
│   │   ├── implement.md
│   │   ├── review.md
│   │   ├── commit.md
│   │   ├── ship.md
│   │   ├── gh-review.md
│   │   └── task.md
│   └── skills/                  # Reusable skills
│       ├── code-review.md
│       └── task-analysis.md
├── src/                         # Source code
├── tests/                       # Test files
├── docs/                        # Documentation
└── ...
```

## 🎯 Coding Standards

### General Rules
- **Language**: TypeScript preferred, JavaScript acceptable
- **Style**: Follow existing codebase conventions
- **Comments**: Meaningful comments only, no obvious ones
- **Naming**: camelCase for variables/functions, PascalCase for classes/components
- **Files**: kebab-case for file names

### AI Code Generation Rules
1. **Không hardcode** - Dùng environment variables cho config
2. **Type-safe** - Luôn define types/interfaces
3. **Error handling** - Mọi async operation phải có error handling
4. **Testing** - Viết test cho mọi business logic
5. **Small commits** - Mỗi commit chỉ làm 1 việc rõ ràng
6. **Conventional commits** - Format: `type(scope): description`

### Commit Message Convention
```
feat(module): add new feature
fix(module): fix bug description
refactor(module): refactor description
docs(module): update documentation
test(module): add/update tests
chore(module): maintenance task
style(module): formatting changes
perf(module): performance improvement
```

## 🔍 Code Review Checklist

Khi review code (workflow `/review`), kiểm tra:

1. **Security** - Không expose secrets, SQL injection, XSS
2. **Performance** - Không N+1 queries, memory leaks
3. **Types** - Đầy đủ type annotations
4. **Error Handling** - Try/catch, error boundaries
5. **Tests** - Coverage đủ cho business logic
6. **DRY** - Không duplicate code
7. **SOLID** - Follow SOLID principles
8. **Accessibility** - a11y cho UI components

## 📝 Task Status Definitions

| Status | Emoji | Meaning |
|--------|-------|---------|
| `todo` | ⬜ | Chưa bắt đầu |
| `in-progress` | 🔵 | Đang làm |
| `review` | 🟡 | Đang review |
| `done` | ✅ | Hoàn thành |
| `blocked` | 🔴 | Bị block |

## 🛡️ Safety Rules

1. **KHÔNG BAO GIỜ** tự động xóa files/directories mà không hỏi user
2. **KHÔNG BAO GIỜ** force push lên branch `main` hoặc `master`
3. **LUÔN** tạo branch mới cho mỗi feature/fix
4. **LUÔN** review diff trước khi commit
5. **LUÔN** chạy tests trước khi push
