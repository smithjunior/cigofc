[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=smithjunior_gitopsfc&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=smithjunior_gitopsfc)
[![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=smithjunior_gitopsfc&metric=duplicated_lines_density)](https://sonarcloud.io/summary/new_code?id=smithjunior_gitopsfc)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=smithjunior_gitopsfc&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=smithjunior_gitopsfc)
# O que é?

É o processo de integrar modificacões do codebase de forma contínua e automatizada,
evitando assim erros humanos de verificacão, garantindo mais agilidade e seguranca no processo de desenvolvimento de um software.

## Principais processos

- execucao de testes
- linter
- verificacões de qualidade de código
- verificacões de seguranca
- geracao de artefatos prontos para o processo de deploy
- identificacão da próxima versão a ser gerada no software
- geracão de tags e release

## Status check

É a garantia de que uma Pull Request não poderá ser mergeanda ao repositório sem antes ter passado pelo processo de CI ou mesmo no processo de Code Review

## Ferramentas Populares

- Jenkins
- Github Actions
- Circle CI
- AWS Code Build
- Azure DevOps
- Google Cloud Build
- GitLab Pipelines / CI

## Dinâmica

### Workflow

- São conjuntos de processos definidos por você. Ex: Fazer o Build + rodar os testes da aplicacão.
- É Possível ter mais do que um workflow por repositório.
- Definidos em arquivos ".yml" em: .github/workflows
- Possui um ou mais "Jobs"
- É iniciado baseado em eventos do GitHub ou através de agendamentos

Evento    ->   Filtros      ->    Ambiente        ->     Acões
on: push        branches:         runs-on: ubuntu       steps:
                - main                                    - uses: action/run-processor
                                                          - run: npm run prod

### Actions

É a acão que de fato será executada em um dos steps de um job em um workflow.
Ela pode ser criada do zero ou ser reutilizada de actions pre-existentes.
