import { useState } from "react";
import logo from "../assets/logo-dark.png";
import { Link } from "react-router-dom";
import { useAuth } from "../hooks/useAuth";
import { logout, logoutAll } from "../auth/authService";

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

const overlayStyle: React.CSSProperties = {
  position: "fixed",
  top: 0,
  left: 0,
  width: "100vw",
  height: "100vh",
  backgroundColor: "rgba(0,0,0,0.4)",
  display: "flex",
  alignItems: "center",
  justifyContent: "center",
  zIndex: 1000,
};

const modalStyle: React.CSSProperties = {
  background: "#fff",
  padding: 20,
  borderRadius: 8,
  display: "flex",
  flexDirection: "column",
  gap: 10,
};

const Header: React.FC = () => {
  const { refreshAuth } = useAuth();
  const [showLogoutModal, setShowLogoutModal] = useState(false);

  const handleLogoutThis = async () => {
    try {
      await logout();
      await refreshAuth();
    } finally {
      setShowLogoutModal(false);
    }
  };

  const handleLogoutAll = async () => {
    try {
      await logoutAll();
      await refreshAuth();
    } finally {
      setShowLogoutModal(false);
    }
  }

  return (
    <header style={headerStyles} >
      <Link to="/" style={brandStyle}>
        <img
          src={logo} 
          alt="Logo"
          style={logoStyles}
        />
        <span style={titleStyle}>TestApp</span>
      </Link>

      {showLogoutModal && (
        <div style={overlayStyle}>
          <div style={modalStyle}>
            <h3>
              Log out of the system?
            </h3>

            <button onClick={handleLogoutThis}>
              Only from this device
            </button>

            <button onClick={handleLogoutAll}>
              From all device
            </button>

            <button onClick={() => setShowLogoutModal(false)}>
              cancel
            </button>
          </div>
        </div>
      )}
    </header>
  );
};

export default Header;
