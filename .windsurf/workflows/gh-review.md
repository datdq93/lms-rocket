---
description: Đọc GitHub PR review comments, phân tích feedback, tự động fix và push lại. Dùng sau khi reviewer comment trên GitHub.
---

# /gh-review - GitHub Review Response Workflow

Khi user gọi `/gh-review` hoặc `/gh-review <pr_url>`, thực hiện các bước sau:

## Bước 1: Lấy PR Information

- Nếu user cung cấp PR URL → extract owner/repo/pr_number
- Nếu không → hỏi user PR URL hoặc PR number
// turbo
- Chạy `gh pr view <pr_number> --json title,body,state,reviewDecision,reviews,comments` để lấy PR info
- Nếu `gh` CLI chưa cài → hướng dẫn user cài: `brew install gh && gh auth login`

## Bước 2: Lấy Review Comments

// turbo
- Chạy `gh pr view <pr_number> --json reviews --jq '.reviews[] | {author: .author.login, state: .state, body: .body}'`
// turbo
- Chạy `gh api repos/{owner}/{repo}/pulls/<pr_number>/comments --jq '.[] | {path: .path, line: .line, body: .body, author: .user.login}'`
- Phân loại comments:
  - **Action Required**: Comments yêu cầu thay đổi code
  - **Questions**: Comments hỏi về implementation
  - **Suggestions**: Gợi ý cải thiện
  - **Approved**: Comments approve

## Bước 3: Hiển thị Review Summary

```
📋 PR REVIEW SUMMARY - PR #42
══════════════════════════════

🔴 CHANGES REQUESTED by @reviewer1

📝 Comments (5):

1. 🔧 [ACTION] src/auth/refresh.ts:25
   "Should validate refresh token expiry before rotation"
   
2. ❓ [QUESTION] src/auth/types.ts:10
   "Why not use a union type here?"

3. 💡 [SUGGESTION] src/middleware/auth.ts:45
   "Consider using a guard clause to reduce nesting"

4. 🔧 [ACTION] tests/auth/refresh.test.ts:30
   "Missing test for expired token scenario"

5. ✅ [APPROVED] General
   "LGTM on the overall architecture"
```

## Bước 4: Xử lý từng comment

Với mỗi **Action Required** comment:
1. Đọc file và dòng code liên quan
2. Hiểu context và yêu cầu của reviewer
3. Implement fix theo reviewer feedback
4. Log lại đã fix gì

Với mỗi **Question** comment:
1. Phân tích code liên quan
2. Chuẩn bị câu trả lời
3. Nếu question hợp lý và cần thay đổi code → implement
4. Hiển thị suggested reply cho user

Với mỗi **Suggestion** comment:
1. Đánh giá suggestion có hợp lý không
2. Nếu hợp lý → implement
3. Nếu không → giải thích lý do cho user quyết định

## Bước 5: Implement Fixes

- Thực hiện tất cả fixes đã xác định
- Đảm bảo fixes không break existing code
- Chạy tests sau khi fix

## Bước 6: Review lại fixes

- Tự review lại các changes vừa làm (mini review)
- Đảm bảo fixes address đúng reviewer concerns

## Bước 7: Commit & Push

- Generate commit message:
  ```
  fix(scope): address PR review feedback
  
  - Fix: [description of fix 1]
  - Fix: [description of fix 2]
  - Answer: [response to question]
  
  PR: #42
  ```
- Stage, commit, push lên cùng branch

## Bước 8: Prepare Reply Comments

Chuẩn bị reply cho từng review comment:

```
💬 SUGGESTED REPLIES
════════════════════

Comment 1 (src/auth/refresh.ts:25):
  Reply: "Fixed! Added refresh token expiry validation before rotation. See commit abc1234."

Comment 2 (src/auth/types.ts:10):
  Reply: "Good point! Changed to union type for better type safety. See commit abc1234."

Comment 3 (src/middleware/auth.ts:45):
  Reply: "Applied guard clause pattern, much cleaner now. Thanks!"

Comment 4 (tests/auth/refresh.test.ts:30):
  Reply: "Added test case for expired token scenario. See commit abc1234."
```

User có thể copy replies này để trả lời trên GitHub.

## Bước 9: Summary

```
🔄 GH-REVIEW COMPLETE
══════════════════════

📋 PR #42: feat(auth): add JWT token refresh

📊 Review Response:
  🔧 Fixed: 3 action items
  ❓ Answered: 1 question  
  💡 Applied: 1 suggestion

📝 Commit: def5678 - fix(auth): address PR review feedback
📤 Pushed to: origin/feat/auth-refresh

💬 Reply templates ready (see above)

👉 Next:
  • Post reply comments on GitHub
  • Wait for re-review
  • Run /gh-review again if new comments
```

## Lưu ý

- **KHÔNG** argue với reviewer - fix theo feedback
- Nếu không đồng ý với feedback → trình bày cho user quyết định
- **KHÔNG** thay đổi code ngoài scope của review comments
- Nếu fix phức tạp → hỏi user trước khi implement
- Giữ commit history clean - 1 commit cho tất cả review fixes
