# Voyna

![License: AGPL](https://img.shields.io/badge/license-AGPL-%233897f0)

Voyna is an open-source search engine, presently in the works.

## Why?

Today's search engines do not always yield great results--even Google. Google's algorithm (the singular "algorithm" is somewhat of a misnomer, as there is never a single algorithm for a complex system), Hummingbird, emphasizes context and meaning of text appearing on websites. Needless to say, there are very likely multiple other parts to Hummingbird--PageRank and "authority centers", for instance. Let us have a look at some sub-par (subjectively) results returned by Google.

Assume for a moment that you are writing critical banking software in Python. You happen to require language support for manipulating date and time. You open up Google, and search for "date and time python". Here is what you will (likely) see:

![google-date-and-time-python](/docs/images/google-date-and-time-python.jpeg)

The top two results point to hosts "www.w3schools.com" and "www.geeksforgeeks.org". Why could that be? Perhaps they contain "natural" language, and are linked to by a relatively large number of websites (PageRank, that is). However, both these results are not "authoritative"--that is, while they may be easily and readily comprehended, and relatively more succinct, they are less valuable as results in many cases--writing critical banking software is one such case.  The authoritative source, Python's official documentation, is ranked third (I had to zoom out on my 1920x1080 monitor):

![google-date-and-time-python-zoomed-out](/docs/images/google-date-and-time-python-zoomed-out.jpeg)

There might also be cases where some other website, perhaps Guido van Rossum's blog, or the musings of a talented Python hacker, ought to be ranked higher. While Google does give us options to tweak our search parameters, such as returning results only from specific sites using the "site:" directive, the painstaking job of curating authoritative and relevant domains falls on its users. Google probably does not explicitly set authoritative sources for all categories of search--perhaps it only does so for "critical" topics like health-care and medicine. As long as Google Search is closed-source, others cannot set up their own instances with Hummingbird tweaked to favour their own sets of authoritative sources.

For subjectively excellent results, we require a federated deployment of a search algorithm, with each node curating its own authority sources. Those unhappy with one node can always switch over to another node, or even better, set up their own nodes. I first came across this idea in a [blog post](https://drewdevault.com/2020/11/17/Better-than-DuckDuckGo.html) by Drew Devault, where he proposes a federated, tier-based search engine. The owner of a node first sets their own authoritative domains, or "tier 1" websites. These tier 1 domains serve as the search engine's "seed", or where it begins crawling from. Links pointed to by tier 1 pages are assigned tier 2, and so on, up to a configurable "n" tiers. When responding to a search query, results with higher tiers are ranked higher. This would immensely help deliver subjectively better results: for example, an instance dedicated to programming might choose to choose the following as their authoritative sources:

* https://*.python.org
* https://*.go.dev
* https://developer.mozilla.org
. . .

[Voyna](https://codeberg.org/voyna/voyna) is currently a primitive prototype of the tier-based architecture, created to get a glimpse of how it would work in practice. What is remarkable is how quickly results "deteriorate" with each increasing tier. Even with "n" (the maximum depth) set to 3, we go from good, authoritative sources to websites that are notorious for serving "clickbaity" content. Such is the connectedness of the web.

Apart from getting better search results, there is another reason why we require an alternative to Google (and other such closed-source search engines): fundamental code infrastructure should ideally be open-source--it is remarkable that we do not have an open-source project for something as fundamental as a search engine! I think of search engines as tooling--and most of tooling is open-source: compilers, compiler back-ends, the Linux kernel, crucial cryptography code, and so on. A transparent, open-source, and federated search engine would certainly change the world for the better.

## Development

Note that Voyna is in its very nascency--it is not suitable for your day-to-day
usage. As of now, the crawler will almost certainly crash after creating a spate
of goroutines.

Run `go run .` in the project's root to start the crawler and processor. Then
run `go run ./server/` to start the search server. Make sure to set your `DEV`
environment variable to `true`: this gives you easy-to-read results. You can
then query Voyna in the following way: `curl 'localhost:8080/search?q=openbsd'`.
The 'q' query parameter represents your search query.

## License

Voyna (A search engine.)

Copyright (C) 2022 Ajay R, Tristan B

This program is free software: you can redistribute it and/or modify it under
the terms of the GNU Affero General Public License as published by the Free
Software Foundation, either version 3 of the License, or (at your option) any
later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY
WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A
PARTICULAR PURPOSE.  See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along
with this program.  If not, see <https://www.gnu.org/licenses/>.
