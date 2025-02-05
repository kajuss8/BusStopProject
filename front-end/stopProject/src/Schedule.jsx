import React, { useState } from 'react';
import axios from 'axios';

function Schedule() {
  const [stopData, setStopData] = useState(null);
  const [routeData, setRouteData] = useState(null);
  const [stopId, setStopId] = useState('');
  const [routeId, setRouteId] = useState('');
  
  const handleStopIdInput = (event) => {
    setStopId(event.target.value);
  };

  const handleRouteInput = (event) => {
    setRouteId(event.target.value);
  };

  const handleButtonClickStopId = () => {
    if (stopId) {
      axios.get(`http://localhost:8080/StopSchedle/${stopId}`)
        .then(function (response) {
          console.log("Hello");
          setStopData(response.data.stopSchedule);
        })
        .catch(function (error) {
          console.log(error);
        });
    }
  };

  const handleRouteInputRouteId = () => {
    if (routeId) {
      axios.get(`http://localhost:8080/RouteSchedule/${routeId}`)
        .then(function (response) {
          console.log("Hello");
          console.log(response.data)
          setRouteData(response.data.routeSchedules);
        })
        .catch(function (error) {
          console.log(error);
        });
    }
  };




  return (
    <div>
        <div>
            <input type="text" value={stopId} onChange={handleStopIdInput} placeholder="Enter stop ID" />
            <button onClick={handleButtonClickStopId}>Get Schedule</button>
            <input type="text" value={routeId} onChange={handleRouteInput} placeholder="Enter stop ID" />
            <button onClick={handleRouteInputRouteId}>Get Schedule</button>
        </div>
      
      {stopData ? (
        <div>
          <h1>{stopData.stopName}</h1>
          {stopData.stopInformation.map((info, index) => (
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


      {routeData ? (
              <div>
                {routeData.map((routeSchedule, _) => (
                  <div> {routeSchedule.routeLongName} 
                        {routeSchedule.shapeId}
                        {routeSchedule.routeInfo.map((routeInfo, _) => (
                          <div>
                            <div>{routeInfo.workDays}</div>
                            {routeInfo.stopInfo.map((stopInfo, _) => (
                              <div>
                                <div>{stopInfo.stopName}</div>
                                <div>{stopInfo.departureTime.join(`, `)}</div>
                              </div>
                            ))}
                          </div>
                        ))}
                  </div>
                  
                ))}
               
              </div>
            ) : (
              <p>Press the button to load stop information.</p>
            )}


    </div>

    
    
  );
}

export default Schedule;
