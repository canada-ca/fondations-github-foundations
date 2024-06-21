# GitHub Foundations

[français](#fondations-github)

## Securing Your GitHub Environment at Scale

### Quick Start: [Get Started Now!](./README.md#getting-started)

Managing GitHub for multiple groups introduces complex security and consistency challenges. Misaligned permission levels, team structures, incomplete change rollouts across groups, and status reporting are just a few of the obstacles that can manifest with manual processes.  The GitHub Foundations Toolkit offers a secure and efficient way to manage your organization's GitHub environment through automation and centralized control.

## Features:

### Automate Secure Infrastructure throughout your organization using CI/CD:
- Implement DevSecOps using **Terraform and Terragrunt** to apply Infrastructure as Code (IaC) principles in managing multiple groups under a single GitHub Enterprise account.
- Establish and enforce security best practices by default.
- Employ drift detection to promptly identify and rectify unauthorized changes, guaranteeing continuous security in your configurations.

### Centralize Control:
- Gain a comprehensive overview for managing groups, repositories, and teams across your organization.
- Streamline updates and security policy enforcement across your entire organization, reducing the need for manual intervention.
- Simplify administration by consistently and efficiently managing teams and their memberships across your organization.
- Enhance your security with branch protection through customizable **[Rulesets](https://docs.github.com/en/repositories/configuring-branches-and-merges-in-your-repository/managing-rulesets/about-rulesets#about-rulesets)**.

### Rapid Rollouts:
- Push policy updates across the entire organization in moments rather than waiting days for each group to schedule individual exercises.
- Streamline your workflow by automating the creation, deletion, and configuration of both private and public repositories.

### Reduce Security Risks:
- Employ GitHub Advanced Security (GHAS) for auto code and dependency scanning on public repositories or private repositories where GHAS has been purchased.
- Enforce consistent security policies across all groups to minimize vulnerabilities and protect against attacks.
- Protect your data by securely storing and managing sensitive secrets directly in your GitHub environment.
- Have a unified view of potential vulnerabilities to prevent gaps opening in one group or another.

### Accelerate your ITSG-33 Controls
- Out of the box, GitHub Foundations has a scaffolding to assist with meeting controls in the following areas:
  - Authentication
  - Cryptographic Protection
  - Information Monitoring and Protection
  - Incident Monitoring and Response
  - Configuration Change Control
  - Security Attributes

## Getting Started:
The repository is organized into two layers for ease of setup and management:

### [Bootstrap Layer](./bootstrap/README.md)

Initial setup of your state file backend, and creation of all organizations under your GitHub Enterprise account.

### [Organizations Layer](./organizations/README.md)

Management of organizations, repositories, and teams, Review results of drift detection, and execution of pull request plans for your organizations.

#### Included Tools:
- **Drift Detection:** Detects when someone makes a change to configuration, outside of the source-controlled configuration. Gives the ability to reapply the correct state.
- **Deletion Protection:** When a PR change requests resources be deleted, this tool forces the user to confirm the action
- **GitHub Advanced Security (GHAS) checks:** Checks the state of GHAS for the repos that have it enabled. Reports all of the GHAS scans in one report.
- **Assessment tool:** Used to assess the readiness of your repo, before importing it with the toolkit. Can be used to check whether toolkit guardrails are already in place in the repo.
- **Import tool:** Import repos not currently managed by the toolkit.


## How to Contribute

See [CONTRIBUTING.md](./CONTRIBUTING.md)

## License

Unless otherwise noted, the source code of this project is distributed under the [MIT License](./LICENSE.md).

The Canada wordmark and related graphics associated with this distribution are protected under trademark law and copyright law. No permission is granted to use them outside the parameters of the Government of Canada's corporate identity program. For more information, see [Federal identity requirements](https://www.canada.ca/en/treasury-board-secretariat/topics/government-communications/federal-identity-requirements.html).

---
---

# Fondations GitHub

[English](#github-foundations)

## Sécurité de votre environnement GitHub à grande échelle

### Démarrage rapide [Bien démarrer maintenant!](#bien-démarrer)

La gestion manuelle de GitHub pour plusieurs groupes à la fois présente des défis complexes de sécurité et d’uniformité. Les problèmes d’alignement entre les niveaux d’autorisation et les structures des équipes, le déploiement incomplet des changements dans les groupes et les rapports d’état n’en sont que quelques exemples. La trousse GitHub Foundations permet une gestion efficace et sécurisée de votre organisation par l’automatisation et la centralisation dans l’environnement GitHub.

## Fonctionnalités

### Automatisation de l’infrastructure sécurisée dans votre organisation au moyen de l’intégration continue et de la livraison continue

* Mettez en œuvre l’approche DevSecOps (développement, sécurité et opérations) à l’aide des outils **Terraform et Terragrunt** pour appliquer les principes de l’infrastructure en tant que code (IaC) à la gestion de multiples groupes au moyen d’un seul compte GitHub Enterprise.
* Déterminez les pratiques exemplaires à appliquer par défaut en matière de sécurité.
* Décelez les changements non autorisés au moyen de la détection de dérive et
rectifiezles pour assurer la sécurité continue de vos configurations.

### Contrôle centralisé

* Gérez les groupes, les dépôts et les équipes dans votre organisation au moyen d’une vue d’ensemble complète.
* Rationalisez les mises à jour et l’application des politiques de sécurité dans l’ensemble de votre organisation, afin de réduire la nécessité d’avoir recours à des interventions manuelles.
* Simplifiez l’administration de votre organisation par la gestion uniforme et efficace des équipes et des personnes qui en font partie.
* Utilisez les **[ensembles de règles](https://docs.github.com/en/repositories/configuring-branches-and-merges-in-your-repository/managing-rulesets/about-rulesets#about-rulesets)** personnalisés pour améliorer votre sécurité au moyen des branches protégées.

### Déploiements accélérés

* Transmettez rapidement à l’ensemble de l’organisation les mises à jour des politiques, sans attendre plusieurs jours que chaque groupe planifie l’application des changements.
* Automatisez la création, la suppression et la configuration des dépôts privés et publics pour rationaliser votre flux de travail.

### Réduction des risques liés à la sécurité

* Lorsqu’elle a été achetée, utilisez la fonctionnalité GHAS (GitHub Advanced Security) afin d’analyser automatiquement les codes et les dépendances des dépôts publics ou privés.
* Appliquez uniformément les politiques sur la sécurité à l’ensemble des groupes pour limiter les vulnérabilités et vous protéger contre les attaques.
* Entreposez et gérez en toute sécurité, directement dans l’environnement GitHub, vos documents de nature secrète et délicate pour protéger vos données.
* Obtenez une vue globale des vulnérabilités potentielles pour éviter que des lacunes voient le jour dans les groupes.

### Accélération de la conformité aux contrôles du document ITSG-33

Prête à l’emploi, la trousse GitHub Foundations permet d’étayer la conformité aux contrôles dans les domaines suivants :
* l’authentification;
* la protection cryptographique;
* la surveillance et la protection de l’information;
* la surveillance des incidents et les interventions connexes;
* le contrôle des modifications à la configuration;
* les attributs de sécurité.

## Bien démarrer

Le dépôt est réparti en deux couches pour en faciliter l’installation et la gestion.

### [Couche Bootstrap](./bootstrap/README.md)

Installation initiale du programme dorsal de vos fichiers d’état et création des organisations dans votre compte GitHub Enterprise.

### [Couche Organizations](./organizations/README.md)

* Gestion des organisations, des dépôts et des équipes.
* Examen des résultats liés à la détection de dérive et à l’exécution des plans de demande de tirage dans vos organisations.


#### Outils compris

* **Détection de dérive :** pour déceler les changements apportés à la configuration audelà de celle contrôlée par la source et restaurer l’état qui convient.
* **Protection contre la suppression :** pour confirmer la suppression des ressources si celle-ci répond à un changement dans une demande de tirage.
* **Vérifications de la fonctionnalité GHAS :** pour vérifier l’état de la fonctionnalité GHAS dans les dépôts où elle est activée et faire état des analyses GHAS dans un seul rapport.
* **Outil d’évaluation :** pour évaluer la préparation de votre dépôt avant son importation avec la trousse et possiblement vérifier si les mesures de protection de la trousse sont déjà mises en place dans le dépôt.
* **Outil d’importation :** pour importer des dépôts qui ne sont pas actuellement gérés par la trousse.

## Comment apporter sa contribution

Voir [CONTRIBUTING.md](./CONTRIBUTING.md)

## License

Sauf indication contraire, le code source de ce projet est distribué sous la [licence MIT](./LICENCE.md).

Le mot-symbole « Canada » et les éléments graphiques liés à cette distribution sont protégés en vertu des lois portant sur les marques de commerce et le droit d’auteur. Aucune autorisation n’est accordée pour leur utilisation à l’extérieur des paramètres de coordination du Programme fédéral de l’image de marque du gouvernement du Canada. Pour obtenir davantage de renseignements à ce sujet, veuillez consulter les [Exigences pour l’image de marque](https://www.canada.ca/fr/secretariat-conseil-tresor/sujets/communications-gouvernementales/exigences-image-marque.html).
