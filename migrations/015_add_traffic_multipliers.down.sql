ALTER TABLE usage_ledgers
    DROP COLUMN billed_total,
    DROP COLUMN billed_download,
    DROP COLUMN billed_upload,
    DROP COLUMN billing_multiplier;

ALTER TABLE plans
    DROP COLUMN residential_traffic_multiplier,
    DROP COLUMN normal_traffic_multiplier;
