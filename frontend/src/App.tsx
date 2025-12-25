import { useEffect, useState } from "react";
import { getHealth } from "./api/health";

function App() {
  const [health, setHealth] = useState("loading...");

  useEffect(() => {
    getHealth().then(setHealth).catch(() => setHealth("error"));
  }, []);

  return <h1>Backend status: {health}</h1>;
}

export default App;
