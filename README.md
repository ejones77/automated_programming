# automated_programming
Testing different automated programming methods on a previous Go assignment

This reposiory tests three different automated programming methods on a previous assignment repository `go_with_stats` to see how well the AI stacks up in approaching the problem when prompted.

The code here is not intended to be compiled, rather showcase how automated programming currently works (and in some cases doesn't work) for Go code.

1 - Programmatic code generation with the `jen` package in `main.go`

- This seems useful for constructing boilerplate code like with testing for example. The code I generated at first did not pass the linting stage.

2 - Assistive programming with Github Copilot in `gh_copilot.go`

- I used the "fix this" feature to correct my broken jennifer code, and provided a comment prompt to create an overall function for the anscombe quartet summary statistics. 

- This seemed to work quite well, but it does have the context of the previous code and could repeat it out verbatim. This is a helpful feature because copilot is able to pick up on existing code when you're modifying it, not just generating it. 

3 - Code generations via LLM in `llm_stats.go`

- The first thing I tried was to prompt CodeLlama-70b-Instruct-hf using huggingface chat. I learned that Llama does not quite understand the incentive of being tipped in [my first attempt](https://hf.co/chat/r/lqqYRHc) and did not actually generate code.

- The second attempt with CodeLlama did produce generated code, but it was assuming that large portions of the problem were [already written](https://hf.co/chat/r/uzFLU04).

- The third attempt is what's provided in `llm_stats.go` where I returned to the financial incentive-based approach using a different model, mistralai/Mixtral-8x7B-Instruct-v0.1. This approach worked better in terms of understanding the scope of the request, but [ultimately hallucinates](https://hf.co/chat/r/enHahTG) functions associated with the statistics package. 

Overall I think copilot offers the most benefit to the developer workflow out of these options. I think that LLMs are most useful in tasks that are limited in scope, i.e. assistance generating portions of a program as opposed to generating a response to the whole problem. 

Resources used: 
- [HuggingChat](https://huggingface.co/chat/) -- A free ChatGPT alternative using a variety of open source LLMs -- there is a limit on the number of responses before being required to make an account.

- [Github Copilot](https://github.com/features/copilot) -- very easy to integrate with my existing github acount and VS code -- inputting payment info is requried but there is a 30 day free trial. 

- [Jen examples](https://www.carlomaiorano.me/golang/2019/10/03/generating-code-golang.html) -- A useful blog post on code generations with Jennifer