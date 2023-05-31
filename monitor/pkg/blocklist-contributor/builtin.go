package blocklist_contributor

// DefaultBlocklists to check
var DefaultBlocklists = []string{
	// 0spam
	"bl.0spam.org",
	"nbl.0spam.org",
	"url.0spam.org",
	// Abuse.ro
	"rbl.abuse.ro",
	"pbl.abuse.ro",
	// Abusix
	"combined.mail.abusix.zone",
	"black.mail.abusix.zone",
	"exploit.mail.abusix.zone",
	"dynamic.mail.abusix.zone",
	// Anomails
	"spam.dnsbl.anonmails.de",
	// BackScatterer
	"ips.backscatterer.org",
	// Barracuda
	"b.barracudacentral.org",
	// Blocklist.de
	"bl.blocklist.de",
	// Calivent
	"dnsbl.calivent.com.pe",
	// Cymru Bogons
	"bogons.cymru.com",
	// DAN TOR
	"torexit.dan.me.uk",
	"tor.dan.me.uk",
	// DNS Servicos
	"rbl.dns-servicios.com",
	// DRMX
	" bl.drmx.org",
	// DRONE BL
	"dnsbl.dronebl.org",
	// FabelSources
	"spamsources.fabel.dk",
	// HIL
	"hil.habeas.com",
	// Hostkarma
	"hostkarma.junkemailfilter.com",
	// IBM
	"dnsbl.cobion.com",
	// Improware
	"dnsrbl.imp.ch",
	"spamrbl.imp.ch",
	"wormrbl.imp.ch",
	// Interserver
	"rblspamassassin.interserver.net",
	// ivmSIP
	"sip.invalument.com",
	// Lashback
	"ubl.unsubscore.com",
	// LNSG
	"spamguard.leadmon.net",
	// Madavi
	"dnsbl.madavi.de",
	// Mailspike
	"bl.mailspike.net",
	// MRSBL
	"virus.rbl.msrbl.net",
	"phishing.rbl.msrbl.net",
	"images.rbl.msrbl.net",
	"spam.rbl.msrbl.net",
	"combined.rbl.msrbl.net",
	// Nether
	"relays.nether.net",
	// NordSpam
	"bl.nordspam.com",
	"dbl.nordspam.com",
	// NoSolicitado
	"bl.worst.nosolicitado.org",
	"bl.nosolicitado.org",
	// OrevDB
	"rsbl.aupads.org",
	// PSBL
	"psbl.surriel.com",
	// Rats
	"noptr.spamrats.com",
	"spam.spamrats.com",
	"auth.spamrats.com",
	"dyna.spamrats.com",
	// Schulte
	"rbl.schulte.org",
	// SEM
	"backscatter.spameatingmonkey.net",
	"bl.spameatingmonkey.net",
	"netbl.spameatingmonkey.net",
	"urired.spameatingmonkey.net",
	"ixhash.spameatingmonkey.net",
	// Sender Score Reputation Network
	"bl.score.senderscore.com",
	// ServicesNET
	"korea.services.net",
	// SORBS has been skipped
	// SpamCop
	"bl.spamcop.net",
	// SpamHaus
	"zen.spamhaus.org",
	// SPFBL
	"dnsbl.spfbl.net",
	// SuomiSpam
	"dbl.suomispam.net",
	// TRIUMF
	"rbl2.triumf.ca",
	// Truncate
	"truncate.gbudb.net",
	// UCEPROTECT
	"dnsbl-1.uceprotect.net",
	"dnsbl-2.uceprotect.net",
	"dnsbl-3.uceprotect.net",
	// Woodys
	"blacklist.woody.ch",
	// WPBL
	"db.wpbl.info",
	// ZapBL
	"dnsbl.zapbl.net",
}

func init() {
	registerProvider(BuiltinBlocklistContributor{})
}

type BuiltinBlocklistContributor struct {
}

func (b BuiltinBlocklistContributor) Name() string {
	return "builtin"
}

func (b BuiltinBlocklistContributor) Contribute(blockLists []string) []string {
	return append(blockLists, DefaultBlocklists...)
}

func (b BuiltinBlocklistContributor) Priority() int {
	return PriorityAdd
}
