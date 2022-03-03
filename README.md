# cellosaurus-api

The [Cellosaurus](https://web.expasy.org/cellosaurus/) is a knowledge resource on cell lines. It attempts to describe all cell lines used in biomedical research. This API aims to make the data provided by Cellosaurus as integrable as possible, by providing programmatic access to the full database.

Follow the steps below to get the API up and running in your local environment.

## Getting Started

1. Clone this repository:

```bash
git clone https://github.com/assefamaru/cellosaurus-api
```

2. Move into [_scripts_](scripts) directory:

```bash
cd cellosaurus-api/scripts
```

3. Run setup script:

```bash
./setup.sh
```

_This will initialize submodules, parse raw cellosaurus text files to generate csv formats, and setup local mysql seeded with csv data._

4. Export environment variables:

```bash
export MYSQL_SERVICE_USER=xyz  # the username you used in step 3
export MYSQL_SERVICE_PASS=xyz  # the password you used in step 3
export MYSQL_SERVICE_DB=xyz    # the database name you provided in step 3
export MYSQL_SERVICE_HOST=xyz  # eg. localhost
export MYSQL_SERVICE_PORT=xyz  # eg. 3306
export PORT=xyz                # eg. 8080
```

5. Run the API locally:

```bash
./run.sh
```

Alternatively, you can also build the api first `./build.sh`, then run one of the generated executables inside `build` directory at the root of the project.

## Endpoints

The following endpoints are currently supported:

| Method | Endpoint             | Parameter(s)                | Example                                                      |
| :----: | :------------------- | :-------------------------- | :----------------------------------------------------------- |
|  GET   | **/cells**           | `page`, `perPage`, `indent` | http://localhost:8080/v40/cells?page=8&perPage=20            |
|  GET   | **/cell-lines**      | `page`, `perPage`, `indent` | http://localhost:8080/v40/cell-lines?page=3&perPage=20       |
|  GET   | **/cell_lines**      | `page`, `perPage`, `indent` | http://localhost:8080/v40/cell_lines?page=5&perPage=20       |
|  GET   | **/cells/{id}**      | `indent`                    | http://localhost:8080/v40/cells/mcf-7?indent=true            |
|  GET   | **/cell-lines/{id}** | `indent`                    | http://localhost:8080/v40/cell-lines/mcf-7?indent=true       |
|  GET   | **/cell_lines/{id}** | `indent`                    | http://localhost:8080/v40/cell_lines/mcf-7?indent=true       |
|  GET   | **/refs**            | `page`, `perPage`, `indent` | http://localhost:8080/v40/refs?page=1&perPage=10&indent=true |
|  GET   | **/xrefs**           | `indent`                    | http://localhost:8080/v40/xrefs?indent=true                  |
|  GET   | **/stats**           | `indent`                    | http://localhost:8080/v40/stats                              |

Always prefix endpoints with the current version number when making a request (ie. the `v40` in `http://localhost:8080/v40/<endpoint>`).

Parameters need not be present in request URLs. When parameters are not included in request, they are set to their default values. The following are their default values:

| Parameters | Default Values |
| :--------- | :------------: |
| `page`     |      `1`       |
| `perPage`  |      `10`      |
| `indent`   |     `true`     |

All responses are in `json` format. Endpoints that contain `page` and `perPage` parameters will have a `meta` field in their response containing pagination information, as well as the total number of records under the requested resource type. Such a response will look as follows:

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

## Contributing

You can contribute in various ways:

- Send PRs (pull requests) with new feature implementations, fixed bugs etc.
- Offer suggestions, request new features, or report any errors by creating [new](https://github.com/assefamaru/cellosaurus-api/issues/new) issues and assigning them appropriate labels.

## License

This project is under the MIT License - see [LICENSE](LICENSE) for details.
