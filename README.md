# cellosaurus-api

The [Cellosaurus](https://web.expasy.org/cellosaurus/) is a knowledge resource on cell lines. It attempts to describe all cell lines used in biomedical research. This API aims to make the data provided by Cellosaurus as integrable as possible, by providing programmatic access to the full database.

NOTE: The live version of the API is hosted on a free tier plan on heroku, and [sleeps after 30 minutes of inactivity](https://devcenter.heroku.com/articles/free-dyno-hours#dyno-sleeping). As a result, you might experience a lag on your first request.

## Accessing the API

All calls are made to the following base URL, adding required endpoints for specific services.

```
https://api.cellosaur.us/v41
```

All responses are in `json` format.

The following previous versions are also available for the live API:

| Version |                            Base URL                            |
| :-----: | :------------------------------------------------------------: |
|  `40`   | [`https://api.cellosaur.us/v40`](https://api.cellosaur.us/v40) |

See [Tags](https://github.com/assefamaru/cellosaurus-api/tags) to build any previous versions locally.

## Endpoints

The following endpoints are currently supported:

| Method | Endpoint             | Parameter(s)                | Example                                                         |
| :----: | :------------------- | :-------------------------- | :-------------------------------------------------------------- |
|  GET   | **/cells**           | `page`, `perPage`, `indent` | https://api.cellosaur.us/v41/cells?page=8&perPage=10            |
|  GET   | **/cell-lines**      | `page`, `perPage`, `indent` | https://api.cellosaur.us/v41/cell-lines?page=3&perPage=20       |
|  GET   | **/cell_lines**      | `page`, `perPage`, `indent` | https://api.cellosaur.us/v41/cell_lines?page=5&perPage=30       |
|  GET   | **/cells/{id}**      | `indent`                    | https://api.cellosaur.us/v41/cells/mcf-7?indent=true            |
|  GET   | **/cell-lines/{id}** | `indent`                    | https://api.cellosaur.us/v41/cell-lines/mcf-7?indent=true       |
|  GET   | **/cell_lines/{id}** | `indent`                    | https://api.cellosaur.us/v41/cell_lines/mcf-7?indent=true       |
|  GET   | **/refs**            | `page`, `perPage`, `indent` | https://api.cellosaur.us/v41/refs?page=1&perPage=10&indent=true |
|  GET   | **/xrefs**           | `indent`                    | https://api.cellosaur.us/v41/xrefs?indent=true                  |
|  GET   | **/stats**           | `indent`                    | https://api.cellosaur.us/v41/stats                              |

Endpoints should always be prefixed with the current version number when making a request (ie. the `v41` in `https://api.cellosaur.us/v41/<endpoint>`).

Parameters are not required in request URLs. When parameters are not included in request, they are set to their default values:

| Parameter | Default Value |
| :-------- | :-----------: |
| `page`    |      `1`      |
| `perPage` |     `10`      |
| `indent`  |    `true`     |

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

## Getting Started Locally

You can follow the steps below to setup the API locally.

1. Clone this repository:

```bash
git clone https://github.com/assefamaru/cellosaurus-api
```

2. Run provision script:

```bash
./scripts/provision.sh
```

This will initialize submodules, parse raw cellosaurus text files to generate csv formats, and setup local mysql seeded with csv data.

3. Export environment variables:

```bash
export MYSQL_SERVICE_USER=xyz  # the username you used in step 3
export MYSQL_SERVICE_PASS=xyz  # the password you used in step 3
export MYSQL_SERVICE_DB=xyz    # the database name you provided in step 3
export MYSQL_SERVICE_HOST=xyz  # eg. localhost
export MYSQL_SERVICE_PORT=xyz  # eg. 3306
export PORT=xyz                # eg. 8080
```

4. Run the API locally:

```bash
./run.sh
```

Alternatively, you can run the API via one of the generated builds inside `bin` directory.

## Contributing

You can contribute in various ways:

- Send PRs (pull requests) with new feature implementations, fixed bugs etc.
- Offer suggestions, request new features, or report any errors by creating [new](https://github.com/assefamaru/cellosaurus-api/issues/new) issues and assigning them appropriate labels.

## License

This project is under the MIT License - see [LICENSE](LICENSE) for details.
