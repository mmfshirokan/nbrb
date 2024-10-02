CREATE TABLE currency (
    id int,
    cur_date TIMESTAMP,
    abbreviation CHAR(3),
    scale int,
    name VARCHAR(100),
    officialRate decimal(10, 4),
    PRIMARY KEY (id, cur_date)
);