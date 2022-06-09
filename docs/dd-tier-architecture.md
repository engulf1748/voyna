Here are some useful extracts from
https://drewdevault.com/2020/11/17/Better-than-DuckDuckGo.html

> Crucially, I would not have it crawling the entire web from the outset.
Instead, it should crawl a whitelist of domains, or “tier 1” domains. These
would be the limited mainly to authoritative or high-quality sources for their
respective specializations, and would be weighed upwards in search results.
Pages that these sites link to would be crawled as well, and given tier 2
status, recursively up to an arbitrary N tiers. Users who want to find, say, a
blog post about a subject rather than the documentation on that subject, would
have to be more specific: “$subject blog posts”.

> An advantage of this design is that it would be easy for anyone to take the
software stack and plop it on their own servers, with their own whitelist of
tier 1 domains, to easily create a domain-specific search engine. Independent
groups could create search engines which specialize in academia, open standards,
specific fandoms, and so on. They could tweak their precise approach to
indexing, tokenization, and so on to better suit their domain.

> We should also prepare the software to boldly lead the way on new internet
standards. Crawling and indexing non-HTTP data sources (Gemini? Man pages? Linux
distribution repositories?), supporting non-traditional network stacks (Tor?
Yggdrasil? cjdns?) and third-party name systems (OpenNIC?), and anything else we
could leverage our influence to give a leg up on.
