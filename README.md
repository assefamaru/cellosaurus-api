# cellosaurus-api

The [Cellosaurus](https://github.com/calipho-sib/cellosaurus) is a knowledge resource on cell lines. It attempts to describe all cell lines used in biomedical research. This API aims to make the data provided by Cellosaurus as integrable as possible, by providing programmatic access to the full database.

NOTE: The live version of the API is hosted on a free tier plan on heroku, and [sleeps after 30 minutes of inactivity](https://devcenter.heroku.com/articles/free-dyno-hours#dyno-sleeping). As a result, you might experience a lag on your first request.

## Accessing the API

All calls are made to the following base URL, adding required endpoints for specific services.

```bash
https://api.cellosaur.us/api/v42/
```

All responses are in `json` format.

The following previous versions are also available for the live API:

| Version | Base URL                                                       |
| :-----: | :------------------------------------------------------------- |
|  `41`   | [`https://api.cellosaur.us/v41`](https://api.cellosaur.us/v41) |
|  `40`   | [`https://api.cellosaur.us/v40`](https://api.cellosaur.us/v40) |

See [Tags](https://github.com/assefamaru/cellosaurus-api/tags) to build any previous versions locally.

## Endpoints

The following endpoints are currently supported:

| Method | Endpoint              | Parameter(s)      | Example                                                     |
| :----: | :-------------------- | :---------------- | :---------------------------------------------------------- |
|  GET   | **/cells**            | `page`, `perPage` | <https://api.cellosaur.us/v41/cells?page=8&perPage=10>      |
|  GET   | **/cell-lines**       | `page`, `perPage` | <https://api.cellosaur.us/v41/cell-lines?page=3&perPage=20> |
|  GET   | **/cells/{id}**       |                   | <https://api.cellosaur.us/v41/cells/mcf-7>                  |
|  GET   | **/cell-lines/{id}**  |                   | <https://api.cellosaur.us/v41/cell-lines/mcf-7>             |
|  GET   | **/refs**             | `page`, `perPage` | <https://api.cellosaur.us/v41/refs?page=1&perPage=10>       |
|  GET   | **/references**       | `page`, `perPage` | <https://api.cellosaur.us/v41/references?page=1&perPage=10> |
|  GET   | **/xrefs**            |                   | <https://api.cellosaur.us/v41/xrefs>                        |
|  GET   | **/cross-references** |                   | <https://api.cellosaur.us/v41/cross-references>             |
|  GET   | **/stats**            |                   | <https://api.cellosaur.us/v41/stats>                        |
|  GET   | **/statistics**       |                   | <https://api.cellosaur.us/v41/statistics>                   |

Endpoints should always be prefixed with the current version number when making requests.

Parameters are not required in request URLs. When parameters are not included, they are set to their default values:

| Parameter | Default Value |
| :-------- | :-----------: |
| `page`    |      `1`      |
| `perPage` |     `10`      |

Endpoints that contain `page` and `perPage` parameters will have a `meta` field in their response containing pagination information, as well as the total number of records under the requested resource type. Such a response will look as follows:

```json
{
    "meta": {
        "page": ,
        "perPage": ,
        "lastPage": ,
        "total":
    },
    "data": []
}
```

The maximum number of records that can be returned in a single request (by modifying `perPage`) is `1000`.

## Getting Started Locally

To setup locally, first install Go and MySQL locally. Then, follow the steps below:

1. Clone this repository:

```bash
git clone https://github.com/assefamaru/cellosaurus-api
```

2. Run provision script:

```bash
./scripts/provision.sh
```

3. Export environment variables:

```bash
export MYSQL_SERVICE_USER=xyz  # the username you used in step 3
export MYSQL_SERVICE_PASS=xyz  # the password you used in step 3
export MYSQL_SERVICE_DB=xyz    # the database name you provided in step 3
export MYSQL_SERVICE_HOST=xyz  # eg. localhost
export MYSQL_SERVICE_PORT=xyz  # eg. 3306
export PORT=xyz                # eg. 8080
```

4. Run the locally built API inside bin directory:

```bash
./bin/cellosaurus-api-<os>-<platform>
```

5. Access the locally running API:

```bash
curl localhost:8080/api/v42/
```

## Troubleshooting

When loading data to local mysql instance, you may encounter an error such as:

```bash
Loading local data is disabled; this must be enabled on both the client and server side
```

You can fix this by enabling local infile:

```bash
mysql> SET GLOBAL local_infile=1;
Query OK, 0 rows affected (0.00 sec)
```

## Contributing

You can contribute in various ways:

- Send PRs (pull requests) with new feature implementations, fixed bugs etc.
- Offer suggestions, request new features, or report any errors by creating [new](https://github.com/assefamaru/cellosaurus-api/issues/new) issues and assigning them appropriate labels.

## License

This project is under the MIT License - see [LICENSE](LICENSE) for details.
