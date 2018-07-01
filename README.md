# Payment Gateway Golang Example

In our simplified model of a prepaid card the card holds a balance in GBP and users can make
transactions in GBP. To be able to use the card the user must first load some money onto the
card. This increases the available balance on the card. When a user goes and makes a
transaction using the card, the merchant (e.g your local independent coffee shop) sends the
card an Authorization request to check if the user loaded enough money on the account to
pay for their coffee.
 
If the user has loaded enough, the Authorization request is approved and
the amount the merchant requested is earmarked (or blocked). The merchant at a later point in
the future can decide to Capture ​the transaction at which point we send the merchant the
money. At any point the user should be able to see how much he loaded onto the card, the
current available balance, and the current amount that is blocked on there (waiting to be
captured).

The merchant can decide to only capture part of the amount or capture the amount multiple
times. In this model, the merchant can’t capture more than we initially authorized him to. The
merchant can decide to reverse the whole or part of the initial Authorization at which point they
can no longer capture the full amount (only the amount that is still authorized). The merchant
can Refund ​the user after they capture the funds. They can’t refund the user more than they
captured. The user can then use the refunded amount to buy more coffee.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

What things you need to install the software and how to install them

```
https://golang.org/doc/install
```

### Installing

A step by step series of examples that tell you how to get a development env running

Download the GO Project

```
https://github.com/angelospillos/gateway.git
```

Place the gateway folder in the MAIN GoPath src to look like

```
go/src/gateway/main.go
```

Then use the following command to run it

```
go run *.go
```

## Running the tests

Do to time constrains there are no automated tests for this system
but you can use the following endpoints to test the system

### Accounts Create
It will generate a random Account.

Endpoint

```
GET http://127.0.0.1:8000/v1/accounts/create
```

Response

```
{
    "id": "4666336021280850",
    "available": 0,
    "blocked": 0,
    "deposited": 0,
    "withdrawn": 0,
    "currency": "GBP",
    "card_name": "Mr Payment",
    "card_type": "VISA",
    "card_number": 4666336021280850,
    "card_expiry_month": 2,
    "card_expiry_year": 23,
    "card_security_code": 547,
    "statement": null,
    "creation_time": "2018-07-02T00:31:04.673878069+03:00"
}
```

### Accounts Deposit

Use the the Id of an Account to deposit.

Request

```
 {
 	"id" : "4666336021280850",
 	"amount":1000
 }
```

Endpoint

```
POST http://127.0.0.1:8000/v1/accounts/deposit
```

Response
```
{
    "account_id": "4666336021280850",
    "status": "0",
    "description": "Approved"
}
```

### Accounts Detail

Details of the Current Account

Request

```

{
	"id" : "4666677990633308"
}

```

Endpoint

```
http://127.0.0.1:8000/v1/accounts/detail
```

Response
```
{
    "id": "",
    "available": 0,
    "blocked": 0,
    "deposited": 0,
    "withdrawn": 0,
    "currency": "",
    "card_name": "",
    "card_type": "",
    "card_number": 0,
    "card_expiry_month": 0,
    "card_expiry_year": 0,
    "card_security_code": 0,
    "statement": null,
    "creation_time": "0001-01-01T00:00:00Z"
}
```

### Accounts Statement

A list of all payments involved with the current Account

Request
```
{
	"id" : "4666744292130045"
}
```

Endpoint
```
http://127.0.0.1:8000/v1/accounts/statement
```

Response
```
{
    "statement": [{}]
}
```

### Payments Authorization

Authorization is the first point of entry to start working with payment operations.
Once an authorization is successful you can execute successive operations.


Headers
```
From 4666557263649072
Content-Type application/json
```

Request
```
{
	"orderId" : "123",
	"amount":900,
	"currency":"GBP",
	"card_name":"Mr Payment",
	"card_number":4666677990633308,
	"card_expiry_month": 10,
	"card_expiry_year" : 19,
	"card_security_code" : 916
}
```

Endpoint

```
POST http://127.0.0.1:8000/v1/payments/authorization

```

Response
```
{
    "referenceId": "9da2df89-f0e4-406d-ac0d-d7670607e306",
    "status": "0",
    "description": "Approved"
}
```

### Successive Requests
Headers
```
From 4666557263649072 // Account Id to get the money
Content-Type application/json
```

Example Request
```
{
	"orderId" : "123",
	"amount":50
}
```

Endpoint
```
http://127.0.0.1:8000/v1/payments/capture/{authorization_id}
http://127.0.0.1:8000/v1/payments/reversal/{authorization_id}
http://127.0.0.1:8000/v1/payments/refund/{capture_id}
```

Example Response
```
{
    "referenceId": "ee1a3371-1781-4de9-bf7c-f72394c811ba",
    "status": "0",
    "description": "Approved"
}
```

## Deployment

To Deploy this on a live system follow http://www.blog.labouardy.com/deploying-go-app-to-aws-ec2/

## Built With

* [Golang](https://golang.org/) - The Go Programming Language.
* [MUX](https://github.com/gorilla/mux) - A powerful URL router and dispatcher for golang.
* [BoltDb](https://github.com/boltdb/bolt) - An embedded key/value database for Go.

## Versioning

For the versions available, see the [tags on this repository](https://github.com/angelospillos/gateway/tags). 

## Authors

* **Angelos Pillos** - (https://www.angelospillos.com)

See also the list of [contributors](https://github.com/angelospillos/gateway/graphs/contributors) who participated in this project.

## License

This project is licensed under the GNU Affero General Public License v3.0 - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* Special thanks for the inspiration to Curve (https://www.imaginecurve.com/)
