# Secure Endpoints

([Back](../../GUARDRAILS.md))

## Objective

Implement increased levels of protection for management interfaces.

This section contains the Guardrails that address controls in the following areas:

- Access Control (AC)
- Auditing and Accountability (AU)
- Identification and Authentication (IA)
- System Information and Integrity (SI)

## Mandatory Requirements

| Activity | Validation |
| --- | --- |
| The system Monitors and Controls Remote Access Methods | <ul><li>Set, via configuration, the rules around remote access</li><li>Log and audit all remote access</li></ul> |
| Encryption of Remote Access Sessions | <ul><li>Ensure that all remote access sessions are encrypted</li><li>Ensure that the encryption is in accordance with the GC [Encryption Guidance](https://www.cyber.gc.ca/en/guidance/cryptographic-algorithms-unclassified-protected-protected-b-information-itsp40111)</li></ul> |

## Conditional Requirements

### Self-hosting considerations

| Activity | Validation |
| --- | --- |
| Implement Endpoint Management Configuration Requirements | <ul><li>Implement the [Endpoint Management Configuration Requirements](https://www.canada.ca/en/government/system/digital-government/policies-standards/enterprise-it-service-common-configurations/endpoint.html)</li></ul> |
| Implement antivirus and anti-malware solutions on endpoints that access the SCM. Regularly update these solutions to protect against the latest threats. | <ul><li>Provide evidence that antivirus and anti-malware solutions are implemented on endpoints that access the SCM.</li></ul> |
| Use firewalls to control inbound and outbound network traffic to the SCM. Only allow necessary traffic. | <ul><li>Provide evidence that firewalls are used to control inbound and outbound network traffic to the SCM.</li></ul> |
| Ensure any exposed endpoints are secure. This includes using secure protocols (like HTTPS), validating input, and handling errors securely. | <ul><li>Use secure protocols (like HTTPS), validate input, and handle errors securely.</li></ul> |
|  Regularly update the SCM and any related software to fix known security vulnerabilities. | <ul><li>Provide evidence that the SCM and related software are regularly updated.</li></ul> |
| Isolate the network segment that the SCM resides on to limit the potential impact of a breach.| <ul><li>Provide evidence that the network segment that the SCM resides on is isolated from other network segments.</li></ul> |

## References

- [Direction on the Secure Use of Commercial Cloud Services: Security Policy Implementation Notice](https://www.canada.ca/en/treasury-board-secretariat/services/access-information-privacy/security-identity-management/direction-secure-use-commercial-cloud-services-spin.html) (SPIN) 2017-01, subsection 6.2.3
- Cyber Centre’s top 10 IT security actions, number 2
- [Recommendations for Two-Factor User Authentication Within the Government of Canada Enterprise Domain](https://intranet.canada.ca/wg-tg/rtua-rafu-eng.asp) (accessible only on the Government of Canada network)
- [Directive on Service and Digital, Appendix G: Standard on Enterprise Information Technology Service Common Configurations](https://www.tbs-sct.canada.ca/pol/doc-eng.aspx?id=32713)
- [Endpoint Management Configuration Requirements](https://www.canada.ca/en/government/system/digital-government/policies-standards/enterprise-it-service-common-configurations/endpoint.html)
- [Guidance on Defence in Depth for Cloud-Based Services (ITSP.50.104)](https://cyber.gc.ca/en/guidance/itsp50104-guidance-defence-depth-cloud-based-services), subsection 4.9
- [User Authentication Guidance for Information Technology Systems (ITSP.30.031 v3)](https://cyber.gc.ca/en/guidance/user-authentication-guidance-information-technology-systems-itsp30031-v3)
- [Cryptographic algorithms for UNCLASSIFIED, PROTECTED A, and PROTECTED B information (ITSP.40.111)](https://www.cyber.gc.ca/en/guidance/cryptographic-algorithms-unclassified-protected-protected-b-information-itsp40111)

## Related security controls from ITSG-33

AC3, AC-3(7), AC-4, AC-5, AC-6, AC6(5), AC-6(10), AC-17(1), AC-17(2), AC-19, AC-20(3), AU-6, AU-12, IA-2, IA-2(1), IA-2(11), IA-4, IA-5, IA-5(1), SC-8, SC-8(1), SI-4
