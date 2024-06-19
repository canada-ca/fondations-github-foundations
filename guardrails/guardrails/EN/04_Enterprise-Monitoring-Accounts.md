
# Enterprise Monitoring Accounts

([Back](../../GUARDRAILS.md))

## Objective

Create role-based accounts to enable enterprise monitoring and visibility.

This section contains the Guardrails that address controls in the following contexts:

- Access Control (AC)
- Audit and Accountability (AU)
- Security Assessment and Authorization (CA)
- Identification and Authentication (IA)
- Incident Response (IR)


## Mandatory Requirements

| Activity  | Validation |
| --- | --- |
| Create role-based accounts to enable enterprise monitoring and visibility for cloud environments that are procured via the GC Cloud Broker or are included in the scope of centralized guardrails validation. | <ul><li>Verify that roles required to enable visibility in the GC have been provisioned or assigned.</li></ul> |
| Maintain healthy management of associated roles. | <ul><li>Verify the implementation of controls listed under [Manage Role Access](./02_Manage-Role-Access.md)</li></ul> |
| Have a plan in place for responding to security incidents detected through monitoring, including roles and responsibilities, communication plans, and recovery procedures. | <ul><li>Provide evidence that a plan is in place for responding to security incidents detected through monitoring, including roles and responsibilities, communication plans, and recovery procedures.</li><li>The plan should be included and coordinated with the organization's overall incident response plan</ul> |

## Conditional Requirements

| Activity  | Validation |
| --- | --- |
|  Integrate the SCM with a Security Information and Event Management (SIEM) system if available to centralize log analysis and alerting.| <ul><li>Provide evidence that the SCM is integrated with a SIEM system to centralize log analysis and alerting.</li></ul> |


## References

- [Direction on the Secure Use of Commercial Cloud Services: Security Policy Implementation Notice](https://www.canada.ca/en/treasury-board-secretariat/services/access-information-privacy/security-identity-management/direction-secure-use-commercial-cloud-services-spin.html) (SPIN) 2017-01, subsection 6.2.3
- Cyber Centre’s top 10 IT security actions, number 2
- [IT Security Risk Management: A Lifecycle Approach (ITSG-33), Annex 3A: Security Control Catalogue](https://cyber.gc.ca/en/guidance/it-security-risk-management-lifecycle-approach-itsg-33)

## Related security controls from ITSG-33

AC-1, AC-2, AC-2(1), AC-2(4), AC-2(7), AC-2(9), AC-5, AC-6, AC-6(5),
AU-1, AU-2, AU-2(3), AU-6, AU-6(1), AU-6(3),
CA-7,
IA-1, IA-4, IA-4(2),
IR-1, IR-2, IR-3, IR-3(2), IR-4, IR-4(3), IR-5, IR-6, IR-6(1), IR-7(1), IR-7(2), IR-8
