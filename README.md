# Desafio Fast Beer Development da Qbem

Esse desafio tem como objetivo melhorar a percepção do tempo, planejamento e execução de projetos, Sendo realizado em um tempo
extremamente curto, com tecnologias nuncas utilizadas pelos desenvolvedores e com uma pitada de embreagues.

## Organização do Dia

- *09h00min ás 09h15min*: Alinhamento de como será o hackton
- *09h15min ás 10h00min*: Planejamento do projeto agendamento
- *10h00min ás 11h00min*: Desenvolvimento do Mock da Api e esboço das telas de agendamento e agenda do médico e deploy
- *11h00min ás 12h00min*: Realização dos CRUD dos dados da API ao banco de dados e integração com a API
- *12h00min ás 13h30min*: Almoço
- *13h30min ás 14h00min*: Alinhamento do que foi realizado e próximos passos
- *14h00min ás 15h00min*: Aplicação das regras de negócio na API & Melhoria na experiencia do usuário da aplicação
- *15h00min ás 17h00min*: Refatoração e aplicação de testes unitários faltantes no projeto, melhoria da pipeline
- *17h00min ás 17h30min*: Review com o stakeholders e retro ;)

## O Projeto

Trata-se de desenvolver o front e o backend de uma micro-ecossistema de agendamento seguindo a história do usuário definida.

## Personas para desenvolvimento
```yml

---
Paciente:
  Nome: Edgar Nishi-myojo
  Email: edmyojo@tatsunoko.com
  Whatsapp: 11999559999

---
Médico:
Nome: Bruno Cardeninho
Especialidade: Proctologista
Atendimento: horário comercial
Slot de tempo: 1h

```

## Histórias do usuário

```cucumber
#language: pt-BR

Funcionalidade: Agendamento de consulta

  "Eu como paciente, quero agendar uma consulta com o médico de especialidade da minha escolha através de um site, 
  para que eu possa agendar sem ter que ligar ou sair de casa. 
  Assim podendo escolher um melhor horário para o meu comparecimento."
  
  Cenário: Agendamento realizado com sucesso
    Dado que eu seja um paciente
    E que eu escolha uma especialidade
    E que eu escolha eu médico
    E que eu informe meu nome, email e numero do whatsapp corretamente
    E que eu escolha uma agenda disponível
    Quando eu solicitar o agendamento
    Então apresente uma mensagem de agendamento salvo com sucesso
    E inclua essa data na agenda do médico
    E inclua os meus dados para contato posteriores
    
  Cenário: Agendamento não realizado com sucesso devido a data indisponível
    Dado que eu seja um paciente
    E que eu escolha uma especialidade
    E que eu escolha eu médico
    E que eu informe meu nome, email e numero do whatsapp corretamente
    E que eu escolha uma data
    E que essa data não esteja disponivel
    Quando eu solicitar o agendamento
    Então apresente que essa data não esteja disponível
    E não faça o agendamento
    E destaque outros dias disponiveis
  
  
---
Funcionalidade: Listagem de agendas para o médico

  "Eu, como um médico da Clinica Deus me Acuda, 
  quero uma forma de conferir as agendas que eu tenho para o dia de hoje e para o próximo dia, 
  assim conseguindo organizar-me antes das consultas"
  
  Cenário: Listagem apresentada com sucesso
    Dado que eu seja um medico
    Quando eu solicitar a minha agenda
    Então a lista de consultas de hoje 
    E de manhã
    E destaque outros dias disponiveis

  Cenário: Sem agendamento para ser apresentado
    Dado que eu seja um medico
    E não tenha agendamentos para o dia de hoje e amanhã
    Quando eu solicitar a minha agenda
    Então informe que não há agendamentos para esses dias.

```

