import styled from "styled-components";
import { theme } from "../theme";

export const Footer = styled.footer`
  position: "fixed";
  bottom: 0;
  left: 0;
  height: "32px";
  width: "100%";
  backgroundColor: ${theme.colors.background};

  display: "flex";
  alignItems: "center";
  justifyContent: "center";
`