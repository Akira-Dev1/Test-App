import styled from "styled-components";
import { Card } from "../styles/component/Card";
import { Center } from "../styles/layout/Center";
import { Stack } from "../styles/primitives/Stack";
import { Text } from "../styles/primitives/Text";

const CardHome = styled(Card)`
  width: 60%;
`

const HomePage = () => {
  return (
    <Center>
      <CardHome>
        <Stack gap={16}>
          <Text variant="h1">HomePage</Text>
        </Stack>
      </CardHome>
    </Center>
  );
};

export default HomePage;