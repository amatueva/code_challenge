# Search

This project will allow search of both the Tickets and Users datasets. A search
will be done on a particular field. Exact equality matching of values is all
that is needed (eg: searching for "Tim" would not return "Timothy").

## Running the tests

The test suite uses Ruby's built in Test::Unit. There is a script that will
execute the tests:

    ./scripts/run-tests
  
## Data

There are two data files. Tickets have an assigned user via the `assignee_id`
field on the ticket. Please include associated data in the results.
