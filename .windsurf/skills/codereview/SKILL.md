---
name: codereview
description: A brief description, shown to the model to help it understand when to use this skill
---

---
description: Skill hướng dẫn cách review code hiệu quả theo chuẩn LMS Rocket project
---

# Code Review Skill

## Mục đích
Cung cấp hướng dẫn chi tiết để AI thực hiện code review một cách nhất quán và hiệu quả.

## Review Process

### 1. Security Scan
```
Priority: CRITICAL
Auto-fix: NO (cần user approve)
```

Tìm kiếm patterns nguy hiểm:
- **Hardcoded secrets**: Regex `(password|secret|api_key|token)\s*[:=]\s*['"][^'"]+['"]`
- **SQL Injection**: Raw string concatenation trong SQL queries
- **XSS**: `dangerouslySetInnerHTML`, unescaped user input trong HTML
- **Path traversal**: User input trong file paths không sanitize
- **Insecure crypto**: MD5, SHA1 cho passwords (nên dùng bcrypt/argon2)
- **CORS misconfiguration**: `origin: '*'` trong production

### 2. TypeScript Quality
```
Priority: HIGH
Auto-fix: YES
```

- Missing type annotations → Infer và thêm type
- `any` type → Replace với proper type
- Missing return types trên public functions
- Unused variables/imports → Remove
- Non-null assertions (`!`) không cần thiết → Replace với optional chaining

### 3. Error Handling
```
Priority: HIGH
Auto-fix: YES (basic cases)
```

- Async functions không có try/catch → Wrap trong try/catch
- Empty catch blocks → Thêm error logging
- Swallowed errors → Thêm proper error propagation
- Missing error boundaries (React) → Flag warning

### 4. Performance
```
Priority: MEDIUM
Auto-fix: PARTIAL
```

- N+1 queries → Flag, suggest batch query
- Missing `useMemo`/`useCallback` cho expensive operations
- Large bundle imports → Suggest tree-shaking / dynamic import
- Missing indexes hints → Flag for database queries
- Unnecessary re-renders → Add React.memo hoặc useMemo

### 5. Code Style
```
Priority: LOW
Auto-fix: YES
```

- Naming conventions: camelCase, PascalCase, kebab-case
- Import ordering: builtin → external → internal → relative
- Consistent quotes (single vs double)
- Trailing commas
- Console.log statements → Remove

## Auto-fix Decision Matrix

| Issue Type | Severity | Auto-fix? | Reason |
|-----------|----------|-----------|--------|
| Unused imports | Low | ✅ Yes | Safe, no logic change |
| Missing types | Medium | ✅ Yes | Infer from usage |
| Console.log | Low | ✅ Yes | Debug artifact |
| Missing error handling | High | ✅ Yes (basic) | Add try/catch wrapper |
| Security vulnerability | Critical | ❌ No | Needs human review |
| Business logic change | High | ❌ No | Needs human decision |
| Architecture issue | Medium | ❌ No | Needs discussion |
| Performance optimization | Medium | ⚠️ Partial | Simple cases only |

## Review Output Format

Mỗi issue phải có:
1. **File path và line number**
2. **Severity**: Critical / High / Medium / Low
3. **Category**: Security / Types / Error Handling / Performance / Style
4. **Description**: Mô tả ngắn gọn vấn đề
5. **Fix**: Đã auto-fix hoặc suggested fix
6. **Reference**: Link đến best practice (nếu có)
