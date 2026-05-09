ALTER TABLE plans
    ADD COLUMN normal_traffic_multiplier DECIMAL(8,3) NOT NULL DEFAULT 1.000 AFTER residential_traffic_limit,
    ADD COLUMN residential_traffic_multiplier DECIMAL(8,3) NOT NULL DEFAULT 1.000 AFTER normal_traffic_multiplier;

ALTER TABLE usage_ledgers
    ADD COLUMN billing_multiplier DECIMAL(8,3) NOT NULL DEFAULT 1.000 AFTER traffic_pool,
    ADD COLUMN billed_upload BIGINT UNSIGNED NOT NULL DEFAULT 0 AFTER delta_upload,
    ADD COLUMN billed_download BIGINT UNSIGNED NOT NULL DEFAULT 0 AFTER delta_download,
    ADD COLUMN billed_total BIGINT UNSIGNED NOT NULL DEFAULT 0 AFTER delta_total;

UPDATE plans
SET normal_traffic_multiplier = 1.000
WHERE normal_traffic_multiplier IS NULL OR normal_traffic_multiplier <= 0;

UPDATE plans
SET residential_traffic_multiplier = 1.000
WHERE residential_traffic_multiplier IS NULL OR residential_traffic_multiplier <= 0;

UPDATE usage_ledgers
SET billing_multiplier = 1.000,
    billed_upload = delta_upload,
    billed_download = delta_download,
    billed_total = delta_total
WHERE billed_total = 0 AND delta_total > 0;
