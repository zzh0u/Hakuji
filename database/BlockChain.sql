CREATE TABLE book (
    id INT PRIMARY KEY AUTO_INCREMENT,
    isbn VARCHAR(18) UNIQUE NOT NULL COMMENT '国际标准书号，唯一标识一本书',
    title VARCHAR(255) NOT NULL COMMENT '书名以及副标题',
    author VARCHAR(100) NOT NULL COMMENT '作者姓名，支持多个作者和译者',
    cover_url VARCHAR(255) DEFAULT '' COMMENT '封面图片 CDN 相对路径格式: /covers/{book_id}_{timestamp}.jpg',
    cover_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '最后更新时间，用于客户端缓存刷新',
    publisher VARCHAR(100) NOT NULL COMMENT '出版社名称',
    published_date DATE NOT NULL COMMENT '出版日期',
    category VARCHAR(50) NOT NULL COMMENT '图书分类，如玄幻、科幻、数学、物理、计算机等',
    content_summary TEXT COMMENT '图书内容简介',
    rating DECIMAL(3,1) DEFAULT 0.0 COMMENT '图书评分，范围 0.0-9.9，步长0.1',
    download_count INT DEFAULT 0 COMMENT '下载次数',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '图书表'

CREATE TABLE user (
    id INT PRIMARY KEY AUTO_INCREMENT,
    avatar_url VARCHAR(255) DEFAULT '' COMMENT 'CDN 相对路径格式: /avatars/{user_id}_{timestamp}.jpg',
    avatar_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '最后更新时间，用于客户端缓存刷新',
    username VARCHAR(50) UNIQUE NOT NULL COMMENT '登录用户名',
    password_hash CHAR(64) NOT NULL COMMENT 'SHA256+HMAC 加密，PBKDF2 算法迭代 10000 次（salt=${salt}）',
    salt CHAR(64) NOT NULL COMMENT 'CSPRNG 生成的 32 字节加密盐，每秒生成限制 100 次',
    email VARCHAR(100) UNIQUE NOT NULL COMMENT '邮箱',
	phone VARCHAR(20) UNIQUE NOT NULL COMMENT '手机号',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '用户表'

CREATE TABLE block(
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    hash VARCHAR(64) NOT NULL COMMENT '区块哈希值',
    pre_hash VARCHAR(64) NOT NULL COMMENT '前一个区块哈希值',
    timestamp BIGINT NOT NULL COMMENT '时间戳',
    data TEXT NOT NULL COMMENT '区块数据',
    nonce INTEGER NOT NULL COMMENT '随机数，用于调整区块哈希以满足难度目标',
    miner_id INTEGER NOT NULL COMMENT '挖矿者 ID',
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4

CREATE TABLE checkout (
    id INTEGER PRIMARY AUTO_INCREMENT,
    book_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL
) DEFAULT CHARSET=utf8mb4 COMMENT '借阅表';

CREATE TABLE favorite (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id INT,
    book_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    deleted_at TIMESTAMP DEFAULT NULL COMMENT '删除时间，NULL 表示未删除'
) DEFAULT CHARSET=utf8mb4 COMMENT '收藏表'

CREATE TABLE comment (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id INT,
    book_id INT,
    parent_id BIGINT DEFAULT NULL COMMENT '父评论 ID',
    content TEXT NOT NULL COMMENT '评论内容',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    deleted_at TIMESTAMP DEFAULT NULL COMMENT '删除时间，NULL 表示未删除'
) DEFAULT CHARSET=utf8mb4 COMMENT '评论表'

-- 评论表外键
ALTER TABLE comments
    ADD CONSTRAINT fk_comments_user_id    -- 用户外键
    FOREIGN KEY (user_id) 
    REFERENCES user(id) 
    ON DELETE SET NULL 
    ON UPDATE CASCADE,

    ADD CONSTRAINT fk_comments_book_id    -- 图书外键
    FOREIGN KEY (book_id) 
    REFERENCES book(id) 
    ON DELETE CASCADE 
    ON UPDATE CASCADE,

    ADD CONSTRAINT fk_comments_parent_id  -- 父评论自关联外键
    FOREIGN KEY (parent_id) 
    REFERENCES comment(id) 
    ON DELETE CASCADE 
    ON UPDATE CASCADE;

-- 收藏表外键
ALTER TABLE favorites
    ADD CONSTRAINT fk_favorites_user_id   -- 用户外键
    FOREIGN KEY (user_id) 
    REFERENCES user(id) 
    ON DELETE CASCADE 
    ON UPDATE CASCADE,

    ADD CONSTRAINT fk_favorites_book_id   -- 图书外键
    FOREIGN KEY (book_id) 
    REFERENCES book(id) 
    ON DELETE CASCADE 
    ON UPDATE CASCADE;
