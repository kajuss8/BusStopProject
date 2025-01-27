import React, { useState } from 'react';
import axios from 'axios';

interface StopSchedule {
  stopName: string;
  stopInformations: StopInformation[];
}

interface StopInformation {
  routeShortName: string;
  routeLongName: string;
  routeType: string;
  calendarWorkDays: number[];
  arrivalTimes: string[];
}

function App() {
  const [data, setData] = useState<StopSchedule>();
  const [id, setId] = useState('');

  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setId(event.target.value);
  };

  const handleButtonClick = () => {
    if (id) {
      axios.get(`http://localhost:8080/StopSchedle/${id}`)
        .then(function (response) {
          console.log("Hello");
          console.log(data)
          setData(response.data as StopSchedule);
        })
        .catch(function (error) {
          console.log(error);
        });
    }
  };
  
  return (
    <div>
      <input type="text" value={id} onChange={handleInputChange} placeholder="Enter stop ID" />
      <button onClick={handleButtonClick}>Get Schedule</button>
          {data && data.stopInformations && (
      <div>
        <h1>{data.stopName}</h1>
        {data.stopInformations.map((info, index) => (
          <div key={index}>
            <p>Route Short Name: {info.routeShortName}</p>
            <p>Route Long Name: {info.routeLongName}</p>
            <p>Route Type: {info.routeType}</p>
            <p>Calendar Work Days: {info.calendarWorkDays.join(', ')}</p>
            <p>Arrival Times: {info.arrivalTimes.join(', ')}</p>
          </div>
        ))}
      </div>
    )}
    </div>
  );
}

export default App;