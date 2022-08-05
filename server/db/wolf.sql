-- Model setup
CREATE TABLE IF NOT EXISTS Model (
    Id INT AUTO_INCREMENT NOT NULL,
    Name VARCHAR(30)      NOT NULL,

    PRIMARY KEY (Id)
);

INSERT INTO Model VALUES (1, "normal");

CREATE TABLE IF NOT EXISTS CharacterR (
    Id   INT AUTO_INCREMENT NOT NULL,
    Name VARCHAR(30),

    PRIMARY KEY (Id)
);

INSERT INTO CharacterR VALUES (1, "villager1");
INSERT INTO CharacterR VALUES (2, "villager2");
INSERT INTO CharacterR VALUES (3, "villager3");
INSERT INTO CharacterR VALUES (4, "villager4");
INSERT INTO CharacterR VALUES (5, "wolf1");
INSERT INTO CharacterR VALUES (6, "wolf2");
INSERT INTO CharacterR VALUES (7, "wolf3");
INSERT INTO CharacterR VALUES (8, "wolf4");
INSERT INTO CharacterR VALUES (9, "seer");
INSERT INTO CharacterR VALUES (10, "witch");
INSERT INTO CharacterR VALUES (11, "hunter");
INSERT INTO CharacterR VALUES (12, "idiot");

CREATE TABLE IF NOT EXISTS ModelCharacter (
    Id          INT AUTO_INCREMENT NOT NULL,
    ModelId     INT NOT NULL,
    CharacterId INT NOT NULL,

    PRIMARY KEY (Id),
    FOREIGN KEY (ModelId) REFERENCES Model(id),
    FOREIGN KEY (CharacterId) REFERENCES CharacterR(id)
);

INSERT INTO ModelCharacter (ModelId, CharacterId) VALUES (1, 1);
INSERT INTO ModelCharacter (ModelId, CharacterId) VALUES (1, 2);
INSERT INTO ModelCharacter (ModelId, CharacterId) VALUES (1, 3);
INSERT INTO ModelCharacter (ModelId, CharacterId) VALUES (1, 4);
INSERT INTO ModelCharacter (ModelId, CharacterId) VALUES (1, 5);
INSERT INTO ModelCharacter (ModelId, CharacterId) VALUES (1, 6);
INSERT INTO ModelCharacter (ModelId, CharacterId) VALUES (1, 7);
INSERT INTO ModelCharacter (ModelId, CharacterId) VALUES (1, 8);
INSERT INTO ModelCharacter (ModelId, CharacterId) VALUES (1, 9);
INSERT INTO ModelCharacter (ModelId, CharacterId) VALUES (1, 10);
INSERT INTO ModelCharacter (ModelId, CharacterId) VALUES (1, 11);
INSERT INTO ModelCharacter (ModelId, CharacterId) VALUES (1, 12);


CREATE TABLE IF NOT EXISTS User (
    Id           INT AUTO_INCREMENT NOT NULL,
    UserName     VARCHAR(20)  UNIQUE,
    PasswordHash VARCHAR(200),

    PRIMARY KEY (Id)
);

CREATE TABLE IF NOT EXISTS Game (
    Id      INT AUTO_INCREMENT NOT NULL,
    Name    VARCHAR(20),
    ModelId INT,
    Date    DATETIME,

    PRIMARY KEY (Id),
    FOREIGN KEY (ModelId) REFERENCES Model(id)
);

CREATE TABLE IF NOT EXISTS UserStatistics (
    Id          INT AUTO_INCREMENT NOT NULL,
    UserId    INT,
    GameId      INT,
    Date        DATETIME,
    SeatId      INT,
    CharacterId INT,
    -- Win or Lose
    Result      VARCHAR(5),

    PRIMARY KEY (Id),
    FOREIGN KEY (UserId) REFERENCES User(id),
    FOREIGN KEY (CharacterId) REFERENCES CharacterR(id),
    FOREIGN KEY (GameId) REFERENCES Game(id)
);

-- There can be only 12 players in one game 
CREATE TABLE IF NOT EXISTS Player (
    Id           INT AUTO_INCREMENT NOT NULL,
    UserId       INT,
    GameId       INT,
    -- 1 to 12
    SeatId       INT,
    CharacterId  INT,

    -- alive or dead
    Status       VARCHAR(5),
    -- win or lose
    Result       VARCHAR(5),


    PRIMARY KEY (Id),
    FOREIGN KEY (UserId) REFERENCES User(id),
    FOREIGN KEY (CharacterId) REFERENCES CharacterR(id),
    FOREIGN KEY (GameId) REFERENCES Game(id)
);