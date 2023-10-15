Take home assignment for a job interview with Redia
10/10-23

Over-engineered calculator:

Must have:
1. Create a basic calculator with history.
2. There should be a clear separation of logic.
3. All logic should be covered by unit tests.
4. The calculator should expose a RESTful interface deployed via a backend technology fx Firebase, but a docker container on eg. Heroku will also suffice or something else you are familar with.
5. The interface should be documented with an accompanying Postman package for easy testing.
6. All source code should be available via your personal Github account.

Nice to have:
1. Best effort micro-service architecture (given the timeframe).
2. Auth for using the service (email/password will suffice).
3. A small webpage utilizing the calculator.

As the title says, it is an over-engineered calculator, which means:
You should go all-in on design patterns and best practices as well as making sure to fulfill the "must have" requirements.
The coding language is also up to you, but GO lang would be preferred, since it is part of our tech stack.


version2 contains the following changes:
- a postloghandler has been created to handle updating the logs
- expression handler sends a logpost instead of accessing our logs directly
- you can click on a given log to revisit it
- i started implementing a login solution
    - can create users stored in the servers runtime
    - hashes passwords and salt
    - compares hashes to check if pw is right
    - logging in creates a jwt that is returned to the client
    - cannot create a user with a username that is already in use
    - authorization is not implemented
    - storing and accessing logs with a specific user is not implemented
    - email verification is not implemented
    - testing is unfortunately non-existent