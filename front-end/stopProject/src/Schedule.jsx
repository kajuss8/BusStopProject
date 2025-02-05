import React, { useState } from 'react';
import axios from 'axios';
import "bootstrap/dist/css/bootstrap.min.css";

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

  const handleButtonClickStopId = async () => {
    if (stopId) {
      await axios.get(`http://localhost:8080/StopSchedle/${stopId}`)
        .then(function (response) {
          console.log("Hello");
          setStopData(response.data.stopSchedule);
          setRouteData(null)
        })
        .catch(function (error) {
          console.log(error);
        });
    }
  };

  const handleRouteInputRouteId = async () => {
    if (routeId) {
      await axios.get(`http://localhost:8080/RouteSchedule/${routeId}`)
        .then(function (response) {
          console.log("Hello");
          setRouteData(response.data.routeSchedules);
          setStopData(null)
        })
        .catch(function (error) {
          console.log(error);
        });
    }
  };

  return (
    <div>
        <div className="container p-3">
          <div className="row justify-content-center">
            <div className="col-auto me-4">
              <input type="text" value={stopId} onChange={handleStopIdInput} placeholder="Enter stop ID" />
              <button onClick={handleButtonClickStopId}>Get Schedule</button>
            </div>
            <div className="col-auto">
              <input type="text" value={routeId} onChange={handleRouteInput} placeholder="Enter route ID" />
              <button onClick={handleRouteInputRouteId}>Get Schedule</button>
            </div>
          </div>
        </div>
      
      {stopData ? (
        <div className='container'>
          <h1 className='p-4'>{stopData.stopName}</h1>
          {stopData.stopInformation.map((info, index) => (
            <div key={index} >
              <div className='container'>
                <div className="row p-3">
                  <div className="col-1 p-3 ">{info.routeType} {info.routeShortName}</div>
                  <div className='col-4 p-3'>{info.routeLongName}</div>
                  <div className="col p-3">{info.workDays.join(' ')}</div>
                  <div className=''>{info.arrivalTimes.join(', ')}</div>
                </div>
              </div>
            </div>
          ))}
        </div>
      ) : (
        <p></p>
      )}

      {routeData ? (
              <div className='container'>
                
                <select className='form-select mb-3 w-auto'>
                {routeData.map((routeSchedule, index) => (
                  <option key={index} value={routeSchedule.routeLongName}>
                    <h3>{routeSchedule.routeLongName}</h3>
                  </option>
                ))}
              </select>
                
              {routeData.map((routeSchedule, index) => (
                <div key ={index}> 
                  {routeSchedule.routeInfo.map((routeInfo, i) => (
                    <div key={i}>
                      <div>{routeInfo.workDays.join(", ")}</div>
                      {routeInfo.stopInfo.map((stopInfo, j) => (
                        <div key={j} className='row align-items-center mb-2'>
                          <dl className="col-6">
                          <dt>{stopInfo.stopName}</dt>
                          </dl>
                          <div className="col-6">
                          <div>{stopInfo.departureTime.join(', ')}</div>
                          </div>
                        </div>
                      ))}
                    </div>
                  ))}
                </div>
              ))}
              </div>
            ) : (
              <p></p>
            )}
    </div>
  );
}

export default Schedule;
