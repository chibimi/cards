
CREATE DATABASE cards_db;
USE cards_db;

CREATE TABLE cards (
    id int unsigned not null auto_increment, 
    main_card_id int GENERATED ALWAYS AS (data->"$.main_card_id") VIRTUAL,
    faction_id int GENERATED ALWAYS AS (data->"$.faction_id") VIRTUAL,
    category_id int GENERATED ALWAYS AS (data->"$.category_id") VIRTUAL,
    status varchar(10) GENERATED ALWAYS AS (data->"$.status") VIRTUAL,
    data json,
    PRIMARY KEY (id)
);

INSERT INTO cards (data) VALUES ('{"id": 0, "faction_id":11, "category_id":5, "name": "Blighted Nyss Swordmen", "properties":"Unité de la légion", "models_cnt":"6", "models_max":"10", "cost":"9", "cost_max":"15", "fa":"2", "status":"wip"}');
INSERT INTO cards (data) VALUES ('{"id": 0, "main_card_id":"1", "faction_id":11, "category_id":5, "name": "Blighted Nyss Legionaires", "properties":"Unité de la légion", "models_cnt":"6", "models_max":"10", "cost":"9", "cost_max":"15", "fa":"2", "status":"wip"}');
INSERT INTO cards (data) VALUES ('{"id": 0, "faction_id":11, "category_id":5, "name": "Hellmouth", "properties":"Unité de la légion", "models_cnt":"4", "cost":"6", "fa":"2", "status":"done"}');
INSERT INTO cards (data) VALUES ('{"id": 0, "faction_id":11, "category_id":5, "name": "Blackfrost Shard", "properties":"Unité de nyss corrompus de la légion", "models_cnt":"3", "cost":"9", "fa":"C", "status":"wip"}');
UPDATE cards SET data = json_set(data, '$.id', id);


CREATE TABLE models (
    id int unsigned not null auto_increment, 
    card_id int GENERATED ALWAYS AS (data->"$.card_id") VIRTUAL,
    data json, 
    PRIMARY KEY (id)
);

-- INSERT INTO models (data) VALUES ('{"card_id":1, "name":"", "spd":"", "str":"", "mat":"", "rat":"", "def":"", "arm":"", "cmd":"", "magic_ability":"", "damage":"", "resource":"", "threshold":"", "base_size":"", "advantages": []}');
INSERT INTO models (data) VALUES ('{"card_id":1, "name":"Chef & recrues", "spd":"5", "str":"6", "mat":"6", "rat":"4", "def":"12", "arm":"15", "cmd":"8", "magic_ability":"", "damage":"", "resource":"", "threshold":"", "base_size":"30", "advantages": ["cma"]}');
INSERT INTO models (data) VALUES ('{"card_id":2, "name":"Chef & recrues", "spd":"7", "str":"7", "mat":"7", "rat":"4", "def":"14", "arm":"13", "cmd":"7", "magic_ability":"", "damage":"", "resource":"", "threshold":"", "base_size":"30", "advantages": []}');
INSERT INTO models (data) VALUES ('{"card_id":3, "name":"Bouche", "spd":"-", "str":"10", "mat":"6", "rat":"-", "def":"10", "arm":"18", "cmd":"8", "magic_ability":"", "damage":"", "resource":"", "threshold":"", "base_size":"50", "advantages": ["officer","soulless"]}');
INSERT INTO models (data) VALUES ('{"card_id":3, "name":"Tentacule", "spd":"5", "str":"6", "mat":"6", "rat":"-", "def":"10", "arm":"15", "cmd":"-", "magic_ability":"", "damage":"", "resource":"", "threshold":"", "base_size":"40", "advantages": ["pathfinder","soulless"]}');
INSERT INTO models (data) VALUES ('{"card_id":4, "name":"Sevryn1", "spd":"6", "str":"5", "mat":"7", "rat":"4", "def":"14", "arm":"12", "cmd":"6", "magic_ability":"7", "damage":"8", "resource":"", "threshold":"", "base_size":"30", "advantages": ["immunity_frost","officer","pathfinder"]}');
INSERT INTO models (data) VALUES ('{"card_id":4, "name":"Rhylyss1", "spd":"6", "str":"5", "mat":"7", "rat":"4", "def":"14", "arm":"12", "cmd":"6", "magic_ability":"7", "damage":"8", "resource":"", "threshold":"", "base_size":"30", "advantages": ["immunity_frost","pathfinder"]}');
INSERT INTO models (data) VALUES ('{"card_id":4, "name":"Vysarr1", "spd":"6", "str":"5", "mat":"7", "rat":"4", "def":"14", "arm":"12", "cmd":"6", "magic_ability":"7", "damage":"8", "resource":"", "threshold":"", "base_size":"30", "advantages": ["immunity_frost","pathfinder"]}');
UPDATE models SET data = json_set(data, '$.id', id);


CREATE TABLE weapons (
    id int unsigned not null auto_increment, 
    model_id int GENERATED ALWAYS AS (data->"$.model_id") VIRTUAL,
    type int GENERATED ALWAYS AS (data->"$.type") VIRTUAL,
    data json, 
    PRIMARY KEY (id)
);
INSERT INTO weapons (data) VALUES ('{"model_id":1, "type":"1", "name":"Claymore Nyss", "rng":"1", "pow":"4", "rof":"", "aoe":"", "loc":"", "cnt":"", "advantages":["weaponmaster"]}');
INSERT INTO weapons (data) VALUES ('{"model_id":3, "type":"1", "name":"Bouche", "rng":"2", "pow":"5", "rof":"", "aoe":"", "loc":"", "cnt":"", "advantages":[]}');
INSERT INTO weapons (data) VALUES ('{"model_id":4, "type":"1", "name":"Frappe tentaculaire", "rng":"2", "pow":"4", "rof":"", "aoe":"", "loc":"", "cnt":"", "advantages":[]}');
INSERT INTO weapons (data) VALUES ('{"model_id":5, "type":"1", "name":"Claymore Nyss", "rng":"1", "pow":"4", "rof":"", "aoe":"", "loc":"", "cnt":"", "advantages":["weaponmaster"]}');
INSERT INTO weapons (data) VALUES ('{"model_id":6, "type":"1", "name":"Claymore Nyss", "rng":"1", "pow":"4", "rof":"", "aoe":"", "loc":"", "cnt":"", "advantages":["weaponmaster"]}');
INSERT INTO weapons (data) VALUES ('{"model_id":7, "type":"1", "name":"Claymore Nyss", "rng":"1", "pow":"4", "rof":"", "aoe":"", "loc":"", "cnt":"", "advantages":["weaponmaster"]}');
UPDATE weapons SET data = json_set(data, '$.id', id);


CREATE TABLE abilities (
    id int unsigned not null auto_increment, 
    magical int GENERATED ALWAYS AS (data->"$.magical") VIRTUAL,
    data json, 
    PRIMARY KEY (id)
);

INSERT INTO abilities (data) VALUES ('{"original_name":"Ice bolt", "name":"Eclair gelé (*attaque)", "magical":true, "description":"Eclair Gelé est une attaque magique de POR 10. La figurine touchée subit un jet de dégats de froid PUI 12. Sur une touche critique la figurine touchée devient stationnaire à moins qu''elle ait Immunité Froid."}');
INSERT INTO abilities (data) VALUES ('{"original_name":"Ice cage", "name":"Prison gelée (*attaque)", "magical":true, "description":"Prison Gelée est une attaque magique de POR 10. La figurine touchée subit une malus cumulatif de -2 DEF pour un tour à moins qu''elle ait Immunité Froid. Quand une figurine sans imunité froid est touchée 3 fois ou plus par Prison Gelée pendant le même tour, elle devient stationnaire."}');
INSERT INTO abilities (data) VALUES ('{"original_name":"kiss of Lyliss", "name":"Baiser de Lyliss (*attaque)", "magical":true, "description":"Baiser de Lyliss est une attaque magique de POR 10. Quand une figurine de faction alliée fait un jet de dégats contre une figurine/unité touchée par Baiser de Lyliss, ajouter +2 au résultat des dés. Baiser de Lyliss dure 1 tour."}');
INSERT INTO abilities (data) VALUES ('{"original_name":"Cloak of mist", "name":"Manteau Brumeux (*action)", "magical":true, "description":"Les figurines de cette unité gagne dissimulation. Les figurines hors formation ne sont pas affectées. Manteau Brumeux dure 1 round."}');
INSERT INTO abilities (data) VALUES ('{"original_name":"Disbinding", "name":"Détachement (*action)", "magical":true, "description":"Les sorts à entretien ennemi et les animi sur les figurines de cette unité expirent immédiatement."}');
INSERT INTO abilities (data) VALUES ('{"original_name":"Vengeance", "name":"Vengeance", "magical":false, "description":"Pendant la phase de maintenance, si une figurine ou plus de cette unité a été endommagé par une attaque ennemie lors du round précédent, chaque modèle dans cette unité peut avancer de 3 et faire une attaque de mêlée basique."}');
INSERT INTO abilities (data) VALUES ('{"original_name":"wall of steel", "name":"Mur d''Acier", "magical":false, "description":"Tant que la figurine est SàS avec une figurine ou plus de cette unité, elle gagne +2 ARM."}');
INSERT INTO abilities (data) VALUES ('{"original_name":"consume", "name":"Consumer", "magical":false, "description":"Si cette attaque touche une figurine de petite base non-warlock, non-warcaster, la figurine touchée est retirée du jeu."}');
INSERT INTO abilities (data) VALUES ('{"original_name":"back attack", "name":"Attaque arrière", "magical":false, "description":"Cette figurine peut cibler des modèles dans son arc arrière lorsqu''elle déclare une attaque avec cette arme et sa porté de mêlée n''est pas limité à son arc avant avec cette arme."}');
INSERT INTO abilities (data) VALUES ('{"original_name":"grip", "name":"Etreinte", "magical":false, "description":"Si cette arme touche une figurine ennemie sur une base large ou plus petite, imédiatement après la résolution de l''attaque cette figurine peut être retirée du jeu. Quand c''est le cas, la figurine touchée est poussée directement vers la Bouche jusqu''à ce qu''elle contacte une figurine, un obstacle ou une obstruction. Après avoir bougé la figurine, la Bouche peut immédiatement faire une attaque de mêlée basique ciblant celle-ci."}');
INSERT INTO abilities (data) VALUES ('{"original_name":"mercenary", "name":"Mercenaire", "magical":false, "description":"Cette figurine travaille pour Legion, Cryx"}');
UPDATE abilities SET data = json_set(data, '$.id', id);


CREATE TABLE card_ability (
    card_id int not null,
    ability_id int not null,
    PRIMARY KEY (card_id, ability_id)
);

INSERT INTO card_ability VALUES (2, 11);

CREATE TABLE model_ability (
    model_id int not null,
    ability_id int not null,
    PRIMARY KEY (model_id, ability_id)
);

INSERT INTO model_ability VALUES(1, 6);
INSERT INTO model_ability VALUES(1, 7);
INSERT INTO model_ability VALUES(5, 1);
INSERT INTO model_ability VALUES(5, 2);
INSERT INTO model_ability VALUES(5, 3);
INSERT INTO model_ability VALUES(5, 8);
INSERT INTO model_ability VALUES(6, 1);
INSERT INTO model_ability VALUES(6, 2);
INSERT INTO model_ability VALUES(6, 4);
INSERT INTO model_ability VALUES(6, 8);
INSERT INTO model_ability VALUES(7, 1);
INSERT INTO model_ability VALUES(7, 2);
INSERT INTO model_ability VALUES(7, 5);
INSERT INTO model_ability VALUES(7, 8);


CREATE TABLE weapon_ability (
    weapon_id int not null,
    ability_id int not null,
    PRIMARY KEY (weapon_id, ability_id)
);

INSERT INTO weapon_ability VALUES(2, 8);
INSERT INTO weapon_ability VALUES(2, 9);
INSERT INTO weapon_ability VALUES(3, 10);

CREATE TABLE spells (
    id int unsigned not null auto_increment, 
    data json,
    PRIMARY KEY (id)
);

INSERT INTO spells (data) VALUES ('{"original_name":"Admonition", "name":"Sommation", "cost":"2", "rng":"6", "aoe":"-", "pow":"-", "dur":"UP", "off":"NO", "description":"Super dodge gotcha!"}');
UPDATE spells SET data = json_set(data, '$.id', id);

CREATE TABLE card_spell (
    card_id int not null,
    spell_id int not null,
    PRIMARY KEY (card_id, spell_id)
);

INSERT INTO card_spell VALUES (9, 1);

CREATE TABLE feats (
    id int unsigned not null auto_increment, 
    card_id int GENERATED ALWAYS AS (data->"$.card_id") VIRTUAL,
    data json,
    PRIMARY KEY (id)
);
