ALTER TABLE follows
ADD UNIQUE uq_user_id_follow_user_id (user_id, follow_user_id);
