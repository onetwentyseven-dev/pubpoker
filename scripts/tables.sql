CREATE TABLE tournaments (
	id varchar(100) not null,
	tournamentID varchar(100) not null,
	venueID varchar(100) not null,
	approved tinyint(1) not null default '0',
	approvedByUserID varchar(100) null,
	createdAt timestamp not null,
	PRIMARY KEY (id)
) COLLATE = 'utf8_unicode_ci' ENGINE=InnoDB;

CREATE TABLE tournament_players (
    id varchar(100) not null,
    tournamentID varchar(100) not null,
    playerID varchar(100) not null,
    finalPosition mediumint unsigned not null default '0',
    `name` varchar(100),
    createdAt timestamp not null,
    PRIMARY KEY (id, tournamentID),
    CONSTRAINT `tournament_foreign_key`  (`tournament_id`)
) COLLATE = 'utf8_unicode_ci' ENGINE=InnoDB;