# Harbor Take Home Project

Welcome to the Harbor take home project.

## The Challenge (refined from REQUIREMENTS.md)

Building a REST API for calendly. Features supported:

 - User can register (with unique `username`)
 - CRUD for users (did not include auth for users since that is not the focus of this assignment)
 - Some basic validations on user creation and updates and propagate DB errors to api responses (input validation / uniqueness / create,update ts) - didn't dive too much into it though
 - **(Initial requirement)** User can set their own availability, two ways (endpoints) to *set availability*
   - general availability for each **day of the week**
     - User can set any length time slots on each day
     - If no time slots on a day, means user is not available
     - Input is given as array of time strings for each day that is interpreted inside 24 hour time window
     - FE can have a nice display and can enforce some consistency in selection box like time slots of min size 15 min or always starting at the hour mark but BE can support arbitrary start/end times
   - override availability on **particular dates**
     - User can *override* their availability on particular dates using this
     - Complete schedule for this date is overridden
   - I did it this way to avoid creating a new row for every date's availability. This ensures that most users will just have 7 rows for day availability and only separate rows for specific DATES overridden.
 - User can *see* their availability between given dates (can extend it in the future to fetch availability for 1 week, 1 month and so on)
 - Given two users, can get their *schedule overlap*
 - **Note:** All input and output timestamps are assumed to be in UTC
 - Since all timestamps stored in DB are in UTC, future timezone logic can be on the app layer so that even if user travels to different timezone, their preferences can be correctly interpreted
 - Comprehensive api docs (swagger) are generated and you can interact with the app at: [localhost swagger](http://localhost:2090/api/docs/index.html)


 - **Additional features: (for later)** Meetings can also be supported, users can book meetings with another after first checking the availability (validation left out for now)
 - Meetings are stored along with a join table for many-to-many relationship with the users - to enable support for group meetings in the future
 - Obviously, existing user meetings will also be taken into consideration while getting user availability or schedule overlap

## Screenshots: API docs and DB
<img width="1512" alt="calendly-api-docs" src="https://github.com/user-attachments/assets/5e2f1398-08f4-40bd-91f0-6347f32bf579" />

<img width="1202" alt="calendly-api-db" src="https://github.com/user-attachments/assets/6c0df257-eb78-418b-bb57-fb62cc171231" />


## Running the app

 - App is containerized, so can run locally as well
 - Clone the repo/branch locally and invoke `docker compose up`
 - Access app at [localhost swagger](http://localhost:2090/api/docs/index.html)
 - Didn't spend time on deploying it for now since running locally is straightforward

## Assumptions and trade offs

 - Did not include support for timezone in MVP, assuming all users are in same timezone
 - No auth of any kind since it is not the focus right now
 - All timestamps are stored in UTC - this ensures consistency and easy to extend the logic to support timezones in future
 - IMP: /availability/ endpoints for day/date/user should ideally get user info from jwt token but right now, it's through username in request body which does an extra DB call - hack to avoid auth for now
 - Schedule/availability endpoints expect slots info (minutes since midnight) directly for now, so more work for FE but can later be extended so that BE does the processing and can also support seconds since midnight (instead of minutes) 
 - No pagination in REST endpoints yet
 - In storing availabilities, we do not support the seconds precision since it's not very useful
 - Day boundaries not handled for now in meetings (availability includes cross day support), will need to do some extra handling. Currently, it can be handled by splitting the meeting slot across two days or storing start/end dates for meetings


## Future work

 - User authentication and authorization
 - Basic Timezone support can be added at FE layer by using the timezone field in user profile
 - Validation for timezone field in user object
 - Validation in input slots to avoid duplication or overlap of input slots
 - Pagination in GET user/users/availability endpoints
 - More nuanced support for fetching own availability and overlap between schedules
 - Caching layer for improving latencies
   - Can cache each users's availability in memcache/redis (update it when user updates their availability or a new meeting is created for this user)
 - DB indexes
 - Meeting reminder jobs that send email notifications before the meeting
 - Have *smart meetings* feature like Google calendar - leave buffer time at the end of the meetings by reducing the duration
 - Support to modify/delete a meeting
 - Support to modify availability times - both the general day-wise and specific date overrides
 - Support recurring meetings (daily/weekly/monthly)

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

