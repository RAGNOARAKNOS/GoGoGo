# App Notes

## Dependencies

```shell
docker pull postgres
docker pull dpage/pgadmin4
```

```shell
docker run --name postgres -p 5432:5432 -e POSTGRES_USER=devuser -e POSTGRES_PASSWORD=password -e POSTGRES_DB=devdb -d postgres

docker run --name pgadmin -p 82:80 -e 'PGADMIN_DEFAULT_EMAIL=test@dev.com' -e 'PGADMIN_DEFAULT_PASSWORD=pass123' -d dpage/pgadmin4
```

REMEMBER: host.docker.internal is the DNS name of the host from within containers (172.17.0.1)

REMEMBER: To access PGADMIN4 use link [PGAdmin dashboard](http://localhost:82/browser/)

## Config Values

A .ENV file located in the root is read to configure the below values, EnvVars can ALSO be used to override

| Variable | Description | Default |
| --- | --- | --- |
| GAMELIB_DB_HOST | Hostname of Postgres instance | 127.0.0.1 |
| GAMELIB_DB_USER | Postgres username | devuser |
| GAMELIB_DB_PASSWORD | Postgres password | password|
| GAMELIB_DB_NAME | Postgres database instance name | devdb|
| GAMELIB_DB_PORT | Postgres port | 5432 |
| GAMELIB_REST_PORT | Application REST endpoint port | 9999 |
| GAMELIB_RUNTIME | sets deployment location, used for debug flags | dev |

## Tags

### Genres

- Abstract Strategy - Games like Chess or Go, focusing on pure strategy without theme or narrative.
- Cooperative - Players work together toward a common goal, e.g., Pandemic.
- Deck-Building - Players build decks during play, like Dominion or Star Realms.
- Economic - Centered on resource management and trade, e.g., Power Grid.
- Eurogame - Strategy-focused with indirect player interaction, like Settlers of Catan.
- Party - Light, social games for groups, e.g., Codenames or Telestrations.
- Roll-and-Move - Movement based on dice rolls, like Monopoly.
- Social Deduction - Involves hidden roles and bluffing, e.g., Werewolf or The Resistance.
- Thematic - Strong narrative or theme, like Betrayal at House on the Hill.
- Wargame - Simulate conflicts with detailed rules, e.g., Risk or Twilight Struggle.
- Worker Placement - Assign workers to tasks, like Agricola or Lords of Waterdeep.
- Area Control - Compete for territory, e.g., Ticket to Ride or Small World.
- Puzzle - Focus on solving problems, like Sagrada.
- Dexterity - Physical skill-based, e.g., Jenga or Flick 'em Up.
- Legacy - Games with evolving rules and components, like Pandemic Legacy.

## GORM

I don't want to be writing SQL code manually, and given that this is an excercise in learning a language and its ecosystem I have made the _conscious choice_ to use an ORM (against my better judgement).  I will be using [Gorm](http://www.gorm.io).

## JSON Validation

I really want to _pretend_ to follow best practise, so I will have seperate structs for ORM and DTO layers.  Any DTO object presented via the REST JSON API will be validated using markup and the validatorv10 library [here](https://github.com/go-playground/validator) to avoid the unnecessary bolilerplate that a librbaryless approach [like this](https://betterstack.com/community/guides/scaling-go/json-in-go/) would entail.

## Mux

I want to present a simple REST api, I will use the lightweight and minimal Mux library.  A swagger spec will be retrofitted at a later point (if time allows)

## Lessons Learned

### Enums in Go and Postgres - LIMITING

I had initially implemented the GameGenre as an enumeration, which necessitated that I implement a custom postgres type and mirror that in the Go structs.  It is very limiting, and having an ever evolving genres table, or better yet - a reuse of the tags table, allows for many genres to be allocated and NO hardcoding wonkyness.

### Complete tangent

Discovered that WSL2 uses the C drive by default for the distros, but it doesnt implement a storage reclamation subsystem like UNMAP or TRIM.  So you have to manually reclaim space, if you dont want Linux runtime nodes to slowly consume all your disk space, you'll need to reclaim space manually (see below)

```shell
wsl --shutdown
optimize-VHD -Path .\ext4.vhdx -Mode full
```

REMEMBER: replace with all vdhx's, my system may have several
