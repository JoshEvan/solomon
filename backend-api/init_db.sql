CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE IF NOT EXISTS
product(
	id uuid DEFAULT gen_random_uuid() PRIMARY KEY, 
	name VARCHAR, 
	img VARCHAR,
	price NUMERIC,
	create_time timestamp DEFAULT NOW(),
	update_time timestamp DEFAULT NOW()
);


CREATE TABLE IF NOT EXISTS
tag (
	id smallint PRIMARY KEY,
	name VARCHAR
);

-- CREATE TABLE IF NOT EXISTS
-- product_tagging (
-- 	id uuid DEFAULT gen_random_uuid() PRIMARY KEY, 
-- 	product_id uuid FOREIGN KEY product,
-- 	tag_id smallint FOREIGN KEY tag
-- );
