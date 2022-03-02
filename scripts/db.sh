#!/bin/bash

echo "== starting script to recreate db locally =="

echo "--"
read -p "New database name: " database
read -p "MySQL user: " user
read -s -p "MySQL password: " password
echo
echo "--"

echo "== creating database tables, and loading csv data =="
mysql --local_infile=1 -u "$user" -p"$password" <<EOF
DROP DATABASE IF EXISTS $database;
CREATE DATABASE $database;
USE $database;

CREATE TABLE cells(
    identifier VARCHAR(255) NOT NULL,
    accession VARCHAR(20) primary key NOT NULL,
    secondary VARCHAR(500),
    synonyms VARCHAR(500),
    sex VARCHAR(255),
    age VARCHAR(255),
    category VARCHAR(255),
    date VARCHAR(255),
    INDEX identifier (identifier),
    INDEX secondary (secondary),
    INDEX synonyms (synonyms),
    INDEX sex (sex),
    INDEX age (age),
    INDEX category (category),
    INDEX date (date)
);
LOAD DATA LOCAL INFILE '../data/cells.csv' INTO TABLE cells
FIELDS TERMINATED BY ',' ENCLOSED BY '"'
IGNORE 1 LINES;

CREATE TABLE cell_attributes(
    id INT AUTO_INCREMENT primary key NOT NULL,
    accession VARCHAR(20) NOT NULL,
    attribute VARCHAR(20) NOT NULL,
    content VARCHAR(1000) NOT NULL,
    FOREIGN KEY (accession) REFERENCES cells(accession)
);
LOAD DATA LOCAL INFILE '../data/cell_attributes.csv' INTO TABLE cell_attributes
FIELDS TERMINATED BY ',' ENCLOSED BY '"'
IGNORE 1 LINES;

CREATE TABLE refs(
    id INT AUTO_INCREMENT primary key NOT NULL,
    identifier VARCHAR(200) NOT NULL,
    citation VARCHAR(500),
    INDEX identifier (identifier)
);
LOAD DATA LOCAL INFILE '../data/refs.csv' INTO TABLE refs
FIELDS TERMINATED BY ',' ENCLOSED BY '"'
IGNORE 1 LINES;

CREATE TABLE ref_attributes(
    id INT AUTO_INCREMENT primary key NOT NULL,
    identifier VARCHAR(20) NOT NULL,
    attribute VARCHAR(20) NOT NULL,
    content VARCHAR(1000) NOT NULL,
    FOREIGN KEY (identifier) REFERENCES refs(identifier)
);
LOAD DATA LOCAL INFILE '../data/ref_attributes.csv' INTO TABLE ref_attributes
FIELDS TERMINATED BY ',' ENCLOSED BY '"'
IGNORE 1 LINES;

CREATE TABLE statistics(
    id INT AUTO_INCREMENT primary key NOT NULL,
    attribute VARCHAR(255) NOT NULL,
    count VARCHAR(255) NOT NULL
);
LOAD DATA LOCAL INFILE '../data/statistics.csv' INTO TABLE statistics
FIELDS TERMINATED BY ',' ENCLOSED BY '"'
IGNORE 1 LINES;
EOF

echo "== DONE =="
