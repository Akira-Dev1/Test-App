import Header from "../components/Header";
import Footer from "../components/Footer";
import { Outlet } from "react-router-dom";

const mainStyle : React.CSSProperties = {
  paddingTop: "56px",
  paddingBottom: "32px",
};

const MainLayout = () => {
  return (
    <>
      <Header />
      <main style={mainStyle}>
        <Outlet />
      </main>
      <Footer/>
    </>
  );
};

export default MainLayout;
