# Plan for Continuity

([Back](../../GUARDRAILS.md))

## Objective

Ensure that there is a plan for continuity of access and service that accommodates both expected and unexpected events.
This section contains the Guardrails that address controls in the following contexts:

- Access Control (AC)
- Security Assessment and Authorization (CA)
- Contingency Planning (CP)
- Incident Response (IR)

## Conditional Requirements

| Activity | Validation |
| --- | --- |
| Develop an SCM backup strategy that considers where GC data is stored, replicated, or backed up by the SCM service, and the IT continuity plan for the service or application. | <ul><li>Confirm through attestation that the SCM backup strategy is developed and approved by the business owner.</li><li>Verify if there are scripts that support the ability to restore from code (for example, infrastructure as code).</li></ul> |

### Self-hosting considerations

| Activity | Validation |
| --- | --- |
| Document, implement, and test a break glass emergency account management process. | <ul> <li>Verify that an emergency account management procedure has been developed</li><li>Verify that alerts are in place to report any use of emergency accounts</li> <li>Verify that testing of emergency accounts took place, and that periodic testing is included in emergency account management procedures.</li> </ul> |
| Obtain confirmation from the departmental chief information officer (CIO) in collaboration with the designated official for cyber security (DOCS) with signatures that acknowledge and approve the emergency account management procedures. | <ul><li>Confirm through attestation that the departmental CIO, in collaboration with the DOCS, has approved the emergency account management procedure for the SCM service.</li> </ul> |

## References

- [Direction on the Secure Use of Commercial Cloud Services: Security Policy Implementation Notice](https://www.canada.ca/en/treasury-board-secretariat/services/access-information-privacy/security-identity-management/direction-secure-use-commercial-cloud-services-spin.html) (SPIN) 2017-01, subsection 6.2.9
- [Break glass emergency account management procedure](https://gcconnex.gc.ca/file/view/55010566/break-glass-emergency-account-procedure-departments-can-use-to-develop-their-emergency-access-management-controls-for-cloud?language=en) for Azure and Office 365 (accessible only on the Government of Canada network)
- [Cyber Security Event Management Plan template](https://www.gcpedia.gc.ca/gcwiki/images/6/66/Department_CSEMP_Template.docx) for departments (Word file) (accessible only on the Government of Canada network)
- [Directive on Service and Digital](https://www.tbs-sct.canada.ca/pol/doc-eng.aspx?id=32601)

## Related security controls from ITSG-33

AC-1, AC-2(1),
CA-3,
CP-1, CP-2, CP-2(1), CP-2(2), CP-2(5), CP-2(6), CP-2(8), CP-3, CP-4, CP-4(1), CP-4(2), CP-6, CP-6(1), CP-6(2), CP-6(3), CP-7, CP-7(1), CP-7(2), CP-7(3), CP-7(4), CP-8, CP-8(1), CP-8(2), CP-8(3), CP-8(5), CP-9, CP-9(1), CP-9(2), CP-9(3), CP-9(5), CP-9(10), CP-10, CP-10(4)
IR-1, IR-9(3)
