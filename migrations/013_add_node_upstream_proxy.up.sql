ALTER TABLE nodes
    ADD COLUMN outbound_type VARCHAR(32) NOT NULL DEFAULT 'direct' AFTER traffic_pool,
    ADD COLUMN outbound_proxy_url TEXT NULL AFTER outbound_ip;

CREATE INDEX idx_nodes_outbound_type ON nodes (outbound_type);
