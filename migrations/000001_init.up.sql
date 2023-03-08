-- accounts
CREATE TABLE accounts (
	id serial PRIMARY KEY,
	address VARCHAR ( 42 ) UNIQUE NOT NULL,
	username VARCHAR ( 50 ) NOT NULL,
	email VARCHAR ( 255 ) UNIQUE NOT NULL,
	bio TEXT NULL,
	image VARCHAR( 255 ) NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
  last_login TIMESTAMP 
);

-- collections
CREATE TABLE collections (
	id serial PRIMARY KEY,
	contract_address VARCHAR ( 42 ) UNIQUE NOT NULL,
	slug VARCHAR ( 50 ) UNIQUE NOT NULL,
	name VARCHAR ( 50 ) UNIQUE NOT NULL,
	description TEXT NULL,
	image_url VARCHAR( 255 ) NULL,
	royalty INT,
	fee_recipient VARCHAR( 42 ) NOT NULL,
	category VARCHAR( 255 ) NULL,
	website VARCHAR( 50 ) NULL,
	discord VARCHAR( 50 ) NULL,
	twitter_url VARCHAR( 50 ) NULL,
	instagram_url VARCHAR( 50 ) NULL,
	medium_url VARCHAR( 50 ) NULL,
	telegram_url VARCHAR( 50 ) NULL,
	contact_email VARCHAR( 50 ) NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL
);

-- items
CREATE TABLE items (
	id serial PRIMARY KEY,
	content_type VARCHAR ( 50 ) NOT NULL,
	contract_address VARCHAR ( 42 ) NOT NULL,
	image_url VARCHAR ( 256 ) NOT NULL,
	is_appropriate VARCHAR ( 50 ) NOT NULL,
	last_sale_price NUMERIC NULL,
	last_sale_price_in_usd NUMERIC NULL,
	last_sale_price_payment_token VARCHAR ( 42 ),
	liked INT NULL,
	name VARCHAR ( 50 ) NULL,
	payment_token VARCHAR ( 42 ) NULL,
	price BIGINT NULL,
	price_in_usd NUMERIC NULL,
	supply INT NULL,
	thumbnail_path VARCHAR ( 256 ) NULL,
	token_id INT NULL,
	token_type INT NULL,
	token_uri VARCHAR ( 256 ) UNIQUE NOT NULL,
	collection_id INT NULL,
	creator VARCHAR ( 42 ) NOT NULL,
	owner VARCHAR ( 42 ) NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	deleted_at_unix BIGINT NULL
);

-- registries
CREATE TABLE registries (
	id serial PRIMARY KEY,
	contract_type VARCHAR ( 20 ) NOT NULL,
	contract_address VARCHAR ( 42 ) UNIQUE NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	deleted_at_unix BIGINT NULL
);

-- marketplaces
CREATE TABLE marketplaces (
	id serial PRIMARY KEY,
	address_registry VARCHAR ( 42 ) NULL,
	platform_fee VARCHAR ( 20 ) NULL,
	fee_recipient VARCHAR ( 42 ) NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	deleted_at_unix BIGINT NULL
);

-- tokens
CREATE TABLE tokens (
	id serial PRIMARY KEY,
	token VARCHAR ( 42 ) UNIQUE NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	deleted_at_unix BIGINT NULL
);