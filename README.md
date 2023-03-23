This project implements a backend server script with Go using gorilla/mux for routing for an online book store. The Book struct is defined as shown in the example below:-

    {
        "id": "1",
        "isbn": "1341234asdfasdg",
        "title": "SecondBook",
        "author": {
            "firstname": "Michael",
            "Lastname": "Faraday"
        }
    }

The APIs implemented are :-
1. CREATE : Server recieves a new JSON Book object and stores it in cache
2. READ ALL : Server gives a list of all Book objects in the cache
3. READ BY ID : Server gets the JSON object of an existing Book from cache identified by "id", else returns an empty JSON object
4. DELETE BY ID : Deletes an instance of the Book object from cache identified by the given "id" parameter
5. UPDATE : Replaces an existing JSON object identified by "id" with the new JSON object provided in the request


API testing is conducted and verified through postman
