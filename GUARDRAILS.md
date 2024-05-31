# Government of Canada Source Code Management (SCM) Guardrails

## Introduction

The Government of Canada Source Code Management (SCM) Guardrails describe a preliminary set of baseline cyber security controls to ensure that the Source Code Management service environment has a minimum set of configurations. Departments must implement, validate and report on compliance with the guardrails in the first 30 business days of getting access to their SCM account.

Source Code Management services are used to store, manage, and track changes to source code, as well as performing DevOps activities (such as automated deployments, continuous integration, and continuous delivery). The SCM guardrails are designed to ensure that the SCM environment is secure and that the source code is protected from unauthorized access, modification, or deletion.

In order to reduce complexity and avoid repetition, controls and specifications have been listed under the table that is the best fit. As such, the guardrails are intended to be taken in as a whole and not by examining select tables separate from others. Departments are responsible for implementing the minimum configurations identified in the below tables.

## Definitions

For this document the following definitions will be used:

- **Mandatory Requirements:** A set of baseline security controls that departments must implement, validate and report on in the first 30 business days of getting access to their SCM account.
- **Conditional Requirements:** Additional security controls that are taken into consideration for a subset of instances. While these controls may not apply to all solutions, they should be taken into account under specified conditions.

## Applicable Service Models
This guardrail document relates to Software as a Service (SaaS), specifically DevOps and Source Code management tools.

## SCM Guardrails

| ID. | SCM Guardrails |
| --- | --- |
| 01  | [Protect User Accounts and Identities](./guardrails/EN/01_Protect-user-accounts-and-identities.md) |
| 02  | [Manage Access](./guardrails/EN/02_Manage-Role-Access.md) |
| 03  | [Secure Endpoints](./guardrails/EN/03_Secure-Endpoints.md) |
| 04  | [Enterprise monitoring accounts](./guardrails/EN/04_Enterprise-Monitoring-Accounts.md) |
| 05  | [Data Protection](./guardrails/EN/05_Data-Protection.md) |
| 06  | [Network security services](./guardrails/EN/06_Network-Security-Services.md) |
| 07  | [Cyber defense services](./guardrails/EN/07_Cyber-Defense-Services.md) |
| 08  | [Logging and monitoring](./guardrails/EN/08_Logging-and-Monitoring.md) |


## After the first 30 business days

Implementing the guardrails is one of the first steps to establishing a secure SCM platform. After the controls established in the first 30 days, departments should be prepared to monitor their solutions and respond to threats, including keeping up to date on patches and updates. By adhering to these guardrails, departments will have a head start on many controls outlined in the below documents. It is expected that they will continue to work towards completing all of the following:

- [Direction on the Secure Use of Commercial Cloud Services: Security Policy Implementation Notice](https://www.canada.ca/en/treasury-board-secretariat/services/access-information-privacy/security-identity-management/direction-secure-use-commercial-cloud-services-spin.html) (SPIN) 2017-01
- [Government of Canada Security Control Profile for Cloud-Based GC Services](https://www.canada.ca/en/government/system/digital-government/modern-emerging-technologies/cloud-services/government-canada-security-control-profile-cloud-based-it-services.html)
- [GC Cloud PBMM Security Control Profile](https://www.canada.ca/en/government/system/digital-government/digital-government-innovations/cloud-services/government-canada-security-control-profile-cloud-based-it-services.html#toc4)

Departments should engage with their IT security risk management teams to obtain advice and guidance on integrating security assessment and authorization activities as part of the implementation of the SCM platform The [Government of Canada Cloud Security Risk Management Approach and Procedures](https://www.canada.ca/en/government/system/digital-government/digital-government-innovations/cloud-services/cloud-security-risk-management-approach-procedures.html) outlines activities for departments to consider as part of risk management.


## How to Contribute

See [CONTRIBUTING.md](./CONTRIBUTING.md)

## License

Unless otherwise noted, the source code of this project is distributed under the [MIT License](./LICENSE.md).

The Canada wordmark and related graphics associated with this distribution are protected under trademark law and copyright law. No permission is granted to use them outside the parameters of the Government of Canada's corporate identity program. For more information, see [Federal identity requirements](https://www.canada.ca/en/treasury-board-secretariat/topics/government-communications/federal-identity-requirements.html).

---
