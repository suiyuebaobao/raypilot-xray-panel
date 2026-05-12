ALTER TABLE user_subscriptions
    ADD COLUMN speed_limit_bps BIGINT UNSIGNED NOT NULL DEFAULT 0 AFTER residential_used_traffic;
