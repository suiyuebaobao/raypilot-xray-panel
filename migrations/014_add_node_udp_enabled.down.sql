DROP INDEX idx_nodes_udp_enabled ON nodes;

ALTER TABLE nodes
    DROP COLUMN udp_enabled;
