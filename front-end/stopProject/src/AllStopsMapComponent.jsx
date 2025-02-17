import React, { useState, useEffect, useRef } from "react";
import { MapContainer, TileLayer, Marker, Popup} from "react-leaflet";
import "leaflet/dist/leaflet.css";

const AllStopsMapComponent = ({ coordinates, searchItem }) => {

const mapRef = useRef(null);

const customIcon = L.icon({
    iconUrl: "/record-button.png",
    iconSize: [15, 15], 
});

const filteredStops = searchItem
? coordinates.filter(stop => stop.stopName.toLowerCase().includes(searchItem.toLowerCase()))
: coordinates;

useEffect(() => {
if (filteredStops.length > 0 && mapRef.current) {
    const { stopLat, stopLon } = filteredStops[0];
    mapRef.current.setView([stopLat, stopLon], 13);
}
}, [filteredStops]);

return (
    <MapContainer
        center={[54.8905, 23.927]}
        zoom={13}
        style={{ height: "90vh", width: "100%" }}
        whenCreated={(map) => (mapRef.current = map)}
    >
        <TileLayer url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png" />
        {filteredStops.map((coord, index) => (
            <Marker
                key={index}
                position={[coord.stopLat, coord.stopLon]}
                icon={customIcon}
            >
                <Popup><strong>Stoptelė:</strong> {coord.stopName}</Popup>
            </Marker>
        ))}
    </MapContainer>
);
};

export default AllStopsMapComponent;