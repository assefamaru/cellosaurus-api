#!/bin/bash

ROOT="$(dirname "$0")"
cd "$ROOT/.."

echo "--"
read -p "New database name: " database
read -p "MySQL user: " user
read -s -p "MySQL password: " password
echo
echo "--"

mysql --local_infile=1 -u"$user" -p"$password" <<EOF
DROP DATABASE IF EXISTS $database;
CREATE DATABASE $database;
USE $database;

CREATE TABLE cells(
    identifier VARCHAR(500) NOT NULL,
    accession VARCHAR(20) primary key NOT NULL,
    secondary VARCHAR(500),
    synonyms VARCHAR(500),
    sex VARCHAR(255),
    age VARCHAR(255),
    category VARCHAR(255),
    date VARCHAR(500),
    INDEX identifier (identifier),
    INDEX secondary (secondary),
    INDEX synonyms (synonyms),
    INDEX sex (sex),
    INDEX age (age),
    INDEX category (category),
    INDEX date (date)
)
ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;
LOAD DATA LOCAL INFILE 'data/cells.csv' INTO TABLE cells
FIELDS TERMINATED BY ','
ENCLOSED BY '"'
IGNORE 1 LINES;

CREATE TABLE cell_attributes(
    id INT AUTO_INCREMENT primary key NOT NULL,
    accession VARCHAR(20) NOT NULL,
    attribute VARCHAR(20) NOT NULL,
    content VARCHAR(1000) NOT NULL,
    FOREIGN KEY (accession) REFERENCES cells(accession)
)
ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;
LOAD DATA LOCAL INFILE 'data/cell_attributes.csv' INTO TABLE cell_attributes
FIELDS TERMINATED BY ','
ENCLOSED BY '"'
IGNORE 1 LINES;

CREATE TABLE refs(
    id INT AUTO_INCREMENT primary key NOT NULL,
    identifier VARCHAR(500) NOT NULL,
    citation VARCHAR(1000),
    INDEX identifier (identifier)
)
ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;
LOAD DATA LOCAL INFILE 'data/refs.csv' INTO TABLE refs
FIELDS TERMINATED BY ','
ENCLOSED BY '"'
IGNORE 1 LINES;

CREATE TABLE ref_attributes(
    id INT AUTO_INCREMENT primary key NOT NULL,
    identifier VARCHAR(500) NOT NULL,
    attribute VARCHAR(20) NOT NULL,
    content VARCHAR(1000) NOT NULL
)
ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;
LOAD DATA LOCAL INFILE 'data/ref_attributes.csv' INTO TABLE ref_attributes
FIELDS TERMINATED BY ','
ENCLOSED BY '"'
IGNORE 1 LINES;

CREATE TABLE xrefs(
    id INT AUTO_INCREMENT primary key NOT NULL,
    abbrev VARCHAR(500) NOT NULL,
    name VARCHAR(500) NOT NULL,
    server VARCHAR(500) NOT NULL,
    url VARCHAR(500) NOT NULL,
    term VARCHAR(500) NOT NULL,
    cat VARCHAR(500) NOT NULL,
    INDEX identifier (abbrev)
)
ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;
LOAD DATA LOCAL INFILE 'data/xrefs.csv' INTO TABLE xrefs
FIELDS TERMINATED BY ','
ENCLOSED BY '"'
IGNORE 1 LINES;
EOF
