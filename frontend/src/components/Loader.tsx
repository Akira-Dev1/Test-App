const Loader = () => {
  return (
    <div style={loaderStyle}>
      <div style={spinnerStyle} />
      <p>loading...</p>
    </div>
  );
};

const loaderStyle: React.CSSProperties = {
  display: "flex",
  //flexDirection: "column",
  alignItems: "center",
  justifyContent: "center",
  height: "100vh",
  width: "100vh",
};

const spinnerStyle: React.CSSProperties = {
  width: 40,
  height: 40,
  border: "4px solid #ccc",
  borderTop: "4px solid #333",
  borderRadius: "50%",
  animation: "spin 1s linear infinite",
};

export default Loader;