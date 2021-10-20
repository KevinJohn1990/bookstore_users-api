CREATE TABLE user(  
    id int NOT NULL primary key AUTO_INCREMENT comment 'primary key',
    first_name VARCHAR(45) NULL,
    last_name VARCHAR(45) NULL,
    email VARCHAR(45) NOT NULL,
    date_created DATETIME COMMENT 'created date',
    UNIQUE INDEX email_unique (email ASC)
) default charset utf8 comment '';