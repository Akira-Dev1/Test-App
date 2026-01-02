import { useEffect, useState } from "react";

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
  return (
    <footer style={footerStyle}>
      <p>Kvartirka 31</p>
    </footer>

  );
}

export default Footer;