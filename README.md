# Voyna

![License: AGPL](https://img.shields.io/badge/license-AGPL-%233897f0)

Voyna is an open-source search engine, presently in the works.

## Why?

TODO

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
