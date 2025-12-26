import { useEffect, useState } from "react";
import { getHealth } from "./api/health";

function App() {
  const [health, setHealth] = useState("loading...");

  useEffect(() => {
    getHealth().then(setHealth).catch(() => setHealth("error"));
  }, []);
  useEffect(() => {
    console.log("VITE_API_URL =", import.meta.env.VITE_API_URL);
  }, []);
  return <h1>Backend status: {health}</h1>; 
}

export default App;
