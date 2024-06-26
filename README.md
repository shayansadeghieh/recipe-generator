# GenAI Recipe Generator

As someone who struggles with digestive issues, finding recipes that meet my dietary needs has always been a challenge. To address this, I developed a GenAI application that generates nutritious recipes tailored to specific dietary requests, complete with ingredients and instructions.

## Demo
![Demo](assets/recipe-generator.gif)


## Tech Stack 

This application is an LLM + RAG application. 

- I used Cohere's Command model + some prompt engineering for the bot.
- I used Cohere's embed-english-v3.0 embedding model to embed the user's request. I used the same model for generating embeddings for the vector DB.
- I used pinecone as my vector DB. Although I used pinecone's go SDK for the creation of Pinecone indexes, their SDK was lacking sufficient documentation, so I created my own pinecone client to query the DB.  
- My database of recipes includes some of my favourites I've found on the internet over the years. 

## Project Structure 
The project is split up into several packages:

`main`: The main entrypoint to the application. It initializes the DB client, cohere client and exposes port 8080 to listen to requests.

`dao`: The data access object package. This package does all things related to vector DBs including creating Pinecone indexes and querying Pinecone indexes. 

`handler`: The handler package. This package handles the incoming API request and routes the request through the appropriate logic.

`wire`: The wire package. This package holds the structure about the data coming in and out of the API. 

## TODO
- Integration tests. This application is primarily API calls linked together. Although I do like mocking, for something as simple as this, I find it to be overkill.