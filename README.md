# effective-octo-computing-machine - My Cohere API Project

This project provides a simple RESTful API interface for interacting with the Cohere AI model using go.

<img width="1070" alt="api_demo" src="https://github.com/user-attachments/assets/b28b738f-2406-4367-bd47-f15cc4f4eb19">


## Getting Started

### Prerequisites

- Go 1.23.2 or later
- An API key from Cohere

### Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/my-cohere-api.git
   cd THIS-RESPOS-NAME
   ```

2. Create a `.env` file in the root of the project directory. This file should contain your Cohere API key. 

   Example `.env` file:
   ```plaintext
   # This is an example .env file. Below is the way you should place your COHERE_API_KEY
   # Assume your COHERE_API_KEY is "123456"
   COHERE_API_KEY=123456
   ```

### Default Model

- The project defaults to the **Command R+ 08-2024** model from Cohere, which cannot be changed in the provided code.

### Running the Application

To start the server, use:

```bash
make start
```

### API Endpoint

- **POST** `/chat`
  - Sends a message to the AI model and receives a response.


### Testing with Postman

To test the API using Postman, follow these steps:

1. Open Postman and create a new request.
2. Set the request type to **POST**.
3. Enter the URL: `http://localhost:8080/chat`.
4. Go to the **Headers** tab and set the **Content-Type** to `application/json`.
5. In the **Body** tab, select the **raw** option and choose **JSON** from the dropdown menu.
6. Enter the JSON payload, for example:

   ```json
   {
     "message": "Hello, AI!"
   }

   ```

7. Send the request and observe the response.

<img width="1050" alt="api_post" src="https://github.com/user-attachments/assets/1563cc76-326e-47ad-8c88-28f234a1bfd1">

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
