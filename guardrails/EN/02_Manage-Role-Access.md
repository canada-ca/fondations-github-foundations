# Manage Role Access

([Back](../../GUARDRAILS.md))

## Objective

Establish access control policies and procedures for management of all accounts at group or role-based levels.

This section contains the Guardrails that address controls in the following areas:

- Access Control (AC)
- Awareness and Training (AT)
- Identification and Authentication (IA)
- Planning (PL)
- Personnel Security (PS)
- System and Communications Protection (SC)
- System and Information Integrity (SI)


## Mandatory Requirements

| Activity | Validation |
| --- | --- |
| Implement a mechanism to enforce access authorizations for all user accounts, based on criteria in the Directive on Service and Digital, Appendix G: Standard on Enterprise Information Technology Service Common Configurations, and in [section 3 of the Account Management Configuration Requirements](https://www.canada.ca/en/government/system/digital-government/policies-standards/enterprise-it-service-common-configurations/account.html#cha3) |<ul><li>Demonstrate access configurations and policies are implemented for different classes of users (non-privileged, and privileged users).</li><li>Confirm that the access authorization mechanisms have been implemented to: <ul><li>Uniquely identify and authenticate users to the cloud service</li> <li>Validate that the least privilege role is assigned</li> <li>Validate that Role Based Access is implemented</li> <li>Terminate role assignment upon job change or termination</li> <li>Perform periodic reviews of role assignment (minimum yearly)</li> <li>Disable default and dormant accounts</li> <li>Avoid using generic user accounts.</li> </ul></li><li>Verify that a review of role assignment for root or global administrator accounts is performed at least every 12 months.</li></ul> |
| Leverage role-based access and configure for least privilege doing so can include built-in roles or custom roles that have been established with only the minimum number of privileges required to perform the job function. | <ul><li>Demonstrate that built-in roles on the SCM platform are configured for least privilege. Custom roles can be used but a rationale should be documented and approved.</li></ul> |
| Establish a guest user access policy and procedures that minimize the number of guest users and that manage the life cycle of such accounts so that such accounts are terminated when they are no longer needed. <br><br><p>**Note:** a guest is someone who is not an employee, student or member of your organization (a guest does not have an existing account with the organization’s SCM platform).<p> | <ul><li>Confirm that only required guest user accounts are enabled (according to the business requirements of the service)</li><li>Provide a list of non-organizational users with elevated privileges.</li><li>Verify that reviews of guest access are performed periodically.</li></ul> |
| Ensure that any security considerations for users or administrators are documented in the organization's security awareness training program. | <ul><li>Provide evidence that security considerations for users are documented in the organization's security awareness training program.</li></ul> |

## Conditional Requirements

| Activity | Validation |
| --- | --- |
| If your SCM uses service accounts for automation, ensure they are securely managed and have the minimum necessary permissions. | <ul><li>Provide evidence that service accounts are securely managed and have the minimum necessary permissions.</li></ul> |
| Document a process for managing accounts, access privileges, and access credentials for organizational users and non-organizational users (if required), based on criteria listed in the Directive on Service and Digital, [Appendix G: Standard on Enterprise Information Technology Service Common Configurations](https://www.tbs-sct.canada.ca/pol/doc-eng.aspx?id=32713), and in [section 3 of the Account Management Configuration Requirements](https://www.canada.ca/en/government/system/digital-government/policies-standards/enterprise-it-service-common-configurations/account.html#cha3). This process should be approved by the chief security officer (or their delegate) and by the designated official for cyber security. | <ul><li>Confirm that the access control procedure for management of administrative accounts has been documented for the SCM platform. The access control procedure:<ul><li>should include provision for any guest accounts and custom accounts</li><li>must refer to the emergency break glass procedure</li></ul></li></ul> |
| Enforce just-in-time access for privileged user accounts to provide time-based and approval-based role activation to mitigate the risks of excessive, unnecessary or misused access permissions. | <ul><li>Confirm just-in-time access for all privileged user accounts to provide time-based and approval-based role activation.</li></ul> |
| Enforce attribute-based access control to restrict access based on a combination of authentication factors, such as devices issued and managed by the GC, device compliance, sign-in and user risks, and location | <ul><li>Provide evidence that attribute-based access control mechanisms are in place to restrict access based on attributes or signals, such as authentication factors, devices issued and managed by the GC, device compliance, sign-in and user risks, and location...</li></ul> |
| <ul><li>Leverage tools, such as privilege access management systems, to enforce access control to privileged functions by configuring roles that require approval for activation</li><li>Choose one or multiple users or groups as delegated approvers</li></ul> | <ul><li>Provide evidence that all role activation for privileged user accounts require approval, and that privilege elevation is temporary (time-bound).</li></ul> |
| Develop an acceptable use policy, ensure that it is communicated to all users, and ensure that it is enforced. | <ul><li>Provide evidence that an acceptable use policy is in place, communicated to all users, and enforced. Ideally, this policy would be made visible to users of the SCM when using the platform</li></ul> |

## References

- [Direction on the Secure Use of Commercial Cloud Services: Security Policy Implementation Notice](https://www.canada.ca/en/treasury-board-secretariat/services/access-information-privacy/security-identity-management/direction-secure-use-commercial-cloud-services-spin.html), (SPIN) 2017-01, subsection 6.2.3
- Cyber Centre’s top 10 IT security actions, number 3
- [User Authentication Guidance for Information Technology Systems (ITSP.30.031 v3)](https://cyber.gc.ca/en/guidance/user-authentication-guidance-information-technology-systems-itsp30031-v3)
- [Guidance on Cloud Authentication for the Government of Canada](https://intranet.canada.ca/wg-tg/cagc-angc-eng.asp) (accessible only on the Government of Canada network)
- [Recommendations for Two-Factor User Authentication Within the Government of Canada Enterprise Domain](https://intranet.canada.ca/wg-tg/rtua-rafu-eng.asp) (accessible only on the Government of Canada network)
- [Directive on Service and Digital, Appendix G: Standard on Enterprise Information Technology Service Common Configurations](https://www.tbs-sct.canada.ca/pol/doc-eng.aspx?id=32601)
- [Account Management Configuration Requirements](https://www.canada.ca/en/government/system/digital-government/policies-standards/enterprise-it-service-common-configurations/account.html)
- [Guidance on Defence in Depth for Cloud-Based Services (ITSP.50.104)](https://cyber.gc.ca/en/guidance/itsp50104-guidance-defence-depth-cloud-based-services), subsection 4.6
- [Password Guidance](https://www.canada.ca/en/government/system/digital-government/online-security-privacy/password-guidance.html)

## Related security controls from ITSG-33

AC-1, AC‑2, AC‑2(1), AC‑2(7), AC‑2(9), AC‑5, AC‑6, AC‑6(1), AC‑6(2), AC‑6(5), AC-14,
AT-1, AT-2, AT-2(2), AT-3, AT-4
IA-1, IA‑2(6), IA-4, IA-4(2), IA-4(3), IA-4(4), IA‑5, IA‑5(3), IA‑5(4), IA‑5(6),
PL-4,
PS-1, PS-2, PS-3, PS-3(3), PS-4, PS-5, PS-6, PS-7,
SC-2,
SI-4(20)
