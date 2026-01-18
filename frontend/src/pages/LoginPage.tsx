import styled from "styled-components";
import { Card } from "../styles/component/Card";
import { Center } from "../styles/layout/Center";
import { Button } from "../styles/primitives/Button";
import { Stack } from "../styles/primitives/Stack";
import { Text } from '../styles/primitives/Text';

const CardLogin = styled(Card)`
  width: 40%;
`

const LoginPage = () => {
  const loginGithub = () => {
    window.location.href = "/login?type=github";
  };

  const loginYandex = () => {
    window.location.href = "/login?type=yandex";
  };

  const loginByCode = () => {
    window.location.href = "/login?type=code";
  };

  return (
    <Center>
      <CardLogin>
        <Stack gap={16}>
          <Text variant="h1">Log in</Text>

          <Button onClick={loginGithub}>
            GitHub
          </Button>

          <Button onClick={loginYandex}>
            Yandex
          </Button>

          <Button onClick={loginByCode}>
            Code
          </Button>
        </Stack>
      </CardLogin>
    </Center>
  );
};

export default LoginPage;
