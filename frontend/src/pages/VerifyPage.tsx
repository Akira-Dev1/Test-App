import { useState } from "react";
import { verifyLogin } from "../auth/authService";
import { useAuth } from "../hooks/useAuth";
import { Button } from "../styles/primitives/Button";
import { Input } from "../styles/primitives/Input";
import { Text } from "../styles/primitives/Text";
import { Center } from "../styles/layout/Center";
import { Stack } from "../styles/primitives/Stack";
import styled from "styled-components";
import { Card } from "../styles/component/Card";
import { theme } from "../styles/theme";

const TextError = styled(Text)`
  color: ${theme.colors.error};
  marginTop: 12px; 
`

const CardVerify = styled(Card)`
  width: 40%;
`

const VerifyPage = () => {
  const [code, setCode] = useState("");
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);

  const { refreshAuth } = useAuth();

  const submitCode = async () => {
    setError(null);
    setLoading(true);

    try {
      await verifyLogin(code);

      await refreshAuth();
    } catch {
      setError("Invalid or outdated code");
    } finally {
      setLoading(false);
    }
  };

  return (
    <Center>
      <CardVerify>
        <Stack gap={16}>  
          <Text variant="h1">Login confirmation</Text>

          <Text>
            Enter the code that was sent to you earlier
          </Text>

          <Input
            type="text"
            placeholder="Confirmation code"
            value={code}
            onChange={(e) => setCode(e.target.value)}
            disabled={loading}
          />

          <Button
            onClick={submitCode}
            disabled={loading || !code}
          >
            {loading ? "Checking..." : "Confirm"}
          </Button>

          {error && (
            <TextError>
              {error}
            </TextError>
          )}
        </Stack>
      </CardVerify>
    </Center>
  );
};

export default VerifyPage;
