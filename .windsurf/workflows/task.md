---
description: Xem và quản lý danh sách task trong tasks.md. Hiển thị trạng thái, cập nhật status task.
---

# /task - Task Management Workflow

Khi user gọi `/task`, thực hiện các bước sau:

## Bước 1: Đọc tasks.md

// turbo
- Đọc file `tasks.md` trong root project
- Nếu file không tồn tại, thông báo cho user và gợi ý chạy `/plan` trước

## Bước 2: Hiển thị Task Board

Hiển thị danh sách task dưới dạng board:

```
📋 TASK BOARD - LMS Rocket
═══════════════════════════

🔵 IN PROGRESS (1)
  → Task 3: Implement auth middleware [M]

⬜ TODO (3)
  • Task 4: Create user profile API [M]
  • Task 5: Add validation layer [S]
  • Task 6: Write integration tests [L]

✅ DONE (2)
  ✓ Task 1: Setup project structure [S]
  ✓ Task 2: Configure database [M]

🔴 BLOCKED (0)
  (none)
```

## Bước 3: Xử lý lệnh phụ (nếu có)

User có thể gọi với tham số:

### `/task` (không tham số)
- Hiển thị task board như trên

### `/task next`
- Tìm task tiếp theo cần làm (task todo đầu tiên)
- Suggest user chạy `/implement` cho task đó

### `/task done <task_number>`
- Đánh dấu task đã hoàn thành (✅)
- Cập nhật tasks.md

### `/task status <task_number> <status>`
- Cập nhật status của task
- Status: `todo`, `in-progress`, `review`, `done`, `blocked`

### `/task add <description>`
- Thêm task mới vào cuối danh sách
- Status mặc định: `todo`

### `/task remove <task_number>`
- Hỏi user xác nhận trước khi xóa
- Xóa task khỏi danh sách

## Bước 4: Cập nhật tasks.md

- Sau mỗi thao tác thay đổi, cập nhật file tasks.md
- Giữ nguyên format và structure
- Cập nhật timestamp last modified

## Lưu ý

- Luôn hiển thị summary sau mỗi thao tác
- Gợi ý next action cho user (implement, review, etc.)
- Nếu tất cả task done → suggest user chạy `/ship`
