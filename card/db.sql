CREATE DATABASE cards_db;
USE cards_db;

CREATE TABLE cards (
    id int unsigned not null auto_increment, 
    main_card_id int, 
    faction_id int not null, 
    category_id int not null, 
    name varchar(100) not null, 
    properties varchar(100),
    models_min varchar(50), 
    models varchar(50), 
    cost_min varchar(2), 
    cost varchar(2), 
    fa varchar(2), 
    status varchar(10), 
    fury varchar(2), 
    focus varchar(2), 
    threshold varchar(2), 
    damage varchar(100),
    damage_grid varchar(100),
    damage_spiral varchar(100),
    PRIMARY KEY (id)
);

INSERT INTO cards VALUES (0, 0, 11, 5, "Blighted Nyss Swordmen", "Unité de la légion", "chef & 5 recrues", "chef & 9 recrues", "9", "15", "2", "TODO", "", "", "", "", "", "");
INSERT INTO cards VALUES (0, 0, 11, 5, "Blighted Nyss Legionaires", "Unité de la légion", "chef & 5 recrues", "chef & 9 recrues", "9", "15", "2", "TODO", "", "", "", "", "", "");
INSERT INTO cards VALUES (0, 0, 11, 5, "Hellmouth", "Unité de la légion", "", "bouche & 3 tentacules", "", "6", "2", "TODO", "", "", "", "8", "", "");
INSERT INTO cards VALUES (0, 0, 11, 5, "Blackfrost Shard", "Unité de nyss corrompus de la légion", "", "Sevryn, Rhylyss & Vysarr", "", "9", "C", "TODO", "", "", "", "5-5-5", "", "");

CREATE TABLE models (
    id int unsigned not null auto_increment, 
    card_id int not null, 
    name varchar(100) not null, 
    spd varchar(2), 
    str varchar(2), 
    mat varchar(2), 
    rat varchar(2), 
    def varchar(2), 
    arm varchar(2), 
    cmd varchar(2), 
    magic_ability varchar(2), 
    advantages varchar(255),
    base_size varchar(2),
    PRIMARY KEY (id)
);

INSERT INTO models VALUES (0, 1, "Chef & recrues", "5", "6", "6", "4", "12", "15", "8", "", "cma", "30");
INSERT INTO models VALUES (0, 2, "Chef & recrues", "6", "7", "7", "4", "14", "13", "7", "", "", "30");
INSERT INTO models VALUES (0, 3, "Bouche", "-", "10", "6", "-", "10", "18", "8", "", "officer,soulless", "50");
INSERT INTO models VALUES (0, 3, "Tentacule", "5", "6", "6", "-", "10", "15", "-", "", "pathfinder,soulless", "40");
INSERT INTO models VALUES (0, 4, "Sevryn1", "6", "5", "7", "4", "14", "12", "6", "7", "immunity_frost,officer,pathfinder", "30");
INSERT INTO models VALUES (0, 4, "Rhylyss1", "6", "5", "7", "4", "14", "12", "6", "7", "immunity_frost,pathfinder", "30");
INSERT INTO models VALUES (0, 4, "Vysarr1", "6", "5", "7", "4", "14", "12", "6", "7", "immunity_frost,pathfinder", "30");
-- INSERT INTO models VALUES (0, 1, "", "", "", "", "", "", "", "", "", "", "");

CREATE TABLE abilities (
    id int unsigned not null auto_increment, 
    type int not null, 
    name varchar(100) not null, 
    description text,
    PRIMARY KEY (id)
);

INSERT INTO abilities VALUES (0, 3, "Eclair gelé (*attaque)", "Eclair Gelé est une attaque magique de POR 10. La figurine touchée subit un jet de dégats de froid PUI 12. Sur une touche critique la figurine touchée devient stationnaire à moins qu'elle ait Immunité Froid.");
INSERT INTO abilities VALUES (0, 3, "Prison gelée (*attaque)", "Prison Gelée est une attaque magique de POR 10. La figurine touchée subit une malus cumulatif de -2 DEF pour un tour à moins qu'elle ait Immunité Froid. Quand une figurine sans imunité froid est touchée 3 fois ou plus par Prison Gelée pendant le même tour, elle devient stationnaire.");
INSERT INTO abilities VALUES (0, 3, "Baiser de Lyliss (*attaque)", "Baiser de Lyliss est une attaque magique de POR 10. Quand une figurine de faction alliée fait un jet de dégats contre une figurine/unité touchée par Baiser de Lyliss, ajouter +2 au résultat des dés. Baiser de Lyliss dure 1 tour.");
INSERT INTO abilities VALUES (0, 3, "Manteau Brumeux (*action)", "Les figurines de cette unité gagne dissimulation. Les figurines hors formation ne sont pas affectées. Manteau Brumeux dure 1 round.");
INSERT INTO abilities VALUES (0, 3, "Détachement (*action)", "Les sorts à entretien ennemi et les animi sur les figurines de cette unité expirent immédiatement.");
INSERT INTO abilities VALUES (0, 2, "Vengeance", "Pendant la phase de maintenance, si une figurine ou plus de cette unité a été endommagé par une attaque ennemie lors du round précédent, chaque modèle dans cette unité peut avancer de 3"" et faire une attaque de mêlée basique.");
INSERT INTO abilities VALUES (0, 2, "Mur d'Acier", "Tant que la figurine est SàS avec une figurine ou plus de cette unité, elle gagne +2 ARM.");
INSERT INTO abilities VALUES (0, 4, "Consumer", "Si cette attaque touche une figurine de petite base non-warlock, non-warcaster, la figurine touchée est retirée du jeu.");
INSERT INTO abilities VALUES (0, 4, "Attaque arrière", "Cette figurine peut cibler des modèles dans son arc arrière lorsqu'elle déclare une attaque avec cette arme et sa porté de mêlée n'est pas limité à son arc avant avec cette arme.");
INSERT INTO abilities VALUES (0, 4, "Etreinte", "Si cette arme touche une figurine ennemie sur une base large ou plus petite, imédiatement après la résolution de l'attaque cette figurine peut être retirée du jeu. Quand c'est le cas, la figurine touchée est poussée directement vers la Bouche jusqu'à ce qu'elle contacte une figurine, un obstacle ou une obstruction. Après avoir bougé la figurine, la Bouche peut immédiatement faire une attaque de mêlée basique ciblant celle-ci.");

CREATE TABLE card_ability (
    card_id int not null,
    ability_id int not null,
    PRIMARY KEY (card_id, ability_id)
);

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

INSERT INTO weapon_ability VALUES(1, 9);
INSERT INTO weapon_ability VALUES(1, 10);
INSERT INTO weapon_ability VALUES(3, 11);
INSERT INTO weapon_ability VALUES(4, 11);


CREATE TABLE weapons (
    id int unsigned not null auto_increment, 
    model_id int not null, 
    type int not null, 
    name varchar(100) not null,
    rng varchar(5),
    pow varchar(5),
    rof varchar(5),
    aoe varchar(5),
    loc varchar(5),
    cnt varchar(5),
    advantages varchar(255),
    PRIMARY KEY (id)
);

INSERT INTO weapons VALUES (0, 3, 1, "Gueule", "2", "5", "", "", "", "", "");
INSERT INTO weapons VALUES (0, 3, 1, "Frappe tentaculaire", "2", "4", "", "", "", "", "");
INSERT INTO weapons VALUES (0, 4, 1, "Frappe tentaculaire", "2", "4", "", "", "", "", "");