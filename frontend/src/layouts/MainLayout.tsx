import HeaderBar from "../components/Header";
import FooterBar from "../components/Footer";
import { Page } from "../styles/layout/Page"
import { Outlet } from "react-router-dom";

const MainLayout = () => {
  return (
    <>
      <HeaderBar />
      <Page>
        <Outlet />
      </Page>
      <FooterBar/>
    </>
  );
};

export default MainLayout;
