CREATE TABLE if not exists words (
    ID   int NOT NULL AUTO_INCREMENT,
    Word varchar(255),
    Freq int DEFAULT 0,

    PRIMARY KEY(ID)
);

CREATE TABLE if not exists trie (
    ID int NOT NULL AUTO_INCREMENT,
    trie blob,

    PRIMARY KEY(ID)
);