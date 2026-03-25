---
name: task-analysis
description: A brief description, shown to the model to help it understand when to use this skill
---

---
description: Skill hướng dẫn cách phân tích yêu cầu và chia task hiệu quả cho AI-first development
---

# Task Analysis Skill

## Mục đích
Cung cấp framework để phân tích yêu cầu từ user và chia thành tasks nhỏ, phù hợp cho AI implementation.

## Quy trình phân tích

### 1. Requirement Decomposition

Từ yêu cầu user, extract:
- **What**: Cần làm gì? (functional requirements)
- **Why**: Tại sao cần làm? (business context)
- **Who**: Ai sử dụng? (target users)
- **Where**: Ảnh hưởng ở đâu? (affected modules)
- **When**: Deadline? Priority? (urgency)
- **How**: Constraints kỹ thuật? (technical requirements)

### 2. Scope Assessment

Phân loại scope:
- **XS** (Extra Small): < 50 lines, 1 file → 1 task
- **S** (Small): < 100 lines, 1-2 files → 1-2 tasks
- **M** (Medium): < 300 lines, 3-5 files → 3-5 tasks
- **L** (Large): < 500 lines, 5-10 files → 5-8 tasks
- **XL** (Extra Large): > 500 lines → Chia thành multiple PRs

### 3. Task Splitting Rules

#### Rule 1: Single Responsibility
Mỗi task chỉ làm 1 việc:
- ✅ "Add User model and migration"
- ❌ "Add User model, migration, API, and tests"

#### Rule 2: Vertical Slicing
Chia theo feature slice, không phải layer:
- ✅ "Add login endpoint (route + controller + service + test)"
- ❌ "Add all routes" → quá horizontal

#### Rule 3: Dependency Ordering
Tasks phải có thứ tự dependencies rõ ràng:
```
Task 1: Setup database schema     ← No dependency
Task 2: Create data models         ← Depends on Task 1
Task 3: Implement business logic   ← Depends on Task 2
Task 4: Add API endpoints          ← Depends on Task 3
Task 5: Write integration tests    ← Depends on Task 4
```

#### Rule 4: Testable Output
Mỗi task phải có cách verify kết quả:
- Unit test passes
- API endpoint responds correctly
- UI renders properly
- Build succeeds

#### Rule 5: AI-Friendly Size
Mỗi task phải:
- Không quá 200 dòng code thay đổi
- Không thay đổi quá 5 files
- Có context rõ ràng (không cần quá nhiều codebase knowledge)
- Có acceptance criteria cụ thể

### 4. Task Template

```markdown
- ⬜ **Task N**: [Tên ngắn gọn] `[S/M/L]`
  - Description: [Mô tả chi tiết hơn]
  - Acceptance: [Criteria cụ thể, measurable]
  - Files: [Danh sách files cần tạo/sửa]
  - Dependencies: [Task numbers phụ thuộc]
  - Notes: [Lưu ý kỹ thuật nếu có]
```

### 5. Common Task Patterns

#### New Feature Pattern
1. Define types/interfaces
2. Create database schema/migration
3. Implement core business logic + unit tests
4. Create API endpoints + integration tests
5. Build UI components + component tests
6. Integration & E2E testing
7. Documentation update

#### Bug Fix Pattern
1. Write failing test reproducing the bug
2. Implement fix
3. Verify fix passes test
4. Add regression test

#### Refactor Pattern
1. Write characterization tests (nếu chưa có)
2. Implement refactor (small steps)
3. Verify all tests pass
4. Update documentation

#### Performance Improvement Pattern
1. Add performance benchmark/measurement
2. Implement optimization
3. Verify improvement with benchmark
4. Add performance regression test

### 6. Estimation Guide

| Size | Lines of Code | Files | Time Estimate | Complexity |
|------|--------------|-------|---------------|------------|
| S | < 50 | 1-2 | 5-10 min | Simple, straightforward |
| M | 50-150 | 2-4 | 10-20 min | Some logic, multiple files |
| L | 150-300 | 4-7 | 20-40 min | Complex logic, many files |

### 7. Quality Checklist cho Plan

Trước khi present plan cho user, verify:
- [ ] Mỗi task có acceptance criteria rõ ràng?
- [ ] Dependencies được xác định đúng?
- [ ] Không có task nào quá lớn (> 200 LOC)?
- [ ] Task đầu tiên là setup/scaffolding?
- [ ] Có task cho testing?
- [ ] Total scope phù hợp với yêu cầu (không over-engineer)?
- [ ] Backward compatibility được xem xét?
