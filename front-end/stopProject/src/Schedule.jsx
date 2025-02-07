import React, { useState } from "react";
import axios from "axios";
import "bootstrap/dist/css/bootstrap.min.css";

function Schedule() {
  const [stopData, setStopData] = useState(null);
  const [routeData, setRouteData] = useState(null);
  const [selectedShape, setSelectedShape] = useState(null);
  const [selectedRoute, setSelectedRoute] = useState(null);
  const [stopList, setStopList] = useState(null);
  const [selectedStopindex, setSelectedStopindex] = useState(null);
  const [stopId, setStopId] = useState("");
  const [routeId, setRouteId] = useState("");

  const handleStopIdInput = (event) => {
    setStopId(event.target.value);
  };

  const handleRouteInput = (event) => {
    setRouteId(event.target.value);
  };

  const handleButtonClickStopId = async () => {
    if (stopId) {
      await axios
        .get(`http://localhost:8080/StopSchedle/${stopId}`)
        .then(function (response) {
          console.log("helo")
          setStopData(response.data.stopSchedule);
          setRouteData(null);
        })
        .catch(function (error) {
          console.log(error);
        });
    }
  };

  const handleRouteInputRouteId = async () => {
    if (routeId) {
      await axios
        .get(`http://localhost:8080/RouteSchedule/${routeId}`)
        .then(function (response) {
          console.log("hello")
          setRouteData(response.data.routeSchedules);
          setSelectedShape(response.data.routeSchedules[0].shapeId);
          setSelectedRoute(response.data.routeSchedules[0]);
          setStopList(response.data.routeSchedules[0].routeInfo[0].stopInfo);
          setSelectedStopindex(0);
          setStopData(null);
        })
        .catch(function (error) {
          console.log(error);
        });
    }
  };

  const handleSelectChange = (e) => {
    if(routeData){
      const shapeId = e.target.value;
      const route = routeData.find((route) => route.shapeId === shapeId);
      const stops = route.routeInfo[0].stopInfo;
      setSelectedShape(shapeId);
      setSelectedRoute(route);
      setStopList(stops);
    }
  };

  const handleStopClick = (index) => {
    setSelectedStopindex(index);
  };

  return (

    <div>
      <div className="container p-3">
        <div className="row justify-content-center">
          <div className="col-auto me-4">
            <input
              type="text"
              value={stopId}
              onChange={handleStopIdInput}
              placeholder="Enter stop ID"
            />
            <button onClick={handleButtonClickStopId}>Get Stop Schedule</button>
          </div>
          <div className="col-auto">
            <input
              type="text"
              value={routeId}
              onChange={handleRouteInput}
              placeholder="Enter route ID"
            />
            <button onClick={handleRouteInputRouteId}>Get Route Schedule</button>
          </div>
        </div>
      </div>

      {stopData && (
        <div className="container">
          <h1 className="p-4">{stopData.stopName}</h1>
          {stopData.stopInformation.map((info, index) => (
            <div key={index}>
              <div className="container">
                <div className="row p-3">
                <table className="table table-bordered table-striped">
                <thead>
                  <tr>
                  <th className="d-flex align-items-center">
                  <div className="col-1 p-3 ">
                    {info.routeType} {info.routeShortName}
                  </div>
                  <div className="col-4 p-3">{info.routeLongName}</div>
                  <div className="col p-3">{info.workDays.join(" ")}</div>
                  </th>
                  </tr>
                  </thead>
                  <tbody>
                  <tr>
                  <td>
                  <div className="">{info.arrivalTimes.join(", ")}</div>
                  </td>
                  </tr>
                  </tbody>
                  </table>
                </div>
              </div>
            </div>
            
          ))}
        </div>
      )}

      {routeData && (
        <div className="container">
          <select
            className="form-select mb-3 w-auto"
            onChange={handleSelectChange}
            value={selectedShape}
          >
            {routeData.map((routeSchedule) => (
              <option key={routeSchedule.shapeId} value={routeSchedule.shapeId}>
                {routeSchedule.routeLongName}
              </option>
            ))}
          </select>

          <div className="row">
            <div className="col-4 ">
              {stopList.map((stop, index) => (
                <dl key={index} onClick={() => handleStopClick(index)} className={`p-2 ${selectedStopindex === index ? 'bg-light' : ''}`}>
                  <dt ><a href='#' className="link-dark hover-light link-offset-2 link-underline link-underline-opacity-0 ">{stop.stopName}</a></dt>
                </dl>
              ))}
            </div>

            {selectedStopindex !== null && (
              <div className="col">
                <div className="row d-flex justify-content-between flex-row-reverse">
                {selectedRoute.routeInfo.map((routeInfo, routeIndex) => (
                  <div key={routeIndex} className="col">
                    <table className="table table-bordered text-cente">
                      <thead>
                        <tr>
                          <th>
                            {routeInfo.workDays.join(", ")}
                          </th>
                        </tr>
                      </thead>
                      <tbody>
                        {routeInfo.stopInfo[
                          selectedStopindex
                        ].departureTime.map((time, timeIndex) => (
                            <tr key={timeIndex}>
                              <td>{time}</td>
                            </tr>
                        ))}
                      </tbody>
                    </table>
                  </div>
                ))}
                </div>
              </div>
            )}
          </div>
        </div>
      )}
    </div>
  );
}

export default Schedule;
