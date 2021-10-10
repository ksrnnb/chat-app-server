CREATE TABLE sessions (
    id VARCHAR(255) NOT NULL PRIMARY KEY,
    user_id INTEGER NOT NULL COMMENT 'ユーザーID',
    expired_at DATETIME NOT NULL COMMENT '有効期限',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日',
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日',
    CONSTRAINT FOREIGN KEY (user_id) REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE
);
