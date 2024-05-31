# Data Protection

([Back](../../GUARDRAILS.md))

## Objective

Safeguard information and assets hosted in SCMs, from unauthorized access, use, disclosure, modification, disposal, transmission, or destruction throughout their life cycle.
This section contains the Guardrails that address controls in the following contexts:

- System and Communications Protection (SC)

## Data Location Requirements

| Activity | Validation |
| --- | --- |
| According to subsection 4.4.3.14 of the Directive on Service and Digital: “Ensuring computing facilities located within the geographic boundaries of Canada or within the premises of a Government of Canada department located abroad, such as a diplomatic or consular mission, be identified and evaluated as a principal delivery option for all sensitive electronic information and data under government control that has been categorized as Protected B, Protected C or is Classified.” | <ul><li>Source code most often contains only Unclassified information. There are a few general exceptions, including: <ul><li>those in the [Access to Information Act](http://laws-lois.justice.gc.ca/eng/acts/A-1/),</li><li> code for systems that perform audit and testing tasks,</li><li> and code for systems that handle financial transactions. Outside of these exceptions, it is expected that code will be properly secured by removing any sensitive information (e.g. secrets) and hold the status of Unclassified.</li></ul></li></ul> |

## Mandatory Requirements

| Activity | Validation |
| --- | --- |
| Implement an encryption mechanism to protect the confidentiality and integrity of data when data is at rest in storage.| <ul><li>For SaaS, confirm that the SCM platform has implemented encryption to protect customer data.</li></ul> |
| Use cryptographic algorithms and protocols approved by Communications Security Establishment Canada (CSE) in accordance with ITSP.40.111 and ITSP.40.062.| <ul><li>Cryptographic algorithms and protocols configurable by the consumer are in accordance with ITSP.40.111 and ITSP.40.062.</li><li>For SaaS, confirm that the CSP has implemented algorithms that align with ITSP.40.111 and ITSP.40.062.</li></ul>|
| Enforce the use of _Pull Request_ (PR) reviews, and _Protected Branches_ to ensure that code changes are reviewed and approved by at least one other developer before being merged into the main branch.| <ul><li>Confirm that PR reviews are enforced for all code changes being merged into the default branch of the repository (often called _main_, or _develop_).</li></ul> |

## Conditional Requirements

| Activity | Validation |
| --- | --- |
| When dealing with personal information in cloud-based environments, seek guidance from privacy and access to information officials within institutions.| <ul><li>Confirm that privacy is part of the departmental software development life cycle.</li></ul> |
| When available, leverage an appropriate key management system for the cryptographic protection used in cloud-based services, in accordance with the Government of Canada Considerations for the Use of Cryptography in Commercial Cloud Services and the Cyber Centre’s [Guidance on Cloud Service Cryptography (ITSP.50.106)](https://www.cyber.gc.ca/en/guidance/guidance-cloud-service-cryptography-itsp50106). | <ul><li>Confirm that a key management strategy has been adopted for the SCM platform.</li></ul> |

### Self-hosting considerations

| Activity | Validation |
| --- | --- |
| Implement data protection mechanisms to protect data in transit.| <ul><li>Use TLS (at least version 1.2) to encrypt data in transit.</li></ul> |
| Regularly backup data and test the restoration process to ensure that data can be recovered in the event of data loss.| <ul><li>Provide evidence that data is regularly backed up and that the restoration process has been tested.</li></ul> |
| Implement secure data disposal procedures to ensure that data is completely removed when no longer needed. | <ul><li>Provide evidence that secure data disposal procedures are implemented, according to [IT media sanitization (ITSP.40.006)](https://www.cyber.gc.ca/en/guidance/it-media-sanitization-itsp40006)</li></ul> |
| Implement data loss prevention (DLP) mechanisms to prevent unauthorized data exfiltration.| <ul><li>Provide evidence that DLP mechanisms are implemented to prevent unauthorized data exfiltration.</li></ul> |

## References

- [Direction on the Secure Use of Commercial Cloud Services: Security Policy Implementation Notice](https://www.canada.ca/en/treasury-board-secretariat/services/access-information-privacy/security-identity-management/direction-secure-use-commercial-cloud-services-spin.html) (SPIN) 2017-01, subsections 6.2.3 & 6.2.4
- [Directive on Service and Digital](https://www.tbs-sct.canada.ca/pol/doc-eng.aspx?id=32601), subsection 4.4.3.14
- [Access to Information Act](http://laws-lois.justice.gc.ca/eng/acts/A-1/)
- cryptography guidance in [Cryptographic Algorithms for Unclassified, Protected A and Protected B Information (ITSP.40.111)](https://cyber.gc.ca/en/guidance/cryptographic-algorithms-unclassified-protected-and-protected-b-information-itsp40111) and [Guidance on Securely Configuring Network Protocols (ITSP.40.062)](https://www.cyber.gc.ca/en/guidance/guidance-securely-configuring-network-protocols-itsp40062)
- [Guidance on Cloud Service Cryptography (ITSP.50.106)](https://www.cyber.gc.ca/en/guidance/guidance-cloud-service-cryptography-itsp50106)
- [Guidance on Defence in Depth for Cloud-Based Services (ITSP.50.104)](https://cyber.gc.ca/en/guidance/itsp50104-guidance-defence-depth-cloud-based-services), subsection 4.5
- [IT Media Sanitization (ITSP.40.006)](https://www.cyber.gc.ca/en/guidance/it-media-sanitization-itsp40006)

## Related security controls from ITSG-33

SC-12, SC-13, SC-17, SC-28, SC-28(1)
