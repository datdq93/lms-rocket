---
description: Phân tích yêu cầu từ user, tạo plan chi tiết và chia thành các task nhỏ trong tasks.md
---

# /plan - Planning Workflow

Khi user gọi `/plan`, thực hiện các bước sau:

## Bước 1: Thu thập yêu cầu

- Đọc yêu cầu từ user message
- Nếu yêu cầu chưa rõ ràng, hỏi lại user để làm rõ trước khi tiếp tục
- Xác định scope: feature mới, bug fix, refactor, hay improvement

## Bước 2: Phân tích codebase hiện tại

- Đọc file `AGENTS.md` để hiểu conventions của project
- Đọc file `tasks.md` để biết trạng thái hiện tại của project
- Sử dụng `code_search` để tìm hiểu codebase liên quan đến yêu cầu
- Xác định các files/modules bị ảnh hưởng
- Xác định dependencies và potential conflicts

## Bước 3: Tạo Plan

Tạo plan chi tiết bao gồm:
- **Objective**: Mục tiêu rõ ràng
- **Scope**: Phạm vi thay đổi
- **Approach**: Phương pháp tiếp cận (giải thích WHY chọn approach này)
- **Affected Files**: Danh sách files cần thay đổi
- **Risks**: Rủi ro và cách giảm thiểu
- **Dependencies**: Thư viện/module cần thêm (nếu có)

## Bước 4: Chia Task

Chia plan thành các task nhỏ, mỗi task phải:
- Có thể hoàn thành trong 1 lần implement (không quá lớn)
- Độc lập hoặc có thứ tự rõ ràng
- Có acceptance criteria cụ thể
- Có estimate effort (S/M/L)

## Bước 5: Cập nhật tasks.md

Cập nhật file `tasks.md` theo format sau:

```markdown
## [Feature/Fix Name] - [Date]

### Objective
[Mô tả ngắn gọn mục tiêu]

### Tasks
- ⬜ **Task 1**: [Mô tả] `[S/M/L]`
  - Acceptance: [Criteria]
  - Files: [file1, file2]
- ⬜ **Task 2**: [Mô tả] `[S/M/L]`
  - Acceptance: [Criteria]
  - Files: [file1, file2]
```

## Bước 6: Trình bày cho User

- Hiển thị plan tổng quan cho user
- Liệt kê danh sách tasks đã chia
- Hỏi user confirm hoặc điều chỉnh trước khi bắt đầu implement
- **KHÔNG tự động implement** - chờ user approve plan

## Lưu ý quan trọng

- Mỗi task KHÔNG nên vượt quá 200 dòng code thay đổi
- Task đầu tiên nên là setup/scaffolding nếu là feature mới
- Task cuối cùng nên là integration test hoặc cleanup
- Nếu feature phức tạp, chia thành nhiều PR nhỏ
- Luôn xem xét backward compatibility
