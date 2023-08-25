# GPT powered Testcase Generator

Simply explain what the feature is expected to do, and boom you should ideally get a exhaustive list of testcases that needs to be covered. 


## Setup

    - Inside `scripts/env` folder, clone the `env.sample` as `development.env` and fill the env vars.
    - Run  `make deploy_local`. Will start-up the server at `$PORT`.

## How to test?
    - UI doesn't work yet, use :-

        `curl -Ns localhost:11000/start/generate/scenarios/streaming | python scripts/py/print_resp.py` for the time being. 
        Which returns a response for `A REST API to fetch a list of users.` prompt.

    - Change the prompt inside `StartStreamingGenerateTestcaseSenariosHandler` function 
    at `http_handlers.go` to get another answer.

## TODO:
 - Implementing the UI.


