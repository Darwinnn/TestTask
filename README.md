# Test Task

Application for processing the incoming requests from the 3d-party providers 

## Prerequisite
1. [Docker](https://www.docker.com/get-started)
2. uuidgen (optionally) - part of the linux-utils and comes by default in MacOS too. Used to generate test data
3. [go-swagger](https://github.com/go-swagger/go-swagger) (optionally) - to regenerate the code

## Run
To run the build locally you would need to do two things:

1. `make up` - bootstraps two docker containers, the app itself and the database
2. `make migrage-up seq=2` - runs migrations in the database (Uses [Go Migrate](https://github.com/golang-migrate/migrate/pulse) in docker)

After that, the app should be up and running, listening for external requests on http://localhost:8080

## App structure
### Config
The configuration is done via enviroment variables, all of them are mandatory:
* `DB_CONN_STRING` - database connection string
* `T_MINUTES` - run transaction cancellation every n minutes
* `T_ODD_NUM` - how many last odd transaction to cancel
> See docker-compose.yml for examples

### API overview
The app is build with [go-swagger](https://github.com/go-swagger/go-swagger) which generates all boilerplate code (including all validity checks) specified in **swagger.yml**
> You can regenarate the code with `make generate`

#### API Docs
App serves swagger.json (for swagger UI) on `/swagger.json`

#### Routes
App accepts POST requests on `/api/v1/state` URI, with json body like so:
```json
{"amount": "1.1", "state": "win", "transactionId": "UUID"}
```
win states increase the balance to the specified amount, lost states decrease it. 

### Database structure
Database has two tables to maintain its job:
1. **balances** - stores a balance value and its ID
2. **transactions** - stores transactions, UUIDs, states and amounts (has **constraint** on UUID column, so we won't end up processing same transactionID more than once (`"transactions_uuid_key" UNIQUE CONSTRAINT, btree (uuid)`)

`transactions` relates to `balances` thru `transactions.balance_id=balances.id`

### Workers overview
The app has a worker called **canceller**, which runs every `T_MINUTES` and cancels last `T_ODD_NUM` odd transactions, correcting the balance.

### Health checks
GET `/health` returns 200 if the app is working correctly.

## Test data
There's a simple bash script in the `testdata` directory which generates 20 random transactions (**requires** `genuuid` util)
## Unit test
`make test` bootstraps a testing db, migrations and runs unit tests

# Thoughts 
1. App stores balance in `double` type in database, which is not perfect for storing billing data due to [accuracy problems](https://en.wikipedia.org/wiki/Floating-point_arithmetic#Accuracy_problems)
2. If you run more than one copy of the app within one database, canceller workers may conflict. 
> I would use something like redis distlock to tell other copies not to touch the data when one of them already works on it
> Or [postgres explicit locks](https://www.postgresql.org/docs/11/explicit-locking.html)

3. The canceller worker corrects balance in a for loop, resulting in unnecessary load on database (and increased time). The same can be done with a single query using [CTEs](https://www.postgresql.org/docs/11/queries-with.html), which would probably work way faster. I decided not to do it, since it's basically hiding bussiness logic in the database. The query could be something like this:
```sql
WITH canceling_transactions AS
  (UPDATE transactions
   SET canceled=TRUE
   FROM
     (SELECT id,
             amount,
             state,
             balance_id
      FROM transactions
      WHERE (NOT canceled)
        AND id%2!=0
      ORDER BY id DESC
      LIMIT 10) t
   WHERE transactions.id=t.id RETURNING t.*)
UPDATE balances
SET value= (CASE
                WHEN t.state='win'
                THEN value - t.amount
                ELSE value + t.amount
            END)
FROM canceling_transactions t
WHERE balances.id=t.balance_id
```