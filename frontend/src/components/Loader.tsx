import styled from "styled-components";
import { Center } from "../styles/layout/Center";
import { Stack } from "../styles/primitives/Stack";

const Spinner = styled.div`
  width: 40;
  height: 40;
  border: "4px solid #ccc";
  borderTop: "4px solid #333";
  borderRadius: "50%";
  animation: "spin 1s linear infinite";
`;

const Loader = () => {
  return (
    <Center>
      <Stack>
        <Spinner />
      </Stack>
    </Center>
  );
};

export default Loader;
