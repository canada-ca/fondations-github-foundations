([Back](README.md))

# Compliance

## Table of Contents

* [francais](#table-des-matieres)

* [Compliance](#compliance)
    * [Overview](#overview)
    * [GitHub Foundations Toolkit PBMM Compliance Features](#github-foundations-toolkit-pbmm-compliance-features)
    * [GitHub Security Features Status](#github-security-features-status)
    * [Audit Logs](#audit-logs)
        * [Daily Audit Log Export](#daily-audit-log-export)
        * [Exporting GitHub Enterprise Audit Logs to Azure Sentinel](#exporting-github-enterprise-audit-logs-to-azure-sentinel)
        * [Streaming Audit Logs to Other Platforms](#streaming-audit-logs-to-other-platforms)

## Overview

This document outlines the compliance and security features of the platform.

## GitHub Foundations Toolkit PBMM Compliance Features

The following table maps the features of the GitHub Foundations Toolkit to the [PBMM Controls]()


| Item | What | Where | Controls | Open-Source Alternative |
| --- | --- | --- | --- | --- |
| Encrypted Secrets | Uses GitHub public key to encrypt secrets. Secrets must be encrypted to be used. | <ul><li>modules/organization</li><li>modules/private_repository</li><li>modules/public_repository</li><ul> | IA-5(c)(e)(h)(i), IA-5(1)(c),IA-5(6), IA-5(7), SC-8(1), SC-12, SC-13, SC-17 | <ul><li>[vault](https://www.hashicorp.com/products/vault)</li><li>[confidant](https://lyft.github.io/confidant/)</li> |
| Vulnerability Alerts | <ul><li>Vulnerability alerts are enabled by default on `public` and `private` repos</li><li>Vulnerabilities detection is automatically updated</li></ul> | <ul><li>modules/organization</li><li>modules/private_repository</li><li>modules/public_repository</li> | SI-4(5), SI-4(7), SI-10 | <ul><li>[sbom-tool](https://github.com/microsoft/sbom-tool)</li><li>[syft](https://github.com/anchore/syft)</li><li>[cdxgen](https://github.com/CycloneDX/cdxgen)</li> |
| Secret Scanning | Scanning of the repo for secrets | <ul><li>modules/organization</li><li>modules/private_repository</li><li>modules/public_repository</li><li>modules/repository_base</li><ul> | AC-22, IA-5(7), IR-9, SI-4(5), SI-4(7), SI-10 | <ul><li>[trufflehog](https://trufflesecurity.com/)</li><li>[gitleaks](https://github.com/gitleaks/gitleaks)</li><li>[detect-secrets](https://github.com/Yelp/detect-secrets)</li> |
| Advanced Security | <ul><li>Code scanning</li><li>Manual step. Instructions to be added to README</li></ul> | <ul><li>modules/private_repository</li><li>modules/public_repository</li><li>modules/repository_base</li><ul> | SI-4(5), SI-4(7), SI-10 | <ul><li>[semgrep](https://github.com/semgrep/semgrep)</li><li>[sonarqube](https://www.sonarqube.org/)</li> |
| Protected Branches Ruleset - Pull Requests | <ul><li>Require pull requests</li><li>Require at least 1 reviewer</li><li>When new commits are pushed to an existing PR, any previous approvals are required again.</li></ul> | <ul><li>modules/private_repository</li><li>modules/public_repository</li><li>modules/repository_base</li><ul> | CM-3, CM-4, CM-5, SI-10, SI-12 | <ul><li>[GitLab](https://about.gitlab.com/)</li> |
| Protected Branches Ruleset - Signed Commits<br><br><br>**TODO - Not currently validated** | <ul><li>Require signed commits</li></ul> | <ul><li>modules/private_repository</li><li>modules/public_repository</li><li>modules/repository_base</li><ul> | IA-2, IA-2(11), IA-8, IA-8(100), SC-8, SC-8(1), SC-13  | <ul><li>[GitLab](https://about.gitlab.com/)</li><li>[Git](https://git-scm.com/)</li><li>[Mercurial](https://www.mercurial-scm.org/)</li> |
| Export Audit Material<br><br><br> | Audit material is exported in [JSON format](#daily-audit-log-export), or there are instructions on how to obtain logs by other means. [See here](#audit-logs). | GH Action | AC-2(4), AC-6(9), AC-17(1), AU-2, AU-6, SI-4 | <ul><li>[GitLab](https://about.gitlab.com/)</li> |
| Delete branches on merge | Branches are configured to be deleted after a PR is merged | <ul><li>modules/private_repository</li><li>modules/public_repository</li><li>modules/repository_base | SI-12 | <ul><li>[GitLab](https://about.gitlab.com/)</li> |
| Repository Creation Restrictions | Users can: <ul><li>Create private repos</li><li>Create internal repos</li></ul> User cannot: <ul><li>Create public repos</li><li>Fork Private repos</li></ul> | <ul><li>modules/organization</li> | AC-20(3), AC-22 | <ul><li>[GitLab](https://about.gitlab.com/)</li> |
| Predefined Roles | Predefined roles include: <ul><li>Security Engineer</li><li>Contractor</li><li>Community Manager</li></ul> | <ul><li>modules/organization</li> | AC-2, AC-16(2) | <ul><li>[GitLab](https://about.gitlab.com/)</li><li>[Keycloak](https://www.keycloak.org/)</li> |
| Drift Detection | Tool used to detect if the terraform state has drifted from what's stored in source control | <ul><li>modules/organization</li> | CM-2, CM-3(f)(g), CM-5, CM-6(c)(d), CM-9(d) | <ul><li>[Terraform CLI](https://www.terraform.io/)</li><li>[Terragrunt CLI](https://terragrunt.gruntwork.io/)</li> |
| Resource Deletion Protection | An action that forces a user to acknowledge Terraform plan deletions, before performing them | <ul><li>modules/organization</li> | AC-16(2), CM-3, CM-4, CM-5, CM-6(D), CM-9, SI-10 | <ul><li>[GitLab](https://docs.gitlab.com/runner/)</li> |
| Detect whether GHAS enabled | For public repositories, and repos with GHAS purchased, we recommend that it be turned on. This GH Action runs daily at 2am to check that the setting is enabled in all eligible repos | GH Action | IA-5(7), SI-4(5), SI-4(7), SI-10 | <ul><li>[GitLab runners](https://docs.gitlab.com/runner/)</li> |
| Documentation - Creating resources | The documentation for the GHF toolkit includes the relevant documentation that describes authentication methods for users signing into your enterprise, how to create organizations and teams for repository access and collaboration, and suggested best practices for user security | READMEs | AC-5, AC-6 | <ul><li>Any text editor / code-revision control tool.</li> |
| | | | | |

### Controls that are met by GitHub by default

The following controls are met by GitHub by default and are not explicitly implemented in the toolkit:

| Item | What | Where | Controls | Open-Source Alternative |
| --- | --- | --- | --- | --- |
| Account dormancy policy | GitHub accounts are marked dormant, and made inactive after 90 days of inactivity for Enterprise accounts | [Managing Dormant Accounts](https://docs.github.com/en/enterprise-cloud@latest/admin/managing-accounts-and-repositories/managing-users-in-your-enterprise/managing-dormant-users) | AC-2(3) | <ul><li>[GitLab](https://about.gitlab.com/)</li><li>[Keycloak](https://www.keycloak.org/)</li> |
| HTTPS and SSH access | GitHub enforces the use of HTTPS and/or SSH for committing and pulling code | [GitHub Docs](https://docs.github.com/en/authentication/securing-your-account-with-two-factor-authentication-2fa) | AC-17(2), SC-8, SC-8(1), SC-12, SC-13, SC-17 | <ul><li>[GitLab](https://about.gitlab.com/)</li><li>[Keycloak](https://www.keycloak.org/)</li> |
| | | | | |


## GitHub Security Features Status

Inside the toolkit, we strive to provide the most up-to-date security features that GitHub has to offer. Below is a table that shows the status of the features that we support in the toolkit.
1. The _**Status**_ column is the status of implementing the GitHub feature in the toolkit.
2. The **_Provider Status**_ column is the status of the feature in [Terraform GitHub Provider](https://registry.terraform.io/providers/integrations/github/latest/docs), and
3. the _**GH API Status**_ column is the status of the feature in the [GitHub API](https://docs.github.com/en/rest/reference).

| Feature | Status | Provider Status | GH API Status | Side Notes |
| --- | --- | --- | --- | --- |
| Encrypted secrets | ✅ | ✅ | ✅ | We don't handle encryption directly, but secrets are encrypted with github private keys |
| Vulnerability alerts | ✅ | ✅ | ✅ | |
| Private vulnerability reporting | ❌ | ❌ | ✅ | |
| Secret scanning | ✅ | ✅ | ✅ | |
| Secret scanning push protection | ✅ | ✅ | ✅ | |
| CodeQL codescanning | ❌ | ❌ | ✅ | |
|  |  |  |  |  |
| Rulesets | ✅ | ✅ | ✅ | Although the [GitHub Documentation](https://docs.github.com/en/rest/repos/rules?apiVersion=2022-11-28#create-a-repository-ruleset) lists it as a capability in the api there are some limitations. For example: Organization rulesets can define workflows that must be completed for the ruleset to be considered passed, but it is not possible for repository rulesets. |
| Commit signing enforcement | ✅ | ✅ | ✅ | |
| Delete branches on merge | ✅ | ✅ | ✅ | |
| Repository creation restrictions | ✅ | ✅ | ✅ | |
| | | | | |
| Custom Repository Roles | ✅ | ✅ | ✅ | |
| Custom Organization Roles | ❌ | ❌ | ✅ | Confusing because there is a terraform resource named `github_organization_custom_role` but this resource actually makes custom repository roles. Custom organization roles have not been implemented. But the GH api does support it. |
| Deploy Keys | ❌ | ✅ | ✅ | |
| Organization Member Base Permissions | ✅ | ✅ | ✅ | |
| Custom Properties | ❌ | ❌ | ❌ | |
| | | | | |
| 2FA | ❌ | ❌ | ❌ | Not exposed by the API. This setting is completely manual |
| SAML SSO | ❌ | ❌ | ❌ | Not exposed by the API. This setting is completely manual |
| Team Synchronization | ❌ | ✅ | ✅ | |
| | | | | |



## Audit Logs

Learn how to enable audit logs for your GitHub Enterprise account to track user activity and changes made to your organization.

1. For _GitHub Enterprise_ Audit logs, see [GitHub's documentation](https://docs.github.com/en/enterprise-cloud@latest/admin/monitoring-activity-in-your-enterprise/reviewing-audit-logs-for-your-enterprise/using-the-audit-log-api-for-your-enterprise).
2. To query the audit log API for your _GitHub Enterprise_, see [GitHub's documentation](https://docs.github.com/en/enterprise-cloud@latest/rest/enterprise-admin/audit-log?apiVersion=2022-11-28).
3. For _Organization_ Audit logs, see [GitHub's documentation](https://docs.github.com/en/organizations/keeping-your-organization-secure/reviewing-the-audit-log-for-your-organization).

### Daily Audit Log Export

A GitHub Workflow is run daily to export the audit logs to a JSON file. The JSON file is then uploaded to the action's output artifacts. The JSON file can be downloaded from the artifacts tab in the GitHub Actions page.

### Exporting GitHub Enterprise Audit Logs to Azure Sentinel

To export GitHub Enterprise audit logs to Azure Sentinel, [follow the steps found here](https://learn.microsoft.com/en-us/azure/sentinel/data-connectors/azure-logic-apps).

The Azure Marketplace offers the [Microsoft Sentinel - Continuous Threat Monitoring for GitHub (Preview)](https://azuremarketplace.microsoft.com/en-us/marketplace/apps/microsoftcorporation1622712991604.sentinel4github?tab=Overview) connector to help you get started.

### Streaming Audit Logs to Other Platforms

GitHub supports the [streaming of audit logs](https://docs.github.com/en/enterprise-cloud@latest/admin/monitoring-activity-in-your-enterprise/reviewing-audit-logs-for-your-enterprise/streaming-the-audit-log-for-your-enterprise) to other platforms. The list of currently supported platforms is:

* [Amazon S3](https://docs.github.com/en/enterprise-cloud@latest/admin/monitoring-activity-in-your-enterprise/reviewing-audit-logs-for-your-enterprise/streaming-the-audit-log-for-your-enterprise#setting-up-streaming-to-amazon-s3)
* [Azure Blob Storage](https://docs.github.com/en/enterprise-cloud@latest/admin/monitoring-activity-in-your-enterprise/reviewing-audit-logs-for-your-enterprise/streaming-the-audit-log-for-your-enterprise#setting-up-streaming-to-azure-blob-storage)
* [Azure Event Hubs](https://docs.github.com/en/enterprise-cloud@latest/admin/monitoring-activity-in-your-enterprise/reviewing-audit-logs-for-your-enterprise/streaming-the-audit-log-for-your-enterprise#setting-up-streaming-to-azure-event-hubs)
* [Datadog](https://docs.github.com/en/enterprise-cloud@latest/admin/monitoring-activity-in-your-enterprise/reviewing-audit-logs-for-your-enterprise/streaming-the-audit-log-for-your-enterprise#setting-up-streaming-to-datadog)
* [Google Cloud Storage](https://docs.github.com/en/enterprise-cloud@latest/admin/monitoring-activity-in-your-enterprise/reviewing-audit-logs-for-your-enterprise/streaming-the-audit-log-for-your-enterprise#setting-up-streaming-to-google-cloud-storage)
* [Splunk](https://docs.github.com/en/enterprise-cloud@latest/admin/monitoring-activity-in-your-enterprise/reviewing-audit-logs-for-your-enterprise/streaming-the-audit-log-for-your-enterprise#setting-up-streaming-to-splunk)


---
([retour](./README.md))

# Conformité

## table des matières

* [English](#table-of-contents)

* [Conformité](#conformité)
    * [Aperçu](#aperçu)
    * [Fonctionnalités de conformité PBMM de la trousse d’outils essentiels GitHub](#fonctionnalités-de-conformité-pbmm-de-la-trousse-doutils-essentiels-github)
    * [Statut des fonctionnalités de sécurité de GitHub](#statut-des-fonctionnalités-de-sécurité-de-github)
    * [Journaux d’audit](#journaux-daudit)
        * [Exportation du journal d’audit quotidien](#exportation-du-journal-daudit-quotidien)
        * [Exportation des journaux d’audit de GitHub Enterprise vers Azure Sentinel](#exportation-des-journaux-daudit-de-github-enterprise-vers-azure-sentinel)
        * [Diffusion des journaux d’audit vers d’autres plateformes](#diffusion-des-journaux-daudit-vers-dautres-plateformes)

## Aperçu

Ce document décrit les caractéristiques de conformité et de sécurité de la plateforme.

# Fonctionnalités de conformité PBMM de la trousse d’outils essentiels GitHub

Le tableau suivant met en correspondance les fonctionnalités de la trousse d’outils essentiels de
GitHub avec les [Contrôles PBMM]()

| Article | Quoi | Où | Contrôles | Autre solution à source ouverte |
| --- | --- | --- | --- | --- |
| Secrets chiffrés | Utilise la clé publique de GitHub pour chiffrer les secrets. Les secrets doivent être chiffrés pour être utilisés | <ul><li>modules/organization</li><li>modules/private_repository</li><li>modules/public_repository</li><ul> | IA-5(c)(e)(h)(i), IA-5(1)(c),IA-5(6), IA-5(7), SC-8(1), SC-12, SC-13, SC-17 | <ul><li>[vault](https://www.hashicorp.com/products/vault)</li><li>[confidant](https://lyft.github.io/confidant/)</li> |
| Alertes de vulnérabilité | <ul><li>Les alertes de vulnérabilité sont activées par défaut sur les dépôts `publics` et `privés`</li><li>La détection des vulnérabilités est automatiquement mise à jour</li></ul> | <ul><li>modules/organization</li><li>modules/private_repository</li><li>modules/public_repository</li> | SI-4(5), SI-4(7), SI-10 | <ul><li>[sbom-tool](https://github.com/microsoft/sbom-tool)</li><li>[syft](https://github.com/anchore/syft)</li><li>[cdxgen](https://github.com/CycloneDX/cdxgen)</li> |
| Analyse de secrets | Analyse de secrets dans le référentiel de code | <ul><li>modules/organization</li><li>modules/private_repository</li><li>modules/public_repository</li><li>modules/repository_base</li><ul> | AC-22, IA-5(7), IR-9, SI-4(5), SI-4(7), SI-10 | <ul><li>[trufflehog](https://trufflesecurity.com/)</li><li>[gitleaks](https://github.com/gitleaks/gitleaks)</li><li>[detect-secrets](https://github.com/Yelp/detect-secrets)</li> |
| Sécurité avancée | <ul><li>Analyse du code</li><li>Étape manuelle. Instructions à ajouter au fichier README</li></ul> | <ul><li>modules/private_repository</li><li>modules/public_repository</li><li>modules/repository_base</li><ul> | SI-4(5), SI-4(7), SI-10 | <ul><li>[semgrep](https://github.com/semgrep/semgrep)</li><li>[sonarqube](https://www.sonarqube.org/)</li> |
| Ensemble de règles des branches protégées - Demande de tirage | <ul><li>Exiger des demandes de tirage</li><li>Exiger au moins un évaluateur</li><li>Lorsque de nouvelles validations sont poussées vers un DT existant, toutes les approbations précédentes sont à nouveau requises</li></ul> | <ul><li>modules/private_repository</li><li>modules/public_repository</li><li>modules/repository_base</li><ul> | CM-3, CM-4, CM-5, SI-10, SI-12 | <ul><li>[GitLab](https://about.gitlab.com/)</li> |
| Ensemble de règles des branches protégées - Validations signées<br><br><br>**TODO - Pas encore validé** | <ul><li>Exiger des validations signées</li></ul> | <ul><li>modules/private_repository</li><li>modules/public_repository</li><li>modules/repository_base</li><ul> | IA-2, IA-2(11), IA-8, IA-8(100), SC-8, SC-8(1), SC-13 | <ul><li>[GitLab](https://about.gitlab.com/)</li><li>[Git](https://git-scm.com/)</li><li>[Mercurial](https://www.mercurial-scm.org/)</li> |
| Exportation du matériel d’audit<br><br><br> | Les données d’audit sont exportées au [format JSON](#exportation-du-journal-daudit-quotidien), ou des instructions sont fournies sur la manière d’obtenir les journaux par d’autres moyens. [Voir ici](#journaux-daudit). | GH Action | AC-2(4), AC-6(9), AC-17(1), AU-2, AU-6, SI-4 | <ul><li>[GitLab](https://about.gitlab.com/)</li> |
| Supprimer des branches lors de la fusion | Les branches sont configurées pour être supprimées après la fusion d’un DT | <ul><li>modules/private_repository</li><li>modules/public_repository</li><li>modules/repository_base | SI-12 | <ul><li>[GitLab](https://about.gitlab.com/)</li> |
| Restrictions à la création de dépôts | Les utilisateurs peuvent: <ul><li>Créer des dépôts privés</li><li>Créer des dépôts internes</li></ul>Les utilisateurs ne peuvent pas: <ul><li>Créer des dépôts publics</li><li>Diviser des dépôts privés</li></ul> | <ul><li>modules/organization</li> | AC-20(3), AC-22 | <ul><li>[GitLab](https://about.gitlab.com/)</li> |
| Rôles prédéfinis | Les rôles prédéfinis sont les suivants: <ul><li>Ingénieur en sécurité</li><li>Entrepreneur</li><li>Gestionnaire de communauté</li></ul> | <ul><li>modules/organization</li> | AC-2, AC-16(2) | <ul><li>[GitLab](https://about.gitlab.com/)</li><li>[Keycloak](https://www.keycloak.org/)</li> |
| Détection de dérive | Outil utilisé pour détecter si l’état de terraform s’est écarté de ce qui est stocké dans le contrôle de la source | <ul><li>modules/organization</li> | CM-2, CM-3(f)(g), CM-5, CM-6(c)(d), CM-9(d) | <ul><li>[Terraform CLI](https://www.terraform.io/)</li><li>[Terragrunt CLI](https://terragrunt.gruntwork.io/)</li> |
| Protection contre la suppression de ressources | Une action qui oblige l’utilisateur à reconnaître les suppressions de plans Terraform avant de les effectuer | <ul><li>modules/organization</li> | AC-16(2), CM-3, CM-4, CM-5, CM-6(D), CM-9, SI-10 | <ul><li>[GitLab](https://docs.gitlab.com/runner/)</li> |
| Détecter si le GHAS est activé | Pour les dépôts publics et les dépôts ayant acheté un GHAS, nous recommandons de l’activer. Cette action de GH est exécutée tous les jours à 2 h pour vérifier que le paramètre est activé dans tous les dépôts admissibles | Action GH | IA-5(7), SI-4(5), SI-4(7), SI-10 | <ul><li>[GitLab runners](https://docs.gitlab.com/runner/)</li> |
| Documentation - Créer des ressources | La documentation de la boîte à outils GHF comprend la documentation pertinente qui décrit les méthodes d’authentification pour les utilisateurs qui s’inscrivent dans votre organisation, la façon de créer des organisations et des équipes pour l’accès au dépôt et la collaboration, et les meilleures pratiques suggérées pour la sécurité des utilisateurs | READMEs | AC-5, AC-6 | <ul><li>Tout éditeur de texte / outil de contrôle de la révision du code.</li> |
| | | | | |

### Contrôles rencontrés par défaut par GitHub

Les contrôles suivants sont respectés par défaut par GitHub et ne sont pas explicitement mis en œuvre dans la boîte à outils:

| Article | Quoi | Où | Contrôles | Autre solution à source ouverte |
| --- | --- | --- | --- | --- |
| Politique en matière de comptes inactifs | Les comptes GitHub sont marqués comme dormants et rendus inactifs après 90 jours d’inactivité pour les comptes Enterprise | [Gestion des utilisateurs dormants](https://docs.github.com/fr/enterprise-cloud@latest/admin/managing-accounts-and-repositories/managing-users-in-your-enterprise/managing-dormant-users) | AC-2(3) | <ul><li>[GitLab](https://about.gitlab.com/)</li><li>[Keycloak](https://www.keycloak.org/)</li> |
| Accès HTTPS et SSH | GitHub impose l’utilisation de HTTPS ou SSH pour la transmission et l’extraction de code | [Documentation GitHub](https://docs.github.com/fr/authentication/securing-your-account-with-two-factor-authentication-2fa) | AC-17(2), SC-8, SC-8(1), SC-12, SC-13, SC-17 | <ul><li>[GitLab](https://about.gitlab.com/)</li><li>[Keycloak](https://www.keycloak.org/)</li> |
| | | | | |

## Statut des fonctionnalités de sécurité de GitHub

Dans la boîte à outils, nous nous efforçons de fournir les fonctionnalités de sécurité les plus récentes que GitHub peut offrir. Vous trouverez ci-dessous un tableau indiquant l’état des fonctionnalités que nous prenons en charge dans la boîte à outils.
1.	La colonne _Statut_ indique l’état d’avancement de la mise en œuvre de la fonctionnalité GitHub dans la boîte à outils.
2.	La colonne _Status de Provider_ correspond à l’état de la fonctionnalité dans [Terraform GitHub Provider](https://registry.terraform.io/providers/integrations/github/latest/docs), et
3.	La colonne _Status de l’API GH_ indique le statut de la fonctionnalité dans [l’API GitHub](https://docs.github.com/fr/rest?apiVersion=2022-11-28).

| Fonctionnalité | Statut | Statut de Provider | Statut de l’API GH | Notes d’accompagnement |
| --- | --- | --- | --- | --- |
| Secrets chiffrés | ✅ | ✅ | ✅ | Nous ne gérons pas le chiffrage directement, mais les secrets sont chiffrés avec les clés privées de github |
| Alertes de vulnérabilité | ✅ | ✅ | ✅ | |
| Signalement des vulnérabilités privés | ❌ | ❌ | ✅ | |
| Analyse de secrets | ✅ | ✅ | ✅ | |
| Protection par poussée de l’analyse de secrets | ✅ | ✅ | ✅ | |
| Analyse des codes CodeQL | ❌ | ❌ | ✅ | |
|  |  |  |  |  |
| Ensemble de règles | ✅ | ✅ | ✅ | Bien que la [documentation GitHub](https://docs.github.com/fr/rest/repos/rules?apiVersion=2022-11-28) le mentionne comme une capacité de l’API, il y a quelques limitations. Par exemple : Les ensembles de règles d’organisation peuvent définir des flux de travail qui doivent être exécutés pour que l’ensemble de règles soit considéré comme réussi, mais ce n’est pas possible pour les ensembles de règles de dépôt. |
| Engager l’exécution de la signature | ✅ | ✅ | ✅ | |
| Supprimer des branches lors de la fusion | ✅ | ✅ | ✅ | |
| Restrictions à la création de dépôts | ✅ | ✅ | ✅ | |
| | | | | |
| Rôles personnalisés du dépôt | ✅ | ✅ | ✅ | |
| Rôles personnalisés d’organisation | ❌ | ❌ | ✅ | Confusion, car il y a une ressource terraform nommée `github_organization_custom_role`, mais cette ressource crée en fait des rôles personnalisés de dépôt. Les rôles personnalisés d’organisation n’ont pas été mis en œuvre. Mais l’API de GH le prend en charge. |
| Clés pour qéployer | ❌ | ✅ | ✅ | |
| Permissions de base des membres de l’organisation | ✅ | ✅ | ✅ | |
| Propriétés personnalisées | ❌ | ❌ | ❌ | |
| | | | | |
| 2FA | ❌ | ❌ | ❌ | Non exposée par l’API. Ce réglage est entièrement manuel |
| SAML SSO | ❌ | ❌ | ❌ | Non exposée par l’API. Ce réglage est entièrement manuel |
| Synchronisation des équipes| ❌ | ✅ | ✅ | |
| | | | | |

## Journaux d’audit

Découvrez comment activer les journaux d’audit pour votre compte GitHub Enterprise afin de suivre l’activité des utilisateurs et les modifications apportées à votre organisation.
1.	Pour les journaux d’audit de GitHub Enterprise, voir la [documentation de GitHub](https://docs.github.com/fr/enterprise-cloud%40latest/admin/monitoring-activity-in-your-enterprise/reviewing-audit-logs-for-your-enterprise/using-the-audit-log-api-for-your-enterprise).
2.	Pour interroger l’API du journal d’audit de votre GitHub Entreprise, consultez la [documentation de GitHub](https://docs.github.com/fr/enterprise-cloud%40latest/rest/enterprise-admin/audit-log?apiVersion=2022-11-28).
3.	Pour les journaux d’audit de l’organisation, voir la [documentation de GitHub](https://docs.github.com/fr/organizations/keeping-your-organization-secure/managing-security-settings-for-your-organization/reviewing-the-audit-log-for-your-organization).


### Exportation du journal d’audit quotidien

Un flux de travail GitHub est exécuté quotidiennement pour exporter les journaux d’audit vers un fichier JSON. Le fichier JSON est ensuite téléchargé dans les artefacts de sortie de l’action. Le fichier JSON peut être téléchargé à partir de l’onglet artefacts de la page GitHub Actions.

### Exportation des journaux d’audit de GitHub Enterprise vers Azure Sentinel

Pour exporter les journaux d’audit de GitHub Enterprise vers Azure Sentinel, [suivez les étapes décrites ici](https://learn.microsoft.com/fr-ca/azure/sentinel/data-connectors/azure-logic-apps).
Azure Marketplace propose le connecteur [Microsoft Sentinel – Continuous Threat Monitoring for GitHub (Preview)](https://azuremarketplace.microsoft.com/fr-ca/marketplace/apps/microsoftcorporation1622712991604.sentinel4github?tab=Overview) pour vous aider à démarrer.

### Diffusion des journaux d’audit vers d’autres plateformes

GitHub prend en charge la [diffusion des journaux d’audit](https://docs.github.com/fr/enterprise-cloud@latest/admin/monitoring-activity-in-your-enterprise/reviewing-audit-logs-for-your-enterprise/streaming-the-audit-log-for-your-enterprise) vers d’autres plateformes. La liste des plateformes actuellement prises en charge est la suivante :

* [Amazon S3](https://docs.github.com/fr/enterprise-cloud@latest/admin/monitoring-activity-in-your-enterprise/reviewing-audit-logs-for-your-enterprise/streaming-the-audit-log-for-your-enterprise#setting-up-streaming-to-amazon-s3)
* [Azure Blob Storage](https://docs.github.com/fr/enterprise-cloud@latest/admin/monitoring-activity-in-your-enterprise/reviewing-audit-logs-for-your-enterprise/streaming-the-audit-log-for-your-enterprise#setting-up-streaming-to-azure-blob-storage)
* [Azure Event Hubs](https://docs.github.com/fr/enterprise-cloud@latest/admin/monitoring-activity-in-your-enterprise/reviewing-audit-logs-for-your-enterprise/streaming-the-audit-log-for-your-enterprise#setting-up-streaming-to-azure-event-hubs)
* [Datadog](https://docs.github.com/fr/enterprise-cloud@latest/admin/monitoring-activity-in-your-enterprise/reviewing-audit-logs-for-your-enterprise/streaming-the-audit-log-for-your-enterprise#setting-up-streaming-to-datadog)
* [Google Cloud Storage](https://docs.github.com/fr/enterprise-cloud@latest/admin/monitoring-activity-in-your-enterprise/reviewing-audit-logs-for-your-enterprise/streaming-the-audit-log-for-your-enterprise#setting-up-streaming-to-google-cloud-storage)
* [Splunk](https://docs.github.com/fr/enterprise-cloud@latest/admin/monitoring-activity-in-your-enterprise/reviewing-audit-logs-for-your-enterprise/streaming-the-audit-log-for-your-enterprise#setting-up-streaming-to-splunk)
