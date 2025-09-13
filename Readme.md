# Weather App

## A microservice-based application that fetches weather data for [St. Louis, MO]("https://www.weather.gov/lsx/")
This application vite/react as the frontend framework and build tools. The compiled html is served via a go handler in ./apps/backend/internal/server/server.go

## Application Functionality
The application uses go to send a http GET request to the [NWS API]("https://api.weather.gov") for the latest data from the KSTL station.
The application then serves the recieved data via the apiIndex go http handler.
The compile frontend file uses [React.js]("https://react.dev) useState library to update the Weather component with the current temperature.

## TO-DO
- Style the frontend
- Build out the fetch function to retrieve more data from the api
- Build out database connection
- Add deploy logic and unit tests
/- Modify the api request url to fetch data in degrees farenheight, rather than degrees celcius