# vacation-planner

## Environment setup

You need to have [Go](https://golang.org/)
[Node.js](https://nodejs.org/) and,
[Air](https://github.com/cosmtrek/air).

Verify the tools by running the following commands:

```sh
go version
npm --version
air -v
```
## Start in development mode

Navigate to the `server` folder 

```zsh
cd server
```
Start the backend using air for live reload

'''zsh
air
'''

The back end will serve on http://localhost:8080.

Navigate to the `webapp` folder, install dependencies,
and start the front end development server by running:

```sh
cd webapp
npm install
npm start
```
The application will be available on http://localhost:3000.
