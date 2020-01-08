# Comment utiliser l'app de remplissage ?
Rdv sur http://cards.jackmarshall.fr/editor

## Créer une nouvelle carte
1. Bien vérifier que la carte n'existe pas. J'insiste les doublons ca va être galère.
2. Mettre le nom anglais complet (ex: Fyanna, Torment of Everblight) de la nouvelle carte dans "new ref english name" et appuyer sur créer.
3. Voir la section suivante pour le remplissage

## Modifier une carte existante
1. Dans le menu du haut: sélectionner la faction, le type de figurine et la figurine puis cliquer sur GO ! (Attention Warlock != Warcaster...)
2. Remplir les différents onglets avec les infos en Français.
3.  Bien vérifier que les données présentes sont à jour. Elles viennent de whac qui n'est plus mis à jour depuis le début de MK3

### Ref
Type => Les tags présents sous le nom du model Blighted Nyss Unit par exemple

### Model
Bien penser à ajouter la taille du socle dans l'onglet Model

Vous pouvez cliquer sur le nom du model pour éditer ses caractéristiques et ses armes.

Saisie des PV:
- PV des Warbeasts => 8-7-11 (mind-body-spirit)
- PV des Warjacks => Pas de solution pratique imo, on va reprendre de qu'a fait Schalf pour limiter le nombre de modif à faire.
Ca sera une chaine de 36 caractères comme la grille d'un jack remplie ligne par ligne avec:
	- x => absence de case
	- . => case vide
	- L,R,M,C,.. => système
	exemple:
	```
	x....x
	......
	......
	.L..R.
	LLMCRR
	xMMCCx
	```
	devient
	```
	x....x.............L..R.LLMCRRxMMCCx
	```
	- PV des jacks scyrah avec shield => idem que pour les jacks avec le nombre de point de bouclier avant exemple: `8-x....x.............L..R.LLMCRRxMMCCx`
	- PV des colossal: 2 grilles de warjack séparée par un `/`
	`x....x.............L..R.LLMCRRxMMCCx/x....x.............L..R.LLMCRRxMMCCx`

### Abilities
Lors de la création d'une nouvelle ability la case "English name" propose un système d'auto-completion si l'ability existe deja.

La textbox de description de la capacité est "intelligente" si vous devez traduire le nom d'un avantage avec icone ou ajouter une reference à une autre ability utilisez "#" ca ouvrira un menu avec tous les skills et les formatera de sorte à ce que le générateur de PDF soit capable de reconnaitre le lien. Dans ce cas la pas besoin d'ajouter la description de la capacité liée, ca le fera tout seul.

Exemple avec un lien vers une capacité
```
Maréchal [Reposition 3"]
Les figurines du battlegroup de cette figurine gagnent #123:Reposition 3"#.
```
==> Ce qui donnera sur le PDF
```
Maréchal [Reposition 3"]: Les figurines du battlegroup de cette figurine gagnent #123:Reposition 3"#. (Reposition 3": Cette figurine peut bouger de 3"...)
```

Exemple avec un lien vers un avantage avec icone
```
Occultation
La figurine / unité amie ciblée gagne :stealth:.
```
==> Ce qui donnera sur le PDF
```
Occultation: La figurine / unité amie ciblée gagne furtivité (+icone).
```

### Spells / Animus
Idem que Abilities pour l'autocomplétion du nom et la textbox "intelligente".

Pour les éléments autres que le nom et la description utiliser de l'anglais. Pour la durée : Up = Up et pas Up = Ent

Les spells ne concernent que les casters, les juniors et les animus. Les spells de la BFS par exemple sont des abilities.

### Feat
Idem que Abilities pour l'autocomplétion du nom et la textbox "intelligente".

## Attention :
Il n'y a pas de sauvegarde automatique: il faut appuyer sur le bouton "Save" sous le menu à droite.

La base de donnée a été pré-alimentée avec les données de whac qui n'est plus mise à jour, par conséquent bien verifier que les stats (PV inclus), les abilities et les spells sont à jour. Ca simplifiera grandement le boulot des relecteurs

