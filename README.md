# Search

This project will allow search of both the Tickets and Users datasets. A search
will be done on a particular field. Exact equality matching of values is all
that is needed (eg: searching for "Tim" would not return "Timothy").

## Runnning the app

To run the app execute following command

    make

## Running the tests

The test suite uses Golang's built in Testing library.

    make test
  
## Data

There are two data files. Tickets have an assigned user via the `assignee_id`
field on the ticket. Please include associated data in the results.

### Issues to fix
- Implement polymorphism
- Write tests
- Print associated data
- Fix graceful shutdown