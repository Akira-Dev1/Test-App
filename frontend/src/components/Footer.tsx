import { useEffect, useState } from "react";
import { getHealth } from "../api/health";

const footerStyle : React.CSSProperties = {
  position: "fixed",
  bottom: 0,
  left: 0,
  height: "32px",
  width: "100%",
  backgroundColor: "#1e1e1e",

  display: "flex",
  alignItems: "center",
  justifyContent: "center",
}

const Footer : React.FC = () => { 
  const [health, setHealth] = useState("loading...");

  useEffect(() => {
    getHealth().then(setHealth).catch(() => setHealth("error"));
  }, []);

  return (
    <>
      <footer style={footerStyle}>
        <p>Backend status: {health}</p>
      </footer>
    </>
  );
}

export default Footer;