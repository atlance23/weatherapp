import { useState } from 'react'

export default function Weather() {

    let isWeatherSet = false;
    const [weather, setWeather] = useState(isWeatherSet ? weather : "Click the button to fetch weather!");
    const [units, setUnits] = useState(isWeatherSet ? " °F" : "");

    function fetchWeather() {
        let res = fetch('http://localhost/api')
        .then((response) => response.json())
        .then((data) => {
            isWeatherSet = true;
            setWeather(data.temp_f);
            setUnits(" °F");
        });
    }

    return (
        <div>
            <h1>Weather Component</h1>
            <p>This is where the weather information will be displayed.</p>
            <button onClick={fetchWeather}>Click Me</button>
            <h3>{`${weather}${units}`}</h3>
        </div>
    )
}