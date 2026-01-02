import React from "react";
import logo from "../assets/logo-dark.png";
import { Link } from "react-router-dom";

const headerStyles : React.CSSProperties = {
  position: "fixed",
  top: "0",
  left: "0",
  height: "56px",
  width: "100%",
  backgroundColor: "#1e1e1e",
  zIndex: "100",

  display: "flex",
  alignItems: "center",
  justifyContent: "center",

};

const brandStyle : React.CSSProperties = {
  display: "flex",
  alignItems: "center",
  gap: "10px",
  textDecoration: "none",
  color: "#fff",
  cursor: "pointer",
}

const logoStyles : React.CSSProperties = {
  width: "32px",
  height: "32px",
  objectFit: "contain",
}

const titleStyle : React.CSSProperties = {
  fontSize: "18px",
  fontWeight: 600,
}

const Header: React.FC = () => {
  return (
    <header style={headerStyles} >
      <Link to="/" style={brandStyle}>
       <img src={logo} alt="Logo" style={logoStyles} />
        <span style={titleStyle}>TestApp</span>
      </Link>
    </header>
  );
};

export default Header;
