# Logging and Monitoring

([Back](../../GUARDRAILS.md))

## Objective

Enable logging for the SCM environment and for SCM-based workloads.
This section contains the Guardrails that address controls in the following areas:

- Access Control (AC)
- Configuration Management (CM)
- Identification and Authentication (IA)
- System Information and Integrity (SI)


## Mandatory Requirements

| Activity | Validation |
| --- | --- |
| Implement adequate level of logging and reporting, including a security audit log function in all information systems. | <ul><li>Confirm policy for event logging is implemented.</li><li>Confirm that the following logs are included: <ul><li>Sign-in logs (all users, interactive and non-interactive sign-ins, API sign-ins)</li><li>Access privilege and group changes (including group membership and group privilege assignment)</li><li>Changes in configuration of the SCM tool</li><li>SCM resource provisioning activities</li><li>Actions/runners logs</li><li>Network access logs</li><li>Webhooks logs</li><li>API logs</li></ul></li></ul> |
| Ensure that the appropriate contact information is configured so that the SCM service provider can notify the GC organization of incidents they detect. | <ul><li>Confirm that the security contact record within the account should be completed with the details of at least two appropriate information security personnel (if multiple personnel are permitted by the SCM platform).</li></ul> |
| Enable Secret-scanning and vulnerability scanning. | <ul><li>Turn on automatic secret scanning and vulnerability scanning.</li><li>Implement a solution to monitor for scans requiring action.</li><li>Demonstrate that scan reports are available for review along with scan logs.</li></ul> |
| Enable dependency scanning. | <ul><li>Confirm that automatic dependency scanning is enabled.</li><li>Implement a solution to monitor for scans requiring action.</li><li>Demonstrate that scan reports are available for review along with scan logs.</li></ul> |
| Ensure that resources are assigned to monitor SCM-based events | <ul><li>Demonstrate that the monitoring use cases for the SCM platform have been implemented and have been integrated with the overall security monitoring activities being performed by the department (evidence could include monitoring a checklist or a system generated report).</li></ul> |
| Audit the use of privileged functions | <ul><li>Confirm that auditing of the use of privileged functions is enabled for all user accounts.</li></ul> |


## Conditional Requirements
None

## References

- [Direction on the Secure Use of Commercial Cloud Services: Security Policy Implementation Notice](https://www.canada.ca/en/treasury-board-secretariat/services/access-information-privacy/security-identity-management/direction-secure-use-commercial-cloud-services-spin.html) (SPIN) 2017-01, subsection 6.3.1
- Cyber Centre’s top 10 IT security actions, numbers 1, 5 and 8
- [Event Logging Guidance](https://www.canada.ca/en/government/system/digital-government/online-security-privacy/event-logging-guidance.html)
- [Guidance on Defence in Depth for Cloud-Based Services (ITSP.50.104)](https://cyber.gc.ca/en/guidance/itsp50104-guidance-defence-depth-cloud-based-services), subsection 4.8

## Related security controls from ITSG-33

AC-22, AU-2, AU-6, IA-5(6), IA-5(7)
