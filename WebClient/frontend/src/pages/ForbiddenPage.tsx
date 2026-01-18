import { Link } from "react-router-dom";
import { Text } from "../styles/primitives/Text";
import { Card } from "../styles/component/Card";
import { Center } from "../styles/layout/Center";
import { Stack } from "../styles/primitives/Stack";

const ForbiddenPage = () => {
  return (
    <Center>
      <Card>
        <Stack gap={16}>
          <Text variant="h1">403 — No access</Text>

          <Text>
            You don't have the rights to perform this action.
          </Text>

          <Text>
            If you think that this is an error, please contact the administrator.
          </Text>

          <Link to="/">
            Go back to the Home page
          </Link>
        </Stack>
      </Card>
    </Center>
  );
};

export default ForbiddenPage;
