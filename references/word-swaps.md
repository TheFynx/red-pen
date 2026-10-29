# Word Swaps

When you find a word in the left column, replace it with the word in the right column. If none of the replacements fit the sentence, rewrite the sentence.

A swap must never change the meaning. Words that are precise technical terms are listed separately under [Technical terms](#technical-terms). Keep those when they name the technical thing.

The same logic applies outside that table. Keep a flagged word when it has a technical use in context, for example `optimize` for compiler, SQL, or algorithm optimization. The linter reports swaps as findings, not violations, so a kept word does not fail the run. To drop a false positive from the report, add an ignore directive above the offending line (see [Ignore a line or region](../SKILL.md#ignore-a-line-or-region)).

## Verbs

<!-- red-pen-ignore-start -->
| Do not use | Use instead |
|---|---|
| utilize | use |
| leverage | use |
| facilitate | help, enable |
| enhance | improve |
| streamline | simplify |
| empower | let, allow |
| enable (as marketing verb) | let, allow |
| ensure | make sure |
| establish | set, create |
| initiate | start |
| deactivate | turn off |
| activate | turn on |
| commence | start |
| indicate | show |
| demonstrate | show |
| illustrate | show |
| obtain | get |
| retrieve | get |
| acquire | get |
| modify | change |
| alter | change |
| transform | change, convert |
| navigate to | go to, open |
| specify | set, enter, type |
| encompass | include |
| comprise | include, contain |
| constitute | make up |
| necessitate | need, require |
| ascertain | find out |
| endeavor | try |
| prioritize | rank, sort |
| mitigate | reduce, limit |
| remediate | fix |
| decommission | shut down, remove |
<!-- red-pen-ignore-end -->

## Nouns and noun phrases

<!-- red-pen-ignore-start -->
| Do not use | Use instead |
|---|---|
| functionality | feature, what it does |
| capability | feature, what it can do |
| paradigm | pattern, approach |
| methodology | method |
| framework (non-technical) | approach, system |
| ecosystem | tools, community |
| pipeline (non-CI/CD) | process, flow |
| workflow (when vague) | steps, process |
| mechanism | way, method |
| artifact (non-technical) | file, output, result |
| deliverable | output, result |
| prerequisite | requirement, what you need first |
| best practice | recommendation |
| use case | example, situation |
| end user | user, reader |
| stakeholder | (name the person or role) |
| bandwidth (non-network) | time, capacity |
<!-- red-pen-ignore-end -->

## Adjectives and adverbs

<!-- red-pen-ignore-start -->
| Do not use | Use instead |
|---|---|
| robust | stable, reliable |
| scalable | (state the actual limit or capacity) |
| seamless | smooth |
| comprehensive | full, complete |
| extensive | large, wide |
| cutting-edge | new, recent |
| state-of-the-art | new, current |
| innovative | new |
| powerful | (delete, or state what it can do) |
| elegant | (delete, or describe the design) |
| intuitive | clear, easy to learn |
| granular | detailed, fine |
| holistic | full, complete |
| pivotal | important |
| crucial | important |
| critical (non-error) | important |
| significant | large, important |
| substantial | large |
| fundamental | basic, core |
| intricate | complex, detailed |
| nuanced | (delete, or explain the detail) |
| myriad | many |
| diverse | different, varied |
| dynamic | (delete, or state what changes) |
| leveraging | using |
| utilizing | using |
| essentially | (delete) |
| basically | (delete) |
| effectively | (delete) |
| actually | (delete) |
| literally | (delete) |
| incredibly | very |
| extremely | very |
| highly | very |
| deeply | (delete or rewrite) |
| fully | (delete or rewrite) |
| truly | (delete) |
| really | (delete) |
| certainly | (delete) |
| definitely | (delete) |
| absolutely | (delete) |
| notably | (delete) |
| importantly | (delete) |
| interestingly | (delete) |
| significantly | (delete or use "by [amount]") |
<!-- red-pen-ignore-end -->

## Phrases

<!-- red-pen-ignore-start -->
| Do not use | Use instead |
|---|---|
| in order to | to |
| in terms of | for, about |
| with respect to | for, about |
| in regard to | about |
| as a matter of fact | (delete) |
| it is worth noting that | (delete, start with the fact) |
| it is important to note that | (delete, start with the fact) |
| it should be noted that | (delete, start with the fact) |
| at the end of the day | (delete) |
| at this point in time | now |
| for the purpose of | to, for |
| in the event that | if |
| on a regular basis | regularly |
| a large number of | many |
| a wide range of | many, various |
| in a timely manner | on time, quickly |
| take into consideration | consider |
| is capable of | can |
| has the ability to | can |
| in close proximity to | near |
| due to the fact that | because |
| serves as a testament to | shows |
| plays a vital role in | (rewrite: X does Y) |
| is a key component of | is part of |
| provides a foundation for | supports |
| helps to ensure | helps |
| with a focus on | for, about |
| on the other hand | but |
| as opposed to | instead of, unlike |
| for all intents and purposes | (delete) |
| a number of | some, several |
| the vast majority of | most |
| in light of | because of |
| with the exception of | except |
<!-- red-pen-ignore-end -->

## Technical terms

These words have a precise technical meaning. The plain swap changes that meaning, so keep the word when it names the technical thing. Use the plain swap only in general prose.

<!-- red-pen-ignore-start -->
| Term | Keep when | Plain swap in general prose |
|---|---|---|
| authenticate | Proving identity with a token, key, or certificate. This is not the same as logging in. | log in, sign in |
| deprecate | Marking something as still present but scheduled for removal. Never swap it for "remove." | (no swap) |
| terminate | A platform uses it as a distinct action. An EC2 terminate deletes the instance. A stop does not. | stop, end |
| optimize | Compiler, SQL, or algorithm performance work: "optimizing the query", "the query optimizer" | improve, speed up |
| dependency | A package, library, or service the code needs | requirement |
| deploy | Shipping code to an environment | install, publish, release |
| provision | Creating infrastructure or accounts through a tool | create, set up |
| configure | Setting options in a config file or tool | set up |
| validate | Checking input against a schema or rules | check |
| verify | Checking a signature, checksum, or certificate | check |
| confirm | A dialog or prompt asks the user to confirm | check |
| implement | Writing code for an interface or spec | set up, add, build |
| implementation | The code behind an interface or spec | setup, code |
| execute | A permission or a mode, such as the execute bit | run |
| instantiate | Creating an instance of a class or type | create |
| invoke | Calling a function, Lambda, or command | call, run |
| persist | Writing state to storage that survives a restart | save, store |
| propagate | Errors, DNS records, or changes spreading through a system | spread, send |
| consume | Reading from a queue, stream, or API | use, read |
| iterate | Looping over a collection | repeat, loop |
| orchestrate | Container or workflow orchestration | manage, run |
| bootstrap | A bootstrap script, cluster, or compiler | set up, start |
| select | A UI control or a SQL statement | click, choose, pick |
| architecture | CPU architecture or system design docs | structure, design |
| infrastructure | Infrastructure as code, or the systems a team runs | servers, systems, setup |
| component | A UI component or a named part of a system | part |
| boilerplate | Template code that every project repeats | template, standard |
| guardrails | Named limits in a system, such as LLM guardrails | limits, rules |
<!-- red-pen-ignore-end -->

## AI slop vocabulary

These words appear in AI-generated text at rates far higher than human writing. Avoid them entirely.

<!-- red-pen-ignore-start -->
| Word | Use instead |
|---|---|
| delve | explore, examine, look at |
| tapestry | (delete or describe what you mean) |
| landscape | area, field, market |
| foster | support, build, encourage |
| testament | proof, evidence, sign |
| underscore | show, highlight |
| multifaceted | complex |
| realm | area, field |
| embark | start |
| beacon | (delete) |
| spearhead | lead |
| groundbreaking | new |
| cornerstone | base, foundation |
| underpinning | base, support |
| game-changer | (delete or state the specific improvement) |
| synergy | (delete) |
| paradigm shift | change |
| deep dive | detailed look |
| unpack | explain |
| double-edged sword | tradeoff |
| low-hanging fruit | easy win |
| move the needle | improve |
| north star | goal |
| pain point | problem |
<!-- red-pen-ignore-end -->
