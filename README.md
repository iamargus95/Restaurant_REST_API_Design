# Problem Statement

## Design and implement a backend API system for Restaurants. Think Zomato v0.1
- At a high level, you need to implement following APIs
1. Signup API to sign up a new API client
2. Allow valid clients to retrieve a list of restaurants
3. Allow valid clients to retrieve a restaurant’s menu items

## User Stories:
- Here are the above use cases in a user story format with additional supporting details.

### Story_1: Signup
- This API is the main entrypoint for the application. A user must first sign up before calling any other APIs. As part of this API, you will accept USERemail, name and phone number of the user who wants to use our restaurant APIs.
- After validating that email and phone is correct (use basic validations for these fields), you need to store this
information in database. 
        - If email is not valid (ex. foobar) or phone is not valid (ex. doesn’t have 10 digits), then return appropriate error message along with appropriate http code to the user.
        - If all fields are valid, return an auth-token (32 char GUID string) to the user in the response body. The user is supposed to use this field in the header for all future requests. Auth token GUIDs can be stored where you see fit for this exercise.

### Story_2: Clients are able to retrieve the restaurants from our system
- The user should be able to fetch list of restaurants via an API. 
> Note that, only the valid users (with matching auth-token from signup) can access this API.
- The token is passed as a X-Auth-Token header by the client.
- The request may optionally have a “page” parameter, ex. ?page=2 as part of the URL. This is used for paginating the restaurant records. You can sort the restaurants by their rating when returning a list to the users.
- The other requirements for this API are as follows:
    1. Validate auth token, if not valid, return appropriate error message along with http code to the user.
    2. If auth token is valid, return a list of restaurants to the user.
    3. Use “page” param to return appropriate page of the records. Assume we return 10 records per page. If no page param is passed, just return the first page of restaurants.
    4. A restaurant has following data elements
        - ID
        - Name
        - Description
        - A star rating between 0.0 and 5.0 (i.e. a decimal value ranging from 0.0 to 5.0)
        - Address

### Story_3: Clients are able to retrieve a restaurant’s menu items
- The user should be able to fetch menu items of a single restaurant via an API.
> Note that, only the valid users (with matching auth-token from signup) can access this API.
- The token is passed as a X-Auth-Token header by the client.
- The request must have a restaurant ID as a parameter. The request may optionally have a “page” parameter, ex. ?page=2 as part of the URL. 
- This is used for paginating the restaurant menu items. You can sort the menu items by their category when returning a list to the users (more on it below).
- The other requirements for this API are as follows:
    - 1. Validate auth token, if not valid, return appropriate error message along with http code to the user.
    - 2. If auth token is valid, return a list of menu items for that restaurant to the user.
    - 3. Use “page” param to return appropriate page of the records. Assume we return 10 records per page. If
    no page param is passed, just return the first page of menu items.
    - 4. A Restaurant Menu Item has following data elements
        - ID
        - Name
        - Description
        - Menu category (one of following)
            - Starter, Side-dish, Main-course, Beverage, Dessert
        - Price
    - 5. When returning menu items, sort them by their menu category. i.e. all starters should come first, then
    all side-dishes, followed by main-course, beverage and then dessert.

## Evaluation criteria
- Use any programming language (Java, Golang, Ruby, Python, JavaScript etc) and any relevant andpopular web framework library of your choice
- Write modular code with unit tests for controllers, and other components.
- Use a DB migration tool (ex. Flyway, liquibase, Golang migrate, ActiveRecord, etc) as per language of your choice.
- This should be used to set up the DB and add dummy data in the DB
> Assume if something is unclear and document the assumptions in the readme
> Share a GitHub link to your code. Make small and incremental progress when writing code. Ensure to have multiple commits along the way as you complete individual stories instead of a single giant commit.
- Use Postman for API testing and demoing the application features.
- Bonus points if you can create a docker based setup for deploying this API locally.
- Feel free to use a relevant DB of your choice (relational or document) ex. Postgres, Mysql or MongoDB etc.
- Bonus points if you can upload a quick 5 min recording demonstrating the API usage via Postmant.
- Include a link to this recording in the readme.