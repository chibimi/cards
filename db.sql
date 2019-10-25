
CREATE DATABASE IF NOT EXISTS cards_db;
USE cards_db;

CREATE TABLE refs (
    id INT unsigned NOT NULL unique auto_increment,
    faction_id int,
    category_id int,
    title text,
    main_card_id int,
    models_cnt text,
    models_max text,
    cost text,
    cost_max text,
    fa text,
    minion_for json,
    mercenary_for json,
    PRIMARY KEY (id)
);

INSERT INTO refs (faction_id, category_id, title, main_card_id, models_cnt, models_max, cost, cost_max, fa) VALUES (11, 5, 'Nyss Swordmen', 0, "6", "10", "9", "15", "2");
INSERT INTO refs (faction_id, category_id, title, main_card_id, models_cnt, models_max, cost, cost_max, fa) VALUES (11, 5, 'Nyss Legionnaires', 0, "6", "10", "9", "15", "2");
INSERT INTO refs (faction_id, category_id, title, main_card_id, models_cnt, models_max, cost, cost_max, fa) VALUES (11, 5, 'Hellmouth', 0, "4", "", "6", "", "2");
INSERT INTO refs (faction_id, category_id, title, main_card_id, models_cnt, models_max, cost, cost_max, fa) VALUES (11, 5, 'BlackFrost Shard', 0, "3", "", "9", "", "C");

CREATE TABLE refs_lang (
    ref_id int unsigned not null,
    lang varchar(2),
    status text,
    name text,
    properties text,
    PRIMARY KEY (ref_id, lang)
);

INSERT INTO refs_lang (ref_id, lang, name, properties, status) VALUES (1, "FR", "Epéeistes Nyss", "Unité de Nyss de la légion", "wip");
INSERT INTO refs_lang (ref_id, lang, name, properties, status) VALUES (3, "FR", "Gueule l'enfer", "Unité de la légion", "done");
INSERT INTO refs_lang (ref_id, lang, name, properties, status) VALUES (3, "UK", "Hellmouth", "Legion Unit", "tbv");
INSERT INTO refs_lang (ref_id, lang, name, properties, status) VALUES (4, "FR", "Shard GlaceNoire", "Unité de Nyss de la légion", "wip");


CREATE TABLE feats (
    ref_id int unsigned not null,
    lang varchar(2),
    name text,
    description text,
    fluff text,
    PRIMARY KEY (ref_id, lang)
);


CREATE TABLE abilities (
    id int unsigned not null auto_increment, 
    title text, 
    PRIMARY KEY (id)
);

CREATE TABLE abilities_lang (
    ability_id int unsigned not null,
    lang varchar(2),
    name text,
    description text,
    PRIMARY KEY (ability_id, lang)
);

CREATE TABLE ref_ability (
    ref_id int not null,
    ability_id int not null,
    type int,
    PRIMARY KEY (ref_id, ability_id)
);

CREATE TABLE model_ability (
    model_id int not null,
    ability_id int not null,
    type int,
    PRIMARY KEY (model_id, ability_id)
);

CREATE TABLE weapon_ability (
    weapon_id int not null,
    ability_id int not null,
    type int,
    PRIMARY KEY (weapon_id, ability_id)
);

CREATE TABLE spells (
    id int unsigned not null auto_increment, 
    title text,
    cost text,
    rng text,
    aoe text,
    pow text,
    dur text,
    off text,
    PRIMARY KEY (id)
);

CREATE TABLE spells_lang (
    spell_id int unsigned not null,
    lang varchar(2),
    name text,
    description text,
    PRIMARY KEY (spell_id, lang)
);

CREATE TABLE ref_spell (
    ref_id int not null,
    spell_id int not null,
    PRIMARY KEY (ref_id, spell_id)
);

CREATE TABLE models (
    id int unsigned not null auto_increment, 
    ref_id int not null,
    spd text,
    str text,
    mat text,
    rat text,
    def text,
    arm text,
    cmd text,
    magic_ability text,
    damage text,
    resource text,
    threshold text,
    base_size text,
    advantages json,
    PRIMARY KEY (id)
);

CREATE TABLE models_lang (
    model_id int unsigned not null,
    lang varchar(2),
    name text,
    PRIMARY KEY (model_id, lang)
);

CREATE TABLE weapons (
    id int unsigned not null auto_increment, 
    model_id int not null,
    type text,
    rng text,
    pow text,
    rof text,
    aoe text,
    loc text,
    cnt text,
    advantages json,
    PRIMARY KEY (id)
);

CREATE TABLE weapons_lang (
    weapon_id int unsigned not null,
    lang varchar(2),
    name text,
    PRIMARY KEY (weapon_id, lang)
);
