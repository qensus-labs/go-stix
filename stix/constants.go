package stix

const (
	// STIX specification version.
	STIXVersion = "2.1"

	// Meta object types.
	TypeBundle = "bundle"

	// Cyber Observable Objects (SCO).
	TypeIPv4Addr              = "ipv4-addr"
	TypeIPv6Addr              = "ipv6-addr"
	TypeMACAddr               = "mac-addr"
	TypeDomainName            = "domain-name"
	TypeURL                   = "url"
	TypeEmailAddr             = "email-addr"
	TypeFile                  = "file"
	TypeDirectory             = "directory"
	TypeArtifact              = "artifact"
	TypeAutonomousSystem      = "autonomous-system"
	TypeUserAccount           = "user-account"
	TypeProcess               = "process"
	TypeSoftware              = "software"
	TypeX509Certificate       = "x509-certificate"
	TypeRegistryKey           = "windows-registry-key"
	TypeNetworkTraffic        = "network-traffic"
	TypeMutex                 = "mutex"
	TypeUserAgent             = "user-agent"
	TypeHostname              = "hostname"
	TypeCryptographicKey      = "cryptographic-key"
	TypeCryptocurrencyAddress = "cryptocurrency-address"

	// STIX Domain Objects (SDO).
	TypeAttackPattern   = "attack-pattern"
	TypeCampaign        = "campaign"
	TypeCourseOfAction  = "course-of-action"
	TypeGrouping        = "grouping"
	TypeIdentity        = "identity"
	TypeIncident        = "incident"
	TypeIndicator       = "indicator"
	TypeInfrastructure  = "infrastructure"
	TypeIntrusionSet    = "intrusion-set"
	TypeLocation        = "location"
	TypeMalware         = "malware"
	TypeMalwareAnalysis = "malware-analysis"
	TypeNote            = "note"
	TypeObservedData    = "observed-data"
	TypeOpinion         = "opinion"
	TypeReport          = "report"
	TypeThreatActor     = "threat-actor"
	TypeTool            = "tool"
	TypeVulnerability   = "vulnerability"

	// STIX Relationship Objects (SRO).
	TypeRelationship = "relationship"
	TypeSighting     = "sighting"

	// STIX Cyber Observable extensions / commonly used objects.
	TypeExtensionDefinition = "extension-definition"
)
