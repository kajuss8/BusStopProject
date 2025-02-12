import React, { useEffect, useState } from "react";
import { MapContainer, TileLayer, Marker, Popup } from "react-leaflet";
import "leaflet/dist/leaflet.css";

const MapComponent = ({coordinates}) => {
  return (
    <MapContainer center={[coordinates.lat, coordinates.lng]} zoom={13} style={{ height: "500px", width: "100%" }}>
      <TileLayer url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png" />
      <Marker position={[coordinates.lat, coordinates.lng]}>
        <Popup>Location: {coordinates.lat}, {coordinates.lng}</Popup>
      </Marker>
    </MapContainer>
  );
};

export default MapComponent;