CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE IF NOT EXISTS
product(
	id uuid DEFAULT gen_random_uuid() PRIMARY KEY, 
	name VARCHAR, 
	img VARCHAR,
	metadata VARCHAR
);
