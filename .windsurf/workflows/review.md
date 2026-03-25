---
description: Review code tự động, phát hiện issues và auto-fix trước khi commit. Đảm bảo code quality.
---

# /review - Code Review Workflow

Khi user gọi `/review`, thực hiện các bước sau:

## Bước 1: Xác định scope review

// turbo
- Chạy `git diff --name-only` để lấy danh sách files đã thay đổi
- Chạy `git diff --stat` để xem tổng quan thay đổi
- Nếu không có thay đổi → thông báo "Nothing to review"

## Bước 2: Đọc và phân tích từng file

Với mỗi file đã thay đổi:
// turbo
- Chạy `git diff <file>` để xem chi tiết thay đổi
- Đọc full file để hiểu context
- Phân tích theo Code Review Checklist trong AGENTS.md

## Bước 3: Security Review

Kiểm tra:
- [ ] Không có hardcoded secrets, API keys, passwords
- [ ] Không có SQL injection vulnerabilities
- [ ] Không có XSS vulnerabilities
- [ ] Input validation đầy đủ
- [ ] Không expose sensitive data trong logs
- [ ] Đúng authentication/authorization checks
- [ ] Không có path traversal issues

## Bước 4: Code Quality Review

Kiểm tra:
- [ ] Type annotations đầy đủ (TypeScript)
- [ ] Error handling cho async operations
- [ ] Không có N+1 query problems
- [ ] Không có memory leaks (event listeners, subscriptions)
- [ ] DRY - không duplicate code
- [ ] Naming conventions đúng (camelCase, PascalCase, kebab-case)
- [ ] Imports organized và không unused imports
- [ ] Không có console.log hoặc debug code còn sót
- [ ] Không có TODO/FIXME mà không có task tương ứng

## Bước 5: Test Review

Kiểm tra:
- [ ] Tests cover business logic chính
- [ ] Tests cover edge cases
- [ ] Tests cover error scenarios
- [ ] Test names mô tả rõ ràng behavior
- [ ] Không có test nào bị skip mà không có lý do

## Bước 6: Performance Review

Kiểm tra:
- [ ] Không có unnecessary re-renders (React)
- [ ] Proper memoization khi cần
- [ ] Lazy loading cho heavy components
- [ ] Efficient database queries
- [ ] Proper caching strategy

## Bước 7: Auto-Fix

Nếu phát hiện issues, **tự động fix** những vấn đề sau:
- Unused imports → Remove
- Missing type annotations → Add types
- Console.log statements → Remove
- Formatting issues → Fix formatting
- Missing error handling → Add try/catch
- Missing null checks → Add optional chaining
- Inconsistent naming → Fix to match conventions

Với mỗi auto-fix:
- Thực hiện sửa code
- Log lại đã fix gì

## Bước 8: Báo cáo Review

Hiển thị report cho user:

```
🔍 CODE REVIEW REPORT
══════════════════════

📁 Files reviewed: X
📊 Changes: +XX / -XX lines

✅ PASSED
  • Security: No issues found
  • Types: All typed correctly

⚠️ WARNINGS (non-blocking)
  • [file:line] Consider extracting this to a utility function
  • [file:line] This function is getting complex (cyclomatic complexity > 10)

🔧 AUTO-FIXED
  • [file:line] Removed unused import 'xyz'
  • [file:line] Added missing error handling
  • [file:line] Removed console.log

❌ ISSUES (blocking) - Cần user quyết định
  • [file:line] [Mô tả issue cần human decision]

📋 SUMMARY
  • Auto-fixed: X issues
  • Warnings: X (non-blocking)
  • Blocking: X (cần user review)
```

## Bước 9: Next Steps

- Nếu không có blocking issues → Suggest chạy `/commit`
- Nếu có blocking issues → Liệt kê và chờ user quyết định
- Nếu có warnings → Thông báo nhưng không block

## Lưu ý

- **LUÔN** auto-fix được gì thì fix, giảm thiểu manual work cho user
- **KHÔNG** thay đổi logic business mà không hỏi user
- **KHÔNG** thêm/xóa features ngoài scope
- Review phải khách quan, dựa trên conventions đã define
- Nếu phát hiện architectural issue → flag nhưng không block
