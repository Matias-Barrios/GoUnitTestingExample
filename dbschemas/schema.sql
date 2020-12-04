
CREATE TABLE Users (
	id serial PRIMARY KEY,
	firstname VARCHAR ( 80 ) NOT NULL,
	lastname VARCHAR ( 80 ) NOT NULL,
	email VARCHAR ( 255 ) UNIQUE NOT NULL,
	age INT NOT NULL,
	created_on VARCHAR ( 40 )  NOT NULL
);

