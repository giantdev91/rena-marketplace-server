-- admin password is "admin1"
-- user1 password is "user1"

INSERT INTO accounts (username, email, password, created_at, updated_at) VALUES
('admin', 'admin@email.com', '$2a$10$qUq8XJ.RcxcQvaEipicm/OLVPp5AjJjoPigj4vlRU579Xz0SkZwqu', now(), now()),
('user1', 'user1@email.com', '$2a$10$lsYsLv8nGPM0.R.ft4sgpe3OP7..KL3ZJqqhSVCKTEnSCMUztoUcW', now(), now());

INSERT INTO collections (contract_address, fee_recipient, slug, name, description, image_url, contact_email, created_at, updated_at) VALUES
('0xea2aB5F5da50ecfb9b0424a35BfaeF0d952a0Ae5', '0x990f5Df13b6894FbA311Af4069bcD425eaf954ed', 'rena-default-collection', 'Rena Default Collection', 'This is a rena default collection', 'https://artion.mypinata.cloud/ipfs/Qmf24waMhqm5cqLHnn63yyGGVKwR5PeAdi5UBWD79A5kdq/4309.png', 'hiroki.moto.pro@gmail.com', now(), now());