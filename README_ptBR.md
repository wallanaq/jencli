# JenCLI

JenCLI é uma ferramenta de linha de comando (CLI) projetada para simplificar interações com Jenkins e Keycloak, permitindo autenticação, gerenciamento de sessão e tarefas de automação em um ambiente DevOps.

## Funcionalidades

- **Login e Logout**: Autentique-se usando o Keycloak e gerencie sessões.
- **Gerenciamento de Configurações**: Armazene e recupere facilmente as configurações do CLI.
- **Integração com Jenkins**: Automatize tarefas e acesse recursos do Jenkins com facilidade.

## Instalação

Para instalar, clone o repositório e compile o binário.

```bash
git clone https://github.com/wallanaq/jencli.git
cd jencli
go build -o jencli
```

## Uso

### Autenticação

- **Login**: Inicia o fluxo de autenticação no Keycloak.

```bash
jencli login
```

- **Logout**: Encerra a sessão do Keycloak para o usuário autenticado.

```bash
jencli logout
```

## Configuração

As configurações são armazenadas localmente em cada sistema operacional nos diretórios padrão:

Windows: `%APPDATA%\sicoobcli\db`
Linux/Mac: `~/.config/jencli/db`

## Desenvolvimento

Para contribuir ou modificar este CLI:

1. Faça um fork do repositório.
2. Faça as suas alterações e envie um pull request.

## Licença

Este projeto é licenciado sob a licença MIT. Consulte o arquivo LICENSE para mais detalhes.
