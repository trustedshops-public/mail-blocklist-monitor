FAQ
===

## What does (classification) level mean?

Most spam lists have multiple levels of entries. When a query is sent to a DNS server asking if a ip is listed the
response is a IP address that also works as as code:

For example
the [SpamHaus project](https://www.spamhaus.org/faq/section/DNSBL%20Usage#:~:text=Spamhaus%20Technology%27s%20blog.-,What%20do%20the%20127.*.*.*%20Return%20Codes%20mean%3F,-Spamhaus%20uses%20this)
provides a list of valid return codes.

These codes are mapped inside the code with level names, where `white` is the best and `brown` the worst.

## How does this whole blocklist thing work?

If you have a lot of time you can read [RFC5782](https://datatracker.ietf.org/doc/html/rfc5782).

The short version is:

1. Each blocklist provider generates DNS records for EACH ip that is listed (even for ranges every single IP).
2. The DNS system does its job and propagates the records
3. You can query with A/AAAA-records for listed IPs
    1. Inverse the IP blocks e.g. `127.0.0.2` -> `2.0.0.127`
    2. Select your spam list e.g. `zen.spamhaus.org`
    3. Query it: `dig 74.0.91.223.zen.spamhaus.org a +noall +answer`
    4. *Optional*: When you get an result, check the TXT record for more details (e.g. a link for removal request):
       `dig 74.0.91.223.zen.spamhaus.org txt +noall +answer`
4. Due to the nature of DNS removals are taking sometime and also when you get listed it can take sometime till everyone
   gets the new records.
