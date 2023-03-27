Command Output HTTP API

This is a simple HTTP REST API built using Go that accepts a shell command and returns the output of that command.

Installation

To install and run the application, follow these steps:

    Install Go from the official website.

    Clone the repository or download the source code.

    Navigate to the project directory in a terminal.

    Run the following command to install the necessary dependencies:
        go get -d ./...
    Build the application using the following command:
        go build -o cmdapi
    Run the application using the following command:
        ./cmdapi
By default, the application will listen on port 8080. If you want to use a different port, you can set the PORT environment variable before starting the application:

export PORT=9000
./cmdapi

Usage

The application provides a single endpoint, /api/cmd, that accepts a POST request with a shell command as a query parameter or in the JSON body. The application will execute the command and return the output as a response.

If the command is not found, the application will return an error with a 404 status code.
Request Format

The request body should be in JSON format with the following structure:
    {
        "cmd": "echo 'Hello, World!'"
    }

Alternatively, you can provide the command as a query parameter:
    POST /api/cmd?cmd=echo 'Hello, World!'

Response Format

The response body will contain the output of the command:

If the command is not found, the response body will contain an error message:
    Command not found

Error Handling

If the application encounters an error while executing the command, it will return an error message with a 500 status code.

If the request body is invalid or missing the cmd field, the application will return an error message with a 400 status code.

Testing

To run the unit tests, navigate to the project directory in a terminal and run the following command:

go test -v ./...

This will run all tests in the project and print the results to the console.


