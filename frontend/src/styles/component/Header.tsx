import styled from "styled-components";
import { theme } from "../theme";

export const Header = styled.header`
  position: "fixed";
  top: "0";
  left: "0";
  height: "56px";
  width: "100%";
  backgroundColor: ${theme.colors.background};
  zIndex: "100";

  display: "flex";
  alignItems: "center";
  justifyContent: "center";
`