# Description du Projet

Notre projet consiste à proposer une simulation du modèle économique des plateformes de streaming comme Netflix ou Amazon Prime. Nous proposons ainsi d'avoir d'une part des producteurs (les plateformes de streaming) et d'autre part, les consommateurs (les abonnés, clients lambda) qui interagissent entre eux.

## Membres du Projet

|  **Auteurs** |   **Contact**              |
| :------------ | :------------- |
| Bingqian Shu |  bingqian.shu@etu.utc.fr |
| Samar Mellouki |  samar.mellouki@etu.utc.fr |
| Chenxin Li |  chenxin.li@etu.utc.fr |
| Yanan Fu |  yanan.fu@etu.utc.fr |

## Ce qu'il est censé faire

Le projet doit lancer plusieurs agents Viewer et Platforms. Chacun dispose d'une stratégie spécifique vis-à-vis des ressources communes (Vidéos, Abonnements) et adapte son comportement en fonction des interactions et/ou de ses propres facteurs internes.

Du côté des Viewers, il dispose d'un budget fixe, de préférences et d'une stratégie à appliquer. Par exemple, le Viewer "passionné" accorde plus d'importance à ses préférences contrairement au Viewer "économique" qui cherche avant tout à maximiser la rentabilité de ses abonnements.

Du côté des Platforms, elles ont chacune une liste de vidéos à publier chaque mois mais aussi des objectifs différents représenté par leur stratégie. Cela peut consister à maximiser leur nombre d'abonnés ou à améliorer leur catalogue.

Enfin, ces deux entités interagissent entre elles via les abonnements et les vidéos : les Viewers s'abonnent, visionnent, et se désabonnent des Platforms qui elles, adaptent leurs prix, leurs catalogues et leurs abonnements en fonction de ce qu'il se passe.

## Ce que nous sommes censé.e.s observer
Il y a plusieurs situations à différencier en fonction du nombre d'agents que l'on choisit de lancer.

### Le cas du Monopole (n, 1)
Le Monopole consiste à avoir une unique Plateform pour plusieurs Viewers. Dans cette situation, les Viewers qui ne sont pas intéressés par la plateforme ne s'abonneront pas. Ceux qui sont intéressés s'abonneront. 

La Plateforme ayant ici le plus de "pouvoir" pourra augmenter ses prix ou non, en fonction de si ses objectifs sont atteints ou pas - et cela dépend du nombre d'abonnés intéressés. A priori, il devrait y avoir une limite au-delà de laquelle, la plateforme perdra des abonnés car elle a trop augmenté ses prix en essayant de profiter de la situation. 

Tant qu'elle ne s'essouffle pas (au niveau de son catalogue) et conserve des prix raisonnables, on tendra vers une limite stable. On aura donc un nombre quasiment fixe d'abonnés et un prix stable.
### Le cas de l’Oligopole (n, m)
Le second cas consiste à avoir m plateformes, plus ou moins populaires. Chacune a sa stratégie et elles commencent toutes avec une base neutre (aucun abonné). Les n Viewers vont s'abonner au fur et à mesure des mois et définir la popularité de telle ou telle plateforme et nous pourrions arriver à d'une part, un classement des plateformes, et d'autres part, à l'identification des stratégies "dominantes" parmi les Platforms et les Viewers mais aussi, des "mauvaises" stratégies en observant les plateformes qui feront faillite.

Il se pourrait aussi que l'on arrive à un monde plus ou moins équitable. Tout dépendra des préférences des abonnés dans le monde lancé (random) ou peut-être d'autres paramètres qu'il sera intéressant d'identifier
### Le cas du Monopsone (1, m)
Le dernier cas consisterait à lancer l’inverse du premier : 1 Viewer et m plateformes. Le marché serait alors à la baisse pour essayer de coller à son budget/ses préférences.

## Problématique
Question posée : _Quels sont les facteurs caractéristiques du rapport de force entre Consommateur et Producteur ?_

Nous souhaitons ainsi identifier les paramètres de ce dernier ainsi que l'impact des stratégies des agents sur des facteurs.

# Installation du Projet
## Prérequis
Usage du package [faker]( https://pkg.go.dev/github.com/go-faker/faker/v4) au sein du projet. Il est donc nécessaire de l’installer :
```
go get -u github.com/go-faker/faker/v4
```
## Lancement
Pour lancer le projet, il faut avoir les ports 8080 et 9000-9005 libres puis run le main :
```
cd cmd
go run main.go
```

# Analyse du Projet
## Variation des paramètres
Les paramètres que nous pouvons faire varier sont nombreux ici. Il est ainsi possible de faire varier le monde en termes de nombres d'agent (Viewers et Platforms) mais aussi de comportement (stratégies) voire de ressources à petite échelle (budget par exemple). Nous n’avons cependant pas eu le temps d’implémenter tout cela et nous ne pouvons faire varier que le nombre d’agents et les quotas de chaque stratégie parmi ce nombre.

A plus grande échelle, nous pouvons aussi faire varier le temps de notre projet car nos observations ne peuvent finalement se faire qu'à un instant t et le choix de cet instant est finalement crucial : s'arrêter trop tôt ne montrera rien de pertinent mais la simulation pourrait n'avoir aucune « fin » car le modèle économique s'auto-alimente (le store s’épuise en revanche puisque notre jeu de données est fixe pour les titres de vidéos).


### Cas du Monopole
Etant donné que nous n’avons qu’une seule stratégie du côté des plateformes, il ne se passe en réalité pas grand-chose : les viewers s’abonnent en fonction de leur préférence et la plateforme se contente d’enrichir son catalogue pour conserver ses clients. Ainsi, elle conserve sa clientèle et tout se passe bien du coup.
Pour que ce cas soit plus intéressant, il aurait fallu implémenter un réajustement des prix : la plateforme détecterai qu’elle a le monopole du marché et tenterait d’augmenter ses prix par exemple de 5% chaque mois jusqu’à perdre des clients, elle s’arrêterait donc et aura trouvé le prix limite à ne pas dépasser.

### Cas de l’Oligopole
Si on lance un monde avec 33% de nos 3 stratégies Viewers (Economique, Riche et Passion) sur 12 mois (99 Viewers et 5 plateformes).

On constate que les plateformes les plus chers ont été délaissés. D’un autre côté, les clients de la plateformes étant « riches », ils ne se désabonneront pas (au contraire, ils suivront la tendance). Les passionnés se retrouvent à différents endroits, entre des « économes » et des « riches » car ils chercheront à s’abonner au meilleur abonnement (car meilleure qualité vidéo) tout en respectant leur budget et leur préférences de thèmes. Finalement, ce sont les Viewers les plus « mixtes » car on les retrouve un peu dans tous les abonnements. Enfin, les économiques se sont concentrés en grande quantité sur les plateformes proposant des prix bas. 

Au niveau des plateformes, elles avaient toutes au départ le même budget initial : 1000. On remarque qu’au niveau de la répartition d’abonnés, les plateformes qui s’adressent à plusieurs profils économiques s’enrichissent plus que les autres (qui ne s’adressent qu’au très riche par exemple).
### Cas du Monopsone
Dans le cas du monopsone, tout dépend du comportement du Viewer. S’il s’agit d’un Viewer Riche (que l’on pourrait associer à une sorte d’investisseur finalement), les offrants aurait des fonds chaque mois et finirait par pouvoir se renouveler. 

Pour ce qui est des Viewer Passion et Economie, il restreigne encore davantage les plateformes car certaines n’auront alors pas de clients. L’idéal aurait été de pouvoir ajouter des clients au fur et à mesure selon un paramètre externe ou un facteur de communication au sein des viewers (système d’entourage, etc.).

Dans tous les cas, le marché n’ayant pas suffisamment de demandes, il sera relativement pauvre et les offrants ne disposant pas de moyens, perdront des abonnés car ils ne sont pas en mesure de se renouveler dans notre simulation.

## Points positifs
Pour ce qui est des points positifs, nous pourrions citer l’architecture de notre projet qui est relativement propre. Chaque dossier représente une ressource (strategy, viewer, plateform, website, utils, demo&cmd) et on s’y retrouve facilement. On a également essayé d’injecter nos dépendances (notamment vis-à-vis des stratégies) ce qui rend très facile la programmation de nouvelles stratégies. 

Aussi, nous avons a priori réussi à gérer la concurrence sur nos ressources via synx.Mutex et synx.RWMutex. Le catalogue par exemple ne devait pas être exclusif à 1 lecteur (plusieurs autres lecteurs autorisés) mais exclusif à 1 rédacteur donc usage de RLock quand le lecteur lit et de Lock quand le rédacteur écrit. Idem avec les autres ressources. La gestion du temps est aussi réalisé à l’aide WaitGroups qui synchronise les mois (mais les agents sont bien lancés en parallèle chaque mois – ils devront juste attendre la fin du mois pour tous avant de débuter le prochain mois).

Aussi, la base du projet est fonctionnelle. Chaque mois, les Viewers applique leur stratégie (Riche, Economique ou Passion) et s’abonnent/se désabonnent des plateformes et payent à la fin du mois. Du côte des plateformes, elles réceptionnent bien toutes ces requêtes et applique leur Stratégie.

## Points négatifs
Le projet est malheureusement incomplet et ne nous permet de déduire des résultats concrets. En effet, les plateformes n’implémentent finalement qu’une seule stratégie consistant à fidéliser leur clientèle en alimentant le catalogue (on remarque lors du lancement du serveur que le catalogue contient plus que les 4 vidéos d’initialisations, ils ont donc bien acheté des vidéos). Il nous manque ainsi une stratégie qui ajuste le prix des abonnements afin de pouvoir déduire quelle est la stratégie qui permet le plus de « succès ». 

Le second point négatif concerne les paramètres pris en compte. En effet, le modèle est assez simpliste et réduit par exemple les préférences d’un Viewer à des thèmes. Idem du côté des plateformes, elles ne gèrent pas toutes les ressources immobilières, les employés, etc.).
