# **Work Completed for Sprint 2:**

#### Frontend:
- Display destination information after search
- Linked user login and create user pages with backend, allowing information to be put in the database
- Page routing (i.e., a user will be routed to the login page after completing account creation)

#### Backend:
- Created a function to store a list of saved destinations
- Created a function to save and delete destinations from a list of saved destinations
- Created a function to return a user's saved destinations list
- Set up a new table in the database to store users' saved destinations
- Finished user login and create user functions which handle several situations (i.e. wrong password, email taken, etc.)
- Created tests to validate all functions linked with frontend

# **Unit Tests:**

#### Frontend:

#### Backend:
- create user function test
- login user function test
- get destination function test
- database connection test


# **Backend Documentation:**

## Overview:
This API is designed to store and access user login information in a database and contact a Yelp API to return information about travel destinations, sorting through the Yelp API output to return only relevant information. In addition, users can save locations which they are interested in, and may view and edit this list by adding to it or deleting from it. Generally, this API provides users the ability to search for new travel destinations and save their favorites for future reference.

## Dependencies: 
The only additional dependency this API requires for all of it's features to function properly is the addition of a .env file containing a private Yelp API key. A contributor must be contacted regarding this for security reasons.

## Yelp API integration:
This API integrates with the Yelp API to provide information to users about shopping, restaurants, and entertainment near a searched destination. Currently, the top 10 most highly rated restaurants, stores, and entertainment locations are returned, displaying their rating along with the average price, address, and the type of establishment (i.e. "French Restaurant").

## Endpoints:
### **POST** /createUser

### **POST** /loginUser

### **GET** /newDestination/{location}

### **GET** /updateDestination

### **PUT** /updateDestination

### **DELETE** /updateDestination

- create user: details on what it does, what it needs, what it returns, how to link with it


put places a desination in a list with a user email associated with it for later searching
delete takes a location associated with an email out
get finds all of the locations associated with a particular email and returns them


# **Demo Video:**
This is a link to a video detailing the current functionality of our integrated application:
https://drive.google.com/file/d/1FVE4IiN6Fo1rgTe7MVeshZ89Fl90Uffk/view?usp=sharing

