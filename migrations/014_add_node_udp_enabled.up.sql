ALTER TABLE nodes
    ADD COLUMN udp_enabled BOOLEAN NOT NULL DEFAULT TRUE AFTER flow;

UPDATE nodes
SET udp_enabled = FALSE
WHERE outbound_type = 'socks5';

CREATE INDEX idx_nodes_udp_enabled ON nodes (udp_enabled);
