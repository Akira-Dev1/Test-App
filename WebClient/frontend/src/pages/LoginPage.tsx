import { Card } from "../styles/component/Card";
import { Center } from "../styles/layout/Center";
import { Button } from "../styles/primitives/Button";
import { Stack } from "../styles/primitives/Stack";
import { Text } from '../styles/primitives/Text';

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
      <Card>
        <Stack gap={16}>
          <Text variant="h1">Log in</Text>

          <Button onClick={loginGithub}>
            <Text variant="h2">GitHub</Text>
          </Button>

          <Button onClick={loginYandex}>
            <Text variant="h2">Yandex</Text>
          </Button>

          <Button onClick={loginByCode}>
            <Text variant="h2">Code</Text>
          </Button>
        </Stack>
      </Card>
    </Center>
  );
};

export default LoginPage;
