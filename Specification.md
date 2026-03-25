# 📘 LMS (Learning Management System) - System Design Specification

## 1. Tổng quan hệ thống

Hệ thống LMS (Learning Management System) là nền tảng học trực tuyến cho phép:
- Quản lý khóa học
- Quản lý học sinh, giáo viên
- Thanh toán khóa học
- Theo dõi tiến độ học tập

---

## 2. Kiến trúc hệ thống

### 2.1 System Overview

```
┌─────────────────────────────────────────────────────────────────────────┐
│                           CLIENT LAYER                                   │
│  ┌──────────────────────────────────────────────────────────────────┐   │
│  │  Next.js 14 (App Router)                                          │   │
│  │  ├── React Server Components                                      │   │
│  │  ├── Client Components (Interactivity)                           │   │
│  │  ├── Tailwind CSS + shadcn/ui                                    │   │
│  │  └── React Query / Zustand (State)                               │   │
│  └──────────────────────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────────────────────┘
                                    │
                                    │ HTTPS / REST API / JSON
                                    │
                                    ▼
┌─────────────────────────────────────────────────────────────────────────┐
│                           API GATEWAY LAYER                            │
│  ┌──────────────────────────────────────────────────────────────────┐   │
│  │  Nginx / AWS ALB                                                  │   │
│  │  ├── Rate Limiting                                               │   │
│  │  ├── SSL Termination                                             │   │
│  │  ├── Load Balancing                                              │   │
│  │  └── Request Routing                                             │   │
│  └──────────────────────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────────────────────┘
                                    │
                                    │
                                    ▼
┌─────────────────────────────────────────────────────────────────────────┐
│                          APPLICATION LAYER                               │
│  ┌──────────────────────────────────────────────────────────────────┐   │
│  │  Golang (Gin Framework) - API Server                              │   │
│  │  ├── Middleware (Auth, CORS, Logging)                          │   │
│  │  ├── Handlers (REST Controllers)                                 │   │
│  │  ├── Services (Business Logic)                                   │   │
│  │  ├── Repositories (Data Access)                                  │   │
│  │  └── Workers (Background Jobs)                                   │   │
│  └──────────────────────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────────────────────┘
                                    │
                    ┌───────────────┼───────────────┐
                    │               │               │
                    ▼               ▼               ▼
┌─────────────────────────┐ ┌─────────────────┐ ┌─────────────────────┐
│     DATA LAYER          │ │   EXTERNAL      │ │   STORAGE LAYER     │
│  ┌─────────────────┐    │ │  ┌───────────┐  │ │  ┌─────────────────┐│
│  │ MySQL 8.0       │    │ │  │ Stripe    │  │ │  │ AWS S3 /        ││
│  │ ├── Users       │    │ │  │ API       │  │ │  │ Cloud Storage   ││
│  │ ├── Courses     │    │ │  └───────────┘  │ │  │                 ││
│  │ ├── Enrollments │    │ │  ┌───────────┐  │ │  │ - Videos        ││
│  │ └── Payments    │    │ │  │ SendGrid  │  │ │  │ - Images        ││
│  └─────────────────┘    │ │  │ (Email)   │  │ │  │ - Documents     ││
│                         │ │  └───────────┘  │ │  └─────────────────┘│
│  ┌─────────────────┐    │ └─────────────────┘ └─────────────────────┘
│  │ Redis (Cache)   │    │
│  │ ├── Sessions    │    │
│  │ └── Rate Limit  │    │
│  └─────────────────┘    │
└─────────────────────────┘
```

### 2.2 Tech Stack Detail

| Layer | Technology | Version | Purpose |
|-------|------------|---------|---------|
| **Frontend** | Next.js | 14.x | React framework with SSR/SSG |
| | React | 18.x | UI library |
| | TypeScript | 5.x | Type safety |
| | Tailwind CSS | 3.x | Styling |
| | shadcn/ui | latest | Component library |
| | React Query | 5.x | Server state management |
| | Zustand | 4.x | Client state management |
| | React Hook Form | 7.x | Form handling |
| | Zod | 3.x | Schema validation |
| **Backend** | Go | 1.22+ | Programming language |
| | Gin | 1.9+ | HTTP web framework |
| | GORM | 1.25+ | ORM for MySQL |
| | JWT | v5 | Authentication |
| | bcrypt | latest | Password hashing |
| | go-playground/validator | latest | Input validation |
| | zap | latest | Structured logging |
| **Database** | MySQL | 8.0 | Primary database |
| | Redis | 7.x | Cache & sessions |
| **Storage** | AWS S3 | - | File storage |
| | CloudFront | - | CDN for media |
| **Payment** | Stripe | latest | Payment processing |
| **Email** | SendGrid | latest | Transactional email |
| **Infrastructure** | Docker | latest | Containerization |
| | Kubernetes | latest | Orchestration |
| | Nginx | latest | Reverse proxy |
| | AWS ECS/EKS | - | Container hosting |

### 2.3 Backend Architecture (Clean Architecture)

```
lms-backend/
├── cmd/
│   └── api/
│       └── main.go              # Application entry point
├── internal/
│   ├── config/                  # Configuration management
│   │   ├── config.go
│   │   └── database.go
│   ├── domain/                  # Domain models (entities)
│   │   ├── user.go
│   │   ├── course.go
│   │   ├── lesson.go
│   │   ├── enrollment.go
│   │   └── payment.go
│   ├── dto/                     # Data Transfer Objects
│   │   ├── request/
│   │   │   ├── auth_request.go
│   │   │   ├── user_request.go
│   │   │   ├── course_request.go
│   │   │   └── payment_request.go
│   │   └── response/
│   │       ├── auth_response.go
│   │       ├── user_response.go
│   │       ├── course_response.go
│   │       └── payment_response.go
│   ├── repository/              # Data access layer
│   │   ├── user_repository.go
│   │   ├── course_repository.go
│   │   ├── lesson_repository.go
│   │   └── payment_repository.go
│   ├── service/                 # Business logic layer
│   │   ├── auth_service.go
│   │   ├── user_service.go
│   │   ├── course_service.go
│   │   ├── payment_service.go
│   │   └── stripe_service.go
│   ├── handler/                 # HTTP handlers (controllers)
│   │   ├── auth_handler.go
│   │   ├── user_handler.go
│   │   ├── course_handler.go
│   │   ├── lesson_handler.go
│   │   ├── quiz_handler.go
│   │   ├── enrollment_handler.go
│   │   └── payment_handler.go
│   ├── middleware/              # HTTP middlewares
│   │   ├── auth_middleware.go
│   │   ├── cors_middleware.go
│   │   ├── rate_limit.go
│   │   ├── logger.go
│   │   └── error_handler.go
│   ├── pkg/                     # Shared packages
│   │   ├── jwt/
│   │   ├── bcrypt/
│   │   ├── validator/
│   │   ├── errors/
│   │   └── utils/
│   └── worker/                  # Background workers
│       ├── email_worker.go
│       └── payment_webhook_worker.go
├── migrations/                  # Database migrations
│   ├── 001_create_users.sql
│   ├── 002_create_courses.sql
│   └── ...
├── tests/                       # Test suites
│   ├── integration/
│   ├── unit/
│   └── e2e/
├── .env.example
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── go.sum
```

### 2.4 Frontend Architecture (Next.js 14 App Router)

```
lms-frontend/
├── app/                         # Next.js App Router
│   ├── (auth)/                  # Auth route group
│   │   ├── login/
│   │   ├── register/
│   │   ├── forgot-password/
│   │   └── layout.tsx
│   ├── (main)/                  # Main route group
│   │   ├── page.tsx             # Home page
│   │   ├── courses/
│   │   │   ├── page.tsx         # Course list
│   │   │   └── [slug]/
│   │   │       └── page.tsx     # Course detail
│   │   ├── dashboard/
│   │   │   └── page.tsx
│   │   ├── learn/
│   │   │   └── [courseId]/
│   │   │       └── [lessonId]/
│   │   │           └── page.tsx # Lesson player
│   │   └── layout.tsx
│   ├── admin/                   # Admin routes
│   │   ├── users/
│   │   ├── courses/
│   │   ├── payments/
│   │   └── layout.tsx
│   ├── teacher/                 # Teacher routes
│   │   ├── courses/
│   │   ├── students/
│   │   └── analytics/
│   ├── api/                     # API Routes (Next.js)
│   │   └── webhooks/
│   │       └── stripe/
│   ├── layout.tsx               # Root layout
│   └── globals.css
├── components/                  # React components
│   ├── ui/                      # shadcn/ui components
│   │   ├── button.tsx
│   │   ├── card.tsx
│   │   ├── input.tsx
│   │   └── ...
│   ├── auth/                    # Auth components
│   │   ├── login-form.tsx
│   │   ├── register-form.tsx
│   │   └── auth-guard.tsx
│   ├── courses/                 # Course components
│   │   ├── course-card.tsx
│   │   ├── course-list.tsx
│   │   ├── course-player.tsx
│   │   └── lesson-list.tsx
│   ├── dashboard/               # Dashboard components
│   │   ├── stats-card.tsx
│   │   ├── progress-chart.tsx
│   │   └── revenue-chart.tsx
│   ├── layout/                  # Layout components
│   │   ├── header.tsx
│   │   ├── sidebar.tsx
│   │   └── footer.tsx
│   └── shared/                  # Shared components
│       ├── video-player.tsx
│       ├── file-uploader.tsx
│       └── pagination.tsx
├── hooks/                       # Custom React hooks
│   ├── use-auth.ts
│   ├── use-courses.ts
│   ├── use-enrollment.ts
│   └── use-progress.ts
├── lib/                         # Utilities & configurations
│   ├── api/                     # API client
│   │   ├── client.ts            # Axios/fetch config
│   │   ├── auth-api.ts
│   │   ├── course-api.ts
│   │   └── payment-api.ts
│   ├── utils/
│   │   ├── cn.ts                # Tailwind merge
│   │   ├── format.ts            # Format helpers
│   │   └── validation.ts        # Validation helpers
│   └── constants.ts             # App constants
├── stores/                      # Zustand stores
│   ├── auth-store.ts
│   ├── course-store.ts
│   └── ui-store.ts
├── types/                       # TypeScript types
│   ├── user.ts
│   ├── course.ts
│   ├── lesson.ts
│   └── api.ts
├── public/                      # Static assets
│   └── images/
├── .env.local.example
├── next.config.js
├── tailwind.config.ts
├── tsconfig.json
└── package.json
```

### 2.5 Layer Communication Flow

```
┌─────────────────────────────────────────────────────────────────┐
│                        Request Flow                              │
├─────────────────────────────────────────────────────────────────┤
│                                                                  │
│  Client Request                                                  │
│       │                                                          │
│       ▼                                                          │
│  ┌─────────────┐                                               │
│  │   Router    │  (Gin - route matching)                       │
│  └──────┬──────┘                                               │
│         │                                                        │
│         ▼                                                        │
│  ┌─────────────┐     ┌─────────────┐     ┌─────────────┐       │
│  │ Middleware  │────▶│   Handler   │────▶│   Service   │       │
│  │  (Auth,     │     │ (Controller)│     │  (Business) │       │
│  │   Validate) │     └─────────────┘     └──────┬──────┘       │
│  └─────────────┘                                │               │
│                                                  │               │
│                                         ┌────────▼────────┐      │
│                                         │   Repository    │      │
│                                         │  (Data Access)  │      │
│                                         └────────┬────────┘      │
│                                                  │               │
│                                         ┌────────▼────────┐      │
│                                         │     MySQL       │      │
│                                         │   (Database)    │      │
│                                         └─────────────────┘      │
│                                                                  │
│  Response Flow (Reverse)                                         │
│  MySQL ──▶ Repository ──▶ Service ──▶ Handler ──▶ Middleware ──▶ Client│
│                                                                  │
└─────────────────────────────────────────────────────────────────┘
```

### 2.6 Authentication Flow

```
┌─────────┐                    ┌─────────────┐                    ┌─────────┐
│  Client │                    │   Backend   │                    │  Redis  │
└────┬────┘                    └──────┬──────┘                    └────┬────┘
     │                                │                               │
     │  1. POST /auth/login           │                               │
     │  {email, password}             │                               │
     │───────────────────────────────▶│                               │
     │                                │                               │
     │                                │  2. Validate credentials      │
     │                                │  3. Generate JWT tokens       │
     │                                │                               │
     │                                │  4. Store refresh token       │
     │                                │──────────────────────────────▶│
     │                                │                               │
     │  5. Return {access, refresh}   │                               │
     │◀───────────────────────────────│                               │
     │                                │                               │
     │  6. Store tokens in          │                               │
     │     httpOnly cookies           │                               │
     │                                │                               │
     │  7. Send access token in       │                               │
     │     Authorization header     │                               │
     │───────────────────────────────▶│                               │
     │                                │                               │
     │                                │  8. Validate JWT              │
     │                                │  9. Extract user from claims  │
     │                                │                               │
     │  10. Return protected data     │                               │
     │◀───────────────────────────────│                               │
     │                                │                               │
```

### 2.7 Payment Flow (Stripe)

```
┌─────────┐     ┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│  Client │     │   Backend   │     │   Stripe    │     │   Webhook   │
└────┬────┘     └──────┬──────┘     └──────┬──────┘     └──────┬──────┘
     │                  │                   │                   │
     │ 1. Enroll paid   │                   │                   │
     │    course        │                   │                   │
     │─────────────────▶│                   │                   │
     │                  │                   │                   │
     │                  │ 2. Create         │                   │
     │                  │    Checkout       │                   │
     │                  │    Session        │                   │
     │                  │──────────────────▶│                   │
     │                  │                   │                   │
     │                  │ 3. Return         │                   │
     │                  │    checkout_url   │                   │
     │                  │◀──────────────────│                   │
     │                  │                   │                   │
     │ 4. Redirect to    │                   │                   │
     │    Stripe        │                   │                   │
     │◀─────────────────│                   │                   │
     │                  │                   │                   │
     │ 5. User completes│                   │                   │
     │    payment on    │                   │                   │
     │    Stripe        │                   │                   │
     │                  │                   │                   │
     │                  │                   │ 6. Payment        │
     │                  │                   │    completed     │
     │                  │                   │──────────────────▶│
     │                  │                   │                   │
     │                  │                   │                   │ 7. Validate
     │                  │                   │                   │    & process
     │                  │◀──────────────────────────────────────│
     │                  │                   │                   │
     │                  │ 8. Activate       │                   │
     │                  │    enrollment     │                   │
     │                  │                   │                   │
     │ 9. Redirect back │                   │                   │
     │    to success    │                   │                   │
     │◀─────────────────│                   │                   │
     │                  │                   │                   │
```

---

## 3. Roles & Permissions

### 3.1 Các vai trò

- **Admin**
- **Teacher (Giáo viên)**
- **Student (Học sinh)**

---

### 3.2 Quyền hạn

| Chức năng                  | Admin | Teacher | Student |
|--------------------------|-------|---------|---------|
| Quản lý khóa học          | ✅    | ✅      | ❌      |
| Quản lý học sinh          | ✅    | ❌      | ❌      |
| Quản lý giáo viên         | ✅    | ❌      | ❌      |
| Upload nội dung           | ✅    | ✅      | ❌      |
| Đăng ký khóa học          | ❌    | ❌      | ✅      |
| Học khóa học              | ❌    | ❌      | ✅      |
| Xem doanh thu             | ✅    | ❌      | ❌      |

---

## 4. Chức năng chính

---

### 4.1 Authentication

- Đăng ký (Student / Teacher)
- Đăng nhập
- Quên mật khẩu
- Xác thực email
- JWT Authentication

---

### 4.2 Quản lý người dùng

#### Admin:
- CRUD học sinh
- CRUD giáo viên
- Phân quyền

#### User:
- Cập nhật profile
- Upload avatar

---

### 4.3 Quản lý khóa học

#### Thuộc tính khóa học:
- id
- title
- description
- price (0 = miễn phí)
- thumbnail
- teacher_id
- status (draft/published)

#### Chức năng:
- Admin/Teacher tạo khóa học
- Publish / Unpublish khóa học
- Gán giáo viên

---

### 4.4 Cấu trúc nội dung khóa học

#### Cấu trúc:
Course
├── Section (Chương)
│ ├── Lesson (Bài học)
│ │ ├── Video
│ │ ├── Tài liệu
│ │ ├── Nội dung text
│ │ └── Quiz (bài luyện tập)


---

### 4.5 Lesson (Bài học)

#### Các loại nội dung:
- Video
- Text (Markdown/HTML)
- File download

---

### 4.6 Quiz / Bài luyện tập

#### Sau mỗi bài học sẽ có:
- Câu hỏi trắc nghiệm
- Câu hỏi nhiều lựa chọn
- Điền đáp án

#### Thuộc tính:
- question
- options
- correct_answer
- explanation

#### Chức năng:
- Submit bài
- Chấm điểm tự động
- Lưu kết quả

---

### 4.7 Đăng ký khóa học

#### Flow:

1. User chọn khóa học
2. Nếu:
   - Miễn phí → enroll ngay
   - Có phí → thanh toán Stripe
3. Sau khi thanh toán thành công → unlock khóa học

---

### 4.8 Payment (Stripe)

#### Chức năng:
- Thanh toán 1 lần
- Webhook xử lý:
  - payment_success
  - payment_failed

#### Dữ liệu lưu:
- user_id
- course_id
- amount
- status
- stripe_session_id

---

### 4.9 Tiến độ học tập

- Tracking:
  - Bài đã học
  - % hoàn thành
- Đánh dấu hoàn thành lesson
- Hiển thị progress bar

---

### 4.10 Dashboard

#### Student:
- Danh sách khóa học đã đăng ký
- Tiến độ học
- Lịch sử thanh toán

#### Teacher:
- Khóa học của mình
- Số học sinh
- Doanh thu khóa học

#### Admin:
- Tổng số user
- Tổng doanh thu
- Tổng khóa học

---

## 5. Database Design

### 5.1 ERD Relationships

```
users ||--o{ courses : "creates (as teacher)"
users ||--o{ enrollments : "enrolls"
users ||--o{ payments : "makes"
users ||--o{ progress : "tracks"
users ||--o{ quiz_results : "completes"

courses ||--|{ sections : "contains"
courses ||--o{ enrollments : "has"
courses ||--o{ payments : "receives"

sections ||--|{ lessons : "contains"

lessons ||--o{ quizzes : "has"
lessons ||--o{ progress : "tracked in"

quizzes ||--o{ quiz_results : "has"
```

### 5.2 Tables Schema

#### users
Lưu trữ thông tin người dùng (Student, Teacher, Admin)

| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | `VARCHAR(36)` | PK | UUID v4 |
| email | `VARCHAR(255)` | UNIQUE, NOT NULL, INDEX | Email đăng nhập |
| password_hash | `VARCHAR(255)` | NOT NULL | Bcrypt hash |
| name | `VARCHAR(255)` | NOT NULL | Tên hiển thị |
| role | `ENUM('student','teacher','admin')` | NOT NULL, INDEX | Vai trò |
| avatar_url | `VARCHAR(500)` | NULL | URL ảnh đại diện |
| email_verified | `BOOLEAN` | DEFAULT FALSE | Đã xác thực email |
| reset_token | `VARCHAR(255)` | NULL, INDEX | Token reset password |
| reset_token_expires | `DATETIME` | NULL | Hạn token reset |
| last_login_at | `DATETIME` | NULL | Lần đăng nhập cuối |
| is_active | `BOOLEAN` | DEFAULT TRUE, INDEX | Trạng thái active |
| deleted_at | `DATETIME` | NULL, INDEX | Soft delete |
| created_at | `DATETIME` | DEFAULT CURRENT_TIMESTAMP | Ngày tạo |
| updated_at | `DATETIME` | ON UPDATE CURRENT_TIMESTAMP | Ngày cập nhật |

**Indexes:** `email`, `role`, `deleted_at` (composite với `is_active`)

---

#### courses
Thông tin khóa học

| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | `VARCHAR(36)` | PK | UUID v4 |
| title | `VARCHAR(255)` | NOT NULL, INDEX | Tên khóa học |
| slug | `VARCHAR(255)` | UNIQUE, NOT NULL, INDEX | URL-friendly slug |
| description | `TEXT` | NULL | Mô tả khóa học |
| short_description | `VARCHAR(500)` | NULL | Tóm tắt ngắn |
| price | `DECIMAL(10,2)` | DEFAULT 0 | Giá (0 = miễn phí) |
| currency | `VARCHAR(3)` | DEFAULT 'USD' | Loại tiền |
| thumbnail_url | `VARCHAR(500)` | NULL | URL thumbnail |
| trailer_video_url | `VARCHAR(500)` | NULL | URL video giới thiệu |
| teacher_id | `VARCHAR(36)` | FK → users.id, NOT NULL, INDEX | Giáo viên phụ trách |
| status | `ENUM('draft','published','archived')` | DEFAULT 'draft', INDEX | Trạng thái |
| level | `ENUM('beginner','intermediate','advanced')` | NULL | Cấp độ |
| duration_minutes | `INT` | DEFAULT 0 | Tổng thời lượng (phút) |
| total_lessons | `INT` | DEFAULT 0 | Tổng số bài học |
| published_at | `DATETIME` | NULL | Ngày publish |
| meta_title | `VARCHAR(255)` | NULL | SEO title |
| meta_description | `VARCHAR(500)` | NULL | SEO description |
| is_featured | `BOOLEAN` | DEFAULT FALSE, INDEX | Khóa học nổi bật |
| deleted_at | `DATETIME` | NULL, INDEX | Soft delete |
| created_at | `DATETIME` | DEFAULT CURRENT_TIMESTAMP | Ngày tạo |
| updated_at | `DATETIME` | ON UPDATE CURRENT_TIMESTAMP | Ngày cập nhật |

**Indexes:** `slug`, `teacher_id`, `status`, `is_featured`, `deleted_at`
**Constraints:** `price >= 0`, `duration_minutes >= 0`

---

#### sections
Các chương trong khóa học

| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | `VARCHAR(36)` | PK | UUID v4 |
| course_id | `VARCHAR(36)` | FK → courses.id, NOT NULL, INDEX | Thuộc khóa học |
| title | `VARCHAR(255)` | NOT NULL | Tên chương |
| description | `TEXT` | NULL | Mô tả chương |
| sort_order | `INT` | NOT NULL, DEFAULT 0 | Thứ tự sắp xếp |
| is_published | `BOOLEAN` | DEFAULT TRUE | Được publish |
| created_at | `DATETIME` | DEFAULT CURRENT_TIMESTAMP | Ngày tạo |
| updated_at | `DATETIME` | ON UPDATE CURRENT_TIMESTAMP | Ngày cập nhật |

**Indexes:** `course_id`, `sort_order`
**Constraints:** `sort_order >= 0`

---

#### lessons
Các bài học trong chương

| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | `VARCHAR(36)` | PK | UUID v4 |
| section_id | `VARCHAR(36)` | FK → sections.id, NOT NULL, INDEX | Thuộc chương |
| title | `VARCHAR(255)` | NOT NULL | Tên bài học |
| description | `TEXT` | NULL | Mô tả bài học |
| content_type | `ENUM('video','text','file','quiz')` | NOT NULL | Loại nội dung |
| video_url | `VARCHAR(500)` | NULL | URL video (nếu là video) |
| video_duration | `INT` | DEFAULT 0 | Thời lượng video (giây) |
| text_content | `LONGTEXT` | NULL | Nội dung text/HTML |
| file_url | `VARCHAR(500)` | NULL | URL file download |
| file_name | `VARCHAR(255)` | NULL | Tên file gốc |
| file_size | `INT` | DEFAULT 0 | Kích thước file (bytes) |
| is_free_preview | `BOOLEAN` | DEFAULT FALSE | Xem thử miễn phí |
| sort_order | `INT` | NOT NULL, DEFAULT 0 | Thứ tự trong chương |
| is_published | `BOOLEAN` | DEFAULT TRUE | Được publish |
| created_at | `DATETIME` | DEFAULT CURRENT_TIMESTAMP | Ngày tạo |
| updated_at | `DATETIME` | ON UPDATE CURRENT_TIMESTAMP | Ngày cập nhật |

**Indexes:** `section_id`, `content_type`, `sort_order`
**Constraints:** `video_duration >= 0`, `sort_order >= 0`

---

#### quizzes
Bài luyện tập trắc nghiệm

| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | `VARCHAR(36)` | PK | UUID v4 |
| lesson_id | `VARCHAR(36)` | FK → lessons.id, NOT NULL, INDEX | Thuộc bài học |
| question | `TEXT` | NOT NULL | Câu hỏi |
| question_type | `ENUM('single_choice','multiple_choice','fill_blank')` | NOT NULL | Loại câu hỏi |
| options | `JSON` | NOT NULL | Các lựa chọn `[{"id": "a", "text": "..."}, ...]` |
| correct_answers | `JSON` | NOT NULL | Đáp án đúng `["a", "b"]` hoặc `["text"]` |
| explanation | `TEXT` | NULL | Giải thích đáp án |
| points | `INT` | DEFAULT 1 | Điểm câu hỏi |
| sort_order | `INT` | DEFAULT 0 | Thứ tự |
| created_at | `DATETIME` | DEFAULT CURRENT_TIMESTAMP | Ngày tạo |
| updated_at | `DATETIME` | ON UPDATE CURRENT_TIMESTAMP | Ngày cập nhật |

**Indexes:** `lesson_id`, `sort_order`
**Constraints:** `points > 0`

---

#### enrollments
Đăng ký khóa học của học sinh

| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | `VARCHAR(36)` | PK | UUID v4 |
| user_id | `VARCHAR(36)` | FK → users.id, NOT NULL, INDEX | Học sinh |
| course_id | `VARCHAR(36)` | FK → courses.id, NOT NULL, INDEX | Khóa học |
| status | `ENUM('pending','active','completed','expired','cancelled')` | DEFAULT 'pending', INDEX | Trạng thái |
| progress_percent | `DECIMAL(5,2)` | DEFAULT 0.00 | % hoàn thành |
| completed_lessons | `INT` | DEFAULT 0 | Số bài đã hoàn thành |
| total_lessons | `INT` | DEFAULT 0 | Tổng số bài |
| enrolled_at | `DATETIME` | DEFAULT CURRENT_TIMESTAMP | Ngày đăng ký |
| completed_at | `DATETIME` | NULL | Ngày hoàn thành |
| expires_at | `DATETIME` | NULL | Ngày hết hạn |
| last_accessed_at | `DATETIME` | NULL | Lần truy cập cuối |
| created_at | `DATETIME` | DEFAULT CURRENT_TIMESTAMP | Ngày tạo |
| updated_at | `DATETIME` | ON UPDATE CURRENT_TIMESTAMP | Ngày cập nhật |

**Indexes:** `user_id`, `course_id`, `status`, UNIQUE(`user_id`, `course_id`)
**Constraints:** `progress_percent BETWEEN 0 AND 100`

---

#### progress
Tiến độ học tập chi tiết

| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | `VARCHAR(36)` | PK | UUID v4 |
| user_id | `VARCHAR(36)` | FK → users.id, NOT NULL, INDEX | Học sinh |
| lesson_id | `VARCHAR(36)` | FK → lessons.id, NOT NULL, INDEX | Bài học |
| enrollment_id | `VARCHAR(36)` | FK → enrollments.id, NOT NULL | Enrollment |
| status | `ENUM('not_started','in_progress','completed')` | DEFAULT 'not_started' | Trạng thái |
| watch_time_seconds | `INT` | DEFAULT 0 | Thời gian đã xem (video) |
| last_position_seconds | `INT` | DEFAULT 0 | Vị trí dừng cuối |
| completed_at | `DATETIME` | NULL | Ngày hoàn thành |
| created_at | `DATETIME` | DEFAULT CURRENT_TIMESTAMP | Ngày tạo |
| updated_at | `DATETIME` | ON UPDATE CURRENT_TIMESTAMP | Ngày cập nhật |

**Indexes:** `user_id`, `lesson_id`, `enrollment_id`, UNIQUE(`user_id`, `lesson_id`)

---

#### quiz_results
Kết quả làm bài trắc nghiệm

| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | `VARCHAR(36)` | PK | UUID v4 |
| user_id | `VARCHAR(36)` | FK → users.id, NOT NULL, INDEX | Học sinh |
| quiz_id | `VARCHAR(36)` | FK → quizzes.id, NOT NULL, INDEX | Bài quiz |
| lesson_id | `VARCHAR(36)` | FK → lessons.id, NOT NULL, INDEX | Bài học |
| answers | `JSON` | NOT NULL | Câu trả lời của user `{"question_id": "answer"}` |
| score | `INT` | DEFAULT 0 | Điểm đạt được |
| max_score | `INT` | DEFAULT 0 | Điểm tối đa |
| percentage | `DECIMAL(5,2)` | DEFAULT 0.00 | % điểm |
| is_passed | `BOOLEAN` | DEFAULT FALSE | Đạt/ Không đạt |
| attempt_number | `INT` | DEFAULT 1 | Lần thử thứ mấy |
| time_spent_seconds | `INT` | DEFAULT 0 | Thời gian làm bài |
| submitted_at | `DATETIME` | DEFAULT CURRENT_TIMESTAMP | Ngày nộp |
| created_at | `DATETIME` | DEFAULT CURRENT_TIMESTAMP | Ngày tạo |

**Indexes:** `user_id`, `quiz_id`, `lesson_id`, `submitted_at`
**Constraints:** `score >= 0`, `max_score > 0`, `percentage BETWEEN 0 AND 100`

---

#### payments
Thanh toán khóa học

| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | `VARCHAR(36)` | PK | UUID v4 |
| user_id | `VARCHAR(36)` | FK → users.id, NOT NULL, INDEX | Người thanh toán |
| course_id | `VARCHAR(36)` | FK → courses.id, NOT NULL, INDEX | Khóa học |
| enrollment_id | `VARCHAR(36)` | FK → enrollments.id, NULL | Enrollment (sau khi thanh toán) |
| amount | `DECIMAL(10,2)` | NOT NULL | Số tiền |
| currency | `VARCHAR(3)` | DEFAULT 'USD' | Loại tiền |
| status | `ENUM('pending','completed','failed','refunded','cancelled')` | DEFAULT 'pending', INDEX | Trạng thái |
| payment_method | `ENUM('stripe','paypal')` | DEFAULT 'stripe' | Phương thức |
| stripe_session_id | `VARCHAR(255)` | NULL, UNIQUE | Stripe Checkout Session ID |
| stripe_payment_intent_id | `VARCHAR(255)` | NULL | Stripe Payment Intent ID |
| stripe_customer_id | `VARCHAR(255)` | NULL | Stripe Customer ID |
| paid_at | `DATETIME` | NULL | Ngày thanh toán thành công |
| refunded_at | `DATETIME` | NULL | Ngày hoàn tiền |
| refund_amount | `DECIMAL(10,2)` | DEFAULT 0 | Số tiền hoàn |
| metadata | `JSON` | NULL | Metadata bổ sung |
| created_at | `DATETIME` | DEFAULT CURRENT_TIMESTAMP | Ngày tạo |
| updated_at | `DATETIME` | ON UPDATE CURRENT_TIMESTAMP | Ngày cập nhật |

**Indexes:** `user_id`, `course_id`, `status`, `stripe_session_id`
**Constraints:** `amount > 0`

---

#### refresh_tokens
Lưu trữ refresh tokens cho JWT

| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | `VARCHAR(36)` | PK | UUID v4 |
| user_id | `VARCHAR(36)` | FK → users.id, NOT NULL, INDEX | User |
| token_hash | `VARCHAR(255)` | UNIQUE, NOT NULL, INDEX | Hash của token |
| expires_at | `DATETIME` | NOT NULL, INDEX | Hạn token |
| created_at | `DATETIME` | DEFAULT CURRENT_TIMESTAMP | Ngày tạo |
| revoked_at | `DATETIME` | NULL | Ngày revoke |
| ip_address | `VARCHAR(45)` | NULL | IP tạo token |
| user_agent | `VARCHAR(500)` | NULL | User agent |

**Indexes:** `user_id`, `token_hash`, `expires_at`

---

### 5.3 Indexes Strategy

| Table | Index | Purpose |
|-------|-------|---------|
| users | `idx_email` | Login lookup |
| users | `idx_role_active` | Filter by role |
| users | `idx_deleted_at` | Soft delete queries |
| courses | `idx_teacher_status` | List teacher's courses |
| courses | `idx_slug` | Course detail by slug |
| courses | `idx_featured` | Featured courses list |
| enrollments | `idx_user_enrollment` | My courses list |
| enrollments | `idx_course_enrollment` | Course students count |
| payments | `idx_user_payment` | Payment history |
| payments | `idx_status` | Webhook processing |
| progress | `idx_user_lesson` | Progress lookup |
| quiz_results | `idx_user_quiz` | Quiz attempt history |

---

### 5.4 Soft Delete Strategy

- Tables có soft delete: `users`, `courses`
- Xóa = set `deleted_at = NOW()` thay vì DELETE
- Query mặc định: `WHERE deleted_at IS NULL`
- Admin có thể xem deleted records để restore

### 5.5 Data Integrity Rules

1. **Cascade Delete**: Khi xóa `courses`, tự động xóa `sections` → `lessons` → `quizzes`
2. **Restrict Delete**: Không cho xóa `users` nếu có `courses` (teacher) hoặc `enrollments`
3. **Set NULL**: Khi xóa `sections`, set `section_id = NULL` trong related records (nếu có)
4. **Check Constraints**: Giá tối thiểu 0, progress 0-100, v.v.

---

## 6. API Design (Golang + Next.js)

### 6.1 Base Configuration

| Property | Value |
|----------|-------|
| Base URL | `https://api.lmsrocket.com/api/v1` |
| Protocol | REST API + JSON |
| Auth | Bearer Token (JWT) |
| Content-Type | `application/json` |
| Pagination | Cursor-based cho list APIs |

### 6.2 Common Response Format

#### Success Response
```json
{
  "success": true,
  "data": { ... },
  "meta": {
    "page": 1,
    "limit": 20,
    "total": 100,
    "total_pages": 5
  }
}
```

#### Error Response
```json
{
  "success": false,
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "Invalid input data",
    "details": [
      { "field": "email", "message": "Email is required" }
    ]
  }
}
```

#### Error Codes
| Code | HTTP | Description |
|------|------|-------------|
| `VALIDATION_ERROR` | 400 | Input validation failed |
| `UNAUTHORIZED` | 401 | Authentication required |
| `FORBIDDEN` | 403 | Permission denied |
| `NOT_FOUND` | 404 | Resource not found |
| `CONFLICT` | 409 | Resource conflict (duplicate) |
| `RATE_LIMITED` | 429 | Too many requests |
| `INTERNAL_ERROR` | 500 | Server error |
| `SERVICE_UNAVAILABLE` | 503 | External service error (Stripe, etc.) |

### 6.3 Authentication APIs

#### POST /auth/register
**Description:** Đăng ký tài khoản mới  
**Auth:** ❌ Không cần  
**Roles:** ✅ Student, Teacher

**Request:**
```json
{
  "email": "user@example.com",
  "password": "SecurePassword123!",
  "name": "Nguyen Van A",
  "role": "student"
}
```

**Validation Rules:**
- `email`: valid email, unique, max 255 chars
- `password`: min 8 chars, ít nhất 1 uppercase, 1 lowercase, 1 number
- `name`: min 2 chars, max 255 chars
- `role`: enum [`student`, `teacher`]

**Response 201:**
```json
{
  "success": true,
  "data": {
    "user": {
      "id": "uuid",
      "email": "user@example.com",
      "name": "Nguyen Van A",
      "role": "student",
      "email_verified": false,
      "created_at": "2026-03-25T10:00:00Z"
    },
    "access_token": "eyJhbG...",
    "refresh_token": "dGhpcyBpcyBhIHJlZnJlc2g...",
    "expires_in": 900
  }
}
```

---

#### POST /auth/login
**Description:** Đăng nhập  
**Auth:** ❌ Không cần

**Request:**
```json
{
  "email": "user@example.com",
  "password": "SecurePassword123!"
}
```

**Response 200:**
```json
{
  "success": true,
  "data": {
    "user": { ... },
    "access_token": "eyJhbG...",
    "refresh_token": "dGhpcyBpcyBhIHJlZnJlc2g...",
    "expires_in": 900
  }
}
```

**Error Codes:**
- `INVALID_CREDENTIALS` (401): Email/password không đúng
- `ACCOUNT_DISABLED` (403): Tài khoản bị vô hiệu hóa

---

#### POST /auth/refresh
**Description:** Làm mới access token  
**Auth:** ❌ Không cần (dùng refresh token)

**Request:**
```json
{
  "refresh_token": "dGhpcyBpcyBhIHJlZnJlc2g..."
}
```

**Response 200:**
```json
{
  "success": true,
  "data": {
    "access_token": "eyJhbG...",
    "refresh_token": "bmV3IHJlZnJlc2g...",
    "expires_in": 900
  }
}
```

---

#### POST /auth/logout
**Description:** Đăng xuất (revoke refresh token)  
**Auth:** ✅ Cần

**Response 200:**
```json
{
  "success": true,
  "message": "Logged out successfully"
}
```

---

#### POST /auth/forgot-password
**Description:** Yêu cầu reset password  
**Auth:** ❌ Không cần

**Request:**
```json
{
  "email": "user@example.com"
}
```

**Response 200:** (Luôn trả về success để tránh user enumeration)
```json
{
  "success": true,
  "message": "If email exists, reset instructions sent"
}
```

---

#### POST /auth/reset-password
**Description:** Đặt lại password với token  
**Auth:** ❌ Không cần

**Request:**
```json
{
  "token": "reset_token_from_email",
  "new_password": "NewSecurePassword123!"
}
```

**Validation:** Password rules như register

---

#### POST /auth/verify-email
**Description:** Xác thực email  
**Auth:** ❌ Không cần

**Request:**
```json
{
  "token": "verification_token_from_email"
}
```

---

#### POST /auth/resend-verification
**Description:** Gửi lại email xác thực  
**Auth:** ✅ Cần

---

### 6.4 User APIs

#### GET /users/me
**Description:** Lấy thông tin user hiện tại  
**Auth:** ✅ Cần

**Response 200:**
```json
{
  "success": true,
  "data": {
    "id": "uuid",
    "email": "user@example.com",
    "name": "Nguyen Van A",
    "role": "student",
    "avatar_url": "https://cdn...",
    "email_verified": true,
    "created_at": "2026-03-25T10:00:00Z",
    "stats": {
      "enrolled_courses": 5,
      "completed_courses": 2,
      "total_learning_hours": 48
    }
  }
}
```

---

#### PATCH /users/me
**Description:** Cập nhật profile  
**Auth:** ✅ Cần

**Request:**
```json
{
  "name": "Nguyen Van B",
  "avatar_url": "https://cdn..."
}
```

---

#### POST /users/me/avatar
**Description:** Upload avatar  
**Auth:** ✅ Cần  
**Content-Type:** `multipart/form-data`

**Request:**
```
file: <image_file> (max 5MB, jpg/png/webp)
```

**Response 200:**
```json
{
  "success": true,
  "data": {
    "avatar_url": "https://cdn.../avatar_uuid.jpg"
  }
}
```

---

#### POST /users/me/change-password
**Description:** Đổi password  
**Auth:** ✅ Cần

**Request:**
```json
{
  "current_password": "OldPassword123!",
  "new_password": "NewPassword123!"
}
```

---

### 6.5 Admin - User Management APIs

#### GET /admin/users
**Description:** List tất cả users (Admin only)  
**Auth:** ✅ Cần  
**Roles:** admin

**Query Params:**
- `page`: number (default: 1)
- `limit`: number (default: 20, max: 100)
- `role`: `student` | `teacher` | `admin`
- `status`: `active` | `inactive`
- `search`: string (search by name/email)
- `sort_by`: `created_at` | `name` | `last_login`
- `sort_order`: `asc` | `desc`

**Response 200:**
```json
{
  "success": true,
  "data": {
    "users": [
      {
        "id": "uuid",
        "email": "user@example.com",
        "name": "Nguyen Van A",
        "role": "student",
        "is_active": true,
        "email_verified": true,
        "created_at": "2026-03-25T10:00:00Z",
        "last_login_at": "2026-03-25T12:00:00Z"
      }
    ],
    "meta": {
      "page": 1,
      "limit": 20,
      "total": 150,
      "total_pages": 8
    }
  }
}
```

---

#### GET /admin/users/:id
**Description:** Chi tiết user  
**Auth:** ✅ Cần  
**Roles:** admin

---

#### PATCH /admin/users/:id
**Description:** Cập nhật user (admin can update any field)  
**Auth:** ✅ Cần  
**Roles:** admin

**Request:**
```json
{
  "name": "New Name",
  "role": "teacher",
  "is_active": false
}
```

---

#### DELETE /admin/users/:id
**Description:** Xóa user (soft delete)  
**Auth:** ✅ Cần  
**Roles:** admin

---

### 6.6 Course APIs

#### GET /courses
**Description:** List khóa học published  
**Auth:** ❌ Không cần (Public)

**Query Params:**
- `page`: number
- `limit`: number
- `category`: string
- `level`: `beginner` | `intermediate` | `advanced`
- `price`: `free` | `paid`
- `teacher_id`: string
- `search`: string (search title/description)
- `sort_by`: `created_at` | `price` | `popularity` | `rating`
- `sort_order`: `asc` | `desc`

**Response 200:**
```json
{
  "success": true,
  "data": {
    "courses": [
      {
        "id": "uuid",
        "title": "Go Programming Masterclass",
        "slug": "go-programming-masterclass",
        "short_description": "Learn Go from scratch",
        "price": 49.99,
        "currency": "USD",
        "thumbnail_url": "https://cdn...",
        "level": "intermediate",
        "duration_minutes": 720,
        "total_lessons": 45,
        "teacher": {
          "id": "uuid",
          "name": "John Doe",
          "avatar_url": "https://cdn..."
        },
        "rating": 4.8,
        "students_count": 1250,
        "is_enrolled": false
      }
    ],
    "meta": { ... }
  }
}
```

---

#### GET /courses/:slug
**Description:** Chi tiết khóa học  
**Auth:** ❌ Không cần (Public)

**Response 200:**
```json
{
  "success": true,
  "data": {
    "id": "uuid",
    "title": "Go Programming Masterclass",
    "slug": "go-programming-masterclass",
    "description": "Full description...",
    "price": 49.99,
    "thumbnail_url": "https://cdn...",
    "trailer_video_url": "https://cdn...",
    "level": "intermediate",
    "duration_minutes": 720,
    "total_lessons": 45,
    "status": "published",
    "teacher": { ... },
    "sections": [
      {
        "id": "uuid",
        "title": "Introduction",
        "sort_order": 1,
        "lessons": [
          {
            "id": "uuid",
            "title": "Welcome to the Course",
            "content_type": "video",
            "video_duration": 300,
            "is_free_preview": true,
            "sort_order": 1
          }
        ]
      }
    ],
    "is_enrolled": false,
    "enrollment_count": 1250
  }
}
```

---

#### POST /courses
**Description:** Tạo khóa học mới  
**Auth:** ✅ Cần  
**Roles:** teacher, admin

**Request:**
```json
{
  "title": "New Course",
  "slug": "new-course",
  "description": "Course description...",
  "short_description": "Short version...",
  "price": 29.99,
  "level": "beginner",
  "thumbnail_url": "https://cdn...",
  "trailer_video_url": "https://cdn..."
}
```

---

#### PATCH /courses/:id
**Description:** Cập nhật khóa học  
**Auth:** ✅ Cần  
**Roles:** teacher (own courses), admin

---

#### DELETE /courses/:id
**Description:** Xóa khóa học (soft delete)  
**Auth:** ✅ Cần  
**Roles:** teacher (own courses), admin

---

#### POST /courses/:id/publish
**Description:** Publish khóa học  
**Auth:** ✅ Cần  
**Roles:** teacher (own courses), admin

---

#### POST /courses/:id/unpublish
**Description:** Unpublish khóa học  
**Auth:** ✅ Cần  
**Roles:** teacher (own courses), admin

---

#### GET /courses/:id/students
**Description:** List học sinh trong khóa học  
**Auth:** ✅ Cần  
**Roles:** teacher (own courses), admin

---

### 6.7 Section APIs

#### POST /courses/:id/sections
**Description:** Thêm section vào khóa học  
**Auth:** ✅ Cần  
**Roles:** teacher (own courses), admin

**Request:**
```json
{
  "title": "Chapter 1: Getting Started",
  "description": "Introduction chapter...",
  "sort_order": 1
}
```

---

#### PATCH /sections/:id
**Description:** Cập nhật section  
**Auth:** ✅ Cần  
**Roles:** teacher (own course), admin

---

#### DELETE /sections/:id
**Description:** Xóa section  
**Auth:** ✅ Cần  
**Roles:** teacher (own course), admin

---

### 6.8 Lesson APIs

#### POST /sections/:id/lessons
**Description:** Thêm lesson vào section  
**Auth:** ✅ Cần  
**Roles:** teacher (own course), admin

**Request:**
```json
{
  "title": "Introduction to Variables",
  "description": "Learn about variables...",
  "content_type": "video",
  "video_url": "https://cdn...",
  "video_duration": 600,
  "sort_order": 1,
  "is_free_preview": true
}
```

---

#### PATCH /lessons/:id
**Description:** Cập nhật lesson  
**Auth:** ✅ Cần  
**Roles:** teacher (own course), admin

---

#### DELETE /lessons/:id
**Description:** Xóa lesson  
**Auth:** ✅ Cần  
**Roles:** teacher (own course), admin

---

#### GET /lessons/:id
**Description:** Chi tiết lesson (cần enrollment nếu không phải free preview)  
**Auth:** ✅ Cần (hoặc ❌ nếu là free preview)

**Response 200:**
```json
{
  "success": true,
  "data": {
    "id": "uuid",
    "title": "Introduction to Variables",
    "content_type": "video",
    "video_url": "https://cdn...",
    "video_duration": 600,
    "text_content": null,
    "file_url": null,
    "is_free_preview": true,
    "course": { ... },
    "section": { ... },
    "progress": {
      "status": "in_progress",
      "watch_time_seconds": 120,
      "last_position_seconds": 120,
      "completed_at": null
    },
    "next_lesson": { "id": "uuid", "title": "..." },
    "prev_lesson": { "id": "uuid", "title": "..." }
  }
}
```

---

#### POST /lessons/:id/progress
**Description:** Cập nhật tiến độ học tập  
**Auth:** ✅ Cần  
**Roles:** student (enrolled)

**Request:**
```json
{
  "watch_time_seconds": 300,
  "last_position_seconds": 300,
  "status": "completed"
}
```

---

### 6.9 Quiz APIs

#### POST /lessons/:id/quizzes
**Description:** Thêm quiz vào lesson  
**Auth:** ✅ Cần  
**Roles:** teacher (own course), admin

**Request:**
```json
{
  "question": "What is the output of fmt.Println(1 + 1)?",
  "question_type": "single_choice",
  "options": [
    { "id": "a", "text": "2" },
    { "id": "b", "text": "11" },
    { "id": "c", "text": "Error" }
  ],
  "correct_answers": ["a"],
  "explanation": "1 + 1 = 2 in Go",
  "points": 1,
  "sort_order": 1
}
```

---

#### POST /quizzes/:id/submit
**Description:** Nộp bài quiz  
**Auth:** ✅ Cần  
**Roles:** student (enrolled)

**Request:**
```json
{
  "answers": {
    "quiz_1": "a",
    "quiz_2": ["a", "b"]
  }
}
```

**Response 200:**
```json
{
  "success": true,
  "data": {
    "quiz_results": {
      "id": "uuid",
      "score": 9,
      "max_score": 10,
      "percentage": 90.00,
      "is_passed": true,
      "attempt_number": 1,
      "time_spent_seconds": 300,
      "submitted_at": "2026-03-25T12:00:00Z",
      "details": [
        {
          "quiz_id": "uuid",
          "question": "What is...",
          "user_answer": "a",
          "correct_answer": "a",
          "is_correct": true,
          "points": 1
        }
      ]
    }
  }
}
```

---

### 6.10 Enrollment APIs

#### POST /courses/:id/enroll
**Description:** Đăng ký khóa học  
**Auth:** ✅ Cần  
**Roles:** student

**Response 201 (Free Course):**
```json
{
  "success": true,
  "data": {
    "enrollment": {
      "id": "uuid",
      "course_id": "uuid",
      "status": "active",
      "enrolled_at": "2026-03-25T12:00:00Z"
    }
  }
}
```

**Response 202 (Paid Course - Redirect to Payment):**
```json
{
  "success": true,
  "data": {
    "enrollment": {
      "id": "uuid",
      "status": "pending"
    },
    "payment_required": true,
    "checkout_url": "https://checkout.stripe.com/..."
  }
}
```

---

#### GET /enrollments
**Description:** List khóa học đã đăng ký của user hiện tại  
**Auth:** ✅ Cần  
**Roles:** student

**Query Params:**
- `status`: `active` | `completed` | `expired`
- `progress`: `not_started` | `in_progress` | `all`

**Response 200:**
```json
{
  "success": true,
  "data": {
    "enrollments": [
      {
        "id": "uuid",
        "course": { ... },
        "status": "active",
        "progress_percent": 35.5,
        "completed_lessons": 16,
        "total_lessons": 45,
        "enrolled_at": "2026-03-20T10:00:00Z",
        "last_accessed_at": "2026-03-25T08:00:00Z",
        "next_lesson": { "id": "uuid", "title": "..." }
      }
    ]
  }
}
```

---

### 6.11 Payment APIs

#### POST /payments/checkout
**Description:** Tạo Stripe Checkout Session  
**Auth:** ✅ Cần

**Request:**
```json
{
  "course_id": "uuid"
}
```

**Response 200:**
```json
{
  "success": true,
  "data": {
    "checkout_url": "https://checkout.stripe.com/...",
    "session_id": "cs_test_..."
  }
}
```

---

#### GET /payments/history
**Description:** Lịch sử thanh toán  
**Auth:** ✅ Cần

**Response 200:**
```json
{
  "success": true,
  "data": {
    "payments": [
      {
        "id": "uuid",
        "course": { ... },
        "amount": 49.99,
        "currency": "USD",
        "status": "completed",
        "paid_at": "2026-03-25T12:00:00Z"
      }
    ]
  }
}
```

---

#### POST /payments/webhook
**Description:** Stripe webhook xử lý payment events  
**Auth:** Stripe Signature Verification

---

### 6.12 Dashboard APIs

#### GET /dashboard/student
**Description:** Dashboard cho học sinh  
**Auth:** ✅ Cần  
**Roles:** student

**Response 200:**
```json
{
  "success": true,
  "data": {
    "stats": {
      "total_enrolled": 12,
      "completed_courses": 3,
      "in_progress_courses": 5,
      "total_learning_hours": 48,
      "average_progress": 45.5
    },
    "recent_courses": [ ... ],
    "upcoming_deadlines": [ ... ],
    "achievements": [ ... ]
  }
}
```

---

#### GET /dashboard/teacher
**Description:** Dashboard cho giáo viên  
**Auth:** ✅ Cần  
**Roles:** teacher

**Response 200:**
```json
{
  "success": true,
  "data": {
    "stats": {
      "total_courses": 8,
      "total_students": 1250,
      "total_revenue": 15000.00,
      "this_month_revenue": 2500.00
    },
    "courses": [
      {
        "id": "uuid",
        "title": "...",
        "students_count": 300,
        "revenue": 4500.00,
        "average_rating": 4.8
      }
    ],
    "recent_enrollments": [ ... ]
  }
}
```

---

#### GET /dashboard/admin
**Description:** Dashboard cho admin  
**Auth:** ✅ Cần  
**Roles:** admin

**Response 200:**
```json
{
  "success": true,
  "data": {
    "stats": {
      "total_users": 5000,
      "total_teachers": 50,
      "total_students": 4950,
      "total_courses": 200,
      "total_revenue": 50000.00,
      "this_month_revenue": 8000.00
    },
    "charts": {
      "user_growth": [ ... ],
      "revenue_by_month": [ ... ],
      "top_courses": [ ... ]
    }
  }
}
```

---

### 6.13 Summary: All Endpoints

| Method | Endpoint | Auth | Roles | Description |
|--------|----------|------|-------|-------------|
| POST | /auth/register | ❌ | - | Đăng ký |
| POST | /auth/login | ❌ | - | Đăng nhập |
| POST | /auth/refresh | ❌ | - | Refresh token |
| POST | /auth/logout | ✅ | - | Đăng xuất |
| POST | /auth/forgot-password | ❌ | - | Quên mật khẩu |
| POST | /auth/reset-password | ❌ | - | Reset mật khẩu |
| POST | /auth/verify-email | ❌ | - | Xác thực email |
| POST | /auth/resend-verification | ✅ | - | Gửi lại email xác thực |
| GET | /users/me | ✅ | - | Profile |
| PATCH | /users/me | ✅ | - | Cập nhật profile |
| POST | /users/me/avatar | ✅ | - | Upload avatar |
| POST | /users/me/change-password | ✅ | - | Đổi mật khẩu |
| GET | /admin/users | ✅ | admin | List users |
| GET | /admin/users/:id | ✅ | admin | User detail |
| PATCH | /admin/users/:id | ✅ | admin | Update user |
| DELETE | /admin/users/:id | ✅ | admin | Delete user |
| GET | /courses | ❌ | - | List courses (public) |
| GET | /courses/:slug | ❌ | - | Course detail |
| POST | /courses | ✅ | teacher/admin | Create course |
| PATCH | /courses/:id | ✅ | teacher/admin | Update course |
| DELETE | /courses/:id | ✅ | teacher/admin | Delete course |
| POST | /courses/:id/publish | ✅ | teacher/admin | Publish course |
| POST | /courses/:id/unpublish | ✅ | teacher/admin | Unpublish course |
| GET | /courses/:id/students | ✅ | teacher/admin | List students |
| POST | /courses/:id/sections | ✅ | teacher/admin | Add section |
| PATCH | /sections/:id | ✅ | teacher/admin | Update section |
| DELETE | /sections/:id | ✅ | teacher/admin | Delete section |
| POST | /sections/:id/lessons | ✅ | teacher/admin | Add lesson |
| PATCH | /lessons/:id | ✅ | teacher/admin | Update lesson |
| DELETE | /lessons/:id | ✅ | teacher/admin | Delete lesson |
| GET | /lessons/:id | ✅ | student | Lesson detail |
| POST | /lessons/:id/progress | ✅ | student | Update progress |
| POST | /lessons/:id/quizzes | ✅ | teacher/admin | Add quiz |
| POST | /quizzes/:id/submit | ✅ | student | Submit quiz |
| POST | /courses/:id/enroll | ✅ | student | Enroll course |
| GET | /enrollments | ✅ | student | My enrollments |
| POST | /payments/checkout | ✅ | - | Create checkout |
| GET | /payments/history | ✅ | - | Payment history |
| POST | /payments/webhook | Stripe | - | Stripe webhook |
| GET | /dashboard/student | ✅ | student | Student dashboard |
| GET | /dashboard/teacher | ✅ | teacher | Teacher dashboard |
| GET | /dashboard/admin | ✅ | admin | Admin dashboard |

---

## 12. Security & Authentication Strategy

### 12.1 JWT Token Strategy

#### Token Types

| Token | Location | Expiry | Purpose |
|-------|----------|--------|---------|
| **Access Token** | Authorization Header | 15 minutes | API authentication |
| **Refresh Token** | HttpOnly Cookie | 7 days | Renew access token |

#### JWT Claims Structure

```json
{
  "sub": "user_uuid",
  "email": "user@example.com",
  "role": "student",
  "iat": 1711363200,
  "exp": 1711364100,
  "jti": "unique_token_id",
  "type": "access"
}
```

**Refresh Token Claims:**
```json
{
  "sub": "user_uuid",
  "jti": "unique_refresh_id",
  "iat": 1711363200,
  "exp": 1711968000,
  "type": "refresh",
  "device_id": "device_fingerprint"
}
```

#### Token Storage (Frontend)

```typescript
// Access Token: Memory only (Zustand store)
// Refresh Token: HttpOnly, Secure, SameSite=Strict cookie
// No localStorage/sessionStorage for tokens
```

#### Token Rotation

```
1. Access token expires (15 min)
2. Client detects 401 Unauthorized
3. Client calls /auth/refresh with refresh token cookie
4. Server validates refresh token (from Redis)
5. Server issues NEW access token + NEW refresh token
6. Old refresh token revoked
7. New tokens returned to client
```

### 12.2 Password Security

#### Hashing
- **Algorithm:** bcrypt with cost factor 12
- **Salt:** Automatically generated 16-byte salt
- **Pepper:** Optional additional secret key

```go
// Example Go implementation
import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword(
        []byte(password), 
        bcrypt.DefaultCost, // 10-12
    )
    return string(bytes), err
}

func CheckPassword(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
```

#### Password Requirements
| Rule | Minimum |
|------|---------|
| Length | 8 characters |
| Uppercase | 1 letter |
| Lowercase | 1 letter |
| Number | 1 digit |
| Special | 1 symbol (!@#$%^&*) |

#### Password Reset Flow

```
1. User requests reset → POST /auth/forgot-password
2. Generate cryptographically secure token (32 bytes)
3. Hash token with SHA256, store hash + expiry (1 hour)
4. Send email with plaintext token (never stored plaintext)
5. User clicks link with token
6. Verify token hash matches
7. Allow password change
8. Invalidate all existing sessions
```

### 12.3 Input Validation & Sanitization

#### Validation Rules

| Field | Rules |
|-------|-------|
| Email | RFC 5322 compliant, domain validation |
| Name | 2-255 chars, alphanumeric + spaces |
| Slug | lowercase, alphanumeric + hyphens |
| Price | Positive decimal, max 999999.99 |
| URL | Valid URL format, whitelist domains |
| ID | UUID v4 format |

#### SQL Injection Prevention

```go
// ✅ Safe - Use parameterized queries
result := db.Where("email = ?", email).First(&user)

// ✅ Safe - Use GORM
result := db.Where(map[string]interface{}{"email": email}).First(&user)

// ❌ Unsafe - String concatenation
result := db.Raw("SELECT * FROM users WHERE email = '" + email + "'")
```

#### XSS Prevention

```typescript
// Frontend - Escape all user content
import DOMPurify from 'dompurify';

const sanitized = DOMPurify.sanitize(userContent);

// Backend - Validate and sanitize all inputs
import "github.com/microcosm-cc/bluemonday"

p := bluemonday.UGCPolicy()
clean := p.Sanitize(dirtyInput)
```

### 12.4 Rate Limiting

#### Tiers

| Endpoint Group | Limit | Window | Notes |
|----------------|-------|--------|-------|
| **Public APIs** (courses list) | 100 | 1 minute | Per IP |
| **Auth APIs** (login/register) | 5 | 1 minute | Per IP + email |
| **Authenticated APIs** | 1000 | 1 minute | Per user |
| **File Upload** | 10 | 1 minute | Per user |
| **Payment** | 10 | 1 minute | Per user |

#### Redis Implementation

```go
// Rate limiting middleware
func RateLimitMiddleware(limit int, window time.Duration) gin.HandlerFunc {
    return func(c *gin.Context) {
        key := fmt.Sprintf("rate_limit:%s", c.ClientIP())
        
        count, err := redisClient.Incr(key).Result()
        if err != nil {
            c.Next()
            return
        }
        
        if count == 1 {
            redisClient.Expire(key, window)
        }
        
        if count > int64(limit) {
            c.JSON(429, gin.H{"error": "Rate limit exceeded"})
            c.Abort()
            return
        }
        
        c.Next()
    }
}
```

### 12.5 CORS Policy

```go
// Allowed origins
allowedOrigins := []string{
    "https://lmsrocket.com",
    "https://www.lmsrocket.com",
    "https://app.lmsrocket.com",
}

// Development only
if os.Getenv("APP_ENV") == "development" {
    allowedOrigins = append(allowedOrigins, "http://localhost:3000")
}

config := cors.Config{
    AllowOrigins:     allowedOrigins,
    AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
    AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Request-ID"},
    ExposeHeaders:    []string{"Content-Length", "X-Request-ID"},
    AllowCredentials: true,
    MaxAge:           12 * time.Hour,
}
```

### 12.6 Security Headers

| Header | Value | Purpose |
|--------|-------|---------|
| `X-Content-Type-Options` | `nosniff` | Prevent MIME sniffing |
| `X-Frame-Options` | `DENY` | Prevent clickjacking |
| `X-XSS-Protection` | `1; mode=block` | XSS filter (legacy) |
| `Strict-Transport-Security` | `max-age=31536000; includeSubDomains` | Force HTTPS |
| `Content-Security-Policy` | Custom policy | Resource loading control |
| `Referrer-Policy` | `strict-origin-when-cross-origin` | Referrer control |
| `Permissions-Policy` | `camera=(), microphone=(), geolocation=()` | Feature policy |

### 12.7 Session Management

#### Concurrent Session Limits

| Role | Max Sessions |
|------|-------------|
| Student | 5 |
| Teacher | 3 |
| Admin | 2 |

#### Session Tracking

```go
type Session struct {
    ID        string    `json:"id"`         // Session UUID
    UserID    string    `json:"user_id"`    
    Device    string    `json:"device"`     // User agent
    IP        string    `json:"ip"`         // IP address
    CreatedAt time.Time `json:"created_at"`
    LastUsed  time.Time `json:"last_used"`
    IsCurrent bool      `json:"is_current"` // Current session flag
}
```

### 12.8 File Upload Security

#### Validation Rules

| Check | Rule |
|-------|------|
| File size | Max 100MB for video, 10MB for documents |
| MIME type | Whitelist: video/mp4, image/jpeg/png/webp, application/pdf |
| Extension | Must match MIME type |
| Magic bytes | Verify file signature |
| Virus scan | ClamAV integration |

#### Storage
- **Videos:** S3 with CloudFront CDN
- **Images:** S3 with image optimization (sharp/imgproxy)
- **Documents:** S3 with virus scan
- **Signed URLs:** 15-minute expiry for private content

### 12.9 Payment Security

#### PCI Compliance
- Never store full credit card numbers
- Use Stripe Checkout (PCI compliant hosted page)
- Webhook signature verification

```go
// Stripe webhook signature verification
func VerifyWebhook(payload []byte, signature string, secret string) error {
    return webhook.ConstructEvent(payload, signature, secret)
}
```

#### Fraud Detection
- 3D Secure for high-value transactions
- Velocity checks (max 3 attempts per card per hour)
- IP geolocation mismatch alerts

### 12.10 Audit Logging

#### Logged Events

| Event | Data Captured |
|-------|---------------|
| Login | IP, user agent, success/failure, timestamp |
| Password change | IP, timestamp |
| Payment | Amount, method, status, IP |
| Course enrollment | Course ID, user ID, timestamp |
| Admin actions | Action type, target, before/after values |
| Data exports | What data, who requested, timestamp |

#### Log Retention
- Application logs: 30 days
- Audit logs: 7 years (compliance)
- Payment logs: 7 years

### 12.11 Environment Security

#### Required Environment Variables

```bash
# JWT
JWT_SECRET_KEY=your-256-bit-secret-min-32-chars
JWT_REFRESH_SECRET=another-256-bit-secret

# Database
DB_HOST=localhost
DB_PORT=3306
DB_NAME=lms_rocket
DB_USER=lms_user
DB_PASSWORD=strong-random-password

# Redis
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=redis-password

# Stripe
STRIPE_SECRET_KEY=sk_live_...
STRIPE_WEBHOOK_SECRET=whsec_...

# Email
SENDGRID_API_KEY=SG.xxx
FROM_EMAIL=noreply@lmsrocket.com

# Storage
AWS_ACCESS_KEY_ID=AKIA...
AWS_SECRET_ACCESS_KEY=...
S3_BUCKET_NAME=lmsrocket-assets

# Security
ENCRYPTION_KEY=32-char-key-for-data-encryption
RATE_LIMIT_ENABLED=true
CORS_ORIGIN=https://lmsrocket.com
```

#### Secret Management
- Development: `.env` file (gitignored)
- Staging/Production: AWS Secrets Manager / HashiCorp Vault
- Rotation: Every 90 days for API keys

---

## 7. Frontend (Next.js)

### Pages:
- Home
- Course List
- Course Detail
- Lesson Player
- Dashboard
- Admin Panel

### State Management:
- Redux / Zustand / React Query

---

## 8. Non-functional Requirements

- Scalability
- Secure Payment
- Responsive UI
- Video streaming optimization
- Logging & monitoring

---

## 9. Future Enhancements

- Live class (Zoom integration)
- Certificate
- AI recommendation khóa học
- Discussion / Forum
- Mobile app

---

## 10. Deployment

- Frontend: Vercel
- Backend: Docker + Kubernetes / VPS
- DB: AWS RDS / Cloud SQL

---

## 11. Tổng kết

Hệ thống LMS này:
- Hỗ trợ đầy đủ học online
- Có payment tích hợp Stripe
- Dễ mở rộng
- Phù hợp SaaS education platform
