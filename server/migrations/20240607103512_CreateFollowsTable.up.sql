CREATE TABLE follows (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT NOT NULL COMMENT 'ユーザーID',
    follow_user_id BIGINT NOT NULL COMMENT 'フォローユーザーID',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP, -- noqa: LT05
    INDEX idx_user_id (user_id),
    INDEX idx_follow_user_id (follow_user_id)
) ENGINE = InnoDB COMMENT 'ユーザーフォロー情報テーブル';
