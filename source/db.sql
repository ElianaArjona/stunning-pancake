CREATE TABLE report (
    id UUID PRIMARY KEY,
    bank    COLLATE,
	month   INT,
	year    INT,
	income  FLOAT,
	expense FLOAT,
	balance FLOAT,
	serviceType COLLATE
);