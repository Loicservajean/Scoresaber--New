# BeatScorer

BeatScorer est un site web réalisé en Go, HTML et CSS.  
Il utilise l’API ScoreSaber pour afficher des informations sur les joueurs et les maps du jeu Beat Saber.

---

## Objectif du projet

L’objectif est d’exploiter une API REST pour récupérer des données et les manipuler :

- Une recherche de maps et de joueurs  
- Des filtres cumulables  
- Un système de pagination  
- Une gestion de favoris  
- Une page de détails (profil joueur)  
- Une page À‑propos décrivant le déroulement du projet  

---

## API utilisée : ScoreSaber API

Documentation officielle : https://docs.scoresaber.com/

L’API ScoreSaber fournit des données au format json sur :

- les profils des joueurs  
- leurs scores récents  
- les leaderboards (maps)  

---

## Endpoints exploités
### 1. Profil joueur  
`GET https://scoresaber.com/api/player/{id}/full`  
→ Récupère les informations complètes d’un joueur.

### 2. Scores récents  
`GET https://scoresaber.com/api/player/{id}/scores?limit=10&sort=recent`  
→ Récupère les scores récents d’un joueur.

### 3. Recherche de maps  
`GET https://scoresaber.com/api/leaderboards?search={mapName}`  
→ Recherche une map par nom ou mot‑clé.

---

## Fonctionnalités (FT1 → FT4)
### FT1 — Recherche  
Recherche de maps sur plusieurs propriétés (nom, auteur, difficulté).  
Recherche de joueurs par nom ou ID.

### FT2 — Filtres  
Filtres cumulables sur les maps :  
- difficulté  
- mode  
- statut (ranked / qualified / loved)

### FT3 — Pagination  
Affichage des maps par pages (14 par page).

### FT4 — Favoris  
Ajout / suppression de favoris.  
sauvegarder dans un fichier JSON.

---
