# Harbor Take Home Project

Welcome to the Harbor take home project.

## The Challenge (refined from REQUIREMENTS.md)

Building a REST API for calendly. Features supported:

 - User can register (with unique `username`)
 - CRUD for users (did not include auth for users since that is not the focus of this assignment)
 - Some basic validations on user creation and updates and propagate DB errors to api responses (input validation / uniqueness / create,update ts) - didn't dive too much into it though
 - User can set their own availability, two ways to set availability
   - general availability for each day of the week
   - override availability on particular dates
 - Meetings are stored separately with a join table with users - to enable support for group meetings in the future
 - Endpoints for setting the general availability and datewise override for a user
 - User can see their schedule between given dates (can extend it in the future to fetch availability for 1 week, 1 month and so on)
 - Given two users, can get their schedule overlap
 - Comprehensive api docs (swagger) are generated and you can interact with the app at: [localhost swagger](http://localhost:2090/api/docs/index.html)


## Running the app

 - App is deployed at:
 - Use swagger docs to try out APIs, there should be some sample users already there (GET /users)
 - App is containerized, so can run locally as well
 - Clone the repo/branch locally and invoke `docker compose up`

## Assumptions and trade offs

 - Did not include support for timezone in MVP, all times are stored in UTC
 - Basic Timezone support can be added at FE layer by using the timezone field in user profile
 - No auth of any kind
 - 

## Expectations

We care about

- Have you thought through what a good MVP looks like? Does your API support that?
- What trade-offs are you making in your design?
- Working code - we should be able to pull and hit the code locally. Bonus points if deployed somewhere.
- Any good engineer will make hacks when necessary - what are your hacks and why?

We don't care about

- Authentication
- UI
- Perfection - good and working quickly is better

It is up to you how much time you want to spend on this project. There are likely diminishing returns as the time spent goes up.

