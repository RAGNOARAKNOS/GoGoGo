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

## Enums

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
