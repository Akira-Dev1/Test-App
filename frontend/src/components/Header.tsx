import { useState } from "react";
import logo from "../assets/logo-dark.png";
import { Link } from "react-router-dom";
import { useAuth } from "../hooks/useAuth";
import { logout, logoutAll } from "../auth/authService";
import { Header } from "../styles/component/Header";
import { Logo } from "../styles/primitives/Logo";
import { Button } from "../styles/primitives/Button";
import { Text } from "../styles/primitives/Text";
import { Overlay, Modal } from "../styles/component/ShowModal";

const brandStyle : React.CSSProperties = {
  display: "flex",
  alignItems: "center",
  gap: "10px",
  textDecoration: "none",
  color: "#fff",
  cursor: "pointer",
}

const titleStyle : React.CSSProperties = {
  fontSize: "18px",
  fontWeight: 600,
}

const HeaderBar: React.FC = () => {
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
    <Header >
      <Link to="/" style={brandStyle}>
        <Logo
          src={logo} 
          alt="Logo"
        />
        <span style={titleStyle}>TestApp</span>
      </Link>

      {showLogoutModal && (
        <Overlay>
          <Modal>
            <Text>
              Log out of the system?
            </Text>

            <Button onClick={handleLogoutThis}>
              Only from this device
            </Button>

            <Button onClick={handleLogoutAll}>
              From all device
            </Button>

            <Button onClick={() => setShowLogoutModal(false)}>
              cancel
            </Button>
          </Modal>
        </Overlay>
      )}
    </Header>
  );
};

export default HeaderBar;
