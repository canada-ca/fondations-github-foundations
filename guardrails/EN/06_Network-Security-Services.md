# Network Security Services

([Back](../../GUARDRAILS.md))

## Objective

Establish external and internal network perimeters and monitor network traffic.


## Mandatory Requirements

| Activity | Validation |
 --- | --- |
| Use HTTPS for all network traffic | <ul><li>Ensure that all network traffic to and from the SCM is encrypted.</li></ul> |

## Conditional Requirements

| Activity |    Validation |
| --- | --- |
| When using Public and Private Repositories, keep them separate | <ul><li>Publicly accessible repositories should be separated from private repositories used for internal development</li></ul> |

### Self-hosting considerations

| Activity | Validation |
| --- | --- |
| Network Segmentation | <ul><li>Ensure that the SCM is hosted on a separate network segment from the internal organizational network.</li></ul> |
| Using Managed Interfaces | <ul><li>Connect to the SCM only through managed interfaces, such as a VPN or a secure API.</li></ul> |
| Using Boundary Protection | <ul><li>Use firewalls and other boundary protection devices to control access to the SCM.</li></ul> |
| Performing Regular Audits | <ul><li>Regularly audit the security configuration of the SCM and the boundary protection devices to ensure they are in accordance with the organizational security architecture.</li></ul> |
| Employ Network monitoring tools | <ul><li>Use network monitoring tools to monitor for: <ul><li>unauthorized use of the information system</li><li>attacks and indicators of potential attacks in accordance with monitoring objectives consistent with the GC CSEMP.</li><li>network intrusion</li></ul></li></ul> |
| Provide system monitoring information to necessary stakeholders | <ul><li>As detailed in the GC CSEMP, provide system monitoring information to the Canadian Centre for Cyber Security (Cyber Centre) and other departmental monitoring organizations.</li></ul> |
| Heighten the level of information system monitoring activity whenever there is an indication of increased risk to organizational assets. | <ul><li>Monitor the SCM for suspicious activity and respond to incidents in accordance with the organizational incident response plan. </li></ul>|
|Obtain a legal opinion with regard to system monitoring | <ul><li>Obtain a legal opinion with regard to system monitoring to ensure that it is in compliance with all applicable Government of Canada legislation, and TBS policies.</li></ul> |

## References

- [Direction on the Secure Use of Commercial Cloud Services: Security Policy Implementation Notice](https://www.canada.ca/en/treasury-board-secretariat/services/access-information-privacy/security-identity-management/direction-secure-use-commercial-cloud-services-spin.html) (SPIN) 2017-01, subsection 6.2.4
- Cyber Centre’s top 10 IT security actions, number 1
- network security zoning guidance in [Baseline Security Requirements for Network Security Zones (ITSP.80.022)](https://cyber.gc.ca/en/guidance/baseline-security-requirements-network-security-zones-government-canada-itsg-22) and [Network Security Zoning (ITSG-38)](https://cyber.gc.ca/en/guidance/network-security-zoning-design-considerations-placement-services-within-zones-itsg-38)
- [Guidance on Defence in Depth for Cloud-Based Services (ITSP.50.104)](https://cyber.gc.ca/en/guidance/itsp50104-guidance-defence-depth-cloud-based-services), subsection 4.3

## Related security controls from ITSG-33

SC‑7, SI-4
