CREATE TABLE user (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    login VARCHAR(50) NOT NULL UNIQUE,
    password BLOB NOT NULL
);

CREATE TABLE genre (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE book (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INT NOT NULL,
    title VARCHAR(150) NOT NULL,
    genre_id INT NULL,
    author VARCHAR(150) NULL,
    year INT NULL,
    description TEXT NULL,

    CONSTRAINT book_user_fk
        FOREIGN KEY (user_id)
        REFERENCES user(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    CONSTRAINT book_genre_fk
        FOREIGN KEY (genre_id)
        REFERENCES genre(id)
        ON DELETE SET NULL
        ON UPDATE CASCADE
);
