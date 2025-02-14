import React, { useState, useEffect } from "react";
import axios from "axios";
import "bootstrap/dist/css/bootstrap.min.css";
import MapComponent from "./MapComponent.jsx";

function Schedule() {
  const [allRoutesData, setAllRoutesData] = useState(null);
  const [filteredTypes, setFilteredTypes] = useState(null);
  const [stopData, setStopData] = useState(null);
  const [routeData, setRouteData] = useState(null);
  const [selectedShape, setSelectedShape] = useState(null);
  const [selectedRoute, setSelectedRoute] = useState(null);
  const [stopList, setStopList] = useState(null);
  const [selectedStopindex, setSelectedStopindex] = useState(0);
  const [routeTypeName, SetRouteTypeName] = useState(null);
  const [stopId, setStopId] = useState("");
  const [stopIdInput, setStopIdInput] = useState("");
  const [routeId, setRouteId] = useState("");
  const [routeIdInput, setRouteIdInpute] = useState("");
  const [coordinates, setCoordinates] = useState(null);

  const handleAllTypes = () => {
    setFilteredTypes(allRoutesData);
    SetRouteTypeName("Autobusai ir Troleibusai");
  };

  const handleBusFileter = () => {
    const bus = allRoutesData.filter(
      (route) => route.routeTransportType === "A"
    );
    setFilteredTypes(bus);
    SetRouteTypeName("Autobusai");
  };

  const handleTrolFileter = () => {
    const trol = allRoutesData.filter(
      (route) => route.routeTransportType === "T"
    );
    setFilteredTypes(trol);
    SetRouteTypeName("Troleibusai");
  };

  const handleStopIdInput = (e) => {
    setStopIdInput(e.target.value);
  };

  const handleStopIdButtonClick = () => {
    setStopId(stopIdInput);
  };

  const handleRouteIdInput = (e) => {
    setRouteIdInpute(e.target.value);
  };

  const handleRouteIdButtonClick = () => {
    setRouteId(routeIdInput);
  };

  const handleStopClick = (index) => {
    setSelectedStopindex(index);
  };

  useEffect(() => {
    handleAllRoutesData();
  }, []);

  useEffect(() => {
    handleButtonClickStopId();
    setStopId();
  }, [stopId]);

  useEffect(() => {
    handleRouteInputRouteId();

    setRouteId();
  }, [routeId]);

  const handleAllRoutesData = async () => {
    await axios
      .get(`http://localhost:8080/AllRoutes`)
      .then(function (response) {
        setAllRoutesData(response.data.routesData);
        setFilteredTypes(response.data.routesData);
        SetRouteTypeName("Autobusai ir Troleibusai");
      })
      .catch(function (error) {
        console.log(error);
      });
  };

  const handleButtonClickStopId = async () => {
    if (stopId) {
      await axios
        .get(`http://localhost:8080/StopSchedle/${stopId}`)
        .then(function (response) {
          setStopData(response.data.stopSchedule);
          setRouteData(null);
          if (response) {
            setAllRoutesData();
          }
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
          setRouteData(response.data.routeSchedules);
          setSelectedRoute(response.data.routeSchedules[0]);
          setStopList(response.data.routeSchedules[0].routeInfo[0].stopInfo);
          setSelectedStopindex(0);
          setStopData(null);
          if (!selectedShape) {
            setSelectedShape(response.data.routeSchedules[0].shapeId);
          }
          if(response){
            setAllRoutesData();
            setCoordinates(response.data.routeSchedules[0].routeInfo[0].stopInfo)
          }
        })
        .catch(function (error) {
          console.log(error);
        });
    }
  };

  const handleSelectChange = (e) => {
    if (routeData) {
      const shapeId = e.target.value;
      const route = routeData.find((route) => route.shapeId === shapeId);
      const stops = route.routeInfo[0].stopInfo;
      setCoordinates(stops)
      setSelectedShape(shapeId);
      setSelectedRoute(route);
      setStopList(stops);
      setSelectedStopindex(0);
    }
  };

  const groupeTimes = (times) => {
    const groupedTimes = {};
    const resultArray = [];
    times.forEach((time) => {
      const hour = time.substring(0, 2);
      if (!groupedTimes[hour]) {
        groupedTimes[hour] = [];
        resultArray.push(groupedTimes[hour]);
      }
      groupedTimes[hour].push(time);
    });
    return resultArray;
  };

  return (
    <div>
      <div className="container p-3">
        <div className="row justify-content-center">
          <div className="col-auto me-4">
            <input
              type="text"
              value={stopIdInput}
              onChange={handleStopIdInput}
              placeholder="Stop ID"
            />

            <button onClick={handleStopIdButtonClick}>Get Stop Schedule</button>
          </div>
          <div className="col-auto">
            <input
              type="text"
              value={routeIdInput}
              onChange={handleRouteIdInput}
              placeholder="Route ID"
            />
            <button onClick={handleRouteIdButtonClick}>
              Get Route Schedule
            </button>
          </div>
        </div>
      </div>

      {allRoutesData && (
        <div className="container">
          <div className="d-flex justify-content-center">
            <div
              className="btn-group btn-group-lg"
              role="group"
              aria-label="Large button group"
            >
              <input
                type="radio"
                className="btn-check"
                name="options"
                id="option1"
                autoComplete="off"
                defaultChecked
                onChange={handleAllTypes}
              />
              <label className="btn btn-outline-primary" htmlFor="option1">
                Autobusai ir Troleibusai
              </label>

              <input
                type="radio"
                className="btn-check"
                name="options"
                id="option2"
                autoComplete="off"
                onClick={handleBusFileter}
              />
              <label className="btn btn-outline-primary" htmlFor="option2">
                Autobusai
              </label>

              <input
                type="radio"
                className="btn-check"
                name="options"
                id="option3"
                autoComplete="off"
                onClick={handleTrolFileter}
              />
              <label className="btn btn-outline-primary" htmlFor="option3">
                Troleibusai
              </label>
            </div>
          </div>
          {filteredTypes && (
            <div>
              <h3>{routeTypeName}</h3>
              <table className="table">
                <tbody>
                  {filteredTypes.map((route, index) => (
                    <tr key={index}>
                      <td className="col-4">
                        {route.routeShortName}
                        <a
                          className="col-4 p-3 link-dark hover-light link-offset-2 link-underline link-underline-opacity-0"
                          href="#"
                          onClick={(e) => {
                            setRouteId(route.routeId);
                            e.preventDefault();
                          }}
                        >
                          {route.routeLongName}
                        </a>
                      </td>
                      <td className="col-4">{route.workDays.join(" ")}</td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          )}
        </div>
      )}

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
                          <a
                            className="col-4 p-3 link-dark hover-light link-offset-2 link-underline link-underline-opacity-0"
                            href="#"
                            onClick={(e) => {
                              setSelectedShape(info.shapeId);
                              setRouteId(info.routeId);
                              e.preventDefault();
                            }}
                          >
                            {info.routeLongName}
                          </a>
                          <div className="col p-3">
                            {info.workDays.join(" ")}
                          </div>
                        </th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr>
                        <td>
                          <div >{info.arrivalTimes.join(", ")}</div>
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
        <div className="" >
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

          <div className="row ">
            <div className="col-auto ">
              {stopList.map((stop, index) => (
                <dl
                  key={index}
                  onClick={() => {handleStopClick(index);}}
                  className={`p-2 ${
                    selectedStopindex === index ? "bg-light" : ""
                  }`}
                >
                  <dt>
                    <a
                      href="#"
                      onClick={(e) => {e.preventDefault()}}
                      onDoubleClick={() => setStopId(stop.stopId)}
                      className="link-dark hover-light link-offset-2 link-underline link-underline-opacity-0 "
                    >
                      {stop.stopName}
                    </a>
                  </dt>
                </dl>
              ))}
            </div>

            {selectedStopindex !== null && (
              <div className="col-auto ">
                <div className="row d-flex flex-wrap  justify-content-between flex-row-reverse">
                  {selectedRoute.routeInfo.map((routeInfo, routeIndex) => (
                    <div key={routeIndex} className="col">
                      <table className="table table-bordered text-cente">
                        <thead>
                          <tr>
                            <th>{routeInfo.workDays.join(", ")}</th>
                          </tr>
                        </thead>
                        <tbody>
                          {groupeTimes(
                            routeInfo.stopInfo[selectedStopindex].departureTime
                          ).map((time, timeIndex) => (
                            <tr key={timeIndex}>
                              <td>{time.join(", ")}</td>
                            </tr>
                          ))}
                        </tbody>
                      </table>
                    </div>
                  ))}
                </div>
              </div>
            )}
            <div className="col">
                <MapComponent coordinates={coordinates} selectedIndex={selectedStopindex}/>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}

export default Schedule;
