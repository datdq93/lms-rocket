---
description: Implement task hiện tại hoặc task được chỉ định. AI sẽ tự động code theo plan đã approve.
---

# /implement - Implementation Workflow

Khi user gọi `/implement` hoặc `/implement <task_number>`, thực hiện các bước sau:

## Bước 1: Xác định task cần implement

// turbo
- Đọc file `tasks.md`
- Nếu user chỉ định task number → lấy task đó
- Nếu không → lấy task `in-progress` đầu tiên
- Nếu không có task `in-progress` → lấy task `todo` đầu tiên
- Nếu không có task nào → thông báo và suggest chạy `/plan`

## Bước 2: Phân tích task

- Đọc task description và acceptance criteria
- Đọc danh sách files bị ảnh hưởng
- Sử dụng `code_search` để hiểu context của code liên quan
- Đọc các files liên quan để hiểu structure hiện tại
- Xác định approach cụ thể cho implementation

## Bước 3: Cập nhật task status

- Đánh dấu task là `in-progress` (🔵) trong tasks.md
- Thông báo cho user task nào đang được implement

## Bước 4: Implementation

Thực hiện implementation theo thứ tự:

### 4.1 - Tạo/sửa types và interfaces (nếu cần)
- Define types trước khi implement logic
- Đảm bảo type-safe

### 4.2 - Implement core logic
- Viết code theo conventions trong AGENTS.md
- Follow existing patterns trong codebase
- Thêm error handling cho mọi async operation
- KHÔNG hardcode values - dùng constants hoặc env vars

### 4.3 - Viết tests
- Unit test cho business logic
- Integration test cho API endpoints (nếu có)
- Test edge cases và error scenarios

### 4.4 - Cập nhật imports và dependencies
- Thêm imports cần thiết
- Cập nhật package.json nếu cần thêm dependency mới

## Bước 5: Verify implementation

- Kiểm tra code đã implement đáp ứng acceptance criteria
- Đảm bảo không break existing functionality
- Chạy linter check nếu có config
// turbo
- Chạy tests liên quan: `npm test -- --related` hoặc tương đương

## Bước 6: Cập nhật task status

- Đánh dấu task là `review` (🟡) trong tasks.md
- Hiển thị summary những gì đã implement
- Liệt kê files đã thay đổi
- Suggest user chạy `/review` trước khi commit

## Bước 7: Trình bày kết quả

Hiển thị cho user:
```
✅ Implementation Complete - Task #X
══════════════════════════════════

📝 Changes:
  • [file1] - [mô tả thay đổi]
  • [file2] - [mô tả thay đổi]

🧪 Tests:
  • [test results summary]

📊 Stats:
  • Files changed: X
  • Lines added: +XX
  • Lines removed: -XX

👉 Next: Run /review to check code quality before committing
```

## Lưu ý quan trọng

- **KHÔNG** implement nhiều task cùng lúc
- **KHÔNG** thay đổi code ngoài scope của task
- Nếu phát hiện bug hoặc issue khác → tạo task mới, KHÔNG fix inline
- Nếu task quá lớn → suggest user chia nhỏ hơn
- Nếu bị block bởi task khác → đánh dấu `blocked` và thông báo user
- Prefer sử dụng edit/multi_edit thay vì write_to_file cho existing files
