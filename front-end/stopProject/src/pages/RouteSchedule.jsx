import React, { useState, useEffect } from "react";
import axios from "axios";
import { Link } from "react-router-dom";
import MapComponent from "../MapComponent.jsx";
import { useLocation } from "react-router-dom";
import { use } from "react";

function RouteSchedule() {
    const [routeData, setRouteData] = useState(null);
    const [selectedShape, setSelectedShape] = useState(null);
    const [selectedRoute, setSelectedRoute] = useState(null);
    const [stopList, setStopList] = useState(null);
    const [selectedStopindex, setSelectedStopindex] = useState(0);
    //const [routeId, setRouteId] = useState("");
    const location = useLocation();
    const routeId = location.state.routeId;
    const [coordinates, setCoordinates] = useState(null);

    useEffect(() => {
          handleRouteApi();
    }, [routeId]);
    
    const handleStopClick = (index) => {
        setSelectedStopindex(index);
    };

    const handleRouteApi = async () => {
        if (routeId) {
          await axios
            .get(`http://localhost:8080/RouteSchedule/${routeId}`)
            .then(function (response) {
              setRouteData(response.data.routeSchedules);
              setSelectedRoute(response.data.routeSchedules[0]);
              setStopList(response.data.routeSchedules[0].routeInfo[0].stopInfo);
              setSelectedStopindex(0);
              console.log("hello");
              if (!selectedShape) {
                setSelectedShape(response.data.routeSchedules[0].shapeId);
              }
              if(response){
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

    return(
    <div className="container-fluid mt-4">
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

          <div className="row mt-4">
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
                <MapComponent coordinates={coordinates} selectedIndex={selectedStopindex} onMarkerClick={handleStopClick}/>
            </div>
          </div>
        </div>
      )}
    </div>
    );



}

export default RouteSchedule;