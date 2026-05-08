DROP INDEX idx_nodes_outbound_type ON nodes;

ALTER TABLE nodes
    DROP COLUMN outbound_proxy_url,
    DROP COLUMN outbound_type;
