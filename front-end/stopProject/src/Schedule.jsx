import React, { useState } from 'react';
import axios from 'axios';
import "bootstrap/dist/css/bootstrap.min.css";

function Schedule() {
  const [stopData, setStopData] = useState(null);
  const [routeData, setRouteData] = useState(null);
  const [selectedShape, setSelectedShape] = useState(null);
  const [selectedRoute, setSelectedRoute] = useState(null);
  const [stopList, setStopList] = useState(null);
  const [selectedStopindex, setSelectedStopindex] = useState(null);
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
          setRouteData(response.data.routeSchedules);
          setSelectedShape(response.data.routeSchedules[0].shapeId);
          setSelectedRoute(response.data.routeSchedules[0]);
          setStopList(response.data.routeSchedules[0].routeInfo[0].stopInfo);
          //console.log(stopList)
          setStopData(null)
        })
        .catch(function (error) {
          console.log(error);
        });
    }
  };

  const handleSelectChange = (e) => {
    const shapeId = e.target.value;
    const route = routeData.find(route => route.shapeId === shapeId);
    const stops = route.routeInfo[0].stopInfo
    setSelectedShape(shapeId);
    setSelectedRoute(route);
    setStopList(stops)
    console.log(selectedRoute)
  };

  const handleStopClick = (index) => {
    setSelectedStopindex(index)
    console.log(index)
  }

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
      
      {stopData && (
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
      )} 

      {routeData && (
              <div className='container'>
                
                <select className='form-select mb-3 w-auto' onChange={handleSelectChange} value={selectedShape}> 
                {routeData.map((routeSchedule) => (
                  <option key={routeSchedule.shapeId} value={routeSchedule.shapeId}>
                    {routeSchedule.routeLongName}
                  </option>
                ))}
              </select>

              <div className='row'>
                <div className='col'>
              {stopList.map((stop, index) => (
                <dl key={index} onClick={() => handleStopClick(index)}>
                <dt>{stop.stopName}</dt>
                </dl>

              ))}
                </div>

              {selectedStopindex !== null && (
                <div className='col'>
                  <div></div>
                  {selectedRoute.routeInfo.map((routeInfo) => (
                    <div>
                      <div>{routeInfo.workDays.join(" ")}</div>
                      <div>{routeInfo.stopInfo[selectedStopindex].departureTime.join(", ")}</div>
                    </div>
                  ))}
                </div>
              )}
              </div>
                       
              {/* {selectedRoute.routeInfo.map((routeInfo, i) => (
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
              ))} */}
            </div>
            )}

    </div>
  );
}

export default Schedule;
