# Protect user accounts and identities

([Back](../../GUARDRAILS.md))

## Objective

Protect user accounts and identities.

This section contains the Guardrails that address controls in the following areas:

- Access Control (AC)
- Awareness and Training (AT)
- Audit and Accountability (AU)
- Identification and Authentication (IA)
- System and Services Acquisition (SA)
- System Information and Integrity (SI)

## Mandatory Requirements

| Activity  | Validation |
| --- | --- |
| Implement multi-factor authentication (MFA) for all user accounts. Use phishing resistant MFA where and when available. <br><br> **Note:** User accounts and identities include: <ul><li>root or global administrator (as it has enhanced or the highest level of privilege over the control plane and addresses all aspects of access control).</li><li>Other Administrative user accounts. Refer to<ul><li>Section 4 of the [Directive on Service and Digital](https://www.tbs-sct.canada.ca/pol/doc-eng.aspx?id=32601)</li><li>[Appendix G: Standard on Enterprise Information Technology Service Common Configurations](https://www.tbs-sct.canada.ca/pol/doc-eng.aspx?id=32713)</li><li>[Account Management Configuration Requirements](https://www.canada.ca/en/government/system/digital-government/policies-standards/enterprise-it-service-common-configurations/account.html)</li></ul></li><li>Regular user accounts</li></ul> | <ul><li>Confirm that MFA is implemented according to GC guidance through screenshots, compliance reports, or compliance checks enabled through a reporting tool for all user accounts.</li><li>Confirm that digital policies are in place to ensure that MFA configurations are enforced.</li></li><li>Confirm and report the count of registered root or global administrators (you should have at least two and no more than five).</li></ul> |
| Use separate dedicated accounts for highly privileged roles (for example, domain administrators, global administrators, and root and any domain administrator equivalent access) when administering cloud services to minimize the potential damage. | <ul><li>Provide evidence that there are dedicated user accounts for administration (for example, privileged access).</li> |
| Disable Inactive Accounts (automated) | <ul><li>The system must disable user accounts after **90 days** of inactivity</li></ul> |
| Implement a timeout and a lockout after **3 consecutive invalid logon attempts** by a user during a **15 minute** period for all user accounts. | <ul><li>The system automatically locks the account/node for a **30-minute** period or until released by an administrator when the maximum number of unsuccessful attempts is exceeded.</li><li>Confirm that mechanisms, such as throttling, account lock out policies, monitoring and risk-based authentication, to protect against password brute force attacks have been implemented.</li><li>Provide evidence that secure session management is implemented, including automatic logouts after periods of inactivity and re-authentication for sensitive operations.</li></ul>|
| Enforce the default password policy in accordance with GC [Password Guidance](https://www.canada.ca/en/government/system/digital-government/online-security-privacy/password-guidance.html). | <ul><li>Demonstrate that password policy for the SCM platform has been configured according to the Password Guidance by:<ul><li>requiring passwords that are at least 12 characters long without a maximum length limit</li><li>countering online guessing or brute force of passwords using throttling, account lockout policies, monitoring and multi-factor authentication</li><li>protecting against offline attacks using effective hashing, salting and keyed hashing.</li></ul></li></ul> |
| Ensure that users only have the minimum permissions necessary to perform their tasks | <ul><li>Confirm that users only have the minimum permissions necessary to perform their tasks.</li><li>Confirm that permissions are reviewed and updated regularly.</li></ul> |

## Conditional Requirements

| Activity  | Validation |
| --- | --- |
| Implement a secure process for account recovery to prevent it from being exploited. | <ul><li>Confirm that the process for account recovery is secure and cannot be exploited.</li></ul> |
| If your SCM supports API keys, ensure they are securely managed, regularly rotated, and not shared. | <ul><li>Confirm that API keys are securely managed, regularly rotated, and not shared.</li></ul> |

## References

- [Direction on the Secure Use of Commercial Cloud Services: Security Policy Implementation Notice](https://www.canada.ca/en/treasury-board-secretariat/services/access-information-privacy/security-identity-management/direction-secure-use-commercial-cloud-services-spin.html) (SPIN) 2017-01, subsection 6.2.3
- Cyber Centre’s top 10 IT security actions, number 3
- [Recommendations for Two-Factor User Authentication Within the Government of Canada Enterprise Domain](https://intranet.canada.ca/wg-tg/rtua-rafu-eng.asp) (accessible only on the Government of Canada network)
- [Government of Canada Multi-Factor Authentication (MFA) Considerations and Strategy for GC Enterprise IT Services](https://www.gcpedia.gc.ca/gcwiki/images/9/9e/GC_MFA_Strategy.pdf)
- [Directive on Service and Digital, Appendix G: Standard on Enterprise Information Technology Service Common Configurations](https://www.tbs-sct.canada.ca/pol/doc-eng.aspx?id=32713)
- [Account Management Configuration Requirements](https://www.canada.ca/en/government/system/digital-government/policies-standards/enterprise-it-service-common-configurations/account.html)
- [Event Logging Guidance](https://www.canada.ca/en/government/system/digital-government/online-security-privacy/event-logging-guidance.html)
- [Guidance on Defence in Depth for Cloud-Based Services (ITSP.50.104)](https://cyber.gc.ca/en/guidance/itsp50104-guidance-defence-depth-cloud-based-services), subsection 4.6

## Related security controls from ITSG-33

AC-2, AC-2(7), AC-5, AC-6, AC-6(1), AC-6(2), AC-6(5), AC-19,
AT-3,
AU-9(4),
IA-2, IA-2(6), IA-5, IA-5(6),
SA-4(12),
SI-4(20)
