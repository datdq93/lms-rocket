---
description: Generate commit message theo conventional commits, add files, commit và push lên GitHub.
---

# /commit - Commit & Push Workflow

Khi user gọi `/commit`, thực hiện các bước sau:

## Bước 1: Kiểm tra trạng thái

// turbo
- Chạy `git status` để xem trạng thái working directory
- Chạy `git diff --stat` để xem summary thay đổi
- Nếu không có thay đổi → thông báo "Nothing to commit"

## Bước 2: Kiểm tra branch

// turbo
- Chạy `git branch --show-current` để xem branch hiện tại
- **CẢNH BÁO** nếu đang ở `main` hoặc `master`
- Nếu đang ở main/master → hỏi user có muốn tạo branch mới không
- Suggest branch name theo format: `feat/task-description` hoặc `fix/task-description`

## Bước 3: Review diff trước khi commit

// turbo
- Chạy `git diff` để xem toàn bộ thay đổi
- Chạy `git diff --cached` để xem staged changes (nếu có)
- Phân tích nội dung thay đổi để generate commit message

## Bước 4: Generate Commit Message

Dựa trên diff, generate commit message theo Conventional Commits:

### Format:
```
<type>(<scope>): <subject>

<body>

<footer>
```

### Rules:
- **type**: feat, fix, refactor, docs, test, chore, style, perf
- **scope**: module/area bị ảnh hưởng
- **subject**: Mô tả ngắn gọn (max 72 chars), viết bằng tiếng Anh, lowercase, không dấu chấm cuối
- **body**: Chi tiết thay đổi (optional, khi thay đổi phức tạp)
- **footer**: Breaking changes, issue references (optional)

### Ví dụ:
```
feat(auth): add JWT token refresh mechanism

- Implement token refresh endpoint
- Add refresh token rotation
- Update auth middleware to handle expired tokens

Refs: #42
```

## Bước 5: Trình bày commit message cho user

Hiển thị commit message đã generate:
```
📝 PROPOSED COMMIT
══════════════════

feat(auth): add JWT token refresh mechanism

- Implement token refresh endpoint
- Add refresh token rotation  
- Update auth middleware to handle expired tokens

Refs: #42

📁 Files to commit:
  M src/auth/refresh.ts
  M src/middleware/auth.ts
  A src/auth/types.ts
  A tests/auth/refresh.test.ts
```

Hỏi user: "Proceed with this commit? (yes/edit/cancel)"

## Bước 6: Stage & Commit

Sau khi user approve:
- Chạy `git add -A` để stage tất cả changes
- Chạy `git commit -m "<commit_message>"` để commit
- Hiển thị kết quả commit

## Bước 7: Push lên GitHub

- Chạy `git push origin <current_branch>`
- Nếu branch chưa có upstream → `git push -u origin <current_branch>`
- Hiển thị kết quả push

## Bước 8: Cập nhật tasks.md

- Đánh dấu task liên quan là `done` (✅) trong tasks.md
- Thêm commit hash reference

## Bước 9: Summary

```
🚀 COMMIT & PUSH COMPLETE
══════════════════════════

📝 Commit: abc1234 - feat(auth): add JWT token refresh mechanism
🌿 Branch: feat/auth-refresh
📤 Pushed to: origin/feat/auth-refresh

👉 Next steps:
  • Create PR on GitHub
  • Run /gh-review after PR review
  • Or run /implement for next task
```

## Lưu ý

- **KHÔNG BAO GIỜ** force push
- **KHÔNG BAO GIỜ** commit trực tiếp lên main/master mà không hỏi user
- Mỗi commit chỉ nên chứa 1 task/feature
- Nếu có uncommitted changes từ nhiều tasks → suggest chia thành nhiều commits
- Luôn verify branch trước khi push
