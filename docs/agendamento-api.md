```http
### Realiza novo agendamento

POST http://api.example.com/agendamentos
Content-Type: application/json
Authorization: Bearer hdsjkhfsdjkkldslkfdsjskdfljsdflweewew

{
  "paciente":{
    "cpf":"12345678900"
    "nome":"Edgar Nishi-Myojo",
    "contato":[
      {
        "tipo":"fone",
        "value":"5511999999999",
      },
      {
      "tipo":"email",
        "value":"edmyojo@gtakiyaki.com"
      }
    ]
  },
  
  "profissional":{
    "crm":"12345SP"
    "nome":"Dr. Bruno Caderninho"
  },
  
  "agendamento":{
    "data":"2022-11-14T16:00:00.0000",
    "sala":""
  }
}

### Listagem de agenda do profissional

GET http://api.example.com/agendamentos?data=&crm=
Accept: application/json
Authorization: Bearer hdsjkhfsdjkkldslkfdsjskdfljsdfl


>> 
{

 "agendamentos":[
  "id":"adhdlfkjsdfkljsdfkljsdflk",
  "paciente":{
    "cpf":"12345678900"
    "nome":"Edgar Nishi-Myojo",
    "contato":[
      {
        "tipo":"fone",
        "value":"5511999999999",
      },
      {
      "tipo":"email",
        "value":"edmyojo@gtakiyaki.com"
      }
    ]
  },
  
  "profissional":{
    "crm":"12345SP"
    "nome":"Dr. Bruno Caderninho"
  },
  
  "agendamento":{
    "data":"2022-11-14T16:00:00.0000",
    "sala":""
  }
 ]


}

```
