import { useState } from "react";
import { verifyLogin } from "../auth/authService";
import { useAuth } from "../hooks/useAuth";
import { Wrapper } from "../components/Wrapper";
import { Title } from "../components/Title";
import { Paragraph } from "../components/Paragraph";



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
    <Wrapper>
      <Title>Login confirmation</Title>

      <Paragraph>
        Enter the code that was sent to you earlier
      </Paragraph>

      <input
        type="text"
        placeholder="Confirmation code"
        value={code}
        onChange={(e) => setCode(e.target.value)}
        disabled={loading}
      />

      <button
        onClick={submitCode}
        disabled={loading || !code}
      >
        {loading ? "Checking..." : "Confirm"}
      </button>

      {error && (
        <Paragraph style={{ color: "red", marginTop: 12 }}>
          {error}
        </Paragraph>
      )}
    </Wrapper>
  );
};

export default VerifyPage;
