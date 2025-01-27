import React, { useState, useEffect } from 'react';
import axios from 'axios';

function App() {
  const [data, setData] = useState(null);
  const [id, setId] = useState('');

  const handleInputChange = (event) => {
    setId(event.target.value);
  };

  const handleButtonClick = () => {
    if (id) {
      axios.get(`http://localhost:8080/StopSchedle/${id}`)
        .then(function (response) {
          console.log("Hello");
          setData(response.data.stopSchedule);
        })
        .catch(function (error) {
          console.log(error);
        });
    }
  };

  useEffect(() => {
    if (data) {
      console.log(data); // Logs the updated value of `data`
    }
  }, [data]); // Run whenever `data` changes

  
  return (
    <div>
      <input type="text" value={id} onChange={handleInputChange} placeholder="Enter stop ID" />
      <button onClick={handleButtonClick}>Get Schedule</button>
      {data ? (
        <div>
          <h1>{data.stopName}</h1>
          {data.stopInformation.map((info, index) => (
            <div key={index}>
              <p>Route Short Name: {info.routeShortName}</p>
              <p>Route Long Name: {info.routeLongName}</p>
              <p>Route Type: {info.routeType}</p>
              <p>Calendar Work Days: {info.workDays.join(', ')}</p>
              <p>Arrival Times: {info.arrivalTimes.join(', ')}</p>
            </div>
          ))}
        </div>
      ) : (
      <p>Press the button to load stop information.</p>
      )}
    </div>
  );
}

export default App;
