export default function Weather() {

    function fetchWeather() {
        let data = fetch('http://localhost/api').then((response) => response.json());
        console.log(data);
    }

    return (
        <div>
            <h1>Weather Component</h1>
            <p>This is where the weather information will be displayed.</p>
            <button onClick={fetchWeather}>Click Me</button>
        </div>
    )
}