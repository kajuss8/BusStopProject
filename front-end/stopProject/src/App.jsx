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

  return (
    <div>
      <input type="text" value={id} onChange={handleInputChange} placeholder="Enter stop ID" />
      <button onClick={handleButtonClick}>Get Schedule</button>
      {data ? (
        <div>
          <h1>{data.stopName}</h1>
          {data.stopInformation.map((info, index) => (
            <div key={index} >
              <div className='Contaner'>
                <div className="row p-3 ">
                  <div className="col-1 p-3 ">{info.routeType} {info.routeShortName}</div>
                  <div className='col-3 p-3'>{info.routeLongName}</div>
                  <div className="col p-3">{info.workDays.join(', ')}</div>
                  <div className=''>{info.arrivalTimes.join(', ')}</div>
                </div>
              </div>
              
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
