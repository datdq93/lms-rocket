---
description: Pipeline đầy đủ - review code, auto-fix, generate commit, push lên GitHub. Shortcut cho /review + /commit.
---

# /ship - Ship Pipeline Workflow

Khi user gọi `/ship`, thực hiện pipeline đầy đủ: review → commit → push.

Đây là shortcut cho `/review` + `/commit`.

## Pipeline

### Phase 1: Review (tương đương /review)

1. Chạy `git diff --name-only` để xác định files thay đổi
2. Review từng file theo Code Review Checklist
3. Auto-fix các issues có thể fix tự động
4. Tạo review report

### Phase 2: Decision Gate

- Nếu có **blocking issues** → DỪNG LẠI, hiển thị issues, chờ user quyết định
- Nếu chỉ có warnings hoặc clean → tiếp tục Phase 3

### Phase 3: Commit & Push (tương đương /commit)

1. Kiểm tra branch (không push lên main/master)
2. Generate commit message từ diff
3. Hiển thị commit message cho user approve
4. Stage, commit, push
5. Cập nhật tasks.md

### Phase 4: Summary

```
🚀 SHIP COMPLETE
═════════════════

🔍 Review: ✅ Passed (X auto-fixes applied)
📝 Commit: abc1234 - feat(auth): add JWT refresh
🌿 Branch: feat/auth-refresh
📤 Push: ✅ Success

📋 Task Status:
  ✅ Task 3: Implement auth middleware → DONE

👉 Next:
  • Create PR on GitHub: https://github.com/user/lms-rocket/compare/feat/auth-refresh
  • Next task: Task 4 - Create user profile API
```

## Lưu ý

- Nếu review phát hiện blocking issues → pipeline DỪNG ở Phase 2
- User có thể fix manual rồi chạy `/ship` lại
- Hoặc user chạy `/review` riêng để xem chi tiết issues
- Pipeline này designed cho smooth flow khi code đã sẵn sàng
