CREATE TABLE posts (
    uuid BINARY(16) NOT NULL PRIMARY KEY,
    user_id BIGINT NOT NULL COMMENT 'ユーザーID',
    content VARCHAR(255) NOT NULL COMMENT '登校内容',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP, -- noqa: LT05
    INDEX idx_user_id (user_id)
) ENGINE = InnoDB COMMENT '投稿テーブル';
