CREATE TABLE users (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    account_id VARCHAR(255) NOT NULL COMMENT 'ユーザーのアカウントID',
    username VARCHAR(255) NOT NULL COMMENT 'ユーザー名',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP, -- noqa: LT05
    UNIQUE uq_account_id (account_id),
    INDEX idx_username (username)
) ENGINE = InnoDB COMMENT 'ユーザー情報テーブル';
