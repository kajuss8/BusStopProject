import React, { useState, useEffect } from "react";
import axios from "axios";
import { Link, Route } from "react-router-dom";
import AllStopsMapComponent from "../AllStopsMapComponent.jsx";

function Home() {
  const [allRoutesData, setAllRoutesData] = useState(null);
  const [allStopsData, setAllStopsData] = useState(null);
  const [filteredTypes, setFilteredTypes] = useState(null);
  const [routeTypeName, SetRouteTypeName] = useState(null);
  const [searchItem, setSearchItem] = useState("");
  const [searchResult, setSearchResult] = useState(null);
  const displayData = searchItem ? searchResult : filteredTypes;

  useEffect(() => {
    handleAllRoutesData();
  }, []);

  const handleInputChange = (e) => {
    const searchTerm = e.target.value;
    setSearchItem(searchTerm);

    if (allStopsData) {
      const searchResults = allStopsData.filter((stop) =>
        stop.stopName.toLowerCase().includes(searchTerm.toLowerCase())
      );
      setSearchResult(searchResults);
    } else {
      const searchResults = filteredTypes.filter(
        (route) =>
          route.routeLongName.toLowerCase().includes(searchTerm.toLowerCase()) ||
          route.routeShortName.toLowerCase().includes(searchTerm.toLowerCase())
      );
      setSearchResult(searchResults);
    }
  };

  const handlesCityStopMap = () => {
    handleAllStopsData();
    setSearchItem("")
  };

  const handleAllTypes = () => {
    setAllStopsData(null);
    setFilteredTypes(allRoutesData);
    SetRouteTypeName("Autobusai ir Troleibusai");
    setSearchItem("");
  };

  const handleBusFileter = () => {
    setAllStopsData(null);
    const bus = allRoutesData.filter(
      (route) => route.routeTransportType === "A"
    );
    setFilteredTypes(bus);
    SetRouteTypeName("Autobusai");
    setSearchItem("");
  };

  const handleTrolFileter = () => {
    setAllStopsData(null);
    const trol = allRoutesData.filter(
      (route) => route.routeTransportType === "T"
    );
    setFilteredTypes(trol);
    SetRouteTypeName("Troleibusai");
    setSearchItem("");
  };

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

  const handleAllStopsData = async () => {
    await axios
      .get(`http://localhost:8080/Stops`)
      .then(function (response) {
        setAllStopsData(response.data.stopsData);
        console.log(response.data.stopsData);
      })
      .catch(function (error) {
        console.log(error);
      });
  };

  return (
    <div className="container mt-5">
      {allRoutesData && (
        <div>
          <div className="text-center ">
            <div className="row ">
              <div className="col ">
                <div className="input-group mb-3 ">
                  <span className="input-group-text" id="inputGroup-sizing-default">
                    Search
                  </span>
                  <input
                    value={searchItem}
                    type="text"
                    className="form-control"
                    aria-label="Sizing example input"
                    aria-describedby="inputGroup-sizing-default"
                    onChange={handleInputChange}
                  />
                </div>
              </div>

              <div className="col">
                <div
                  className="btn-group btn-group-lg text-nowrap"
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

                  <input
                    type="radio"
                    className="btn-check"
                    name="options"
                    id="option4"
                    autoComplete="off"
                    onClick={handlesCityStopMap}
                  />
                  <label className="btn btn-outline-primary" htmlFor="option4">
                    žemėlapis
                  </label>
                </div>
              </div>
            </div>
          </div>
          {allStopsData ? (
            <AllStopsMapComponent coordinates={allStopsData} searchItem={searchItem}/>
          ) : (
            displayData && (
              <div>
                <h3>{routeTypeName}</h3>
                <table className="table">
                  <tbody>
                    {displayData.map((route, index) => (
                      <tr key={index}>
                        <td className="col-4">
                          {route.routeShortName}
                          <Link
                            className="col-4 p-3 link-dark hover-light link-offset-2 link-underline link-underline-opacity-0"
                            to="/RouteSchedule"
                            state={{ routeId: route.routeId }}
                          >
                            {route.routeLongName}
                          </Link>
                        </td>
                        <td className="col-4">{route.workDays.join(" ")}</td>
                      </tr>
                    ))}
                  </tbody>
                </table>
              </div>
            )
          )}
        </div>
      )}
    </div>
  );
}

export default Home;
