import styled from "styled-components";

export const Overlay = styled.div`
	position: "fixed";
  top: 0;
  left: 0;
  width: "100vw";
  height: "100vh";
  backgroundColor: "rgba(0,0,0,0.4)";
  display: "flex";
  alignItems: "center";
  justifyContent: "center";
  zIndex: 1000;
`
export const Modal = styled.div`
	background: "#fff",
  padding: 20,
  borderRadius: 8,
  display: "flex",
  flexDirection: "column",
  gap: 10,
`