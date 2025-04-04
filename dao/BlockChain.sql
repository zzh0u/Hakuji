CREATE TABLE book (
    id SERIAL PRIMARY KEY,
    isbn VARCHAR(18) UNIQUE NOT NULL,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(100) NOT NULL,
    cover_url VARCHAR(255) DEFAULT '',
    hash VARCHAR(64) NOT NULL,
    pre_hash VARCHAR(64) NOT NULL,
    publisher VARCHAR(100) NOT NULL,
    published_date DATE NOT NULL,
    category VARCHAR(50) NOT NULL,
    content_summary TEXT,
    rating DECIMAL(3,1) DEFAULT 0.0,
    download_count INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

COMMENT ON TABLE book IS '图书表';
COMMENT ON COLUMN book.isbn IS '国际标准书号，唯一标识一本书';
COMMENT ON COLUMN book.title IS '书名以及副标题';
COMMENT ON COLUMN book.author IS '作者姓名，支持多个作者和译者';
COMMENT ON COLUMN book.cover_url IS '封面图片 CDN 相对路径格式: /covers/{book_id}_{timestamp}.jpg';
COMMENT ON COLUMN book.hash IS '区块哈希值';
COMMENT ON COLUMN book.pre_hash IS '区块哈希值';
COMMENT ON COLUMN book.publisher IS '出版社名称';
COMMENT ON COLUMN book.published_date IS '出版日期';
COMMENT ON COLUMN book.category IS '图书分类，如玄幻、科幻、数学、物理、计算机等';
COMMENT ON COLUMN book.content_summary IS '图书内容简介';
COMMENT ON COLUMN book.rating IS '图书评分，范围 0.0-9.9，步长0.1';
COMMENT ON COLUMN book.download_count IS '下载次数';
COMMENT ON COLUMN book.created_at IS '创建时间';
COMMENT ON COLUMN book.deleted_at IS '删除时间，逻辑删除，非空为已删除';

CREATE TABLE user (
    id SERIAL PRIMARY KEY,
    avatar_url VARCHAR(255) DEFAULT '',
    avatar_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    username VARCHAR(50) UNIQUE NOT NULL,
    password_hash CHAR(64) NOT NULL,
    salt CHAR(64) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    phone VARCHAR(20) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

COMMENT ON TABLE user IS '用户表';
COMMENT ON COLUMN user.avatar_url IS 'CDN 相对路径格式: /avatars/{user_id}_{timestamp}.jpg';
COMMENT ON COLUMN user.avatar_updated_at IS '最后更新时间，用于客户端缓存刷新';
COMMENT ON COLUMN user.username IS '登录用户名';
COMMENT ON COLUMN user.password_hash IS 'SHA256+HMAC 加密，PBKDF2 算法迭代 10000 次（salt=${salt}）';
COMMENT ON COLUMN user.salt IS 'CSPRNG 生成的 32 字节加密盐，每秒生成限制 100 次';
COMMENT ON COLUMN user.email IS '邮箱';
COMMENT ON COLUMN user.phone IS '手机号';
COMMENT ON COLUMN user.created_at IS '创建时间';
COMMENT ON COLUMN user.updated_at IS '更新时间';
COMMENT ON COLUMN user.deleted_at IS '注销用户，逻辑删除，非空为已删除';

CREATE TABLE download (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    book_id INTEGER NOT NULL
);

COMMENT ON TABLE download IS '下载记录表';

CREATE TABLE borrow (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    book_id INTEGER NOT NULL
);

COMMENT ON TABLE borrow IS '借阅记录表';

CREATE TABLE favorite (
    id BIGSERIAL PRIMARY KEY,
    user_id INTEGER,
    book_id INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

COMMENT ON TABLE favorite IS '收藏表';
COMMENT ON COLUMN favorite.created_at IS '创建时间';
COMMENT ON COLUMN favorite.deleted_at IS '删除时间，NULL 表示未删除';

CREATE TABLE comment (
    id BIGSERIAL PRIMARY KEY,
    user_id INTEGER,
    book_id INTEGER,
    parent_id BIGINT DEFAULT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

COMMENT ON TABLE comment IS '评论表';
COMMENT ON COLUMN comment.parent_id IS '父评论 ID';
COMMENT ON COLUMN comment.content IS '评论内容';
COMMENT ON COLUMN comment.created_at IS '创建时间';
COMMENT ON COLUMN comment.deleted_at IS '删除时间，NULL 表示未删除';